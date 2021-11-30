minikube-start:
	minikube start --cpus=4 --memory=8192 --network-plugin=cni --extra-config=apiserver.service-node-port-range=1-65535 --feature-gates=EphemeralContainers=true
minikube-save-image:
	minikube image save $(minikube image ls)
minikube-login:
	docker login ccr.ccs.tencentyun.com --username=100017144341
	docker login hkccr.ccs.tencentyun.com --username=100017144341
init:
	kubectl apply -f ./redis.yaml
image:
	docker pull docker.io/daprio/dashboard:0.9.0
	docker pull docker.io/daprio/dapr:1.5.0
	docker pull docker.io/daprio/daprd:1.5.0
	docker tag docker.io/daprio/dashboard:0.9.0 hkccr.ccs.tencentyun.com/dapr/dashboard:0.9.0
	docker tag docker.io/daprio/dapr:1.5.0 hkccr.ccs.tencentyun.com/dapr/dapr:1.5.0
	docker tag docker.io/daprio/daprd:1.5.0 hkccr.ccs.tencentyun.com/dapr/daprd:1.5.0
	docker pull hkccr.ccs.tencentyun.com/dapr/dashboard:0.9.0
	docker pull hkccr.ccs.tencentyun.com/dapr/dapr:1.5.0
	docker tag hkccr.ccs.tencentyun.com/dapr/dashboard:0.9.0 docker.io/daprio/dashboard:0.9.0 
	docker tag hkccr.ccs.tencentyun.com/dapr/dapr:1.5.0 docker.io/daprio/dapr:1.5.0 
