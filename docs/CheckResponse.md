# CheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the check. | 
**CreatedAt** | Pointer to **time.Time** | The date and time when this check was created. | [optional] 
**Href** | Pointer to **string** | The uri of this resource. | [optional] 
**Status** | Pointer to **string** | The current state of the check in the checking process. | [optional] 
**Result** | Pointer to **string** | The overall result of the check, based on the results of the constituent reports. | [optional] 
**FormUri** | Pointer to **string** | A link to the applicant form, if &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;. | [optional] 
**ResultsUri** | Pointer to **string** | A link to the corresponding results page on the Onfido dashboard. | [optional] 
**ReportIds** | Pointer to **[]string** | An array of report ids. | [optional] 
**Sandbox** | Pointer to **bool** | Indicates whether the object was created in the sandbox or not. | [optional] 

## Methods

### NewCheckResponse

`func NewCheckResponse(id string, ) *CheckResponse`

NewCheckResponse instantiates a new CheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckResponseWithDefaults

`func NewCheckResponseWithDefaults() *CheckResponse`

NewCheckResponseWithDefaults instantiates a new CheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CheckResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CheckResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *CheckResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CheckResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CheckResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CheckResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *CheckResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *CheckResponse) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CheckResponse) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CheckResponse) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CheckResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetFormUri

`func (o *CheckResponse) GetFormUri() string`

GetFormUri returns the FormUri field if non-nil, zero value otherwise.

### GetFormUriOk

`func (o *CheckResponse) GetFormUriOk() (*string, bool)`

GetFormUriOk returns a tuple with the FormUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormUri

`func (o *CheckResponse) SetFormUri(v string)`

SetFormUri sets FormUri field to given value.

### HasFormUri

`func (o *CheckResponse) HasFormUri() bool`

HasFormUri returns a boolean if a field has been set.

### GetResultsUri

`func (o *CheckResponse) GetResultsUri() string`

GetResultsUri returns the ResultsUri field if non-nil, zero value otherwise.

### GetResultsUriOk

`func (o *CheckResponse) GetResultsUriOk() (*string, bool)`

GetResultsUriOk returns a tuple with the ResultsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsUri

`func (o *CheckResponse) SetResultsUri(v string)`

SetResultsUri sets ResultsUri field to given value.

### HasResultsUri

`func (o *CheckResponse) HasResultsUri() bool`

HasResultsUri returns a boolean if a field has been set.

### GetReportIds

`func (o *CheckResponse) GetReportIds() []string`

GetReportIds returns the ReportIds field if non-nil, zero value otherwise.

### GetReportIdsOk

`func (o *CheckResponse) GetReportIdsOk() (*[]string, bool)`

GetReportIdsOk returns a tuple with the ReportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIds

`func (o *CheckResponse) SetReportIds(v []string)`

SetReportIds sets ReportIds field to given value.

### HasReportIds

`func (o *CheckResponse) HasReportIds() bool`

HasReportIds returns a boolean if a field has been set.

### GetSandbox

`func (o *CheckResponse) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *CheckResponse) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *CheckResponse) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *CheckResponse) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


