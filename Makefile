include local.env

LOCAL_BIN:=$(CURDIR)/bin


#LOCAL_MIGRATION_DSN=${MIGRATION_DSN_MAKE}


install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path api/user_v1 \
	--go_out=pkg/user_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

# Для локальной накатки миграций
# local-migration-status:
# 	bin/goose -dir ${MIGRATION_DIR}	postgres ${LOCAL_MIGRATION_DSN} status -v

# local-migration-up:
# 	bin/goose -dir ${MIGRATION_DIR}	postgres ${LOCAL_MIGRATION_DSN} up -v

# local-migration-down:
# 	bin/goose -dir ${MIGRATION_DIR}	postgres ${LOCAL_MIGRATION_DSN} down -v

#bin/goose -dir ./migrations create auth  sql - создание файлов миграций

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

test: 
	go clean -testcache
	go test ./... -covermode count -coverpkg=github.com/Mobo140/microservices/auth/internal/service/..., github.com/Mobo140/microservices/auth/internal/transport/... -count 5

test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode count -coverpkg=github.com/Mobo140/microservices/auth/internal/transport/user/...,github.com/Mobo140/microservices/auth/internal/service/user/... -count 5
	grep -v 'mocks\|config' coverage.tmp.out > coverage.out 
	rm coverage.tmp.out
	go tool cover -html=coverage.out
	go tool cover -func=./coverage.out | grep "total"
	grep -sqFx "/coverage.out" .gitignore || echo "coverage_out" >> .gitignore