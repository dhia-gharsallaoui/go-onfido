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

// checks if the WorkflowRunShared type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRunShared{}

// WorkflowRunShared struct for WorkflowRunShared
type WorkflowRunShared struct {
	// The unique identifier for the Applicant.
	ApplicantId string `json:"applicant_id"`
	// The unique identifier for the Workflow.
	WorkflowId string `json:"workflow_id"`
	// Tags or labels assigned to the workflow run.
	Tags []string `json:"tags,omitempty"`
	// Customer-provided user identifier.
	CustomerUserId *string `json:"customer_user_id,omitempty"`
	Link *WorkflowRunSharedLink `json:"link,omitempty"`
	// The date and time when the Workflow Run was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time when the Workflow Run was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type _WorkflowRunShared WorkflowRunShared

// NewWorkflowRunShared instantiates a new WorkflowRunShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunShared(applicantId string, workflowId string) *WorkflowRunShared {
	this := WorkflowRunShared{}
	this.ApplicantId = applicantId
	this.WorkflowId = workflowId
	return &this
}

// NewWorkflowRunSharedWithDefaults instantiates a new WorkflowRunShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunSharedWithDefaults() *WorkflowRunShared {
	this := WorkflowRunShared{}
	return &this
}

// GetApplicantId returns the ApplicantId field value
func (o *WorkflowRunShared) GetApplicantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicantId
}

// GetApplicantIdOk returns a tuple with the ApplicantId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunShared) GetApplicantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicantId, true
}

// SetApplicantId sets field value
func (o *WorkflowRunShared) SetApplicantId(v string) {
	o.ApplicantId = v
}

// GetWorkflowId returns the WorkflowId field value
func (o *WorkflowRunShared) GetWorkflowId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowId
}

// GetWorkflowIdOk returns a tuple with the WorkflowId field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunShared) GetWorkflowIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowId, true
}

// SetWorkflowId sets field value
func (o *WorkflowRunShared) SetWorkflowId(v string) {
	o.WorkflowId = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRunShared) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRunShared) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WorkflowRunShared) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WorkflowRunShared) SetTags(v []string) {
	o.Tags = v
}

// GetCustomerUserId returns the CustomerUserId field value if set, zero value otherwise.
func (o *WorkflowRunShared) GetCustomerUserId() string {
	if o == nil || IsNil(o.CustomerUserId) {
		var ret string
		return ret
	}
	return *o.CustomerUserId
}

// GetCustomerUserIdOk returns a tuple with the CustomerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunShared) GetCustomerUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerUserId) {
		return nil, false
	}
	return o.CustomerUserId, true
}

// HasCustomerUserId returns a boolean if a field has been set.
func (o *WorkflowRunShared) HasCustomerUserId() bool {
	if o != nil && !IsNil(o.CustomerUserId) {
		return true
	}

	return false
}

// SetCustomerUserId gets a reference to the given string and assigns it to the CustomerUserId field.
func (o *WorkflowRunShared) SetCustomerUserId(v string) {
	o.CustomerUserId = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *WorkflowRunShared) GetLink() WorkflowRunSharedLink {
	if o == nil || IsNil(o.Link) {
		var ret WorkflowRunSharedLink
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunShared) GetLinkOk() (*WorkflowRunSharedLink, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *WorkflowRunShared) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given WorkflowRunSharedLink and assigns it to the Link field.
func (o *WorkflowRunShared) SetLink(v WorkflowRunSharedLink) {
	o.Link = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WorkflowRunShared) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunShared) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WorkflowRunShared) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WorkflowRunShared) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkflowRunShared) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunShared) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkflowRunShared) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WorkflowRunShared) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o WorkflowRunShared) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRunShared) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicant_id"] = o.ApplicantId
	toSerialize["workflow_id"] = o.WorkflowId
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomerUserId) {
		toSerialize["customer_user_id"] = o.CustomerUserId
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

func (o *WorkflowRunShared) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicant_id",
		"workflow_id",
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

	varWorkflowRunShared := _WorkflowRunShared{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowRunShared)

	if err != nil {
		return err
	}

	*o = WorkflowRunShared(varWorkflowRunShared)

	return err
}

type NullableWorkflowRunShared struct {
	value *WorkflowRunShared
	isSet bool
}

func (v NullableWorkflowRunShared) Get() *WorkflowRunShared {
	return v.value
}

func (v *NullableWorkflowRunShared) Set(val *WorkflowRunShared) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunShared) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunShared(val *WorkflowRunShared) *NullableWorkflowRunShared {
	return &NullableWorkflowRunShared{value: val, isSet: true}
}

func (v NullableWorkflowRunShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


