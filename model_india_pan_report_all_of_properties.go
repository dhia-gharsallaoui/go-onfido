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

// checks if the IndiaPanReportAllOfProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndiaPanReportAllOfProperties{}

// IndiaPanReportAllOfProperties struct for IndiaPanReportAllOfProperties
type IndiaPanReportAllOfProperties struct {
	Device *IndiaPanReportAllOfPropertiesDevice `json:"device,omitempty"`
}

// NewIndiaPanReportAllOfProperties instantiates a new IndiaPanReportAllOfProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndiaPanReportAllOfProperties() *IndiaPanReportAllOfProperties {
	this := IndiaPanReportAllOfProperties{}
	return &this
}

// NewIndiaPanReportAllOfPropertiesWithDefaults instantiates a new IndiaPanReportAllOfProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndiaPanReportAllOfPropertiesWithDefaults() *IndiaPanReportAllOfProperties {
	this := IndiaPanReportAllOfProperties{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *IndiaPanReportAllOfProperties) GetDevice() IndiaPanReportAllOfPropertiesDevice {
	if o == nil || IsNil(o.Device) {
		var ret IndiaPanReportAllOfPropertiesDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndiaPanReportAllOfProperties) GetDeviceOk() (*IndiaPanReportAllOfPropertiesDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *IndiaPanReportAllOfProperties) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given IndiaPanReportAllOfPropertiesDevice and assigns it to the Device field.
func (o *IndiaPanReportAllOfProperties) SetDevice(v IndiaPanReportAllOfPropertiesDevice) {
	o.Device = &v
}

func (o IndiaPanReportAllOfProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndiaPanReportAllOfProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return toSerialize, nil
}

type NullableIndiaPanReportAllOfProperties struct {
	value *IndiaPanReportAllOfProperties
	isSet bool
}

func (v NullableIndiaPanReportAllOfProperties) Get() *IndiaPanReportAllOfProperties {
	return v.value
}

func (v *NullableIndiaPanReportAllOfProperties) Set(val *IndiaPanReportAllOfProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableIndiaPanReportAllOfProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableIndiaPanReportAllOfProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndiaPanReportAllOfProperties(val *IndiaPanReportAllOfProperties) *NullableIndiaPanReportAllOfProperties {
	return &NullableIndiaPanReportAllOfProperties{value: val, isSet: true}
}

func (v NullableIndiaPanReportAllOfProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndiaPanReportAllOfProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


