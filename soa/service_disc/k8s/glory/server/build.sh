export GOOS="linux"
export GOARCH="amd64"
go build -o glory-k8s-server .
docker build --no-cache -t glory-k8s-server-image  .
rm ./glory-k8s-server
kubectl create namespace glory
kubectl delete -f ./server.yaml
kubectl create -f ./server.yaml
