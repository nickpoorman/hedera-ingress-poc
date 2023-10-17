# ContractLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The hex encoded EVM address of the contract | [optional] 
**Bloom** | Pointer to [***os.File**](*os.File.md) |  | [optional] 
**ContractId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Data** | Pointer to **NullableString** | The hex encoded data of the contract log | [optional] 
**Index** | Pointer to **int32** | The index of the contract log in the chain of logs for an execution | [optional] 
**Topics** | Pointer to **[]string** | A list of hex encoded topics associated with this log event | [optional] 
**BlockHash** | Pointer to **string** | The hex encoded block (record file chain) hash | [optional] 
**BlockNumber** | Pointer to **int64** | The block height calculated as the number of record files starting from zero since network start. | [optional] 
**RootContractId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**TransactionHash** | Pointer to **string** | A hex encoded transaction hash | [optional] 
**TransactionIndex** | Pointer to **NullableInt32** | The position of the transaction in the block | [optional] 

## Methods

### NewContractLog

`func NewContractLog() *ContractLog`

NewContractLog instantiates a new ContractLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractLogWithDefaults

`func NewContractLogWithDefaults() *ContractLog`

NewContractLogWithDefaults instantiates a new ContractLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractLog) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractLog) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractLog) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractLog) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBloom

`func (o *ContractLog) GetBloom() *os.File`

GetBloom returns the Bloom field if non-nil, zero value otherwise.

### GetBloomOk

`func (o *ContractLog) GetBloomOk() (**os.File, bool)`

GetBloomOk returns a tuple with the Bloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloom

`func (o *ContractLog) SetBloom(v *os.File)`

SetBloom sets Bloom field to given value.

### HasBloom

`func (o *ContractLog) HasBloom() bool`

HasBloom returns a boolean if a field has been set.

### GetContractId

`func (o *ContractLog) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractLog) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractLog) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ContractLog) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### SetContractIdNil

`func (o *ContractLog) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *ContractLog) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetData

`func (o *ContractLog) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContractLog) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContractLog) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ContractLog) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ContractLog) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ContractLog) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetIndex

`func (o *ContractLog) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ContractLog) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ContractLog) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ContractLog) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetTopics

`func (o *ContractLog) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ContractLog) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ContractLog) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *ContractLog) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetBlockHash

`func (o *ContractLog) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ContractLog) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ContractLog) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *ContractLog) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockNumber

`func (o *ContractLog) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ContractLog) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ContractLog) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *ContractLog) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetRootContractId

`func (o *ContractLog) GetRootContractId() string`

GetRootContractId returns the RootContractId field if non-nil, zero value otherwise.

### GetRootContractIdOk

`func (o *ContractLog) GetRootContractIdOk() (*string, bool)`

GetRootContractIdOk returns a tuple with the RootContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootContractId

`func (o *ContractLog) SetRootContractId(v string)`

SetRootContractId sets RootContractId field to given value.

### HasRootContractId

`func (o *ContractLog) HasRootContractId() bool`

HasRootContractId returns a boolean if a field has been set.

### GetTimestamp

`func (o *ContractLog) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContractLog) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContractLog) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContractLog) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTransactionHash

`func (o *ContractLog) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ContractLog) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ContractLog) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *ContractLog) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetTransactionIndex

`func (o *ContractLog) GetTransactionIndex() int32`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *ContractLog) GetTransactionIndexOk() (*int32, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *ContractLog) SetTransactionIndex(v int32)`

SetTransactionIndex sets TransactionIndex field to given value.

### HasTransactionIndex

`func (o *ContractLog) HasTransactionIndex() bool`

HasTransactionIndex returns a boolean if a field has been set.

### SetTransactionIndexNil

`func (o *ContractLog) SetTransactionIndexNil(b bool)`

 SetTransactionIndexNil sets the value for TransactionIndex to be an explicit nil

### UnsetTransactionIndex
`func (o *ContractLog) UnsetTransactionIndex()`

UnsetTransactionIndex ensures that no value is present for TransactionIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


