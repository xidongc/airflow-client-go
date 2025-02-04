# \TaskInstanceApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExtraLinks**](TaskInstanceApi.md#GetExtraLinks) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/links | List extra links
[**GetLog**](TaskInstanceApi.md#GetLog) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/logs/{task_try_number} | Get logs
[**GetTaskInstance**](TaskInstanceApi.md#GetTaskInstance) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id} | Get a task instance
[**GetTaskInstances**](TaskInstanceApi.md#GetTaskInstances) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances | List task instances
[**GetTaskInstancesBatch**](TaskInstanceApi.md#GetTaskInstancesBatch) | **Post** /dags/~/dagRuns/~/taskInstances/list | List task instances (batch)



## GetExtraLinks

> ExtraLinkCollection GetExtraLinks(ctx, dagId, dagRunId, taskId).Execute()

List extra links



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dagId := "dagId_example" // string | The DAG ID.
    dagRunId := "dagRunId_example" // string | The DAG run ID.
    taskId := "taskId_example" // string | The task ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskInstanceApi.GetExtraLinks(context.Background(), dagId, dagRunId, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskInstanceApi.GetExtraLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtraLinks`: ExtraLinkCollection
    fmt.Fprintf(os.Stdout, "Response from `TaskInstanceApi.GetExtraLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 
**taskId** | **string** | The task ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtraLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ExtraLinkCollection**](ExtraLinkCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLog

> InlineResponse200 GetLog(ctx, dagId, dagRunId, taskId, taskTryNumber).FullContent(fullContent).Token(token).Execute()

Get logs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dagId := "dagId_example" // string | The DAG ID.
    dagRunId := "dagRunId_example" // string | The DAG run ID.
    taskId := "taskId_example" // string | The task ID.
    taskTryNumber := int32(56) // int32 | The task try number.
    fullContent := true // bool | A full content will be returned. By default, only the first fragment will be returned.  (optional)
    token := "token_example" // string | A token that allows you to continue fetching logs. If passed, it will specify the location from which the download should be continued.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskInstanceApi.GetLog(context.Background(), dagId, dagRunId, taskId, taskTryNumber).FullContent(fullContent).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskInstanceApi.GetLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLog`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `TaskInstanceApi.GetLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 
**taskId** | **string** | The task ID. | 
**taskTryNumber** | **int32** | The task try number. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fullContent** | **bool** | A full content will be returned. By default, only the first fragment will be returned.  | 
 **token** | **string** | A token that allows you to continue fetching logs. If passed, it will specify the location from which the download should be continued.  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskInstance

> TaskInstance GetTaskInstance(ctx, dagId, dagRunId, taskId).Execute()

Get a task instance

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dagId := "dagId_example" // string | The DAG ID.
    dagRunId := "dagRunId_example" // string | The DAG run ID.
    taskId := "taskId_example" // string | The task ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskInstanceApi.GetTaskInstance(context.Background(), dagId, dagRunId, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskInstanceApi.GetTaskInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskInstance`: TaskInstance
    fmt.Fprintf(os.Stdout, "Response from `TaskInstanceApi.GetTaskInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 
**taskId** | **string** | The task ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TaskInstance**](TaskInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskInstances

> TaskInstanceCollection GetTaskInstances(ctx, dagId, dagRunId).ExecutionDateGte(executionDateGte).ExecutionDateLte(executionDateLte).StartDateGte(startDateGte).StartDateLte(startDateLte).EndDateGte(endDateGte).EndDateLte(endDateLte).DurationGte(durationGte).DurationLte(durationLte).State(state).Pool(pool).Queue(queue).Limit(limit).Offset(offset).Execute()

List task instances



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    dagId := "dagId_example" // string | The DAG ID.
    dagRunId := "dagRunId_example" // string | The DAG run ID.
    executionDateGte := time.Now() // time.Time | Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period.  (optional)
    executionDateLte := time.Now() // time.Time | Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period.  (optional)
    startDateGte := time.Now() // time.Time | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  (optional)
    startDateLte := time.Now() // time.Time | Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  (optional)
    endDateGte := time.Now() // time.Time | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  (optional)
    endDateLte := time.Now() // time.Time | Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  (optional)
    durationGte := float32(8.14) // float32 | Returns objects greater than or equal to the specified values.  This can be combined with duration_lte parameter to receive only the selected period.  (optional)
    durationLte := float32(8.14) // float32 | Returns objects less than or equal to the specified values.  This can be combined with duration_gte parameter to receive only the selected range.  (optional)
    state := []string{"Inner_example"} // []string | The value can be repeated to retrieve multiple matching values (OR condition). (optional)
    pool := []string{"Inner_example"} // []string | The value can be repeated to retrieve multiple matching values (OR condition). (optional)
    queue := []string{"Inner_example"} // []string | The value can be repeated to retrieve multiple matching values (OR condition). (optional)
    limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskInstanceApi.GetTaskInstances(context.Background(), dagId, dagRunId).ExecutionDateGte(executionDateGte).ExecutionDateLte(executionDateLte).StartDateGte(startDateGte).StartDateLte(startDateLte).EndDateGte(endDateGte).EndDateLte(endDateLte).DurationGte(durationGte).DurationLte(durationLte).State(state).Pool(pool).Queue(queue).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskInstanceApi.GetTaskInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskInstances`: TaskInstanceCollection
    fmt.Fprintf(os.Stdout, "Response from `TaskInstanceApi.GetTaskInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dagId** | **string** | The DAG ID. | 
**dagRunId** | **string** | The DAG run ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **executionDateGte** | **time.Time** | Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period.  | 
 **executionDateLte** | **time.Time** | Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period.  | 
 **startDateGte** | **time.Time** | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  | 
 **startDateLte** | **time.Time** | Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  | 
 **endDateGte** | **time.Time** | Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period.  | 
 **endDateLte** | **time.Time** | Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period.  | 
 **durationGte** | **float32** | Returns objects greater than or equal to the specified values.  This can be combined with duration_lte parameter to receive only the selected period.  | 
 **durationLte** | **float32** | Returns objects less than or equal to the specified values.  This can be combined with duration_gte parameter to receive only the selected range.  | 
 **state** | **[]string** | The value can be repeated to retrieve multiple matching values (OR condition). | 
 **pool** | **[]string** | The value can be repeated to retrieve multiple matching values (OR condition). | 
 **queue** | **[]string** | The value can be repeated to retrieve multiple matching values (OR condition). | 
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 

### Return type

[**TaskInstanceCollection**](TaskInstanceCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskInstancesBatch

> TaskInstanceCollection GetTaskInstancesBatch(ctx).ListTaskInstanceForm(listTaskInstanceForm).Execute()

List task instances (batch)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    listTaskInstanceForm := *openapiclient.NewListTaskInstanceForm() // ListTaskInstanceForm | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaskInstanceApi.GetTaskInstancesBatch(context.Background()).ListTaskInstanceForm(listTaskInstanceForm).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskInstanceApi.GetTaskInstancesBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaskInstancesBatch`: TaskInstanceCollection
    fmt.Fprintf(os.Stdout, "Response from `TaskInstanceApi.GetTaskInstancesBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskInstancesBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listTaskInstanceForm** | [**ListTaskInstanceForm**](ListTaskInstanceForm.md) |  | 

### Return type

[**TaskInstanceCollection**](TaskInstanceCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

