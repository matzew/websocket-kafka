BINARY ?= ws-kafka
APP_NAME = ws-kafka
DOCKER_USER=matzew
DOCKER_LATEST_TAG = docker.io/$(DOCKER_USER)/$(APP_NAME):latest
DOCKER_MASTER_TAG = docker.io/$(DOCKER_USER)/$(APP_NAME):master
RELEASE_TAG ?= $(CIRCLE_TAG)
DOCKER_RELEASE_TAG = matzew/$(APP_NAME):$(RELEASE_TAG)

.PHONY: generate
generate:
	./scripts/generate.sh

.PHONY: setup
setup:
	glide install --strip-vendor

.PHONY: build_linux
build_linux:
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY) ./cmd/bridge/main.go

.PHONY: docker_build
docker_build:
	docker build -t $(DOCKER_LATEST_TAG) -f Dockerfile .

.PHONY: docker_build_release
docker_build_release:
	docker build -t $(DOCKER_LATEST_TAG) -t $(DOCKER_RELEASE_TAG) -f Dockerfile .

.PHONY: docker_build_master
docker_build_master:
	docker build -t $(DOCKER_MASTER_TAG) -f Dockerfile .

.PHONY: docker_push_latest
docker_push_latest:
	@docker login -u $(DOCKERHUB_USERNAME) -p $(DOCKERHUB_PASSWORD)
	docker push $(DOCKER_LATEST_TAG)

.PHONY: docker_push_master
docker_push_master:
	@docker login -u $(DOCKERHUB_USERNAME) -p $(DOCKERHUB_PASSWORD)
	docker push $(DOCKER_MASTER_TAG)

.PHONY: docker_push_release
docker_push_release:
	@docker login --username $(DOCKERHUB_USERNAME) --password $(DOCKERHUB_PASSWORD)
	docker push $(DOCKER_LATEST_TAG)
	docker push $(DOCKER_RELEASE_TAG)
