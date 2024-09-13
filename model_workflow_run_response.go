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

// checks if the WorkflowRunResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowRunResponse{}

// WorkflowRunResponse struct for WorkflowRunResponse
type WorkflowRunResponse struct {
	// The unique identifier for the Workflow Run.
	Id string `json:"id"`
	// The identifier for the Workflow version.
	WorkflowVersionId *int32 `json:"workflow_version_id,omitempty"`
	// The URL for viewing the Workflow Run results on your Onfido Dashboard.
	DashboardUrl *string `json:"dashboard_url,omitempty"`
	// The status of the Workflow Run.
	Status *string `json:"status,omitempty"`
	// Output object contains all of the properties configured on the Workflow version.
	Output map[string]interface{} `json:"output,omitempty"`
	// The reasons the Workflow Run outcome was reached. Configurable when creating the Workflow version.
	Reasons []string `json:"reasons,omitempty"`
	Error *WorkflowRunResponseError `json:"error,omitempty"`
	// Client token to use when loading this workflow run in the Onfido SDK.
	SdkToken NullableString `json:"sdk_token,omitempty"`
}

type _WorkflowRunResponse WorkflowRunResponse

// NewWorkflowRunResponse instantiates a new WorkflowRunResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowRunResponse(id string) *WorkflowRunResponse {
	this := WorkflowRunResponse{}
	this.Id = id
	return &this
}

// NewWorkflowRunResponseWithDefaults instantiates a new WorkflowRunResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowRunResponseWithDefaults() *WorkflowRunResponse {
	this := WorkflowRunResponse{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowRunResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowRunResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowRunResponse) SetId(v string) {
	o.Id = v
}

// GetWorkflowVersionId returns the WorkflowVersionId field value if set, zero value otherwise.
func (o *WorkflowRunResponse) GetWorkflowVersionId() int32 {
	if o == nil || IsNil(o.WorkflowVersionId) {
		var ret int32
		return ret
	}
	return *o.WorkflowVersionId
}

// GetWorkflowVersionIdOk returns a tuple with the WorkflowVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunResponse) GetWorkflowVersionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.WorkflowVersionId) {
		return nil, false
	}
	return o.WorkflowVersionId, true
}

// HasWorkflowVersionId returns a boolean if a field has been set.
func (o *WorkflowRunResponse) HasWorkflowVersionId() bool {
	if o != nil && !IsNil(o.WorkflowVersionId) {
		return true
	}

	return false
}

// SetWorkflowVersionId gets a reference to the given int32 and assigns it to the WorkflowVersionId field.
func (o *WorkflowRunResponse) SetWorkflowVersionId(v int32) {
	o.WorkflowVersionId = &v
}

// GetDashboardUrl returns the DashboardUrl field value if set, zero value otherwise.
func (o *WorkflowRunResponse) GetDashboardUrl() string {
	if o == nil || IsNil(o.DashboardUrl) {
		var ret string
		return ret
	}
	return *o.DashboardUrl
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunResponse) GetDashboardUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DashboardUrl) {
		return nil, false
	}
	return o.DashboardUrl, true
}

// HasDashboardUrl returns a boolean if a field has been set.
func (o *WorkflowRunResponse) HasDashboardUrl() bool {
	if o != nil && !IsNil(o.DashboardUrl) {
		return true
	}

	return false
}

// SetDashboardUrl gets a reference to the given string and assigns it to the DashboardUrl field.
func (o *WorkflowRunResponse) SetDashboardUrl(v string) {
	o.DashboardUrl = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WorkflowRunResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WorkflowRunResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WorkflowRunResponse) SetStatus(v string) {
	o.Status = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *WorkflowRunResponse) GetOutput() map[string]interface{} {
	if o == nil || IsNil(o.Output) {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunResponse) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *WorkflowRunResponse) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *WorkflowRunResponse) SetOutput(v map[string]interface{}) {
	o.Output = v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *WorkflowRunResponse) GetReasons() []string {
	if o == nil || IsNil(o.Reasons) {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunResponse) GetReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *WorkflowRunResponse) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *WorkflowRunResponse) SetReasons(v []string) {
	o.Reasons = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WorkflowRunResponse) GetError() WorkflowRunResponseError {
	if o == nil || IsNil(o.Error) {
		var ret WorkflowRunResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowRunResponse) GetErrorOk() (*WorkflowRunResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WorkflowRunResponse) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given WorkflowRunResponseError and assigns it to the Error field.
func (o *WorkflowRunResponse) SetError(v WorkflowRunResponseError) {
	o.Error = &v
}

// GetSdkToken returns the SdkToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowRunResponse) GetSdkToken() string {
	if o == nil || IsNil(o.SdkToken.Get()) {
		var ret string
		return ret
	}
	return *o.SdkToken.Get()
}

// GetSdkTokenOk returns a tuple with the SdkToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowRunResponse) GetSdkTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdkToken.Get(), o.SdkToken.IsSet()
}

// HasSdkToken returns a boolean if a field has been set.
func (o *WorkflowRunResponse) HasSdkToken() bool {
	if o != nil && o.SdkToken.IsSet() {
		return true
	}

	return false
}

// SetSdkToken gets a reference to the given NullableString and assigns it to the SdkToken field.
func (o *WorkflowRunResponse) SetSdkToken(v string) {
	o.SdkToken.Set(&v)
}
// SetSdkTokenNil sets the value for SdkToken to be an explicit nil
func (o *WorkflowRunResponse) SetSdkTokenNil() {
	o.SdkToken.Set(nil)
}

// UnsetSdkToken ensures that no value is present for SdkToken, not even an explicit nil
func (o *WorkflowRunResponse) UnsetSdkToken() {
	o.SdkToken.Unset()
}

func (o WorkflowRunResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowRunResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.WorkflowVersionId) {
		toSerialize["workflow_version_id"] = o.WorkflowVersionId
	}
	if !IsNil(o.DashboardUrl) {
		toSerialize["dashboard_url"] = o.DashboardUrl
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if o.SdkToken.IsSet() {
		toSerialize["sdk_token"] = o.SdkToken.Get()
	}
	return toSerialize, nil
}

func (o *WorkflowRunResponse) UnmarshalJSON(data []byte) (err error) {
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

	varWorkflowRunResponse := _WorkflowRunResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowRunResponse)

	if err != nil {
		return err
	}

	*o = WorkflowRunResponse(varWorkflowRunResponse)

	return err
}

type NullableWorkflowRunResponse struct {
	value *WorkflowRunResponse
	isSet bool
}

func (v NullableWorkflowRunResponse) Get() *WorkflowRunResponse {
	return v.value
}

func (v *NullableWorkflowRunResponse) Set(val *WorkflowRunResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowRunResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowRunResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowRunResponse(val *WorkflowRunResponse) *NullableWorkflowRunResponse {
	return &NullableWorkflowRunResponse{value: val, isSet: true}
}

func (v NullableWorkflowRunResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowRunResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


