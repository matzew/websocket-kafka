# Go WebSocket 2 Kafka Bridge

## build

run `make build_linux`

## run

Execute the binary (./ws-kafka)

> NOTE: export `KAFKA_BOOTSTRAP_SERVERS`, `WEBSOCKET_SERVER` and `KAFKA_TOPIC` to be able to use the client

## deploy to k8s to access internal kafka clustem

```
user=mysuser
DOCKER_USER=${user} make -e docker_build
docker push ${user}/ws-kafka

kubectl run --image ${user}/ws-kafka \
            --generator=run-pod/v1
            --env=WEBSOCKET_SERVER=wss://api.usb.urbanobservatory.ac.uk/stream \
            --env=KAFKA_TOPIC=demo \
            --env=KAFKA_HOST=my-cluster-kafka-bootstrap.kafka \
            ws-kafka-bridge
```
