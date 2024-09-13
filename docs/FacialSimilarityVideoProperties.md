# FacialSimilarityVideoProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | Pointer to **float32** | A floating point number between 0 and 1. The closer the score is to 0, the more likely it is to be a spoof (i.e. videos of digital screens, masks or print-outs). Conversely, the closer it is to 1, the less likely it is to be a spoof.  | [optional] 

## Methods

### NewFacialSimilarityVideoProperties

`func NewFacialSimilarityVideoProperties() *FacialSimilarityVideoProperties`

NewFacialSimilarityVideoProperties instantiates a new FacialSimilarityVideoProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacialSimilarityVideoPropertiesWithDefaults

`func NewFacialSimilarityVideoPropertiesWithDefaults() *FacialSimilarityVideoProperties`

NewFacialSimilarityVideoPropertiesWithDefaults instantiates a new FacialSimilarityVideoProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *FacialSimilarityVideoProperties) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *FacialSimilarityVideoProperties) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *FacialSimilarityVideoProperties) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *FacialSimilarityVideoProperties) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


