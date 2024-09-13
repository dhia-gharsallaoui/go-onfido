# WebhookBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determine if the webhook is active. | [optional] 
**Events** | Pointer to [**[]WebhookEventType**](WebhookEventType.md) | The events that will be published to the webhook. If the events parameter is omitted all the events will be subscribed.  | [optional] 
**Environments** | Pointer to **[]string** | The environments from which the webhook will receive events. Allowed values are “sandbox” and “live”. If the environments parameter is omitted the webhook will receive events from both environments.  | [optional] 
**PayloadVersion** | Pointer to **int32** | Webhook version used to control the payload object when sending webhooks. | [optional] 
**Url** | **string** | The url that will listen to notifications (must be https). | 

## Methods

### NewWebhookBuilder

`func NewWebhookBuilder(url string, ) *WebhookBuilder`

NewWebhookBuilder instantiates a new WebhookBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookBuilderWithDefaults

`func NewWebhookBuilderWithDefaults() *WebhookBuilder`

NewWebhookBuilderWithDefaults instantiates a new WebhookBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *WebhookBuilder) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebhookBuilder) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebhookBuilder) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WebhookBuilder) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvents

`func (o *WebhookBuilder) GetEvents() []WebhookEventType`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WebhookBuilder) GetEventsOk() (*[]WebhookEventType, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WebhookBuilder) SetEvents(v []WebhookEventType)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WebhookBuilder) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetEnvironments

`func (o *WebhookBuilder) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *WebhookBuilder) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *WebhookBuilder) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *WebhookBuilder) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetPayloadVersion

`func (o *WebhookBuilder) GetPayloadVersion() int32`

GetPayloadVersion returns the PayloadVersion field if non-nil, zero value otherwise.

### GetPayloadVersionOk

`func (o *WebhookBuilder) GetPayloadVersionOk() (*int32, bool)`

GetPayloadVersionOk returns a tuple with the PayloadVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadVersion

`func (o *WebhookBuilder) SetPayloadVersion(v int32)`

SetPayloadVersion sets PayloadVersion field to given value.

### HasPayloadVersion

`func (o *WebhookBuilder) HasPayloadVersion() bool`

HasPayloadVersion returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookBuilder) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookBuilder) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookBuilder) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


