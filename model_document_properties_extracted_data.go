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

// checks if the DocumentPropertiesExtractedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentPropertiesExtractedData{}

// DocumentPropertiesExtractedData struct for DocumentPropertiesExtractedData
type DocumentPropertiesExtractedData struct {
	DocumentNumber *string `json:"document_number,omitempty"`
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	DateOfExpiry *string `json:"date_of_expiry,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	MiddleName *string `json:"middle_name,omitempty"`
	FullName *string `json:"full_name,omitempty"`
	SpouseName *string `json:"spouse_name,omitempty"`
	WidowName *string `json:"widow_name,omitempty"`
	AliasName *string `json:"alias_name,omitempty"`
	Gender *string `json:"gender,omitempty"`
	MrzLine1 *string `json:"mrz_line1,omitempty"`
	MrzLine2 *string `json:"mrz_line2,omitempty"`
	MrzLine3 *string `json:"mrz_line3,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
	AddressLine1 *string `json:"address_line_1,omitempty"`
	AddressLine2 *string `json:"address_line_2,omitempty"`
	AddressLine3 *string `json:"address_line_3,omitempty"`
	AddressLine4 *string `json:"address_line_4,omitempty"`
	AddressLine5 *string `json:"address_line_5,omitempty"`
	IssuingAuthority *string `json:"issuing_authority,omitempty"`
}

// NewDocumentPropertiesExtractedData instantiates a new DocumentPropertiesExtractedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentPropertiesExtractedData() *DocumentPropertiesExtractedData {
	this := DocumentPropertiesExtractedData{}
	return &this
}

// NewDocumentPropertiesExtractedDataWithDefaults instantiates a new DocumentPropertiesExtractedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentPropertiesExtractedDataWithDefaults() *DocumentPropertiesExtractedData {
	this := DocumentPropertiesExtractedData{}
	return &this
}

// GetDocumentNumber returns the DocumentNumber field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetDocumentNumber() string {
	if o == nil || IsNil(o.DocumentNumber) {
		var ret string
		return ret
	}
	return *o.DocumentNumber
}

// GetDocumentNumberOk returns a tuple with the DocumentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetDocumentNumberOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentNumber) {
		return nil, false
	}
	return o.DocumentNumber, true
}

// HasDocumentNumber returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasDocumentNumber() bool {
	if o != nil && !IsNil(o.DocumentNumber) {
		return true
	}

	return false
}

// SetDocumentNumber gets a reference to the given string and assigns it to the DocumentNumber field.
func (o *DocumentPropertiesExtractedData) SetDocumentNumber(v string) {
	o.DocumentNumber = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetDateOfBirth() string {
	if o == nil || IsNil(o.DateOfBirth) {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetDateOfBirthOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfBirth) {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasDateOfBirth() bool {
	if o != nil && !IsNil(o.DateOfBirth) {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *DocumentPropertiesExtractedData) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

// GetDateOfExpiry returns the DateOfExpiry field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetDateOfExpiry() string {
	if o == nil || IsNil(o.DateOfExpiry) {
		var ret string
		return ret
	}
	return *o.DateOfExpiry
}

// GetDateOfExpiryOk returns a tuple with the DateOfExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetDateOfExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfExpiry) {
		return nil, false
	}
	return o.DateOfExpiry, true
}

// HasDateOfExpiry returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasDateOfExpiry() bool {
	if o != nil && !IsNil(o.DateOfExpiry) {
		return true
	}

	return false
}

// SetDateOfExpiry gets a reference to the given string and assigns it to the DateOfExpiry field.
func (o *DocumentPropertiesExtractedData) SetDateOfExpiry(v string) {
	o.DateOfExpiry = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *DocumentPropertiesExtractedData) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *DocumentPropertiesExtractedData) SetLastName(v string) {
	o.LastName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *DocumentPropertiesExtractedData) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *DocumentPropertiesExtractedData) SetFullName(v string) {
	o.FullName = &v
}

// GetSpouseName returns the SpouseName field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetSpouseName() string {
	if o == nil || IsNil(o.SpouseName) {
		var ret string
		return ret
	}
	return *o.SpouseName
}

// GetSpouseNameOk returns a tuple with the SpouseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetSpouseNameOk() (*string, bool) {
	if o == nil || IsNil(o.SpouseName) {
		return nil, false
	}
	return o.SpouseName, true
}

// HasSpouseName returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasSpouseName() bool {
	if o != nil && !IsNil(o.SpouseName) {
		return true
	}

	return false
}

// SetSpouseName gets a reference to the given string and assigns it to the SpouseName field.
func (o *DocumentPropertiesExtractedData) SetSpouseName(v string) {
	o.SpouseName = &v
}

// GetWidowName returns the WidowName field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetWidowName() string {
	if o == nil || IsNil(o.WidowName) {
		var ret string
		return ret
	}
	return *o.WidowName
}

// GetWidowNameOk returns a tuple with the WidowName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetWidowNameOk() (*string, bool) {
	if o == nil || IsNil(o.WidowName) {
		return nil, false
	}
	return o.WidowName, true
}

// HasWidowName returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasWidowName() bool {
	if o != nil && !IsNil(o.WidowName) {
		return true
	}

	return false
}

// SetWidowName gets a reference to the given string and assigns it to the WidowName field.
func (o *DocumentPropertiesExtractedData) SetWidowName(v string) {
	o.WidowName = &v
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetAliasName() string {
	if o == nil || IsNil(o.AliasName) {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetAliasNameOk() (*string, bool) {
	if o == nil || IsNil(o.AliasName) {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasAliasName() bool {
	if o != nil && !IsNil(o.AliasName) {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *DocumentPropertiesExtractedData) SetAliasName(v string) {
	o.AliasName = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *DocumentPropertiesExtractedData) SetGender(v string) {
	o.Gender = &v
}

// GetMrzLine1 returns the MrzLine1 field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetMrzLine1() string {
	if o == nil || IsNil(o.MrzLine1) {
		var ret string
		return ret
	}
	return *o.MrzLine1
}

// GetMrzLine1Ok returns a tuple with the MrzLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetMrzLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.MrzLine1) {
		return nil, false
	}
	return o.MrzLine1, true
}

// HasMrzLine1 returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasMrzLine1() bool {
	if o != nil && !IsNil(o.MrzLine1) {
		return true
	}

	return false
}

// SetMrzLine1 gets a reference to the given string and assigns it to the MrzLine1 field.
func (o *DocumentPropertiesExtractedData) SetMrzLine1(v string) {
	o.MrzLine1 = &v
}

// GetMrzLine2 returns the MrzLine2 field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetMrzLine2() string {
	if o == nil || IsNil(o.MrzLine2) {
		var ret string
		return ret
	}
	return *o.MrzLine2
}

// GetMrzLine2Ok returns a tuple with the MrzLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetMrzLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.MrzLine2) {
		return nil, false
	}
	return o.MrzLine2, true
}

// HasMrzLine2 returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasMrzLine2() bool {
	if o != nil && !IsNil(o.MrzLine2) {
		return true
	}

	return false
}

// SetMrzLine2 gets a reference to the given string and assigns it to the MrzLine2 field.
func (o *DocumentPropertiesExtractedData) SetMrzLine2(v string) {
	o.MrzLine2 = &v
}

// GetMrzLine3 returns the MrzLine3 field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetMrzLine3() string {
	if o == nil || IsNil(o.MrzLine3) {
		var ret string
		return ret
	}
	return *o.MrzLine3
}

// GetMrzLine3Ok returns a tuple with the MrzLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetMrzLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.MrzLine3) {
		return nil, false
	}
	return o.MrzLine3, true
}

// HasMrzLine3 returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasMrzLine3() bool {
	if o != nil && !IsNil(o.MrzLine3) {
		return true
	}

	return false
}

// SetMrzLine3 gets a reference to the given string and assigns it to the MrzLine3 field.
func (o *DocumentPropertiesExtractedData) SetMrzLine3(v string) {
	o.MrzLine3 = &v
}

// GetNationality returns the Nationality field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetNationality() string {
	if o == nil || IsNil(o.Nationality) {
		var ret string
		return ret
	}
	return *o.Nationality
}

// GetNationalityOk returns a tuple with the Nationality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetNationalityOk() (*string, bool) {
	if o == nil || IsNil(o.Nationality) {
		return nil, false
	}
	return o.Nationality, true
}

// HasNationality returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasNationality() bool {
	if o != nil && !IsNil(o.Nationality) {
		return true
	}

	return false
}

// SetNationality gets a reference to the given string and assigns it to the Nationality field.
func (o *DocumentPropertiesExtractedData) SetNationality(v string) {
	o.Nationality = &v
}

// GetAddressLine1 returns the AddressLine1 field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetAddressLine1() string {
	if o == nil || IsNil(o.AddressLine1) {
		var ret string
		return ret
	}
	return *o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetAddressLine1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine1) {
		return nil, false
	}
	return o.AddressLine1, true
}

// HasAddressLine1 returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasAddressLine1() bool {
	if o != nil && !IsNil(o.AddressLine1) {
		return true
	}

	return false
}

// SetAddressLine1 gets a reference to the given string and assigns it to the AddressLine1 field.
func (o *DocumentPropertiesExtractedData) SetAddressLine1(v string) {
	o.AddressLine1 = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *DocumentPropertiesExtractedData) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetAddressLine3 returns the AddressLine3 field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetAddressLine3() string {
	if o == nil || IsNil(o.AddressLine3) {
		var ret string
		return ret
	}
	return *o.AddressLine3
}

// GetAddressLine3Ok returns a tuple with the AddressLine3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetAddressLine3Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine3) {
		return nil, false
	}
	return o.AddressLine3, true
}

// HasAddressLine3 returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasAddressLine3() bool {
	if o != nil && !IsNil(o.AddressLine3) {
		return true
	}

	return false
}

// SetAddressLine3 gets a reference to the given string and assigns it to the AddressLine3 field.
func (o *DocumentPropertiesExtractedData) SetAddressLine3(v string) {
	o.AddressLine3 = &v
}

// GetAddressLine4 returns the AddressLine4 field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetAddressLine4() string {
	if o == nil || IsNil(o.AddressLine4) {
		var ret string
		return ret
	}
	return *o.AddressLine4
}

// GetAddressLine4Ok returns a tuple with the AddressLine4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetAddressLine4Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine4) {
		return nil, false
	}
	return o.AddressLine4, true
}

// HasAddressLine4 returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasAddressLine4() bool {
	if o != nil && !IsNil(o.AddressLine4) {
		return true
	}

	return false
}

// SetAddressLine4 gets a reference to the given string and assigns it to the AddressLine4 field.
func (o *DocumentPropertiesExtractedData) SetAddressLine4(v string) {
	o.AddressLine4 = &v
}

// GetAddressLine5 returns the AddressLine5 field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetAddressLine5() string {
	if o == nil || IsNil(o.AddressLine5) {
		var ret string
		return ret
	}
	return *o.AddressLine5
}

// GetAddressLine5Ok returns a tuple with the AddressLine5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetAddressLine5Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine5) {
		return nil, false
	}
	return o.AddressLine5, true
}

// HasAddressLine5 returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasAddressLine5() bool {
	if o != nil && !IsNil(o.AddressLine5) {
		return true
	}

	return false
}

// SetAddressLine5 gets a reference to the given string and assigns it to the AddressLine5 field.
func (o *DocumentPropertiesExtractedData) SetAddressLine5(v string) {
	o.AddressLine5 = &v
}

// GetIssuingAuthority returns the IssuingAuthority field value if set, zero value otherwise.
func (o *DocumentPropertiesExtractedData) GetIssuingAuthority() string {
	if o == nil || IsNil(o.IssuingAuthority) {
		var ret string
		return ret
	}
	return *o.IssuingAuthority
}

// GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPropertiesExtractedData) GetIssuingAuthorityOk() (*string, bool) {
	if o == nil || IsNil(o.IssuingAuthority) {
		return nil, false
	}
	return o.IssuingAuthority, true
}

// HasIssuingAuthority returns a boolean if a field has been set.
func (o *DocumentPropertiesExtractedData) HasIssuingAuthority() bool {
	if o != nil && !IsNil(o.IssuingAuthority) {
		return true
	}

	return false
}

// SetIssuingAuthority gets a reference to the given string and assigns it to the IssuingAuthority field.
func (o *DocumentPropertiesExtractedData) SetIssuingAuthority(v string) {
	o.IssuingAuthority = &v
}

func (o DocumentPropertiesExtractedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentPropertiesExtractedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentNumber) {
		toSerialize["document_number"] = o.DocumentNumber
	}
	if !IsNil(o.DateOfBirth) {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}
	if !IsNil(o.DateOfExpiry) {
		toSerialize["date_of_expiry"] = o.DateOfExpiry
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if !IsNil(o.SpouseName) {
		toSerialize["spouse_name"] = o.SpouseName
	}
	if !IsNil(o.WidowName) {
		toSerialize["widow_name"] = o.WidowName
	}
	if !IsNil(o.AliasName) {
		toSerialize["alias_name"] = o.AliasName
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.MrzLine1) {
		toSerialize["mrz_line1"] = o.MrzLine1
	}
	if !IsNil(o.MrzLine2) {
		toSerialize["mrz_line2"] = o.MrzLine2
	}
	if !IsNil(o.MrzLine3) {
		toSerialize["mrz_line3"] = o.MrzLine3
	}
	if !IsNil(o.Nationality) {
		toSerialize["nationality"] = o.Nationality
	}
	if !IsNil(o.AddressLine1) {
		toSerialize["address_line_1"] = o.AddressLine1
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["address_line_2"] = o.AddressLine2
	}
	if !IsNil(o.AddressLine3) {
		toSerialize["address_line_3"] = o.AddressLine3
	}
	if !IsNil(o.AddressLine4) {
		toSerialize["address_line_4"] = o.AddressLine4
	}
	if !IsNil(o.AddressLine5) {
		toSerialize["address_line_5"] = o.AddressLine5
	}
	if !IsNil(o.IssuingAuthority) {
		toSerialize["issuing_authority"] = o.IssuingAuthority
	}
	return toSerialize, nil
}

type NullableDocumentPropertiesExtractedData struct {
	value *DocumentPropertiesExtractedData
	isSet bool
}

func (v NullableDocumentPropertiesExtractedData) Get() *DocumentPropertiesExtractedData {
	return v.value
}

func (v *NullableDocumentPropertiesExtractedData) Set(val *DocumentPropertiesExtractedData) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentPropertiesExtractedData) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentPropertiesExtractedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentPropertiesExtractedData(val *DocumentPropertiesExtractedData) *NullableDocumentPropertiesExtractedData {
	return &NullableDocumentPropertiesExtractedData{value: val, isSet: true}
}

func (v NullableDocumentPropertiesExtractedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentPropertiesExtractedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


