package main

import (
	"context"
	"fmt"
	"os"

	"github.com/state-cloud/client-go/pkg/openapi/config"
	"github.com/state-cloud/client-go/pkg/protocol"
	"github.com/state-cloud/statecloud-sdk-go/service/eci"
	cg "github.com/state-cloud/statecloud-sdk-go/service/eci/types/containergroup"
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

func BuildCreateOpsTask() *cg.CreateOpsTaskRequest {
	return &cg.CreateOpsTaskRequest{
		RegionId:         "b342b77ef26b11ecb0ac0242ac110002",
		ContainerGroupId: "eci-xvcv485qgvxnlhzz",
		OpsType:          "coredump",
		OpsValue:         "{\"enable\":true,\"volumeName\":\"test-volume\"}",
	}
}

func CreateOpsTask(ctx context.Context, cli eci.ContainerGroupClient) (*cg.CreateOpsTaskResponse, *protocol.Response, error) {
	req := BuildCreateOpsTask()
	return cli.CreateOpsTask(ctx, req)
}

func main() {
	config := &config.OpenapiConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		return
	}

	ctx := context.Background()
	resp, raw, err := CreateOpsTask(ctx, client.ContainerGroup())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
