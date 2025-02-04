# \EventLogApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventLog**](EventLogApi.md#GetEventLog) | **Get** /eventLogs/{event_log_id} | Get a log entry
[**GetEventLogs**](EventLogApi.md#GetEventLogs) | **Get** /eventLogs | List log entries



## GetEventLog

> EventLog GetEventLog(ctx, eventLogId).Execute()

Get a log entry

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
    eventLogId := int32(56) // int32 | The event log ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventLogApi.GetEventLog(context.Background(), eventLogId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventLogApi.GetEventLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventLog`: EventLog
    fmt.Fprintf(os.Stdout, "Response from `EventLogApi.GetEventLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventLogId** | **int32** | The event log ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventLog**](EventLog.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventLogs

> EventLogCollection GetEventLogs(ctx).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()

List log entries



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
    limit := int32(56) // int32 | The numbers of items to return. (optional) (default to 100)
    offset := int32(56) // int32 | The number of items to skip before starting to collect the result set. (optional)
    orderBy := "orderBy_example" // string | The name of the field to order the results by. Prefix a field name with `-` to reverse the sort order.  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventLogApi.GetEventLogs(context.Background()).Limit(limit).Offset(offset).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventLogApi.GetEventLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventLogs`: EventLogCollection
    fmt.Fprintf(os.Stdout, "Response from `EventLogApi.GetEventLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 
 **orderBy** | **string** | The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  | 

### Return type

[**EventLogCollection**](EventLogCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

