# MotionCapture

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier for the motion capture. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the motion capture was uploaded. | [optional] 
**Href** | Pointer to **string** | The uri of this resource. | [optional] 
**DownloadHref** | Pointer to **string** | The uri that can be used to download the motion capture. | [optional] 
**FileName** | Pointer to **string** | The name of the uploaded file. | [optional] 
**FileSize** | Pointer to **int32** | The size of the file in bytes. | [optional] 
**FileType** | Pointer to **string** | The file type of the uploaded file. | [optional] 

## Methods

### NewMotionCapture

`func NewMotionCapture() *MotionCapture`

NewMotionCapture instantiates a new MotionCapture object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotionCaptureWithDefaults

`func NewMotionCaptureWithDefaults() *MotionCapture`

NewMotionCaptureWithDefaults instantiates a new MotionCapture object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MotionCapture) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MotionCapture) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MotionCapture) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MotionCapture) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MotionCapture) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MotionCapture) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MotionCapture) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MotionCapture) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *MotionCapture) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MotionCapture) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MotionCapture) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MotionCapture) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetDownloadHref

`func (o *MotionCapture) GetDownloadHref() string`

GetDownloadHref returns the DownloadHref field if non-nil, zero value otherwise.

### GetDownloadHrefOk

`func (o *MotionCapture) GetDownloadHrefOk() (*string, bool)`

GetDownloadHrefOk returns a tuple with the DownloadHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadHref

`func (o *MotionCapture) SetDownloadHref(v string)`

SetDownloadHref sets DownloadHref field to given value.

### HasDownloadHref

`func (o *MotionCapture) HasDownloadHref() bool`

HasDownloadHref returns a boolean if a field has been set.

### GetFileName

`func (o *MotionCapture) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MotionCapture) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MotionCapture) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MotionCapture) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *MotionCapture) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *MotionCapture) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *MotionCapture) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *MotionCapture) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileType

`func (o *MotionCapture) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *MotionCapture) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *MotionCapture) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *MotionCapture) HasFileType() bool`

HasFileType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


