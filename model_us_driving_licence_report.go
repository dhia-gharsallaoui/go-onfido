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

// checks if the UsDrivingLicenceReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsDrivingLicenceReport{}

// UsDrivingLicenceReport struct for UsDrivingLicenceReport
type UsDrivingLicenceReport struct {
	// The unique identifier for the report. Read-only.
	Id string `json:"id"`
	// The date and time at which the report was first initiated. Read-only.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The API endpoint to retrieve the report. Read-only.
	Href *string `json:"href,omitempty"`
	Status *ReportStatus `json:"status,omitempty"`
	Result *ReportResult `json:"result,omitempty"`
	SubResult *ReportSubResult `json:"sub_result,omitempty"`
	// The ID of the check to which the report belongs. Read-only.
	CheckId *string `json:"check_id,omitempty"`
	// Array of objects with document ids that were used in the Onfido engine. [ONLY POPULATED FOR DOCUMENT AND FACIAL SIMILARITY REPORTS]
	Documents []ReportDocument `json:"documents,omitempty"`
	Name ReportName `json:"name"`
	Breakdown *UsDrivingLicenceBreakdown `json:"breakdown,omitempty"`
	Properties *DocumentProperties `json:"properties,omitempty"`
}

type _UsDrivingLicenceReport UsDrivingLicenceReport

// NewUsDrivingLicenceReport instantiates a new UsDrivingLicenceReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsDrivingLicenceReport(id string, name ReportName) *UsDrivingLicenceReport {
	this := UsDrivingLicenceReport{}
	this.Id = id
	this.Name = name
	return &this
}

// NewUsDrivingLicenceReportWithDefaults instantiates a new UsDrivingLicenceReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsDrivingLicenceReportWithDefaults() *UsDrivingLicenceReport {
	this := UsDrivingLicenceReport{}
	return &this
}

// GetId returns the Id field value
func (o *UsDrivingLicenceReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UsDrivingLicenceReport) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UsDrivingLicenceReport) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *UsDrivingLicenceReport) SetHref(v string) {
	o.Href = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetStatus() ReportStatus {
	if o == nil || IsNil(o.Status) {
		var ret ReportStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetStatusOk() (*ReportStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ReportStatus and assigns it to the Status field.
func (o *UsDrivingLicenceReport) SetStatus(v ReportStatus) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetResult() ReportResult {
	if o == nil || IsNil(o.Result) {
		var ret ReportResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetResultOk() (*ReportResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given ReportResult and assigns it to the Result field.
func (o *UsDrivingLicenceReport) SetResult(v ReportResult) {
	o.Result = &v
}

// GetSubResult returns the SubResult field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetSubResult() ReportSubResult {
	if o == nil || IsNil(o.SubResult) {
		var ret ReportSubResult
		return ret
	}
	return *o.SubResult
}

// GetSubResultOk returns a tuple with the SubResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetSubResultOk() (*ReportSubResult, bool) {
	if o == nil || IsNil(o.SubResult) {
		return nil, false
	}
	return o.SubResult, true
}

// HasSubResult returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasSubResult() bool {
	if o != nil && !IsNil(o.SubResult) {
		return true
	}

	return false
}

// SetSubResult gets a reference to the given ReportSubResult and assigns it to the SubResult field.
func (o *UsDrivingLicenceReport) SetSubResult(v ReportSubResult) {
	o.SubResult = &v
}

// GetCheckId returns the CheckId field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetCheckId() string {
	if o == nil || IsNil(o.CheckId) {
		var ret string
		return ret
	}
	return *o.CheckId
}

// GetCheckIdOk returns a tuple with the CheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetCheckIdOk() (*string, bool) {
	if o == nil || IsNil(o.CheckId) {
		return nil, false
	}
	return o.CheckId, true
}

// HasCheckId returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasCheckId() bool {
	if o != nil && !IsNil(o.CheckId) {
		return true
	}

	return false
}

// SetCheckId gets a reference to the given string and assigns it to the CheckId field.
func (o *UsDrivingLicenceReport) SetCheckId(v string) {
	o.CheckId = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetDocuments() []ReportDocument {
	if o == nil || IsNil(o.Documents) {
		var ret []ReportDocument
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetDocumentsOk() ([]ReportDocument, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []ReportDocument and assigns it to the Documents field.
func (o *UsDrivingLicenceReport) SetDocuments(v []ReportDocument) {
	o.Documents = v
}

// GetName returns the Name field value
func (o *UsDrivingLicenceReport) GetName() ReportName {
	if o == nil {
		var ret ReportName
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetNameOk() (*ReportName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UsDrivingLicenceReport) SetName(v ReportName) {
	o.Name = v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetBreakdown() UsDrivingLicenceBreakdown {
	if o == nil || IsNil(o.Breakdown) {
		var ret UsDrivingLicenceBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetBreakdownOk() (*UsDrivingLicenceBreakdown, bool) {
	if o == nil || IsNil(o.Breakdown) {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasBreakdown() bool {
	if o != nil && !IsNil(o.Breakdown) {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given UsDrivingLicenceBreakdown and assigns it to the Breakdown field.
func (o *UsDrivingLicenceReport) SetBreakdown(v UsDrivingLicenceBreakdown) {
	o.Breakdown = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UsDrivingLicenceReport) GetProperties() DocumentProperties {
	if o == nil || IsNil(o.Properties) {
		var ret DocumentProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsDrivingLicenceReport) GetPropertiesOk() (*DocumentProperties, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UsDrivingLicenceReport) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given DocumentProperties and assigns it to the Properties field.
func (o *UsDrivingLicenceReport) SetProperties(v DocumentProperties) {
	o.Properties = &v
}

func (o UsDrivingLicenceReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsDrivingLicenceReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.SubResult) {
		toSerialize["sub_result"] = o.SubResult
	}
	if !IsNil(o.CheckId) {
		toSerialize["check_id"] = o.CheckId
	}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Breakdown) {
		toSerialize["breakdown"] = o.Breakdown
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

func (o *UsDrivingLicenceReport) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varUsDrivingLicenceReport := _UsDrivingLicenceReport{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsDrivingLicenceReport)

	if err != nil {
		return err
	}

	*o = UsDrivingLicenceReport(varUsDrivingLicenceReport)

	return err
}

type NullableUsDrivingLicenceReport struct {
	value *UsDrivingLicenceReport
	isSet bool
}

func (v NullableUsDrivingLicenceReport) Get() *UsDrivingLicenceReport {
	return v.value
}

func (v *NullableUsDrivingLicenceReport) Set(val *UsDrivingLicenceReport) {
	v.value = val
	v.isSet = true
}

func (v NullableUsDrivingLicenceReport) IsSet() bool {
	return v.isSet
}

func (v *NullableUsDrivingLicenceReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsDrivingLicenceReport(val *UsDrivingLicenceReport) *NullableUsDrivingLicenceReport {
	return &NullableUsDrivingLicenceReport{value: val, isSet: true}
}

func (v NullableUsDrivingLicenceReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsDrivingLicenceReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


