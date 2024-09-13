# CheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportNames** | [**[]ReportName**](ReportName.md) | An array of report names (strings). | 
**DocumentIds** | Pointer to **[]string** | Optional. An array of document ids, for use with Document reports only. If omitted, the Document report will use the most recently uploaded document by default. | [optional] 
**ApplicantProvidesData** | Pointer to **bool** | Send an applicant form to applicant to complete to proceed with check. Defaults to false. | [optional] [default to false]
**Asynchronous** | Pointer to **bool** | Defaults to &#x60;true&#x60;. If set to &#x60;false&#x60;, you will only receive a response when all reports in your check have completed.  | [optional] [default to true]
**SuppressFormEmails** | Pointer to **bool** | For checks where &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;, applicant form will not be automatically sent if &#x60;suppress_form_emails&#x60; is set to &#x60;true&#x60;. You can manually send the form at any time after the check has been created, using the link found in the form_uri attribute of the check object. Write-only. Defaults to false.  | [optional] 
**SubResult** | Pointer to **string** | Triggers responses for particular sub-results for sandbox Document reports. | [optional] 
**Consider** | Pointer to [**[]ReportName**](ReportName.md) | Array of names of particular reports to return consider as their results. This is a feature available in sandbox testing | [optional] 
**UsDrivingLicence** | Pointer to [**UsDrivingLicenceBuilder**](UsDrivingLicenceBuilder.md) |  | [optional] 

## Methods

### NewCheckRequest

`func NewCheckRequest(reportNames []ReportName, ) *CheckRequest`

NewCheckRequest instantiates a new CheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckRequestWithDefaults

`func NewCheckRequestWithDefaults() *CheckRequest`

NewCheckRequestWithDefaults instantiates a new CheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportNames

`func (o *CheckRequest) GetReportNames() []ReportName`

GetReportNames returns the ReportNames field if non-nil, zero value otherwise.

### GetReportNamesOk

`func (o *CheckRequest) GetReportNamesOk() (*[]ReportName, bool)`

GetReportNamesOk returns a tuple with the ReportNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportNames

`func (o *CheckRequest) SetReportNames(v []ReportName)`

SetReportNames sets ReportNames field to given value.


### GetDocumentIds

`func (o *CheckRequest) GetDocumentIds() []string`

GetDocumentIds returns the DocumentIds field if non-nil, zero value otherwise.

### GetDocumentIdsOk

`func (o *CheckRequest) GetDocumentIdsOk() (*[]string, bool)`

GetDocumentIdsOk returns a tuple with the DocumentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIds

`func (o *CheckRequest) SetDocumentIds(v []string)`

SetDocumentIds sets DocumentIds field to given value.

### HasDocumentIds

`func (o *CheckRequest) HasDocumentIds() bool`

HasDocumentIds returns a boolean if a field has been set.

### GetApplicantProvidesData

`func (o *CheckRequest) GetApplicantProvidesData() bool`

GetApplicantProvidesData returns the ApplicantProvidesData field if non-nil, zero value otherwise.

### GetApplicantProvidesDataOk

`func (o *CheckRequest) GetApplicantProvidesDataOk() (*bool, bool)`

GetApplicantProvidesDataOk returns a tuple with the ApplicantProvidesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantProvidesData

`func (o *CheckRequest) SetApplicantProvidesData(v bool)`

SetApplicantProvidesData sets ApplicantProvidesData field to given value.

### HasApplicantProvidesData

`func (o *CheckRequest) HasApplicantProvidesData() bool`

HasApplicantProvidesData returns a boolean if a field has been set.

### GetAsynchronous

`func (o *CheckRequest) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *CheckRequest) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *CheckRequest) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *CheckRequest) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetSuppressFormEmails

`func (o *CheckRequest) GetSuppressFormEmails() bool`

GetSuppressFormEmails returns the SuppressFormEmails field if non-nil, zero value otherwise.

### GetSuppressFormEmailsOk

`func (o *CheckRequest) GetSuppressFormEmailsOk() (*bool, bool)`

GetSuppressFormEmailsOk returns a tuple with the SuppressFormEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressFormEmails

`func (o *CheckRequest) SetSuppressFormEmails(v bool)`

SetSuppressFormEmails sets SuppressFormEmails field to given value.

### HasSuppressFormEmails

`func (o *CheckRequest) HasSuppressFormEmails() bool`

HasSuppressFormEmails returns a boolean if a field has been set.

### GetSubResult

`func (o *CheckRequest) GetSubResult() string`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *CheckRequest) GetSubResultOk() (*string, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *CheckRequest) SetSubResult(v string)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *CheckRequest) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetConsider

`func (o *CheckRequest) GetConsider() []ReportName`

GetConsider returns the Consider field if non-nil, zero value otherwise.

### GetConsiderOk

`func (o *CheckRequest) GetConsiderOk() (*[]ReportName, bool)`

GetConsiderOk returns a tuple with the Consider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsider

`func (o *CheckRequest) SetConsider(v []ReportName)`

SetConsider sets Consider field to given value.

### HasConsider

`func (o *CheckRequest) HasConsider() bool`

HasConsider returns a boolean if a field has been set.

### GetUsDrivingLicence

`func (o *CheckRequest) GetUsDrivingLicence() UsDrivingLicenceBuilder`

GetUsDrivingLicence returns the UsDrivingLicence field if non-nil, zero value otherwise.

### GetUsDrivingLicenceOk

`func (o *CheckRequest) GetUsDrivingLicenceOk() (*UsDrivingLicenceBuilder, bool)`

GetUsDrivingLicenceOk returns a tuple with the UsDrivingLicence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsDrivingLicence

`func (o *CheckRequest) SetUsDrivingLicence(v UsDrivingLicenceBuilder)`

SetUsDrivingLicence sets UsDrivingLicence field to given value.

### HasUsDrivingLicence

`func (o *CheckRequest) HasUsDrivingLicence() bool`

HasUsDrivingLicence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


