package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	// "github.com/appleboy/go-fcm"
	"github.com/gorilla/mux"
)

type ChatterBoxUserMessage struct {
	NotificationBody string `json:"notification_body,omitempty"`
	Title            string `json:"notification_title,omitempty"`
	Username         string `json:"username,omitempty"`
	Name             string `json:"name,omitempty"`
	Counter          string `json:"counter,omitempty"`
	DeviceID         string `json:"device_id,omitempty"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sendMessage", sendMessage)
	router.HandleFunc("/", hello)
	err := http.ListenAndServe(":8090", router)
	fmt.Print("ye to chal")
	if err != nil {
		panic("gand fat gyi")
	} else {
		log.Println("things hould work")
		fmt.Print("sab makkhna chal raha h ")
	}

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")

}

func sendMessage(w http.ResponseWriter, r *http.Request) {

	fmt.Print("chal gaya")

	var (
		err  error
		uReq ChatterBoxUserMessage
	)
	err = json.NewDecoder(r.Body).Decode(&uReq)
	if err != nil {
		panic("error in getting request data")

	}

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := uReq.DeviceID

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: uReq.Title,
			Body:  uReq.NotificationBody,
		},
		Data: map[string]string{
			"username": uReq.Username,
			"name":     uReq.Name,
			"counter":  uReq.Counter,
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)

}

// var httpClient = &http.Client{}

// func verifyIdToken(idToken string) (*oauth2.Tokeninfo, error) {
// 	oauth2Service, err := oauth2.New(httpClient)
// 	tokenInfoCall := oauth2Service.Tokeninfo()
// 	tokenInfoCall.IdToken(idToken)
// 	tokenInfo, err := tokenInfoCall.Do()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return tokenInfo, nil
// }
