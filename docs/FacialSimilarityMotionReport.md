# FacialSimilarityMotionReport

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
**Breakdown** | Pointer to [**FacialSimilarityMotionBreakdown**](FacialSimilarityMotionBreakdown.md) |  | [optional] 
**Properties** | Pointer to [**FacialSimilarityMotionProperties**](FacialSimilarityMotionProperties.md) |  | [optional] 

## Methods

### NewFacialSimilarityMotionReport

`func NewFacialSimilarityMotionReport(id string, name ReportName, ) *FacialSimilarityMotionReport`

NewFacialSimilarityMotionReport instantiates a new FacialSimilarityMotionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacialSimilarityMotionReportWithDefaults

`func NewFacialSimilarityMotionReportWithDefaults() *FacialSimilarityMotionReport`

NewFacialSimilarityMotionReportWithDefaults instantiates a new FacialSimilarityMotionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FacialSimilarityMotionReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FacialSimilarityMotionReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FacialSimilarityMotionReport) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *FacialSimilarityMotionReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FacialSimilarityMotionReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FacialSimilarityMotionReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FacialSimilarityMotionReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *FacialSimilarityMotionReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FacialSimilarityMotionReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FacialSimilarityMotionReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FacialSimilarityMotionReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *FacialSimilarityMotionReport) GetStatus() ReportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FacialSimilarityMotionReport) GetStatusOk() (*ReportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FacialSimilarityMotionReport) SetStatus(v ReportStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FacialSimilarityMotionReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *FacialSimilarityMotionReport) GetResult() ReportResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *FacialSimilarityMotionReport) GetResultOk() (*ReportResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *FacialSimilarityMotionReport) SetResult(v ReportResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *FacialSimilarityMotionReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *FacialSimilarityMotionReport) GetSubResult() ReportSubResult`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *FacialSimilarityMotionReport) GetSubResultOk() (*ReportSubResult, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *FacialSimilarityMotionReport) SetSubResult(v ReportSubResult)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *FacialSimilarityMotionReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *FacialSimilarityMotionReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *FacialSimilarityMotionReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *FacialSimilarityMotionReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *FacialSimilarityMotionReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *FacialSimilarityMotionReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *FacialSimilarityMotionReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *FacialSimilarityMotionReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *FacialSimilarityMotionReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetName

`func (o *FacialSimilarityMotionReport) GetName() ReportName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FacialSimilarityMotionReport) GetNameOk() (*ReportName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FacialSimilarityMotionReport) SetName(v ReportName)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *FacialSimilarityMotionReport) GetBreakdown() FacialSimilarityMotionBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *FacialSimilarityMotionReport) GetBreakdownOk() (*FacialSimilarityMotionBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *FacialSimilarityMotionReport) SetBreakdown(v FacialSimilarityMotionBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *FacialSimilarityMotionReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *FacialSimilarityMotionReport) GetProperties() FacialSimilarityMotionProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FacialSimilarityMotionReport) GetPropertiesOk() (*FacialSimilarityMotionProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FacialSimilarityMotionReport) SetProperties(v FacialSimilarityMotionProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FacialSimilarityMotionReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


