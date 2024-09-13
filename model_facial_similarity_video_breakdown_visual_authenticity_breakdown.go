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

// checks if the FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown{}

// FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown struct for FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown
type FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown struct {
	LivenessDetected *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownLivenessDetected `json:"liveness_detected,omitempty"`
	SpoofingDetection *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection `json:"spoofing_detection,omitempty"`
}

// NewFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown instantiates a new FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown() *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown {
	this := FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown{}
	return &this
}

// NewFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownWithDefaults instantiates a new FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacialSimilarityVideoBreakdownVisualAuthenticityBreakdownWithDefaults() *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown {
	this := FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown{}
	return &this
}

// GetLivenessDetected returns the LivenessDetected field value if set, zero value otherwise.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) GetLivenessDetected() FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownLivenessDetected {
	if o == nil || IsNil(o.LivenessDetected) {
		var ret FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownLivenessDetected
		return ret
	}
	return *o.LivenessDetected
}

// GetLivenessDetectedOk returns a tuple with the LivenessDetected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) GetLivenessDetectedOk() (*FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownLivenessDetected, bool) {
	if o == nil || IsNil(o.LivenessDetected) {
		return nil, false
	}
	return o.LivenessDetected, true
}

// HasLivenessDetected returns a boolean if a field has been set.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) HasLivenessDetected() bool {
	if o != nil && !IsNil(o.LivenessDetected) {
		return true
	}

	return false
}

// SetLivenessDetected gets a reference to the given FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownLivenessDetected and assigns it to the LivenessDetected field.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) SetLivenessDetected(v FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownLivenessDetected) {
	o.LivenessDetected = &v
}

// GetSpoofingDetection returns the SpoofingDetection field value if set, zero value otherwise.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) GetSpoofingDetection() FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection {
	if o == nil || IsNil(o.SpoofingDetection) {
		var ret FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection
		return ret
	}
	return *o.SpoofingDetection
}

// GetSpoofingDetectionOk returns a tuple with the SpoofingDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) GetSpoofingDetectionOk() (*FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection, bool) {
	if o == nil || IsNil(o.SpoofingDetection) {
		return nil, false
	}
	return o.SpoofingDetection, true
}

// HasSpoofingDetection returns a boolean if a field has been set.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) HasSpoofingDetection() bool {
	if o != nil && !IsNil(o.SpoofingDetection) {
		return true
	}

	return false
}

// SetSpoofingDetection gets a reference to the given FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection and assigns it to the SpoofingDetection field.
func (o *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) SetSpoofingDetection(v FacialSimilarityVideoBreakdownVisualAuthenticityBreakdownSpoofingDetection) {
	o.SpoofingDetection = &v
}

func (o FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LivenessDetected) {
		toSerialize["liveness_detected"] = o.LivenessDetected
	}
	if !IsNil(o.SpoofingDetection) {
		toSerialize["spoofing_detection"] = o.SpoofingDetection
	}
	return toSerialize, nil
}

type NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown struct {
	value *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown
	isSet bool
}

func (v NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) Get() *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown {
	return v.value
}

func (v *NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) Set(val *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown(val *FacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) *NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown {
	return &NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown{value: val, isSet: true}
}

func (v NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacialSimilarityVideoBreakdownVisualAuthenticityBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


