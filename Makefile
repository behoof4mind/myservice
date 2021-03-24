APP_NAME = $(name)
APP_VERSION = $(version)

DOCKER_REGISTRY_URL = behoof4mind

build:
	docker build -t ${DOCKER_REGISTRY_URL}/${APP_NAME}:${APP_VERSION} .

test:
	echo "Here we should perform some tests"

push:
	docker push ${DOCKER_REGISTRY_URL}/$(APP_NAME):$(APP_VERSION)

all: build test push
