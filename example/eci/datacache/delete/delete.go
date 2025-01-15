package main

import (
	"context"
	"fmt"
	"os"

	"github.com/state-cloud/client-go/pkg/openapi/config"
	"github.com/state-cloud/client-go/pkg/protocol"
	"github.com/state-cloud/statecloud-sdk-go/service/eci"
	dc "github.com/state-cloud/statecloud-sdk-go/service/eci/types/datacache"
)

func BuildDeleteDataCacheRequest() *dc.DeleteDataCacheRequest {
	return &dc.DeleteDataCacheRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		DataCacheId: "edc-3oytf3ix1hbph9zx",
	}
}

func DeleteDataCache(ctx context.Context, cli eci.DataCacheClient) (*dc.DeleteDataCacheResponse, *protocol.Response, error) {
	req := BuildDeleteDataCacheRequest()
	return cli.DeleteDataCache(ctx, req)
}

func main() {
	baseDomain := "https://eci-global.ctapi.ctyun.cn"
	config := &config.OpenapiConfig{
		AccessKey: os.Getenv("CTAPI_AK"),
		SecretKey: os.Getenv("CTAPI_SK"),
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		return
	}

	ctx := context.Background()
	resp, raw, err := DeleteDataCache(ctx, client.DataCache())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
