# LocationShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | The applicant&#39;s ip address. | [optional] 
**CountryOfResidence** | Pointer to [**CountryCodes**](CountryCodes.md) | The applicant&#39;s country of residence in 3-letter ISO code. | [optional] 

## Methods

### NewLocationShared

`func NewLocationShared() *LocationShared`

NewLocationShared instantiates a new LocationShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationSharedWithDefaults

`func NewLocationSharedWithDefaults() *LocationShared`

NewLocationSharedWithDefaults instantiates a new LocationShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *LocationShared) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LocationShared) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LocationShared) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *LocationShared) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetCountryOfResidence

`func (o *LocationShared) GetCountryOfResidence() CountryCodes`

GetCountryOfResidence returns the CountryOfResidence field if non-nil, zero value otherwise.

### GetCountryOfResidenceOk

`func (o *LocationShared) GetCountryOfResidenceOk() (*CountryCodes, bool)`

GetCountryOfResidenceOk returns a tuple with the CountryOfResidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfResidence

`func (o *LocationShared) SetCountryOfResidence(v CountryCodes)`

SetCountryOfResidence sets CountryOfResidence field to given value.

### HasCountryOfResidence

`func (o *LocationShared) HasCountryOfResidence() bool`

HasCountryOfResidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


