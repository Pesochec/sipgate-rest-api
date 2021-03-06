// Code generated by go-swagger; DO NOT EDIT.

package rtcm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new rtcm API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rtcm API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCalls(params *GetCallsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCallsOK, error)

	HangupCall(params *HangupCallParams, authInfo runtime.ClientAuthInfoWriter) (*HangupCallNoContent, error)

	SendCallDtmf(params *SendCallDtmfParams, authInfo runtime.ClientAuthInfoWriter) (*SendCallDtmfNoContent, error)

	SetCallHold(params *SetCallHoldParams, authInfo runtime.ClientAuthInfoWriter) (*SetCallHoldNoContent, error)

	SetCallMuted(params *SetCallMutedParams, authInfo runtime.ClientAuthInfoWriter) (*SetCallMutedNoContent, error)

	SetCallRecording(params *SetCallRecordingParams, authInfo runtime.ClientAuthInfoWriter) (*SetCallRecordingNoContent, error)

	StartCallAnnouncement(params *StartCallAnnouncementParams, authInfo runtime.ClientAuthInfoWriter) (*StartCallAnnouncementNoContent, error)

	TransferCall(params *TransferCallParams, authInfo runtime.ClientAuthInfoWriter) (*TransferCallNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCalls lists all calls currently established calls the response does not contain ringing calls or currently running voicemail recordings
*/
func (a *Client) GetCalls(params *GetCallsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCallsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCallsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCalls",
		Method:             "GET",
		PathPattern:        "/calls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCallsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCallsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCalls: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  HangupCall hangs up call
*/
func (a *Client) HangupCall(params *HangupCallParams, authInfo runtime.ClientAuthInfoWriter) (*HangupCallNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHangupCallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "hangupCall",
		Method:             "DELETE",
		PathPattern:        "/calls/{callId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HangupCallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HangupCallNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for hangupCall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SendCallDtmf sends d t m f tone sequences to all participants
*/
func (a *Client) SendCallDtmf(params *SendCallDtmfParams, authInfo runtime.ClientAuthInfoWriter) (*SendCallDtmfNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendCallDtmfParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendCallDtmf",
		Method:             "POST",
		PathPattern:        "/calls/{callId}/dtmf",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendCallDtmfReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendCallDtmfNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendCallDtmf: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetCallHold holds resume all participants
*/
func (a *Client) SetCallHold(params *SetCallHoldParams, authInfo runtime.ClientAuthInfoWriter) (*SetCallHoldNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetCallHoldParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setCallHold",
		Method:             "PUT",
		PathPattern:        "/calls/{callId}/hold",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetCallHoldReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetCallHoldNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setCallHold: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetCallMuted mutes unmute yourself
*/
func (a *Client) SetCallMuted(params *SetCallMutedParams, authInfo runtime.ClientAuthInfoWriter) (*SetCallMutedNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetCallMutedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setCallMuted",
		Method:             "PUT",
		PathPattern:        "/calls/{callId}/muted",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetCallMutedReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetCallMutedNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setCallMuted: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetCallRecording starts or stop a call recording
*/
func (a *Client) SetCallRecording(params *SetCallRecordingParams, authInfo runtime.ClientAuthInfoWriter) (*SetCallRecordingNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetCallRecordingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setCallRecording",
		Method:             "PUT",
		PathPattern:        "/calls/{callId}/recording",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetCallRecordingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetCallRecordingNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setCallRecording: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StartCallAnnouncement starts playing a new announcement to all participants

  Currently the sound file needs to be a mono 16bit PCM WAV file with a sampling rate of 8kHz. MP3 conversion example: mpg123 --rate 8000 --mono -w output.wav input.mp3
*/
func (a *Client) StartCallAnnouncement(params *StartCallAnnouncementParams, authInfo runtime.ClientAuthInfoWriter) (*StartCallAnnouncementNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartCallAnnouncementParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startCallAnnouncement",
		Method:             "POST",
		PathPattern:        "/calls/{callId}/announcements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartCallAnnouncementReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartCallAnnouncementNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for startCallAnnouncement: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TransferCall transfers call
*/
func (a *Client) TransferCall(params *TransferCallParams, authInfo runtime.ClientAuthInfoWriter) (*TransferCallNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransferCallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "transferCall",
		Method:             "POST",
		PathPattern:        "/calls/{callId}/transfer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TransferCallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TransferCallNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for transferCall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
