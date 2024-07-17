// This file was auto-generated by Fern from our API Definition.

package client

import (
	basic "go-mod-path/generated/go/basic"
	batch "go-mod-path/generated/go/batch"
	core "go-mod-path/generated/go/core"
	gdpr "go-mod-path/generated/go/gdpr"
	option "go-mod-path/generated/go/option"
	publicobject "go-mod-path/generated/go/publicobject"
	search "go-mod-path/generated/go/search"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Batch        *batch.Client
	Basic        *basic.Client
	PublicObject *publicobject.Client
	Gdpr         *gdpr.Client
	Search       *search.Client
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
		header:       options.ToHeader(),
		Batch:        batch.NewClient(opts...),
		Basic:        basic.NewClient(opts...),
		PublicObject: publicobject.NewClient(opts...),
		Gdpr:         gdpr.NewClient(opts...),
		Search:       search.NewClient(opts...),
	}
}