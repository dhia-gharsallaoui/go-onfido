# LocationBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | The applicant&#39;s ip address. | [optional] 
**CountryOfResidence** | Pointer to [**CountryCodes**](CountryCodes.md) | The applicant&#39;s country of residence in 3-letter ISO code. | [optional] 

## Methods

### NewLocationBuilder

`func NewLocationBuilder() *LocationBuilder`

NewLocationBuilder instantiates a new LocationBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationBuilderWithDefaults

`func NewLocationBuilderWithDefaults() *LocationBuilder`

NewLocationBuilderWithDefaults instantiates a new LocationBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *LocationBuilder) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LocationBuilder) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LocationBuilder) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LocationBuilder) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetCountryOfResidence

`func (o *LocationBuilder) GetCountryOfResidence() CountryCodes`

GetCountryOfResidence returns the CountryOfResidence field if non-nil, zero value otherwise.

### GetCountryOfResidenceOk

`func (o *LocationBuilder) GetCountryOfResidenceOk() (*CountryCodes, bool)`

GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfResidence

`func (o *LocationBuilder) SetCountryOfResidence(v CountryCodes)`

SetCountryOfResidence sets CountryOfResidence field to given value.

### HasCountryOfResidence

`func (o *LocationBuilder) HasCountryOfResidence() bool`

HasCountryOfResidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


