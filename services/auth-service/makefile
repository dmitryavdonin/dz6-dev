
commit: ## build and push container to docker hub| provide tag=(...)
	@docker build -t dmitryavdonin/auth-service:latest . && docker push dmitryavdonin/auth-service:latest
.PHONY: commit



cert: ##generates rsa key
	openssl genrsa -out ./cert/id_rsa 4096 && openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub
.PHONY: cert