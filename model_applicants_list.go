/*
Onfido API v3.6

The Onfido API (v3.6)

API version: v3.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApplicantsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicantsList{}

// ApplicantsList struct for ApplicantsList
type ApplicantsList struct {
	Applicants []Applicant `json:"applicants"`
}

type _ApplicantsList ApplicantsList

// NewApplicantsList instantiates a new ApplicantsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicantsList(applicants []Applicant) *ApplicantsList {
	this := ApplicantsList{}
	this.Applicants = applicants
	return &this
}

// NewApplicantsListWithDefaults instantiates a new ApplicantsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicantsListWithDefaults() *ApplicantsList {
	this := ApplicantsList{}
	return &this
}

// GetApplicants returns the Applicants field value
func (o *ApplicantsList) GetApplicants() []Applicant {
	if o == nil {
		var ret []Applicant
		return ret
	}

	return o.Applicants
}

// GetApplicantsOk returns a tuple with the Applicants field value
// and a boolean to check if the value has been set.
func (o *ApplicantsList) GetApplicantsOk() ([]Applicant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Applicants, true
}

// SetApplicants sets field value
func (o *ApplicantsList) SetApplicants(v []Applicant) {
	o.Applicants = v
}

func (o ApplicantsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicantsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicants"] = o.Applicants
	return toSerialize, nil
}

func (o *ApplicantsList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicants",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicantsList := _ApplicantsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicantsList)

	if err != nil {
		return err
	}

	*o = ApplicantsList(varApplicantsList)

	return err
}

type NullableApplicantsList struct {
	value *ApplicantsList
	isSet bool
}

func (v NullableApplicantsList) Get() *ApplicantsList {
	return v.value
}

func (v *NullableApplicantsList) Set(val *ApplicantsList) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicantsList) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicantsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicantsList(val *ApplicantsList) *NullableApplicantsList {
	return &NullableApplicantsList{value: val, isSet: true}
}

func (v NullableApplicantsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicantsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


