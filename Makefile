
.PHONY: build run

CONTAINER_NAME := testing
IMAGE_NAME := notify-container


build:
	docker buildx build -t notify-container .

kill:
	-docker stop $(CONTAINER_NAME)
	-docker rm $(CONTAINER_NAME)
	-docker rmi $(IMAGE_NAME)

run:
	docker run --name $(CONTAINER_NAME) --env-file .env $(IMAGE_NAME)
