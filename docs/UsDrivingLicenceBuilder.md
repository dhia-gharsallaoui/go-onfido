# UsDrivingLicenceBuilder

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

### NewUsDrivingLicenceBuilder

`func NewUsDrivingLicenceBuilder(idNumber string, issueState string, ) *UsDrivingLicenceBuilder`

NewUsDrivingLicenceBuilder instantiates a new UsDrivingLicenceBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsDrivingLicenceBuilderWithDefaults

`func NewUsDrivingLicenceBuilderWithDefaults() *UsDrivingLicenceBuilder`

NewUsDrivingLicenceBuilderWithDefaults instantiates a new UsDrivingLicenceBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdNumber

`func (o *UsDrivingLicenceBuilder) GetIdNumber() string`

GetIdNumber returns the IdNumber field if non-nil, zero value otherwise.

### GetIdNumberOk

`func (o *UsDrivingLicenceBuilder) GetIdNumberOk() (*string, bool)`

GetIdNumberOk returns a tuple with the IdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumber

`func (o *UsDrivingLicenceBuilder) SetIdNumber(v string)`

SetIdNumber sets IdNumber field to given value.


### GetIssueState

`func (o *UsDrivingLicenceBuilder) GetIssueState() string`

GetIssueState returns the IssueState field if non-nil, zero value otherwise.

### GetIssueStateOk

`func (o *UsDrivingLicenceBuilder) GetIssueStateOk() (*string, bool)`

GetIssueStateOk returns a tuple with the IssueState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueState

`func (o *UsDrivingLicenceBuilder) SetIssueState(v string)`

SetIssueState sets IssueState field to given value.


### GetAddressLine1

`func (o *UsDrivingLicenceBuilder) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *UsDrivingLicenceBuilder) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *UsDrivingLicenceBuilder) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *UsDrivingLicenceBuilder) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetAddressLine2

`func (o *UsDrivingLicenceBuilder) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *UsDrivingLicenceBuilder) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *UsDrivingLicenceBuilder) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *UsDrivingLicenceBuilder) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *UsDrivingLicenceBuilder) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UsDrivingLicenceBuilder) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UsDrivingLicenceBuilder) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UsDrivingLicenceBuilder) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *UsDrivingLicenceBuilder) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *UsDrivingLicenceBuilder) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *UsDrivingLicenceBuilder) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *UsDrivingLicenceBuilder) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDocumentCategory

`func (o *UsDrivingLicenceBuilder) GetDocumentCategory() string`

GetDocumentCategory returns the DocumentCategory field if non-nil, zero value otherwise.

### GetDocumentCategoryOk

`func (o *UsDrivingLicenceBuilder) GetDocumentCategoryOk() (*string, bool)`

GetDocumentCategoryOk returns a tuple with the DocumentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCategory

`func (o *UsDrivingLicenceBuilder) SetDocumentCategory(v string)`

SetDocumentCategory sets DocumentCategory field to given value.

### HasDocumentCategory

`func (o *UsDrivingLicenceBuilder) HasDocumentCategory() bool`

HasDocumentCategory returns a boolean if a field has been set.

### GetExpirationDate

`func (o *UsDrivingLicenceBuilder) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *UsDrivingLicenceBuilder) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *UsDrivingLicenceBuilder) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *UsDrivingLicenceBuilder) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetEyeColorCode

`func (o *UsDrivingLicenceBuilder) GetEyeColorCode() string`

GetEyeColorCode returns the EyeColorCode field if non-nil, zero value otherwise.

### GetEyeColorCodeOk

`func (o *UsDrivingLicenceBuilder) GetEyeColorCodeOk() (*string, bool)`

GetEyeColorCodeOk returns a tuple with the EyeColorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEyeColorCode

`func (o *UsDrivingLicenceBuilder) SetEyeColorCode(v string)`

SetEyeColorCode sets EyeColorCode field to given value.

### HasEyeColorCode

`func (o *UsDrivingLicenceBuilder) HasEyeColorCode() bool`

HasEyeColorCode returns a boolean if a field has been set.

### GetFirstName

`func (o *UsDrivingLicenceBuilder) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UsDrivingLicenceBuilder) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UsDrivingLicenceBuilder) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UsDrivingLicenceBuilder) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGender

`func (o *UsDrivingLicenceBuilder) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UsDrivingLicenceBuilder) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UsDrivingLicenceBuilder) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UsDrivingLicenceBuilder) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetIssueDate

`func (o *UsDrivingLicenceBuilder) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *UsDrivingLicenceBuilder) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *UsDrivingLicenceBuilder) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *UsDrivingLicenceBuilder) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetLastName

`func (o *UsDrivingLicenceBuilder) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UsDrivingLicenceBuilder) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UsDrivingLicenceBuilder) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UsDrivingLicenceBuilder) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMiddleName

`func (o *UsDrivingLicenceBuilder) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UsDrivingLicenceBuilder) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UsDrivingLicenceBuilder) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UsDrivingLicenceBuilder) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetNameSuffix

`func (o *UsDrivingLicenceBuilder) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *UsDrivingLicenceBuilder) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *UsDrivingLicenceBuilder) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *UsDrivingLicenceBuilder) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### GetPostalCode

`func (o *UsDrivingLicenceBuilder) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UsDrivingLicenceBuilder) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UsDrivingLicenceBuilder) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UsDrivingLicenceBuilder) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetState

`func (o *UsDrivingLicenceBuilder) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UsDrivingLicenceBuilder) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UsDrivingLicenceBuilder) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UsDrivingLicenceBuilder) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWeightMeasure

`func (o *UsDrivingLicenceBuilder) GetWeightMeasure() int32`

GetWeightMeasure returns the WeightMeasure field if non-nil, zero value otherwise.

### GetWeightMeasureOk

`func (o *UsDrivingLicenceBuilder) GetWeightMeasureOk() (*int32, bool)`

GetWeightMeasureOk returns a tuple with the WeightMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightMeasure

`func (o *UsDrivingLicenceBuilder) SetWeightMeasure(v int32)`

SetWeightMeasure sets WeightMeasure field to given value.

### HasWeightMeasure

`func (o *UsDrivingLicenceBuilder) HasWeightMeasure() bool`

HasWeightMeasure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


