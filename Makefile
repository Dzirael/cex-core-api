PHONY: gen


lint:
	gofumpt -l -w .

gen:
	go generate ./...

test-cover:
	go test -coverprofile=coverage.out ./app/...
	go tool cover -html=coverage.out

test-repo-%:
	go test -v ./app/internal/storages/postgres/repository/$*