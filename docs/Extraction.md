# Extraction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentId** | Pointer to **string** | The unique identifier of the uploaded document. | [optional] 
**DocumentClassification** | Pointer to [**ExtractionDocumentClassification**](ExtractionDocumentClassification.md) |  | [optional] 
**ExtractedData** | Pointer to [**ExtractionExtractedData**](ExtractionExtractedData.md) |  | [optional] 

## Methods

### NewExtraction

`func NewExtraction() *Extraction`

NewExtraction instantiates a new Extraction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionWithDefaults

`func NewExtractionWithDefaults() *Extraction`

NewExtractionWithDefaults instantiates a new Extraction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentId

`func (o *Extraction) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *Extraction) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *Extraction) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *Extraction) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetDocumentClassification

`func (o *Extraction) GetDocumentClassification() ExtractionDocumentClassification`

GetDocumentClassification returns the DocumentClassification field if non-nil, zero value otherwise.

### GetDocumentClassificationOk

`func (o *Extraction) GetDocumentClassificationOk() (*ExtractionDocumentClassification, bool)`

GetDocumentClassificationOk returns a tuple with the DocumentClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentClassification

`func (o *Extraction) SetDocumentClassification(v ExtractionDocumentClassification)`

SetDocumentClassification sets DocumentClassification field to given value.

### HasDocumentClassification

`func (o *Extraction) HasDocumentClassification() bool`

HasDocumentClassification returns a boolean if a field has been set.

### GetExtractedData

`func (o *Extraction) GetExtractedData() ExtractionExtractedData`

GetExtractedData returns the ExtractedData field if non-nil, zero value otherwise.

### GetExtractedDataOk

`func (o *Extraction) GetExtractedDataOk() (*ExtractionExtractedData, bool)`

GetExtractedDataOk returns a tuple with the ExtractedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedData

`func (o *Extraction) SetExtractedData(v ExtractionExtractedData)`

SetExtractedData sets ExtractedData field to given value.

### HasExtractedData

`func (o *Extraction) HasExtractedData() bool`

HasExtractedData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


