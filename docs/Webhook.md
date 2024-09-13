# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Determine if the webhook is active. | [optional] 
**Events** | Pointer to [**[]WebhookEventType**](WebhookEventType.md) | The events that will be published to the webhook. If the events parameter is omitted all the events will be subscribed.  | [optional] 
**Environments** | Pointer to **[]string** | The environments from which the webhook will receive events. Allowed values are “sandbox” and “live”. If the environments parameter is omitted the webhook will receive events from both environments.  | [optional] 
**PayloadVersion** | Pointer to **int32** | Webhook version used to control the payload object when sending webhooks. | [optional] 
**Id** | **string** | The unique identifier of the webhook. | 
**Url** | Pointer to **string** | The url that will listen to notifications (must be https). | [optional] 
**Token** | Pointer to **string** | Webhook secret token used to sign the webhook&#39;s payload. | [optional] 
**Href** | Pointer to **string** | The API endpoint to retrieve the webhook. | [optional] 

## Methods

### NewWebhook

`func NewWebhook(id string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Webhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvents

`func (o *Webhook) GetEvents() []WebhookEventType`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Webhook) GetEventsOk() (*[]WebhookEventType, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Webhook) SetEvents(v []WebhookEventType)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Webhook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetEnvironments

`func (o *Webhook) GetEnvironments() []string`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *Webhook) GetEnvironmentsOk() (*[]string, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *Webhook) SetEnvironments(v []string)`

SetEnvironments sets Environments field to given value.

### HasEnvironments

`func (o *Webhook) HasEnvironments() bool`

HasEnvironments returns a boolean if a field has been set.

### GetPayloadVersion

`func (o *Webhook) GetPayloadVersion() int32`

GetPayloadVersion returns the PayloadVersion field if non-nil, zero value otherwise.

### GetPayloadVersionOk

`func (o *Webhook) GetPayloadVersionOk() (*int32, bool)`

GetPayloadVersionOk returns a tuple with the PayloadVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadVersion

`func (o *Webhook) SetPayloadVersion(v int32)`

SetPayloadVersion sets PayloadVersion field to given value.

### HasPayloadVersion

`func (o *Webhook) HasPayloadVersion() bool`

HasPayloadVersion returns a boolean if a field has been set.

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetToken

`func (o *Webhook) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Webhook) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Webhook) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *Webhook) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetHref

`func (o *Webhook) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Webhook) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Webhook) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Webhook) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


