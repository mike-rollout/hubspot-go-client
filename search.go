// This file was auto-generated by Fern from our API Definition.

package api

type PublicObjectSearchRequest struct {
	Query        *string        `json:"query,omitempty" url:"-"`
	Limit        int            `json:"limit" url:"-"`
	After        string         `json:"after" url:"-"`
	Sorts        []string       `json:"sorts,omitempty" url:"-"`
	Properties   []string       `json:"properties,omitempty" url:"-"`
	FilterGroups []*FilterGroup `json:"filterGroups,omitempty" url:"-"`
}
