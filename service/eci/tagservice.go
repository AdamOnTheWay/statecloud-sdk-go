// Code generated by crafter generator. DO NOT EDIT.

package eci

import (
	"context"
	"fmt"
	"net/http"

	"github.com/state-cloud/client-go/pkg/common/config"
	"github.com/state-cloud/client-go/pkg/openapi"
	"github.com/state-cloud/client-go/pkg/protocol"

	tag "github.com/state-cloud/statecloud-sdk-go/service/eci/types/tag"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
)

type TagClient interface {
	BindTag(context context.Context, req *tag.BindTagRequest, reqOpt ...config.RequestOption) (resp *tag.BindTagResponse, rawResponse *protocol.Response, err error)

	ListTag(context context.Context, req *tag.ListTagRequest, reqOpt ...config.RequestOption) (resp *tag.ListTagResponse, rawResponse *protocol.Response, err error)

	UnbindTag(context context.Context, req *tag.UnbindTagRequest, reqOpt ...config.RequestOption) (resp *tag.UnbindTagResponse, rawResponse *protocol.Response, err error)
}

type tagClient struct {
	client *HttpClient
}

func NewTagClient(hostUrl string, ops ...Option) (TagClient, error) {
	opts := GetOptions(append(ops, WithHostUrl(hostUrl))...)
	cli, err := NewHttpClient(opts)
	if err != nil {
		return nil, err
	}
	return &tagClient{
		client: cli,
	}, nil
}

func (s *tagClient) BindTag(ctx context.Context, req *tag.BindTagRequest, reqOpt ...config.RequestOption) (resp *tag.BindTagResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.OpenapiResponse{}
	openapiResp.ReturnObj = &resp
	ret, err := s.client.R().
		SetContext(ctx).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodPost, "/eci/api/v1/tag/bindTag")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *tagClient) ListTag(ctx context.Context, req *tag.ListTagRequest, reqOpt ...config.RequestOption) (resp *tag.ListTagResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.OpenapiResponse{}
	openapiResp.ReturnObj = &resp
	ret, err := s.client.R().
		SetContext(ctx).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodPost, "/eci/api/v1/tag/listTag")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

func (s *tagClient) UnbindTag(ctx context.Context, req *tag.UnbindTagRequest, reqOpt ...config.RequestOption) (resp *tag.UnbindTagResponse, rawResponse *protocol.Response, err error) {
	openapiResp := &openapi.OpenapiResponse{}
	openapiResp.ReturnObj = &resp
	ret, err := s.client.R().
		SetContext(ctx).
		SetBodyParam(req).
		SetRequestOption(reqOpt...).
		SetResult(openapiResp).
		Execute(http.MethodPost, "/eci/api/v1/tag/unbindTag")
	if err != nil {
		return nil, nil, err
	}

	rawResponse = ret.RawResponse
	return resp, rawResponse, nil
}

var defaultTagClient, _ = NewTagClient(baseDomain)

func ConfigDefaultTagClient(ops ...Option) (err error) {
	defaultTagClient, err = NewTagClient(baseDomain, ops...)
	return
}

func BindTag(context context.Context, req *tag.BindTagRequest, reqOpt ...config.RequestOption) (resp *tag.BindTagResponse, rawResponse *protocol.Response, err error) {
	return defaultTagClient.BindTag(context, req, reqOpt...)
}

func ListTag(context context.Context, req *tag.ListTagRequest, reqOpt ...config.RequestOption) (resp *tag.ListTagResponse, rawResponse *protocol.Response, err error) {
	return defaultTagClient.ListTag(context, req, reqOpt...)
}

func UnbindTag(context context.Context, req *tag.UnbindTagRequest, reqOpt ...config.RequestOption) (resp *tag.UnbindTagResponse, rawResponse *protocol.Response, err error) {
	return defaultTagClient.UnbindTag(context, req, reqOpt...)
}
