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

// checks if the DeviceIntelligenceBreakdownPropertiesGeolocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceIntelligenceBreakdownPropertiesGeolocation{}

// DeviceIntelligenceBreakdownPropertiesGeolocation struct for DeviceIntelligenceBreakdownPropertiesGeolocation
type DeviceIntelligenceBreakdownPropertiesGeolocation struct {
	// City location of the IP address.
	City *string `json:"city,omitempty"`
	// Region location of the IP address.
	Region *string `json:"region,omitempty"`
	// Country location of the IP address in a three letter format.
	Country *CountryCodes `json:"country,omitempty"`
}

// NewDeviceIntelligenceBreakdownPropertiesGeolocation instantiates a new DeviceIntelligenceBreakdownPropertiesGeolocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIntelligenceBreakdownPropertiesGeolocation() *DeviceIntelligenceBreakdownPropertiesGeolocation {
	this := DeviceIntelligenceBreakdownPropertiesGeolocation{}
	return &this
}

// NewDeviceIntelligenceBreakdownPropertiesGeolocationWithDefaults instantiates a new DeviceIntelligenceBreakdownPropertiesGeolocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIntelligenceBreakdownPropertiesGeolocationWithDefaults() *DeviceIntelligenceBreakdownPropertiesGeolocation {
	this := DeviceIntelligenceBreakdownPropertiesGeolocation{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) SetCity(v string) {
	o.City = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) SetRegion(v string) {
	o.Region = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetCountry() CountryCodes {
	if o == nil || IsNil(o.Country) {
		var ret CountryCodes
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetCountryOk() (*CountryCodes, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given CountryCodes and assigns it to the Country field.
func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) SetCountry(v CountryCodes) {
	o.Country = &v
}

func (o DeviceIntelligenceBreakdownPropertiesGeolocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceIntelligenceBreakdownPropertiesGeolocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	return toSerialize, nil
}

type NullableDeviceIntelligenceBreakdownPropertiesGeolocation struct {
	value *DeviceIntelligenceBreakdownPropertiesGeolocation
	isSet bool
}

func (v NullableDeviceIntelligenceBreakdownPropertiesGeolocation) Get() *DeviceIntelligenceBreakdownPropertiesGeolocation {
	return v.value
}

func (v *NullableDeviceIntelligenceBreakdownPropertiesGeolocation) Set(val *DeviceIntelligenceBreakdownPropertiesGeolocation) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIntelligenceBreakdownPropertiesGeolocation) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIntelligenceBreakdownPropertiesGeolocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIntelligenceBreakdownPropertiesGeolocation(val *DeviceIntelligenceBreakdownPropertiesGeolocation) *NullableDeviceIntelligenceBreakdownPropertiesGeolocation {
	return &NullableDeviceIntelligenceBreakdownPropertiesGeolocation{value: val, isSet: true}
}

func (v NullableDeviceIntelligenceBreakdownPropertiesGeolocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIntelligenceBreakdownPropertiesGeolocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


