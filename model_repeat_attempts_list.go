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

// checks if the RepeatAttemptsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RepeatAttemptsList{}

// RepeatAttemptsList struct for RepeatAttemptsList
type RepeatAttemptsList struct {
	// The unique identifier of the completed Document report.
	ReportId *string `json:"report_id,omitempty"`
	// An array of repeat attempt objects. If no repeat attempts were found, the array will be empty. The number of objects returned can increase over time if more matches are received. 
	RepeatAttempts []RepeatAttemptsListRepeatAttemptsInner `json:"repeat_attempts"`
	// The total number of attempts using the same document, including the current report under assessment.
	AttemptsCount *int32 `json:"attempts_count,omitempty"`
	// A number between 0 and 1 which indicates the proportion of attempts that have been cleared, including the current report under assessment.
	AttemptsClearRate *float32 `json:"attempts_clear_rate,omitempty"`
	// The number of unique entries in the repeat_attempts field for which at least one of the fields is a mismatch.
	UniqueMismatchesCount *int32 `json:"unique_mismatches_count,omitempty"`
}

type _RepeatAttemptsList RepeatAttemptsList

// NewRepeatAttemptsList instantiates a new RepeatAttemptsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepeatAttemptsList(repeatAttempts []RepeatAttemptsListRepeatAttemptsInner) *RepeatAttemptsList {
	this := RepeatAttemptsList{}
	this.RepeatAttempts = repeatAttempts
	return &this
}

// NewRepeatAttemptsListWithDefaults instantiates a new RepeatAttemptsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepeatAttemptsListWithDefaults() *RepeatAttemptsList {
	this := RepeatAttemptsList{}
	return &this
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *RepeatAttemptsList) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepeatAttemptsList) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *RepeatAttemptsList) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *RepeatAttemptsList) SetReportId(v string) {
	o.ReportId = &v
}

// GetRepeatAttempts returns the RepeatAttempts field value
func (o *RepeatAttemptsList) GetRepeatAttempts() []RepeatAttemptsListRepeatAttemptsInner {
	if o == nil {
		var ret []RepeatAttemptsListRepeatAttemptsInner
		return ret
	}

	return o.RepeatAttempts
}

// GetRepeatAttemptsOk returns a tuple with the RepeatAttempts field value
// and a boolean to check if the value has been set.
func (o *RepeatAttemptsList) GetRepeatAttemptsOk() ([]RepeatAttemptsListRepeatAttemptsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepeatAttempts, true
}

// SetRepeatAttempts sets field value
func (o *RepeatAttemptsList) SetRepeatAttempts(v []RepeatAttemptsListRepeatAttemptsInner) {
	o.RepeatAttempts = v
}

// GetAttemptsCount returns the AttemptsCount field value if set, zero value otherwise.
func (o *RepeatAttemptsList) GetAttemptsCount() int32 {
	if o == nil || IsNil(o.AttemptsCount) {
		var ret int32
		return ret
	}
	return *o.AttemptsCount
}

// GetAttemptsCountOk returns a tuple with the AttemptsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepeatAttemptsList) GetAttemptsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AttemptsCount) {
		return nil, false
	}
	return o.AttemptsCount, true
}

// HasAttemptsCount returns a boolean if a field has been set.
func (o *RepeatAttemptsList) HasAttemptsCount() bool {
	if o != nil && !IsNil(o.AttemptsCount) {
		return true
	}

	return false
}

// SetAttemptsCount gets a reference to the given int32 and assigns it to the AttemptsCount field.
func (o *RepeatAttemptsList) SetAttemptsCount(v int32) {
	o.AttemptsCount = &v
}

// GetAttemptsClearRate returns the AttemptsClearRate field value if set, zero value otherwise.
func (o *RepeatAttemptsList) GetAttemptsClearRate() float32 {
	if o == nil || IsNil(o.AttemptsClearRate) {
		var ret float32
		return ret
	}
	return *o.AttemptsClearRate
}

// GetAttemptsClearRateOk returns a tuple with the AttemptsClearRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepeatAttemptsList) GetAttemptsClearRateOk() (*float32, bool) {
	if o == nil || IsNil(o.AttemptsClearRate) {
		return nil, false
	}
	return o.AttemptsClearRate, true
}

// HasAttemptsClearRate returns a boolean if a field has been set.
func (o *RepeatAttemptsList) HasAttemptsClearRate() bool {
	if o != nil && !IsNil(o.AttemptsClearRate) {
		return true
	}

	return false
}

// SetAttemptsClearRate gets a reference to the given float32 and assigns it to the AttemptsClearRate field.
func (o *RepeatAttemptsList) SetAttemptsClearRate(v float32) {
	o.AttemptsClearRate = &v
}

// GetUniqueMismatchesCount returns the UniqueMismatchesCount field value if set, zero value otherwise.
func (o *RepeatAttemptsList) GetUniqueMismatchesCount() int32 {
	if o == nil || IsNil(o.UniqueMismatchesCount) {
		var ret int32
		return ret
	}
	return *o.UniqueMismatchesCount
}

// GetUniqueMismatchesCountOk returns a tuple with the UniqueMismatchesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RepeatAttemptsList) GetUniqueMismatchesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UniqueMismatchesCount) {
		return nil, false
	}
	return o.UniqueMismatchesCount, true
}

// HasUniqueMismatchesCount returns a boolean if a field has been set.
func (o *RepeatAttemptsList) HasUniqueMismatchesCount() bool {
	if o != nil && !IsNil(o.UniqueMismatchesCount) {
		return true
	}

	return false
}

// SetUniqueMismatchesCount gets a reference to the given int32 and assigns it to the UniqueMismatchesCount field.
func (o *RepeatAttemptsList) SetUniqueMismatchesCount(v int32) {
	o.UniqueMismatchesCount = &v
}

func (o RepeatAttemptsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RepeatAttemptsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportId) {
		toSerialize["report_id"] = o.ReportId
	}
	toSerialize["repeat_attempts"] = o.RepeatAttempts
	if !IsNil(o.AttemptsCount) {
		toSerialize["attempts_count"] = o.AttemptsCount
	}
	if !IsNil(o.AttemptsClearRate) {
		toSerialize["attempts_clear_rate"] = o.AttemptsClearRate
	}
	if !IsNil(o.UniqueMismatchesCount) {
		toSerialize["unique_mismatches_count"] = o.UniqueMismatchesCount
	}
	return toSerialize, nil
}

func (o *RepeatAttemptsList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"repeat_attempts",
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

	varRepeatAttemptsList := _RepeatAttemptsList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRepeatAttemptsList)

	if err != nil {
		return err
	}

	*o = RepeatAttemptsList(varRepeatAttemptsList)

	return err
}

type NullableRepeatAttemptsList struct {
	value *RepeatAttemptsList
	isSet bool
}

func (v NullableRepeatAttemptsList) Get() *RepeatAttemptsList {
	return v.value
}

func (v *NullableRepeatAttemptsList) Set(val *RepeatAttemptsList) {
	v.value = val
	v.isSet = true
}

func (v NullableRepeatAttemptsList) IsSet() bool {
	return v.isSet
}

func (v *NullableRepeatAttemptsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepeatAttemptsList(val *RepeatAttemptsList) *NullableRepeatAttemptsList {
	return &NullableRepeatAttemptsList{value: val, isSet: true}
}

func (v NullableRepeatAttemptsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepeatAttemptsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


