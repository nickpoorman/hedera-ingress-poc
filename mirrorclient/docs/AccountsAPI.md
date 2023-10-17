# \AccountsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountByIdOrAliasOrEvmAddress**](AccountsAPI.md#GetAccountByIdOrAliasOrEvmAddress) | **Get** /api/v1/accounts/{idOrAliasOrEvmAddress} | Get account by alias, id, or evm address
[**ListAccounts**](AccountsAPI.md#ListAccounts) | **Get** /api/v1/accounts | List account entities on network
[**ListCryptoAllowancesByAccountId**](AccountsAPI.md#ListCryptoAllowancesByAccountId) | **Get** /api/v1/accounts/{idOrAliasOrEvmAddress}/allowances/crypto | Get crypto allowances for an account info
[**ListNftByAccountId**](AccountsAPI.md#ListNftByAccountId) | **Get** /api/v1/accounts/{idOrAliasOrEvmAddress}/nfts | Get nfts for an account info
[**ListStakingRewardsByAccountId**](AccountsAPI.md#ListStakingRewardsByAccountId) | **Get** /api/v1/accounts/{idOrAliasOrEvmAddress}/rewards | Get past staking reward payouts for an account
[**ListTokenAllowancesByAccountId**](AccountsAPI.md#ListTokenAllowancesByAccountId) | **Get** /api/v1/accounts/{idOrAliasOrEvmAddress}/allowances/tokens | Get fungible token allowances for an account
[**ListTokenRelationshipByAccountId**](AccountsAPI.md#ListTokenRelationshipByAccountId) | **Get** /api/v1/accounts/{idOrAliasOrEvmAddress}/tokens | Get token relationships info for an account



## GetAccountByIdOrAliasOrEvmAddress

> AccountBalanceTransactions GetAccountByIdOrAliasOrEvmAddress(ctx, idOrAliasOrEvmAddress).Limit(limit).Order(order).Timestamp(timestamp).Transactiontype(transactiontype).Transactions(transactions).Execute()

Get account by alias, id, or evm address



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
    idOrAliasOrEvmAddress := "HIQQEXWKW53RKN4W6XXC4Q232SYNZ3SZANVZZSUME5B5PRGXL663UAQA" // string | Account alias or account id or evm address
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)
    transactiontype := openapiclient.TransactionTypes("CONSENSUSCREATETOPIC") // TransactionTypes |  (optional)
    transactions := true // bool | If provided and set to false transactions will not be included in the response (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.GetAccountByIdOrAliasOrEvmAddress(context.Background(), idOrAliasOrEvmAddress).Limit(limit).Order(order).Timestamp(timestamp).Transactiontype(transactiontype).Transactions(transactions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountByIdOrAliasOrEvmAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountByIdOrAliasOrEvmAddress`: AccountBalanceTransactions
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountByIdOrAliasOrEvmAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrAliasOrEvmAddress** | **string** | Account alias or account id or evm address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountByIdOrAliasOrEvmAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 
 **transactiontype** | [**TransactionTypes**](TransactionTypes.md) |  | 
 **transactions** | **bool** | If provided and set to false transactions will not be included in the response | [default to true]

### Return type

[**AccountBalanceTransactions**](AccountBalanceTransactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> AccountsResponse ListAccounts(ctx).AccountBalance(accountBalance).AccountId(accountId).AccountPublickey(accountPublickey).Balance(balance).Limit(limit).Order(order).Execute()

List account entities on network



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
    accountBalance := "accountBalance_example" // string | The optional balance value to compare against (optional)
    accountId := "accountId_example" // string | The ID of the account to return information for (optional)
    accountPublickey := "3c3d546321ff6f63d701d2ec5c277095874e19f4a235bee1e6bb19258bf362be" // string | The account's public key to compare against (optional)
    balance := true // bool | Whether to include balance information or not. If included, token balances are limited to at most 50 per account as outlined in HIP-367. (optional) (default to true)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListAccounts(context.Background()).AccountBalance(accountBalance).AccountId(accountId).AccountPublickey(accountPublickey).Balance(balance).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: AccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountBalance** | **string** | The optional balance value to compare against | 
 **accountId** | **string** | The ID of the account to return information for | 
 **accountPublickey** | **string** | The account&#39;s public key to compare against | 
 **balance** | **bool** | Whether to include balance information or not. If included, token balances are limited to at most 50 per account as outlined in HIP-367. | [default to true]
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]

### Return type

[**AccountsResponse**](AccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCryptoAllowancesByAccountId

> CryptoAllowancesResponse ListCryptoAllowancesByAccountId(ctx, idOrAliasOrEvmAddress).Limit(limit).Order(order).SpenderId(spenderId).Execute()

Get crypto allowances for an account info



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
    idOrAliasOrEvmAddress := "HIQQEXWKW53RKN4W6XXC4Q232SYNZ3SZANVZZSUME5B5PRGXL663UAQA" // string | Account alias or account id or evm address
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    spenderId := "spenderId_example" // string | The ID of the spender to return information for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListCryptoAllowancesByAccountId(context.Background(), idOrAliasOrEvmAddress).Limit(limit).Order(order).SpenderId(spenderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListCryptoAllowancesByAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCryptoAllowancesByAccountId`: CryptoAllowancesResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListCryptoAllowancesByAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrAliasOrEvmAddress** | **string** | Account alias or account id or evm address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCryptoAllowancesByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **spenderId** | **string** | The ID of the spender to return information for | 

### Return type

[**CryptoAllowancesResponse**](CryptoAllowancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNftByAccountId

> Nfts ListNftByAccountId(ctx, idOrAliasOrEvmAddress).Limit(limit).Order(order).Serialnumber(serialnumber).SpenderId(spenderId).TokenId(tokenId).Execute()

Get nfts for an account info



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
    idOrAliasOrEvmAddress := "HIQQEXWKW53RKN4W6XXC4Q232SYNZ3SZANVZZSUME5B5PRGXL663UAQA" // string | Account alias or account id or evm address
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    serialnumber := "serialnumber_example" // string | The nft serial number (64 bit type). Requires a tokenId value also be populated. (optional)
    spenderId := "spenderId_example" // string | The ID of the spender to return information for (optional)
    tokenId := "tokenId_example" // string | The ID of the token to return information for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListNftByAccountId(context.Background(), idOrAliasOrEvmAddress).Limit(limit).Order(order).Serialnumber(serialnumber).SpenderId(spenderId).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListNftByAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNftByAccountId`: Nfts
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListNftByAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrAliasOrEvmAddress** | **string** | Account alias or account id or evm address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNftByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **serialnumber** | **string** | The nft serial number (64 bit type). Requires a tokenId value also be populated. | 
 **spenderId** | **string** | The ID of the spender to return information for | 
 **tokenId** | **string** | The ID of the token to return information for | 

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


## ListStakingRewardsByAccountId

> StakingRewardsResponse ListStakingRewardsByAccountId(ctx, idOrAliasOrEvmAddress).Limit(limit).Order(order).Timestamp(timestamp).Execute()

Get past staking reward payouts for an account



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
    idOrAliasOrEvmAddress := "HIQQEXWKW53RKN4W6XXC4Q232SYNZ3SZANVZZSUME5B5PRGXL663UAQA" // string | Account alias or account id or evm address
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListStakingRewardsByAccountId(context.Background(), idOrAliasOrEvmAddress).Limit(limit).Order(order).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListStakingRewardsByAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStakingRewardsByAccountId`: StakingRewardsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListStakingRewardsByAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrAliasOrEvmAddress** | **string** | Account alias or account id or evm address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStakingRewardsByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**StakingRewardsResponse**](StakingRewardsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenAllowancesByAccountId

> TokenAllowancesResponse ListTokenAllowancesByAccountId(ctx, idOrAliasOrEvmAddress).Limit(limit).Order(order).SpenderId(spenderId).TokenId(tokenId).Execute()

Get fungible token allowances for an account



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
    idOrAliasOrEvmAddress := "HIQQEXWKW53RKN4W6XXC4Q232SYNZ3SZANVZZSUME5B5PRGXL663UAQA" // string | Account alias or account id or evm address
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")
    spenderId := "spenderId_example" // string | The ID of the spender to return information for (optional)
    tokenId := "tokenId_example" // string | The ID of the token to return information for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListTokenAllowancesByAccountId(context.Background(), idOrAliasOrEvmAddress).Limit(limit).Order(order).SpenderId(spenderId).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListTokenAllowancesByAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokenAllowancesByAccountId`: TokenAllowancesResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListTokenAllowancesByAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrAliasOrEvmAddress** | **string** | Account alias or account id or evm address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenAllowancesByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]
 **spenderId** | **string** | The ID of the spender to return information for | 
 **tokenId** | **string** | The ID of the token to return information for | 

### Return type

[**TokenAllowancesResponse**](TokenAllowancesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokenRelationshipByAccountId

> TokenRelationshipResponse ListTokenRelationshipByAccountId(ctx, idOrAliasOrEvmAddress).Limit(limit).Order(order).TokenId(tokenId).Execute()

Get token relationships info for an account



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
    idOrAliasOrEvmAddress := "HIQQEXWKW53RKN4W6XXC4Q232SYNZ3SZANVZZSUME5B5PRGXL663UAQA" // string | Account alias or account id or evm address
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")
    tokenId := "tokenId_example" // string | The ID of the token to return information for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountsAPI.ListTokenRelationshipByAccountId(context.Background(), idOrAliasOrEvmAddress).Limit(limit).Order(order).TokenId(tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ListTokenRelationshipByAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokenRelationshipByAccountId`: TokenRelationshipResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.ListTokenRelationshipByAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrAliasOrEvmAddress** | **string** | Account alias or account id or evm address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTokenRelationshipByAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]
 **tokenId** | **string** | The ID of the token to return information for | 

### Return type

[**TokenRelationshipResponse**](TokenRelationshipResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

