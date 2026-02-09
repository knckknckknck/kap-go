package kap_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/knckknckknck/kap-go"
)

func ExampleNewClient() {
	// Production client with API key.
	client := kap.NewClient(os.Getenv("MKK_API_KEY"))
	_ = client

	// Test environment client with basic auth.
	testClient := kap.NewClient("",
		kap.WithBasicAuth("user", "pass"),
		kap.WithBaseURL("https://apigwdev.mkk.com.tr"),
	)
	_ = testClient
}

func ExampleClient_GenerateToken() {
	client := kap.NewClient(os.Getenv("MKK_API_KEY"))

	token, err := client.GenerateToken(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Token length:", len(token))
}

func ExampleClient_Disclosures() {
	client := kap.NewClient("", kap.WithBasicAuth("user", "pass"))

	disclosures, err := client.Disclosures(context.Background(), 1092228, &kap.DisclosureListParams{
		DisclosureClass: "DG",
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range disclosures {
		fmt.Println(d.DisclosureIndex, d.Title)
	}
}

func ExampleClient_DisclosureDetail() {
	client := kap.NewClient("", kap.WithBasicAuth("user", "pass"))

	detail, err := client.DisclosureDetail(context.Background(), 1211180, "data", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(detail.SenderTitle)
}
