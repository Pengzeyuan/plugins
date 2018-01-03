// Code generated with goa v2.0.0-wip, DO NOT EDIT.
//
// storage client
//
// Command:
// $ goa gen goa.design/plugins/cors/examples/cellar/design

package storage

import (
	"context"

	goa "goa.design/goa"
)

// Client is the "storage" service client.
type Client struct {
	ListEndpoint   goa.Endpoint
	ShowEndpoint   goa.Endpoint
	AddEndpoint    goa.Endpoint
	RemoveEndpoint goa.Endpoint
}

// NewClient initializes a "storage" service client given the endpoints.
func NewClient(list, show, add, remove goa.Endpoint) *Client {
	return &Client{
		ListEndpoint:   list,
		ShowEndpoint:   show,
		AddEndpoint:    add,
		RemoveEndpoint: remove,
	}
}

// List calls the "list" endpoint of the "storage" service.
func (c *Client) List(ctx context.Context) (res StoredBottleCollection, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(StoredBottleCollection), nil
}

// Show calls the "show" endpoint of the "storage" service.
// Show can return the following error types:
//	- *NotFound: Bottle not found
//	- error: generic transport error.
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *StoredBottle, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*StoredBottle), nil
}

// Add calls the "add" endpoint of the "storage" service.
func (c *Client) Add(ctx context.Context, p *Bottle) (res string, err error) {
	var ires interface{}
	ires, err = c.AddEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Remove calls the "remove" endpoint of the "storage" service.
// Remove can return the following error types:
//	- *NotFound: Bottle not found
//	- error: generic transport error.
func (c *Client) Remove(ctx context.Context, p *RemovePayload) (err error) {
	_, err = c.RemoveEndpoint(ctx, p)
	return
}
