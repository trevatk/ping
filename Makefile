
build:
	docker build -t trevatk/ping:client-0.1.0 -f ./docker/client.Dockerfile .
	docker build -t trevatk/ping:server-0.1.0 -f ./docker/server.Dockerfile .