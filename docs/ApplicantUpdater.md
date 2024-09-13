# ApplicantUpdater

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The applicant&#39;s email address. Required if doing a US check, or a UK check for which &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;. | [optional] 
**Dob** | Pointer to **string** | The applicant&#39;s date of birth | [optional] 
**IdNumbers** | Pointer to [**[]IdNumber**](IdNumber.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** | The applicant&#39;s phone number | [optional] 
**Consents** | Pointer to [**ConsentsBuilder**](ConsentsBuilder.md) |  | [optional] 
**Address** | Pointer to [**AddressBuilder**](AddressBuilder.md) |  | [optional] 
**Location** | Pointer to [**LocationBuilder**](LocationBuilder.md) |  | [optional] 
**FirstName** | Pointer to **string** | The applicant&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The applicant&#39;s surname | [optional] 

## Methods

### NewApplicantUpdater

`func NewApplicantUpdater() *ApplicantUpdater`

NewApplicantUpdater instantiates a new ApplicantUpdater object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicantUpdaterWithDefaults

`func NewApplicantUpdaterWithDefaults() *ApplicantUpdater`

NewApplicantUpdaterWithDefaults instantiates a new ApplicantUpdater object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ApplicantUpdater) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApplicantUpdater) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApplicantUpdater) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApplicantUpdater) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDob

`func (o *ApplicantUpdater) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *ApplicantUpdater) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *ApplicantUpdater) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *ApplicantUpdater) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetIdNumbers

`func (o *ApplicantUpdater) GetIdNumbers() []IdNumber`

GetIdNumbers returns the IdNumbers field if non-nil, zero value otherwise.

### GetIdNumbersOk

`func (o *ApplicantUpdater) GetIdNumbersOk() (*[]IdNumber, bool)`

GetIdNumbersOk returns a tuple with the IdNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumbers

`func (o *ApplicantUpdater) SetIdNumbers(v []IdNumber)`

SetIdNumbers sets IdNumbers field to given value.

### HasIdNumbers

`func (o *ApplicantUpdater) HasIdNumbers() bool`

HasIdNumbers returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ApplicantUpdater) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ApplicantUpdater) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ApplicantUpdater) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ApplicantUpdater) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetConsents

`func (o *ApplicantUpdater) GetConsents() ConsentsBuilder`

GetConsents returns the Consents field if non-nil, zero value otherwise.

### GetConsentsOk

`func (o *ApplicantUpdater) GetConsentsOk() (*ConsentsBuilder, bool)`

GetConsentsOk returns a tuple with the Consents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsents

`func (o *ApplicantUpdater) SetConsents(v ConsentsBuilder)`

SetConsents sets Consents field to given value.

### HasConsents

`func (o *ApplicantUpdater) HasConsents() bool`

HasConsents returns a boolean if a field has been set.

### GetAddress

`func (o *ApplicantUpdater) GetAddress() AddressBuilder`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ApplicantUpdater) GetAddressOk() (*AddressBuilder, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ApplicantUpdater) SetAddress(v AddressBuilder)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ApplicantUpdater) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLocation

`func (o *ApplicantUpdater) GetLocation() LocationBuilder`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ApplicantUpdater) GetLocationOk() (*LocationBuilder, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ApplicantUpdater) SetLocation(v LocationBuilder)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ApplicantUpdater) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetFirstName

`func (o *ApplicantUpdater) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ApplicantUpdater) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ApplicantUpdater) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ApplicantUpdater) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ApplicantUpdater) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ApplicantUpdater) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ApplicantUpdater) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ApplicantUpdater) HasLastName() bool`

HasLastName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


