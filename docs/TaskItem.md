# TaskItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier for the Task. | [optional] 
**TaskDefId** | Pointer to **string** | The identifier for the Task Definition. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time when the Task was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time when the Task was last updated. | [optional] 

## Methods

### NewTaskItem

`func NewTaskItem() *TaskItem`

NewTaskItem instantiates a new TaskItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskItemWithDefaults

`func NewTaskItemWithDefaults() *TaskItem`

NewTaskItemWithDefaults instantiates a new TaskItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaskDefId

`func (o *TaskItem) GetTaskDefId() string`

GetTaskDefId returns the TaskDefId field if non-nil, zero value otherwise.

### GetTaskDefIdOk

`func (o *TaskItem) GetTaskDefIdOk() (*string, bool)`

GetTaskDefIdOk returns a tuple with the TaskDefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskDefId

`func (o *TaskItem) SetTaskDefId(v string)`

SetTaskDefId sets TaskDefId field to given value.

### HasTaskDefId

`func (o *TaskItem) HasTaskDefId() bool`

HasTaskDefId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TaskItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TaskItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TaskItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TaskItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


