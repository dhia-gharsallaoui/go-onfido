# WatchlistMonitorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the monitor. | 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the monitor was created. | [optional] 
**DeletedAt** | Pointer to **time.Time** | The date and time at which the monitor was deleted. If the monitor is still active, this field will be null. | [optional] 
**IsSandbox** | Pointer to **bool** | Indicates whether the object was created in the sandbox or not. | [optional] [default to false]

## Methods

### NewWatchlistMonitorResponse

`func NewWatchlistMonitorResponse(id string, ) *WatchlistMonitorResponse`

NewWatchlistMonitorResponse instantiates a new WatchlistMonitorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchlistMonitorResponseWithDefaults

`func NewWatchlistMonitorResponseWithDefaults() *WatchlistMonitorResponse`

NewWatchlistMonitorResponseWithDefaults instantiates a new WatchlistMonitorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WatchlistMonitorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WatchlistMonitorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WatchlistMonitorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *WatchlistMonitorResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WatchlistMonitorResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WatchlistMonitorResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WatchlistMonitorResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *WatchlistMonitorResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WatchlistMonitorResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WatchlistMonitorResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WatchlistMonitorResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetIsSandbox

`func (o *WatchlistMonitorResponse) GetIsSandbox() bool`

GetIsSandbox returns the IsSandbox field if non-nil, zero value otherwise.

### GetIsSandboxOk

`func (o *WatchlistMonitorResponse) GetIsSandboxOk() (*bool, bool)`

GetIsSandboxOk returns a tuple with the IsSandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSandbox

`func (o *WatchlistMonitorResponse) SetIsSandbox(v bool)`

SetIsSandbox sets IsSandbox field to given value.

### HasIsSandbox

`func (o *WatchlistMonitorResponse) HasIsSandbox() bool`

HasIsSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


