PROJECT?=github.com/iceoskara/goworkshop
APP?=gophercon
PORT?=8000
INTERNAL_PORT?=3001

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

CONTAINER_IMAGE?=iceoskara/goworkshop/${APP}

GOOS?=linux
GOARCH?=amd64


clean:
	rm -f ./bin/${GOOS}-${GOARCH}/${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	go build \
	-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} \
		-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ./bin/${GOOS}-${GOARCH}/${APP} ${PROJECT}/cmd/gophercon 

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: container
	docker stop $(CONTAINER_IMAGE):$(RELEASE) || true && docker rm $(CONTAINER_IMAGE):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} -p ${INTERNAL_PORT}:${INTERNAL_PORT} --rm \
		-e "SERVICE_PORT=${PORT}" -e "INTERNAL_PORT=${INTERNAL_PORT}" \
		$(CONTAINER_IMAGE):$(RELEASE)

test:
	go test -race ./...

push: build
	docker push $(CONTAINER_IMAGE):$(RELEASE)

deploy: push
	helm upgrade ${CONTAINER_NAME} -f charts/${VALUES}.yaml charts --kube-context ${KUBE_CONTEXT} --namespace ${NAMESPACE} --version=${RELEASE} -i --wait