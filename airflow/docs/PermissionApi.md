# \PermissionApi

All URIs are relative to *http://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPermissions**](PermissionApi.md#GetPermissions) | **Get** /permissions | List permissions



## GetPermissions

> ActionCollection GetPermissions(ctx).Limit(limit).Offset(offset).Execute()

List permissions

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PermissionApi.GetPermissions(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.GetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermissions`: ActionCollection
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.GetPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The numbers of items to return. | [default to 100]
 **offset** | **int32** | The number of items to skip before starting to collect the result set. | 

### Return type

[**ActionCollection**](ActionCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

