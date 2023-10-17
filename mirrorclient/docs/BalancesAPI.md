# \BalancesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAccountBalances**](BalancesAPI.md#ListAccountBalances) | **Get** /api/v1/balances | List account balances



## ListAccountBalances

> BalancesResponse ListAccountBalances(ctx).AccountId(accountId).AccountBalance(accountBalance).AccountPublickey(accountPublickey).Limit(limit).Order(order).Timestamp(timestamp).Execute()

List account balances



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
    accountId := "HIQQEXWKW53RKN4W6XXC4Q232SYNZ3SZANVZZSUME5B5PRGXL663UAQA" // string | Account id or account alias with no shard realm or evm address with no shard realm (optional)
    accountBalance := "accountBalance_example" // string | The optional balance value to compare against (optional)
    accountPublickey := "3c3d546321ff6f63d701d2ec5c277095874e19f4a235bee1e6bb19258bf362be" // string | The account's public key to compare against (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BalancesAPI.ListAccountBalances(context.Background()).AccountId(accountId).AccountBalance(accountBalance).AccountPublickey(accountPublickey).Limit(limit).Order(order).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BalancesAPI.ListAccountBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountBalances`: BalancesResponse
    fmt.Fprintf(os.Stdout, "Response from `BalancesAPI.ListAccountBalances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | Account id or account alias with no shard realm or evm address with no shard realm | 
 **accountBalance** | **string** | The optional balance value to compare against | 
 **accountPublickey** | **string** | The account&#39;s public key to compare against | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**BalancesResponse**](BalancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

