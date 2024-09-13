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

// checks if the DocumentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentResponse{}

// DocumentResponse struct for DocumentResponse
type DocumentResponse struct {
	// The unique identifier for the document
	Id string `json:"id"`
	// The date and time at which the document was uploaded
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The uri of this resource
	Href *string `json:"href,omitempty"`
	// The uri that can be used to download the document
	DownloadHref *string `json:"download_href,omitempty"`
	// The name of the uploaded file
	FileName *string `json:"file_name,omitempty"`
	// The size of the file in bytes
	FileSize *int32 `json:"file_size,omitempty"`
}

type _DocumentResponse DocumentResponse

// NewDocumentResponse instantiates a new DocumentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentResponse(id string) *DocumentResponse {
	this := DocumentResponse{}
	this.Id = id
	return &this
}

// NewDocumentResponseWithDefaults instantiates a new DocumentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentResponseWithDefaults() *DocumentResponse {
	this := DocumentResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DocumentResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DocumentResponse) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DocumentResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DocumentResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DocumentResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *DocumentResponse) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *DocumentResponse) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *DocumentResponse) SetHref(v string) {
	o.Href = &v
}

// GetDownloadHref returns the DownloadHref field value if set, zero value otherwise.
func (o *DocumentResponse) GetDownloadHref() string {
	if o == nil || IsNil(o.DownloadHref) {
		var ret string
		return ret
	}
	return *o.DownloadHref
}

// GetDownloadHrefOk returns a tuple with the DownloadHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetDownloadHrefOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadHref) {
		return nil, false
	}
	return o.DownloadHref, true
}

// HasDownloadHref returns a boolean if a field has been set.
func (o *DocumentResponse) HasDownloadHref() bool {
	if o != nil && !IsNil(o.DownloadHref) {
		return true
	}

	return false
}

// SetDownloadHref gets a reference to the given string and assigns it to the DownloadHref field.
func (o *DocumentResponse) SetDownloadHref(v string) {
	o.DownloadHref = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *DocumentResponse) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *DocumentResponse) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *DocumentResponse) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *DocumentResponse) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentResponse) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *DocumentResponse) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *DocumentResponse) SetFileSize(v int32) {
	o.FileSize = &v
}

func (o DocumentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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
	return toSerialize, nil
}

func (o *DocumentResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDocumentResponse := _DocumentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDocumentResponse)

	if err != nil {
		return err
	}

	*o = DocumentResponse(varDocumentResponse)

	return err
}

type NullableDocumentResponse struct {
	value *DocumentResponse
	isSet bool
}

func (v NullableDocumentResponse) Get() *DocumentResponse {
	return v.value
}

func (v *NullableDocumentResponse) Set(val *DocumentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentResponse(val *DocumentResponse) *NullableDocumentResponse {
	return &NullableDocumentResponse{value: val, isSet: true}
}

func (v NullableDocumentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


