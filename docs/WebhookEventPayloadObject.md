# WebhookEventPayloadObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the resource. | 
**Status** | Pointer to **string** | The current state of the object, if available. | [optional] 
**StartedAtIso8601** | Pointer to **time.Time** | The date and time when the operation was started, if available. | [optional] 
**CompletedAtIso8601** | Pointer to **time.Time** | The date and time when the operation was completed, if available. | [optional] 
**Href** | **string** | The uri of the resource. | 

## Methods

### NewWebhookEventPayloadObject

`func NewWebhookEventPayloadObject(id string, href string, ) *WebhookEventPayloadObject`

NewWebhookEventPayloadObject instantiates a new WebhookEventPayloadObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookEventPayloadObjectWithDefaults

`func NewWebhookEventPayloadObjectWithDefaults() *WebhookEventPayloadObject`

NewWebhookEventPayloadObjectWithDefaults instantiates a new WebhookEventPayloadObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookEventPayloadObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookEventPayloadObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookEventPayloadObject) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *WebhookEventPayloadObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookEventPayloadObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookEventPayloadObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookEventPayloadObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAtIso8601

`func (o *WebhookEventPayloadObject) GetStartedAtIso8601() time.Time`

GetStartedAtIso8601 returns the StartedAtIso8601 field if non-nil, zero value otherwise.

### GetStartedAtIso8601Ok

`func (o *WebhookEventPayloadObject) GetStartedAtIso8601Ok() (*time.Time, bool)`

GetStartedAtIso8601Ok returns a tuple with the StartedAtIso8601 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtIso8601

`func (o *WebhookEventPayloadObject) SetStartedAtIso8601(v time.Time)`

SetStartedAtIso8601 sets StartedAtIso8601 field to given value.

### HasStartedAtIso8601

`func (o *WebhookEventPayloadObject) HasStartedAtIso8601() bool`

HasStartedAtIso8601 returns a boolean if a field has been set.

### GetCompletedAtIso8601

`func (o *WebhookEventPayloadObject) GetCompletedAtIso8601() time.Time`

GetCompletedAtIso8601 returns the CompletedAtIso8601 field if non-nil, zero value otherwise.

### GetCompletedAtIso8601Ok

`func (o *WebhookEventPayloadObject) GetCompletedAtIso8601Ok() (*time.Time, bool)`

GetCompletedAtIso8601Ok returns a tuple with the CompletedAtIso8601 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAtIso8601

`func (o *WebhookEventPayloadObject) SetCompletedAtIso8601(v time.Time)`

SetCompletedAtIso8601 sets CompletedAtIso8601 field to given value.

### HasCompletedAtIso8601

`func (o *WebhookEventPayloadObject) HasCompletedAtIso8601() bool`

HasCompletedAtIso8601 returns a boolean if a field has been set.

### GetHref

`func (o *WebhookEventPayloadObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WebhookEventPayloadObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WebhookEventPayloadObject) SetHref(v string)`

SetHref sets Href field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


