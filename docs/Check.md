# Check

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookIds** | Pointer to **[]string** | An array of webhook ids describing which webhooks to trigger for this check. | [optional] 
**ApplicantId** | **string** | The ID of the applicant to do the check on. | 
**ApplicantProvidesData** | Pointer to **bool** | Send an applicant form to applicant to complete to proceed with check. Defaults to false.  | [optional] 
**Tags** | Pointer to **[]string** | Array of tags being assigned to this check. | [optional] 
**RedirectUri** | Pointer to **string** | For checks where &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;, redirect to this URI when the applicant has submitted their data. | [optional] 
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

### NewCheck

`func NewCheck(applicantId string, id string, ) *Check`

NewCheck instantiates a new Check object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWithDefaults

`func NewCheckWithDefaults() *Check`

NewCheckWithDefaults instantiates a new Check object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookIds

`func (o *Check) GetWebhookIds() []string`

GetWebhookIds returns the WebhookIds field if non-nil, zero value otherwise.

### GetWebhookIdsOk

`func (o *Check) GetWebhookIdsOk() (*[]string, bool)`

GetWebhookIdsOk returns a tuple with the WebhookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIds

`func (o *Check) SetWebhookIds(v []string)`

SetWebhookIds sets WebhookIds field to given value.

### HasWebhookIds

`func (o *Check) HasWebhookIds() bool`

HasWebhookIds returns a boolean if a field has been set.

### GetApplicantId

`func (o *Check) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *Check) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *Check) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicantProvidesData

`func (o *Check) GetApplicantProvidesData() bool`

GetApplicantProvidesData returns the ApplicantProvidesData field if non-nil, zero value otherwise.

### GetApplicantProvidesDataOk

`func (o *Check) GetApplicantProvidesDataOk() (*bool, bool)`

GetApplicantProvidesDataOk returns a tuple with the ApplicantProvidesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantProvidesData

`func (o *Check) SetApplicantProvidesData(v bool)`

SetApplicantProvidesData sets ApplicantProvidesData field to given value.

### HasApplicantProvidesData

`func (o *Check) HasApplicantProvidesData() bool`

HasApplicantProvidesData returns a boolean if a field has been set.

### GetTags

`func (o *Check) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Check) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Check) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Check) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRedirectUri

`func (o *Check) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *Check) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *Check) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *Check) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetId

`func (o *Check) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Check) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Check) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Check) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Check) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Check) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Check) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *Check) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Check) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Check) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Check) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *Check) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Check) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Check) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Check) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *Check) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Check) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Check) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Check) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetFormUri

`func (o *Check) GetFormUri() string`

GetFormUri returns the FormUri field if non-nil, zero value otherwise.

### GetFormUriOk

`func (o *Check) GetFormUriOk() (*string, bool)`

GetFormUriOk returns a tuple with the FormUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormUri

`func (o *Check) SetFormUri(v string)`

SetFormUri sets FormUri field to given value.

### HasFormUri

`func (o *Check) HasFormUri() bool`

HasFormUri returns a boolean if a field has been set.

### GetResultsUri

`func (o *Check) GetResultsUri() string`

GetResultsUri returns the ResultsUri field if non-nil, zero value otherwise.

### GetResultsUriOk

`func (o *Check) GetResultsUriOk() (*string, bool)`

GetResultsUriOk returns a tuple with the ResultsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsUri

`func (o *Check) SetResultsUri(v string)`

SetResultsUri sets ResultsUri field to given value.

### HasResultsUri

`func (o *Check) HasResultsUri() bool`

HasResultsUri returns a boolean if a field has been set.

### GetReportIds

`func (o *Check) GetReportIds() []string`

GetReportIds returns the ReportIds field if non-nil, zero value otherwise.

### GetReportIdsOk

`func (o *Check) GetReportIdsOk() (*[]string, bool)`

GetReportIdsOk returns a tuple with the ReportIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIds

`func (o *Check) SetReportIds(v []string)`

SetReportIds sets ReportIds field to given value.

### HasReportIds

`func (o *Check) HasReportIds() bool`

HasReportIds returns a boolean if a field has been set.

### GetSandbox

`func (o *Check) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Check) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Check) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *Check) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


