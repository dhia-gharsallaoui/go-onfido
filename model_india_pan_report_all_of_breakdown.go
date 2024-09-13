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

// checks if the IndiaPanReportAllOfBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndiaPanReportAllOfBreakdown{}

// IndiaPanReportAllOfBreakdown struct for IndiaPanReportAllOfBreakdown
type IndiaPanReportAllOfBreakdown struct {
	Device *IndiaPanReportAllOfBreakdownDevice `json:"device,omitempty"`
}

// NewIndiaPanReportAllOfBreakdown instantiates a new IndiaPanReportAllOfBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndiaPanReportAllOfBreakdown() *IndiaPanReportAllOfBreakdown {
	this := IndiaPanReportAllOfBreakdown{}
	return &this
}

// NewIndiaPanReportAllOfBreakdownWithDefaults instantiates a new IndiaPanReportAllOfBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndiaPanReportAllOfBreakdownWithDefaults() *IndiaPanReportAllOfBreakdown {
	this := IndiaPanReportAllOfBreakdown{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *IndiaPanReportAllOfBreakdown) GetDevice() IndiaPanReportAllOfBreakdownDevice {
	if o == nil || IsNil(o.Device) {
		var ret IndiaPanReportAllOfBreakdownDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndiaPanReportAllOfBreakdown) GetDeviceOk() (*IndiaPanReportAllOfBreakdownDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *IndiaPanReportAllOfBreakdown) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given IndiaPanReportAllOfBreakdownDevice and assigns it to the Device field.
func (o *IndiaPanReportAllOfBreakdown) SetDevice(v IndiaPanReportAllOfBreakdownDevice) {
	o.Device = &v
}

func (o IndiaPanReportAllOfBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndiaPanReportAllOfBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return toSerialize, nil
}

type NullableIndiaPanReportAllOfBreakdown struct {
	value *IndiaPanReportAllOfBreakdown
	isSet bool
}

func (v NullableIndiaPanReportAllOfBreakdown) Get() *IndiaPanReportAllOfBreakdown {
	return v.value
}

func (v *NullableIndiaPanReportAllOfBreakdown) Set(val *IndiaPanReportAllOfBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableIndiaPanReportAllOfBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableIndiaPanReportAllOfBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndiaPanReportAllOfBreakdown(val *IndiaPanReportAllOfBreakdown) *NullableIndiaPanReportAllOfBreakdown {
	return &NullableIndiaPanReportAllOfBreakdown{value: val, isSet: true}
}

func (v NullableIndiaPanReportAllOfBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndiaPanReportAllOfBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


