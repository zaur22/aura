tests: tests-unit tests-integrations test-build

tests-unit:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down --volumes
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit unit

tests-integrations:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down --volumes
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit postgres local-migrations
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit postgres integration

test-build:
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down --volumes
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit build