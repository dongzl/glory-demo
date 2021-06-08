export GOOS="linux"
export GOARCH="amd64"
go build -o glory-k8s-client .
docker build --no-cache -t glory-k8s-client-image  .
rm ./glory-k8s-client
kubectl create namespace glory
kubectl delete -f ./client.yaml
kubectl create -f ./client.yaml
