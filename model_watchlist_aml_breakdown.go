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

// checks if the WatchlistAmlBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WatchlistAmlBreakdown{}

// WatchlistAmlBreakdown struct for WatchlistAmlBreakdown
type WatchlistAmlBreakdown struct {
	Sanction *WatchlistAmlBreakdownSanction `json:"sanction,omitempty"`
	PoliticallyExposedPerson *WatchlistAmlBreakdownPoliticallyExposedPerson `json:"politically_exposed_person,omitempty"`
	LegalAndRegulatoryWarnings *WatchlistAmlBreakdownLegalAndRegulatoryWarnings `json:"legal_and_regulatory_warnings,omitempty"`
	AdverseMedia *WatchlistAmlBreakdownAdverseMedia `json:"adverse_media,omitempty"`
}

// NewWatchlistAmlBreakdown instantiates a new WatchlistAmlBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWatchlistAmlBreakdown() *WatchlistAmlBreakdown {
	this := WatchlistAmlBreakdown{}
	return &this
}

// NewWatchlistAmlBreakdownWithDefaults instantiates a new WatchlistAmlBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWatchlistAmlBreakdownWithDefaults() *WatchlistAmlBreakdown {
	this := WatchlistAmlBreakdown{}
	return &this
}

// GetSanction returns the Sanction field value if set, zero value otherwise.
func (o *WatchlistAmlBreakdown) GetSanction() WatchlistAmlBreakdownSanction {
	if o == nil || IsNil(o.Sanction) {
		var ret WatchlistAmlBreakdownSanction
		return ret
	}
	return *o.Sanction
}

// GetSanctionOk returns a tuple with the Sanction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistAmlBreakdown) GetSanctionOk() (*WatchlistAmlBreakdownSanction, bool) {
	if o == nil || IsNil(o.Sanction) {
		return nil, false
	}
	return o.Sanction, true
}

// HasSanction returns a boolean if a field has been set.
func (o *WatchlistAmlBreakdown) HasSanction() bool {
	if o != nil && !IsNil(o.Sanction) {
		return true
	}

	return false
}

// SetSanction gets a reference to the given WatchlistAmlBreakdownSanction and assigns it to the Sanction field.
func (o *WatchlistAmlBreakdown) SetSanction(v WatchlistAmlBreakdownSanction) {
	o.Sanction = &v
}

// GetPoliticallyExposedPerson returns the PoliticallyExposedPerson field value if set, zero value otherwise.
func (o *WatchlistAmlBreakdown) GetPoliticallyExposedPerson() WatchlistAmlBreakdownPoliticallyExposedPerson {
	if o == nil || IsNil(o.PoliticallyExposedPerson) {
		var ret WatchlistAmlBreakdownPoliticallyExposedPerson
		return ret
	}
	return *o.PoliticallyExposedPerson
}

// GetPoliticallyExposedPersonOk returns a tuple with the PoliticallyExposedPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistAmlBreakdown) GetPoliticallyExposedPersonOk() (*WatchlistAmlBreakdownPoliticallyExposedPerson, bool) {
	if o == nil || IsNil(o.PoliticallyExposedPerson) {
		return nil, false
	}
	return o.PoliticallyExposedPerson, true
}

// HasPoliticallyExposedPerson returns a boolean if a field has been set.
func (o *WatchlistAmlBreakdown) HasPoliticallyExposedPerson() bool {
	if o != nil && !IsNil(o.PoliticallyExposedPerson) {
		return true
	}

	return false
}

// SetPoliticallyExposedPerson gets a reference to the given WatchlistAmlBreakdownPoliticallyExposedPerson and assigns it to the PoliticallyExposedPerson field.
func (o *WatchlistAmlBreakdown) SetPoliticallyExposedPerson(v WatchlistAmlBreakdownPoliticallyExposedPerson) {
	o.PoliticallyExposedPerson = &v
}

// GetLegalAndRegulatoryWarnings returns the LegalAndRegulatoryWarnings field value if set, zero value otherwise.
func (o *WatchlistAmlBreakdown) GetLegalAndRegulatoryWarnings() WatchlistAmlBreakdownLegalAndRegulatoryWarnings {
	if o == nil || IsNil(o.LegalAndRegulatoryWarnings) {
		var ret WatchlistAmlBreakdownLegalAndRegulatoryWarnings
		return ret
	}
	return *o.LegalAndRegulatoryWarnings
}

// GetLegalAndRegulatoryWarningsOk returns a tuple with the LegalAndRegulatoryWarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistAmlBreakdown) GetLegalAndRegulatoryWarningsOk() (*WatchlistAmlBreakdownLegalAndRegulatoryWarnings, bool) {
	if o == nil || IsNil(o.LegalAndRegulatoryWarnings) {
		return nil, false
	}
	return o.LegalAndRegulatoryWarnings, true
}

// HasLegalAndRegulatoryWarnings returns a boolean if a field has been set.
func (o *WatchlistAmlBreakdown) HasLegalAndRegulatoryWarnings() bool {
	if o != nil && !IsNil(o.LegalAndRegulatoryWarnings) {
		return true
	}

	return false
}

// SetLegalAndRegulatoryWarnings gets a reference to the given WatchlistAmlBreakdownLegalAndRegulatoryWarnings and assigns it to the LegalAndRegulatoryWarnings field.
func (o *WatchlistAmlBreakdown) SetLegalAndRegulatoryWarnings(v WatchlistAmlBreakdownLegalAndRegulatoryWarnings) {
	o.LegalAndRegulatoryWarnings = &v
}

// GetAdverseMedia returns the AdverseMedia field value if set, zero value otherwise.
func (o *WatchlistAmlBreakdown) GetAdverseMedia() WatchlistAmlBreakdownAdverseMedia {
	if o == nil || IsNil(o.AdverseMedia) {
		var ret WatchlistAmlBreakdownAdverseMedia
		return ret
	}
	return *o.AdverseMedia
}

// GetAdverseMediaOk returns a tuple with the AdverseMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WatchlistAmlBreakdown) GetAdverseMediaOk() (*WatchlistAmlBreakdownAdverseMedia, bool) {
	if o == nil || IsNil(o.AdverseMedia) {
		return nil, false
	}
	return o.AdverseMedia, true
}

// HasAdverseMedia returns a boolean if a field has been set.
func (o *WatchlistAmlBreakdown) HasAdverseMedia() bool {
	if o != nil && !IsNil(o.AdverseMedia) {
		return true
	}

	return false
}

// SetAdverseMedia gets a reference to the given WatchlistAmlBreakdownAdverseMedia and assigns it to the AdverseMedia field.
func (o *WatchlistAmlBreakdown) SetAdverseMedia(v WatchlistAmlBreakdownAdverseMedia) {
	o.AdverseMedia = &v
}

func (o WatchlistAmlBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WatchlistAmlBreakdown) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdverseMedia) {
		toSerialize["adverse_media"] = o.AdverseMedia
	}
	return toSerialize, nil
}

type NullableWatchlistAmlBreakdown struct {
	value *WatchlistAmlBreakdown
	isSet bool
}

func (v NullableWatchlistAmlBreakdown) Get() *WatchlistAmlBreakdown {
	return v.value
}

func (v *NullableWatchlistAmlBreakdown) Set(val *WatchlistAmlBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableWatchlistAmlBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableWatchlistAmlBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWatchlistAmlBreakdown(val *WatchlistAmlBreakdown) *NullableWatchlistAmlBreakdown {
	return &NullableWatchlistAmlBreakdown{value: val, isSet: true}
}

func (v NullableWatchlistAmlBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWatchlistAmlBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


