docker-build:
	docker build -t http-echo .

docker-run:
	docker run --rm http-echo

test:
	curl 127.0.0.1:8000
