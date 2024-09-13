# WebhooksResendItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | ID of the resource whose webhooks are to be retriggered. | 
**Event** | [**WebhookEventType**](WebhookEventType.md) | The events that should retrigger webhooks. Accepts values check.completed. | 

## Methods

### NewWebhooksResendItem

`func NewWebhooksResendItem(resourceId string, event WebhookEventType, ) *WebhooksResendItem`

NewWebhooksResendItem instantiates a new WebhooksResendItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhooksResendItemWithDefaults

`func NewWebhooksResendItemWithDefaults() *WebhooksResendItem`

NewWebhooksResendItemWithDefaults instantiates a new WebhooksResendItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *WebhooksResendItem) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *WebhooksResendItem) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *WebhooksResendItem) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetEvent

`func (o *WebhooksResendItem) GetEvent() WebhookEventType`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhooksResendItem) GetEventOk() (*WebhookEventType, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhooksResendItem) SetEvent(v WebhookEventType)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


