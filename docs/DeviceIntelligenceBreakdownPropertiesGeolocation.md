# DeviceIntelligenceBreakdownPropertiesGeolocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | City location of the IP address. | [optional] 
**Region** | Pointer to **string** | Region location of the IP address. | [optional] 
**Country** | Pointer to [**CountryCodes**](CountryCodes.md) | Country location of the IP address in a three letter format. | [optional] 

## Methods

### NewDeviceIntelligenceBreakdownPropertiesGeolocation

`func NewDeviceIntelligenceBreakdownPropertiesGeolocation() *DeviceIntelligenceBreakdownPropertiesGeolocation`

NewDeviceIntelligenceBreakdownPropertiesGeolocation instantiates a new DeviceIntelligenceBreakdownPropertiesGeolocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceIntelligenceBreakdownPropertiesGeolocationWithDefaults

`func NewDeviceIntelligenceBreakdownPropertiesGeolocationWithDefaults() *DeviceIntelligenceBreakdownPropertiesGeolocation`

NewDeviceIntelligenceBreakdownPropertiesGeolocationWithDefaults instantiates a new DeviceIntelligenceBreakdownPropertiesGeolocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCountry

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetCountry() CountryCodes`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) GetCountryOk() (*CountryCodes, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) SetCountry(v CountryCodes)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *DeviceIntelligenceBreakdownPropertiesGeolocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


