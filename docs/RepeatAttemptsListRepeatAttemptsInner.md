# RepeatAttemptsListRepeatAttemptsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **string** | The unique identifier of the matching Document report. | [optional] 
**ApplicantId** | Pointer to **string** | The unique identifier of the applicant for the matching Document report. | [optional] 
**DateOfBirth** | Pointer to **string** | Whether the dates of birth are exactly the same or are different. | [optional] 
**Names** | Pointer to **string** | Whether the names are exactly the same or are different. | [optional] 
**Result** | Pointer to **string** | The report result of this attempt. | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the matching report was created. | [optional] 
**CompletedAt** | Pointer to **time.Time** | When the matching report was completed. | [optional] 

## Methods

### NewRepeatAttemptsListRepeatAttemptsInner

`func NewRepeatAttemptsListRepeatAttemptsInner() *RepeatAttemptsListRepeatAttemptsInner`

NewRepeatAttemptsListRepeatAttemptsInner instantiates a new RepeatAttemptsListRepeatAttemptsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepeatAttemptsListRepeatAttemptsInnerWithDefaults

`func NewRepeatAttemptsListRepeatAttemptsInnerWithDefaults() *RepeatAttemptsListRepeatAttemptsInner`

NewRepeatAttemptsListRepeatAttemptsInnerWithDefaults instantiates a new RepeatAttemptsListRepeatAttemptsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *RepeatAttemptsListRepeatAttemptsInner) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *RepeatAttemptsListRepeatAttemptsInner) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetApplicantId

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *RepeatAttemptsListRepeatAttemptsInner) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.

### HasApplicantId

`func (o *RepeatAttemptsListRepeatAttemptsInner) HasApplicantId() bool`

HasApplicantId returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *RepeatAttemptsListRepeatAttemptsInner) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *RepeatAttemptsListRepeatAttemptsInner) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetNames

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetNames() string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetNamesOk() (*string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *RepeatAttemptsListRepeatAttemptsInner) SetNames(v string)`

SetNames sets Names field to given value.

### HasNames

`func (o *RepeatAttemptsListRepeatAttemptsInner) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetResult

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *RepeatAttemptsListRepeatAttemptsInner) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *RepeatAttemptsListRepeatAttemptsInner) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RepeatAttemptsListRepeatAttemptsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RepeatAttemptsListRepeatAttemptsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *RepeatAttemptsListRepeatAttemptsInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *RepeatAttemptsListRepeatAttemptsInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *RepeatAttemptsListRepeatAttemptsInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


