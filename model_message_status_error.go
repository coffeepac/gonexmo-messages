/*
 * Messages API
 *
 * The Messaging API is a new API that consolidates all messaging channels. It encapsulates a user (developer) from having to use multiple APIs to interact with our various channels such as SMS, MMS, Viber, Facebook Messenger, etc. The API normalises information across all channels to abstracted to, from and content. This API is currently Beta.
 *
 * API version: 0.3.0
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type MessageStatusError struct {
	// The error code. See [our errors list](https://developer.nexmo.com/api-errors/messages-olympus) for a list of possible errors
	Code int32 `json:"code,omitempty"`
	// Text describing the error. See [our errors list](https://developer.nexmo.com/api-errors/messages-olympus) for a list of possible errors
	Reason string `json:"reason,omitempty"`
}