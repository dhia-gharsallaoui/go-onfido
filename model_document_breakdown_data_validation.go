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

// checks if the DocumentBreakdownDataValidation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentBreakdownDataValidation{}

// DocumentBreakdownDataValidation Asserts whether algorithmically validatable elements are correct.
type DocumentBreakdownDataValidation struct {
	Result *string `json:"result,omitempty"`
	Breakdown *DocumentBreakdownDataValidationBreakdown `json:"breakdown,omitempty"`
}

// NewDocumentBreakdownDataValidation instantiates a new DocumentBreakdownDataValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentBreakdownDataValidation() *DocumentBreakdownDataValidation {
	this := DocumentBreakdownDataValidation{}
	return &this
}

// NewDocumentBreakdownDataValidationWithDefaults instantiates a new DocumentBreakdownDataValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentBreakdownDataValidationWithDefaults() *DocumentBreakdownDataValidation {
	this := DocumentBreakdownDataValidation{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DocumentBreakdownDataValidation) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownDataValidation) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DocumentBreakdownDataValidation) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *DocumentBreakdownDataValidation) SetResult(v string) {
	o.Result = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *DocumentBreakdownDataValidation) GetBreakdown() DocumentBreakdownDataValidationBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret DocumentBreakdownDataValidationBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentBreakdownDataValidation) GetBreakdownOk() (*DocumentBreakdownDataValidationBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *DocumentBreakdownDataValidation) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given DocumentBreakdownDataValidationBreakdown and assigns it to the Breakdown field.
func (o *DocumentBreakdownDataValidation) SetBreakdown(v DocumentBreakdownDataValidationBreakdown) {
	o.Breakdown = &v
}

func (o DocumentBreakdownDataValidation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentBreakdownDataValidation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	return toSerialize, nil
}

type NullableDocumentBreakdownDataValidation struct {
	value *DocumentBreakdownDataValidation
	isSet bool
}

func (v NullableDocumentBreakdownDataValidation) Get() *DocumentBreakdownDataValidation {
	return v.value
}

func (v *NullableDocumentBreakdownDataValidation) Set(val *DocumentBreakdownDataValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentBreakdownDataValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentBreakdownDataValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentBreakdownDataValidation(val *DocumentBreakdownDataValidation) *NullableDocumentBreakdownDataValidation {
	return &NullableDocumentBreakdownDataValidation{value: val, isSet: true}
}

func (v NullableDocumentBreakdownDataValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentBreakdownDataValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


