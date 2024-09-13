# CheckShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookIds** | Pointer to **[]string** | An array of webhook ids describing which webhooks to trigger for this check. | [optional] 
**ApplicantId** | **string** | The ID of the applicant to do the check on. | 
**ApplicantProvidesData** | Pointer to **bool** | Send an applicant form to applicant to complete to proceed with check. Defaults to false.  | [optional] 
**Tags** | Pointer to **[]string** | Array of tags being assigned to this check. | [optional] 
**RedirectUri** | Pointer to **string** | For checks where &#x60;applicant_provides_data&#x60; is &#x60;true&#x60;, redirect to this URI when the applicant has submitted their data. | [optional] 

## Methods

### NewCheckShared

`func NewCheckShared(applicantId string, ) *CheckShared`

NewCheckShared instantiates a new CheckShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckSharedWithDefaults

`func NewCheckSharedWithDefaults() *CheckShared`

NewCheckSharedWithDefaults instantiates a new CheckShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookIds

`func (o *CheckShared) GetWebhookIds() []string`

GetWebhookIds returns the WebhookIds field if non-nil, zero value otherwise.

### GetWebhookIdsOk

`func (o *CheckShared) GetWebhookIdsOk() (*[]string, bool)`

GetWebhookIdsOk returns a tuple with the WebhookIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookIds

`func (o *CheckShared) SetWebhookIds(v []string)`

SetWebhookIds sets WebhookIds field to given value.

### HasWebhookIds

`func (o *CheckShared) HasWebhookIds() bool`

HasWebhookIds returns a boolean if a field has been set.

### GetApplicantId

`func (o *CheckShared) GetApplicantId() string`

GetApplicantId returns the ApplicantId field if non-nil, zero value otherwise.

### GetApplicantIdOk

`func (o *CheckShared) GetApplicantIdOk() (*string, bool)`

GetApplicantIdOk returns a tuple with the ApplicantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantId

`func (o *CheckShared) SetApplicantId(v string)`

SetApplicantId sets ApplicantId field to given value.


### GetApplicantProvidesData

`func (o *CheckShared) GetApplicantProvidesData() bool`

GetApplicantProvidesData returns the ApplicantProvidesData field if non-nil, zero value otherwise.

### GetApplicantProvidesDataOk

`func (o *CheckShared) GetApplicantProvidesDataOk() (*bool, bool)`

GetApplicantProvidesDataOk returns a tuple with the ApplicantProvidesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantProvidesData

`func (o *CheckShared) SetApplicantProvidesData(v bool)`

SetApplicantProvidesData sets ApplicantProvidesData field to given value.

### HasApplicantProvidesData

`func (o *CheckShared) HasApplicantProvidesData() bool`

HasApplicantProvidesData returns a boolean if a field has been set.

### GetTags

`func (o *CheckShared) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckShared) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckShared) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckShared) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRedirectUri

`func (o *CheckShared) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *CheckShared) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *CheckShared) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *CheckShared) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


