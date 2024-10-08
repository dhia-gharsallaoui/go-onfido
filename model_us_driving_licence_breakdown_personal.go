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

// checks if the UsDrivingLicenceBreakdownPersonal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsDrivingLicenceBreakdownPersonal{}

// UsDrivingLicenceBreakdownPersonal Asserts whether the personal data provided matches a real driving license in the DMV driver's license database.
type UsDrivingLicenceBreakdownPersonal struct {
	Result *string `json:"result,omitempty"`
	Breakdown *UsDrivingLicenceBreakdownPersonalBreakdown `json:"breakdown,omitempty"`
}

// NewUsDrivingLicenceBreakdownPersonal instantiates a new UsDrivingLicenceBreakdownPersonal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsDrivingLicenceBreakdownPersonal() *UsDrivingLicenceBreakdownPersonal {
	this := UsDrivingLicenceBreakdownPersonal{}
	return &this
}

// NewUsDrivingLicenceBreakdownPersonalWithDefaults instantiates a new UsDrivingLicenceBreakdownPersonal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsDrivingLicenceBreakdownPersonalWithDefaults() *UsDrivingLicenceBreakdownPersonal {
	this := UsDrivingLicenceBreakdownPersonal{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonal) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonal) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonal) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *UsDrivingLicenceBreakdownPersonal) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonal) GetBreakdown() UsDrivingLicenceBreakdownPersonalBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret UsDrivingLicenceBreakdownPersonalBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonal) GetBreakdownOk() (*UsDrivingLicenceBreakdownPersonalBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonal) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given UsDrivingLicenceBreakdownPersonalBreakdown and assigns it to the Breakdown field.
func (o *UsDrivingLicenceBreakdownPersonal) SetBreakdown(v UsDrivingLicenceBreakdownPersonalBreakdown) {
	o.Breakdown = &v
}

func (o UsDrivingLicenceBreakdownPersonal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsDrivingLicenceBreakdownPersonal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableUsDrivingLicenceBreakdownPersonal struct {
	value *UsDrivingLicenceBreakdownPersonal
	isSet bool
}

func (v NullableUsDrivingLicenceBreakdownPersonal) Get() *UsDrivingLicenceBreakdownPersonal {
	return v.value
}

func (v *NullableUsDrivingLicenceBreakdownPersonal) Set(val *UsDrivingLicenceBreakdownPersonal) {
	v.value = val
	v.isSet = true
}

func (v NullableUsDrivingLicenceBreakdownPersonal) IsSet() bool {
	return v.isSet
}

func (v *NullableUsDrivingLicenceBreakdownPersonal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsDrivingLicenceBreakdownPersonal(val *UsDrivingLicenceBreakdownPersonal) *NullableUsDrivingLicenceBreakdownPersonal {
	return &NullableUsDrivingLicenceBreakdownPersonal{value: val, isSet: true}
}

func (v NullableUsDrivingLicenceBreakdownPersonal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsDrivingLicenceBreakdownPersonal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


