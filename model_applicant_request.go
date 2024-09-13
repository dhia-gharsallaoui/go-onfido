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

// checks if the ApplicantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicantRequest{}

// ApplicantRequest struct for ApplicantRequest
type ApplicantRequest struct {
	Consents *ConsentsBuilder `json:"consents,omitempty"`
	Address *AddressBuilder `json:"address,omitempty"`
	Location *LocationBuilder `json:"location,omitempty"`
}

// NewApplicantRequest instantiates a new ApplicantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicantRequest() *ApplicantRequest {
	this := ApplicantRequest{}
	return &this
}

// NewApplicantRequestWithDefaults instantiates a new ApplicantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicantRequestWithDefaults() *ApplicantRequest {
	this := ApplicantRequest{}
	return &this
}

// GetConsents returns the Consents field value if set, zero value otherwise.
func (o *ApplicantRequest) GetConsents() ConsentsBuilder {
	if o == nil || IsNil(o.Consents) {
		var ret ConsentsBuilder
		return ret
	}
	return *o.Consents
}

// GetConsentsOk returns a tuple with the Consents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantRequest) GetConsentsOk() (*ConsentsBuilder, bool) {
	if o == nil || IsNil(o.Consents) {
		return nil, false
	}
	return o.Consents, true
}

// HasConsents returns a boolean if a field has been set.
func (o *ApplicantRequest) HasConsents() bool {
	if o != nil && !IsNil(o.Consents) {
		return true
	}

	return false
}

// SetConsents gets a reference to the given ConsentsBuilder and assigns it to the Consents field.
func (o *ApplicantRequest) SetConsents(v ConsentsBuilder) {
	o.Consents = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ApplicantRequest) GetAddress() AddressBuilder {
	if o == nil || IsNil(o.Address) {
		var ret AddressBuilder
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantRequest) GetAddressOk() (*AddressBuilder, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ApplicantRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AddressBuilder and assigns it to the Address field.
func (o *ApplicantRequest) SetAddress(v AddressBuilder) {
	o.Address = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ApplicantRequest) GetLocation() LocationBuilder {
	if o == nil || IsNil(o.Location) {
		var ret LocationBuilder
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicantRequest) GetLocationOk() (*LocationBuilder, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ApplicantRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given LocationBuilder and assigns it to the Location field.
func (o *ApplicantRequest) SetLocation(v LocationBuilder) {
	o.Location = &v
}

func (o ApplicantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Consents) {
		toSerialize["consents"] = o.Consents
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	return toSerialize, nil
}

type NullableApplicantRequest struct {
	value *ApplicantRequest
	isSet bool
}

func (v NullableApplicantRequest) Get() *ApplicantRequest {
	return v.value
}

func (v *NullableApplicantRequest) Set(val *ApplicantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicantRequest(val *ApplicantRequest) *NullableApplicantRequest {
	return &NullableApplicantRequest{value: val, isSet: true}
}

func (v NullableApplicantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


