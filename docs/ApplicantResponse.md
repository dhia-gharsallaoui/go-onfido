# ApplicantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The applicant&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The applicant&#39;s surname | [optional] 
**Id** | **string** | The unique identifier for the applicant. | 
**CreatedAt** | Pointer to **time.Time** | The date and time when this applicant was created. | [optional] 
**DeleteAt** | Pointer to **time.Time** | The date and time when this applicant is scheduled to be deleted. | [optional] 
**Href** | Pointer to **string** | The uri of this resource. | [optional] 
**Sandbox** | Pointer to **bool** |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 

## Methods

### NewApplicantResponse

`func NewApplicantResponse(id string, ) *ApplicantResponse`

NewApplicantResponse instantiates a new ApplicantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicantResponseWithDefaults

`func NewApplicantResponseWithDefaults() *ApplicantResponse`

NewApplicantResponseWithDefaults instantiates a new ApplicantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *ApplicantResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ApplicantResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ApplicantResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ApplicantResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ApplicantResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ApplicantResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ApplicantResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ApplicantResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetId

`func (o *ApplicantResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicantResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicantResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ApplicantResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicantResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicantResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicantResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeleteAt

`func (o *ApplicantResponse) GetDeleteAt() time.Time`

GetDeleteAt returns the DeleteAt field if non-nil, zero value otherwise.

### GetDeleteAtOk

`func (o *ApplicantResponse) GetDeleteAtOk() (*time.Time, bool)`

GetDeleteAtOk returns a tuple with the DeleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAt

`func (o *ApplicantResponse) SetDeleteAt(v time.Time)`

SetDeleteAt sets DeleteAt field to given value.

### HasDeleteAt

`func (o *ApplicantResponse) HasDeleteAt() bool`

HasDeleteAt returns a boolean if a field has been set.

### GetHref

`func (o *ApplicantResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApplicantResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApplicantResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApplicantResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetSandbox

`func (o *ApplicantResponse) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *ApplicantResponse) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *ApplicantResponse) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *ApplicantResponse) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetAddress

`func (o *ApplicantResponse) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ApplicantResponse) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ApplicantResponse) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ApplicantResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLocation

`func (o *ApplicantResponse) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApplicantResponse) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApplicantResponse) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApplicantResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


