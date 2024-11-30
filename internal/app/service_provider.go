package app

import (
	"context"
	"log"

	"github.com/Mobo140/auth/internal/client/cache"
	"github.com/Mobo140/auth/internal/client/cache/redis"
	"github.com/Mobo140/auth/internal/config"
	"github.com/Mobo140/auth/internal/config/env"
	"github.com/Mobo140/auth/internal/repository"
	accessCacheRepository "github.com/Mobo140/auth/internal/repository/access/cache"
	accessDBRepository "github.com/Mobo140/auth/internal/repository/access/db"
	logRepository "github.com/Mobo140/auth/internal/repository/logs"
	userCacheRepository "github.com/Mobo140/auth/internal/repository/user/cache"
	userDBRepository "github.com/Mobo140/auth/internal/repository/user/db"
	"github.com/Mobo140/auth/internal/service"
	accessService "github.com/Mobo140/auth/internal/service/access"
	authService "github.com/Mobo140/auth/internal/service/auth"
	userService "github.com/Mobo140/auth/internal/service/user"
	"github.com/Mobo140/auth/internal/transport/access"
	"github.com/Mobo140/auth/internal/transport/auth"
	"github.com/Mobo140/auth/internal/transport/user"
	"github.com/Mobo140/platform_common/pkg/closer"
	"github.com/Mobo140/platform_common/pkg/db"
	"github.com/Mobo140/platform_common/pkg/db/pg"
	"github.com/Mobo140/platform_common/pkg/db/transaction"
	redigo "github.com/gomodule/redigo/redis"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	redisConfig   config.RedisConfig
	storageConfig config.StorageConfig
	secretConfig  config.SecretConfig

	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	redisPool *redigo.Pool

	redisClient           cache.Client
	dbClient              db.Client
	txManager             db.TxManager
	userDBRepository      repository.UserDBRepository
	userCacheRepository   repository.UserCacheRepository
	accessDBRepository    repository.AccessDBRepository
	accessCacheRepository repository.AccessCacheRepository
	logRepository         repository.LogRepository

	userService   service.UserService
	authService   service.AuthService
	accessService service.AccessService

	userImplementation   *user.Implementation
	authImplementation   *auth.Implementation
	accessImplementation *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *user.Implementation {
	if s.userImplementation == nil {
		s.userImplementation = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImplementation
}

func (s *serviceProvider) AuthImplementation(ctx context.Context) *auth.Implementation {
	if s.authImplementation == nil {
		s.authImplementation = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImplementation
}

func (s *serviceProvider) AccessImplementation(ctx context.Context) *access.Implementation {
	if s.accessImplementation == nil {
		s.accessImplementation = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImplementation
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserDBRepository(ctx),
			s.LogRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.UserDBRepository(ctx),
			s.UserCacheRepository(ctx),
			s.LogRepository(ctx),
			s.TxManager(ctx),
			s.SecretConfig(),
		)
	}

	return s.authService
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.AccessDBRepository(ctx),
			s.AccessCacheRepository(ctx),
			s.LogRepository(ctx),
			s.TxManager(ctx),
			s.SecretConfig(),
		)
	}

	return s.accessService
}

func (s *serviceProvider) UserDBRepository(ctx context.Context) repository.UserDBRepository {
	if s.userDBRepository == nil {
		s.userDBRepository = userDBRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userDBRepository
}

func (s *serviceProvider) UserCacheRepository(_ context.Context) repository.UserCacheRepository {
	if s.userCacheRepository == nil {
		s.userCacheRepository = userCacheRepository.NewRepository(s.RedisClient())
	}

	return s.userCacheRepository
}

func (s *serviceProvider) AccessDBRepository(ctx context.Context) repository.AccessDBRepository {
	if s.accessDBRepository == nil {
		s.accessDBRepository = accessDBRepository.NewRepository(s.DBClient(ctx))
	}

	return s.accessDBRepository
}

func (s *serviceProvider) AccessCacheRepository(_ context.Context) repository.AccessCacheRepository {
	if s.accessCacheRepository == nil {
		s.accessCacheRepository = accessCacheRepository.NewRepository(s.RedisClient())
	}

	return s.accessCacheRepository
}

func (s *serviceProvider) SecretConfig() config.SecretConfig {
	if s.secretConfig == nil {
		cfg, err := env.NewSecretConfig()
		if err != nil {
			log.Fatalf("failed to get secret config: %s", err.Error())
		}

		s.secretConfig = cfg
	}

	return s.secretConfig
}

func (s *serviceProvider) StorageConfig() config.StorageConfig {
	if s.storageConfig == nil {
		cfg, err := env.NewStorageConfig()
		if err != nil {
			log.Fatalf("failed to get storage config: %s", err.Error())
		}

		s.storageConfig = cfg
	}

	return s.storageConfig
}

func (s *serviceProvider) RedisClient() cache.Client {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(s.RedisPool(), s.RedisConfig())
	}

	return s.redisClient
}

func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.RedisConfig().MaxIdle(),
			IdleTimeout: s.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.RedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.NewClient(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.DBClient(ctx))
	}

	return s.logRepository
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %v", err)
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %v", err)
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}
