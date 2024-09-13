# \DefaultAPI

All URIs are relative to *https://api.eu.onfido.com/v3.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReport**](DefaultAPI.md#CancelReport) | **Post** /reports/{report_id}/cancel | Cancel report
[**CompleteTask**](DefaultAPI.md#CompleteTask) | **Post** /workflow_runs/{workflow_run_id}/tasks/{task_id}/complete | Complete Task
[**CreateApplicant**](DefaultAPI.md#CreateApplicant) | **Post** /applicants | Create Applicant
[**CreateCheck**](DefaultAPI.md#CreateCheck) | **Post** /checks | Create a check
[**CreateTimelineFile**](DefaultAPI.md#CreateTimelineFile) | **Post** /workflow_runs/{workflow_run_id}/timeline_file | Create Timeline File for Workflow Run
[**CreateWatchlistMonitor**](DefaultAPI.md#CreateWatchlistMonitor) | **Post** /watchlist_monitors | Create monitor
[**CreateWebhook**](DefaultAPI.md#CreateWebhook) | **Post** /webhooks | Register webhook
[**CreateWorkflowRun**](DefaultAPI.md#CreateWorkflowRun) | **Post** /workflow_runs | Create a Workflow Run
[**DeleteApplicant**](DefaultAPI.md#DeleteApplicant) | **Delete** /applicants/{applicant_id} | Delete Applicant
[**DeleteWatchlistMonitor**](DefaultAPI.md#DeleteWatchlistMonitor) | **Delete** /watchlist_monitors/{monitor_id} | Delete monitor
[**DeleteWebhook**](DefaultAPI.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
[**DownloadCheck**](DefaultAPI.md#DownloadCheck) | **Get** /checks/{check_id}/download | Download check
[**DownloadDocument**](DefaultAPI.md#DownloadDocument) | **Get** /documents/{document_id}/download | Download document
[**DownloadDocumentVideo**](DefaultAPI.md#DownloadDocumentVideo) | **Get** /documents/{document_id}/video/download | Download document video
[**DownloadIdPhoto**](DefaultAPI.md#DownloadIdPhoto) | **Get** /id_photos/{id_photo_id}/download | Download ID photo
[**DownloadLivePhoto**](DefaultAPI.md#DownloadLivePhoto) | **Get** /live_photos/{live_photo_id}/download | Download live photo
[**DownloadLiveVideo**](DefaultAPI.md#DownloadLiveVideo) | **Get** /live_videos/{live_video_id}/download | Download live video
[**DownloadLiveVideoFrame**](DefaultAPI.md#DownloadLiveVideoFrame) | **Get** /live_videos/{live_video_id}/frame | Download live video frame
[**DownloadMotionCapture**](DefaultAPI.md#DownloadMotionCapture) | **Get** /motion_captures/{motion_capture_id}/download | Download motion capture
[**DownloadMotionCaptureFrame**](DefaultAPI.md#DownloadMotionCaptureFrame) | **Get** /motion_captures/{motion_capture_id}/frame | Download motion capture frame
[**DownloadQesDocument**](DefaultAPI.md#DownloadQesDocument) | **Get** /qualified_electronic_signature/documents | Retrieves the signed document or application form
[**DownloadSignedEvidenceFile**](DefaultAPI.md#DownloadSignedEvidenceFile) | **Get** /workflow_runs/{workflow_run_id}/signed_evidence_file | Retrieve Workflow Run Evidence Summary File
[**Extract**](DefaultAPI.md#Extract) | **Post** /extractions | Autofill
[**FindAddresses**](DefaultAPI.md#FindAddresses) | **Get** /addresses/pick | Address Picker
[**FindApplicant**](DefaultAPI.md#FindApplicant) | **Get** /applicants/{applicant_id} | Retrieve Applicant
[**FindCheck**](DefaultAPI.md#FindCheck) | **Get** /checks/{check_id} | Retrieve a Check
[**FindDocument**](DefaultAPI.md#FindDocument) | **Get** /documents/{document_id} | Retrieve document
[**FindIdPhoto**](DefaultAPI.md#FindIdPhoto) | **Get** /id_photos/{id_photo_id} | Retrieve ID photo
[**FindLivePhoto**](DefaultAPI.md#FindLivePhoto) | **Get** /live_photos/{live_photo_id} | Retrieve live photo
[**FindLiveVideo**](DefaultAPI.md#FindLiveVideo) | **Get** /live_videos/{live_video_id} | Retrieve live video
[**FindMotionCapture**](DefaultAPI.md#FindMotionCapture) | **Get** /motion_captures/{motion_capture_id} | Retrieve motion capture
[**FindReport**](DefaultAPI.md#FindReport) | **Get** /reports/{report_id} | Retrieve report
[**FindTask**](DefaultAPI.md#FindTask) | **Get** /workflow_runs/{workflow_run_id}/tasks/{task_id} | Retrieve Task
[**FindTimelineFile**](DefaultAPI.md#FindTimelineFile) | **Get** /workflow_runs/{workflow_run_id}/timeline_file/{timeline_file_id} | Retrieve Timeline File for Workflow Run
[**FindWatchlistMonitor**](DefaultAPI.md#FindWatchlistMonitor) | **Get** /watchlist_monitors/{monitor_id} | Retrieve monitor
[**FindWebhook**](DefaultAPI.md#FindWebhook) | **Get** /webhooks/{webhook_id} | Retrieve a Webhook
[**FindWorkflowRun**](DefaultAPI.md#FindWorkflowRun) | **Get** /workflow_runs/{workflow_run_id} | Retrieve Workflow Run
[**ForceReportCreationFromWatchlistMonitor**](DefaultAPI.md#ForceReportCreationFromWatchlistMonitor) | **Post** /watchlist_monitors/{monitor_id}/new_report | Force new report creation (BETA)
[**GenerateSdkToken**](DefaultAPI.md#GenerateSdkToken) | **Post** /sdk_token | Generate a SDK token
[**ListApplicants**](DefaultAPI.md#ListApplicants) | **Get** /applicants | List Applicants
[**ListChecks**](DefaultAPI.md#ListChecks) | **Get** /checks | Retrieve Checks
[**ListDocuments**](DefaultAPI.md#ListDocuments) | **Get** /documents | List documents
[**ListIdPhotos**](DefaultAPI.md#ListIdPhotos) | **Get** /id_photos | List ID photos
[**ListLivePhotos**](DefaultAPI.md#ListLivePhotos) | **Get** /live_photos | List live photos
[**ListLiveVideos**](DefaultAPI.md#ListLiveVideos) | **Get** /live_videos | List live videos
[**ListMotionCaptures**](DefaultAPI.md#ListMotionCaptures) | **Get** /motion_captures | List motion captures
[**ListRepeatAttempts**](DefaultAPI.md#ListRepeatAttempts) | **Get** /repeat_attempts/{report_id} | Retrieve repeat attempts
[**ListReports**](DefaultAPI.md#ListReports) | **Get** /reports | List reports
[**ListTasks**](DefaultAPI.md#ListTasks) | **Get** /workflow_runs/{workflow_run_id}/tasks | List Tasks
[**ListWatchlistMonitorMatches**](DefaultAPI.md#ListWatchlistMonitorMatches) | **Get** /watchlist_monitors/{monitor_id}/matches | List matches (BETA)
[**ListWatchlistMonitors**](DefaultAPI.md#ListWatchlistMonitors) | **Get** /watchlist_monitors | List monitors
[**ListWebhooks**](DefaultAPI.md#ListWebhooks) | **Get** /webhooks | List webhooks
[**ListWorkflowRuns**](DefaultAPI.md#ListWorkflowRuns) | **Get** /workflow_runs | List Workflow Runs
[**Ping**](DefaultAPI.md#Ping) | **Get** /ping | Ping
[**PostResultsFeedback**](DefaultAPI.md#PostResultsFeedback) | **Post** /results_feedback | Fraud reporting (ALPHA)
[**ResendWebhooks**](DefaultAPI.md#ResendWebhooks) | **Post** /webhooks/resend | Resends webhooks
[**RestoreApplicant**](DefaultAPI.md#RestoreApplicant) | **Post** /applicants/{applicant_id}/restore | Restore Applicant
[**ResumeCheck**](DefaultAPI.md#ResumeCheck) | **Post** /checks/{check_id}/resume | Resume a Check
[**ResumeReport**](DefaultAPI.md#ResumeReport) | **Post** /reports/{report_id}/resume | Resume report
[**UpdateApplicant**](DefaultAPI.md#UpdateApplicant) | **Put** /applicants/{applicant_id} | Update Applicant
[**UpdateWatchlistMonitorMatch**](DefaultAPI.md#UpdateWatchlistMonitorMatch) | **Patch** /watchlist_monitors/{monitor_id}/matches | Set match status (BETA)
[**UpdateWebhook**](DefaultAPI.md#UpdateWebhook) | **Put** /webhooks/{webhook_id} | Edit a webhook
[**UploadDocument**](DefaultAPI.md#UploadDocument) | **Post** /documents | Upload a document
[**UploadIdPhoto**](DefaultAPI.md#UploadIdPhoto) | **Post** /id_photos | Upload ID photo
[**UploadLivePhoto**](DefaultAPI.md#UploadLivePhoto) | **Post** /live_photos | Upload live photo



## CancelReport

> CancelReport(ctx, reportId).Execute()

Cancel report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.CancelReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CancelReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteTask

> CompleteTask(ctx, workflowRunId, taskId).CompleteTaskBuilder(completeTaskBuilder).Execute()

Complete Task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the Workflow Run to which the Task belongs.
	taskId := "taskId_example" // string | The identifier of the Task you want to complete.
	completeTaskBuilder := *openapiclient.NewCompleteTaskBuilder(openapiclient.Complete_Task_Data_Builder{ArrayOfMapmapOfStringAny: new([]map[string]interface{})}) // CompleteTaskBuilder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.CompleteTask(context.Background(), workflowRunId, taskId).CompleteTaskBuilder(completeTaskBuilder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CompleteTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run to which the Task belongs. | 
**taskId** | **string** | The identifier of the Task you want to complete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **completeTaskBuilder** | [**CompleteTaskBuilder**](CompleteTaskBuilder.md) |  | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicant

> Applicant CreateApplicant(ctx).ApplicantBuilder(applicantBuilder).Execute()

Create Applicant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantBuilder := *openapiclient.NewApplicantBuilder("FirstName_example", "LastName_example") // ApplicantBuilder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateApplicant(context.Background()).ApplicantBuilder(applicantBuilder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicant`: Applicant
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateApplicant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantBuilder** | [**ApplicantBuilder**](ApplicantBuilder.md) |  | 

### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCheck

> Check CreateCheck(ctx).CheckBuilder(checkBuilder).Execute()

Create a check



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	checkBuilder := *openapiclient.NewCheckBuilder("ApplicantId_example", []openapiclient.ReportName{openapiclient.report_name("document")}) // CheckBuilder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateCheck(context.Background()).CheckBuilder(checkBuilder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCheck`: Check
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkBuilder** | [**CheckBuilder**](CheckBuilder.md) |  | 

### Return type

[**Check**](Check.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTimelineFile

> TimelineFileReference CreateTimelineFile(ctx, workflowRunId).Execute()

Create Timeline File for Workflow Run



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the Workflow Run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateTimelineFile(context.Background(), workflowRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateTimelineFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTimelineFile`: TimelineFileReference
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateTimelineFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTimelineFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TimelineFileReference**](TimelineFileReference.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWatchlistMonitor

> WatchlistMonitor CreateWatchlistMonitor(ctx).WatchlistMonitorBuilder(watchlistMonitorBuilder).Execute()

Create monitor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	watchlistMonitorBuilder := *openapiclient.NewWatchlistMonitorBuilder("ApplicantId_example", "ReportName_example") // WatchlistMonitorBuilder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateWatchlistMonitor(context.Background()).WatchlistMonitorBuilder(watchlistMonitorBuilder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateWatchlistMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWatchlistMonitor`: WatchlistMonitor
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateWatchlistMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWatchlistMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **watchlistMonitorBuilder** | [**WatchlistMonitorBuilder**](WatchlistMonitorBuilder.md) |  | 

### Return type

[**WatchlistMonitor**](WatchlistMonitor.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> Webhook CreateWebhook(ctx).WebhookBuilder(webhookBuilder).Execute()

Register webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	webhookBuilder := *openapiclient.NewWebhookBuilder("Url_example") // WebhookBuilder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateWebhook(context.Background()).WebhookBuilder(webhookBuilder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookBuilder** | [**WebhookBuilder**](WebhookBuilder.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowRun

> WorkflowRun CreateWorkflowRun(ctx).WorkflowRunBuilder(workflowRunBuilder).Execute()

Create a Workflow Run



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunBuilder := *openapiclient.NewWorkflowRunBuilder("ApplicantId_example", "WorkflowId_example") // WorkflowRunBuilder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateWorkflowRun(context.Background()).WorkflowRunBuilder(workflowRunBuilder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateWorkflowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflowRun`: WorkflowRun
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateWorkflowRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowRunBuilder** | [**WorkflowRunBuilder**](WorkflowRunBuilder.md) |  | 

### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicant

> DeleteApplicant(ctx, applicantId).Execute()

Delete Applicant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteApplicant(context.Background(), applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWatchlistMonitor

> DeleteWatchlistMonitor(ctx, monitorId).Execute()

Delete monitor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	monitorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteWatchlistMonitor(context.Background(), monitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWatchlistMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWatchlistMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, webhookId).Execute()

Delete a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadCheck

> *os.File DownloadCheck(ctx, checkId).Execute()

Download check



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	checkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadCheck(context.Background(), checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadCheck`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDocument

> *os.File DownloadDocument(ctx, documentId).Execute()

Download document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadDocument(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDocument`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDocumentVideo

> *os.File DownloadDocumentVideo(ctx, documentId).Execute()

Download document video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	documentId := "documentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadDocumentVideo(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadDocumentVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDocumentVideo`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadDocumentVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDocumentVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadIdPhoto

> *os.File DownloadIdPhoto(ctx, idPhotoId).Execute()

Download ID photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	idPhotoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID photo's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadIdPhoto(context.Background(), idPhotoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadIdPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadIdPhoto`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadIdPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idPhotoId** | **string** | The ID photo&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadIdPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLivePhoto

> *os.File DownloadLivePhoto(ctx, livePhotoId).Execute()

Download live photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	livePhotoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The live photo's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadLivePhoto(context.Background(), livePhotoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadLivePhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLivePhoto`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadLivePhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**livePhotoId** | **string** | The live photo&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLiveVideo

> *os.File DownloadLiveVideo(ctx, liveVideoId).Execute()

Download live video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	liveVideoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The live video's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadLiveVideo(context.Background(), liveVideoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadLiveVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLiveVideo`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadLiveVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLiveVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLiveVideoFrame

> *os.File DownloadLiveVideoFrame(ctx, liveVideoId).Execute()

Download live video frame



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	liveVideoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The live video's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadLiveVideoFrame(context.Background(), liveVideoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadLiveVideoFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLiveVideoFrame`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadLiveVideoFrame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLiveVideoFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadMotionCapture

> *os.File DownloadMotionCapture(ctx, motionCaptureId).Execute()

Download motion capture



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	motionCaptureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The motion capture's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadMotionCapture(context.Background(), motionCaptureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadMotionCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadMotionCapture`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadMotionCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**motionCaptureId** | **string** | The motion capture&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMotionCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadMotionCaptureFrame

> *os.File DownloadMotionCaptureFrame(ctx, motionCaptureId).Execute()

Download motion capture frame



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	motionCaptureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The motion capture's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadMotionCaptureFrame(context.Background(), motionCaptureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadMotionCaptureFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadMotionCaptureFrame`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadMotionCaptureFrame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**motionCaptureId** | **string** | The motion capture&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadMotionCaptureFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadQesDocument

> *os.File DownloadQesDocument(ctx).WorkflowRunId(workflowRunId).FileId(fileId).Execute()

Retrieves the signed document or application form



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the Workflow Run for which you want to retrieve the signed document.
	fileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the file which you want to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadQesDocument(context.Background()).WorkflowRunId(workflowRunId).FileId(fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadQesDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadQesDocument`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadQesDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadQesDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowRunId** | **string** | The unique identifier of the Workflow Run for which you want to retrieve the signed document. | 
 **fileId** | **string** | The unique identifier of the file which you want to retrieve. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadSignedEvidenceFile

> *os.File DownloadSignedEvidenceFile(ctx, workflowRunId).Execute()

Retrieve Workflow Run Evidence Summary File



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow Run ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadSignedEvidenceFile(context.Background(), workflowRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadSignedEvidenceFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadSignedEvidenceFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadSignedEvidenceFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | Workflow Run ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadSignedEvidenceFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Extract

> Extraction Extract(ctx).ExtractRequest(extractRequest).Execute()

Autofill



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	extractRequest := *openapiclient.NewExtractRequest("DocumentId_example") // ExtractRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Extract(context.Background()).ExtractRequest(extractRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Extract``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Extract`: Extraction
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Extract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExtractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **extractRequest** | [**ExtractRequest**](ExtractRequest.md) |  | 

### Return type

[**Extraction**](Extraction.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAddresses

> AddressesList FindAddresses(ctx).Postcode(postcode).Execute()

Address Picker



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	postcode := "postcode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindAddresses(context.Background()).Postcode(postcode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindAddresses`: AddressesList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postcode** | **string** |  | 

### Return type

[**AddressesList**](AddressesList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindApplicant

> Applicant FindApplicant(ctx, applicantId).Execute()

Retrieve Applicant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindApplicant(context.Background(), applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindApplicant`: Applicant
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindApplicant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCheck

> Check FindCheck(ctx, checkId).Execute()

Retrieve a Check



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	checkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindCheck(context.Background(), checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindCheck`: Check
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Check**](Check.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDocument

> Document FindDocument(ctx, documentId).Execute()

Retrieve document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	documentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindDocument(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindDocument`: Document
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Document**](Document.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindIdPhoto

> IdPhoto FindIdPhoto(ctx, idPhotoId).Execute()

Retrieve ID photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	idPhotoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID photo's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindIdPhoto(context.Background(), idPhotoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindIdPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindIdPhoto`: IdPhoto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindIdPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idPhotoId** | **string** | The ID photo&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindIdPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IdPhoto**](IdPhoto.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindLivePhoto

> LivePhoto FindLivePhoto(ctx, livePhotoId).Execute()

Retrieve live photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	livePhotoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The live photo's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindLivePhoto(context.Background(), livePhotoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindLivePhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindLivePhoto`: LivePhoto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindLivePhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**livePhotoId** | **string** | The live photo&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LivePhoto**](LivePhoto.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindLiveVideo

> LiveVideo FindLiveVideo(ctx, liveVideoId).Execute()

Retrieve live video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	liveVideoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The live video's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindLiveVideo(context.Background(), liveVideoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindLiveVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindLiveVideo`: LiveVideo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindLiveVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindLiveVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LiveVideo**](LiveVideo.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindMotionCapture

> MotionCapture FindMotionCapture(ctx, motionCaptureId).Execute()

Retrieve motion capture



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	motionCaptureId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The motion capture's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindMotionCapture(context.Background(), motionCaptureId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindMotionCapture``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindMotionCapture`: MotionCapture
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindMotionCapture`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**motionCaptureId** | **string** | The motion capture&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindMotionCaptureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MotionCapture**](MotionCapture.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindReport

> Report FindReport(ctx, reportId).Execute()

Retrieve report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindReport`: Report
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Report**](Report.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTask

> Task FindTask(ctx, workflowRunId, taskId).Execute()

Retrieve Task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the Workflow Run to which the Task belongs.
	taskId := "taskId_example" // string | The identifier of the Task you want to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindTask(context.Background(), workflowRunId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run to which the Task belongs. | 
**taskId** | **string** | The identifier of the Task you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Task**](Task.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindTimelineFile

> *os.File FindTimelineFile(ctx, workflowRunId, timelineFileId).Execute()

Retrieve Timeline File for Workflow Run



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the Workflow Run.
	timelineFileId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier for the Timefile File.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindTimelineFile(context.Background(), workflowRunId, timelineFileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindTimelineFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindTimelineFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindTimelineFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run. | 
**timelineFileId** | **string** | The unique identifier for the Timefile File. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindTimelineFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWatchlistMonitor

> WatchlistMonitor FindWatchlistMonitor(ctx, monitorId).Execute()

Retrieve monitor



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	monitorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The watchlist monitor's unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindWatchlistMonitor(context.Background(), monitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindWatchlistMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindWatchlistMonitor`: WatchlistMonitor
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindWatchlistMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** | The watchlist monitor&#39;s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWatchlistMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WatchlistMonitor**](WatchlistMonitor.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebhook

> Webhook FindWebhook(ctx, webhookId).Execute()

Retrieve a Webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWorkflowRun

> WorkflowRun FindWorkflowRun(ctx, workflowRunId).Execute()

Retrieve Workflow Run



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the Workflow Run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindWorkflowRun(context.Background(), workflowRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindWorkflowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindWorkflowRun`: WorkflowRun
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForceReportCreationFromWatchlistMonitor

> ForceReportCreationFromWatchlistMonitor(ctx, monitorId).Execute()

Force new report creation (BETA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	monitorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ForceReportCreationFromWatchlistMonitor(context.Background(), monitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ForceReportCreationFromWatchlistMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiForceReportCreationFromWatchlistMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSdkToken

> SdkToken GenerateSdkToken(ctx).SdkTokenBuilder(sdkTokenBuilder).Execute()

Generate a SDK token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	sdkTokenBuilder := *openapiclient.NewSdkTokenBuilder("ApplicantId_example") // SdkTokenBuilder | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GenerateSdkToken(context.Background()).SdkTokenBuilder(sdkTokenBuilder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenerateSdkToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSdkToken`: SdkToken
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenerateSdkToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSdkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sdkTokenBuilder** | [**SdkTokenBuilder**](SdkTokenBuilder.md) |  | 

### Return type

[**SdkToken**](SdkToken.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicants

> ApplicantsList ListApplicants(ctx).Page(page).PerPage(perPage).IncludeDeleted(includeDeleted).Execute()

List Applicants



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	page := int32(56) // int32 | The page to return. The first page is `page=1` (optional) (default to 1)
	perPage := int32(56) // int32 | The number of objects per page. (optional) (default to 20)
	includeDeleted := true // bool | Whether to also include applicants scheduled for deletion. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListApplicants(context.Background()).Page(page).PerPage(perPage).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListApplicants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicants`: ApplicantsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListApplicants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page to return. The first page is &#x60;page&#x3D;1&#x60; | [default to 1]
 **perPage** | **int32** | The number of objects per page. | [default to 20]
 **includeDeleted** | **bool** | Whether to also include applicants scheduled for deletion. | [default to false]

### Return type

[**ApplicantsList**](ApplicantsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChecks

> ChecksList ListChecks(ctx).ApplicantId(applicantId).Execute()

Retrieve Checks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListChecks(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChecks`: ChecksList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** |  | 

### Return type

[**ChecksList**](ChecksList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocuments

> DocumentsList ListDocuments(ctx).ApplicantId(applicantId).Execute()

List documents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListDocuments(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDocuments`: DocumentsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** |  | 

### Return type

[**DocumentsList**](DocumentsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIdPhotos

> IdPhotosList ListIdPhotos(ctx).ApplicantId(applicantId).Execute()

List ID photos



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the applicant the ID photos belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListIdPhotos(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListIdPhotos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIdPhotos`: IdPhotosList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListIdPhotos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIdPhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the ID photos belong to. | 

### Return type

[**IdPhotosList**](IdPhotosList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLivePhotos

> LivePhotosList ListLivePhotos(ctx).ApplicantId(applicantId).Execute()

List live photos



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the applicant the live photos belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListLivePhotos(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListLivePhotos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLivePhotos`: LivePhotosList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListLivePhotos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLivePhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the live photos belong to. | 

### Return type

[**LivePhotosList**](LivePhotosList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLiveVideos

> LiveVideosList ListLiveVideos(ctx).ApplicantId(applicantId).Execute()

List live videos



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the applicant the live videos belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListLiveVideos(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListLiveVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLiveVideos`: LiveVideosList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListLiveVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLiveVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the live videos belong to. | 

### Return type

[**LiveVideosList**](LiveVideosList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMotionCaptures

> MotionCapturesList ListMotionCaptures(ctx).ApplicantId(applicantId).Execute()

List motion captures



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the applicant the motion captures belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListMotionCaptures(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListMotionCaptures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMotionCaptures`: MotionCapturesList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListMotionCaptures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMotionCapturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the motion captures belong to. | 

### Return type

[**MotionCapturesList**](MotionCapturesList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepeatAttempts

> RepeatAttemptsList ListRepeatAttempts(ctx, reportId).Execute()

Retrieve repeat attempts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListRepeatAttempts(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListRepeatAttempts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRepeatAttempts`: RepeatAttemptsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListRepeatAttempts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRepeatAttemptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepeatAttemptsList**](RepeatAttemptsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReports

> ReportsList ListReports(ctx).CheckId(checkId).Execute()

List reports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	checkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListReports(context.Background()).CheckId(checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReports`: ReportsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkId** | **string** |  | 

### Return type

[**ReportsList**](ReportsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> []TaskItem ListTasks(ctx, workflowRunId).Execute()

List Tasks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	workflowRunId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the Workflow Run to which the Tasks belong.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListTasks(context.Background(), workflowRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasks`: []TaskItem
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run to which the Tasks belong. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TaskItem**](TaskItem.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWatchlistMonitorMatches

> WatchlistMonitorMatchesList ListWatchlistMonitorMatches(ctx, monitorId).Execute()

List matches (BETA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	monitorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWatchlistMonitorMatches(context.Background(), monitorId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWatchlistMonitorMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWatchlistMonitorMatches`: WatchlistMonitorMatchesList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWatchlistMonitorMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWatchlistMonitorMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WatchlistMonitorMatchesList**](WatchlistMonitorMatchesList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWatchlistMonitors

> WatchlistMonitorsList ListWatchlistMonitors(ctx).ApplicantId(applicantId).IncludeDeleted(includeDeleted).Execute()

List monitors



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the applicant the watchlist monitors belong to. If omitted, all monitors for the account will be listed.
	includeDeleted := true // bool | Whether to also include deleted (inactive) monitors. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWatchlistMonitors(context.Background()).ApplicantId(applicantId).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWatchlistMonitors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWatchlistMonitors`: WatchlistMonitorsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWatchlistMonitors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWatchlistMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the watchlist monitors belong to. If omitted, all monitors for the account will be listed. | 
 **includeDeleted** | **bool** | Whether to also include deleted (inactive) monitors. | [default to false]

### Return type

[**WatchlistMonitorsList**](WatchlistMonitorsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> WebhooksList ListWebhooks(ctx).Execute()

List webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooks`: WebhooksList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


### Return type

[**WebhooksList**](WebhooksList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowRuns

> []WorkflowRun ListWorkflowRuns(ctx).Page(page).Status(status).CreatedAtGt(createdAtGt).CreatedAtLt(createdAtLt).Sort(sort).Execute()

List Workflow Runs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	page := int32(56) // int32 | The number of the page to be retrieved. If not specified, defaults to 1. (optional) (default to 1)
	status := "status_example" // string | A list of comma separated status values to filter the results. Possible values are 'processing', 'awaiting_input', 'approved', 'declined', 'review', 'abandoned' and 'error'. (optional)
	createdAtGt := time.Now() // time.Time | A ISO-8601 date to filter results with a created date greater than (after) the one provided. (optional)
	createdAtLt := time.Now() // time.Time | A ISO-8601 date to filter results with a created date less than (before) the one provided. (optional)
	sort := "sort_example" // string | A string with the value 'desc' or 'asc' that allows to sort the returned list by the completed datetime either descending or ascending, respectively. If not specified, defaults to 'desc'. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWorkflowRuns(context.Background()).Page(page).Status(status).CreatedAtGt(createdAtGt).CreatedAtLt(createdAtLt).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWorkflowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflowRuns`: []WorkflowRun
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWorkflowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The number of the page to be retrieved. If not specified, defaults to 1. | [default to 1]
 **status** | **string** | A list of comma separated status values to filter the results. Possible values are &#39;processing&#39;, &#39;awaiting_input&#39;, &#39;approved&#39;, &#39;declined&#39;, &#39;review&#39;, &#39;abandoned&#39; and &#39;error&#39;. | 
 **createdAtGt** | **time.Time** | A ISO-8601 date to filter results with a created date greater than (after) the one provided. | 
 **createdAtLt** | **time.Time** | A ISO-8601 date to filter results with a created date less than (before) the one provided. | 
 **sort** | **string** | A string with the value &#39;desc&#39; or &#39;asc&#39; that allows to sort the returned list by the completed datetime either descending or ascending, respectively. If not specified, defaults to &#39;desc&#39;. | [default to &quot;desc&quot;]

### Return type

[**[]WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> string Ping(ctx).Execute()

Ping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Ping`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

**string**

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostResultsFeedback

> ResultsFeedback PostResultsFeedback(ctx).ResultsFeedback(resultsFeedback).Execute()

Fraud reporting (ALPHA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	resultsFeedback := *openapiclient.NewResultsFeedback() // ResultsFeedback | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.PostResultsFeedback(context.Background()).ResultsFeedback(resultsFeedback).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostResultsFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResultsFeedback`: ResultsFeedback
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostResultsFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostResultsFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resultsFeedback** | [**ResultsFeedback**](ResultsFeedback.md) |  | 

### Return type

[**ResultsFeedback**](ResultsFeedback.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendWebhooks

> ResendWebhooks(ctx).WebhookResend(webhookResend).Execute()

Resends webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	webhookResend := *openapiclient.NewWebhookResend() // WebhookResend | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResendWebhooks(context.Background()).WebhookResend(webhookResend).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResendWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookResend** | [**WebhookResend**](WebhookResend.md) |  | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreApplicant

> RestoreApplicant(ctx, applicantId).Execute()

Restore Applicant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RestoreApplicant(context.Background(), applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RestoreApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeCheck

> ResumeCheck(ctx, checkId).Execute()

Resume a Check



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	checkId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResumeCheck(context.Background(), checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResumeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeReport

> ResumeReport(ctx, reportId).Execute()

Resume report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResumeReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResumeReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicant

> Applicant UpdateApplicant(ctx, applicantId).ApplicantUpdater(applicantUpdater).Execute()

Update Applicant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	applicantUpdater := *openapiclient.NewApplicantUpdater() // ApplicantUpdater | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApplicant(context.Background(), applicantId).ApplicantUpdater(applicantUpdater).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicant`: Applicant
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApplicant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicantUpdater** | [**ApplicantUpdater**](ApplicantUpdater.md) |  | 

### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWatchlistMonitorMatch

> WatchlistMonitorMatchesList UpdateWatchlistMonitorMatch(ctx, monitorId).WatchlistMonitorMatchesUpdater(watchlistMonitorMatchesUpdater).Execute()

Set match status (BETA)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	monitorId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	watchlistMonitorMatchesUpdater := *openapiclient.NewWatchlistMonitorMatchesUpdater() // WatchlistMonitorMatchesUpdater | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateWatchlistMonitorMatch(context.Background(), monitorId).WatchlistMonitorMatchesUpdater(watchlistMonitorMatchesUpdater).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateWatchlistMonitorMatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWatchlistMonitorMatch`: WatchlistMonitorMatchesList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateWatchlistMonitorMatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWatchlistMonitorMatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **watchlistMonitorMatchesUpdater** | [**WatchlistMonitorMatchesUpdater**](WatchlistMonitorMatchesUpdater.md) |  | 

### Return type

[**WatchlistMonitorMatchesList**](WatchlistMonitorMatchesList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebhook

> Webhook UpdateWebhook(ctx, webhookId).WebhookUpdater(webhookUpdater).Execute()

Edit a webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	webhookUpdater := *openapiclient.NewWebhookUpdater() // WebhookUpdater | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateWebhook(context.Background(), webhookId).WebhookUpdater(webhookUpdater).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhookUpdater** | [**WebhookUpdater**](WebhookUpdater.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadDocument

> Document UploadDocument(ctx).Type_(type_).ApplicantId(applicantId).File(file).FileType(fileType).Side(side).IssuingCountry(issuingCountry).ValidateImageQuality(validateImageQuality).Location(location).Execute()

Upload a document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	type_ := "type__example" // string | The type of document
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the applicant whose document is being uploaded.
	file := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded.
	fileType := "fileType_example" // string | The file type of the uploaded file (optional)
	side := "side_example" // string | The side of the document, if applicable. The possible values are front and back (optional)
	issuingCountry := "issuingCountry_example" // CountryCodes | The issuing country of the document, a 3-letter ISO code. (optional)
	validateImageQuality := true // bool | Defaults to false. When true the submitted image will undergo an image quality validation which may take up to 5 seconds. (optional)
	location := *openapiclient.NewLocationBuilder() // LocationBuilder |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UploadDocument(context.Background()).Type_(type_).ApplicantId(applicantId).File(file).FileType(fileType).Side(side).IssuingCountry(issuingCountry).ValidateImageQuality(validateImageQuality).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UploadDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadDocument`: Document
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UploadDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | The type of document | 
 **applicantId** | **string** | The ID of the applicant whose document is being uploaded. | 
 **file** | ***os.File** | The file to be uploaded. | 
 **fileType** | **string** | The file type of the uploaded file | 
 **side** | **string** | The side of the document, if applicable. The possible values are front and back | 
 **issuingCountry** | **CountryCodes** | The issuing country of the document, a 3-letter ISO code. | 
 **validateImageQuality** | **bool** | Defaults to false. When true the submitted image will undergo an image quality validation which may take up to 5 seconds. | 
 **location** | [**LocationBuilder**](LocationBuilder.md) |  | 

### Return type

[**Document**](Document.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadIdPhoto

> IdPhoto UploadIdPhoto(ctx).ApplicantId(applicantId).File(file).Execute()

Upload ID photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the applicant whose ID photo is being uploaded. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UploadIdPhoto(context.Background()).ApplicantId(applicantId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UploadIdPhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadIdPhoto`: IdPhoto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UploadIdPhoto`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadIdPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The ID of the applicant whose ID photo is being uploaded. | 
 **file** | ***os.File** | The file to be uploaded. | 

### Return type

[**IdPhoto**](IdPhoto.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLivePhoto

> LivePhoto UploadLivePhoto(ctx).ApplicantId(applicantId).File(file).AdvancedValidation(advancedValidation).Execute()

Upload live photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dhia-gharsallaoui/go-onfido"
)

func main() {
	applicantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the applicant whose live photo is being uploaded. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded. (optional)
	advancedValidation := true // bool | Validates that the live photo contains exactly one face. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UploadLivePhoto(context.Background()).ApplicantId(applicantId).File(file).AdvancedValidation(advancedValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UploadLivePhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLivePhoto`: LivePhoto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UploadLivePhoto`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The ID of the applicant whose live photo is being uploaded. | 
 **file** | ***os.File** | The file to be uploaded. | 
 **advancedValidation** | **bool** | Validates that the live photo contains exactly one face. | [default to true]

### Return type

[**LivePhoto**](LivePhoto.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

