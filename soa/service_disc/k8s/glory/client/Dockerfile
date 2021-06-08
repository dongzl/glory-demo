FROM golang:latest as build
ENV GO111MODULE on
ADD glory-k8s-client /
ADD config/glory.yaml /config/glory.yaml
ENV GLORY_CONFIG_PATH "/config/glory.yaml"
ENTRYPOINT exec /glory-k8s-client