/*
Onfido API v3.6

The Onfido API (v3.6)

API version: v3.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Webhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Webhook{}

// Webhook struct for Webhook
type Webhook struct {
	// Determine if the webhook is active.
	Enabled *bool `json:"enabled,omitempty"`
	// The events that will be published to the webhook. If the events parameter is omitted all the events will be subscribed. 
	Events []WebhookEventType `json:"events,omitempty"`
	// The environments from which the webhook will receive events. Allowed values are “sandbox” and “live”. If the environments parameter is omitted the webhook will receive events from both environments. 
	Environments []string `json:"environments,omitempty"`
	// Webhook version used to control the payload object when sending webhooks.
	PayloadVersion *int32 `json:"payload_version,omitempty"`
	// The unique identifier of the webhook.
	Id string `json:"id"`
	// The url that will listen to notifications (must be https).
	Url *string `json:"url,omitempty"`
	// Webhook secret token used to sign the webhook's payload.
	Token *string `json:"token,omitempty"`
	// The API endpoint to retrieve the webhook.
	Href *string `json:"href,omitempty"`
}

type _Webhook Webhook

// NewWebhook instantiates a new Webhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhook(id string) *Webhook {
	this := Webhook{}
	this.Id = id
	return &this
}

// NewWebhookWithDefaults instantiates a new Webhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookWithDefaults() *Webhook {
	this := Webhook{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Webhook) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Webhook) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Webhook) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *Webhook) GetEvents() []WebhookEventType {
	if o == nil || IsNil(o.Events) {
		var ret []WebhookEventType
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEventsOk() ([]WebhookEventType, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *Webhook) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []WebhookEventType and assigns it to the Events field.
func (o *Webhook) SetEvents(v []WebhookEventType) {
	o.Events = v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *Webhook) GetEnvironments() []string {
	if o == nil || IsNil(o.Environments) {
		var ret []string
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetEnvironmentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Environments) {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *Webhook) HasEnvironments() bool {
	if o != nil && !IsNil(o.Environments) {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []string and assigns it to the Environments field.
func (o *Webhook) SetEnvironments(v []string) {
	o.Environments = v
}

// GetPayloadVersion returns the PayloadVersion field value if set, zero value otherwise.
func (o *Webhook) GetPayloadVersion() int32 {
	if o == nil || IsNil(o.PayloadVersion) {
		var ret int32
		return ret
	}
	return *o.PayloadVersion
}

// GetPayloadVersionOk returns a tuple with the PayloadVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetPayloadVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.PayloadVersion) {
		return nil, false
	}
	return o.PayloadVersion, true
}

// HasPayloadVersion returns a boolean if a field has been set.
func (o *Webhook) HasPayloadVersion() bool {
	if o != nil && !IsNil(o.PayloadVersion) {
		return true
	}

	return false
}

// SetPayloadVersion gets a reference to the given int32 and assigns it to the PayloadVersion field.
func (o *Webhook) SetPayloadVersion(v int32) {
	o.PayloadVersion = &v
}

// GetId returns the Id field value
func (o *Webhook) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Webhook) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Webhook) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Webhook) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Webhook) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Webhook) SetUrl(v string) {
	o.Url = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Webhook) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Webhook) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Webhook) SetToken(v string) {
	o.Token = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Webhook) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Webhook) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Webhook) SetHref(v string) {
	o.Href = &v
}

func (o Webhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Webhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Environments) {
		toSerialize["environments"] = o.Environments
	}
	if !IsNil(o.PayloadVersion) {
		toSerialize["payload_version"] = o.PayloadVersion
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

func (o *Webhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhook := _Webhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhook)

	if err != nil {
		return err
	}

	*o = Webhook(varWebhook)

	return err
}

type NullableWebhook struct {
	value *Webhook
	isSet bool
}

func (v NullableWebhook) Get() *Webhook {
	return v.value
}

func (v *NullableWebhook) Set(val *Webhook) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhook(val *Webhook) *NullableWebhook {
	return &NullableWebhook{value: val, isSet: true}
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


