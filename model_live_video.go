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
)

// checks if the LiveVideo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveVideo{}

// LiveVideo struct for LiveVideo
type LiveVideo struct {
	// The unique identifier for the video.
	Id *string `json:"id,omitempty"`
	// The date and time at which the video was uploaded.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The uri of this resource.
	Href *string `json:"href,omitempty"`
	// The uri that can be used to download the video.
	DownloadHref *string `json:"download_href,omitempty"`
	// The name of the uploaded file.
	FileName *string `json:"file_name,omitempty"`
	// The size of the file in bytes.
	FileSize *int32 `json:"file_size,omitempty"`
	// The file type of the uploaded file.
	FileType *string `json:"file_type,omitempty"`
	// Challenge the end user was asked to perform during the video recording.
	Challenge []map[string]interface{} `json:"challenge,omitempty"`
}

// NewLiveVideo instantiates a new LiveVideo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveVideo() *LiveVideo {
	this := LiveVideo{}
	return &this
}

// NewLiveVideoWithDefaults instantiates a new LiveVideo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveVideoWithDefaults() *LiveVideo {
	this := LiveVideo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LiveVideo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveVideo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveVideo) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LiveVideo) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideo) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LiveVideo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LiveVideo) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *LiveVideo) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideo) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *LiveVideo) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *LiveVideo) SetHref(v string) {
	o.Href = &v
}

// GetDownloadHref returns the DownloadHref field value if set, zero value otherwise.
func (o *LiveVideo) GetDownloadHref() string {
	if o == nil || IsNil(o.DownloadHref) {
		var ret string
		return ret
	}
	return *o.DownloadHref
}

// GetDownloadHrefOk returns a tuple with the DownloadHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideo) GetDownloadHrefOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadHref) {
		return nil, false
	}
	return o.DownloadHref, true
}

// HasDownloadHref returns a boolean if a field has been set.
func (o *LiveVideo) HasDownloadHref() bool {
	if o != nil && !IsNil(o.DownloadHref) {
		return true
	}

	return false
}

// SetDownloadHref gets a reference to the given string and assigns it to the DownloadHref field.
func (o *LiveVideo) SetDownloadHref(v string) {
	o.DownloadHref = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *LiveVideo) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideo) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *LiveVideo) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *LiveVideo) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *LiveVideo) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideo) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *LiveVideo) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *LiveVideo) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *LiveVideo) GetFileType() string {
	if o == nil || IsNil(o.FileType) {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideo) GetFileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileType) {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *LiveVideo) HasFileType() bool {
	if o != nil && !IsNil(o.FileType) {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *LiveVideo) SetFileType(v string) {
	o.FileType = &v
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *LiveVideo) GetChallenge() []map[string]interface{} {
	if o == nil || IsNil(o.Challenge) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideo) GetChallengeOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *LiveVideo) HasChallenge() bool {
	if o != nil && !IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given []map[string]interface{} and assigns it to the Challenge field.
func (o *LiveVideo) SetChallenge(v []map[string]interface{}) {
	o.Challenge = v
}

func (o LiveVideo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveVideo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.DownloadHref) {
		toSerialize["download_href"] = o.DownloadHref
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.FileSize) {
		toSerialize["file_size"] = o.FileSize
	}
	if !IsNil(o.FileType) {
		toSerialize["file_type"] = o.FileType
	}
	if !IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	return toSerialize, nil
}

type NullableLiveVideo struct {
	value *LiveVideo
	isSet bool
}

func (v NullableLiveVideo) Get() *LiveVideo {
	return v.value
}

func (v *NullableLiveVideo) Set(val *LiveVideo) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveVideo(val *LiveVideo) *NullableLiveVideo {
	return &NullableLiveVideo{value: val, isSet: true}
}

func (v NullableLiveVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


