// Code generated by go-swagger; DO NOT EDIT.

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new acl API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for acl API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateACL(params *CreateACLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateACLCreated, *CreateACLAccepted, error)

	DeleteACL(params *DeleteACLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteACLAccepted, *DeleteACLNoContent, error)

	GetACL(params *GetACLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetACLOK, error)

	GetAcls(params *GetAclsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAclsOK, error)

	ReplaceACL(params *ReplaceACLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceACLOK, *ReplaceACLAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateACL adds a new ACL line

  Adds a new ACL line of the specified type in the specified parent.
*/
func (a *Client) CreateACL(params *CreateACLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateACLCreated, *CreateACLAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateACLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createAcl",
		Method:             "POST",
		PathPattern:        "/services/haproxy/configuration/acls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateACLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateACLCreated:
		return value, nil, nil
	case *CreateACLAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateACLDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteACL deletes a ACL line

  Deletes a ACL line configuration by it's index from the specified parent.
*/
func (a *Client) DeleteACL(params *DeleteACLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteACLAccepted, *DeleteACLNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteACLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteAcl",
		Method:             "DELETE",
		PathPattern:        "/services/haproxy/configuration/acls/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteACLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteACLAccepted:
		return value, nil, nil
	case *DeleteACLNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteACLDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetACL returns one ACL line

  Returns one ACL line configuration by it's index in the specified parent.
*/
func (a *Client) GetACL(params *GetACLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetACLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetACLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAcl",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/acls/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetACLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetACLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetACLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetAcls returns an array of all ACL lines

  Returns all ACL lines that are configured in specified parent.
*/
func (a *Client) GetAcls(params *GetAclsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAclsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAclsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAcls",
		Method:             "GET",
		PathPattern:        "/services/haproxy/configuration/acls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAclsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAclsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetAclsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ReplaceACL replaces a ACL line

  Replaces a ACL line configuration by it's index in the specified parent.
*/
func (a *Client) ReplaceACL(params *ReplaceACLParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReplaceACLOK, *ReplaceACLAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReplaceACLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "replaceAcl",
		Method:             "PUT",
		PathPattern:        "/services/haproxy/configuration/acls/{index}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReplaceACLReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ReplaceACLOK:
		return value, nil, nil
	case *ReplaceACLAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReplaceACLDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
