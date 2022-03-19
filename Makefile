docker-build:
	docker build -t api-challenge .

docker-run:
	docker run -d -p 16453:16453 --name challengeAPI api-challenge

docker-run-attached:
	docker run -p 16453:16453 --name challengeAPI api-challenge

docker-remove:
	docker rm -fv challengeAPI