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

// checks if the UsDrivingLicenceBreakdownAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsDrivingLicenceBreakdownAddress{}

// UsDrivingLicenceBreakdownAddress Asserts whether the address data provided matches a real driving license in the DMV driver's license database.
type UsDrivingLicenceBreakdownAddress struct {
	Result *string `json:"result,omitempty"`
	Breakdown *UsDrivingLicenceBreakdownAddressBreakdown `json:"breakdown,omitempty"`
}

// NewUsDrivingLicenceBreakdownAddress instantiates a new UsDrivingLicenceBreakdownAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsDrivingLicenceBreakdownAddress() *UsDrivingLicenceBreakdownAddress {
	this := UsDrivingLicenceBreakdownAddress{}
	return &this
}

// NewUsDrivingLicenceBreakdownAddressWithDefaults instantiates a new UsDrivingLicenceBreakdownAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsDrivingLicenceBreakdownAddressWithDefaults() *UsDrivingLicenceBreakdownAddress {
	this := UsDrivingLicenceBreakdownAddress{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownAddress) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownAddress) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownAddress) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *UsDrivingLicenceBreakdownAddress) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownAddress) GetBreakdown() UsDrivingLicenceBreakdownAddressBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret UsDrivingLicenceBreakdownAddressBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownAddress) GetBreakdownOk() (*UsDrivingLicenceBreakdownAddressBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownAddress) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given UsDrivingLicenceBreakdownAddressBreakdown and assigns it to the Breakdown field.
func (o *UsDrivingLicenceBreakdownAddress) SetBreakdown(v UsDrivingLicenceBreakdownAddressBreakdown) {
	o.Breakdown = &v
}

func (o UsDrivingLicenceBreakdownAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsDrivingLicenceBreakdownAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableUsDrivingLicenceBreakdownAddress struct {
	value *UsDrivingLicenceBreakdownAddress
	isSet bool
}

func (v NullableUsDrivingLicenceBreakdownAddress) Get() *UsDrivingLicenceBreakdownAddress {
	return v.value
}

func (v *NullableUsDrivingLicenceBreakdownAddress) Set(val *UsDrivingLicenceBreakdownAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableUsDrivingLicenceBreakdownAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableUsDrivingLicenceBreakdownAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsDrivingLicenceBreakdownAddress(val *UsDrivingLicenceBreakdownAddress) *NullableUsDrivingLicenceBreakdownAddress {
	return &NullableUsDrivingLicenceBreakdownAddress{value: val, isSet: true}
}

func (v NullableUsDrivingLicenceBreakdownAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsDrivingLicenceBreakdownAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


