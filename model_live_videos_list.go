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

// checks if the LiveVideosList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveVideosList{}

// LiveVideosList struct for LiveVideosList
type LiveVideosList struct {
	LiveVideos []LiveVideo `json:"live_videos"`
}

type _LiveVideosList LiveVideosList

// NewLiveVideosList instantiates a new LiveVideosList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveVideosList(liveVideos []LiveVideo) *LiveVideosList {
	this := LiveVideosList{}
	this.LiveVideos = liveVideos
	return &this
}

// NewLiveVideosListWithDefaults instantiates a new LiveVideosList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveVideosListWithDefaults() *LiveVideosList {
	this := LiveVideosList{}
	return &this
}

// GetLiveVideos returns the LiveVideos field value
func (o *LiveVideosList) GetLiveVideos() []LiveVideo {
	if o == nil {
		var ret []LiveVideo
		return ret
	}

	return o.LiveVideos
}

// GetLiveVideosOk returns a tuple with the LiveVideos field value
// and a boolean to check if the value has been set.
func (o *LiveVideosList) GetLiveVideosOk() ([]LiveVideo, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveVideos, true
}

// SetLiveVideos sets field value
func (o *LiveVideosList) SetLiveVideos(v []LiveVideo) {
	o.LiveVideos = v
}

func (o LiveVideosList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveVideosList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["live_videos"] = o.LiveVideos
	return toSerialize, nil
}

func (o *LiveVideosList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"live_videos",
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

	varLiveVideosList := _LiveVideosList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLiveVideosList)

	if err != nil {
		return err
	}

	*o = LiveVideosList(varLiveVideosList)

	return err
}

type NullableLiveVideosList struct {
	value *LiveVideosList
	isSet bool
}

func (v NullableLiveVideosList) Get() *LiveVideosList {
	return v.value
}

func (v *NullableLiveVideosList) Set(val *LiveVideosList) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveVideosList) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveVideosList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveVideosList(val *LiveVideosList) *NullableLiveVideosList {
	return &NullableLiveVideosList{value: val, isSet: true}
}

func (v NullableLiveVideosList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveVideosList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


