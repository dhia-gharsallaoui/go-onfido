# DocumentProperties

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

## Methods

### NewDocumentProperties

`func NewDocumentProperties() *DocumentProperties`

NewDocumentProperties instantiates a new DocumentProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentPropertiesWithDefaults

`func NewDocumentPropertiesWithDefaults() *DocumentProperties`

NewDocumentPropertiesWithDefaults instantiates a new DocumentProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOfBirth

`func (o *DocumentProperties) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *DocumentProperties) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *DocumentProperties) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *DocumentProperties) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDateOfExpiry

`func (o *DocumentProperties) GetDateOfExpiry() string`

GetDateOfExpiry returns the DateOfExpiry field if non-nil, zero value otherwise.

### GetDateOfExpiryOk

`func (o *DocumentProperties) GetDateOfExpiryOk() (*string, bool)`

GetDateOfExpiryOk returns a tuple with the DateOfExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfExpiry

`func (o *DocumentProperties) SetDateOfExpiry(v string)`

SetDateOfExpiry sets DateOfExpiry field to given value.

### HasDateOfExpiry

`func (o *DocumentProperties) HasDateOfExpiry() bool`

HasDateOfExpiry returns a boolean if a field has been set.

### GetPersonalNumber

`func (o *DocumentProperties) GetPersonalNumber() string`

GetPersonalNumber returns the PersonalNumber field if non-nil, zero value otherwise.

### GetPersonalNumberOk

`func (o *DocumentProperties) GetPersonalNumberOk() (*string, bool)`

GetPersonalNumberOk returns a tuple with the PersonalNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalNumber

`func (o *DocumentProperties) SetPersonalNumber(v string)`

SetPersonalNumber sets PersonalNumber field to given value.

### HasPersonalNumber

`func (o *DocumentProperties) HasPersonalNumber() bool`

HasPersonalNumber returns a boolean if a field has been set.

### GetDocumentNumbers

`func (o *DocumentProperties) GetDocumentNumbers() []DocumentPropertiesDocumentNumbersInner`

GetDocumentNumbers returns the DocumentNumbers field if non-nil, zero value otherwise.

### GetDocumentNumbersOk

`func (o *DocumentProperties) GetDocumentNumbersOk() (*[]DocumentPropertiesDocumentNumbersInner, bool)`

GetDocumentNumbersOk returns a tuple with the DocumentNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentNumbers

`func (o *DocumentProperties) SetDocumentNumbers(v []DocumentPropertiesDocumentNumbersInner)`

SetDocumentNumbers sets DocumentNumbers field to given value.

### HasDocumentNumbers

`func (o *DocumentProperties) HasDocumentNumbers() bool`

HasDocumentNumbers returns a boolean if a field has been set.

### GetDocumentType

`func (o *DocumentProperties) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *DocumentProperties) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *DocumentProperties) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *DocumentProperties) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetFirstName

`func (o *DocumentProperties) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DocumentProperties) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DocumentProperties) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DocumentProperties) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleName

`func (o *DocumentProperties) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *DocumentProperties) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *DocumentProperties) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *DocumentProperties) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetLastName

`func (o *DocumentProperties) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DocumentProperties) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DocumentProperties) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DocumentProperties) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetGender

`func (o *DocumentProperties) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *DocumentProperties) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *DocumentProperties) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *DocumentProperties) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetIssuingCountry

`func (o *DocumentProperties) GetIssuingCountry() string`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *DocumentProperties) GetIssuingCountryOk() (*string, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *DocumentProperties) SetIssuingCountry(v string)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *DocumentProperties) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### GetNationality

`func (o *DocumentProperties) GetNationality() string`

GetNationality returns the Nationality field if non-nil, zero value otherwise.

### GetNationalityOk

`func (o *DocumentProperties) GetNationalityOk() (*string, bool)`

GetNationalityOk returns a tuple with the Nationality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationality

`func (o *DocumentProperties) SetNationality(v string)`

SetNationality sets Nationality field to given value.

### HasNationality

`func (o *DocumentProperties) HasNationality() bool`

HasNationality returns a boolean if a field has been set.

### GetIssuingState

`func (o *DocumentProperties) GetIssuingState() string`

GetIssuingState returns the IssuingState field if non-nil, zero value otherwise.

### GetIssuingStateOk

`func (o *DocumentProperties) GetIssuingStateOk() (*string, bool)`

GetIssuingStateOk returns a tuple with the IssuingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingState

`func (o *DocumentProperties) SetIssuingState(v string)`

SetIssuingState sets IssuingState field to given value.

### HasIssuingState

`func (o *DocumentProperties) HasIssuingState() bool`

HasIssuingState returns a boolean if a field has been set.

### GetIssuingDate

`func (o *DocumentProperties) GetIssuingDate() string`

GetIssuingDate returns the IssuingDate field if non-nil, zero value otherwise.

### GetIssuingDateOk

`func (o *DocumentProperties) GetIssuingDateOk() (*string, bool)`

GetIssuingDateOk returns a tuple with the IssuingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingDate

`func (o *DocumentProperties) SetIssuingDate(v string)`

SetIssuingDate sets IssuingDate field to given value.

### HasIssuingDate

`func (o *DocumentProperties) HasIssuingDate() bool`

HasIssuingDate returns a boolean if a field has been set.

### GetCategorisation

`func (o *DocumentProperties) GetCategorisation() string`

GetCategorisation returns the Categorisation field if non-nil, zero value otherwise.

### GetCategorisationOk

`func (o *DocumentProperties) GetCategorisationOk() (*string, bool)`

GetCategorisationOk returns a tuple with the Categorisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorisation

`func (o *DocumentProperties) SetCategorisation(v string)`

SetCategorisation sets Categorisation field to given value.

### HasCategorisation

`func (o *DocumentProperties) HasCategorisation() bool`

HasCategorisation returns a boolean if a field has been set.

### GetMrzLine1

`func (o *DocumentProperties) GetMrzLine1() string`

GetMrzLine1 returns the MrzLine1 field if non-nil, zero value otherwise.

### GetMrzLine1Ok

`func (o *DocumentProperties) GetMrzLine1Ok() (*string, bool)`

GetMrzLine1Ok returns a tuple with the MrzLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine1

`func (o *DocumentProperties) SetMrzLine1(v string)`

SetMrzLine1 sets MrzLine1 field to given value.

### HasMrzLine1

`func (o *DocumentProperties) HasMrzLine1() bool`

HasMrzLine1 returns a boolean if a field has been set.

### GetMrzLine2

`func (o *DocumentProperties) GetMrzLine2() string`

GetMrzLine2 returns the MrzLine2 field if non-nil, zero value otherwise.

### GetMrzLine2Ok

`func (o *DocumentProperties) GetMrzLine2Ok() (*string, bool)`

GetMrzLine2Ok returns a tuple with the MrzLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine2

`func (o *DocumentProperties) SetMrzLine2(v string)`

SetMrzLine2 sets MrzLine2 field to given value.

### HasMrzLine2

`func (o *DocumentProperties) HasMrzLine2() bool`

HasMrzLine2 returns a boolean if a field has been set.

### GetMrzLine3

`func (o *DocumentProperties) GetMrzLine3() string`

GetMrzLine3 returns the MrzLine3 field if non-nil, zero value otherwise.

### GetMrzLine3Ok

`func (o *DocumentProperties) GetMrzLine3Ok() (*string, bool)`

GetMrzLine3Ok returns a tuple with the MrzLine3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrzLine3

`func (o *DocumentProperties) SetMrzLine3(v string)`

SetMrzLine3 sets MrzLine3 field to given value.

### HasMrzLine3

`func (o *DocumentProperties) HasMrzLine3() bool`

HasMrzLine3 returns a boolean if a field has been set.

### GetAddress

`func (o *DocumentProperties) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DocumentProperties) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DocumentProperties) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DocumentProperties) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPlaceOfBirth

`func (o *DocumentProperties) GetPlaceOfBirth() string`

GetPlaceOfBirth returns the PlaceOfBirth field if non-nil, zero value otherwise.

### GetPlaceOfBirthOk

`func (o *DocumentProperties) GetPlaceOfBirthOk() (*string, bool)`

GetPlaceOfBirthOk returns a tuple with the PlaceOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceOfBirth

`func (o *DocumentProperties) SetPlaceOfBirth(v string)`

SetPlaceOfBirth sets PlaceOfBirth field to given value.

### HasPlaceOfBirth

`func (o *DocumentProperties) HasPlaceOfBirth() bool`

HasPlaceOfBirth returns a boolean if a field has been set.

### GetSpouseName

`func (o *DocumentProperties) GetSpouseName() string`

GetSpouseName returns the SpouseName field if non-nil, zero value otherwise.

### GetSpouseNameOk

`func (o *DocumentProperties) GetSpouseNameOk() (*string, bool)`

GetSpouseNameOk returns a tuple with the SpouseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseName

`func (o *DocumentProperties) SetSpouseName(v string)`

SetSpouseName sets SpouseName field to given value.

### HasSpouseName

`func (o *DocumentProperties) HasSpouseName() bool`

HasSpouseName returns a boolean if a field has been set.

### GetWidowName

`func (o *DocumentProperties) GetWidowName() string`

GetWidowName returns the WidowName field if non-nil, zero value otherwise.

### GetWidowNameOk

`func (o *DocumentProperties) GetWidowNameOk() (*string, bool)`

GetWidowNameOk returns a tuple with the WidowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidowName

`func (o *DocumentProperties) SetWidowName(v string)`

SetWidowName sets WidowName field to given value.

### HasWidowName

`func (o *DocumentProperties) HasWidowName() bool`

HasWidowName returns a boolean if a field has been set.

### GetAliasName

`func (o *DocumentProperties) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *DocumentProperties) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *DocumentProperties) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *DocumentProperties) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetIssuingAuthority

`func (o *DocumentProperties) GetIssuingAuthority() string`

GetIssuingAuthority returns the IssuingAuthority field if non-nil, zero value otherwise.

### GetIssuingAuthorityOk

`func (o *DocumentProperties) GetIssuingAuthorityOk() (*string, bool)`

GetIssuingAuthorityOk returns a tuple with the IssuingAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingAuthority

`func (o *DocumentProperties) SetIssuingAuthority(v string)`

SetIssuingAuthority sets IssuingAuthority field to given value.

### HasIssuingAuthority

`func (o *DocumentProperties) HasIssuingAuthority() bool`

HasIssuingAuthority returns a boolean if a field has been set.

### GetRemarks

`func (o *DocumentProperties) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *DocumentProperties) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *DocumentProperties) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *DocumentProperties) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetCivilState

`func (o *DocumentProperties) GetCivilState() string`

GetCivilState returns the CivilState field if non-nil, zero value otherwise.

### GetCivilStateOk

`func (o *DocumentProperties) GetCivilStateOk() (*string, bool)`

GetCivilStateOk returns a tuple with the CivilState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivilState

`func (o *DocumentProperties) SetCivilState(v string)`

SetCivilState sets CivilState field to given value.

### HasCivilState

`func (o *DocumentProperties) HasCivilState() bool`

HasCivilState returns a boolean if a field has been set.

### GetExpatriation

`func (o *DocumentProperties) GetExpatriation() string`

GetExpatriation returns the Expatriation field if non-nil, zero value otherwise.

### GetExpatriationOk

`func (o *DocumentProperties) GetExpatriationOk() (*string, bool)`

GetExpatriationOk returns a tuple with the Expatriation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpatriation

`func (o *DocumentProperties) SetExpatriation(v string)`

SetExpatriation sets Expatriation field to given value.

### HasExpatriation

`func (o *DocumentProperties) HasExpatriation() bool`

HasExpatriation returns a boolean if a field has been set.

### GetFatherName

`func (o *DocumentProperties) GetFatherName() string`

GetFatherName returns the FatherName field if non-nil, zero value otherwise.

### GetFatherNameOk

`func (o *DocumentProperties) GetFatherNameOk() (*string, bool)`

GetFatherNameOk returns a tuple with the FatherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatherName

`func (o *DocumentProperties) SetFatherName(v string)`

SetFatherName sets FatherName field to given value.

### HasFatherName

`func (o *DocumentProperties) HasFatherName() bool`

HasFatherName returns a boolean if a field has been set.

### GetMotherName

`func (o *DocumentProperties) GetMotherName() string`

GetMotherName returns the MotherName field if non-nil, zero value otherwise.

### GetMotherNameOk

`func (o *DocumentProperties) GetMotherNameOk() (*string, bool)`

GetMotherNameOk returns a tuple with the MotherName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotherName

`func (o *DocumentProperties) SetMotherName(v string)`

SetMotherName sets MotherName field to given value.

### HasMotherName

`func (o *DocumentProperties) HasMotherName() bool`

HasMotherName returns a boolean if a field has been set.

### GetReligion

`func (o *DocumentProperties) GetReligion() string`

GetReligion returns the Religion field if non-nil, zero value otherwise.

### GetReligionOk

`func (o *DocumentProperties) GetReligionOk() (*string, bool)`

GetReligionOk returns a tuple with the Religion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReligion

`func (o *DocumentProperties) SetReligion(v string)`

SetReligion sets Religion field to given value.

### HasReligion

`func (o *DocumentProperties) HasReligion() bool`

HasReligion returns a boolean if a field has been set.

### GetTypeOfPermit

`func (o *DocumentProperties) GetTypeOfPermit() string`

GetTypeOfPermit returns the TypeOfPermit field if non-nil, zero value otherwise.

### GetTypeOfPermitOk

`func (o *DocumentProperties) GetTypeOfPermitOk() (*string, bool)`

GetTypeOfPermitOk returns a tuple with the TypeOfPermit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfPermit

`func (o *DocumentProperties) SetTypeOfPermit(v string)`

SetTypeOfPermit sets TypeOfPermit field to given value.

### HasTypeOfPermit

`func (o *DocumentProperties) HasTypeOfPermit() bool`

HasTypeOfPermit returns a boolean if a field has been set.

### GetVersionNumber

`func (o *DocumentProperties) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *DocumentProperties) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *DocumentProperties) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *DocumentProperties) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetDocumentSubtype

`func (o *DocumentProperties) GetDocumentSubtype() string`

GetDocumentSubtype returns the DocumentSubtype field if non-nil, zero value otherwise.

### GetDocumentSubtypeOk

`func (o *DocumentProperties) GetDocumentSubtypeOk() (*string, bool)`

GetDocumentSubtypeOk returns a tuple with the DocumentSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentSubtype

`func (o *DocumentProperties) SetDocumentSubtype(v string)`

SetDocumentSubtype sets DocumentSubtype field to given value.

### HasDocumentSubtype

`func (o *DocumentProperties) HasDocumentSubtype() bool`

HasDocumentSubtype returns a boolean if a field has been set.

### GetProfession

`func (o *DocumentProperties) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *DocumentProperties) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *DocumentProperties) SetProfession(v string)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *DocumentProperties) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetSecurityDocumentNumber

`func (o *DocumentProperties) GetSecurityDocumentNumber() string`

GetSecurityDocumentNumber returns the SecurityDocumentNumber field if non-nil, zero value otherwise.

### GetSecurityDocumentNumberOk

`func (o *DocumentProperties) GetSecurityDocumentNumberOk() (*string, bool)`

GetSecurityDocumentNumberOk returns a tuple with the SecurityDocumentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityDocumentNumber

`func (o *DocumentProperties) SetSecurityDocumentNumber(v string)`

SetSecurityDocumentNumber sets SecurityDocumentNumber field to given value.

### HasSecurityDocumentNumber

`func (o *DocumentProperties) HasSecurityDocumentNumber() bool`

HasSecurityDocumentNumber returns a boolean if a field has been set.

### GetTaxNumber

`func (o *DocumentProperties) GetTaxNumber() string`

GetTaxNumber returns the TaxNumber field if non-nil, zero value otherwise.

### GetTaxNumberOk

`func (o *DocumentProperties) GetTaxNumberOk() (*string, bool)`

GetTaxNumberOk returns a tuple with the TaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxNumber

`func (o *DocumentProperties) SetTaxNumber(v string)`

SetTaxNumber sets TaxNumber field to given value.

### HasTaxNumber

`func (o *DocumentProperties) HasTaxNumber() bool`

HasTaxNumber returns a boolean if a field has been set.

### GetNistIdentityEvidenceStrength

`func (o *DocumentProperties) GetNistIdentityEvidenceStrength() string`

GetNistIdentityEvidenceStrength returns the NistIdentityEvidenceStrength field if non-nil, zero value otherwise.

### GetNistIdentityEvidenceStrengthOk

`func (o *DocumentProperties) GetNistIdentityEvidenceStrengthOk() (*string, bool)`

GetNistIdentityEvidenceStrengthOk returns a tuple with the NistIdentityEvidenceStrength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNistIdentityEvidenceStrength

`func (o *DocumentProperties) SetNistIdentityEvidenceStrength(v string)`

SetNistIdentityEvidenceStrength sets NistIdentityEvidenceStrength field to given value.

### HasNistIdentityEvidenceStrength

`func (o *DocumentProperties) HasNistIdentityEvidenceStrength() bool`

HasNistIdentityEvidenceStrength returns a boolean if a field has been set.

### GetHasIssuanceConfirmation

`func (o *DocumentProperties) GetHasIssuanceConfirmation() string`

GetHasIssuanceConfirmation returns the HasIssuanceConfirmation field if non-nil, zero value otherwise.

### GetHasIssuanceConfirmationOk

`func (o *DocumentProperties) GetHasIssuanceConfirmationOk() (*string, bool)`

GetHasIssuanceConfirmationOk returns a tuple with the HasIssuanceConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIssuanceConfirmation

`func (o *DocumentProperties) SetHasIssuanceConfirmation(v string)`

SetHasIssuanceConfirmation sets HasIssuanceConfirmation field to given value.

### HasHasIssuanceConfirmation

`func (o *DocumentProperties) HasHasIssuanceConfirmation() bool`

HasHasIssuanceConfirmation returns a boolean if a field has been set.

### GetRealIdCompliance

`func (o *DocumentProperties) GetRealIdCompliance() bool`

GetRealIdCompliance returns the RealIdCompliance field if non-nil, zero value otherwise.

### GetRealIdComplianceOk

`func (o *DocumentProperties) GetRealIdComplianceOk() (*bool, bool)`

GetRealIdComplianceOk returns a tuple with the RealIdCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealIdCompliance

`func (o *DocumentProperties) SetRealIdCompliance(v bool)`

SetRealIdCompliance sets RealIdCompliance field to given value.

### HasRealIdCompliance

`func (o *DocumentProperties) HasRealIdCompliance() bool`

HasRealIdCompliance returns a boolean if a field has been set.

### GetSecurityTier

`func (o *DocumentProperties) GetSecurityTier() string`

GetSecurityTier returns the SecurityTier field if non-nil, zero value otherwise.

### GetSecurityTierOk

`func (o *DocumentProperties) GetSecurityTierOk() (*string, bool)`

GetSecurityTierOk returns a tuple with the SecurityTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityTier

`func (o *DocumentProperties) SetSecurityTier(v string)`

SetSecurityTier sets SecurityTier field to given value.

### HasSecurityTier

`func (o *DocumentProperties) HasSecurityTier() bool`

HasSecurityTier returns a boolean if a field has been set.

### GetAddressLines

`func (o *DocumentProperties) GetAddressLines() DocumentPropertiesAddressLines`

GetAddressLines returns the AddressLines field if non-nil, zero value otherwise.

### GetAddressLinesOk

`func (o *DocumentProperties) GetAddressLinesOk() (*DocumentPropertiesAddressLines, bool)`

GetAddressLinesOk returns a tuple with the AddressLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLines

`func (o *DocumentProperties) SetAddressLines(v DocumentPropertiesAddressLines)`

SetAddressLines sets AddressLines field to given value.

### HasAddressLines

`func (o *DocumentProperties) HasAddressLines() bool`

HasAddressLines returns a boolean if a field has been set.

### GetBarcode

`func (o *DocumentProperties) GetBarcode() []DocumentPropertiesBarcodeInner`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *DocumentProperties) GetBarcodeOk() (*[]DocumentPropertiesBarcodeInner, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *DocumentProperties) SetBarcode(v []DocumentPropertiesBarcodeInner)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *DocumentProperties) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetNfc

`func (o *DocumentProperties) GetNfc() DocumentPropertiesNfc`

GetNfc returns the Nfc field if non-nil, zero value otherwise.

### GetNfcOk

`func (o *DocumentProperties) GetNfcOk() (*DocumentPropertiesNfc, bool)`

GetNfcOk returns a tuple with the Nfc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfc

`func (o *DocumentProperties) SetNfc(v DocumentPropertiesNfc)`

SetNfc sets Nfc field to given value.

### HasNfc

`func (o *DocumentProperties) HasNfc() bool`

HasNfc returns a boolean if a field has been set.

### GetDrivingLicenceInformation

`func (o *DocumentProperties) GetDrivingLicenceInformation() DocumentPropertiesDrivingLicenceInformation`

GetDrivingLicenceInformation returns the DrivingLicenceInformation field if non-nil, zero value otherwise.

### GetDrivingLicenceInformationOk

`func (o *DocumentProperties) GetDrivingLicenceInformationOk() (*DocumentPropertiesDrivingLicenceInformation, bool)`

GetDrivingLicenceInformationOk returns a tuple with the DrivingLicenceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivingLicenceInformation

`func (o *DocumentProperties) SetDrivingLicenceInformation(v DocumentPropertiesDrivingLicenceInformation)`

SetDrivingLicenceInformation sets DrivingLicenceInformation field to given value.

### HasDrivingLicenceInformation

`func (o *DocumentProperties) HasDrivingLicenceInformation() bool`

HasDrivingLicenceInformation returns a boolean if a field has been set.

### GetDocumentClassification

`func (o *DocumentProperties) GetDocumentClassification() DocumentPropertiesDocumentClassification`

GetDocumentClassification returns the DocumentClassification field if non-nil, zero value otherwise.

### GetDocumentClassificationOk

`func (o *DocumentProperties) GetDocumentClassificationOk() (*DocumentPropertiesDocumentClassification, bool)`

GetDocumentClassificationOk returns a tuple with the DocumentClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentClassification

`func (o *DocumentProperties) SetDocumentClassification(v DocumentPropertiesDocumentClassification)`

SetDocumentClassification sets DocumentClassification field to given value.

### HasDocumentClassification

`func (o *DocumentProperties) HasDocumentClassification() bool`

HasDocumentClassification returns a boolean if a field has been set.

### GetExtractedData

`func (o *DocumentProperties) GetExtractedData() DocumentPropertiesExtractedData`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *DocumentProperties) GetExtractedDataOk() (*DocumentPropertiesExtractedData, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *DocumentProperties) SetExtractedData(v DocumentPropertiesExtractedData)`

SetExtractedData sets ExtractedData field to given value.

### HasExtractedData

`func (o *DocumentProperties) HasExtractedData() bool`

HasExtractedData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


