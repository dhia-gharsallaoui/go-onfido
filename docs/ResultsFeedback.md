# ResultsFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedResult** | Pointer to **string** | The expected result for the check or report. | [optional] 
**CheckId** | Pointer to **string** | The ID of the check (only if report_id is not provided). | [optional] 
**ReportId** | Pointer to **string** | The ID of the check (only if check_id is not provided). | [optional] 
**FeedbackNotes** | Pointer to **string** | Any additional information or feedback. | [optional] 

## Methods

### NewResultsFeedback

`func NewResultsFeedback() *ResultsFeedback`

NewResultsFeedback instantiates a new ResultsFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsFeedbackWithDefaults

`func NewResultsFeedbackWithDefaults() *ResultsFeedback`

NewResultsFeedbackWithDefaults instantiates a new ResultsFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedResult

`func (o *ResultsFeedback) GetExpectedResult() string`

GetExpectedResult returns the ExpectedResult field if non-nil, zero value otherwise.

### GetExpectedResultOk

`func (o *ResultsFeedback) GetExpectedResultOk() (*string, bool)`

GetExpectedResultOk returns a tuple with the ExpectedResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedResult

`func (o *ResultsFeedback) SetExpectedResult(v string)`

SetExpectedResult sets ExpectedResult field to given value.

### HasExpectedResult

`func (o *ResultsFeedback) HasExpectedResult() bool`

HasExpectedResult returns a boolean if a field has been set.

### GetCheckId

`func (o *ResultsFeedback) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *ResultsFeedback) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *ResultsFeedback) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *ResultsFeedback) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetReportId

`func (o *ResultsFeedback) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ResultsFeedback) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ResultsFeedback) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *ResultsFeedback) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetFeedbackNotes

`func (o *ResultsFeedback) GetFeedbackNotes() string`

GetFeedbackNotes returns the FeedbackNotes field if non-nil, zero value otherwise.

### GetFeedbackNotesOk

`func (o *ResultsFeedback) GetFeedbackNotesOk() (*string, bool)`

GetFeedbackNotesOk returns a tuple with the FeedbackNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackNotes

`func (o *ResultsFeedback) SetFeedbackNotes(v string)`

SetFeedbackNotes sets FeedbackNotes field to given value.

### HasFeedbackNotes

`func (o *ResultsFeedback) HasFeedbackNotes() bool`

HasFeedbackNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


