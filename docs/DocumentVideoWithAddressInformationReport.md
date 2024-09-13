# DocumentVideoWithAddressInformationReport

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
**Breakdown** | Pointer to [**DocumentBreakdown**](DocumentBreakdown.md) |  | [optional] 
**Properties** | Pointer to [**DocumentProperties**](DocumentProperties.md) |  | [optional] 

## Methods

### NewDocumentVideoWithAddressInformationReport

`func NewDocumentVideoWithAddressInformationReport(id string, name ReportName, ) *DocumentVideoWithAddressInformationReport`

NewDocumentVideoWithAddressInformationReport instantiates a new DocumentVideoWithAddressInformationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentVideoWithAddressInformationReportWithDefaults

`func NewDocumentVideoWithAddressInformationReportWithDefaults() *DocumentVideoWithAddressInformationReport`

NewDocumentVideoWithAddressInformationReportWithDefaults instantiates a new DocumentVideoWithAddressInformationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocumentVideoWithAddressInformationReport) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentVideoWithAddressInformationReport) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentVideoWithAddressInformationReport) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DocumentVideoWithAddressInformationReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentVideoWithAddressInformationReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentVideoWithAddressInformationReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentVideoWithAddressInformationReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *DocumentVideoWithAddressInformationReport) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DocumentVideoWithAddressInformationReport) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DocumentVideoWithAddressInformationReport) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DocumentVideoWithAddressInformationReport) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentVideoWithAddressInformationReport) GetStatus() ReportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentVideoWithAddressInformationReport) GetStatusOk() (*ReportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentVideoWithAddressInformationReport) SetStatus(v ReportStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentVideoWithAddressInformationReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *DocumentVideoWithAddressInformationReport) GetResult() ReportResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DocumentVideoWithAddressInformationReport) GetResultOk() (*ReportResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DocumentVideoWithAddressInformationReport) SetResult(v ReportResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DocumentVideoWithAddressInformationReport) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSubResult

`func (o *DocumentVideoWithAddressInformationReport) GetSubResult() ReportSubResult`

GetSubResult returns the SubResult field if non-nil, zero value otherwise.

### GetSubResultOk

`func (o *DocumentVideoWithAddressInformationReport) GetSubResultOk() (*ReportSubResult, bool)`

GetSubResultOk returns a tuple with the SubResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubResult

`func (o *DocumentVideoWithAddressInformationReport) SetSubResult(v ReportSubResult)`

SetSubResult sets SubResult field to given value.

### HasSubResult

`func (o *DocumentVideoWithAddressInformationReport) HasSubResult() bool`

HasSubResult returns a boolean if a field has been set.

### GetCheckId

`func (o *DocumentVideoWithAddressInformationReport) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *DocumentVideoWithAddressInformationReport) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *DocumentVideoWithAddressInformationReport) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *DocumentVideoWithAddressInformationReport) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDocuments

`func (o *DocumentVideoWithAddressInformationReport) GetDocuments() []ReportDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *DocumentVideoWithAddressInformationReport) GetDocumentsOk() (*[]ReportDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *DocumentVideoWithAddressInformationReport) SetDocuments(v []ReportDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *DocumentVideoWithAddressInformationReport) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetName

`func (o *DocumentVideoWithAddressInformationReport) GetName() ReportName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentVideoWithAddressInformationReport) GetNameOk() (*ReportName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentVideoWithAddressInformationReport) SetName(v ReportName)`

SetName sets Name field to given value.


### GetBreakdown

`func (o *DocumentVideoWithAddressInformationReport) GetBreakdown() DocumentBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *DocumentVideoWithAddressInformationReport) GetBreakdownOk() (*DocumentBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *DocumentVideoWithAddressInformationReport) SetBreakdown(v DocumentBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *DocumentVideoWithAddressInformationReport) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetProperties

`func (o *DocumentVideoWithAddressInformationReport) GetProperties() DocumentProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DocumentVideoWithAddressInformationReport) GetPropertiesOk() (*DocumentProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DocumentVideoWithAddressInformationReport) SetProperties(v DocumentProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DocumentVideoWithAddressInformationReport) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


