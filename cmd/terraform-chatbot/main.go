package main

import (
	"context"
	"fmt"
	"log"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	dialogflowpb "cloud.google.com/go/dialogflow/apiv2/dialogflowpb"
)

func main() {
	// Set up Dialogflow client
	ctx := context.Background()
	client, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Dialogflow client: %v", err)
	}
	defer client.Close()

	// Set up session parameters
	sessionID := "your-session-id"
	projectID := "your-project-id"
	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)

	// Set up the query parameters
	queryText := "I need a new server in AWS"
	queryInput := &dialogflowpb.QueryInput{
		Input: &dialogflowpb.QueryInput_Text{
			Text: &dialogflowpb.TextInput{
				Text:         queryText,
				LanguageCode: "en-US",
			},
		},
	}

	// Send the query to Dialogflow
	response, err := client.DetectIntent(ctx, &dialogflowpb.DetectIntentRequest{
		Session:    sessionPath,
		QueryInput: queryInput,
	})
	if err != nil {
		log.Fatalf("Failed to detect intent: %v", err)
	}

	// Parse the response
	queryResult := response.GetQueryResult()
	intentDisplayName := queryResult.Intent.DisplayName
	parameters := queryResult.Parameters.Fields

	// Deploy infrastructure using Terraform based on the user's input
	if intentDisplayName == "deploy-infra" {
		_ = parameters["server_type"].GetStringValue()
		_ = parameters["cloud_provider"].GetStringValue()

		// Use Terraform to deploy infrastructure based on the user's input
		// ...
	}
}
