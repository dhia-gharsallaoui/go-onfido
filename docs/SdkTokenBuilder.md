# SdkTokenBuilder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **string** | The unique identifier of the applicant | 
**Referrer** | Pointer to **string** | The referrer URL pattern | [optional] 
**ApplicationId** | Pointer to **string** | The application ID (iOS or Android) | [optional] 
**CrossDeviceUrl** | Pointer to **string** | The URL to be used by the Web SDK for the cross device flow. | [optional] 

## Methods

### NewSdkTokenBuilder

`func NewSdkTokenBuilder(applicantId string, ) *SdkTokenBuilder`

NewSdkTokenBuilder instantiates a new SdkTokenBuilder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkTokenBuilderWithDefaults

`func NewSdkTokenBuilderWithDefaults() *SdkTokenBuilder`

NewSdkTokenBuilderWithDefaults instantiates a new SdkTokenBuilder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *SdkTokenBuilder) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *SdkTokenBuilder) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *SdkTokenBuilder) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetReferrer

`func (o *SdkTokenBuilder) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *SdkTokenBuilder) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *SdkTokenBuilder) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *SdkTokenBuilder) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetApplicationId

`func (o *SdkTokenBuilder) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *SdkTokenBuilder) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *SdkTokenBuilder) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *SdkTokenBuilder) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetCrossDeviceUrl

`func (o *SdkTokenBuilder) GetCrossDeviceUrl() string`

GetCrossDeviceUrl returns the CrossDeviceUrl field if non-nil, zero value otherwise.

### GetCrossDeviceUrlOk

`func (o *SdkTokenBuilder) GetCrossDeviceUrlOk() (*string, bool)`

GetCrossDeviceUrlOk returns a tuple with the CrossDeviceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossDeviceUrl

`func (o *SdkTokenBuilder) SetCrossDeviceUrl(v string)`

SetCrossDeviceUrl sets CrossDeviceUrl field to given value.

### HasCrossDeviceUrl

`func (o *SdkTokenBuilder) HasCrossDeviceUrl() bool`

HasCrossDeviceUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


