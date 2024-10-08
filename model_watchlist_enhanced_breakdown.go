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

// checks if the WatchlistEnhancedBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchlistEnhancedBreakdown{}

// WatchlistEnhancedBreakdown struct for WatchlistEnhancedBreakdown
type WatchlistEnhancedBreakdown struct {
	PoliticallyExposedPerson *WatchlistAmlBreakdownPoliticallyExposedPerson `json:"politically_exposed_person,omitempty"`
	Sanction *WatchlistAmlBreakdownSanction `json:"sanction,omitempty"`
	AdverseMedia *WatchlistAmlBreakdownAdverseMedia `json:"adverse_media,omitempty"`
	MonitoredLists *WatchlistAmlBreakdownLegalAndRegulatoryWarnings `json:"monitored_lists,omitempty"`
}

// NewWatchlistEnhancedBreakdown instantiates a new WatchlistEnhancedBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistEnhancedBreakdown() *WatchlistEnhancedBreakdown {
	this := WatchlistEnhancedBreakdown{}
	return &this
}

// NewWatchlistEnhancedBreakdownWithDefaults instantiates a new WatchlistEnhancedBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistEnhancedBreakdownWithDefaults() *WatchlistEnhancedBreakdown {
	this := WatchlistEnhancedBreakdown{}
	return &this
}

// GetPoliticallyExposedPerson returns the PoliticallyExposedPerson field value if set, zero value otherwise.
func (o *WatchlistEnhancedBreakdown) GetPoliticallyExposedPerson() WatchlistAmlBreakdownPoliticallyExposedPerson {
	if o == nil || IsNil(o.PoliticallyExposedPerson) {
		var ret WatchlistAmlBreakdownPoliticallyExposedPerson
		return ret
	}
	return *o.PoliticallyExposedPerson
}

// GetPoliticallyExposedPersonOk returns a tuple with the PoliticallyExposedPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistEnhancedBreakdown) GetPoliticallyExposedPersonOk() (*WatchlistAmlBreakdownPoliticallyExposedPerson, bool) {
	if o == nil || IsNil(o.PoliticallyExposedPerson) {
		return nil, false
	}
	return o.PoliticallyExposedPerson, true
}

// HasPoliticallyExposedPerson returns a boolean if a field has been set.
func (o *WatchlistEnhancedBreakdown) HasPoliticallyExposedPerson() bool {
	if o != nil && !IsNil(o.PoliticallyExposedPerson) {
		return true
	}

	return false
}

// SetPoliticallyExposedPerson gets a reference to the given WatchlistAmlBreakdownPoliticallyExposedPerson and assigns it to the PoliticallyExposedPerson field.
func (o *WatchlistEnhancedBreakdown) SetPoliticallyExposedPerson(v WatchlistAmlBreakdownPoliticallyExposedPerson) {
	o.PoliticallyExposedPerson = &v
}

// GetSanction returns the Sanction field value if set, zero value otherwise.
func (o *WatchlistEnhancedBreakdown) GetSanction() WatchlistAmlBreakdownSanction {
	if o == nil || IsNil(o.Sanction) {
		var ret WatchlistAmlBreakdownSanction
		return ret
	}
	return *o.Sanction
}

// GetSanctionOk returns a tuple with the Sanction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistEnhancedBreakdown) GetSanctionOk() (*WatchlistAmlBreakdownSanction, bool) {
	if o == nil || IsNil(o.Sanction) {
		return nil, false
	}
	return o.Sanction, true
}

// HasSanction returns a boolean if a field has been set.
func (o *WatchlistEnhancedBreakdown) HasSanction() bool {
	if o != nil && !IsNil(o.Sanction) {
		return true
	}

	return false
}

// SetSanction gets a reference to the given WatchlistAmlBreakdownSanction and assigns it to the Sanction field.
func (o *WatchlistEnhancedBreakdown) SetSanction(v WatchlistAmlBreakdownSanction) {
	o.Sanction = &v
}

// GetAdverseMedia returns the AdverseMedia field value if set, zero value otherwise.
func (o *WatchlistEnhancedBreakdown) GetAdverseMedia() WatchlistAmlBreakdownAdverseMedia {
	if o == nil || IsNil(o.AdverseMedia) {
		var ret WatchlistAmlBreakdownAdverseMedia
		return ret
	}
	return *o.AdverseMedia
}

// GetAdverseMediaOk returns a tuple with the AdverseMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistEnhancedBreakdown) GetAdverseMediaOk() (*WatchlistAmlBreakdownAdverseMedia, bool) {
	if o == nil || IsNil(o.AdverseMedia) {
		return nil, false
	}
	return o.AdverseMedia, true
}

// HasAdverseMedia returns a boolean if a field has been set.
func (o *WatchlistEnhancedBreakdown) HasAdverseMedia() bool {
	if o != nil && !IsNil(o.AdverseMedia) {
		return true
	}

	return false
}

// SetAdverseMedia gets a reference to the given WatchlistAmlBreakdownAdverseMedia and assigns it to the AdverseMedia field.
func (o *WatchlistEnhancedBreakdown) SetAdverseMedia(v WatchlistAmlBreakdownAdverseMedia) {
	o.AdverseMedia = &v
}

// GetMonitoredLists returns the MonitoredLists field value if set, zero value otherwise.
func (o *WatchlistEnhancedBreakdown) GetMonitoredLists() WatchlistAmlBreakdownLegalAndRegulatoryWarnings {
	if o == nil || IsNil(o.MonitoredLists) {
		var ret WatchlistAmlBreakdownLegalAndRegulatoryWarnings
		return ret
	}
	return *o.MonitoredLists
}

// GetMonitoredListsOk returns a tuple with the MonitoredLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistEnhancedBreakdown) GetMonitoredListsOk() (*WatchlistAmlBreakdownLegalAndRegulatoryWarnings, bool) {
	if o == nil || IsNil(o.MonitoredLists) {
		return nil, false
	}
	return o.MonitoredLists, true
}

// HasMonitoredLists returns a boolean if a field has been set.
func (o *WatchlistEnhancedBreakdown) HasMonitoredLists() bool {
	if o != nil && !IsNil(o.MonitoredLists) {
		return true
	}

	return false
}

// SetMonitoredLists gets a reference to the given WatchlistAmlBreakdownLegalAndRegulatoryWarnings and assigns it to the MonitoredLists field.
func (o *WatchlistEnhancedBreakdown) SetMonitoredLists(v WatchlistAmlBreakdownLegalAndRegulatoryWarnings) {
	o.MonitoredLists = &v
}

func (o WatchlistEnhancedBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchlistEnhancedBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PoliticallyExposedPerson) {
		toSerialize["politically_exposed_person"] = o.PoliticallyExposedPerson
	}
	if !IsNil(o.Sanction) {
		toSerialize["sanction"] = o.Sanction
	}
	if !IsNil(o.AdverseMedia) {
		toSerialize["adverse_media"] = o.AdverseMedia
	}
	if !IsNil(o.MonitoredLists) {
		toSerialize["monitored_lists"] = o.MonitoredLists
	}
	return toSerialize, nil
}

type NullableWatchlistEnhancedBreakdown struct {
	value *WatchlistEnhancedBreakdown
	isSet bool
}

func (v NullableWatchlistEnhancedBreakdown) Get() *WatchlistEnhancedBreakdown {
	return v.value
}

func (v *NullableWatchlistEnhancedBreakdown) Set(val *WatchlistEnhancedBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistEnhancedBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistEnhancedBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistEnhancedBreakdown(val *WatchlistEnhancedBreakdown) *NullableWatchlistEnhancedBreakdown {
	return &NullableWatchlistEnhancedBreakdown{value: val, isSet: true}
}

func (v NullableWatchlistEnhancedBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistEnhancedBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


