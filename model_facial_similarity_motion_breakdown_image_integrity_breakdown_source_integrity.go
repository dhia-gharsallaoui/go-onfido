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

// checks if the FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity{}

// FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity Asserts whether the motion capture is trustworthy - e.g. not from a fake webcam.
type FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity struct {
	Result *string `json:"result,omitempty"`
	Properties *VideoReasons `json:"properties,omitempty"`
}

// NewFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity instantiates a new FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity() *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity {
	this := FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity{}
	return &this
}

// NewFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrityWithDefaults instantiates a new FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrityWithDefaults() *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity {
	this := FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) SetResult(v string) {
	o.Result = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) GetProperties() VideoReasons {
	if o == nil || IsNil(o.Properties) {
		var ret VideoReasons
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) GetPropertiesOk() (*VideoReasons, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given VideoReasons and assigns it to the Properties field.
func (o *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) SetProperties(v VideoReasons) {
	o.Properties = &v
}

func (o FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity struct {
	value *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity
	isSet bool
}

func (v NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) Get() *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity {
	return v.value
}

func (v *NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) Set(val *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) {
	v.value = val
	v.isSet = true
}

func (v NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) IsSet() bool {
	return v.isSet
}

func (v *NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity(val *FacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) *NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity {
	return &NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity{value: val, isSet: true}
}

func (v NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacialSimilarityMotionBreakdownImageIntegrityBreakdownSourceIntegrity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


