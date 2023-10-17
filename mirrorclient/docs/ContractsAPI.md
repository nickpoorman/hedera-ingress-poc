# \ContractsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContractsCall**](ContractsAPI.md#ContractsCall) | **Post** /api/v1/contracts/call | Invoke a smart contract
[**GetContractActionsByTransactionIdOrHash**](ContractsAPI.md#GetContractActionsByTransactionIdOrHash) | **Get** /api/v1/contracts/results/{transactionIdOrHash}/actions | Get the contract actions from a contract on the network for a given transactionId or ethereum transaction hash
[**GetContractById**](ContractsAPI.md#GetContractById) | **Get** /api/v1/contracts/{contractIdOrAddress} | Get contract by id
[**GetContractResultByIdAndTimestamp**](ContractsAPI.md#GetContractResultByIdAndTimestamp) | **Get** /api/v1/contracts/{contractIdOrAddress}/results/{timestamp} | Get the contract result from a contract on the network executed at a given timestamp
[**GetContractResultByTransactionIdOrHash**](ContractsAPI.md#GetContractResultByTransactionIdOrHash) | **Get** /api/v1/contracts/results/{transactionIdOrHash} | Get the contract result from a contract on the network for a given transactionId or ethereum transaction hash
[**ListAllContractsResults**](ContractsAPI.md#ListAllContractsResults) | **Get** /api/v1/contracts/results | List contract results from all contracts on the network
[**ListContractLogs**](ContractsAPI.md#ListContractLogs) | **Get** /api/v1/contracts/{contractIdOrAddress}/results/logs | List contract logs from a contract on the network
[**ListContractResults**](ContractsAPI.md#ListContractResults) | **Get** /api/v1/contracts/{contractIdOrAddress}/results | List contract results from a contract on the network
[**ListContractState**](ContractsAPI.md#ListContractState) | **Get** /api/v1/contracts/{contractIdOrAddress}/state | The contract state from a contract on the network
[**ListContracts**](ContractsAPI.md#ListContracts) | **Get** /api/v1/contracts | List contract entities on network
[**ListContractsLogs**](ContractsAPI.md#ListContractsLogs) | **Get** /api/v1/contracts/results/logs | List contracts logs across many contracts on the network



## ContractsCall

> ContractCallResponse ContractsCall(ctx).ContractCallRequest(contractCallRequest).Execute()

Invoke a smart contract



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
    contractCallRequest := *openapiclient.NewContractCallRequest("TODO") // ContractCallRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.ContractsCall(context.Background()).ContractCallRequest(contractCallRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.ContractsCall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ContractsCall`: ContractCallResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.ContractsCall`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContractsCallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractCallRequest** | [**ContractCallRequest**](ContractCallRequest.md) |  | 

### Return type

[**ContractCallResponse**](ContractCallResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractActionsByTransactionIdOrHash

> ContractActionsResponse GetContractActionsByTransactionIdOrHash(ctx, transactionIdOrHash).Index(index).Limit(limit).Order(order).Execute()

Get the contract actions from a contract on the network for a given transactionId or ethereum transaction hash



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
    transactionIdOrHash := "0.0.10-1234567890-000000000" // string | Transaction Id or a 32 byte hash with optional 0x prefix
    index := "index_example" // string | The index of a contract action (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.GetContractActionsByTransactionIdOrHash(context.Background(), transactionIdOrHash).Index(index).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractActionsByTransactionIdOrHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractActionsByTransactionIdOrHash`: ContractActionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractActionsByTransactionIdOrHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionIdOrHash** | **string** | Transaction Id or a 32 byte hash with optional 0x prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractActionsByTransactionIdOrHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **index** | **string** | The index of a contract action | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]

### Return type

[**ContractActionsResponse**](ContractActionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractById

> ContractResponse GetContractById(ctx, contractIdOrAddress).Timestamp(timestamp).Execute()

Get contract by id



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
    contractIdOrAddress := "contractIdOrAddress_example" // string | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract.
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.GetContractById(context.Background(), contractIdOrAddress).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractById`: ContractResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractIdOrAddress** | **string** | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**ContractResponse**](ContractResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractResultByIdAndTimestamp

> ContractResultDetails GetContractResultByIdAndTimestamp(ctx, contractIdOrAddress, timestamp).Execute()

Get the contract result from a contract on the network executed at a given timestamp



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
    contractIdOrAddress := "contractIdOrAddress_example" // string | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract.
    timestamp := "1234567890.0000007" // string | The timestamp at which the associated transaction reached consensus

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.GetContractResultByIdAndTimestamp(context.Background(), contractIdOrAddress, timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractResultByIdAndTimestamp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractResultByIdAndTimestamp`: ContractResultDetails
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractResultByIdAndTimestamp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractIdOrAddress** | **string** | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract. | 
**timestamp** | **string** | The timestamp at which the associated transaction reached consensus | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractResultByIdAndTimestampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ContractResultDetails**](ContractResultDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContractResultByTransactionIdOrHash

> ContractResultDetails GetContractResultByTransactionIdOrHash(ctx, transactionIdOrHash).Nonce(nonce).Execute()

Get the contract result from a contract on the network for a given transactionId or ethereum transaction hash



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
    transactionIdOrHash := "0.0.10-1234567890-000000000" // string | Transaction Id or a 32 byte hash with optional 0x prefix
    nonce := int32(1) // int32 | Filter the query result by the nonce of the transaction. A zero nonce represents user submitted transactions while a non-zero nonce is generated by main nodes. The filter honors the last value. Default is 0 when not specified. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.GetContractResultByTransactionIdOrHash(context.Background(), transactionIdOrHash).Nonce(nonce).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.GetContractResultByTransactionIdOrHash``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContractResultByTransactionIdOrHash`: ContractResultDetails
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.GetContractResultByTransactionIdOrHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionIdOrHash** | **string** | Transaction Id or a 32 byte hash with optional 0x prefix | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContractResultByTransactionIdOrHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nonce** | **int32** | Filter the query result by the nonce of the transaction. A zero nonce represents user submitted transactions while a non-zero nonce is generated by main nodes. The filter honors the last value. Default is 0 when not specified. | [default to 0]

### Return type

[**ContractResultDetails**](ContractResultDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllContractsResults

> ContractResultsResponse ListAllContractsResults(ctx).From(from).BlockHash(blockHash).BlockNumber(blockNumber).Internal(internal).Limit(limit).Order(order).Timestamp(timestamp).TransactionIndex(transactionIndex).Execute()

List contract results from all contracts on the network



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
    from := "from_example" // string | Account ID or EVM address executing the contract (optional)
    blockHash := "blockHash_example" // string | The block's hash (optional)
    blockNumber := "blockNumber_example" // string | The block's number (optional)
    internal := true // bool | Whether to include child transactions or not (optional) (default to false)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)
    transactionIndex := int32(1) // int32 | The transaction index in the block (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.ListAllContractsResults(context.Background()).From(from).BlockHash(blockHash).BlockNumber(blockNumber).Internal(internal).Limit(limit).Order(order).Timestamp(timestamp).TransactionIndex(transactionIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.ListAllContractsResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllContractsResults`: ContractResultsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.ListAllContractsResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllContractsResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | Account ID or EVM address executing the contract | 
 **blockHash** | **string** | The block&#39;s hash | 
 **blockNumber** | **string** | The block&#39;s number | 
 **internal** | **bool** | Whether to include child transactions or not | [default to false]
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 
 **transactionIndex** | **int32** | The transaction index in the block | 

### Return type

[**ContractResultsResponse**](ContractResultsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContractLogs

> ContractLogsResponse ListContractLogs(ctx, contractIdOrAddress).Index(index).Limit(limit).Order(order).Timestamp(timestamp).Topic0(topic0).Topic1(topic1).Topic2(topic2).Topic3(topic3).Execute()

List contract logs from a contract on the network



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
    contractIdOrAddress := "contractIdOrAddress_example" // string | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract.
    index := "index_example" // string | Contract log index (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)
    topic0 := []string{"Inner_example"} // []string | The first topic associated with a contract log. Requires a timestamp range also be populated. (optional)
    topic1 := []string{"Inner_example"} // []string | The second topic associated with a contract log. Requires a timestamp range also be populated. (optional)
    topic2 := []string{"Inner_example"} // []string | The third topic associated with a contract log. Requires a timestamp range also be populated. (optional)
    topic3 := []string{"Inner_example"} // []string | The fourth topic associated with a contract log. Requires a timestamp range also be populated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.ListContractLogs(context.Background(), contractIdOrAddress).Index(index).Limit(limit).Order(order).Timestamp(timestamp).Topic0(topic0).Topic1(topic1).Topic2(topic2).Topic3(topic3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.ListContractLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContractLogs`: ContractLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.ListContractLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractIdOrAddress** | **string** | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContractLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **index** | **string** | Contract log index | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 
 **topic0** | **[]string** | The first topic associated with a contract log. Requires a timestamp range also be populated. | 
 **topic1** | **[]string** | The second topic associated with a contract log. Requires a timestamp range also be populated. | 
 **topic2** | **[]string** | The third topic associated with a contract log. Requires a timestamp range also be populated. | 
 **topic3** | **[]string** | The fourth topic associated with a contract log. Requires a timestamp range also be populated. | 

### Return type

[**ContractLogsResponse**](ContractLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContractResults

> ContractResultsResponse ListContractResults(ctx, contractIdOrAddress).BlockHash(blockHash).BlockNumber(blockNumber).From(from).Internal(internal).Limit(limit).Order(order).Timestamp(timestamp).TransactionIndex(transactionIndex).Execute()

List contract results from a contract on the network



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
    contractIdOrAddress := "contractIdOrAddress_example" // string | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract.
    blockHash := "blockHash_example" // string | The block's hash (optional)
    blockNumber := "blockNumber_example" // string | The block's number (optional)
    from := "from_example" // string | Account ID or EVM address executing the contract (optional)
    internal := true // bool | Whether to include child transactions or not (optional) (default to false)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)
    transactionIndex := int32(1) // int32 | The transaction index in the block (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.ListContractResults(context.Background(), contractIdOrAddress).BlockHash(blockHash).BlockNumber(blockNumber).From(from).Internal(internal).Limit(limit).Order(order).Timestamp(timestamp).TransactionIndex(transactionIndex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.ListContractResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContractResults`: ContractResultsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.ListContractResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractIdOrAddress** | **string** | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContractResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockHash** | **string** | The block&#39;s hash | 
 **blockNumber** | **string** | The block&#39;s number | 
 **from** | **string** | Account ID or EVM address executing the contract | 
 **internal** | **bool** | Whether to include child transactions or not | [default to false]
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 
 **transactionIndex** | **int32** | The transaction index in the block | 

### Return type

[**ContractResultsResponse**](ContractResultsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContractState

> ContractStateResponse ListContractState(ctx, contractIdOrAddress).Limit(limit).Order(order).Slot(slot).Timestamp(timestamp).Execute()

The contract state from a contract on the network



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
    contractIdOrAddress := "contractIdOrAddress_example" // string | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract.
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "desc" // string | The order in which items are listed (optional) (default to "asc")
    slot := "slot_example" // string | The slot's number (optional)
    timestamp := "1234567890.0000007" // string | The timestamp at which the contract state is (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.ListContractState(context.Background(), contractIdOrAddress).Limit(limit).Order(order).Slot(slot).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.ListContractState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContractState`: ContractStateResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.ListContractState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractIdOrAddress** | **string** | The ID or hex encoded EVM address (with or without 0x prefix) associated with this contract. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListContractStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]
 **slot** | **string** | The slot&#39;s number | 
 **timestamp** | **string** | The timestamp at which the contract state is | 

### Return type

[**ContractStateResponse**](ContractStateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContracts

> ContractsResponse ListContracts(ctx).ContractId(contractId).Limit(limit).Order(order).Execute()

List contract entities on network



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
    contractId := "contractId_example" // string | The ID of the smart contract (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.ListContracts(context.Background()).ContractId(contractId).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.ListContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContracts`: ContractsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.ListContracts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContractsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contractId** | **string** | The ID of the smart contract | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]

### Return type

[**ContractsResponse**](ContractsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListContractsLogs

> ContractLogsResponse ListContractsLogs(ctx).Index(index).Limit(limit).Order(order).Timestamp(timestamp).Topic0(topic0).Topic1(topic1).Topic2(topic2).Topic3(topic3).Execute()

List contracts logs across many contracts on the network



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
    index := "index_example" // string | Contract log index (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)
    topic0 := []string{"Inner_example"} // []string | The first topic associated with a contract log. Requires a timestamp range also be populated. (optional)
    topic1 := []string{"Inner_example"} // []string | The second topic associated with a contract log. Requires a timestamp range also be populated. (optional)
    topic2 := []string{"Inner_example"} // []string | The third topic associated with a contract log. Requires a timestamp range also be populated. (optional)
    topic3 := []string{"Inner_example"} // []string | The fourth topic associated with a contract log. Requires a timestamp range also be populated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ContractsAPI.ListContractsLogs(context.Background()).Index(index).Limit(limit).Order(order).Timestamp(timestamp).Topic0(topic0).Topic1(topic1).Topic2(topic2).Topic3(topic3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsAPI.ListContractsLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListContractsLogs`: ContractLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsAPI.ListContractsLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListContractsLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index** | **string** | Contract log index | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 
 **topic0** | **[]string** | The first topic associated with a contract log. Requires a timestamp range also be populated. | 
 **topic1** | **[]string** | The second topic associated with a contract log. Requires a timestamp range also be populated. | 
 **topic2** | **[]string** | The third topic associated with a contract log. Requires a timestamp range also be populated. | 
 **topic3** | **[]string** | The fourth topic associated with a contract log. Requires a timestamp range also be populated. | 

### Return type

[**ContractLogsResponse**](ContractLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

