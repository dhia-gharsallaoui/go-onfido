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

// checks if the ProofOfAddressBreakdownDataComparison type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProofOfAddressBreakdownDataComparison{}

// ProofOfAddressBreakdownDataComparison Asserts whether the first name, last name and address provided by the applicant match those on the PoA document.
type ProofOfAddressBreakdownDataComparison struct {
	Result *string `json:"result,omitempty"`
	Breakdown *ProofOfAddressBreakdownDataComparisonBreakdown `json:"breakdown,omitempty"`
}

// NewProofOfAddressBreakdownDataComparison instantiates a new ProofOfAddressBreakdownDataComparison object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProofOfAddressBreakdownDataComparison() *ProofOfAddressBreakdownDataComparison {
	this := ProofOfAddressBreakdownDataComparison{}
	return &this
}

// NewProofOfAddressBreakdownDataComparisonWithDefaults instantiates a new ProofOfAddressBreakdownDataComparison object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProofOfAddressBreakdownDataComparisonWithDefaults() *ProofOfAddressBreakdownDataComparison {
	this := ProofOfAddressBreakdownDataComparison{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ProofOfAddressBreakdownDataComparison) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressBreakdownDataComparison) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ProofOfAddressBreakdownDataComparison) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ProofOfAddressBreakdownDataComparison) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *ProofOfAddressBreakdownDataComparison) GetBreakdown() ProofOfAddressBreakdownDataComparisonBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret ProofOfAddressBreakdownDataComparisonBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressBreakdownDataComparison) GetBreakdownOk() (*ProofOfAddressBreakdownDataComparisonBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *ProofOfAddressBreakdownDataComparison) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given ProofOfAddressBreakdownDataComparisonBreakdown and assigns it to the Breakdown field.
func (o *ProofOfAddressBreakdownDataComparison) SetBreakdown(v ProofOfAddressBreakdownDataComparisonBreakdown) {
	o.Breakdown = &v
}

func (o ProofOfAddressBreakdownDataComparison) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProofOfAddressBreakdownDataComparison) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableProofOfAddressBreakdownDataComparison struct {
	value *ProofOfAddressBreakdownDataComparison
	isSet bool
}

func (v NullableProofOfAddressBreakdownDataComparison) Get() *ProofOfAddressBreakdownDataComparison {
	return v.value
}

func (v *NullableProofOfAddressBreakdownDataComparison) Set(val *ProofOfAddressBreakdownDataComparison) {
	v.value = val
	v.isSet = true
}

func (v NullableProofOfAddressBreakdownDataComparison) IsSet() bool {
	return v.isSet
}

func (v *NullableProofOfAddressBreakdownDataComparison) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProofOfAddressBreakdownDataComparison(val *ProofOfAddressBreakdownDataComparison) *NullableProofOfAddressBreakdownDataComparison {
	return &NullableProofOfAddressBreakdownDataComparison{value: val, isSet: true}
}

func (v NullableProofOfAddressBreakdownDataComparison) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProofOfAddressBreakdownDataComparison) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


