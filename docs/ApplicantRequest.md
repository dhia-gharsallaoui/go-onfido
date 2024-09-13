# ApplicantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consents** | Pointer to [**ConsentsBuilder**](ConsentsBuilder.md) |  | [optional] 
**Address** | Pointer to [**AddressBuilder**](AddressBuilder.md) |  | [optional] 
**Location** | Pointer to [**LocationBuilder**](LocationBuilder.md) |  | [optional] 

## Methods

### NewApplicantRequest

`func NewApplicantRequest() *ApplicantRequest`

NewApplicantRequest instantiates a new ApplicantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicantRequestWithDefaults

`func NewApplicantRequestWithDefaults() *ApplicantRequest`

NewApplicantRequestWithDefaults instantiates a new ApplicantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsents

`func (o *ApplicantRequest) GetConsents() ConsentsBuilder`

GetConsents returns the Consents field if non-nil, zero value otherwise.

### GetConsentsOk

`func (o *ApplicantRequest) GetConsentsOk() (*ConsentsBuilder, bool)`

GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsents

`func (o *ApplicantRequest) SetConsents(v ConsentsBuilder)`

SetConsents sets Consents field to given value.

### HasConsents

`func (o *ApplicantRequest) HasConsents() bool`

HasConsents returns a boolean if a field has been set.

### GetAddress

`func (o *ApplicantRequest) GetAddress() AddressBuilder`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ApplicantRequest) GetAddressOk() (*AddressBuilder, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ApplicantRequest) SetAddress(v AddressBuilder)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ApplicantRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLocation

`func (o *ApplicantRequest) GetLocation() LocationBuilder`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApplicantRequest) GetLocationOk() (*LocationBuilder, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApplicantRequest) SetLocation(v LocationBuilder)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApplicantRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


