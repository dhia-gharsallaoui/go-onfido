# WebhookEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Indicates the resource affected by this event. | 
**Action** | Pointer to [**WebhookEventType**](WebhookEventType.md) | The event that triggered this webhook. | [optional] 
**Object** | Pointer to [**WebhookEventPayloadObject**](WebhookEventPayloadObject.md) |  | [optional] 
**Resource** | Pointer to **map[string]interface{}** | The resource affected by this event. | [optional] 

## Methods

### NewWebhookEventPayload

`func NewWebhookEventPayload(resourceType string, ) *WebhookEventPayload`

NewWebhookEventPayload instantiates a new WebhookEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventPayloadWithDefaults

`func NewWebhookEventPayloadWithDefaults() *WebhookEventPayload`

NewWebhookEventPayloadWithDefaults instantiates a new WebhookEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *WebhookEventPayload) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *WebhookEventPayload) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *WebhookEventPayload) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetAction

`func (o *WebhookEventPayload) GetAction() WebhookEventType`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookEventPayload) GetActionOk() (*WebhookEventType, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookEventPayload) SetAction(v WebhookEventType)`

SetAction sets Action field to given value.

### HasAction

`func (o *WebhookEventPayload) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetObject

`func (o *WebhookEventPayload) GetObject() WebhookEventPayloadObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookEventPayload) GetObjectOk() (*WebhookEventPayloadObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookEventPayload) SetObject(v WebhookEventPayloadObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *WebhookEventPayload) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetResource

`func (o *WebhookEventPayload) GetResource() map[string]interface{}`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *WebhookEventPayload) GetResourceOk() (*map[string]interface{}, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *WebhookEventPayload) SetResource(v map[string]interface{})`

SetResource sets Resource field to given value.

### HasResource

`func (o *WebhookEventPayload) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


