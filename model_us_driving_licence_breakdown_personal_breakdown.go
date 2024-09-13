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

// checks if the UsDrivingLicenceBreakdownPersonalBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsDrivingLicenceBreakdownPersonalBreakdown{}

// UsDrivingLicenceBreakdownPersonalBreakdown struct for UsDrivingLicenceBreakdownPersonalBreakdown
type UsDrivingLicenceBreakdownPersonalBreakdown struct {
	FirstName *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"first_name,omitempty"`
	NameSuffix *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"name_suffix,omitempty"`
	Height *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"height,omitempty"`
	Weight *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"weight,omitempty"`
	SexCode *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"sex_code,omitempty"`
	EyeColor *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"eye_color,omitempty"`
	DateOfBirth *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"date_of_birth,omitempty"`
	LastName *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"last_name,omitempty"`
	MiddleName *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"middle_name,omitempty"`
	FirstNameFuzzy *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"first_name_fuzzy,omitempty"`
	MiddleNameFuzzy *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"middle_name_fuzzy,omitempty"`
	LastNameFuzzy *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"last_name_fuzzy,omitempty"`
	MiddleInitial *DocumentBreakdownDataComparisonBreakdownIssuingCountry `json:"middle_initial,omitempty"`
}

// NewUsDrivingLicenceBreakdownPersonalBreakdown instantiates a new UsDrivingLicenceBreakdownPersonalBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsDrivingLicenceBreakdownPersonalBreakdown() *UsDrivingLicenceBreakdownPersonalBreakdown {
	this := UsDrivingLicenceBreakdownPersonalBreakdown{}
	return &this
}

// NewUsDrivingLicenceBreakdownPersonalBreakdownWithDefaults instantiates a new UsDrivingLicenceBreakdownPersonalBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsDrivingLicenceBreakdownPersonalBreakdownWithDefaults() *UsDrivingLicenceBreakdownPersonalBreakdown {
	this := UsDrivingLicenceBreakdownPersonalBreakdown{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetFirstName() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.FirstName) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetFirstNameOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the FirstName field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetFirstName(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.FirstName = &v
}

// GetNameSuffix returns the NameSuffix field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetNameSuffix() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.NameSuffix) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.NameSuffix
}

// GetNameSuffixOk returns a tuple with the NameSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetNameSuffixOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.NameSuffix) {
		return nil, false
	}
	return o.NameSuffix, true
}

// HasNameSuffix returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasNameSuffix() bool {
	if o != nil && !IsNil(o.NameSuffix) {
		return true
	}

	return false
}

// SetNameSuffix gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the NameSuffix field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetNameSuffix(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.NameSuffix = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetHeight() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.Height) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetHeightOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the Height field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetHeight(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.Height = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetWeight() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.Weight) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetWeightOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the Weight field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetWeight(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.Weight = &v
}

// GetSexCode returns the SexCode field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetSexCode() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.SexCode) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.SexCode
}

// GetSexCodeOk returns a tuple with the SexCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetSexCodeOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.SexCode) {
		return nil, false
	}
	return o.SexCode, true
}

// HasSexCode returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasSexCode() bool {
	if o != nil && !IsNil(o.SexCode) {
		return true
	}

	return false
}

// SetSexCode gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the SexCode field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetSexCode(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.SexCode = &v
}

// GetEyeColor returns the EyeColor field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetEyeColor() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.EyeColor) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.EyeColor
}

// GetEyeColorOk returns a tuple with the EyeColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetEyeColorOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.EyeColor) {
		return nil, false
	}
	return o.EyeColor, true
}

// HasEyeColor returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasEyeColor() bool {
	if o != nil && !IsNil(o.EyeColor) {
		return true
	}

	return false
}

// SetEyeColor gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the EyeColor field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetEyeColor(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.EyeColor = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetDateOfBirth() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetDateOfBirthOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the DateOfBirth field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetDateOfBirth(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.DateOfBirth = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetLastName() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.LastName) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetLastNameOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the LastName field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetLastName(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.LastName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetMiddleName() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.MiddleName) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetMiddleNameOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the MiddleName field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetMiddleName(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.MiddleName = &v
}

// GetFirstNameFuzzy returns the FirstNameFuzzy field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetFirstNameFuzzy() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.FirstNameFuzzy) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.FirstNameFuzzy
}

// GetFirstNameFuzzyOk returns a tuple with the FirstNameFuzzy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetFirstNameFuzzyOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.FirstNameFuzzy) {
		return nil, false
	}
	return o.FirstNameFuzzy, true
}

// HasFirstNameFuzzy returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasFirstNameFuzzy() bool {
	if o != nil && !IsNil(o.FirstNameFuzzy) {
		return true
	}

	return false
}

// SetFirstNameFuzzy gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the FirstNameFuzzy field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetFirstNameFuzzy(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.FirstNameFuzzy = &v
}

// GetMiddleNameFuzzy returns the MiddleNameFuzzy field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetMiddleNameFuzzy() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.MiddleNameFuzzy) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.MiddleNameFuzzy
}

// GetMiddleNameFuzzyOk returns a tuple with the MiddleNameFuzzy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetMiddleNameFuzzyOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.MiddleNameFuzzy) {
		return nil, false
	}
	return o.MiddleNameFuzzy, true
}

// HasMiddleNameFuzzy returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasMiddleNameFuzzy() bool {
	if o != nil && !IsNil(o.MiddleNameFuzzy) {
		return true
	}

	return false
}

// SetMiddleNameFuzzy gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the MiddleNameFuzzy field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetMiddleNameFuzzy(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.MiddleNameFuzzy = &v
}

// GetLastNameFuzzy returns the LastNameFuzzy field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetLastNameFuzzy() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.LastNameFuzzy) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.LastNameFuzzy
}

// GetLastNameFuzzyOk returns a tuple with the LastNameFuzzy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetLastNameFuzzyOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.LastNameFuzzy) {
		return nil, false
	}
	return o.LastNameFuzzy, true
}

// HasLastNameFuzzy returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasLastNameFuzzy() bool {
	if o != nil && !IsNil(o.LastNameFuzzy) {
		return true
	}

	return false
}

// SetLastNameFuzzy gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the LastNameFuzzy field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetLastNameFuzzy(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.LastNameFuzzy = &v
}

// GetMiddleInitial returns the MiddleInitial field value if set, zero value otherwise.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetMiddleInitial() DocumentBreakdownDataComparisonBreakdownIssuingCountry {
	if o == nil || IsNil(o.MiddleInitial) {
		var ret DocumentBreakdownDataComparisonBreakdownIssuingCountry
		return ret
	}
	return *o.MiddleInitial
}

// GetMiddleInitialOk returns a tuple with the MiddleInitial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) GetMiddleInitialOk() (*DocumentBreakdownDataComparisonBreakdownIssuingCountry, bool) {
	if o == nil || IsNil(o.MiddleInitial) {
		return nil, false
	}
	return o.MiddleInitial, true
}

// HasMiddleInitial returns a boolean if a field has been set.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) HasMiddleInitial() bool {
	if o != nil && !IsNil(o.MiddleInitial) {
		return true
	}

	return false
}

// SetMiddleInitial gets a reference to the given DocumentBreakdownDataComparisonBreakdownIssuingCountry and assigns it to the MiddleInitial field.
func (o *UsDrivingLicenceBreakdownPersonalBreakdown) SetMiddleInitial(v DocumentBreakdownDataComparisonBreakdownIssuingCountry) {
	o.MiddleInitial = &v
}

func (o UsDrivingLicenceBreakdownPersonalBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsDrivingLicenceBreakdownPersonalBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.NameSuffix) {
		toSerialize["name_suffix"] = o.NameSuffix
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.SexCode) {
		toSerialize["sex_code"] = o.SexCode
	}
	if !IsNil(o.EyeColor) {
		toSerialize["eye_color"] = o.EyeColor
	}
	if !IsNil(o.DateOfBirth) {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.FirstNameFuzzy) {
		toSerialize["first_name_fuzzy"] = o.FirstNameFuzzy
	}
	if !IsNil(o.MiddleNameFuzzy) {
		toSerialize["middle_name_fuzzy"] = o.MiddleNameFuzzy
	}
	if !IsNil(o.LastNameFuzzy) {
		toSerialize["last_name_fuzzy"] = o.LastNameFuzzy
	}
	if !IsNil(o.MiddleInitial) {
		toSerialize["middle_initial"] = o.MiddleInitial
	}
	return toSerialize, nil
}

type NullableUsDrivingLicenceBreakdownPersonalBreakdown struct {
	value *UsDrivingLicenceBreakdownPersonalBreakdown
	isSet bool
}

func (v NullableUsDrivingLicenceBreakdownPersonalBreakdown) Get() *UsDrivingLicenceBreakdownPersonalBreakdown {
	return v.value
}

func (v *NullableUsDrivingLicenceBreakdownPersonalBreakdown) Set(val *UsDrivingLicenceBreakdownPersonalBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableUsDrivingLicenceBreakdownPersonalBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableUsDrivingLicenceBreakdownPersonalBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsDrivingLicenceBreakdownPersonalBreakdown(val *UsDrivingLicenceBreakdownPersonalBreakdown) *NullableUsDrivingLicenceBreakdownPersonalBreakdown {
	return &NullableUsDrivingLicenceBreakdownPersonalBreakdown{value: val, isSet: true}
}

func (v NullableUsDrivingLicenceBreakdownPersonalBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsDrivingLicenceBreakdownPersonalBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


