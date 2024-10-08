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

// checks if the LocationBuilder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationBuilder{}

// LocationBuilder struct for LocationBuilder
type LocationBuilder struct {
	// The applicant's ip address.
	IpAddress *string `json:"ip_address,omitempty"`
	// The applicant's country of residence in 3-letter ISO code.
	CountryOfResidence *CountryCodes `json:"country_of_residence,omitempty"`
}

// NewLocationBuilder instantiates a new LocationBuilder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationBuilder() *LocationBuilder {
	this := LocationBuilder{}
	return &this
}

// NewLocationBuilderWithDefaults instantiates a new LocationBuilder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationBuilderWithDefaults() *LocationBuilder {
	this := LocationBuilder{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *LocationBuilder) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationBuilder) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *LocationBuilder) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *LocationBuilder) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetCountryOfResidence returns the CountryOfResidence field value if set, zero value otherwise.
func (o *LocationBuilder) GetCountryOfResidence() CountryCodes {
	if o == nil || IsNil(o.CountryOfResidence) {
		var ret CountryCodes
		return ret
	}
	return *o.CountryOfResidence
}

// GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationBuilder) GetCountryOfResidenceOk() (*CountryCodes, bool) {
	if o == nil || IsNil(o.CountryOfResidence) {
		return nil, false
	}
	return o.CountryOfResidence, true
}

// HasCountryOfResidence returns a boolean if a field has been set.
func (o *LocationBuilder) HasCountryOfResidence() bool {
	if o != nil && !IsNil(o.CountryOfResidence) {
		return true
	}

	return false
}

// SetCountryOfResidence gets a reference to the given CountryCodes and assigns it to the CountryOfResidence field.
func (o *LocationBuilder) SetCountryOfResidence(v CountryCodes) {
	o.CountryOfResidence = &v
}

func (o LocationBuilder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationBuilder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.CountryOfResidence) {
		toSerialize["country_of_residence"] = o.CountryOfResidence
	}
	return toSerialize, nil
}

type NullableLocationBuilder struct {
	value *LocationBuilder
	isSet bool
}

func (v NullableLocationBuilder) Get() *LocationBuilder {
	return v.value
}

func (v *NullableLocationBuilder) Set(val *LocationBuilder) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationBuilder) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationBuilder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationBuilder(val *LocationBuilder) *NullableLocationBuilder {
	return &NullableLocationBuilder{value: val, isSet: true}
}

func (v NullableLocationBuilder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationBuilder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


