# ExtractRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | **string** | The unique identifier of the uploaded document to run extraction on | 

## Methods

### NewExtractRequest

`func NewExtractRequest(documentId string, ) *ExtractRequest`

NewExtractRequest instantiates a new ExtractRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractRequestWithDefaults

`func NewExtractRequestWithDefaults() *ExtractRequest`

NewExtractRequestWithDefaults instantiates a new ExtractRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *ExtractRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *ExtractRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *ExtractRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


