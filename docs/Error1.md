# Error1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | The unique identifier of the uploaded document | [optional] 
**Error** | Pointer to [**ErrorProperties1**](ErrorProperties1.md) |  | [optional] 

## Methods

### NewError1

`func NewError1() *Error1`

NewError1 instantiates a new Error1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewError1WithDefaults

`func NewError1WithDefaults() *Error1`

NewError1WithDefaults instantiates a new Error1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *Error1) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *Error1) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *Error1) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *Error1) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetError

`func (o *Error1) GetError() ErrorProperties1`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Error1) GetErrorOk() (*ErrorProperties1, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Error1) SetError(v ErrorProperties1)`

SetError sets Error field to given value.

### HasError

`func (o *Error1) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


