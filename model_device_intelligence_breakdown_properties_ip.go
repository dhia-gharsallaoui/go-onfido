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

// checks if the DeviceIntelligenceBreakdownPropertiesIp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceIntelligenceBreakdownPropertiesIp{}

// DeviceIntelligenceBreakdownPropertiesIp struct for DeviceIntelligenceBreakdownPropertiesIp
type DeviceIntelligenceBreakdownPropertiesIp struct {
	// The IP address that uploaded the media.
	Address *string `json:"address,omitempty"`
	// The likelihood of the network connection being a VPN.
	// Deprecated
	VpnDetection *string `json:"vpn_detection,omitempty"`
	// The likelihood of the network connection being a Proxy.
	// Deprecated
	ProxyDetection *string `json:"proxy_detection,omitempty"`
	// The type of organization that owns this IP address.
	// Deprecated
	Type *string `json:"type,omitempty"`
}

// NewDeviceIntelligenceBreakdownPropertiesIp instantiates a new DeviceIntelligenceBreakdownPropertiesIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceIntelligenceBreakdownPropertiesIp() *DeviceIntelligenceBreakdownPropertiesIp {
	this := DeviceIntelligenceBreakdownPropertiesIp{}
	return &this
}

// NewDeviceIntelligenceBreakdownPropertiesIpWithDefaults instantiates a new DeviceIntelligenceBreakdownPropertiesIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceIntelligenceBreakdownPropertiesIpWithDefaults() *DeviceIntelligenceBreakdownPropertiesIp {
	this := DeviceIntelligenceBreakdownPropertiesIp{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *DeviceIntelligenceBreakdownPropertiesIp) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceIntelligenceBreakdownPropertiesIp) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *DeviceIntelligenceBreakdownPropertiesIp) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *DeviceIntelligenceBreakdownPropertiesIp) SetAddress(v string) {
	o.Address = &v
}

// GetVpnDetection returns the VpnDetection field value if set, zero value otherwise.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) GetVpnDetection() string {
	if o == nil || IsNil(o.VpnDetection) {
		var ret string
		return ret
	}
	return *o.VpnDetection
}

// GetVpnDetectionOk returns a tuple with the VpnDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) GetVpnDetectionOk() (*string, bool) {
	if o == nil || IsNil(o.VpnDetection) {
		return nil, false
	}
	return o.VpnDetection, true
}

// HasVpnDetection returns a boolean if a field has been set.
func (o *DeviceIntelligenceBreakdownPropertiesIp) HasVpnDetection() bool {
	if o != nil && !IsNil(o.VpnDetection) {
		return true
	}

	return false
}

// SetVpnDetection gets a reference to the given string and assigns it to the VpnDetection field.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) SetVpnDetection(v string) {
	o.VpnDetection = &v
}

// GetProxyDetection returns the ProxyDetection field value if set, zero value otherwise.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) GetProxyDetection() string {
	if o == nil || IsNil(o.ProxyDetection) {
		var ret string
		return ret
	}
	return *o.ProxyDetection
}

// GetProxyDetectionOk returns a tuple with the ProxyDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) GetProxyDetectionOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyDetection) {
		return nil, false
	}
	return o.ProxyDetection, true
}

// HasProxyDetection returns a boolean if a field has been set.
func (o *DeviceIntelligenceBreakdownPropertiesIp) HasProxyDetection() bool {
	if o != nil && !IsNil(o.ProxyDetection) {
		return true
	}

	return false
}

// SetProxyDetection gets a reference to the given string and assigns it to the ProxyDetection field.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) SetProxyDetection(v string) {
	o.ProxyDetection = &v
}

// GetType returns the Type field value if set, zero value otherwise.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceIntelligenceBreakdownPropertiesIp) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
// Deprecated
func (o *DeviceIntelligenceBreakdownPropertiesIp) SetType(v string) {
	o.Type = &v
}

func (o DeviceIntelligenceBreakdownPropertiesIp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceIntelligenceBreakdownPropertiesIp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.VpnDetection) {
		toSerialize["vpn_detection"] = o.VpnDetection
	}
	if !IsNil(o.ProxyDetection) {
		toSerialize["proxy_detection"] = o.ProxyDetection
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDeviceIntelligenceBreakdownPropertiesIp struct {
	value *DeviceIntelligenceBreakdownPropertiesIp
	isSet bool
}

func (v NullableDeviceIntelligenceBreakdownPropertiesIp) Get() *DeviceIntelligenceBreakdownPropertiesIp {
	return v.value
}

func (v *NullableDeviceIntelligenceBreakdownPropertiesIp) Set(val *DeviceIntelligenceBreakdownPropertiesIp) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceIntelligenceBreakdownPropertiesIp) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceIntelligenceBreakdownPropertiesIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceIntelligenceBreakdownPropertiesIp(val *DeviceIntelligenceBreakdownPropertiesIp) *NullableDeviceIntelligenceBreakdownPropertiesIp {
	return &NullableDeviceIntelligenceBreakdownPropertiesIp{value: val, isSet: true}
}

func (v NullableDeviceIntelligenceBreakdownPropertiesIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceIntelligenceBreakdownPropertiesIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


