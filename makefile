minikube-start:
	@minikube start --cpus=2 --memory=4g --cni=flannel

helm-install:
	@helm install app ./chart

apigw:
	@minikube service -n default app-traefik

run-newman:
	@newman run ./collection.postman.json
helm-remove:
	@helm uninstall app
