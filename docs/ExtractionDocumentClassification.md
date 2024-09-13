# ExtractionDocumentClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuingCountry** | Pointer to [**CountryCodes**](CountryCodes.md) | Document country in 3-letter ISO code. | [optional] 
**DocumentType** | Pointer to [**DocumentTypes**](DocumentTypes.md) | Type of document. | [optional] 
**IssuingState** | Pointer to **string** | The state that issued the document (if available). | [optional] 
**Subtype** | Pointer to **string** | The document subtype (if available). | [optional] 
**Version** | Pointer to **string** | The document issuing version (if available). | [optional] 

## Methods

### NewExtractionDocumentClassification

`func NewExtractionDocumentClassification() *ExtractionDocumentClassification`

NewExtractionDocumentClassification instantiates a new ExtractionDocumentClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionDocumentClassificationWithDefaults

`func NewExtractionDocumentClassificationWithDefaults() *ExtractionDocumentClassification`

NewExtractionDocumentClassificationWithDefaults instantiates a new ExtractionDocumentClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuingCountry

`func (o *ExtractionDocumentClassification) GetIssuingCountry() CountryCodes`

GetIssuingCountry returns the IssuingCountry field if non-nil, zero value otherwise.

### GetIssuingCountryOk

`func (o *ExtractionDocumentClassification) GetIssuingCountryOk() (*CountryCodes, bool)`

GetIssuingCountryOk returns a tuple with the IssuingCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingCountry

`func (o *ExtractionDocumentClassification) SetIssuingCountry(v CountryCodes)`

SetIssuingCountry sets IssuingCountry field to given value.

### HasIssuingCountry

`func (o *ExtractionDocumentClassification) HasIssuingCountry() bool`

HasIssuingCountry returns a boolean if a field has been set.

### GetDocumentType

`func (o *ExtractionDocumentClassification) GetDocumentType() DocumentTypes`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *ExtractionDocumentClassification) GetDocumentTypeOk() (*DocumentTypes, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *ExtractionDocumentClassification) SetDocumentType(v DocumentTypes)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *ExtractionDocumentClassification) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetIssuingState

`func (o *ExtractionDocumentClassification) GetIssuingState() string`

GetIssuingState returns the IssuingState field if non-nil, zero value otherwise.

### GetIssuingStateOk

`func (o *ExtractionDocumentClassification) GetIssuingStateOk() (*string, bool)`

GetIssuingStateOk returns a tuple with the IssuingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuingState

`func (o *ExtractionDocumentClassification) SetIssuingState(v string)`

SetIssuingState sets IssuingState field to given value.

### HasIssuingState

`func (o *ExtractionDocumentClassification) HasIssuingState() bool`

HasIssuingState returns a boolean if a field has been set.

### GetSubtype

`func (o *ExtractionDocumentClassification) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *ExtractionDocumentClassification) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *ExtractionDocumentClassification) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *ExtractionDocumentClassification) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetVersion

`func (o *ExtractionDocumentClassification) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtractionDocumentClassification) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtractionDocumentClassification) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtractionDocumentClassification) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


