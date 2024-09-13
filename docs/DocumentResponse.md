# DocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the document | 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the document was uploaded | [optional] 
**Href** | Pointer to **string** | The uri of this resource | [optional] 
**DownloadHref** | Pointer to **string** | The uri that can be used to download the document | [optional] 
**FileName** | Pointer to **string** | The name of the uploaded file | [optional] 
**FileSize** | Pointer to **int32** | The size of the file in bytes | [optional] 

## Methods

### NewDocumentResponse

`func NewDocumentResponse(id string, ) *DocumentResponse`

NewDocumentResponse instantiates a new DocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentResponseWithDefaults

`func NewDocumentResponseWithDefaults() *DocumentResponse`

NewDocumentResponseWithDefaults instantiates a new DocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DocumentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *DocumentResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DocumentResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DocumentResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DocumentResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetDownloadHref

`func (o *DocumentResponse) GetDownloadHref() string`

GetDownloadHref returns the DownloadHref field if non-nil, zero value otherwise.

### GetDownloadHrefOk

`func (o *DocumentResponse) GetDownloadHrefOk() (*string, bool)`

GetDownloadHrefOk returns a tuple with the DownloadHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadHref

`func (o *DocumentResponse) SetDownloadHref(v string)`

SetDownloadHref sets DownloadHref field to given value.

### HasDownloadHref

`func (o *DocumentResponse) HasDownloadHref() bool`

HasDownloadHref returns a boolean if a field has been set.

### GetFileName

`func (o *DocumentResponse) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *DocumentResponse) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *DocumentResponse) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *DocumentResponse) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *DocumentResponse) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *DocumentResponse) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *DocumentResponse) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *DocumentResponse) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


