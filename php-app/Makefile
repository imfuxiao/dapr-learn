.PHONY: build
build:
	docker build -t ccr.ccs.tencentyun.com/dapr/myphp:v0.1.0 .
	docker push ccr.ccs.tencentyun.com/dapr/myphp:v0.1.0
.PHONY: test
test:
	curl http://192.168.64.9:38080/abc/def\?param\=123
.PHONY: delete
delete:
	kubectl delete -f php-app.yaml
.PHONY: apply
apply:
	kubectl apply -f php-app.yaml
debug:
	kubectl run -it --rm --restart=Never busybox --image=ccr.ccs.tencentyun.com/k8s-fx/busybox:latest sh
	kubectl debug phpapp-7bcc8fb57-xxrg9 -it --image=ccr.ccs.tencentyun.com/k8s-fx/busybox:latest --target=daprd sh
k8s-curl:
	wget phpapp-dapr:3500/v1.0/invoke/phpapp/method/abc?test