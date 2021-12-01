package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dapr/dapr/pkg/runtime/pubsub"
	"io/ioutil"
	"log"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"
)

// var pubsubName = os.Getenv("DAPR_PUBSUB_NAME")

const (
	pubsubName = "pubsub"
	topic      = "testTopic"
)

func main() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}

	defer client.Close()

	http.HandleFunc("/set", func(rw http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("/set req body:", string(b))

		// 不使用SDK, 直接调用Dapr暴露的API: https://docs.dapr.io/reference/api/pubsub_api/#http-request
		//resp, err := http.Post(`http://localhost:3501/v1.0/publish/pubsub/testTopic`, "application/cloudevents+json", bytes.NewReader(b))
		//
		//defer resp.Body.Close()
		//
		//b, err = ioutil.ReadAll(r.Body)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//fmt.Println("pubsub resp body:", string(b))

		err = client.PublishEvent(context.Background(), pubsubName, topic, b, dapr.PublishEventWithContentType("application/cloudevents+json"))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Fprint(rw, map[string]interface{}{"ok": true})
	})

	// 可选的应用程序（用户代码）路由: https://docs.dapr.io/reference/api/pubsub_api/#optional-application-user-code-routes
	http.HandleFunc("/dapr/subscribe", func(rw http.ResponseWriter, r *http.Request) {

		subs := []pubsub.SubscriptionJSON{
			{
				PubsubName: pubsubName,
				Topic:      topic,
				//Metadata: map[string]string{
				//	"testName": "testValue",
				//},
				Routes: pubsub.RoutesJSON{
					Rules: []*pubsub.RuleJSON{
						{
							Match: `event.type == "widget"`,
							Path:  "widgets",
						},
						{
							Match: `event.type == "gadget"`,
							Path:  "gadgets",
						},
					},
					Default: "widgets",
				},
			},
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(subs)
	})

	http.HandleFunc("/widgets", func(rw http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("widgets req body:", string(b))
	})

	http.HandleFunc("/gadgets", func(rw http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("gadgets req body:", string(b))
	})

	log.Fatalln(http.ListenAndServe(":18080", nil))
}
