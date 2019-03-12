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

type FromProperty struct {
	// The type of message that you want to send.
	Type string `json:"type"`
	// Your ID for the platform that you are sending from.  **Messenger**: This value should be the `to.id` value you received in the inbound messenger event.  **Viber**: This is your Service Message ID given to you by Nexmo Account Manager. To find out more please visit [nexmo.com/products/messages](https://www.nexmo.com/products/messages). 
	Id string `json:"id,omitempty"`
	// **SMS**: The phone number of the message recipient in the [E.164](https://en.wikipedia.org/wiki/E.164) format. Don't use a leading + or 00 when entering a phone number, start with the country code, for example, 447700900000.  **WhatsApp**: This is your WhatsApp Business Number  given to you by Nexmo Account Manager. To find out more please visit [nexmo.com/products/messages](https://www.nexmo.com/products/messages).  **MMS**: US shortcode 
	Number string `json:"number,omitempty"`
}