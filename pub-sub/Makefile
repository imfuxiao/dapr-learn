init:
	kubectl apply -f ./config/redis.yaml
	kubectl apply -f ./config/pubsub.yaml
local-pub:
	# 外部声明
	#dapr run --log-level debug --app-id pub --dapr-http-port 3501 --app-port 18080 --config config.yaml --components-path ./config -- go run pub/main.go
	# 代码声明
	dapr run --log-level debug --app-id pub --dapr-http-port 3501 --app-port 18080 --config config.yaml   -- go run pub/main.go
local-sub:
	dapr run --app-id sub --dapr-http-port 3502 --app-port 18888 --config config.yaml --components-path ./config -- go run sub/main.go
local-sub2:
	dapr run --app-id sub2 --dapr-http-port 3503 --app-port 18889 --config config.yaml --components-path ./config -- go run sub2/main.go
docker-build:
	docker build -t ccr.ccs.tencentyun.com/dapr/pubsubapp:v0.1.0 .