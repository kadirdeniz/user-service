run:
	bash scripts/run.sh

tests:
	bash scripts/test.sh

docker-run-mongo:
	bash scripts/docker-run-mongo.sh

docker-run-redis:
	bash scripts/docker-run-redis.sh

docker-build-user-service:
	bash scripts/docker-build-user-service.sh

docker-run-user-service:
	bash scripts/docker-run-user-service.sh