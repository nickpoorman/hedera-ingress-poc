# \TokensAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenById**](TokensAPI.md#GetTokenById) | **Get** /api/v1/tokens/{tokenId} | Get token by id
[**ListNftBySerialnumber**](TokensAPI.md#ListNftBySerialnumber) | **Get** /api/v1/tokens/{tokenId}/nfts/{serialNumber} | Get nft info
[**ListNftTransactions**](TokensAPI.md#ListNftTransactions) | **Get** /api/v1/tokens/{tokenId}/nfts/{serialNumber}/transactions | Get an nfts transction history
[**ListNfts**](TokensAPI.md#ListNfts) | **Get** /api/v1/tokens/{tokenId}/nfts | List nfts
[**ListTokenBalancesById**](TokensAPI.md#ListTokenBalancesById) | **Get** /api/v1/tokens/{tokenId}/balances | List token balances
[**ListTokens**](TokensAPI.md#ListTokens) | **Get** /api/v1/tokens | List tokens



## GetTokenById

> TokenInfo GetTokenById(ctx, tokenId).Timestamp(timestamp).Execute()

Get token by id



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
    tokenId := "tokenId_example" // string | Token id
    timestamp := "timestamp_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensAPI.GetTokenById(context.Background(), tokenId).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokenById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenById`: TokenInfo
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokenById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timestamp** | **string** |  | 

### Return type

[**TokenInfo**](TokenInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNftBySerialnumber

> Nft ListNftBySerialnumber(ctx, tokenId, serialNumber).Execute()

Get nft info



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
    tokenId := "tokenId_example" // string | Token id
    serialNumber := int64(1) // int64 | The nft serial number (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensAPI.ListNftBySerialnumber(context.Background(), tokenId, serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ListNftBySerialnumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNftBySerialnumber`: Nft
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.ListNftBySerialnumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token id | 
**serialNumber** | **int64** | The nft serial number | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiListNftBySerialnumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Nft**](Nft.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNftTransactions

> NftTransactionHistory ListNftTransactions(ctx, tokenId, serialNumber).Limit(limit).Order(order).Timestamp(timestamp).Execute()

Get an nfts transction history



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
    tokenId := "tokenId_example" // string | Token id
    serialNumber := int64(1) // int64 | The nft serial number (default to 1)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensAPI.ListNftTransactions(context.Background(), tokenId, serialNumber).Limit(limit).Order(order).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ListNftTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNftTransactions`: NftTransactionHistory
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.ListNftTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token id | 
**serialNumber** | **int64** | The nft serial number | [default to 1]

### Other Parameters

Other parameters are passed through a pointer to a apiListNftTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**NftTransactionHistory**](NftTransactionHistory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNfts

> Nfts ListNfts(ctx, tokenId).AccountId(accountId).Limit(limit).Order(order).Serialnumber(serialnumber).Execute()

List nfts



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
    tokenId := "tokenId_example" // string | Token id
    accountId := "accountId_example" // string | The ID of the account to return information for (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    serialnumber := "serialnumber_example" // string | The nft serial number (64 bit type). Requires a tokenId value also be populated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensAPI.ListNfts(context.Background(), tokenId).AccountId(accountId).Limit(limit).Order(order).Serialnumber(serialnumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ListNfts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNfts`: Nfts
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.ListNfts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **string** | The ID of the account to return information for | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **serialnumber** | **string** | The nft serial number (64 bit type). Requires a tokenId value also be populated. | 

### Return type

[**Nfts**](Nfts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenBalancesById

> TokenBalancesResponse ListTokenBalancesById(ctx, tokenId).AccountBalance(accountBalance).AccountId(accountId).AccountPublickey(accountPublickey).Limit(limit).Order(order).Timestamp(timestamp).Execute()

List token balances



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
    tokenId := "tokenId_example" // string | Token id
    accountBalance := "accountBalance_example" // string | The optional balance value to compare against (optional)
    accountId := "accountId_example" // string | The ID of the account to return information for (optional)
    accountPublickey := "3c3d546321ff6f63d701d2ec5c277095874e19f4a235bee1e6bb19258bf362be" // string | The account's public key to compare against (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensAPI.ListTokenBalancesById(context.Background(), tokenId).AccountBalance(accountBalance).AccountId(accountId).AccountPublickey(accountPublickey).Limit(limit).Order(order).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ListTokenBalancesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokenBalancesById`: TokenBalancesResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.ListTokenBalancesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenBalancesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountBalance** | **string** | The optional balance value to compare against | 
 **accountId** | **string** | The ID of the account to return information for | 
 **accountPublickey** | **string** | The account&#39;s public key to compare against | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**TokenBalancesResponse**](TokenBalancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> TokensResponse ListTokens(ctx).AccountId(accountId).Limit(limit).Order(order).Publickey(publickey).TokenId(tokenId).Type_(type_).Execute()

List tokens



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
    accountId := "accountId_example" // string | The ID of the account to return information for (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")
    publickey := "3c3d546321ff6f63d701d2ec5c277095874e19f4a235bee1e6bb19258bf362be" // string | The public key to compare against (optional)
    tokenId := "tokenId_example" // string | The ID of the token to return information for (optional)
    type_ := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensAPI.ListTokens(context.Background()).AccountId(accountId).Limit(limit).Order(order).Publickey(publickey).TokenId(tokenId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ListTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokens`: TokensResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensAPI.ListTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The ID of the account to return information for | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]
 **publickey** | **string** | The public key to compare against | 
 **tokenId** | **string** | The ID of the token to return information for | 
 **type_** | **[]string** |  | 

### Return type

[**TokensResponse**](TokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

