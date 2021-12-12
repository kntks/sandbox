package gcp

import (
	"context"
	"fmt"
	"io"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	"go.uber.org/zap/buffer"
	"google.golang.org/api/iterator"
	securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"
)

func listAllAssets(w io.Writer, orgID string) error {
	// orgID := "12321311"
	// Instantiate a context and a security service client to make API calls.
	ctx := context.Background()
	client, err := securitycenter.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("securitycenter.NewClient: %v", err)
	}
	defer client.Close() // Closing the client safely cleans up background resources.

	req := &securitycenterpb.ListAssetsRequest{
		Parent: fmt.Sprintf("organizations/%s", orgID),
	}

	assetsFound := 0
	it := client.ListAssets(ctx, req)
	for {
		result, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("ListAssets: %v", err)
		}
		asset := result.Asset
		properties := asset.SecurityCenterProperties
		fmt.Fprintf(w, "Asset Name: %s,", asset.Name)
		fmt.Fprintf(w, "Resource Name %s,", properties.ResourceName)
		fmt.Fprintf(w, "Resource Type %s\n", properties.ResourceType)
		fmt.Println(properties)
		assetsFound++
	}
	return nil
}

func MainScc(orgID string) {
	var buf buffer.Buffer
	if err := listAllAssets(&buf, orgID); err != nil {
		panic(err)
	}
	// fmt.Println(buf.String())
}
