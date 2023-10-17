# \NetworkAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkExchangeRate**](NetworkAPI.md#GetNetworkExchangeRate) | **Get** /api/v1/network/exchangerate | Get the network exchange rate to estimate costs
[**GetNetworkFees**](NetworkAPI.md#GetNetworkFees) | **Get** /api/v1/network/fees | Get the network fees
[**GetNetworkNodes**](NetworkAPI.md#GetNetworkNodes) | **Get** /api/v1/network/nodes | Get the network address book nodes
[**GetNetworkStake**](NetworkAPI.md#GetNetworkStake) | **Get** /api/v1/network/stake | Get network stake information
[**GetNetworkSupply**](NetworkAPI.md#GetNetworkSupply) | **Get** /api/v1/network/supply | Get the network supply



## GetNetworkExchangeRate

> NetworkExchangeRateSetResponse GetNetworkExchangeRate(ctx).Timestamp(timestamp).Execute()

Get the network exchange rate to estimate costs



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
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkAPI.GetNetworkExchangeRate(context.Background()).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkExchangeRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkExchangeRate`: NetworkExchangeRateSetResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkExchangeRate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkExchangeRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**NetworkExchangeRateSetResponse**](NetworkExchangeRateSetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFees

> NetworkFeesResponse GetNetworkFees(ctx).Order(order).Timestamp(timestamp).Execute()

Get the network fees



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
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkAPI.GetNetworkFees(context.Background()).Order(order).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkFees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFees`: NetworkFeesResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkFees`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**NetworkFeesResponse**](NetworkFeesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkNodes

> NetworkNodesResponse GetNetworkNodes(ctx).FileId(fileId).Limit(limit).NodeId(nodeId).Order(order).Execute()

Get the network address book nodes



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
    fileId := "fileId_example" // string | The ID of the file entity (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    nodeId := "nodeId_example" // string | The ID of the node (optional)
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkAPI.GetNetworkNodes(context.Background()).FileId(fileId).Limit(limit).NodeId(nodeId).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkNodes`: NetworkNodesResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkNodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileId** | **string** | The ID of the file entity | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **nodeId** | **string** | The ID of the node | 
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]

### Return type

[**NetworkNodesResponse**](NetworkNodesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkStake

> NetworkStakeResponse GetNetworkStake(ctx).Execute()

Get network stake information



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkAPI.GetNetworkStake(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkStake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkStake`: NetworkStakeResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkStake`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkStakeRequest struct via the builder pattern


### Return type

[**NetworkStakeResponse**](NetworkStakeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSupply

> NetworkSupplyResponse GetNetworkSupply(ctx).Timestamp(timestamp).Execute()

Get the network supply



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
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkAPI.GetNetworkSupply(context.Background()).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkAPI.GetNetworkSupply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSupply`: NetworkSupplyResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkAPI.GetNetworkSupply`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSupplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**NetworkSupplyResponse**](NetworkSupplyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

