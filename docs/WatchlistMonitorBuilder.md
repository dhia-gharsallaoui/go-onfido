# WatchlistMonitorBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **string** | The ID for the applicant associated with the monitor. | 
**ReportName** | **string** | The name of the report type the monitor creates. | 
**Tags** | Pointer to **[]string** | A list of tags associated with this monitor. These tags will be applied to each check this monitor creates. | [optional] 

## Methods

### NewWatchlistMonitorBuilder

`func NewWatchlistMonitorBuilder(applicantId string, reportName string, ) *WatchlistMonitorBuilder`

NewWatchlistMonitorBuilder instantiates a new WatchlistMonitorBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistMonitorBuilderWithDefaults

`func NewWatchlistMonitorBuilderWithDefaults() *WatchlistMonitorBuilder`

NewWatchlistMonitorBuilderWithDefaults instantiates a new WatchlistMonitorBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *WatchlistMonitorBuilder) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *WatchlistMonitorBuilder) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *WatchlistMonitorBuilder) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetReportName

`func (o *WatchlistMonitorBuilder) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *WatchlistMonitorBuilder) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *WatchlistMonitorBuilder) SetReportName(v string)`

SetReportName sets ReportName field to given value.


### GetTags

`func (o *WatchlistMonitorBuilder) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WatchlistMonitorBuilder) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WatchlistMonitorBuilder) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WatchlistMonitorBuilder) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


