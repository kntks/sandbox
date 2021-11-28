package gcp

import (
	"context"
	"fmt"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"google.golang.org/api/iterator"

	resourcemanagerpb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"
)

// https://pkg.go.dev/cloud.google.com/go/resourcemanager@v1.0.0/apiv3#FoldersClient.ListFolders
func Main(orgID string) {
	ctx := context.Background()
	c, err := resourcemanager.NewFoldersClient(ctx)
	if err != nil {
		panic(err)
	}
	defer c.Close()
	req := &resourcemanagerpb.ListFoldersRequest{
		Parent: fmt.Sprintf("organizations/%s", orgID),
	}
	it := c.ListFolders(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", resp)
	}
}
