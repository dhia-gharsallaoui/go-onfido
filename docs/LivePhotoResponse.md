# LivePhotoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the photo. | 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the photo was uploaded. | [optional] 
**Href** | Pointer to **string** | The uri of this resource. | [optional] 
**DownloadHref** | Pointer to **string** | The uri that can be used to download the photo. | [optional] 
**FileName** | Pointer to **string** | The name of the uploaded file. | [optional] 
**FileType** | Pointer to **string** | The file type of the uploaded file. | [optional] 
**FileSize** | Pointer to **int32** | The size of the file in bytes. | [optional] 

## Methods

### NewLivePhotoResponse

`func NewLivePhotoResponse(id string, ) *LivePhotoResponse`

NewLivePhotoResponse instantiates a new LivePhotoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLivePhotoResponseWithDefaults

`func NewLivePhotoResponseWithDefaults() *LivePhotoResponse`

NewLivePhotoResponseWithDefaults instantiates a new LivePhotoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LivePhotoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LivePhotoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LivePhotoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *LivePhotoResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LivePhotoResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LivePhotoResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LivePhotoResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *LivePhotoResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LivePhotoResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LivePhotoResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LivePhotoResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetDownloadHref

`func (o *LivePhotoResponse) GetDownloadHref() string`

GetDownloadHref returns the DownloadHref field if non-nil, zero value otherwise.

### GetDownloadHrefOk

`func (o *LivePhotoResponse) GetDownloadHrefOk() (*string, bool)`

GetDownloadHrefOk returns a tuple with the DownloadHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadHref

`func (o *LivePhotoResponse) SetDownloadHref(v string)`

SetDownloadHref sets DownloadHref field to given value.

### HasDownloadHref

`func (o *LivePhotoResponse) HasDownloadHref() bool`

HasDownloadHref returns a boolean if a field has been set.

### GetFileName

`func (o *LivePhotoResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *LivePhotoResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *LivePhotoResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *LivePhotoResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *LivePhotoResponse) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *LivePhotoResponse) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *LivePhotoResponse) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *LivePhotoResponse) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFileSize

`func (o *LivePhotoResponse) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *LivePhotoResponse) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *LivePhotoResponse) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *LivePhotoResponse) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


