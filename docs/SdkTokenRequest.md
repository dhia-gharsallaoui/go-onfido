# SdkTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantId** | **string** | The unique identifier of the applicant | 
**Referrer** | Pointer to **string** | The referrer URL pattern | [optional] 
**ApplicationId** | Pointer to **string** | The application ID (iOS or Android) | [optional] 
**CrossDeviceUrl** | Pointer to **string** | The URL to be used by the Web SDK for the cross device flow. | [optional] 

## Methods

### NewSdkTokenRequest

`func NewSdkTokenRequest(applicantId string, ) *SdkTokenRequest`

NewSdkTokenRequest instantiates a new SdkTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdkTokenRequestWithDefaults

`func NewSdkTokenRequestWithDefaults() *SdkTokenRequest`

NewSdkTokenRequestWithDefaults instantiates a new SdkTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantId

`func (o *SdkTokenRequest) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *SdkTokenRequest) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *SdkTokenRequest) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetReferrer

`func (o *SdkTokenRequest) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *SdkTokenRequest) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *SdkTokenRequest) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *SdkTokenRequest) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetApplicationId

`func (o *SdkTokenRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *SdkTokenRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *SdkTokenRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *SdkTokenRequest) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetCrossDeviceUrl

`func (o *SdkTokenRequest) GetCrossDeviceUrl() string`

GetCrossDeviceUrl returns the CrossDeviceUrl field if non-nil, zero value otherwise.

### GetCrossDeviceUrlOk

`func (o *SdkTokenRequest) GetCrossDeviceUrlOk() (*string, bool)`

GetCrossDeviceUrlOk returns a tuple with the CrossDeviceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossDeviceUrl

`func (o *SdkTokenRequest) SetCrossDeviceUrl(v string)`

SetCrossDeviceUrl sets CrossDeviceUrl field to given value.

### HasCrossDeviceUrl

`func (o *SdkTokenRequest) HasCrossDeviceUrl() bool`

HasCrossDeviceUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


