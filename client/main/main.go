package main

import (
	"context"
	"io"
	"log"
	"time"

	openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
)

// Helper functions to simplify pointer initialization

func Bool(b bool) *bool {
	return &b
}

func String(s string) *string {
	return &s
}

func StringToTime(s string) *time.Time {
	t, _ := time.Parse(time.RFC3339, s)
	return &t
}

// CreateSubscription creates a subscription using the provided data
// func Subscription(ctx context.Context, c openapi.APIClient) {
// 	data := openapi.SdmSubscription{
// 		NfInstanceId:          "3fa85f64-5717-4562-b3fc-2c963f66afa6",
// 		ImplicitUnsubscribe:   Bool(true),
// 		Expires:               StringToTime("2024-01-30T09:07:01.162Z"),
// 		CallbackReference:     "dsjvb",
// 		MonitoredResourceUris: []string{"string"},

// 		Dnn:            String("string"),
// 		SubscriptionId: String("string"),
// 		PlmnId: &openapi.PlmnId{
// 			Mcc: "326",
// 			Mnc: "495",
// 		},
// 		ImmediateReport:   Bool(false),
// 		SupportedFeatures: String("a0fcB2fdF28Bbc215Bf6cfEC4966Bed1A28b4dAc253c6aE4b7dEBf0Eab5cFEC6fb6F4Ec718414c"),

// 		NfChangeFilter:             Bool(false),
// 		UniqueSubscription:         Bool(true),
// 		ResetIds:                   []string{"string"},
// 		DataRestorationCallbackUri: String("string"),
// 	}

// 	r := c.SubscriptionCreationAPI.Subscribe(ctx, "832gf")
// 	rs := r.SdmSubscription(data)
// 	_, resp, err := rs.Execute()
// 	if err != nil {
// 		log.Print(err)
// 	} else {
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Fatal("Error reading response body:", err)
// 		}

// 		log.Printf("Response Status: %s", resp.Status)
// 		log.Printf("Response Headers: %v", resp.Header)
// 		log.Printf("Response Body: %s", body)
// 	}
// }

// func Unsubscribe(ctx context.Context, c openapi.APIClient) {
// 	r := c.SubscriptionDeletionAPI.Unsubscribe(ctx, "egvas", "fdhg")
// 	resp, err := r.Execute()
// 	if err != nil {
// 		log.Print(err)
// 	} else {
// 		log.Printf("Response Status: %s", resp.Status)
// 	}
// }

func GetSession(ctx context.Context, c openapi.APIClient) {
	r := c.SessionManagementSubscriptionDataRetrievalAPI.GetSmData(ctx, "erb32")
	_, resp, err := r.Execute()
	if err != nil {
		log.Print(err)
	} else {
		// log.Print(data)
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading response body:", err)
		}
		log.Printf("Response Status: %s", resp.Status)
		log.Printf("Response Body: %s", body)
	}
}

// func Modify(ctx context.Context, c openapi.APIClient) {
// 	r := c.SubscriptionModificationAPI.Modify(ctx, "vcdsv", "dvsv")
// 	data := openapi.SdmSubsModification{
// 		Expires:               StringToTime("2024-02-07T17:38:07.814Z"),
// 		MonitoredResourceUris: []string{"string"},
// 	}
// 	rs := r.SdmSubsModification(data)
// 	_, resp, err := rs.Execute()
// 	if err != nil {
// 		log.Print(err)
// 	} else {
// 		// body, err := io.ReadAll(resp.Body)
// 		// if err != nil {
// 		// 	log.Fatal("Error reading response body:", err)
// 		// }

// 		log.Printf("Response Status: %s", resp.Status)
// 		log.Printf("Response Headers: %v", resp.Header)
// 		// log.Printf("Response Body: %s", body)
// 	}
// }

func main() {
	ctx := context.Background()
	config := openapi.NewConfiguration()
	client := openapi.NewAPIClient(config)
	// Subscription(ctx, *client)
	// Unsubscribe(ctx, *client)
	// Modify(ctx, *client)
	GetSession(ctx, *client)
}
