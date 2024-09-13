/*
Onfido API v3.6

The Onfido API (v3.6)

API version: v3.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the Applicant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Applicant{}

// Applicant struct for Applicant
type Applicant struct {
	// The applicant's email address. Required if doing a US check, or a UK check for which `applicant_provides_data` is `true`.
	Email *string `json:"email,omitempty"`
	// The applicant's date of birth
	Dob *string `json:"dob,omitempty"`
	IdNumbers []IdNumber `json:"id_numbers,omitempty"`
	// The applicant's phone number
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The applicant's first name
	FirstName *string `json:"first_name,omitempty" validate:"regexp=^[^!#$%*=<>;{}\\"]*$"`
	// The applicant's surname
	LastName *string `json:"last_name,omitempty" validate:"regexp=^[^!#$%*=<>;{}\\"]*$"`
	// The unique identifier for the applicant.
	Id string `json:"id"`
	// The date and time when this applicant was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time when this applicant is scheduled to be deleted.
	DeleteAt *time.Time `json:"delete_at,omitempty"`
	// The uri of this resource.
	Href *string `json:"href,omitempty"`
	Sandbox *bool `json:"sandbox,omitempty"`
	Address *Address `json:"address,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type _Applicant Applicant

// NewApplicant instantiates a new Applicant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicant(id string) *Applicant {
	this := Applicant{}
	this.Id = id
	return &this
}

// NewApplicantWithDefaults instantiates a new Applicant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicantWithDefaults() *Applicant {
	this := Applicant{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Applicant) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Applicant) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Applicant) SetEmail(v string) {
	o.Email = &v
}

// GetDob returns the Dob field value if set, zero value otherwise.
func (o *Applicant) GetDob() string {
	if o == nil || IsNil(o.Dob) {
		var ret string
		return ret
	}
	return *o.Dob
}

// GetDobOk returns a tuple with the Dob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetDobOk() (*string, bool) {
	if o == nil || IsNil(o.Dob) {
		return nil, false
	}
	return o.Dob, true
}

// HasDob returns a boolean if a field has been set.
func (o *Applicant) HasDob() bool {
	if o != nil && !IsNil(o.Dob) {
		return true
	}

	return false
}

// SetDob gets a reference to the given string and assigns it to the Dob field.
func (o *Applicant) SetDob(v string) {
	o.Dob = &v
}

// GetIdNumbers returns the IdNumbers field value if set, zero value otherwise.
func (o *Applicant) GetIdNumbers() []IdNumber {
	if o == nil || IsNil(o.IdNumbers) {
		var ret []IdNumber
		return ret
	}
	return o.IdNumbers
}

// GetIdNumbersOk returns a tuple with the IdNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetIdNumbersOk() ([]IdNumber, bool) {
	if o == nil || IsNil(o.IdNumbers) {
		return nil, false
	}
	return o.IdNumbers, true
}

// HasIdNumbers returns a boolean if a field has been set.
func (o *Applicant) HasIdNumbers() bool {
	if o != nil && !IsNil(o.IdNumbers) {
		return true
	}

	return false
}

// SetIdNumbers gets a reference to the given []IdNumber and assigns it to the IdNumbers field.
func (o *Applicant) SetIdNumbers(v []IdNumber) {
	o.IdNumbers = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *Applicant) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *Applicant) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *Applicant) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Applicant) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Applicant) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Applicant) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Applicant) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Applicant) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Applicant) SetLastName(v string) {
	o.LastName = &v
}

// GetId returns the Id field value
func (o *Applicant) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Applicant) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Applicant) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Applicant) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Applicant) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Applicant) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeleteAt returns the DeleteAt field value if set, zero value otherwise.
func (o *Applicant) GetDeleteAt() time.Time {
	if o == nil || IsNil(o.DeleteAt) {
		var ret time.Time
		return ret
	}
	return *o.DeleteAt
}

// GetDeleteAtOk returns a tuple with the DeleteAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetDeleteAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeleteAt) {
		return nil, false
	}
	return o.DeleteAt, true
}

// HasDeleteAt returns a boolean if a field has been set.
func (o *Applicant) HasDeleteAt() bool {
	if o != nil && !IsNil(o.DeleteAt) {
		return true
	}

	return false
}

// SetDeleteAt gets a reference to the given time.Time and assigns it to the DeleteAt field.
func (o *Applicant) SetDeleteAt(v time.Time) {
	o.DeleteAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Applicant) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Applicant) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Applicant) SetHref(v string) {
	o.Href = &v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *Applicant) GetSandbox() bool {
	if o == nil || IsNil(o.Sandbox) {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetSandboxOk() (*bool, bool) {
	if o == nil || IsNil(o.Sandbox) {
		return nil, false
	}
	return o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *Applicant) HasSandbox() bool {
	if o != nil && !IsNil(o.Sandbox) {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *Applicant) SetSandbox(v bool) {
	o.Sandbox = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Applicant) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Applicant) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Applicant) SetAddress(v Address) {
	o.Address = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Applicant) GetLocation() Location {
	if o == nil || IsNil(o.Location) {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Applicant) GetLocationOk() (*Location, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Applicant) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *Applicant) SetLocation(v Location) {
	o.Location = &v
}

func (o Applicant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Applicant) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DeleteAt) {
		toSerialize["delete_at"] = o.DeleteAt
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Sandbox) {
		toSerialize["sandbox"] = o.Sandbox
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	return toSerialize, nil
}

func (o *Applicant) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varApplicant := _Applicant{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicant)

	if err != nil {
		return err
	}

	*o = Applicant(varApplicant)

	return err
}

type NullableApplicant struct {
	value *Applicant
	isSet bool
}

func (v NullableApplicant) Get() *Applicant {
	return v.value
}

func (v *NullableApplicant) Set(val *Applicant) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicant) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicant(val *Applicant) *NullableApplicant {
	return &NullableApplicant{value: val, isSet: true}
}

func (v NullableApplicant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


