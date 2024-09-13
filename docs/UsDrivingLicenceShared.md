# UsDrivingLicenceShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdNumber** | **string** | Driving licence ID number | 
**IssueState** | **string** | Two letter code of issuing state (state-issued driving licenses only) | 
**AddressLine1** | Pointer to **string** | Line 1 of the address | [optional] 
**AddressLine2** | Pointer to **string** | Line 2 of the address | [optional] 
**City** | Pointer to **string** | The city of the owner&#39;s address | [optional] 
**DateOfBirth** | Pointer to **string** | Date of birth in yyyy-mm-dd format | [optional] 
**DocumentCategory** | Pointer to **string** | Document category. | [optional] 
**ExpirationDate** | Pointer to **string** | Expiration date of the driving licence in yyyy-mm-dd format | [optional] 
**EyeColorCode** | Pointer to **string** | Eye color code. | [optional] 
**FirstName** | Pointer to **string** | The owner&#39;s first name | [optional] 
**Gender** | Pointer to **string** |  | [optional] [readonly] 
**IssueDate** | Pointer to **string** | Issue date in yyyy-mm-dd format | [optional] 
**LastName** | Pointer to **string** | The owner&#39;s surname | [optional] 
**MiddleName** | Pointer to **string** | The owner&#39;s middle name | [optional] 
**NameSuffix** | Pointer to **string** | The owner&#39;s name suffix | [optional] 
**PostalCode** | Pointer to **string** | The postcode or ZIP of the owner&#39;s address | [optional] 
**State** | Pointer to **string** | 2-characters state code | [optional] 
**WeightMeasure** | Pointer to **int32** | Weight in pounds | [optional] 

## Methods

### NewUsDrivingLicenceShared

`func NewUsDrivingLicenceShared(idNumber string, issueState string, ) *UsDrivingLicenceShared`

NewUsDrivingLicenceShared instantiates a new UsDrivingLicenceShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsDrivingLicenceSharedWithDefaults

`func NewUsDrivingLicenceSharedWithDefaults() *UsDrivingLicenceShared`

NewUsDrivingLicenceSharedWithDefaults instantiates a new UsDrivingLicenceShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *UsDrivingLicenceShared) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *UsDrivingLicenceShared) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *UsDrivingLicenceShared) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.


### GetIssueState

`func (o *UsDrivingLicenceShared) GetIssueState() string`

GetIssueState returns the IssueState field if non-nil, zero value otherwise.

### GetIssueStateOk

`func (o *UsDrivingLicenceShared) GetIssueStateOk() (*string, bool)`

GetIssueStateOk returns a tuple with the IssueState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueState

`func (o *UsDrivingLicenceShared) SetIssueState(v string)`

SetIssueState sets IssueState field to given value.


### GetAddressLine1

`func (o *UsDrivingLicenceShared) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *UsDrivingLicenceShared) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *UsDrivingLicenceShared) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *UsDrivingLicenceShared) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *UsDrivingLicenceShared) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *UsDrivingLicenceShared) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *UsDrivingLicenceShared) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *UsDrivingLicenceShared) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *UsDrivingLicenceShared) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UsDrivingLicenceShared) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UsDrivingLicenceShared) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UsDrivingLicenceShared) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *UsDrivingLicenceShared) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UsDrivingLicenceShared) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UsDrivingLicenceShared) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *UsDrivingLicenceShared) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDocumentCategory

`func (o *UsDrivingLicenceShared) GetDocumentCategory() string`

GetDocumentCategory returns the DocumentCategory field if non-nil, zero value otherwise.

### GetDocumentCategoryOk

`func (o *UsDrivingLicenceShared) GetDocumentCategoryOk() (*string, bool)`

GetDocumentCategoryOk returns a tuple with the DocumentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCategory

`func (o *UsDrivingLicenceShared) SetDocumentCategory(v string)`

SetDocumentCategory sets DocumentCategory field to given value.

### HasDocumentCategory

`func (o *UsDrivingLicenceShared) HasDocumentCategory() bool`

HasDocumentCategory returns a boolean if a field has been set.

### GetExpirationDate

`func (o *UsDrivingLicenceShared) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *UsDrivingLicenceShared) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *UsDrivingLicenceShared) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *UsDrivingLicenceShared) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetEyeColorCode

`func (o *UsDrivingLicenceShared) GetEyeColorCode() string`

GetEyeColorCode returns the EyeColorCode field if non-nil, zero value otherwise.

### GetEyeColorCodeOk

`func (o *UsDrivingLicenceShared) GetEyeColorCodeOk() (*string, bool)`

GetEyeColorCodeOk returns a tuple with the EyeColorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEyeColorCode

`func (o *UsDrivingLicenceShared) SetEyeColorCode(v string)`

SetEyeColorCode sets EyeColorCode field to given value.

### HasEyeColorCode

`func (o *UsDrivingLicenceShared) HasEyeColorCode() bool`

HasEyeColorCode returns a boolean if a field has been set.

### GetFirstName

`func (o *UsDrivingLicenceShared) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsDrivingLicenceShared) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsDrivingLicenceShared) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UsDrivingLicenceShared) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGender

`func (o *UsDrivingLicenceShared) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UsDrivingLicenceShared) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UsDrivingLicenceShared) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UsDrivingLicenceShared) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetIssueDate

`func (o *UsDrivingLicenceShared) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *UsDrivingLicenceShared) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *UsDrivingLicenceShared) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *UsDrivingLicenceShared) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetLastName

`func (o *UsDrivingLicenceShared) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsDrivingLicenceShared) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsDrivingLicenceShared) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UsDrivingLicenceShared) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *UsDrivingLicenceShared) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UsDrivingLicenceShared) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UsDrivingLicenceShared) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UsDrivingLicenceShared) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNameSuffix

`func (o *UsDrivingLicenceShared) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *UsDrivingLicenceShared) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *UsDrivingLicenceShared) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *UsDrivingLicenceShared) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### GetPostalCode

`func (o *UsDrivingLicenceShared) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UsDrivingLicenceShared) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UsDrivingLicenceShared) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UsDrivingLicenceShared) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *UsDrivingLicenceShared) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UsDrivingLicenceShared) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UsDrivingLicenceShared) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UsDrivingLicenceShared) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWeightMeasure

`func (o *UsDrivingLicenceShared) GetWeightMeasure() int32`

GetWeightMeasure returns the WeightMeasure field if non-nil, zero value otherwise.

### GetWeightMeasureOk

`func (o *UsDrivingLicenceShared) GetWeightMeasureOk() (*int32, bool)`

GetWeightMeasureOk returns a tuple with the WeightMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightMeasure

`func (o *UsDrivingLicenceShared) SetWeightMeasure(v int32)`

SetWeightMeasure sets WeightMeasure field to given value.

### HasWeightMeasure

`func (o *UsDrivingLicenceShared) HasWeightMeasure() bool`

HasWeightMeasure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


