package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"), config.WithSharedConfigProfile("kntks"))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	// client, ok := cfg.HTTPClient.(*awshttp.BuildableClient)
	// fmt.Printf("%v  %v\n", client, ok)
	client := iam.NewFromConfig(cfg)
	outputs, err := client.ListRoles(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, role := range outputs.Roles {
		fmt.Printf("%s\n", *role.RoleName)
	}
}
