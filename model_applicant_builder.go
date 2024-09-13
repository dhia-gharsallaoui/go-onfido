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

// checks if the ApplicantBuilder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicantBuilder{}

// ApplicantBuilder struct for ApplicantBuilder
type ApplicantBuilder struct {
	// The applicant's email address. Required if doing a US check, or a UK check for which `applicant_provides_data` is `true`.
	Email *string `json:"email,omitempty"`
	// The applicant's date of birth
	Dob *string `json:"dob,omitempty"`
	IdNumbers []IdNumber `json:"id_numbers,omitempty"`
	// The applicant's phone number
	PhoneNumber *string `json:"phone_number,omitempty"`
	Consents *ConsentsBuilder `json:"consents,omitempty"`
	Address *AddressBuilder `json:"address,omitempty"`
	Location *LocationBuilder `json:"location,omitempty"`
	// The applicant's first name
	FirstName string `json:"first_name" validate:"regexp=^[^!#$%*=<>;{}\\"]*$"`
	// The applicant's surname
	LastName string `json:"last_name" validate:"regexp=^[^!#$%*=<>;{}\\"]*$"`
}

type _ApplicantBuilder ApplicantBuilder

// NewApplicantBuilder instantiates a new ApplicantBuilder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicantBuilder(firstName string, lastName string) *ApplicantBuilder {
	this := ApplicantBuilder{}
	this.FirstName = firstName
	this.LastName = lastName
	return &this
}

// NewApplicantBuilderWithDefaults instantiates a new ApplicantBuilder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicantBuilderWithDefaults() *ApplicantBuilder {
	this := ApplicantBuilder{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ApplicantBuilder) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ApplicantBuilder) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ApplicantBuilder) SetEmail(v string) {
	o.Email = &v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *ApplicantBuilder) GetDob() string {
	if o == nil || IsNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetDobOk() (*string, bool) {
	if o == nil || IsNil(o.Dob) {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *ApplicantBuilder) HasDob() bool {
	if o != nil && !IsNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *ApplicantBuilder) SetDob(v string) {
	o.Dob = &v
}

// GetIdNumbers returns the IdNumbers field value if set, zero value otherwise.
func (o *ApplicantBuilder) GetIdNumbers() []IdNumber {
	if o == nil || IsNil(o.IdNumbers) {
		var ret []IdNumber
		return ret
	}
	return o.IdNumbers
}

// GetIdNumbersOk returns a tuple with the IdNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetIdNumbersOk() ([]IdNumber, bool) {
	if o == nil || IsNil(o.IdNumbers) {
		return nil, false
	}
	return o.IdNumbers, true
}

// HasIdNumbers returns a boolean if a field has been set.
func (o *ApplicantBuilder) HasIdNumbers() bool {
	if o != nil && !IsNil(o.IdNumbers) {
		return true
	}

	return false
}

// SetIdNumbers gets a reference to the given []IdNumber and assigns it to the IdNumbers field.
func (o *ApplicantBuilder) SetIdNumbers(v []IdNumber) {
	o.IdNumbers = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ApplicantBuilder) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ApplicantBuilder) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ApplicantBuilder) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetConsents returns the Consents field value if set, zero value otherwise.
func (o *ApplicantBuilder) GetConsents() ConsentsBuilder {
	if o == nil || IsNil(o.Consents) {
		var ret ConsentsBuilder
		return ret
	}
	return *o.Consents
}

// GetConsentsOk returns a tuple with the Consents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetConsentsOk() (*ConsentsBuilder, bool) {
	if o == nil || IsNil(o.Consents) {
		return nil, false
	}
	return o.Consents, true
}

// HasConsents returns a boolean if a field has been set.
func (o *ApplicantBuilder) HasConsents() bool {
	if o != nil && !IsNil(o.Consents) {
		return true
	}

	return false
}

// SetConsents gets a reference to the given ConsentsBuilder and assigns it to the Consents field.
func (o *ApplicantBuilder) SetConsents(v ConsentsBuilder) {
	o.Consents = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ApplicantBuilder) GetAddress() AddressBuilder {
	if o == nil || IsNil(o.Address) {
		var ret AddressBuilder
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetAddressOk() (*AddressBuilder, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ApplicantBuilder) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AddressBuilder and assigns it to the Address field.
func (o *ApplicantBuilder) SetAddress(v AddressBuilder) {
	o.Address = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ApplicantBuilder) GetLocation() LocationBuilder {
	if o == nil || IsNil(o.Location) {
		var ret LocationBuilder
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetLocationOk() (*LocationBuilder, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ApplicantBuilder) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationBuilder and assigns it to the Location field.
func (o *ApplicantBuilder) SetLocation(v LocationBuilder) {
	o.Location = &v
}

// GetFirstName returns the FirstName field value
func (o *ApplicantBuilder) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ApplicantBuilder) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *ApplicantBuilder) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ApplicantBuilder) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ApplicantBuilder) SetLastName(v string) {
	o.LastName = v
}

func (o ApplicantBuilder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicantBuilder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Dob) {
		toSerialize["dob"] = o.Dob
	}
	if !IsNil(o.IdNumbers) {
		toSerialize["id_numbers"] = o.IdNumbers
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Consents) {
		toSerialize["consents"] = o.Consents
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	return toSerialize, nil
}

func (o *ApplicantBuilder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"first_name",
		"last_name",
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

	varApplicantBuilder := _ApplicantBuilder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicantBuilder)

	if err != nil {
		return err
	}

	*o = ApplicantBuilder(varApplicantBuilder)

	return err
}

type NullableApplicantBuilder struct {
	value *ApplicantBuilder
	isSet bool
}

func (v NullableApplicantBuilder) Get() *ApplicantBuilder {
	return v.value
}

func (v *NullableApplicantBuilder) Set(val *ApplicantBuilder) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicantBuilder) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicantBuilder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicantBuilder(val *ApplicantBuilder) *NullableApplicantBuilder {
	return &NullableApplicantBuilder{value: val, isSet: true}
}

func (v NullableApplicantBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicantBuilder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


