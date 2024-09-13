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

// checks if the ProofOfAddressBreakdownImageIntegrity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProofOfAddressBreakdownImageIntegrity{}

// ProofOfAddressBreakdownImageIntegrity Asserts whether the quality of the uploaded document was sufficient to verify the address.
type ProofOfAddressBreakdownImageIntegrity struct {
	Result *string `json:"result,omitempty"`
	Breakdown *ProofOfAddressBreakdownImageIntegrityBreakdown `json:"breakdown,omitempty"`
}

// NewProofOfAddressBreakdownImageIntegrity instantiates a new ProofOfAddressBreakdownImageIntegrity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProofOfAddressBreakdownImageIntegrity() *ProofOfAddressBreakdownImageIntegrity {
	this := ProofOfAddressBreakdownImageIntegrity{}
	return &this
}

// NewProofOfAddressBreakdownImageIntegrityWithDefaults instantiates a new ProofOfAddressBreakdownImageIntegrity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProofOfAddressBreakdownImageIntegrityWithDefaults() *ProofOfAddressBreakdownImageIntegrity {
	this := ProofOfAddressBreakdownImageIntegrity{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *ProofOfAddressBreakdownImageIntegrity) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressBreakdownImageIntegrity) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *ProofOfAddressBreakdownImageIntegrity) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *ProofOfAddressBreakdownImageIntegrity) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *ProofOfAddressBreakdownImageIntegrity) GetBreakdown() ProofOfAddressBreakdownImageIntegrityBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret ProofOfAddressBreakdownImageIntegrityBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfAddressBreakdownImageIntegrity) GetBreakdownOk() (*ProofOfAddressBreakdownImageIntegrityBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *ProofOfAddressBreakdownImageIntegrity) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given ProofOfAddressBreakdownImageIntegrityBreakdown and assigns it to the Breakdown field.
func (o *ProofOfAddressBreakdownImageIntegrity) SetBreakdown(v ProofOfAddressBreakdownImageIntegrityBreakdown) {
	o.Breakdown = &v
}

func (o ProofOfAddressBreakdownImageIntegrity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProofOfAddressBreakdownImageIntegrity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableProofOfAddressBreakdownImageIntegrity struct {
	value *ProofOfAddressBreakdownImageIntegrity
	isSet bool
}

func (v NullableProofOfAddressBreakdownImageIntegrity) Get() *ProofOfAddressBreakdownImageIntegrity {
	return v.value
}

func (v *NullableProofOfAddressBreakdownImageIntegrity) Set(val *ProofOfAddressBreakdownImageIntegrity) {
	v.value = val
	v.isSet = true
}

func (v NullableProofOfAddressBreakdownImageIntegrity) IsSet() bool {
	return v.isSet
}

func (v *NullableProofOfAddressBreakdownImageIntegrity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProofOfAddressBreakdownImageIntegrity(val *ProofOfAddressBreakdownImageIntegrity) *NullableProofOfAddressBreakdownImageIntegrity {
	return &NullableProofOfAddressBreakdownImageIntegrity{value: val, isSet: true}
}

func (v NullableProofOfAddressBreakdownImageIntegrity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProofOfAddressBreakdownImageIntegrity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


