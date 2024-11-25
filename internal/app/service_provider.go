package app

import (
	"context"
	"log"

	"github.com/Mobo140/microservices/auth/internal/client/cache"
	"github.com/Mobo140/microservices/auth/internal/client/cache/redis"
	"github.com/Mobo140/microservices/auth/internal/client/db"
	"github.com/Mobo140/microservices/auth/internal/client/db/pg"
	"github.com/Mobo140/microservices/auth/internal/client/db/transaction"
	"github.com/Mobo140/microservices/auth/internal/closer"
	"github.com/Mobo140/microservices/auth/internal/config"
	"github.com/Mobo140/microservices/auth/internal/config/env"
	"github.com/Mobo140/microservices/auth/internal/repository"
	logRepository "github.com/Mobo140/microservices/auth/internal/repository/logs"
	userCacheRepository "github.com/Mobo140/microservices/auth/internal/repository/user/cache"
	userDBRepository "github.com/Mobo140/microservices/auth/internal/repository/user/db"
	"github.com/Mobo140/microservices/auth/internal/service"
	authService "github.com/Mobo140/microservices/auth/internal/service/auth"
	userService "github.com/Mobo140/microservices/auth/internal/service/user"
	"github.com/Mobo140/microservices/auth/internal/transport/user"
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

	redisClient         cache.Client
	dbClient            db.Client
	txManager           db.TxManager
	userDBRepository    repository.UserDBRepository
	userCacheRepository repository.UserCacheRepository
	logRepository       repository.LogRepository

	userService service.UserService
	authService service.AuthService

	userImplementation *user.Implementation
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
			s.SecretConfig(),
		)
	}

	return s.authService
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
