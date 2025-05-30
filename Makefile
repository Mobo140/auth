include env/local.env

LOCAL_BIN := $(CURDIR)/bin

# Setup and run the project: install dependencies, generate code, start services, then run the server
setup: install-deps generate up
	go run cmd/grpc-server/main.go --config-path=env/local.env -l=debug

# Start all services in detached mode using docker-compose
up:
	docker-compose up -d

# Stop all services
down:
	docker-compose down

# Install required CLI tools for protobuf, migrations, grpc-gateway, validation, and linting
install-deps:
	@test -f $(LOCAL_BIN)/protoc-gen-go || GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	@test -f $(LOCAL_BIN)/protoc-gen-go-grpc || GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@test -f $(LOCAL_BIN)/goose || GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	@test -f $(LOCAL_BIN)/protoc-gen-validate || GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.0.4
	@test -f $(LOCAL_BIN)/protoc-gen-grpc-gateway || GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.20.0
	@test -f $(LOCAL_BIN)/protoc-gen-openapiv2 || GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.20.0
	@test -f $(LOCAL_BIN)/statik || GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7
	@test -f $(LOCAL_BIN)/golangci-lint || GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

# Generate protobuf, gRPC, validation, grpc-gateway, and swagger code for user, auth and access APIs
generate:
	mkdir -p pkg/swagger
	make generate-user-api
	make generate-auth-api
	make generate-access-api
	$(LOCAL_BIN)/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

# Generate user API protobuf code and gateway
generate-user-api:
	mkdir -p pkg/user_v1
	protoc --proto_path=api/user_v1 --proto_path=vendor.protogen \
		--go_out=pkg/user_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/user_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out=lang=go:pkg/user_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		--grpc-gateway_out=pkg/user_v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
		--openapiv2_out=allow_merge=true,merge_file_name=user_api:pkg/swagger \
		--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
		api/user_v1/user.proto

# Generate auth API protobuf code and gateway
generate-auth-api:
	mkdir -p pkg/auth_v1
	protoc --proto_path=api/auth_v1 --proto_path=vendor.protogen \
		--go_out=pkg/auth_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/auth_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out=lang=go:pkg/auth_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		--grpc-gateway_out=pkg/auth_v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
		--openapiv2_out=allow_merge=true,merge_file_name=auth_api:pkg/swagger/ \
		--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
		api/auth_v1/auth.proto

# Generate access API protobuf code and gateway
generate-access-api:
	mkdir -p pkg/access_v1
	protoc --proto_path=api/access_v1 --proto_path=vendor.protogen \
		--go_out=pkg/access_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/access_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out=lang=go:pkg/access_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		--grpc-gateway_out=pkg/access_v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
		--openapiv2_out=allow_merge=true,merge_file_name=access_api:pkg/swagger \
		--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
		api/access_v1/access.proto

# Clone third-party protobuf dependencies if missing
vendor-proto:
	@if [ ! -d vendor.protogen/validate ]; then \
		mkdir -p vendor.protogen/validate && \
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate && \
		mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate && \
		rm -rf vendor.protogen/protoc-gen-validate ; \
	fi
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis && \
		mkdir -p vendor.protogen/google && \
		mv vendor.protogen/googleapis/google/api vendor.protogen/google && \
		rm -rf vendor.protogen/googleapis ; \
	fi
	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
		mkdir -p vendor.protogen/protoc-gen-openapiv2/options && \
		git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 && \
		mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options && \
		rm -rf vendor.protogen/openapiv2 ; \
	fi

# Generate self-signed TLS certificates for local development/testing
gen-cert:
	mkdir -p secure
	openssl genrsa -out secure/ca.key 4096
	openssl req -new -x509 -key secure/ca.key -sha256 -subj "/C=US/ST=NJ/O=CA, Inc." -days 365 -out secure/ca.cert
	openssl genrsa -out secure/service.key 4096
	openssl req -new -key secure/service.key -out secure/service.csr -config certificate.conf
	openssl x509 -req -in secure/service.csr -CA secure/ca.cert -CAkey secure/ca.key -CAcreateserial \
		-out secure/service.pem -days 365 -sha256 -extfile certificate.conf -extensions req_ext

# Run golangci-lint using pipeline config
lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

# Run tests with coverage and retry 5 times
test:
	go clean -testcache
	go test ./... -covermode=count -coverpkg=github.com/Mobo140/auth/internal/service/...,github.com/Mobo140/auth/internal/transport/... -count=5

# Run tests with coverage report, open HTML report and show summary
test-coverage:
	go clean -testcache
	go test ./... -coverprofile=coverage.tmp.out -covermode=count -coverpkg=github.com/Mobo140/auth/internal/transport/user/...,github.com/Mobo140/auth/internal/service/user/... -count=5
	grep -v 'mocks\|config' coverage.tmp.out > coverage.out
	rm coverage.tmp.out
	go tool cover -html=coverage.out
	go tool cover -func=coverage.out | grep "total"
	grep -sqFx "/coverage.out" .gitignore || echo "coverage_out" >> .gitignore

# Run gRPC load test with TLS certs and rate limiting
grpc-load-test:
	ghz --proto api/test_user_v1/user_test.proto \
		--call user_v1.UserV1.Get \
		--data '{"id": 1}' \
		--rps 100 --total 3000 \
		--cacert ca.cert --cert service.pem --key service.key \
		localhost:${GRPC_PORT}

# Run gRPC load test expecting errors (e.g. invalid user id) with TLS certs
grpc-error-load-test:
	ghz --proto api/test_user_v1/user_test.proto \
		--call user_v1.UserV1.Get \
		--data '{"id": 100}' \
		--rps 100 --total 3000 \
		--cacert ca.cert --cert service.pem --key service.key \
		localhost:${GRPC_PORT}

# Format all Go files using goimports
format:
	find . -name '*.go' -exec goimports -w {} +
