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

// checks if the WatchlistStandardBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchlistStandardBreakdown{}

// WatchlistStandardBreakdown struct for WatchlistStandardBreakdown
type WatchlistStandardBreakdown struct {
	Sanction *WatchlistAmlBreakdownSanction `json:"sanction,omitempty"`
	PoliticallyExposedPerson *WatchlistAmlBreakdownPoliticallyExposedPerson `json:"politically_exposed_person,omitempty"`
	LegalAndRegulatoryWarnings *WatchlistAmlBreakdownLegalAndRegulatoryWarnings `json:"legal_and_regulatory_warnings,omitempty"`
}

// NewWatchlistStandardBreakdown instantiates a new WatchlistStandardBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistStandardBreakdown() *WatchlistStandardBreakdown {
	this := WatchlistStandardBreakdown{}
	return &this
}

// NewWatchlistStandardBreakdownWithDefaults instantiates a new WatchlistStandardBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistStandardBreakdownWithDefaults() *WatchlistStandardBreakdown {
	this := WatchlistStandardBreakdown{}
	return &this
}

// GetSanction returns the Sanction field value if set, zero value otherwise.
func (o *WatchlistStandardBreakdown) GetSanction() WatchlistAmlBreakdownSanction {
	if o == nil || IsNil(o.Sanction) {
		var ret WatchlistAmlBreakdownSanction
		return ret
	}
	return *o.Sanction
}

// GetSanctionOk returns a tuple with the Sanction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistStandardBreakdown) GetSanctionOk() (*WatchlistAmlBreakdownSanction, bool) {
	if o == nil || IsNil(o.Sanction) {
		return nil, false
	}
	return o.Sanction, true
}

// HasSanction returns a boolean if a field has been set.
func (o *WatchlistStandardBreakdown) HasSanction() bool {
	if o != nil && !IsNil(o.Sanction) {
		return true
	}

	return false
}

// SetSanction gets a reference to the given WatchlistAmlBreakdownSanction and assigns it to the Sanction field.
func (o *WatchlistStandardBreakdown) SetSanction(v WatchlistAmlBreakdownSanction) {
	o.Sanction = &v
}

// GetPoliticallyExposedPerson returns the PoliticallyExposedPerson field value if set, zero value otherwise.
func (o *WatchlistStandardBreakdown) GetPoliticallyExposedPerson() WatchlistAmlBreakdownPoliticallyExposedPerson {
	if o == nil || IsNil(o.PoliticallyExposedPerson) {
		var ret WatchlistAmlBreakdownPoliticallyExposedPerson
		return ret
	}
	return *o.PoliticallyExposedPerson
}

// GetPoliticallyExposedPersonOk returns a tuple with the PoliticallyExposedPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistStandardBreakdown) GetPoliticallyExposedPersonOk() (*WatchlistAmlBreakdownPoliticallyExposedPerson, bool) {
	if o == nil || IsNil(o.PoliticallyExposedPerson) {
		return nil, false
	}
	return o.PoliticallyExposedPerson, true
}

// HasPoliticallyExposedPerson returns a boolean if a field has been set.
func (o *WatchlistStandardBreakdown) HasPoliticallyExposedPerson() bool {
	if o != nil && !IsNil(o.PoliticallyExposedPerson) {
		return true
	}

	return false
}

// SetPoliticallyExposedPerson gets a reference to the given WatchlistAmlBreakdownPoliticallyExposedPerson and assigns it to the PoliticallyExposedPerson field.
func (o *WatchlistStandardBreakdown) SetPoliticallyExposedPerson(v WatchlistAmlBreakdownPoliticallyExposedPerson) {
	o.PoliticallyExposedPerson = &v
}

// GetLegalAndRegulatoryWarnings returns the LegalAndRegulatoryWarnings field value if set, zero value otherwise.
func (o *WatchlistStandardBreakdown) GetLegalAndRegulatoryWarnings() WatchlistAmlBreakdownLegalAndRegulatoryWarnings {
	if o == nil || IsNil(o.LegalAndRegulatoryWarnings) {
		var ret WatchlistAmlBreakdownLegalAndRegulatoryWarnings
		return ret
	}
	return *o.LegalAndRegulatoryWarnings
}

// GetLegalAndRegulatoryWarningsOk returns a tuple with the LegalAndRegulatoryWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistStandardBreakdown) GetLegalAndRegulatoryWarningsOk() (*WatchlistAmlBreakdownLegalAndRegulatoryWarnings, bool) {
	if o == nil || IsNil(o.LegalAndRegulatoryWarnings) {
		return nil, false
	}
	return o.LegalAndRegulatoryWarnings, true
}

// HasLegalAndRegulatoryWarnings returns a boolean if a field has been set.
func (o *WatchlistStandardBreakdown) HasLegalAndRegulatoryWarnings() bool {
	if o != nil && !IsNil(o.LegalAndRegulatoryWarnings) {
		return true
	}

	return false
}

// SetLegalAndRegulatoryWarnings gets a reference to the given WatchlistAmlBreakdownLegalAndRegulatoryWarnings and assigns it to the LegalAndRegulatoryWarnings field.
func (o *WatchlistStandardBreakdown) SetLegalAndRegulatoryWarnings(v WatchlistAmlBreakdownLegalAndRegulatoryWarnings) {
	o.LegalAndRegulatoryWarnings = &v
}

func (o WatchlistStandardBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchlistStandardBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sanction) {
		toSerialize["sanction"] = o.Sanction
	}
	if !IsNil(o.PoliticallyExposedPerson) {
		toSerialize["politically_exposed_person"] = o.PoliticallyExposedPerson
	}
	if !IsNil(o.LegalAndRegulatoryWarnings) {
		toSerialize["legal_and_regulatory_warnings"] = o.LegalAndRegulatoryWarnings
	}
	return toSerialize, nil
}

type NullableWatchlistStandardBreakdown struct {
	value *WatchlistStandardBreakdown
	isSet bool
}

func (v NullableWatchlistStandardBreakdown) Get() *WatchlistStandardBreakdown {
	return v.value
}

func (v *NullableWatchlistStandardBreakdown) Set(val *WatchlistStandardBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistStandardBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistStandardBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistStandardBreakdown(val *WatchlistStandardBreakdown) *NullableWatchlistStandardBreakdown {
	return &NullableWatchlistStandardBreakdown{value: val, isSet: true}
}

func (v NullableWatchlistStandardBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistStandardBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


