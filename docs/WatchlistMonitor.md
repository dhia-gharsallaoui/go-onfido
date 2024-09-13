# WatchlistMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **string** | The ID for the applicant associated with the monitor. | 
**ReportName** | **string** | The name of the report type the monitor creates. | 
**Tags** | Pointer to **[]string** | A list of tags associated with this monitor. These tags will be applied to each check this monitor creates. | [optional] 
**Id** | **string** | The unique identifier for the monitor. | 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the monitor was created. | [optional] 
**DeletedAt** | Pointer to **time.Time** | The date and time at which the monitor was deleted. If the monitor is still active, this field will be null. | [optional] 
**IsSandbox** | Pointer to **bool** | Indicates whether the object was created in the sandbox or not. | [optional] [default to false]

## Methods

### NewWatchlistMonitor

`func NewWatchlistMonitor(applicantId string, reportName string, id string, ) *WatchlistMonitor`

NewWatchlistMonitor instantiates a new WatchlistMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistMonitorWithDefaults

`func NewWatchlistMonitorWithDefaults() *WatchlistMonitor`

NewWatchlistMonitorWithDefaults instantiates a new WatchlistMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *WatchlistMonitor) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *WatchlistMonitor) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *WatchlistMonitor) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetReportName

`func (o *WatchlistMonitor) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *WatchlistMonitor) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *WatchlistMonitor) SetReportName(v string)`

SetReportName sets ReportName field to given value.


### GetTags

`func (o *WatchlistMonitor) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WatchlistMonitor) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WatchlistMonitor) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WatchlistMonitor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetId

`func (o *WatchlistMonitor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WatchlistMonitor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WatchlistMonitor) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *WatchlistMonitor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WatchlistMonitor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WatchlistMonitor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WatchlistMonitor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *WatchlistMonitor) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WatchlistMonitor) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WatchlistMonitor) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WatchlistMonitor) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetIsSandbox

`func (o *WatchlistMonitor) GetIsSandbox() bool`

GetIsSandbox returns the IsSandbox field if non-nil, zero value otherwise.

### GetIsSandboxOk

`func (o *WatchlistMonitor) GetIsSandboxOk() (*bool, bool)`

GetIsSandboxOk returns a tuple with the IsSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSandbox

`func (o *WatchlistMonitor) SetIsSandbox(v bool)`

SetIsSandbox sets IsSandbox field to given value.

### HasIsSandbox

`func (o *WatchlistMonitor) HasIsSandbox() bool`

HasIsSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


