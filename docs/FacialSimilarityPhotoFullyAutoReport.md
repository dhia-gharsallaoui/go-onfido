# FacialSimilarityPhotoFullyAutoReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier for the report. Read-only. | 
**CreatedAt** | Pointer to **time.Time** | The date and time at which the report was first initiated. Read-only. | [optional] 
**Href** | Pointer to **string** | The API endpoint to retrieve the report. Read-only. | [optional] 
**Status** | Pointer to [**ReportStatus**](ReportStatus.md) |  | [optional] 
**Result** | Pointer to [**ReportResult**](ReportResult.md) |  | [optional] 
**SubResult** | Pointer to [**ReportSubResult**](ReportSubResult.md) |  | [optional] 
**CheckId** | Pointer to **string** | The ID of the check to which the report belongs. Read-only. | [optional] 
**Documents** | Pointer to [**[]ReportDocument**](ReportDocument.md) | Array of objects with document ids that were used in the Onfido engine. [ONLY POPULATED FOR DOCUMENT AND FACIAL SIMILARITY REPORTS] | [optional] 
**Name** | [**ReportName**](ReportName.md) |  | 
**Breakdown** | Pointer to [**FacialSimilarityPhotoFullyAutoBreakdown**](FacialSimilarityPhotoFullyAutoBreakdown.md) |  | [optional] 
**Properties** | Pointer to [**FacialSimilarityPhotoFullyAutoProperties**](FacialSimilarityPhotoFullyAutoProperties.md) |  | [optional] 

## Methods

### NewFacialSimilarityPhotoFullyAutoReport

`func NewFacialSimilarityPhotoFullyAutoReport(id string, name ReportName, ) *FacialSimilarityPhotoFullyAutoReport`

NewFacialSimilarityPhotoFullyAutoReport instantiates a new FacialSimilarityPhotoFullyAutoReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacialSimilarityPhotoFullyAutoReportWithDefaults

`func NewFacialSimilarityPhotoFullyAutoReportWithDefaults() *FacialSimilarityPhotoFullyAutoReport`

NewFacialSimilarityPhotoFullyAutoReportWithDefaults instantiates a new FacialSimilarityPhotoFullyAutoReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FacialSimilarityPhotoFullyAutoReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FacialSimilarityPhotoFullyAutoReport) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *FacialSimilarityPhotoFullyAutoReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FacialSimilarityPhotoFullyAutoReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FacialSimilarityPhotoFullyAutoReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *FacialSimilarityPhotoFullyAutoReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FacialSimilarityPhotoFullyAutoReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FacialSimilarityPhotoFullyAutoReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *FacialSimilarityPhotoFullyAutoReport) GetStatus() ReportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetStatusOk() (*ReportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FacialSimilarityPhotoFullyAutoReport) SetStatus(v ReportStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FacialSimilarityPhotoFullyAutoReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *FacialSimilarityPhotoFullyAutoReport) GetResult() ReportResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetResultOk() (*ReportResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FacialSimilarityPhotoFullyAutoReport) SetResult(v ReportResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *FacialSimilarityPhotoFullyAutoReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *FacialSimilarityPhotoFullyAutoReport) GetSubResult() ReportSubResult`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetSubResultOk() (*ReportSubResult, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *FacialSimilarityPhotoFullyAutoReport) SetSubResult(v ReportSubResult)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *FacialSimilarityPhotoFullyAutoReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *FacialSimilarityPhotoFullyAutoReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *FacialSimilarityPhotoFullyAutoReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *FacialSimilarityPhotoFullyAutoReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *FacialSimilarityPhotoFullyAutoReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *FacialSimilarityPhotoFullyAutoReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *FacialSimilarityPhotoFullyAutoReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetName

`func (o *FacialSimilarityPhotoFullyAutoReport) GetName() ReportName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetNameOk() (*ReportName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FacialSimilarityPhotoFullyAutoReport) SetName(v ReportName)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *FacialSimilarityPhotoFullyAutoReport) GetBreakdown() FacialSimilarityPhotoFullyAutoBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetBreakdownOk() (*FacialSimilarityPhotoFullyAutoBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *FacialSimilarityPhotoFullyAutoReport) SetBreakdown(v FacialSimilarityPhotoFullyAutoBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *FacialSimilarityPhotoFullyAutoReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *FacialSimilarityPhotoFullyAutoReport) GetProperties() FacialSimilarityPhotoFullyAutoProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FacialSimilarityPhotoFullyAutoReport) GetPropertiesOk() (*FacialSimilarityPhotoFullyAutoProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FacialSimilarityPhotoFullyAutoReport) SetProperties(v FacialSimilarityPhotoFullyAutoProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FacialSimilarityPhotoFullyAutoReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


