# ExtractionExtractedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentNumber** | Pointer to **string** | The official document number. | [optional] 
**FirstName** | Pointer to **string** | First name. | [optional] 
**LastName** | Pointer to **string** | Last name. | [optional] 
**FullName** | Pointer to **string** | Full name. | [optional] 
**SpouseName** | Pointer to **string** | Spouse name (French documents only). | [optional] 
**WidowName** | Pointer to **string** | Widow name (French documents only). | [optional] 
**AliasName** | Pointer to **string** | Alias name (French documents only). | [optional] 
**Gender** | Pointer to **string** | Gender (Valid values are Male and Female). | [optional] 
**DateOfBirth** | Pointer to **string** | Date of birth in YYYY-MM-DD format. | [optional] 
**DateOfExpiry** | Pointer to **string** | Date of expiry in YYYY-MM-DD format. | [optional] 
**ExpiryDate** | Pointer to **string** | Date of expiry in YYYY-MM-DD format. | [optional] 
**Nationality** | Pointer to **string** | Nationality in 3-letter ISO code. | [optional] 
**MrzLine1** | Pointer to **string** | Line 1 of the MRZ code. | [optional] 
**MrzLine2** | Pointer to **string** | Line 2 of the MRZ code. | [optional] 
**MrzLine3** | Pointer to **string** | Line 3 of the MRZ code. | [optional] 
**Address1** | Pointer to **string** | Line 1 of the address. | [optional] 
**Address2** | Pointer to **string** | Line 2 of the address. | [optional] 
**Address3** | Pointer to **string** | Line 3 of the address. | [optional] 
**Address4** | Pointer to **string** | Line 4 of the address. | [optional] 
**Address5** | Pointer to **string** | Line 5 of the address. | [optional] 
**IssuingAuthority** | Pointer to **string** | Issuing authority. | [optional] 
**IssuingCountry** | Pointer to [**CountryCodes**](CountryCodes.md) | Document country in 3-letter ISO code. | [optional] 
**DocumentType** | Pointer to [**DocumentTypes**](DocumentTypes.md) | Type of document. | [optional] 
**PlaceOfBirth** | Pointer to **string** | Place of birth. | [optional] 
**IssuingState** | Pointer to **string** | The state that issued the document. | [optional] 
**IssuingDate** | Pointer to **string** | Issuing date in YYYY-MM-DD format. | [optional] 
**PersonalNumber** | Pointer to **string** | The owner&#39;s unique identification number. | [optional] 

## Methods

### NewExtractionExtractedData

`func NewExtractionExtractedData() *ExtractionExtractedData`

NewExtractionExtractedData instantiates a new ExtractionExtractedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionExtractedDataWithDefaults

`func NewExtractionExtractedDataWithDefaults() *ExtractionExtractedData`

NewExtractionExtractedDataWithDefaults instantiates a new ExtractionExtractedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentNumber

`func (o *ExtractionExtractedData) GetDocumentNumber() string`

GetDocumentNumber returns the DocumentNumber field if non-nil, zero value otherwise.

### GetDocumentNumberOk

`func (o *ExtractionExtractedData) GetDocumentNumberOk() (*string, bool)`

GetDocumentNumberOk returns a tuple with the DocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumber

`func (o *ExtractionExtractedData) SetDocumentNumber(v string)`

SetDocumentNumber sets DocumentNumber field to given value.

### HasDocumentNumber

`func (o *ExtractionExtractedData) HasDocumentNumber() bool`

HasDocumentNumber returns a boolean if a field has been set.

### GetFirstName

`func (o *ExtractionExtractedData) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ExtractionExtractedData) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ExtractionExtractedData) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ExtractionExtractedData) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ExtractionExtractedData) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ExtractionExtractedData) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ExtractionExtractedData) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ExtractionExtractedData) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFullName

`func (o *ExtractionExtractedData) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ExtractionExtractedData) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ExtractionExtractedData) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *ExtractionExtractedData) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetSpouseName

`func (o *ExtractionExtractedData) GetSpouseName() string`

GetSpouseName returns the SpouseName field if non-nil, zero value otherwise.

### GetSpouseNameOk

`func (o *ExtractionExtractedData) GetSpouseNameOk() (*string, bool)`

GetSpouseNameOk returns a tuple with the SpouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseName

`func (o *ExtractionExtractedData) SetSpouseName(v string)`

SetSpouseName sets SpouseName field to given value.

### HasSpouseName

`func (o *ExtractionExtractedData) HasSpouseName() bool`

HasSpouseName returns a boolean if a field has been set.

### GetWidowName

`func (o *ExtractionExtractedData) GetWidowName() string`

GetWidowName returns the WidowName field if non-nil, zero value otherwise.

### GetWidowNameOk

`func (o *ExtractionExtractedData) GetWidowNameOk() (*string, bool)`

GetWidowNameOk returns a tuple with the WidowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidowName

`func (o *ExtractionExtractedData) SetWidowName(v string)`

SetWidowName sets WidowName field to given value.

### HasWidowName

`func (o *ExtractionExtractedData) HasWidowName() bool`

HasWidowName returns a boolean if a field has been set.

### GetAliasName

`func (o *ExtractionExtractedData) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *ExtractionExtractedData) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *ExtractionExtractedData) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *ExtractionExtractedData) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetGender

`func (o *ExtractionExtractedData) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ExtractionExtractedData) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ExtractionExtractedData) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ExtractionExtractedData) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *ExtractionExtractedData) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *ExtractionExtractedData) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *ExtractionExtractedData) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *ExtractionExtractedData) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDateOfExpiry

`func (o *ExtractionExtractedData) GetDateOfExpiry() string`

GetDateOfExpiry returns the DateOfExpiry field if non-nil, zero value otherwise.

### GetDateOfExpiryOk

`func (o *ExtractionExtractedData) GetDateOfExpiryOk() (*string, bool)`

GetDateOfExpiryOk returns a tuple with the DateOfExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfExpiry

`func (o *ExtractionExtractedData) SetDateOfExpiry(v string)`

SetDateOfExpiry sets DateOfExpiry field to given value.

### HasDateOfExpiry

`func (o *ExtractionExtractedData) HasDateOfExpiry() bool`

HasDateOfExpiry returns a boolean if a field has been set.

### GetExpiryDate

`func (o *ExtractionExtractedData) GetExpiryDate() string`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *ExtractionExtractedData) GetExpiryDateOk() (*string, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *ExtractionExtractedData) SetExpiryDate(v string)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *ExtractionExtractedData) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetNationality

`func (o *ExtractionExtractedData) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *ExtractionExtractedData) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *ExtractionExtractedData) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *ExtractionExtractedData) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetMrzLine1

`func (o *ExtractionExtractedData) GetMrzLine1() string`

GetMrzLine1 returns the MrzLine1 field if non-nil, zero value otherwise.

### GetMrzLine1Ok

`func (o *ExtractionExtractedData) GetMrzLine1Ok() (*string, bool)`

GetMrzLine1Ok returns a tuple with the MrzLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine1

`func (o *ExtractionExtractedData) SetMrzLine1(v string)`

SetMrzLine1 sets MrzLine1 field to given value.

### HasMrzLine1

`func (o *ExtractionExtractedData) HasMrzLine1() bool`

HasMrzLine1 returns a boolean if a field has been set.

### GetMrzLine2

`func (o *ExtractionExtractedData) GetMrzLine2() string`

GetMrzLine2 returns the MrzLine2 field if non-nil, zero value otherwise.

### GetMrzLine2Ok

`func (o *ExtractionExtractedData) GetMrzLine2Ok() (*string, bool)`

GetMrzLine2Ok returns a tuple with the MrzLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine2

`func (o *ExtractionExtractedData) SetMrzLine2(v string)`

SetMrzLine2 sets MrzLine2 field to given value.

### HasMrzLine2

`func (o *ExtractionExtractedData) HasMrzLine2() bool`

HasMrzLine2 returns a boolean if a field has been set.

### GetMrzLine3

`func (o *ExtractionExtractedData) GetMrzLine3() string`

GetMrzLine3 returns the MrzLine3 field if non-nil, zero value otherwise.

### GetMrzLine3Ok

`func (o *ExtractionExtractedData) GetMrzLine3Ok() (*string, bool)`

GetMrzLine3Ok returns a tuple with the MrzLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine3

`func (o *ExtractionExtractedData) SetMrzLine3(v string)`

SetMrzLine3 sets MrzLine3 field to given value.

### HasMrzLine3

`func (o *ExtractionExtractedData) HasMrzLine3() bool`

HasMrzLine3 returns a boolean if a field has been set.

### GetAddress1

`func (o *ExtractionExtractedData) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ExtractionExtractedData) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ExtractionExtractedData) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *ExtractionExtractedData) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *ExtractionExtractedData) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ExtractionExtractedData) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ExtractionExtractedData) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ExtractionExtractedData) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAddress3

`func (o *ExtractionExtractedData) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *ExtractionExtractedData) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *ExtractionExtractedData) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *ExtractionExtractedData) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### GetAddress4

`func (o *ExtractionExtractedData) GetAddress4() string`

GetAddress4 returns the Address4 field if non-nil, zero value otherwise.

### GetAddress4Ok

`func (o *ExtractionExtractedData) GetAddress4Ok() (*string, bool)`

GetAddress4Ok returns a tuple with the Address4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress4

`func (o *ExtractionExtractedData) SetAddress4(v string)`

SetAddress4 sets Address4 field to given value.

### HasAddress4

`func (o *ExtractionExtractedData) HasAddress4() bool`

HasAddress4 returns a boolean if a field has been set.

### GetAddress5

`func (o *ExtractionExtractedData) GetAddress5() string`

GetAddress5 returns the Address5 field if non-nil, zero value otherwise.

### GetAddress5Ok

`func (o *ExtractionExtractedData) GetAddress5Ok() (*string, bool)`

GetAddress5Ok returns a tuple with the Address5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress5

`func (o *ExtractionExtractedData) SetAddress5(v string)`

SetAddress5 sets Address5 field to given value.

### HasAddress5

`func (o *ExtractionExtractedData) HasAddress5() bool`

HasAddress5 returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *ExtractionExtractedData) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *ExtractionExtractedData) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *ExtractionExtractedData) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *ExtractionExtractedData) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### GetIssuingCountry

`func (o *ExtractionExtractedData) GetIssuingCountry() CountryCodes`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *ExtractionExtractedData) GetIssuingCountryOk() (*CountryCodes, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *ExtractionExtractedData) SetIssuingCountry(v CountryCodes)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *ExtractionExtractedData) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### GetDocumentType

`func (o *ExtractionExtractedData) GetDocumentType() DocumentTypes`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ExtractionExtractedData) GetDocumentTypeOk() (*DocumentTypes, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ExtractionExtractedData) SetDocumentType(v DocumentTypes)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *ExtractionExtractedData) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *ExtractionExtractedData) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *ExtractionExtractedData) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *ExtractionExtractedData) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *ExtractionExtractedData) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetIssuingState

`func (o *ExtractionExtractedData) GetIssuingState() string`

GetIssuingState returns the IssuingState field if non-nil, zero value otherwise.

### GetIssuingStateOk

`func (o *ExtractionExtractedData) GetIssuingStateOk() (*string, bool)`

GetIssuingStateOk returns a tuple with the IssuingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingState

`func (o *ExtractionExtractedData) SetIssuingState(v string)`

SetIssuingState sets IssuingState field to given value.

### HasIssuingState

`func (o *ExtractionExtractedData) HasIssuingState() bool`

HasIssuingState returns a boolean if a field has been set.

### GetIssuingDate

`func (o *ExtractionExtractedData) GetIssuingDate() string`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *ExtractionExtractedData) GetIssuingDateOk() (*string, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *ExtractionExtractedData) SetIssuingDate(v string)`

SetIssuingDate sets IssuingDate field to given value.

### HasIssuingDate

`func (o *ExtractionExtractedData) HasIssuingDate() bool`

HasIssuingDate returns a boolean if a field has been set.

### GetPersonalNumber

`func (o *ExtractionExtractedData) GetPersonalNumber() string`

GetPersonalNumber returns the PersonalNumber field if non-nil, zero value otherwise.

### GetPersonalNumberOk

`func (o *ExtractionExtractedData) GetPersonalNumberOk() (*string, bool)`

GetPersonalNumberOk returns a tuple with the PersonalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumber

`func (o *ExtractionExtractedData) SetPersonalNumber(v string)`

SetPersonalNumber sets PersonalNumber field to given value.

### HasPersonalNumber

`func (o *ExtractionExtractedData) HasPersonalNumber() bool`

HasPersonalNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


