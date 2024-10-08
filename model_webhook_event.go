/*
Onfido API v3.6

The Onfido API (v3.6)

API version: v3.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the WebhookEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEvent{}

// WebhookEvent struct for WebhookEvent
type WebhookEvent struct {
	Payload *WebhookEventPayload `json:"payload,omitempty"`
}

// NewWebhookEvent instantiates a new WebhookEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEvent() *WebhookEvent {
	this := WebhookEvent{}
	return &this
}

// NewWebhookEventWithDefaults instantiates a new WebhookEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEventWithDefaults() *WebhookEvent {
	this := WebhookEvent{}
	return &this
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *WebhookEvent) GetPayload() WebhookEventPayload {
	if o == nil || IsNil(o.Payload) {
		var ret WebhookEventPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookEvent) GetPayloadOk() (*WebhookEventPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *WebhookEvent) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given WebhookEventPayload and assigns it to the Payload field.
func (o *WebhookEvent) SetPayload(v WebhookEventPayload) {
	o.Payload = &v
}

func (o WebhookEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableWebhookEvent struct {
	value *WebhookEvent
	isSet bool
}

func (v NullableWebhookEvent) Get() *WebhookEvent {
	return v.value
}

func (v *NullableWebhookEvent) Set(val *WebhookEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEvent(val *WebhookEvent) *NullableWebhookEvent {
	return &NullableWebhookEvent{value: val, isSet: true}
}

func (v NullableWebhookEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


