# DocumentWithDriverVerificationReportAllOfProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateOfBirth** | Pointer to **string** |  | [optional] 
**DateOfExpiry** | Pointer to **string** |  | [optional] 
**PersonalNumber** | Pointer to **string** |  | [optional] 
**DocumentNumbers** | Pointer to [**[]DocumentPropertiesDocumentNumbersInner**](DocumentPropertiesDocumentNumbersInner.md) |  | [optional] 
**DocumentType** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**IssuingCountry** | Pointer to **string** |  | [optional] 
**Nationality** | Pointer to **string** |  | [optional] 
**IssuingState** | Pointer to **string** |  | [optional] 
**IssuingDate** | Pointer to **string** |  | [optional] 
**Categorisation** | Pointer to **string** |  | [optional] 
**MrzLine1** | Pointer to **string** |  | [optional] 
**MrzLine2** | Pointer to **string** |  | [optional] 
**MrzLine3** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**PlaceOfBirth** | Pointer to **string** |  | [optional] 
**SpouseName** | Pointer to **string** |  | [optional] 
**WidowName** | Pointer to **string** |  | [optional] 
**AliasName** | Pointer to **string** |  | [optional] 
**IssuingAuthority** | Pointer to **string** |  | [optional] 
**Remarks** | Pointer to **string** |  | [optional] 
**CivilState** | Pointer to **string** |  | [optional] 
**Expatriation** | Pointer to **string** |  | [optional] 
**FatherName** | Pointer to **string** |  | [optional] 
**MotherName** | Pointer to **string** |  | [optional] 
**Religion** | Pointer to **string** |  | [optional] 
**TypeOfPermit** | Pointer to **string** |  | [optional] 
**VersionNumber** | Pointer to **string** |  | [optional] 
**DocumentSubtype** | Pointer to **string** |  | [optional] 
**Profession** | Pointer to **string** |  | [optional] 
**SecurityDocumentNumber** | Pointer to **string** |  | [optional] 
**TaxNumber** | Pointer to **string** |  | [optional] 
**NistIdentityEvidenceStrength** | Pointer to **string** |  | [optional] 
**HasIssuanceConfirmation** | Pointer to **string** |  | [optional] 
**RealIdCompliance** | Pointer to **bool** |  | [optional] 
**SecurityTier** | Pointer to **string** |  | [optional] 
**AddressLines** | Pointer to [**DocumentPropertiesAddressLines**](DocumentPropertiesAddressLines.md) |  | [optional] 
**Barcode** | Pointer to [**[]DocumentPropertiesBarcodeInner**](DocumentPropertiesBarcodeInner.md) |  | [optional] 
**Nfc** | Pointer to [**DocumentPropertiesNfc**](DocumentPropertiesNfc.md) |  | [optional] 
**DrivingLicenceInformation** | Pointer to [**DocumentPropertiesDrivingLicenceInformation**](DocumentPropertiesDrivingLicenceInformation.md) |  | [optional] 
**DocumentClassification** | Pointer to [**DocumentPropertiesDocumentClassification**](DocumentPropertiesDocumentClassification.md) |  | [optional] 
**ExtractedData** | Pointer to [**DocumentPropertiesExtractedData**](DocumentPropertiesExtractedData.md) |  | [optional] 
**DriversLicence** | Pointer to **bool** | True for **non-restricted** driving licences | [optional] 
**RestrictedLicence** | Pointer to **bool** | True for **limited/restricted** driving license, including learner&#39;s permits | [optional] 
**RawLicenceCategory** | Pointer to **string** | Underlying, non-normalised, licence category (e.g. \&quot;Junior operators license\&quot;) | [optional] 
**RawVehicleClasses** | Pointer to **string** | Comma-separated vehicle classes that the user is qualified for | [optional] 
**VehicleClassDetails** | Pointer to [**[]DocumentWithDriverVerificationReportAllOfPropertiesAllOfVehicleClassDetailsInner**](DocumentWithDriverVerificationReportAllOfPropertiesAllOfVehicleClassDetailsInner.md) | Detailed classes/categories information | [optional] 
**PassengerVehicle** | Pointer to [**DocumentWithDriverVerificationReportAllOfPropertiesAllOfPassengerVehicle**](DocumentWithDriverVerificationReportAllOfPropertiesAllOfPassengerVehicle.md) |  | [optional] 

## Methods

### NewDocumentWithDriverVerificationReportAllOfProperties

`func NewDocumentWithDriverVerificationReportAllOfProperties() *DocumentWithDriverVerificationReportAllOfProperties`

NewDocumentWithDriverVerificationReportAllOfProperties instantiates a new DocumentWithDriverVerificationReportAllOfProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDriverVerificationReportAllOfPropertiesWithDefaults

`func NewDocumentWithDriverVerificationReportAllOfPropertiesWithDefaults() *DocumentWithDriverVerificationReportAllOfProperties`

NewDocumentWithDriverVerificationReportAllOfPropertiesWithDefaults instantiates a new DocumentWithDriverVerificationReportAllOfProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOfBirth

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDateOfExpiry

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDateOfExpiry() string`

GetDateOfExpiry returns the DateOfExpiry field if non-nil, zero value otherwise.

### GetDateOfExpiryOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDateOfExpiryOk() (*string, bool)`

GetDateOfExpiryOk returns a tuple with the DateOfExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfExpiry

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetDateOfExpiry(v string)`

SetDateOfExpiry sets DateOfExpiry field to given value.

### HasDateOfExpiry

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasDateOfExpiry() bool`

HasDateOfExpiry returns a boolean if a field has been set.

### GetPersonalNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetPersonalNumber() string`

GetPersonalNumber returns the PersonalNumber field if non-nil, zero value otherwise.

### GetPersonalNumberOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetPersonalNumberOk() (*string, bool)`

GetPersonalNumberOk returns a tuple with the PersonalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetPersonalNumber(v string)`

SetPersonalNumber sets PersonalNumber field to given value.

### HasPersonalNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasPersonalNumber() bool`

HasPersonalNumber returns a boolean if a field has been set.

### GetDocumentNumbers

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDocumentNumbers() []DocumentPropertiesDocumentNumbersInner`

GetDocumentNumbers returns the DocumentNumbers field if non-nil, zero value otherwise.

### GetDocumentNumbersOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDocumentNumbersOk() (*[]DocumentPropertiesDocumentNumbersInner, bool)`

GetDocumentNumbersOk returns a tuple with the DocumentNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumbers

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetDocumentNumbers(v []DocumentPropertiesDocumentNumbersInner)`

SetDocumentNumbers sets DocumentNumbers field to given value.

### HasDocumentNumbers

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasDocumentNumbers() bool`

HasDocumentNumbers returns a boolean if a field has been set.

### GetDocumentType

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetFirstName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetGender

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetIssuingCountry

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### GetNationality

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetIssuingState

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetIssuingState() string`

GetIssuingState returns the IssuingState field if non-nil, zero value otherwise.

### GetIssuingStateOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetIssuingStateOk() (*string, bool)`

GetIssuingStateOk returns a tuple with the IssuingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingState

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetIssuingState(v string)`

SetIssuingState sets IssuingState field to given value.

### HasIssuingState

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasIssuingState() bool`

HasIssuingState returns a boolean if a field has been set.

### GetIssuingDate

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetIssuingDate() string`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetIssuingDateOk() (*string, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetIssuingDate(v string)`

SetIssuingDate sets IssuingDate field to given value.

### HasIssuingDate

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasIssuingDate() bool`

HasIssuingDate returns a boolean if a field has been set.

### GetCategorisation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetCategorisation() string`

GetCategorisation returns the Categorisation field if non-nil, zero value otherwise.

### GetCategorisationOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetCategorisationOk() (*string, bool)`

GetCategorisationOk returns a tuple with the Categorisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorisation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetCategorisation(v string)`

SetCategorisation sets Categorisation field to given value.

### HasCategorisation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasCategorisation() bool`

HasCategorisation returns a boolean if a field has been set.

### GetMrzLine1

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMrzLine1() string`

GetMrzLine1 returns the MrzLine1 field if non-nil, zero value otherwise.

### GetMrzLine1Ok

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMrzLine1Ok() (*string, bool)`

GetMrzLine1Ok returns a tuple with the MrzLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine1

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetMrzLine1(v string)`

SetMrzLine1 sets MrzLine1 field to given value.

### HasMrzLine1

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasMrzLine1() bool`

HasMrzLine1 returns a boolean if a field has been set.

### GetMrzLine2

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMrzLine2() string`

GetMrzLine2 returns the MrzLine2 field if non-nil, zero value otherwise.

### GetMrzLine2Ok

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMrzLine2Ok() (*string, bool)`

GetMrzLine2Ok returns a tuple with the MrzLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine2

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetMrzLine2(v string)`

SetMrzLine2 sets MrzLine2 field to given value.

### HasMrzLine2

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasMrzLine2() bool`

HasMrzLine2 returns a boolean if a field has been set.

### GetMrzLine3

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMrzLine3() string`

GetMrzLine3 returns the MrzLine3 field if non-nil, zero value otherwise.

### GetMrzLine3Ok

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMrzLine3Ok() (*string, bool)`

GetMrzLine3Ok returns a tuple with the MrzLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine3

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetMrzLine3(v string)`

SetMrzLine3 sets MrzLine3 field to given value.

### HasMrzLine3

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasMrzLine3() bool`

HasMrzLine3 returns a boolean if a field has been set.

### GetAddress

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetSpouseName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetSpouseName() string`

GetSpouseName returns the SpouseName field if non-nil, zero value otherwise.

### GetSpouseNameOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetSpouseNameOk() (*string, bool)`

GetSpouseNameOk returns a tuple with the SpouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetSpouseName(v string)`

SetSpouseName sets SpouseName field to given value.

### HasSpouseName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasSpouseName() bool`

HasSpouseName returns a boolean if a field has been set.

### GetWidowName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetWidowName() string`

GetWidowName returns the WidowName field if non-nil, zero value otherwise.

### GetWidowNameOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetWidowNameOk() (*string, bool)`

GetWidowNameOk returns a tuple with the WidowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidowName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetWidowName(v string)`

SetWidowName sets WidowName field to given value.

### HasWidowName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasWidowName() bool`

HasWidowName returns a boolean if a field has been set.

### GetAliasName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### GetRemarks

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetCivilState

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetCivilState() string`

GetCivilState returns the CivilState field if non-nil, zero value otherwise.

### GetCivilStateOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetCivilStateOk() (*string, bool)`

GetCivilStateOk returns a tuple with the CivilState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivilState

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetCivilState(v string)`

SetCivilState sets CivilState field to given value.

### HasCivilState

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasCivilState() bool`

HasCivilState returns a boolean if a field has been set.

### GetExpatriation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetExpatriation() string`

GetExpatriation returns the Expatriation field if non-nil, zero value otherwise.

### GetExpatriationOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetExpatriationOk() (*string, bool)`

GetExpatriationOk returns a tuple with the Expatriation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpatriation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetExpatriation(v string)`

SetExpatriation sets Expatriation field to given value.

### HasExpatriation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasExpatriation() bool`

HasExpatriation returns a boolean if a field has been set.

### GetFatherName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetFatherName() string`

GetFatherName returns the FatherName field if non-nil, zero value otherwise.

### GetFatherNameOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetFatherNameOk() (*string, bool)`

GetFatherNameOk returns a tuple with the FatherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatherName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetFatherName(v string)`

SetFatherName sets FatherName field to given value.

### HasFatherName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasFatherName() bool`

HasFatherName returns a boolean if a field has been set.

### GetMotherName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMotherName() string`

GetMotherName returns the MotherName field if non-nil, zero value otherwise.

### GetMotherNameOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetMotherNameOk() (*string, bool)`

GetMotherNameOk returns a tuple with the MotherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotherName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetMotherName(v string)`

SetMotherName sets MotherName field to given value.

### HasMotherName

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasMotherName() bool`

HasMotherName returns a boolean if a field has been set.

### GetReligion

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetReligion() string`

GetReligion returns the Religion field if non-nil, zero value otherwise.

### GetReligionOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetReligionOk() (*string, bool)`

GetReligionOk returns a tuple with the Religion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReligion

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetReligion(v string)`

SetReligion sets Religion field to given value.

### HasReligion

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasReligion() bool`

HasReligion returns a boolean if a field has been set.

### GetTypeOfPermit

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetTypeOfPermit() string`

GetTypeOfPermit returns the TypeOfPermit field if non-nil, zero value otherwise.

### GetTypeOfPermitOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetTypeOfPermitOk() (*string, bool)`

GetTypeOfPermitOk returns a tuple with the TypeOfPermit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfPermit

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetTypeOfPermit(v string)`

SetTypeOfPermit sets TypeOfPermit field to given value.

### HasTypeOfPermit

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasTypeOfPermit() bool`

HasTypeOfPermit returns a boolean if a field has been set.

### GetVersionNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetDocumentSubtype

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDocumentSubtype() string`

GetDocumentSubtype returns the DocumentSubtype field if non-nil, zero value otherwise.

### GetDocumentSubtypeOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDocumentSubtypeOk() (*string, bool)`

GetDocumentSubtypeOk returns a tuple with the DocumentSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSubtype

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetDocumentSubtype(v string)`

SetDocumentSubtype sets DocumentSubtype field to given value.

### HasDocumentSubtype

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasDocumentSubtype() bool`

HasDocumentSubtype returns a boolean if a field has been set.

### GetProfession

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetSecurityDocumentNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetSecurityDocumentNumber() string`

GetSecurityDocumentNumber returns the SecurityDocumentNumber field if non-nil, zero value otherwise.

### GetSecurityDocumentNumberOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetSecurityDocumentNumberOk() (*string, bool)`

GetSecurityDocumentNumberOk returns a tuple with the SecurityDocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityDocumentNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetSecurityDocumentNumber(v string)`

SetSecurityDocumentNumber sets SecurityDocumentNumber field to given value.

### HasSecurityDocumentNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasSecurityDocumentNumber() bool`

HasSecurityDocumentNumber returns a boolean if a field has been set.

### GetTaxNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumber

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasTaxNumber() bool`

HasTaxNumber returns a boolean if a field has been set.

### GetNistIdentityEvidenceStrength

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetNistIdentityEvidenceStrength() string`

GetNistIdentityEvidenceStrength returns the NistIdentityEvidenceStrength field if non-nil, zero value otherwise.

### GetNistIdentityEvidenceStrengthOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetNistIdentityEvidenceStrengthOk() (*string, bool)`

GetNistIdentityEvidenceStrengthOk returns a tuple with the NistIdentityEvidenceStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNistIdentityEvidenceStrength

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetNistIdentityEvidenceStrength(v string)`

SetNistIdentityEvidenceStrength sets NistIdentityEvidenceStrength field to given value.

### HasNistIdentityEvidenceStrength

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasNistIdentityEvidenceStrength() bool`

HasNistIdentityEvidenceStrength returns a boolean if a field has been set.

### GetHasIssuanceConfirmation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetHasIssuanceConfirmation() string`

GetHasIssuanceConfirmation returns the HasIssuanceConfirmation field if non-nil, zero value otherwise.

### GetHasIssuanceConfirmationOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetHasIssuanceConfirmationOk() (*string, bool)`

GetHasIssuanceConfirmationOk returns a tuple with the HasIssuanceConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssuanceConfirmation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetHasIssuanceConfirmation(v string)`

SetHasIssuanceConfirmation sets HasIssuanceConfirmation field to given value.

### HasHasIssuanceConfirmation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasHasIssuanceConfirmation() bool`

HasHasIssuanceConfirmation returns a boolean if a field has been set.

### GetRealIdCompliance

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRealIdCompliance() bool`

GetRealIdCompliance returns the RealIdCompliance field if non-nil, zero value otherwise.

### GetRealIdComplianceOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRealIdComplianceOk() (*bool, bool)`

GetRealIdComplianceOk returns a tuple with the RealIdCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealIdCompliance

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetRealIdCompliance(v bool)`

SetRealIdCompliance sets RealIdCompliance field to given value.

### HasRealIdCompliance

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasRealIdCompliance() bool`

HasRealIdCompliance returns a boolean if a field has been set.

### GetSecurityTier

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetSecurityTier() string`

GetSecurityTier returns the SecurityTier field if non-nil, zero value otherwise.

### GetSecurityTierOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetSecurityTierOk() (*string, bool)`

GetSecurityTierOk returns a tuple with the SecurityTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityTier

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetSecurityTier(v string)`

SetSecurityTier sets SecurityTier field to given value.

### HasSecurityTier

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasSecurityTier() bool`

HasSecurityTier returns a boolean if a field has been set.

### GetAddressLines

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetAddressLines() DocumentPropertiesAddressLines`

GetAddressLines returns the AddressLines field if non-nil, zero value otherwise.

### GetAddressLinesOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetAddressLinesOk() (*DocumentPropertiesAddressLines, bool)`

GetAddressLinesOk returns a tuple with the AddressLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLines

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetAddressLines(v DocumentPropertiesAddressLines)`

SetAddressLines sets AddressLines field to given value.

### HasAddressLines

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasAddressLines() bool`

HasAddressLines returns a boolean if a field has been set.

### GetBarcode

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetBarcode() []DocumentPropertiesBarcodeInner`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetBarcodeOk() (*[]DocumentPropertiesBarcodeInner, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetBarcode(v []DocumentPropertiesBarcodeInner)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetNfc

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetNfc() DocumentPropertiesNfc`

GetNfc returns the Nfc field if non-nil, zero value otherwise.

### GetNfcOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetNfcOk() (*DocumentPropertiesNfc, bool)`

GetNfcOk returns a tuple with the Nfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfc

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetNfc(v DocumentPropertiesNfc)`

SetNfc sets Nfc field to given value.

### HasNfc

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasNfc() bool`

HasNfc returns a boolean if a field has been set.

### GetDrivingLicenceInformation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDrivingLicenceInformation() DocumentPropertiesDrivingLicenceInformation`

GetDrivingLicenceInformation returns the DrivingLicenceInformation field if non-nil, zero value otherwise.

### GetDrivingLicenceInformationOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDrivingLicenceInformationOk() (*DocumentPropertiesDrivingLicenceInformation, bool)`

GetDrivingLicenceInformationOk returns a tuple with the DrivingLicenceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingLicenceInformation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetDrivingLicenceInformation(v DocumentPropertiesDrivingLicenceInformation)`

SetDrivingLicenceInformation sets DrivingLicenceInformation field to given value.

### HasDrivingLicenceInformation

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasDrivingLicenceInformation() bool`

HasDrivingLicenceInformation returns a boolean if a field has been set.

### GetDocumentClassification

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDocumentClassification() DocumentPropertiesDocumentClassification`

GetDocumentClassification returns the DocumentClassification field if non-nil, zero value otherwise.

### GetDocumentClassificationOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDocumentClassificationOk() (*DocumentPropertiesDocumentClassification, bool)`

GetDocumentClassificationOk returns a tuple with the DocumentClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentClassification

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetDocumentClassification(v DocumentPropertiesDocumentClassification)`

SetDocumentClassification sets DocumentClassification field to given value.

### HasDocumentClassification

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasDocumentClassification() bool`

HasDocumentClassification returns a boolean if a field has been set.

### GetExtractedData

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetExtractedData() DocumentPropertiesExtractedData`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetExtractedDataOk() (*DocumentPropertiesExtractedData, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetExtractedData(v DocumentPropertiesExtractedData)`

SetExtractedData sets ExtractedData field to given value.

### HasExtractedData

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasExtractedData() bool`

HasExtractedData returns a boolean if a field has been set.

### GetDriversLicence

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDriversLicence() bool`

GetDriversLicence returns the DriversLicence field if non-nil, zero value otherwise.

### GetDriversLicenceOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetDriversLicenceOk() (*bool, bool)`

GetDriversLicenceOk returns a tuple with the DriversLicence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriversLicence

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetDriversLicence(v bool)`

SetDriversLicence sets DriversLicence field to given value.

### HasDriversLicence

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasDriversLicence() bool`

HasDriversLicence returns a boolean if a field has been set.

### GetRestrictedLicence

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRestrictedLicence() bool`

GetRestrictedLicence returns the RestrictedLicence field if non-nil, zero value otherwise.

### GetRestrictedLicenceOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRestrictedLicenceOk() (*bool, bool)`

GetRestrictedLicenceOk returns a tuple with the RestrictedLicence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedLicence

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetRestrictedLicence(v bool)`

SetRestrictedLicence sets RestrictedLicence field to given value.

### HasRestrictedLicence

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasRestrictedLicence() bool`

HasRestrictedLicence returns a boolean if a field has been set.

### GetRawLicenceCategory

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRawLicenceCategory() string`

GetRawLicenceCategory returns the RawLicenceCategory field if non-nil, zero value otherwise.

### GetRawLicenceCategoryOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRawLicenceCategoryOk() (*string, bool)`

GetRawLicenceCategoryOk returns a tuple with the RawLicenceCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawLicenceCategory

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetRawLicenceCategory(v string)`

SetRawLicenceCategory sets RawLicenceCategory field to given value.

### HasRawLicenceCategory

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasRawLicenceCategory() bool`

HasRawLicenceCategory returns a boolean if a field has been set.

### GetRawVehicleClasses

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRawVehicleClasses() string`

GetRawVehicleClasses returns the RawVehicleClasses field if non-nil, zero value otherwise.

### GetRawVehicleClassesOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetRawVehicleClassesOk() (*string, bool)`

GetRawVehicleClassesOk returns a tuple with the RawVehicleClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawVehicleClasses

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetRawVehicleClasses(v string)`

SetRawVehicleClasses sets RawVehicleClasses field to given value.

### HasRawVehicleClasses

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasRawVehicleClasses() bool`

HasRawVehicleClasses returns a boolean if a field has been set.

### GetVehicleClassDetails

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetVehicleClassDetails() []DocumentWithDriverVerificationReportAllOfPropertiesAllOfVehicleClassDetailsInner`

GetVehicleClassDetails returns the VehicleClassDetails field if non-nil, zero value otherwise.

### GetVehicleClassDetailsOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetVehicleClassDetailsOk() (*[]DocumentWithDriverVerificationReportAllOfPropertiesAllOfVehicleClassDetailsInner, bool)`

GetVehicleClassDetailsOk returns a tuple with the VehicleClassDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleClassDetails

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetVehicleClassDetails(v []DocumentWithDriverVerificationReportAllOfPropertiesAllOfVehicleClassDetailsInner)`

SetVehicleClassDetails sets VehicleClassDetails field to given value.

### HasVehicleClassDetails

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasVehicleClassDetails() bool`

HasVehicleClassDetails returns a boolean if a field has been set.

### GetPassengerVehicle

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetPassengerVehicle() DocumentWithDriverVerificationReportAllOfPropertiesAllOfPassengerVehicle`

GetPassengerVehicle returns the PassengerVehicle field if non-nil, zero value otherwise.

### GetPassengerVehicleOk

`func (o *DocumentWithDriverVerificationReportAllOfProperties) GetPassengerVehicleOk() (*DocumentWithDriverVerificationReportAllOfPropertiesAllOfPassengerVehicle, bool)`

GetPassengerVehicleOk returns a tuple with the PassengerVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassengerVehicle

`func (o *DocumentWithDriverVerificationReportAllOfProperties) SetPassengerVehicle(v DocumentWithDriverVerificationReportAllOfPropertiesAllOfPassengerVehicle)`

SetPassengerVehicle sets PassengerVehicle field to given value.

### HasPassengerVehicle

`func (o *DocumentWithDriverVerificationReportAllOfProperties) HasPassengerVehicle() bool`

HasPassengerVehicle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


