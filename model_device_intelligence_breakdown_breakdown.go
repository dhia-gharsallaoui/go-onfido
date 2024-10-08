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

// checks if the DeviceIntelligenceBreakdownBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceIntelligenceBreakdownBreakdown{}

// DeviceIntelligenceBreakdownBreakdown struct for DeviceIntelligenceBreakdownBreakdown
type DeviceIntelligenceBreakdownBreakdown struct {
	Device *DeviceIntelligenceBreakdownBreakdownDevice `json:"device,omitempty"`
}

// NewDeviceIntelligenceBreakdownBreakdown instantiates a new DeviceIntelligenceBreakdownBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIntelligenceBreakdownBreakdown() *DeviceIntelligenceBreakdownBreakdown {
	this := DeviceIntelligenceBreakdownBreakdown{}
	return &this
}

// NewDeviceIntelligenceBreakdownBreakdownWithDefaults instantiates a new DeviceIntelligenceBreakdownBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIntelligenceBreakdownBreakdownWithDefaults() *DeviceIntelligenceBreakdownBreakdown {
	this := DeviceIntelligenceBreakdownBreakdown{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *DeviceIntelligenceBreakdownBreakdown) GetDevice() DeviceIntelligenceBreakdownBreakdownDevice {
	if o == nil || IsNil(o.Device) {
		var ret DeviceIntelligenceBreakdownBreakdownDevice
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntelligenceBreakdownBreakdown) GetDeviceOk() (*DeviceIntelligenceBreakdownBreakdownDevice, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *DeviceIntelligenceBreakdownBreakdown) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given DeviceIntelligenceBreakdownBreakdownDevice and assigns it to the Device field.
func (o *DeviceIntelligenceBreakdownBreakdown) SetDevice(v DeviceIntelligenceBreakdownBreakdownDevice) {
	o.Device = &v
}

func (o DeviceIntelligenceBreakdownBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceIntelligenceBreakdownBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	return toSerialize, nil
}

type NullableDeviceIntelligenceBreakdownBreakdown struct {
	value *DeviceIntelligenceBreakdownBreakdown
	isSet bool
}

func (v NullableDeviceIntelligenceBreakdownBreakdown) Get() *DeviceIntelligenceBreakdownBreakdown {
	return v.value
}

func (v *NullableDeviceIntelligenceBreakdownBreakdown) Set(val *DeviceIntelligenceBreakdownBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIntelligenceBreakdownBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIntelligenceBreakdownBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIntelligenceBreakdownBreakdown(val *DeviceIntelligenceBreakdownBreakdown) *NullableDeviceIntelligenceBreakdownBreakdown {
	return &NullableDeviceIntelligenceBreakdownBreakdown{value: val, isSet: true}
}

func (v NullableDeviceIntelligenceBreakdownBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIntelligenceBreakdownBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


