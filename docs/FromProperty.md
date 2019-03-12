# FromProperty

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of message that you want to send. | 
**Id** | **string** | Your ID for the platform that you are sending from.  **Messenger**: This value should be the &#x60;to.id&#x60; value you received in the inbound messenger event.  **Viber**: This is your Service Message ID given to you by Nexmo Account Manager. To find out more please visit [nexmo.com/products/messages](https://www.nexmo.com/products/messages).  | [optional] 
**Number** | **string** | **SMS**: The phone number of the message recipient in the [E.164](https://en.wikipedia.org/wiki/E.164) format. Don&#39;t use a leading + or 00 when entering a phone number, start with the country code, for example, 447700900000.  **WhatsApp**: This is your WhatsApp Business Number  given to you by Nexmo Account Manager. To find out more please visit [nexmo.com/products/messages](https://www.nexmo.com/products/messages).  **MMS**: US shortcode  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


