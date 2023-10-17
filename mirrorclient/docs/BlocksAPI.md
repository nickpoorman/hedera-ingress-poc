# \BlocksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByHashOrNumber**](BlocksAPI.md#GetByHashOrNumber) | **Get** /api/v1/blocks/{hashOrNumber} | Get block by hash or number
[**ListBlocks**](BlocksAPI.md#ListBlocks) | **Get** /api/v1/blocks | List blocks



## GetByHashOrNumber

> Block GetByHashOrNumber(ctx, hashOrNumber).Execute()

Get block by hash or number



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nickpoorman/hedera-ingress/mirrorclient"
)

func main() {
    hashOrNumber := "hashOrNumber_example" // string | Accepts both eth and hedera hash format or block number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocksAPI.GetByHashOrNumber(context.Background(), hashOrNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.GetByHashOrNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByHashOrNumber`: Block
    fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.GetByHashOrNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hashOrNumber** | **string** | Accepts both eth and hedera hash format or block number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByHashOrNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Block**](Block.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlocks

> BlocksResponse ListBlocks(ctx).BlockNumber(blockNumber).Limit(limit).Order(order).Timestamp(timestamp).Execute()

List blocks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nickpoorman/hedera-ingress/mirrorclient"
)

func main() {
    blockNumber := "blockNumber_example" // string | The block's number (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlocksAPI.ListBlocks(context.Background()).BlockNumber(blockNumber).Limit(limit).Order(order).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlocksAPI.ListBlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBlocks`: BlocksResponse
    fmt.Fprintf(os.Stdout, "Response from `BlocksAPI.ListBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockNumber** | **string** | The block&#39;s number | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**BlocksResponse**](BlocksResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

