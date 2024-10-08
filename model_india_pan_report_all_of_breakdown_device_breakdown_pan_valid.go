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

// checks if the IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid{}

// IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid struct for IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid
type IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid struct {
	Result *string `json:"result,omitempty"`
}

// NewIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid instantiates a new IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid() *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid {
	this := IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid{}
	return &this
}

// NewIndiaPanReportAllOfBreakdownDeviceBreakdownPanValidWithDefaults instantiates a new IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndiaPanReportAllOfBreakdownDeviceBreakdownPanValidWithDefaults() *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid {
	this := IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) SetResult(v string) {
	o.Result = &v
}

func (o IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid struct {
	value *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid
	isSet bool
}

func (v NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) Get() *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid {
	return v.value
}

func (v *NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) Set(val *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) {
	v.value = val
	v.isSet = true
}

func (v NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) IsSet() bool {
	return v.isSet
}

func (v *NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid(val *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) *NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid {
	return &NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid{value: val, isSet: true}
}

func (v NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


