# AddressBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlatNumber** | Pointer to **string** | The flat number of this address | [optional] 
**BuildingNumber** | Pointer to **string** | The building number of this address | [optional] 
**BuildingName** | Pointer to **string** | The building name of this address | [optional] 
**Street** | Pointer to **string** | The street of the applicant&#39;s address | [optional] 
**SubStreet** | Pointer to **string** | The sub-street of the applicant&#39;s address | [optional] 
**Town** | Pointer to **string** | The town of the applicant&#39;s address | [optional] 
**Postcode** | **string** | The postcode or ZIP of the applicant&#39;s address | 
**Country** | [**CountryCodes**](CountryCodes.md) | The 3 character ISO country code of this address. For example, GBR is the country code for the United Kingdom | 
**State** | Pointer to **string** | The address state. US states must use the USPS abbreviation (see also ISO 3166-2:US), for example AK, CA, or TX. | [optional] 
**Line1** | Pointer to **NullableString** | Line 1 of the applicant&#39;s address | [optional] 
**Line2** | Pointer to **NullableString** | Line 2 of the applicant&#39;s address | [optional] 
**Line3** | Pointer to **NullableString** | Line 3 of the applicant&#39;s address | [optional] 

## Methods

### NewAddressBuilder

`func NewAddressBuilder(postcode string, country CountryCodes, ) *AddressBuilder`

NewAddressBuilder instantiates a new AddressBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressBuilderWithDefaults

`func NewAddressBuilderWithDefaults() *AddressBuilder`

NewAddressBuilderWithDefaults instantiates a new AddressBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlatNumber

`func (o *AddressBuilder) GetFlatNumber() string`

GetFlatNumber returns the FlatNumber field if non-nil, zero value otherwise.

### GetFlatNumberOk

`func (o *AddressBuilder) GetFlatNumberOk() (*string, bool)`

GetFlatNumberOk returns a tuple with the FlatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatNumber

`func (o *AddressBuilder) SetFlatNumber(v string)`

SetFlatNumber sets FlatNumber field to given value.

### HasFlatNumber

`func (o *AddressBuilder) HasFlatNumber() bool`

HasFlatNumber returns a boolean if a field has been set.

### GetBuildingNumber

`func (o *AddressBuilder) GetBuildingNumber() string`

GetBuildingNumber returns the BuildingNumber field if non-nil, zero value otherwise.

### GetBuildingNumberOk

`func (o *AddressBuilder) GetBuildingNumberOk() (*string, bool)`

GetBuildingNumberOk returns a tuple with the BuildingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumber

`func (o *AddressBuilder) SetBuildingNumber(v string)`

SetBuildingNumber sets BuildingNumber field to given value.

### HasBuildingNumber

`func (o *AddressBuilder) HasBuildingNumber() bool`

HasBuildingNumber returns a boolean if a field has been set.

### GetBuildingName

`func (o *AddressBuilder) GetBuildingName() string`

GetBuildingName returns the BuildingName field if non-nil, zero value otherwise.

### GetBuildingNameOk

`func (o *AddressBuilder) GetBuildingNameOk() (*string, bool)`

GetBuildingNameOk returns a tuple with the BuildingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingName

`func (o *AddressBuilder) SetBuildingName(v string)`

SetBuildingName sets BuildingName field to given value.

### HasBuildingName

`func (o *AddressBuilder) HasBuildingName() bool`

HasBuildingName returns a boolean if a field has been set.

### GetStreet

`func (o *AddressBuilder) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AddressBuilder) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AddressBuilder) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *AddressBuilder) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetSubStreet

`func (o *AddressBuilder) GetSubStreet() string`

GetSubStreet returns the SubStreet field if non-nil, zero value otherwise.

### GetSubStreetOk

`func (o *AddressBuilder) GetSubStreetOk() (*string, bool)`

GetSubStreetOk returns a tuple with the SubStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStreet

`func (o *AddressBuilder) SetSubStreet(v string)`

SetSubStreet sets SubStreet field to given value.

### HasSubStreet

`func (o *AddressBuilder) HasSubStreet() bool`

HasSubStreet returns a boolean if a field has been set.

### GetTown

`func (o *AddressBuilder) GetTown() string`

GetTown returns the Town field if non-nil, zero value otherwise.

### GetTownOk

`func (o *AddressBuilder) GetTownOk() (*string, bool)`

GetTownOk returns a tuple with the Town field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTown

`func (o *AddressBuilder) SetTown(v string)`

SetTown sets Town field to given value.

### HasTown

`func (o *AddressBuilder) HasTown() bool`

HasTown returns a boolean if a field has been set.

### GetPostcode

`func (o *AddressBuilder) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *AddressBuilder) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *AddressBuilder) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetCountry

`func (o *AddressBuilder) GetCountry() CountryCodes`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressBuilder) GetCountryOk() (*CountryCodes, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressBuilder) SetCountry(v CountryCodes)`

SetCountry sets Country field to given value.


### GetState

`func (o *AddressBuilder) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressBuilder) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressBuilder) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressBuilder) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLine1

`func (o *AddressBuilder) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *AddressBuilder) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *AddressBuilder) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *AddressBuilder) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *AddressBuilder) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *AddressBuilder) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *AddressBuilder) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *AddressBuilder) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *AddressBuilder) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *AddressBuilder) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *AddressBuilder) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *AddressBuilder) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *AddressBuilder) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *AddressBuilder) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *AddressBuilder) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *AddressBuilder) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### SetLine3Nil

`func (o *AddressBuilder) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *AddressBuilder) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


