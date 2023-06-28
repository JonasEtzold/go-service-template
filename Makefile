docs:
	swag init --dir cmd/api --parseDependency --output docs

build:
	go build -o ./out/app cmd/api/main.go

run:
	go run cmd/api/main.go

build-docker: build
	docker build . -t <%= serviceName %>

run-docker: build-docker
	docker run -p 3000:3000 gostarter
