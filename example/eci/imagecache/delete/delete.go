package main

import (
	"context"
	"fmt"
	"os"

	"github.com/state-cloud/client-go/pkg/openapi/config"
	"github.com/state-cloud/statecloud-sdk-go/service/eci"
	"github.com/state-cloud/statecloud-sdk-go/service/eci/types/imagecache"
)

var (
	accessKey  = ""
	secretKey  = ""
	baseDomain = "https://eci-global.ctapi.ctyun.cn"
)

func init() {
	accessKey = os.Getenv("CTAPI_AK")
	secretKey = os.Getenv("CTAPI_SK")
	domain := os.Getenv("CTAPI_ECI_DOMAIN")
	if domain != "" {
		baseDomain = domain
	}
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: "10df4de7266e49fdbe7958e7c4920b56",
		SecretKey: "c7f655a8952b49cba5edba79d099ef77",
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		return
	}

	ctx := context.Background()
	req := &imagecache.DeleteImageCacheRequest{
		ImageCacheId: "imc-w8rlb89a2ahuflqa",
		RegionId:     "bb9fdb42056f11eda1610242ac110002",
	}

	resp, raw, err := client.ImageCache().DeleteImageCache(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
