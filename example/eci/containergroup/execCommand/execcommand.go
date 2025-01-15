package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"

	cli "github.com/state-cloud/client-go/pkg/client"
	"github.com/state-cloud/client-go/pkg/openapi/config"
	"github.com/state-cloud/statecloud-sdk-go/service/eci"
	containergroup "github.com/state-cloud/statecloud-sdk-go/service/eci/types/containergroup"
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
		AccessKey: accessKey,
		SecretKey: secretKey,
	}

	options := []eci.Option{
		eci.WithClientConfig(config),
		eci.WithClientOption(cli.WithTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		})),
	}
	client, err := eci.NewClientSet(baseDomain, options...)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()
	req := &containergroup.ExecCommandRequest{
		// Fill in the request parameters
		ContainerGroupId: "eci-xxx",
		ContainerName:    "container-x",
		Command:          []string{"ls", "/"},
		TTY:              false,
		Stdin:            false,
		Sync:             true,
	}
	resp, raw, err := client.ContainerGroup().ExecCommand(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("raw: %v\nresp: %v\n", string(raw.Body()), resp)
}
