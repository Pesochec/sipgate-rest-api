// Code generated by go-swagger; DO NOT EDIT.

package autorecording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new autorecording API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for autorecording API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteGreeting(params *DeleteGreetingParams, authInfo runtime.ClientAuthInfoWriter) error

	GetGreeting(params *GetGreetingParams, authInfo runtime.ClientAuthInfoWriter) (*GetGreetingOK, error)

	GetSetting(params *GetSettingParams, authInfo runtime.ClientAuthInfoWriter) (*GetSettingOK, error)

	UpdateSetting(params *UpdateSettingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSettingOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteGreeting deletes the greeting used for automated call recordings
*/
func (a *Client) DeleteGreeting(params *DeleteGreetingParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGreetingParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGreeting",
		Method:             "DELETE",
		PathPattern:        "/autorecordings/greetings/{greetingId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteGreetingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  GetGreeting fetches the current automated call recording greeting
*/
func (a *Client) GetGreeting(params *GetGreetingParams, authInfo runtime.ClientAuthInfoWriter) (*GetGreetingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGreetingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGreeting",
		Method:             "GET",
		PathPattern:        "/autorecordings/greetings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGreetingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGreetingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGreeting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSetting gets the recording setting of a device
*/
func (a *Client) GetSetting(params *GetSettingParams, authInfo runtime.ClientAuthInfoWriter) (*GetSettingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSetting",
		Method:             "GET",
		PathPattern:        "/autorecordings/{extension}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSettingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSetting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateSetting sets the recording setting of a device
*/
func (a *Client) UpdateSetting(params *UpdateSettingParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateSettingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSettingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateSetting",
		Method:             "PUT",
		PathPattern:        "/autorecordings/{extension}/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateSettingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSettingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateSetting: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
