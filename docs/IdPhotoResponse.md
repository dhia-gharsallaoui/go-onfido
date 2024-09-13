# IdPhotoResponse

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

### NewIdPhotoResponse

`func NewIdPhotoResponse(id string, ) *IdPhotoResponse`

NewIdPhotoResponse instantiates a new IdPhotoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdPhotoResponseWithDefaults

`func NewIdPhotoResponseWithDefaults() *IdPhotoResponse`

NewIdPhotoResponseWithDefaults instantiates a new IdPhotoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdPhotoResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdPhotoResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdPhotoResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *IdPhotoResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdPhotoResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdPhotoResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdPhotoResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *IdPhotoResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IdPhotoResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IdPhotoResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IdPhotoResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetDownloadHref

`func (o *IdPhotoResponse) GetDownloadHref() string`

GetDownloadHref returns the DownloadHref field if non-nil, zero value otherwise.

### GetDownloadHrefOk

`func (o *IdPhotoResponse) GetDownloadHrefOk() (*string, bool)`

GetDownloadHrefOk returns a tuple with the DownloadHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadHref

`func (o *IdPhotoResponse) SetDownloadHref(v string)`

SetDownloadHref sets DownloadHref field to given value.

### HasDownloadHref

`func (o *IdPhotoResponse) HasDownloadHref() bool`

HasDownloadHref returns a boolean if a field has been set.

### GetFileName

`func (o *IdPhotoResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *IdPhotoResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *IdPhotoResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *IdPhotoResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileType

`func (o *IdPhotoResponse) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *IdPhotoResponse) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *IdPhotoResponse) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *IdPhotoResponse) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFileSize

`func (o *IdPhotoResponse) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *IdPhotoResponse) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *IdPhotoResponse) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *IdPhotoResponse) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


