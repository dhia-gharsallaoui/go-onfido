# DocumentShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileType** | Pointer to **string** | The file type of the uploaded file | [optional] 
**Type** | Pointer to **string** | The type of document | [optional] 
**Side** | Pointer to **string** | The side of the document, if applicable. The possible values are front and back | [optional] 
**IssuingCountry** | Pointer to [**CountryCodes**](CountryCodes.md) | The issuing country of the document, a 3-letter ISO code. | [optional] 
**ApplicantId** | Pointer to **string** | The ID of the applicant whose document is being uploaded. | [optional] 

## Methods

### NewDocumentShared

`func NewDocumentShared() *DocumentShared`

NewDocumentShared instantiates a new DocumentShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentSharedWithDefaults

`func NewDocumentSharedWithDefaults() *DocumentShared`

NewDocumentSharedWithDefaults instantiates a new DocumentShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileType

`func (o *DocumentShared) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *DocumentShared) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *DocumentShared) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *DocumentShared) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetType

`func (o *DocumentShared) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DocumentShared) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DocumentShared) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DocumentShared) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSide

`func (o *DocumentShared) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *DocumentShared) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *DocumentShared) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *DocumentShared) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetIssuingCountry

`func (o *DocumentShared) GetIssuingCountry() CountryCodes`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *DocumentShared) GetIssuingCountryOk() (*CountryCodes, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *DocumentShared) SetIssuingCountry(v CountryCodes)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *DocumentShared) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### GetApplicantId

`func (o *DocumentShared) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *DocumentShared) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *DocumentShared) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.

### HasApplicantId

`func (o *DocumentShared) HasApplicantId() bool`

HasApplicantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


