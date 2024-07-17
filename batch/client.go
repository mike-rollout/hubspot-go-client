// This file was auto-generated by Fern from our API Definition.

package batch

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

func (c *Client) PostCrmV3ObjectsContactsBatchReadRead(
	ctx context.Context,
	request *generatedgo.BatchReadInputSimplePublicObjectId,
	opts ...option.RequestOption,
) (*generatedgo.BatchResponseSimplePublicObject, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.hubapi.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/crm/v3/objects/contacts/batch/read"

	queryParams, err := core.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *generatedgo.BatchResponseSimplePublicObject
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

func (c *Client) PostCrmV3ObjectsContactsBatchArchiveArchive(
	ctx context.Context,
	request *generatedgo.BatchInputSimplePublicObjectId,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.hubapi.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/crm/v3/objects/contacts/batch/archive"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:         endpointURL,
			Method:      http.MethodPost,
			MaxAttempts: options.MaxAttempts,
			Headers:     headers,
			Client:      options.HTTPClient,
			Request:     request,
		},
	); err != nil {
		return err
	}
	return nil
}

func (c *Client) PostCrmV3ObjectsContactsBatchCreateCreate(
	ctx context.Context,
	request *generatedgo.BatchInputSimplePublicObjectInputForCreate,
	opts ...option.RequestOption,
) (*generatedgo.BatchResponseSimplePublicObject, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.hubapi.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/crm/v3/objects/contacts/batch/create"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *generatedgo.BatchResponseSimplePublicObject
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

func (c *Client) PostCrmV3ObjectsContactsBatchUpdateUpdate(
	ctx context.Context,
	request *generatedgo.BatchInputSimplePublicObjectBatchInput,
	opts ...option.RequestOption,
) (*generatedgo.BatchResponseSimplePublicObject, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.hubapi.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/crm/v3/objects/contacts/batch/update"

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *generatedgo.BatchResponseSimplePublicObject
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