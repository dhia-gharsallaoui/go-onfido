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

// checks if the DocumentPropertiesNfc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentPropertiesNfc{}

// DocumentPropertiesNfc struct for DocumentPropertiesNfc
type DocumentPropertiesNfc struct {
	DocumentType *string `json:"document_type,omitempty"`
	IssuingCountry *string `json:"issuing_country,omitempty"`
	FullName *string `json:"full_name,omitempty"`
	DocumentNumber *string `json:"document_number,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	Gender *string `json:"gender,omitempty"`
	DateOfExpiry *string `json:"date_of_expiry,omitempty"`
	PersonalNumber *string `json:"personal_number,omitempty"`
	PlaceOfBirth *string `json:"place_of_birth,omitempty"`
	Address *string `json:"address,omitempty"`
	IssuingDate *string `json:"issuing_date,omitempty"`
	IssuingAuthority *string `json:"issuing_authority,omitempty"`
}

// NewDocumentPropertiesNfc instantiates a new DocumentPropertiesNfc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentPropertiesNfc() *DocumentPropertiesNfc {
	this := DocumentPropertiesNfc{}
	return &this
}

// NewDocumentPropertiesNfcWithDefaults instantiates a new DocumentPropertiesNfc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentPropertiesNfcWithDefaults() *DocumentPropertiesNfc {
	this := DocumentPropertiesNfc{}
	return &this
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetDocumentType() string {
	if o == nil || IsNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetDocumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasDocumentType() bool {
	if o != nil && !IsNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *DocumentPropertiesNfc) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetIssuingCountry returns the IssuingCountry field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetIssuingCountry() string {
	if o == nil || IsNil(o.IssuingCountry) {
		var ret string
		return ret
	}
	return *o.IssuingCountry
}

// GetIssuingCountryOk returns a tuple with the IssuingCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetIssuingCountryOk() (*string, bool) {
	if o == nil || IsNil(o.IssuingCountry) {
		return nil, false
	}
	return o.IssuingCountry, true
}

// HasIssuingCountry returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasIssuingCountry() bool {
	if o != nil && !IsNil(o.IssuingCountry) {
		return true
	}

	return false
}

// SetIssuingCountry gets a reference to the given string and assigns it to the IssuingCountry field.
func (o *DocumentPropertiesNfc) SetIssuingCountry(v string) {
	o.IssuingCountry = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *DocumentPropertiesNfc) SetFullName(v string) {
	o.FullName = &v
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetDocumentNumber() string {
	if o == nil || IsNil(o.DocumentNumber) {
		var ret string
		return ret
	}
	return *o.DocumentNumber
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetDocumentNumberOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentNumber) {
		return nil, false
	}
	return o.DocumentNumber, true
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasDocumentNumber() bool {
	if o != nil && !IsNil(o.DocumentNumber) {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given string and assigns it to the DocumentNumber field.
func (o *DocumentPropertiesNfc) SetDocumentNumber(v string) {
	o.DocumentNumber = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *DocumentPropertiesNfc) SetNationality(v string) {
	o.Nationality = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetDateOfBirth() string {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetDateOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *DocumentPropertiesNfc) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *DocumentPropertiesNfc) SetGender(v string) {
	o.Gender = &v
}

// GetDateOfExpiry returns the DateOfExpiry field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetDateOfExpiry() string {
	if o == nil || IsNil(o.DateOfExpiry) {
		var ret string
		return ret
	}
	return *o.DateOfExpiry
}

// GetDateOfExpiryOk returns a tuple with the DateOfExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetDateOfExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfExpiry) {
		return nil, false
	}
	return o.DateOfExpiry, true
}

// HasDateOfExpiry returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasDateOfExpiry() bool {
	if o != nil && !IsNil(o.DateOfExpiry) {
		return true
	}

	return false
}

// SetDateOfExpiry gets a reference to the given string and assigns it to the DateOfExpiry field.
func (o *DocumentPropertiesNfc) SetDateOfExpiry(v string) {
	o.DateOfExpiry = &v
}

// GetPersonalNumber returns the PersonalNumber field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetPersonalNumber() string {
	if o == nil || IsNil(o.PersonalNumber) {
		var ret string
		return ret
	}
	return *o.PersonalNumber
}

// GetPersonalNumberOk returns a tuple with the PersonalNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetPersonalNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PersonalNumber) {
		return nil, false
	}
	return o.PersonalNumber, true
}

// HasPersonalNumber returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasPersonalNumber() bool {
	if o != nil && !IsNil(o.PersonalNumber) {
		return true
	}

	return false
}

// SetPersonalNumber gets a reference to the given string and assigns it to the PersonalNumber field.
func (o *DocumentPropertiesNfc) SetPersonalNumber(v string) {
	o.PersonalNumber = &v
}

// GetPlaceOfBirth returns the PlaceOfBirth field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetPlaceOfBirth() string {
	if o == nil || IsNil(o.PlaceOfBirth) {
		var ret string
		return ret
	}
	return *o.PlaceOfBirth
}

// GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetPlaceOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceOfBirth) {
		return nil, false
	}
	return o.PlaceOfBirth, true
}

// HasPlaceOfBirth returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasPlaceOfBirth() bool {
	if o != nil && !IsNil(o.PlaceOfBirth) {
		return true
	}

	return false
}

// SetPlaceOfBirth gets a reference to the given string and assigns it to the PlaceOfBirth field.
func (o *DocumentPropertiesNfc) SetPlaceOfBirth(v string) {
	o.PlaceOfBirth = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DocumentPropertiesNfc) SetAddress(v string) {
	o.Address = &v
}

// GetIssuingDate returns the IssuingDate field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetIssuingDate() string {
	if o == nil || IsNil(o.IssuingDate) {
		var ret string
		return ret
	}
	return *o.IssuingDate
}

// GetIssuingDateOk returns a tuple with the IssuingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetIssuingDateOk() (*string, bool) {
	if o == nil || IsNil(o.IssuingDate) {
		return nil, false
	}
	return o.IssuingDate, true
}

// HasIssuingDate returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasIssuingDate() bool {
	if o != nil && !IsNil(o.IssuingDate) {
		return true
	}

	return false
}

// SetIssuingDate gets a reference to the given string and assigns it to the IssuingDate field.
func (o *DocumentPropertiesNfc) SetIssuingDate(v string) {
	o.IssuingDate = &v
}

// GetIssuingAuthority returns the IssuingAuthority field value if set, zero value otherwise.
func (o *DocumentPropertiesNfc) GetIssuingAuthority() string {
	if o == nil || IsNil(o.IssuingAuthority) {
		var ret string
		return ret
	}
	return *o.IssuingAuthority
}

// GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesNfc) GetIssuingAuthorityOk() (*string, bool) {
	if o == nil || IsNil(o.IssuingAuthority) {
		return nil, false
	}
	return o.IssuingAuthority, true
}

// HasIssuingAuthority returns a boolean if a field has been set.
func (o *DocumentPropertiesNfc) HasIssuingAuthority() bool {
	if o != nil && !IsNil(o.IssuingAuthority) {
		return true
	}

	return false
}

// SetIssuingAuthority gets a reference to the given string and assigns it to the IssuingAuthority field.
func (o *DocumentPropertiesNfc) SetIssuingAuthority(v string) {
	o.IssuingAuthority = &v
}

func (o DocumentPropertiesNfc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentPropertiesNfc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentType) {
		toSerialize["document_type"] = o.DocumentType
	}
	if !IsNil(o.IssuingCountry) {
		toSerialize["issuing_country"] = o.IssuingCountry
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.DocumentNumber) {
		toSerialize["document_number"] = o.DocumentNumber
	}
	if !IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !IsNil(o.DateOfBirth) {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.DateOfExpiry) {
		toSerialize["date_of_expiry"] = o.DateOfExpiry
	}
	if !IsNil(o.PersonalNumber) {
		toSerialize["personal_number"] = o.PersonalNumber
	}
	if !IsNil(o.PlaceOfBirth) {
		toSerialize["place_of_birth"] = o.PlaceOfBirth
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.IssuingDate) {
		toSerialize["issuing_date"] = o.IssuingDate
	}
	if !IsNil(o.IssuingAuthority) {
		toSerialize["issuing_authority"] = o.IssuingAuthority
	}
	return toSerialize, nil
}

type NullableDocumentPropertiesNfc struct {
	value *DocumentPropertiesNfc
	isSet bool
}

func (v NullableDocumentPropertiesNfc) Get() *DocumentPropertiesNfc {
	return v.value
}

func (v *NullableDocumentPropertiesNfc) Set(val *DocumentPropertiesNfc) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentPropertiesNfc) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentPropertiesNfc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentPropertiesNfc(val *DocumentPropertiesNfc) *NullableDocumentPropertiesNfc {
	return &NullableDocumentPropertiesNfc{value: val, isSet: true}
}

func (v NullableDocumentPropertiesNfc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentPropertiesNfc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


