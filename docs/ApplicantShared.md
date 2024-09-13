# ApplicantShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The applicant&#39;s email address. Required if doing a US check, or a UK check for which &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;. | [optional] 
**Dob** | Pointer to **string** | The applicant&#39;s date of birth | [optional] 
**IdNumbers** | Pointer to [**[]IdNumber**](IdNumber.md) |  | [optional] 
**PhoneNumber** | Pointer to **string** | The applicant&#39;s phone number | [optional] 

## Methods

### NewApplicantShared

`func NewApplicantShared() *ApplicantShared`

NewApplicantShared instantiates a new ApplicantShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicantSharedWithDefaults

`func NewApplicantSharedWithDefaults() *ApplicantShared`

NewApplicantSharedWithDefaults instantiates a new ApplicantShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ApplicantShared) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ApplicantShared) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ApplicantShared) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ApplicantShared) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDob

`func (o *ApplicantShared) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *ApplicantShared) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *ApplicantShared) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *ApplicantShared) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetIdNumbers

`func (o *ApplicantShared) GetIdNumbers() []IdNumber`

GetIdNumbers returns the IdNumbers field if non-nil, zero value otherwise.

### GetIdNumbersOk

`func (o *ApplicantShared) GetIdNumbersOk() (*[]IdNumber, bool)`

GetIdNumbersOk returns a tuple with the IdNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNumbers

`func (o *ApplicantShared) SetIdNumbers(v []IdNumber)`

SetIdNumbers sets IdNumbers field to given value.

### HasIdNumbers

`func (o *ApplicantShared) HasIdNumbers() bool`

HasIdNumbers returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *ApplicantShared) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *ApplicantShared) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *ApplicantShared) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *ApplicantShared) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


