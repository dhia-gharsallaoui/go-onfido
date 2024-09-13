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

// checks if the WorkflowRunSharedLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRunSharedLink{}

// WorkflowRunSharedLink Object for the configuration of the Workflow Run link.
type WorkflowRunSharedLink struct {
	// Link to access the Workflow Run without the need to integrate with Onfido's SDKs.
	Url *string `json:"url,omitempty"`
	// When the interactive section of the Workflow Run has completed successfully, the user will be redirected to this URL instead of seeing the default Onfido 'thank you' page.
	CompletedRedirectUrl *string `json:"completed_redirect_url,omitempty"`
	// When the link has expired, the user will be immediately redirected to this URL instead of seeing the default Onfido error message.
	ExpiredRedirectUrl *string `json:"expired_redirect_url,omitempty"`
	// Date and time when the link will expire.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// The code for the language when the workflow run is acessed using the link.
	Language *string `json:"language,omitempty"`
}

// NewWorkflowRunSharedLink instantiates a new WorkflowRunSharedLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunSharedLink() *WorkflowRunSharedLink {
	this := WorkflowRunSharedLink{}
	return &this
}

// NewWorkflowRunSharedLinkWithDefaults instantiates a new WorkflowRunSharedLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunSharedLinkWithDefaults() *WorkflowRunSharedLink {
	this := WorkflowRunSharedLink{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WorkflowRunSharedLink) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunSharedLink) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WorkflowRunSharedLink) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WorkflowRunSharedLink) SetUrl(v string) {
	o.Url = &v
}

// GetCompletedRedirectUrl returns the CompletedRedirectUrl field value if set, zero value otherwise.
func (o *WorkflowRunSharedLink) GetCompletedRedirectUrl() string {
	if o == nil || IsNil(o.CompletedRedirectUrl) {
		var ret string
		return ret
	}
	return *o.CompletedRedirectUrl
}

// GetCompletedRedirectUrlOk returns a tuple with the CompletedRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunSharedLink) GetCompletedRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CompletedRedirectUrl) {
		return nil, false
	}
	return o.CompletedRedirectUrl, true
}

// HasCompletedRedirectUrl returns a boolean if a field has been set.
func (o *WorkflowRunSharedLink) HasCompletedRedirectUrl() bool {
	if o != nil && !IsNil(o.CompletedRedirectUrl) {
		return true
	}

	return false
}

// SetCompletedRedirectUrl gets a reference to the given string and assigns it to the CompletedRedirectUrl field.
func (o *WorkflowRunSharedLink) SetCompletedRedirectUrl(v string) {
	o.CompletedRedirectUrl = &v
}

// GetExpiredRedirectUrl returns the ExpiredRedirectUrl field value if set, zero value otherwise.
func (o *WorkflowRunSharedLink) GetExpiredRedirectUrl() string {
	if o == nil || IsNil(o.ExpiredRedirectUrl) {
		var ret string
		return ret
	}
	return *o.ExpiredRedirectUrl
}

// GetExpiredRedirectUrlOk returns a tuple with the ExpiredRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunSharedLink) GetExpiredRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiredRedirectUrl) {
		return nil, false
	}
	return o.ExpiredRedirectUrl, true
}

// HasExpiredRedirectUrl returns a boolean if a field has been set.
func (o *WorkflowRunSharedLink) HasExpiredRedirectUrl() bool {
	if o != nil && !IsNil(o.ExpiredRedirectUrl) {
		return true
	}

	return false
}

// SetExpiredRedirectUrl gets a reference to the given string and assigns it to the ExpiredRedirectUrl field.
func (o *WorkflowRunSharedLink) SetExpiredRedirectUrl(v string) {
	o.ExpiredRedirectUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *WorkflowRunSharedLink) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunSharedLink) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *WorkflowRunSharedLink) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *WorkflowRunSharedLink) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *WorkflowRunSharedLink) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunSharedLink) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *WorkflowRunSharedLink) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *WorkflowRunSharedLink) SetLanguage(v string) {
	o.Language = &v
}

func (o WorkflowRunSharedLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRunSharedLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.CompletedRedirectUrl) {
		toSerialize["completed_redirect_url"] = o.CompletedRedirectUrl
	}
	if !IsNil(o.ExpiredRedirectUrl) {
		toSerialize["expired_redirect_url"] = o.ExpiredRedirectUrl
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableWorkflowRunSharedLink struct {
	value *WorkflowRunSharedLink
	isSet bool
}

func (v NullableWorkflowRunSharedLink) Get() *WorkflowRunSharedLink {
	return v.value
}

func (v *NullableWorkflowRunSharedLink) Set(val *WorkflowRunSharedLink) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunSharedLink) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunSharedLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunSharedLink(val *WorkflowRunSharedLink) *NullableWorkflowRunSharedLink {
	return &NullableWorkflowRunSharedLink{value: val, isSet: true}
}

func (v NullableWorkflowRunSharedLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunSharedLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


