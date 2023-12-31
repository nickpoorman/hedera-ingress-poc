# \TransactionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStateproofInfo**](TransactionsAPI.md#GetStateproofInfo) | **Get** /api/v1/transactions/{transactionId}/stateproof | Get stateproof information
[**GetTransactionById**](TransactionsAPI.md#GetTransactionById) | **Get** /api/v1/transactions/{transactionId} | Get transaction by id
[**ListTransactions**](TransactionsAPI.md#ListTransactions) | **Get** /api/v1/transactions | List transactions



## GetStateproofInfo

> GetStateproofInfo200Response GetStateproofInfo(ctx, transactionId).Nonce(nonce).Scheduled(scheduled).Execute()

Get stateproof information



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
    transactionId := "0.0.10-1234567890-000000000" // string | Transaction id
    nonce := int32(1) // int32 | Filter the query result by the nonce of the transaction. A zero nonce represents user submitted transactions while a non-zero nonce is generated by main nodes. The filter honors the last value. Default is 0 when not specified. (optional) (default to 0)
    scheduled := true // bool | Filter transactions by the scheduled flag. If true, return information for the scheduled transaction. If false, return information for the non-scheduled transaction. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.GetStateproofInfo(context.Background(), transactionId).Nonce(nonce).Scheduled(scheduled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetStateproofInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStateproofInfo`: GetStateproofInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetStateproofInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStateproofInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonce** | **int32** | Filter the query result by the nonce of the transaction. A zero nonce represents user submitted transactions while a non-zero nonce is generated by main nodes. The filter honors the last value. Default is 0 when not specified. | [default to 0]
 **scheduled** | **bool** | Filter transactions by the scheduled flag. If true, return information for the scheduled transaction. If false, return information for the non-scheduled transaction. | [default to false]

### Return type

[**GetStateproofInfo200Response**](GetStateproofInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransactionById

> TransactionByIdResponse GetTransactionById(ctx, transactionId).Nonce(nonce).Scheduled(scheduled).Execute()

Get transaction by id



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
    transactionId := "0.0.10-1234567890-000000000" // string | Transaction id
    nonce := int32(0) // int32 | Filter the query result by the nonce of the transaction. A zero nonce represents user submitted transactions while a non-zero nonce is generated by main nodes. The filter honors the last value. If not specified, all transactions with specified payer account ID and valid start timestamp match. (optional)
    scheduled := true // bool | Filter transactions by the scheduled flag. If true, return information for the scheduled transaction. If false, return information for the non-scheduled transaction. If not present, return information for all transactions matching transactionId. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.GetTransactionById(context.Background(), transactionId).Nonce(nonce).Scheduled(scheduled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.GetTransactionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransactionById`: TransactionByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.GetTransactionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonce** | **int32** | Filter the query result by the nonce of the transaction. A zero nonce represents user submitted transactions while a non-zero nonce is generated by main nodes. The filter honors the last value. If not specified, all transactions with specified payer account ID and valid start timestamp match. | 
 **scheduled** | **bool** | Filter transactions by the scheduled flag. If true, return information for the scheduled transaction. If false, return information for the non-scheduled transaction. If not present, return information for all transactions matching transactionId. | 

### Return type

[**TransactionByIdResponse**](TransactionByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransactions

> TransactionsResponse ListTransactions(ctx).AccountId(accountId).Limit(limit).Order(order).Timestamp(timestamp).Transactiontype(transactiontype).Result(result).Type_(type_).Execute()

List transactions



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
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)
    transactiontype := openapiclient.TransactionTypes("CONSENSUSCREATETOPIC") // TransactionTypes |  (optional)
    result := "result_example" // string | The transaction success type. (optional)
    type_ := "type__example" // string | The transaction account balance modification type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransactionsAPI.ListTransactions(context.Background()).AccountId(accountId).Limit(limit).Order(order).Timestamp(timestamp).Transactiontype(transactiontype).Result(result).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.ListTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransactions`: TransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.ListTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The ID of the account to return information for | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 
 **transactiontype** | [**TransactionTypes**](TransactionTypes.md) |  | 
 **result** | **string** | The transaction success type. | 
 **type_** | **string** | The transaction account balance modification type. | 

### Return type

[**TransactionsResponse**](TransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

