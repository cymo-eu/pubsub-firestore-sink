package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"

	firebase "firebase.google.com/go"
)

func main() {

	// Find and initialize the runtime configuration variables
	env := retrieveEnvironmentRunParams()
	ctx := context.Background()

	// Initialize the pubsub client
	pubSubClient, err := pubsub.NewClient(ctx, env.pubsubProjectId)
	if err != nil {
		log.Fatal(err)
	}
	defer pubSubClient.Close()

	// Initialize the firebase client
	conf := &firebase.Config{ProjectID: env.firestoreProjectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer firestoreClient.Close()

	// Attach to the subscription
	sub := pubSubClient.Subscription(env.subscription)
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		_, _, err := firestoreClient.Collection(env.firestoreCollection).Add(ctx, map[string]interface{}{
			"messageId":       m.ID,
			"attributes":      m.Attributes,
			"orderingKey":     m.OrderingKey,
			"publishTime":     m.PublishTime,
			"payload":         m.Data,
			"deliveryAttempt": m.DeliveryAttempt,
		})
		if err != nil {
			log.Fatalf("Failed adding message: %v", err)
			m.Nack()
		}
		m.Ack()
	})
	if err != nil {
		log.Println(err)
	}
}

func retrieveEnvironmentRunParams() environment {
	projectId, found := os.LookupEnv("PROJECT_ID")
	if !found {
		panic("PROJECT_ID not found. Exiting.")
	}
	subscription, found := os.LookupEnv("SUBSCRIPTION")
	if !found {
		panic("SUBSCRIPTION not found. Exiting.")
	}
	firestoreCollection, found := os.LookupEnv("FIRESTORE_COLLECTION")
	if !found {
		panic("FIRESTORE_COLLECTION not found. Exiting.")
	}
	firestoreProjectId, found := os.LookupEnv("FIRESTORE_PROJECT_ID")
	if !found {
		panic("FIRESTORE_PROJECT_ID not found. Exiting.")
	}
	return environment{pubsubProjectId: projectId, firestoreProjectId: firestoreProjectId, subscription: subscription, firestoreCollection: firestoreCollection}
}

type environment struct {
	pubsubProjectId     string
	firestoreProjectId  string
	subscription        string
	firestoreCollection string
}
