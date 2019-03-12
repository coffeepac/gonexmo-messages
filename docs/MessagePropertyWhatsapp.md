# MessagePropertyWhatsapp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | There are two policies that you can specify when sending an Message Template: &#x60;deterministic&#x60; and &#x60;fallback&#x60;. &#x60;deterministic&#x60; — Deliver the Message Template in exactly the language and locale asked for. &#x60;fallback&#x60; — Deliver the Message Template in the language that matches users language/locale setting on device. If one can not be found, deliver using the specified fallback language. | [optional] 
**Locale** | **string** | We are using the industry standard, BCP 47, for locales. So in your API call to the /messages API you will need to put “en-GB” and this will refer to the “en_GB” template for WhatsApp. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


