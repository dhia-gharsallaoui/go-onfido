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

// checks if the FacialSimilarityVideoBreakdownImageIntegrity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacialSimilarityVideoBreakdownImageIntegrity{}

// FacialSimilarityVideoBreakdownImageIntegrity Asserts whether the quality and integrity of the uploaded files were sufficient to perform a face comparison.
type FacialSimilarityVideoBreakdownImageIntegrity struct {
	Result *string `json:"result,omitempty"`
	Breakdown *FacialSimilarityVideoBreakdownImageIntegrityBreakdown `json:"breakdown,omitempty"`
}

// NewFacialSimilarityVideoBreakdownImageIntegrity instantiates a new FacialSimilarityVideoBreakdownImageIntegrity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacialSimilarityVideoBreakdownImageIntegrity() *FacialSimilarityVideoBreakdownImageIntegrity {
	this := FacialSimilarityVideoBreakdownImageIntegrity{}
	return &this
}

// NewFacialSimilarityVideoBreakdownImageIntegrityWithDefaults instantiates a new FacialSimilarityVideoBreakdownImageIntegrity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacialSimilarityVideoBreakdownImageIntegrityWithDefaults() *FacialSimilarityVideoBreakdownImageIntegrity {
	this := FacialSimilarityVideoBreakdownImageIntegrity{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FacialSimilarityVideoBreakdownImageIntegrity) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityVideoBreakdownImageIntegrity) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FacialSimilarityVideoBreakdownImageIntegrity) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *FacialSimilarityVideoBreakdownImageIntegrity) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *FacialSimilarityVideoBreakdownImageIntegrity) GetBreakdown() FacialSimilarityVideoBreakdownImageIntegrityBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret FacialSimilarityVideoBreakdownImageIntegrityBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityVideoBreakdownImageIntegrity) GetBreakdownOk() (*FacialSimilarityVideoBreakdownImageIntegrityBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *FacialSimilarityVideoBreakdownImageIntegrity) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given FacialSimilarityVideoBreakdownImageIntegrityBreakdown and assigns it to the Breakdown field.
func (o *FacialSimilarityVideoBreakdownImageIntegrity) SetBreakdown(v FacialSimilarityVideoBreakdownImageIntegrityBreakdown) {
	o.Breakdown = &v
}

func (o FacialSimilarityVideoBreakdownImageIntegrity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacialSimilarityVideoBreakdownImageIntegrity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableFacialSimilarityVideoBreakdownImageIntegrity struct {
	value *FacialSimilarityVideoBreakdownImageIntegrity
	isSet bool
}

func (v NullableFacialSimilarityVideoBreakdownImageIntegrity) Get() *FacialSimilarityVideoBreakdownImageIntegrity {
	return v.value
}

func (v *NullableFacialSimilarityVideoBreakdownImageIntegrity) Set(val *FacialSimilarityVideoBreakdownImageIntegrity) {
	v.value = val
	v.isSet = true
}

func (v NullableFacialSimilarityVideoBreakdownImageIntegrity) IsSet() bool {
	return v.isSet
}

func (v *NullableFacialSimilarityVideoBreakdownImageIntegrity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacialSimilarityVideoBreakdownImageIntegrity(val *FacialSimilarityVideoBreakdownImageIntegrity) *NullableFacialSimilarityVideoBreakdownImageIntegrity {
	return &NullableFacialSimilarityVideoBreakdownImageIntegrity{value: val, isSet: true}
}

func (v NullableFacialSimilarityVideoBreakdownImageIntegrity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacialSimilarityVideoBreakdownImageIntegrity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


