# AddressShared

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

### NewAddressShared

`func NewAddressShared(postcode string, country CountryCodes, ) *AddressShared`

NewAddressShared instantiates a new AddressShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressSharedWithDefaults

`func NewAddressSharedWithDefaults() *AddressShared`

NewAddressSharedWithDefaults instantiates a new AddressShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlatNumber

`func (o *AddressShared) GetFlatNumber() string`

GetFlatNumber returns the FlatNumber field if non-nil, zero value otherwise.

### GetFlatNumberOk

`func (o *AddressShared) GetFlatNumberOk() (*string, bool)`

GetFlatNumberOk returns a tuple with the FlatNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatNumber

`func (o *AddressShared) SetFlatNumber(v string)`

SetFlatNumber sets FlatNumber field to given value.

### HasFlatNumber

`func (o *AddressShared) HasFlatNumber() bool`

HasFlatNumber returns a boolean if a field has been set.

### GetBuildingNumber

`func (o *AddressShared) GetBuildingNumber() string`

GetBuildingNumber returns the BuildingNumber field if non-nil, zero value otherwise.

### GetBuildingNumberOk

`func (o *AddressShared) GetBuildingNumberOk() (*string, bool)`

GetBuildingNumberOk returns a tuple with the BuildingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingNumber

`func (o *AddressShared) SetBuildingNumber(v string)`

SetBuildingNumber sets BuildingNumber field to given value.

### HasBuildingNumber

`func (o *AddressShared) HasBuildingNumber() bool`

HasBuildingNumber returns a boolean if a field has been set.

### GetBuildingName

`func (o *AddressShared) GetBuildingName() string`

GetBuildingName returns the BuildingName field if non-nil, zero value otherwise.

### GetBuildingNameOk

`func (o *AddressShared) GetBuildingNameOk() (*string, bool)`

GetBuildingNameOk returns a tuple with the BuildingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingName

`func (o *AddressShared) SetBuildingName(v string)`

SetBuildingName sets BuildingName field to given value.

### HasBuildingName

`func (o *AddressShared) HasBuildingName() bool`

HasBuildingName returns a boolean if a field has been set.

### GetStreet

`func (o *AddressShared) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AddressShared) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AddressShared) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *AddressShared) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetSubStreet

`func (o *AddressShared) GetSubStreet() string`

GetSubStreet returns the SubStreet field if non-nil, zero value otherwise.

### GetSubStreetOk

`func (o *AddressShared) GetSubStreetOk() (*string, bool)`

GetSubStreetOk returns a tuple with the SubStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubStreet

`func (o *AddressShared) SetSubStreet(v string)`

SetSubStreet sets SubStreet field to given value.

### HasSubStreet

`func (o *AddressShared) HasSubStreet() bool`

HasSubStreet returns a boolean if a field has been set.

### GetTown

`func (o *AddressShared) GetTown() string`

GetTown returns the Town field if non-nil, zero value otherwise.

### GetTownOk

`func (o *AddressShared) GetTownOk() (*string, bool)`

GetTownOk returns a tuple with the Town field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTown

`func (o *AddressShared) SetTown(v string)`

SetTown sets Town field to given value.

### HasTown

`func (o *AddressShared) HasTown() bool`

HasTown returns a boolean if a field has been set.

### GetPostcode

`func (o *AddressShared) GetPostcode() string`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *AddressShared) GetPostcodeOk() (*string, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *AddressShared) SetPostcode(v string)`

SetPostcode sets Postcode field to given value.


### GetCountry

`func (o *AddressShared) GetCountry() CountryCodes`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressShared) GetCountryOk() (*CountryCodes, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressShared) SetCountry(v CountryCodes)`

SetCountry sets Country field to given value.


### GetState

`func (o *AddressShared) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AddressShared) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AddressShared) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AddressShared) HasState() bool`

HasState returns a boolean if a field has been set.

### GetLine1

`func (o *AddressShared) GetLine1() string`

GetLine1 returns the Line1 field if non-nil, zero value otherwise.

### GetLine1Ok

`func (o *AddressShared) GetLine1Ok() (*string, bool)`

GetLine1Ok returns a tuple with the Line1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine1

`func (o *AddressShared) SetLine1(v string)`

SetLine1 sets Line1 field to given value.

### HasLine1

`func (o *AddressShared) HasLine1() bool`

HasLine1 returns a boolean if a field has been set.

### SetLine1Nil

`func (o *AddressShared) SetLine1Nil(b bool)`

 SetLine1Nil sets the value for Line1 to be an explicit nil

### UnsetLine1
`func (o *AddressShared) UnsetLine1()`

UnsetLine1 ensures that no value is present for Line1, not even an explicit nil
### GetLine2

`func (o *AddressShared) GetLine2() string`

GetLine2 returns the Line2 field if non-nil, zero value otherwise.

### GetLine2Ok

`func (o *AddressShared) GetLine2Ok() (*string, bool)`

GetLine2Ok returns a tuple with the Line2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine2

`func (o *AddressShared) SetLine2(v string)`

SetLine2 sets Line2 field to given value.

### HasLine2

`func (o *AddressShared) HasLine2() bool`

HasLine2 returns a boolean if a field has been set.

### SetLine2Nil

`func (o *AddressShared) SetLine2Nil(b bool)`

 SetLine2Nil sets the value for Line2 to be an explicit nil

### UnsetLine2
`func (o *AddressShared) UnsetLine2()`

UnsetLine2 ensures that no value is present for Line2, not even an explicit nil
### GetLine3

`func (o *AddressShared) GetLine3() string`

GetLine3 returns the Line3 field if non-nil, zero value otherwise.

### GetLine3Ok

`func (o *AddressShared) GetLine3Ok() (*string, bool)`

GetLine3Ok returns a tuple with the Line3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine3

`func (o *AddressShared) SetLine3(v string)`

SetLine3 sets Line3 field to given value.

### HasLine3

`func (o *AddressShared) HasLine3() bool`

HasLine3 returns a boolean if a field has been set.

### SetLine3Nil

`func (o *AddressShared) SetLine3Nil(b bool)`

 SetLine3Nil sets the value for Line3 to be an explicit nil

### UnsetLine3
`func (o *AddressShared) UnsetLine3()`

UnsetLine3 ensures that no value is present for Line3, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


