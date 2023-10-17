# \SchedulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScheduleById**](SchedulesAPI.md#GetScheduleById) | **Get** /api/v1/schedules/{scheduleId} | Get schedule by id
[**ListSchedules**](SchedulesAPI.md#ListSchedules) | **Get** /api/v1/schedules | List schedules entities



## GetScheduleById

> Schedule GetScheduleById(ctx, scheduleId).Execute()

Get schedule by id



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
    scheduleId := "scheduleId_example" // string | Schedule id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesAPI.GetScheduleById(context.Background(), scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.GetScheduleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScheduleById`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.GetScheduleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scheduleId** | **string** | Schedule id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Schedule**](Schedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedules

> SchedulesResponse ListSchedules(ctx).AccountId(accountId).Limit(limit).Order(order).ScheduleId(scheduleId).Execute()

List schedules entities



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
    scheduleId := "scheduleId_example" // string | The ID of the schedule to return information for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulesAPI.ListSchedules(context.Background()).AccountId(accountId).Limit(limit).Order(order).ScheduleId(scheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulesAPI.ListSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchedules`: SchedulesResponse
    fmt.Fprintf(os.Stdout, "Response from `SchedulesAPI.ListSchedules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The ID of the account to return information for | 
 **limit** | **int32** | The maximum number of items to return | [default to 25]
 **order** | **string** | The order in which items are listed | [default to &quot;asc&quot;]
 **scheduleId** | **string** | The ID of the schedule to return information for | 

### Return type

[**SchedulesResponse**](SchedulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

