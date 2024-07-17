// This file was auto-generated by Fern from our API Definition.

package search

import (
	context "context"
	generatedgo "go-mod-path/generated/go"
	core "go-mod-path/generated/go/core"
	option "go-mod-path/generated/go/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

func (c *Client) PostCrmV3ObjectsContactsSearchDoSearch(
	ctx context.Context,
	request *generatedgo.PublicObjectSearchRequest,
	opts ...option.RequestOption,
) (*generatedgo.CollectionResponseWithTotalSimplePublicObjectForwardPaging, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.hubapi.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/crm/v3/objects/contacts/search"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *generatedgo.CollectionResponseWithTotalSimplePublicObjectForwardPaging
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
			Response:    &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
