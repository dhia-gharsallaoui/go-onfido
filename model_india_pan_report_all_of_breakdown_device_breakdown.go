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

// checks if the IndiaPanReportAllOfBreakdownDeviceBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IndiaPanReportAllOfBreakdownDeviceBreakdown{}

// IndiaPanReportAllOfBreakdownDeviceBreakdown struct for IndiaPanReportAllOfBreakdownDeviceBreakdown
type IndiaPanReportAllOfBreakdownDeviceBreakdown struct {
	PanValid *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid `json:"pan_valid,omitempty"`
	NameMatch *IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid `json:"name_match,omitempty"`
}

// NewIndiaPanReportAllOfBreakdownDeviceBreakdown instantiates a new IndiaPanReportAllOfBreakdownDeviceBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndiaPanReportAllOfBreakdownDeviceBreakdown() *IndiaPanReportAllOfBreakdownDeviceBreakdown {
	this := IndiaPanReportAllOfBreakdownDeviceBreakdown{}
	return &this
}

// NewIndiaPanReportAllOfBreakdownDeviceBreakdownWithDefaults instantiates a new IndiaPanReportAllOfBreakdownDeviceBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndiaPanReportAllOfBreakdownDeviceBreakdownWithDefaults() *IndiaPanReportAllOfBreakdownDeviceBreakdown {
	this := IndiaPanReportAllOfBreakdownDeviceBreakdown{}
	return &this
}

// GetPanValid returns the PanValid field value if set, zero value otherwise.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdown) GetPanValid() IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid {
	if o == nil || IsNil(o.PanValid) {
		var ret IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid
		return ret
	}
	return *o.PanValid
}

// GetPanValidOk returns a tuple with the PanValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdown) GetPanValidOk() (*IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid, bool) {
	if o == nil || IsNil(o.PanValid) {
		return nil, false
	}
	return o.PanValid, true
}

// HasPanValid returns a boolean if a field has been set.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdown) HasPanValid() bool {
	if o != nil && !IsNil(o.PanValid) {
		return true
	}

	return false
}

// SetPanValid gets a reference to the given IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid and assigns it to the PanValid field.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdown) SetPanValid(v IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) {
	o.PanValid = &v
}

// GetNameMatch returns the NameMatch field value if set, zero value otherwise.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdown) GetNameMatch() IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid {
	if o == nil || IsNil(o.NameMatch) {
		var ret IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid
		return ret
	}
	return *o.NameMatch
}

// GetNameMatchOk returns a tuple with the NameMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdown) GetNameMatchOk() (*IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid, bool) {
	if o == nil || IsNil(o.NameMatch) {
		return nil, false
	}
	return o.NameMatch, true
}

// HasNameMatch returns a boolean if a field has been set.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdown) HasNameMatch() bool {
	if o != nil && !IsNil(o.NameMatch) {
		return true
	}

	return false
}

// SetNameMatch gets a reference to the given IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid and assigns it to the NameMatch field.
func (o *IndiaPanReportAllOfBreakdownDeviceBreakdown) SetNameMatch(v IndiaPanReportAllOfBreakdownDeviceBreakdownPanValid) {
	o.NameMatch = &v
}

func (o IndiaPanReportAllOfBreakdownDeviceBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IndiaPanReportAllOfBreakdownDeviceBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PanValid) {
		toSerialize["pan_valid"] = o.PanValid
	}
	if !IsNil(o.NameMatch) {
		toSerialize["name_match"] = o.NameMatch
	}
	return toSerialize, nil
}

type NullableIndiaPanReportAllOfBreakdownDeviceBreakdown struct {
	value *IndiaPanReportAllOfBreakdownDeviceBreakdown
	isSet bool
}

func (v NullableIndiaPanReportAllOfBreakdownDeviceBreakdown) Get() *IndiaPanReportAllOfBreakdownDeviceBreakdown {
	return v.value
}

func (v *NullableIndiaPanReportAllOfBreakdownDeviceBreakdown) Set(val *IndiaPanReportAllOfBreakdownDeviceBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableIndiaPanReportAllOfBreakdownDeviceBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableIndiaPanReportAllOfBreakdownDeviceBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndiaPanReportAllOfBreakdownDeviceBreakdown(val *IndiaPanReportAllOfBreakdownDeviceBreakdown) *NullableIndiaPanReportAllOfBreakdownDeviceBreakdown {
	return &NullableIndiaPanReportAllOfBreakdownDeviceBreakdown{value: val, isSet: true}
}

func (v NullableIndiaPanReportAllOfBreakdownDeviceBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndiaPanReportAllOfBreakdownDeviceBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


