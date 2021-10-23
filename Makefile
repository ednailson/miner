start:
	cp .env.docker.sample .env
	docker-compose up -d --build