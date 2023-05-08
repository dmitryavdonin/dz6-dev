minikube-start:
	@minikube start --cpus=2 --memory=4g --cni=flannel


dependency-build:
	@cd chart && helm dependency build


helm-install:
	@helm install app ./chart

apigw:
	@minikube service -n default app-traefik

run-newman:
	@newman run ./collection.postman.json
helm-remove:
	@helm uninstall app
