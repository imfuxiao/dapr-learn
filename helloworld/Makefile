run:
	go run main.go
dapr-run:
	dapr run --app-id hello-dapr --app-port 8088 --dapr-http-port 8089 -- go run main.go
get:
	curl http://127.0.0.1:8088/greeting
dapr-get:
	curl http://127.0.0.1:8089/v1.0/invoke/hello-dapr/method/greeting
