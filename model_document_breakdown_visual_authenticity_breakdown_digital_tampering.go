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

// checks if the DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering{}

// DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering Indication of digital tampering in the image.
type DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering struct {
	Result *string `json:"result,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// NewDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering instantiates a new DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering() *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering {
	this := DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering{}
	return &this
}

// NewDocumentBreakdownVisualAuthenticityBreakdownDigitalTamperingWithDefaults instantiates a new DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentBreakdownVisualAuthenticityBreakdownDigitalTamperingWithDefaults() *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering {
	this := DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) SetResult(v string) {
	o.Result = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering struct {
	value *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering
	isSet bool
}

func (v NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) Get() *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering {
	return v.value
}

func (v *NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) Set(val *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering(val *DocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) *NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering {
	return &NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering{value: val, isSet: true}
}

func (v NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentBreakdownVisualAuthenticityBreakdownDigitalTampering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


