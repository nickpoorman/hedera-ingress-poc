# ContractLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]ContractLog**](ContractLog.md) |  | [optional] 

## Methods

### NewContractLogsResponse

`func NewContractLogsResponse() *ContractLogsResponse`

NewContractLogsResponse instantiates a new ContractLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractLogsResponseWithDefaults

`func NewContractLogsResponseWithDefaults() *ContractLogsResponse`

NewContractLogsResponseWithDefaults instantiates a new ContractLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *ContractLogsResponse) GetLogs() []ContractLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ContractLogsResponse) GetLogsOk() (*[]ContractLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ContractLogsResponse) SetLogs(v []ContractLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ContractLogsResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


