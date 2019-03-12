# MessagePropertyContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of message that you are sending.  **Messenger**: supports all types.  **Viber Service Messages**: supports &#x60;image&#x60; and &#x60;text&#x60; and &#x60;custom&#x60;.  **WhatsApp**: supports &#x60;template&#x60; and &#x60;text&#x60;.  **SMS**: supports &#x60;text&#x60;.  **MMS**: supports &#x60;image&#x60;.  | [optional] 
**Text** | **string** | The text of the message.  **Messenger**: Is limited to 640 characters  **SMS** or **Viber**: Is 1000 characters  **WhatsApp**: is 4096 characters  | [optional] 
**Image** | [**ImageProperty**](ImageProperty.md) |  | [optional] 
**Audio** | [**AudioProperty**](AudioProperty.md) |  | [optional] 
**Video** | [**VideoProperty**](VideoProperty.md) |  | [optional] 
**File** | [**FileProperty**](FileProperty.md) |  | [optional] 
**Template** | [**TemplateProperty**](TemplateProperty.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


