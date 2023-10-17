# ContractResultLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The hex encoded EVM address of the contract | [optional] 
**Bloom** | Pointer to [***os.File**](*os.File.md) |  | [optional] 
**ContractId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Data** | Pointer to **NullableString** | The hex encoded data of the contract log | [optional] 
**Index** | Pointer to **int32** | The index of the contract log in the chain of logs for an execution | [optional] 
**Topics** | Pointer to **[]string** | A list of hex encoded topics associated with this log event | [optional] 

## Methods

### NewContractResultLog

`func NewContractResultLog() *ContractResultLog`

NewContractResultLog instantiates a new ContractResultLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractResultLogWithDefaults

`func NewContractResultLogWithDefaults() *ContractResultLog`

NewContractResultLogWithDefaults instantiates a new ContractResultLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ContractResultLog) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractResultLog) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractResultLog) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractResultLog) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBloom

`func (o *ContractResultLog) GetBloom() *os.File`

GetBloom returns the Bloom field if non-nil, zero value otherwise.

### GetBloomOk

`func (o *ContractResultLog) GetBloomOk() (**os.File, bool)`

GetBloomOk returns a tuple with the Bloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloom

`func (o *ContractResultLog) SetBloom(v *os.File)`

SetBloom sets Bloom field to given value.

### HasBloom

`func (o *ContractResultLog) HasBloom() bool`

HasBloom returns a boolean if a field has been set.

### GetContractId

`func (o *ContractResultLog) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractResultLog) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractResultLog) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ContractResultLog) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### SetContractIdNil

`func (o *ContractResultLog) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *ContractResultLog) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetData

`func (o *ContractResultLog) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContractResultLog) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContractResultLog) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ContractResultLog) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ContractResultLog) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ContractResultLog) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetIndex

`func (o *ContractResultLog) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ContractResultLog) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ContractResultLog) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ContractResultLog) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetTopics

`func (o *ContractResultLog) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *ContractResultLog) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *ContractResultLog) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *ContractResultLog) HasTopics() bool`

HasTopics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


