start:
	cp .env.docker.sample .env
	docker-compose up -d --build

tests:
	touch count.out
	go test -covermode=count -coverprofile=count.out -v ./...
	go tool cover -func=count.out

mock-generate:
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir /app/internal/infra/datastore
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir /app/internal/domain/usecase
