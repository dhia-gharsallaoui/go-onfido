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

// checks if the FacialSimilarityVideoProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacialSimilarityVideoProperties{}

// FacialSimilarityVideoProperties struct for FacialSimilarityVideoProperties
type FacialSimilarityVideoProperties struct {
	// A floating point number between 0 and 1. The closer the score is to 0, the more likely it is to be a spoof (i.e. videos of digital screens, masks or print-outs). Conversely, the closer it is to 1, the less likely it is to be a spoof. 
	Score *float32 `json:"score,omitempty"`
}

// NewFacialSimilarityVideoProperties instantiates a new FacialSimilarityVideoProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacialSimilarityVideoProperties() *FacialSimilarityVideoProperties {
	this := FacialSimilarityVideoProperties{}
	return &this
}

// NewFacialSimilarityVideoPropertiesWithDefaults instantiates a new FacialSimilarityVideoProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacialSimilarityVideoPropertiesWithDefaults() *FacialSimilarityVideoProperties {
	this := FacialSimilarityVideoProperties{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *FacialSimilarityVideoProperties) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityVideoProperties) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *FacialSimilarityVideoProperties) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *FacialSimilarityVideoProperties) SetScore(v float32) {
	o.Score = &v
}

func (o FacialSimilarityVideoProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacialSimilarityVideoProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	return toSerialize, nil
}

type NullableFacialSimilarityVideoProperties struct {
	value *FacialSimilarityVideoProperties
	isSet bool
}

func (v NullableFacialSimilarityVideoProperties) Get() *FacialSimilarityVideoProperties {
	return v.value
}

func (v *NullableFacialSimilarityVideoProperties) Set(val *FacialSimilarityVideoProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableFacialSimilarityVideoProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableFacialSimilarityVideoProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacialSimilarityVideoProperties(val *FacialSimilarityVideoProperties) *NullableFacialSimilarityVideoProperties {
	return &NullableFacialSimilarityVideoProperties{value: val, isSet: true}
}

func (v NullableFacialSimilarityVideoProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacialSimilarityVideoProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


