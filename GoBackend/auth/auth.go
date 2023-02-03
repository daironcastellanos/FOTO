// firebase.go (Golang)
package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App

func init() {
	ctx := context.Background()
	config := &firebase.Config{
		DatabaseURL: "https://<DATABASE_NAME>.firebaseio.com",
	}

	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Error initializing app: %v\n", err)
	}
}

func Auth() *firebase.Auth {
	ctx := context.Background()
	auth, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("Error getting Auth client: %v\n", err)
	}

	return auth
}