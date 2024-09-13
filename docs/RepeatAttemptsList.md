# RepeatAttemptsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **string** | The unique identifier of the completed Document report. | [optional] 
**RepeatAttempts** | [**[]RepeatAttemptsListRepeatAttemptsInner**](RepeatAttemptsListRepeatAttemptsInner.md) | An array of repeat attempt objects. If no repeat attempts were found, the array will be empty. The number of objects returned can increase over time if more matches are received.  | 
**AttemptsCount** | Pointer to **int32** | The total number of attempts using the same document, including the current report under assessment. | [optional] 
**AttemptsClearRate** | Pointer to **float32** | A number between 0 and 1 which indicates the proportion of attempts that have been cleared, including the current report under assessment. | [optional] 
**UniqueMismatchesCount** | Pointer to **int32** | The number of unique entries in the repeat_attempts field for which at least one of the fields is a mismatch. | [optional] 

## Methods

### NewRepeatAttemptsList

`func NewRepeatAttemptsList(repeatAttempts []RepeatAttemptsListRepeatAttemptsInner, ) *RepeatAttemptsList`

NewRepeatAttemptsList instantiates a new RepeatAttemptsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepeatAttemptsListWithDefaults

`func NewRepeatAttemptsListWithDefaults() *RepeatAttemptsList`

NewRepeatAttemptsListWithDefaults instantiates a new RepeatAttemptsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *RepeatAttemptsList) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *RepeatAttemptsList) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *RepeatAttemptsList) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *RepeatAttemptsList) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetRepeatAttempts

`func (o *RepeatAttemptsList) GetRepeatAttempts() []RepeatAttemptsListRepeatAttemptsInner`

GetRepeatAttempts returns the RepeatAttempts field if non-nil, zero value otherwise.

### GetRepeatAttemptsOk

`func (o *RepeatAttemptsList) GetRepeatAttemptsOk() (*[]RepeatAttemptsListRepeatAttemptsInner, bool)`

GetRepeatAttemptsOk returns a tuple with the RepeatAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatAttempts

`func (o *RepeatAttemptsList) SetRepeatAttempts(v []RepeatAttemptsListRepeatAttemptsInner)`

SetRepeatAttempts sets RepeatAttempts field to given value.


### GetAttemptsCount

`func (o *RepeatAttemptsList) GetAttemptsCount() int32`

GetAttemptsCount returns the AttemptsCount field if non-nil, zero value otherwise.

### GetAttemptsCountOk

`func (o *RepeatAttemptsList) GetAttemptsCountOk() (*int32, bool)`

GetAttemptsCountOk returns a tuple with the AttemptsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsCount

`func (o *RepeatAttemptsList) SetAttemptsCount(v int32)`

SetAttemptsCount sets AttemptsCount field to given value.

### HasAttemptsCount

`func (o *RepeatAttemptsList) HasAttemptsCount() bool`

HasAttemptsCount returns a boolean if a field has been set.

### GetAttemptsClearRate

`func (o *RepeatAttemptsList) GetAttemptsClearRate() float32`

GetAttemptsClearRate returns the AttemptsClearRate field if non-nil, zero value otherwise.

### GetAttemptsClearRateOk

`func (o *RepeatAttemptsList) GetAttemptsClearRateOk() (*float32, bool)`

GetAttemptsClearRateOk returns a tuple with the AttemptsClearRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsClearRate

`func (o *RepeatAttemptsList) SetAttemptsClearRate(v float32)`

SetAttemptsClearRate sets AttemptsClearRate field to given value.

### HasAttemptsClearRate

`func (o *RepeatAttemptsList) HasAttemptsClearRate() bool`

HasAttemptsClearRate returns a boolean if a field has been set.

### GetUniqueMismatchesCount

`func (o *RepeatAttemptsList) GetUniqueMismatchesCount() int32`

GetUniqueMismatchesCount returns the UniqueMismatchesCount field if non-nil, zero value otherwise.

### GetUniqueMismatchesCountOk

`func (o *RepeatAttemptsList) GetUniqueMismatchesCountOk() (*int32, bool)`

GetUniqueMismatchesCountOk returns a tuple with the UniqueMismatchesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueMismatchesCount

`func (o *RepeatAttemptsList) SetUniqueMismatchesCount(v int32)`

SetUniqueMismatchesCount sets UniqueMismatchesCount field to given value.

### HasUniqueMismatchesCount

`func (o *RepeatAttemptsList) HasUniqueMismatchesCount() bool`

HasUniqueMismatchesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


