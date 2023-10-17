# \TopicsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTopicMessageByIdAndSequenceNumber**](TopicsAPI.md#GetTopicMessageByIdAndSequenceNumber) | **Get** /api/v1/topics/{topicId}/messages/{sequenceNumber} | Get topic message by id and sequence number
[**GetTopicMessagesByConsensusTimestamp**](TopicsAPI.md#GetTopicMessagesByConsensusTimestamp) | **Get** /api/v1/topics/messages/{timestamp} | Get topic message by consensusTimestamp
[**ListTopicMessagesById**](TopicsAPI.md#ListTopicMessagesById) | **Get** /api/v1/topics/{topicId}/messages | List topic messages by id



## GetTopicMessageByIdAndSequenceNumber

> TopicMessagesResponse GetTopicMessageByIdAndSequenceNumber(ctx, topicId, sequenceNumber).Execute()

Get topic message by id and sequence number



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
    topicId := "topicId_example" // string | Topic id
    sequenceNumber := int64(2) // int64 | Topic message sequence number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsAPI.GetTopicMessageByIdAndSequenceNumber(context.Background(), topicId, sequenceNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.GetTopicMessageByIdAndSequenceNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopicMessageByIdAndSequenceNumber`: TopicMessagesResponse
    fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.GetTopicMessageByIdAndSequenceNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicId** | **string** | Topic id | 
**sequenceNumber** | **int64** | Topic message sequence number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicMessageByIdAndSequenceNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TopicMessagesResponse**](TopicMessagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTopicMessagesByConsensusTimestamp

> TopicMessage GetTopicMessagesByConsensusTimestamp(ctx, timestamp).Execute()

Get topic message by consensusTimestamp



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
    timestamp := "1234567890.0000007" // string | The timestamp at which the associated transaction reached consensus

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsAPI.GetTopicMessagesByConsensusTimestamp(context.Background(), timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.GetTopicMessagesByConsensusTimestamp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTopicMessagesByConsensusTimestamp`: TopicMessage
    fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.GetTopicMessagesByConsensusTimestamp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timestamp** | **string** | The timestamp at which the associated transaction reached consensus | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTopicMessagesByConsensusTimestampRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TopicMessage**](TopicMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTopicMessagesById

> TopicMessagesResponse ListTopicMessagesById(ctx, topicId).Encoding(encoding).Limit(limit).Order(order).Sequencenumber(sequencenumber).Timestamp(timestamp).Execute()

List topic messages by id



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
    topicId := "topicId_example" // string | Topic id
    encoding := "base64" // string |  (optional)
    limit := int32(2) // int32 | The maximum number of items to return (optional) (default to 25)
    order := "asc" // string | The order in which items are listed (optional) (default to "desc")
    sequencenumber := int64(2) // int64 |  (optional)
    timestamp := []string{"Inner_example"} // []string | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsAPI.ListTopicMessagesById(context.Background(), topicId).Encoding(encoding).Limit(limit).Order(order).Sequencenumber(sequencenumber).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsAPI.ListTopicMessagesById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTopicMessagesById`: TopicMessagesResponse
    fmt.Fprintf(os.Stdout, "Response from `TopicsAPI.ListTopicMessagesById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topicId** | **string** | Topic id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTopicMessagesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **encoding** | **string** |  | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;desc&quot;]
 **sequencenumber** | **int64** |  | 
 **timestamp** | **[]string** | The consensus timestamp in seconds.nanoseconds format with an optional comparison operator | 

### Return type

[**TopicMessagesResponse**](TopicMessagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

