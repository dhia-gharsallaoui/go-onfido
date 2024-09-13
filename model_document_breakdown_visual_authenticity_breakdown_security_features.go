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

// checks if the DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures{}

// DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures Security features expected on the document are missing or wrong.
type DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures struct {
	Result *string `json:"result,omitempty"`
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// NewDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures instantiates a new DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures() *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures {
	this := DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures{}
	return &this
}

// NewDocumentBreakdownVisualAuthenticityBreakdownSecurityFeaturesWithDefaults instantiates a new DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentBreakdownVisualAuthenticityBreakdownSecurityFeaturesWithDefaults() *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures {
	this := DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) SetResult(v string) {
	o.Result = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

func (o DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures struct {
	value *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures
	isSet bool
}

func (v NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) Get() *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures {
	return v.value
}

func (v *NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) Set(val *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures(val *DocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) *NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures {
	return &NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures{value: val, isSet: true}
}

func (v NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentBreakdownVisualAuthenticityBreakdownSecurityFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


