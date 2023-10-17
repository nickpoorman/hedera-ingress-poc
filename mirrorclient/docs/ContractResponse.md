# ContractResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminKey** | Pointer to [**NullableKey**](Key.md) |  | [optional] 
**AutoRenewAccount** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**AutoRenewPeriod** | Pointer to **NullableInt64** |  | [optional] 
**ContractId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**CreatedTimestamp** | Pointer to **NullableString** |  | [optional] 
**Deleted** | Pointer to **bool** |  | [optional] 
**EvmAddress** | Pointer to ***os.File** | A network entity encoded as an EVM address in hex. | [optional] 
**ExpirationTimestamp** | Pointer to **NullableString** |  | [optional] 
**FileId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**MaxAutomaticTokenAssociations** | Pointer to **NullableInt32** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**ObtainerId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**PermanentRemoval** | Pointer to **NullableBool** |  | [optional] 
**ProxyAccountId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**Timestamp** | Pointer to [**TimestampRange**](TimestampRange.md) |  | [optional] 
**Bytecode** | Pointer to **Nullable*os.File** | The contract bytecode in hex during deployment | [optional] 
**RuntimeBytecode** | Pointer to **Nullable*os.File** | The contract bytecode in hex after deployment | [optional] 

## Methods

### NewContractResponse

`func NewContractResponse() *ContractResponse`

NewContractResponse instantiates a new ContractResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractResponseWithDefaults

`func NewContractResponseWithDefaults() *ContractResponse`

NewContractResponseWithDefaults instantiates a new ContractResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminKey

`func (o *ContractResponse) GetAdminKey() Key`

GetAdminKey returns the AdminKey field if non-nil, zero value otherwise.

### GetAdminKeyOk

`func (o *ContractResponse) GetAdminKeyOk() (*Key, bool)`

GetAdminKeyOk returns a tuple with the AdminKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminKey

`func (o *ContractResponse) SetAdminKey(v Key)`

SetAdminKey sets AdminKey field to given value.

### HasAdminKey

`func (o *ContractResponse) HasAdminKey() bool`

HasAdminKey returns a boolean if a field has been set.

### SetAdminKeyNil

`func (o *ContractResponse) SetAdminKeyNil(b bool)`

 SetAdminKeyNil sets the value for AdminKey to be an explicit nil

### UnsetAdminKey
`func (o *ContractResponse) UnsetAdminKey()`

UnsetAdminKey ensures that no value is present for AdminKey, not even an explicit nil
### GetAutoRenewAccount

`func (o *ContractResponse) GetAutoRenewAccount() string`

GetAutoRenewAccount returns the AutoRenewAccount field if non-nil, zero value otherwise.

### GetAutoRenewAccountOk

`func (o *ContractResponse) GetAutoRenewAccountOk() (*string, bool)`

GetAutoRenewAccountOk returns a tuple with the AutoRenewAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewAccount

`func (o *ContractResponse) SetAutoRenewAccount(v string)`

SetAutoRenewAccount sets AutoRenewAccount field to given value.

### HasAutoRenewAccount

`func (o *ContractResponse) HasAutoRenewAccount() bool`

HasAutoRenewAccount returns a boolean if a field has been set.

### SetAutoRenewAccountNil

`func (o *ContractResponse) SetAutoRenewAccountNil(b bool)`

 SetAutoRenewAccountNil sets the value for AutoRenewAccount to be an explicit nil

### UnsetAutoRenewAccount
`func (o *ContractResponse) UnsetAutoRenewAccount()`

UnsetAutoRenewAccount ensures that no value is present for AutoRenewAccount, not even an explicit nil
### GetAutoRenewPeriod

`func (o *ContractResponse) GetAutoRenewPeriod() int64`

GetAutoRenewPeriod returns the AutoRenewPeriod field if non-nil, zero value otherwise.

### GetAutoRenewPeriodOk

`func (o *ContractResponse) GetAutoRenewPeriodOk() (*int64, bool)`

GetAutoRenewPeriodOk returns a tuple with the AutoRenewPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRenewPeriod

`func (o *ContractResponse) SetAutoRenewPeriod(v int64)`

SetAutoRenewPeriod sets AutoRenewPeriod field to given value.

### HasAutoRenewPeriod

`func (o *ContractResponse) HasAutoRenewPeriod() bool`

HasAutoRenewPeriod returns a boolean if a field has been set.

### SetAutoRenewPeriodNil

`func (o *ContractResponse) SetAutoRenewPeriodNil(b bool)`

 SetAutoRenewPeriodNil sets the value for AutoRenewPeriod to be an explicit nil

### UnsetAutoRenewPeriod
`func (o *ContractResponse) UnsetAutoRenewPeriod()`

UnsetAutoRenewPeriod ensures that no value is present for AutoRenewPeriod, not even an explicit nil
### GetContractId

`func (o *ContractResponse) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractResponse) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractResponse) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ContractResponse) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### SetContractIdNil

`func (o *ContractResponse) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *ContractResponse) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetCreatedTimestamp

`func (o *ContractResponse) GetCreatedTimestamp() string`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ContractResponse) GetCreatedTimestampOk() (*string, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ContractResponse) SetCreatedTimestamp(v string)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ContractResponse) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### SetCreatedTimestampNil

`func (o *ContractResponse) SetCreatedTimestampNil(b bool)`

 SetCreatedTimestampNil sets the value for CreatedTimestamp to be an explicit nil

### UnsetCreatedTimestamp
`func (o *ContractResponse) UnsetCreatedTimestamp()`

UnsetCreatedTimestamp ensures that no value is present for CreatedTimestamp, not even an explicit nil
### GetDeleted

`func (o *ContractResponse) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *ContractResponse) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *ContractResponse) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *ContractResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEvmAddress

`func (o *ContractResponse) GetEvmAddress() *os.File`

GetEvmAddress returns the EvmAddress field if non-nil, zero value otherwise.

### GetEvmAddressOk

`func (o *ContractResponse) GetEvmAddressOk() (**os.File, bool)`

GetEvmAddressOk returns a tuple with the EvmAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvmAddress

`func (o *ContractResponse) SetEvmAddress(v *os.File)`

SetEvmAddress sets EvmAddress field to given value.

### HasEvmAddress

`func (o *ContractResponse) HasEvmAddress() bool`

HasEvmAddress returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *ContractResponse) GetExpirationTimestamp() string`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *ContractResponse) GetExpirationTimestampOk() (*string, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *ContractResponse) SetExpirationTimestamp(v string)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *ContractResponse) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### SetExpirationTimestampNil

`func (o *ContractResponse) SetExpirationTimestampNil(b bool)`

 SetExpirationTimestampNil sets the value for ExpirationTimestamp to be an explicit nil

### UnsetExpirationTimestamp
`func (o *ContractResponse) UnsetExpirationTimestamp()`

UnsetExpirationTimestamp ensures that no value is present for ExpirationTimestamp, not even an explicit nil
### GetFileId

`func (o *ContractResponse) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ContractResponse) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ContractResponse) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ContractResponse) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *ContractResponse) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *ContractResponse) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetMaxAutomaticTokenAssociations

`func (o *ContractResponse) GetMaxAutomaticTokenAssociations() int32`

GetMaxAutomaticTokenAssociations returns the MaxAutomaticTokenAssociations field if non-nil, zero value otherwise.

### GetMaxAutomaticTokenAssociationsOk

`func (o *ContractResponse) GetMaxAutomaticTokenAssociationsOk() (*int32, bool)`

GetMaxAutomaticTokenAssociationsOk returns a tuple with the MaxAutomaticTokenAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAutomaticTokenAssociations

`func (o *ContractResponse) SetMaxAutomaticTokenAssociations(v int32)`

SetMaxAutomaticTokenAssociations sets MaxAutomaticTokenAssociations field to given value.

### HasMaxAutomaticTokenAssociations

`func (o *ContractResponse) HasMaxAutomaticTokenAssociations() bool`

HasMaxAutomaticTokenAssociations returns a boolean if a field has been set.

### SetMaxAutomaticTokenAssociationsNil

`func (o *ContractResponse) SetMaxAutomaticTokenAssociationsNil(b bool)`

 SetMaxAutomaticTokenAssociationsNil sets the value for MaxAutomaticTokenAssociations to be an explicit nil

### UnsetMaxAutomaticTokenAssociations
`func (o *ContractResponse) UnsetMaxAutomaticTokenAssociations()`

UnsetMaxAutomaticTokenAssociations ensures that no value is present for MaxAutomaticTokenAssociations, not even an explicit nil
### GetMemo

`func (o *ContractResponse) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ContractResponse) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ContractResponse) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ContractResponse) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetObtainerId

`func (o *ContractResponse) GetObtainerId() string`

GetObtainerId returns the ObtainerId field if non-nil, zero value otherwise.

### GetObtainerIdOk

`func (o *ContractResponse) GetObtainerIdOk() (*string, bool)`

GetObtainerIdOk returns a tuple with the ObtainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObtainerId

`func (o *ContractResponse) SetObtainerId(v string)`

SetObtainerId sets ObtainerId field to given value.

### HasObtainerId

`func (o *ContractResponse) HasObtainerId() bool`

HasObtainerId returns a boolean if a field has been set.

### SetObtainerIdNil

`func (o *ContractResponse) SetObtainerIdNil(b bool)`

 SetObtainerIdNil sets the value for ObtainerId to be an explicit nil

### UnsetObtainerId
`func (o *ContractResponse) UnsetObtainerId()`

UnsetObtainerId ensures that no value is present for ObtainerId, not even an explicit nil
### GetPermanentRemoval

`func (o *ContractResponse) GetPermanentRemoval() bool`

GetPermanentRemoval returns the PermanentRemoval field if non-nil, zero value otherwise.

### GetPermanentRemovalOk

`func (o *ContractResponse) GetPermanentRemovalOk() (*bool, bool)`

GetPermanentRemovalOk returns a tuple with the PermanentRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentRemoval

`func (o *ContractResponse) SetPermanentRemoval(v bool)`

SetPermanentRemoval sets PermanentRemoval field to given value.

### HasPermanentRemoval

`func (o *ContractResponse) HasPermanentRemoval() bool`

HasPermanentRemoval returns a boolean if a field has been set.

### SetPermanentRemovalNil

`func (o *ContractResponse) SetPermanentRemovalNil(b bool)`

 SetPermanentRemovalNil sets the value for PermanentRemoval to be an explicit nil

### UnsetPermanentRemoval
`func (o *ContractResponse) UnsetPermanentRemoval()`

UnsetPermanentRemoval ensures that no value is present for PermanentRemoval, not even an explicit nil
### GetProxyAccountId

`func (o *ContractResponse) GetProxyAccountId() string`

GetProxyAccountId returns the ProxyAccountId field if non-nil, zero value otherwise.

### GetProxyAccountIdOk

`func (o *ContractResponse) GetProxyAccountIdOk() (*string, bool)`

GetProxyAccountIdOk returns a tuple with the ProxyAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAccountId

`func (o *ContractResponse) SetProxyAccountId(v string)`

SetProxyAccountId sets ProxyAccountId field to given value.

### HasProxyAccountId

`func (o *ContractResponse) HasProxyAccountId() bool`

HasProxyAccountId returns a boolean if a field has been set.

### SetProxyAccountIdNil

`func (o *ContractResponse) SetProxyAccountIdNil(b bool)`

 SetProxyAccountIdNil sets the value for ProxyAccountId to be an explicit nil

### UnsetProxyAccountId
`func (o *ContractResponse) UnsetProxyAccountId()`

UnsetProxyAccountId ensures that no value is present for ProxyAccountId, not even an explicit nil
### GetTimestamp

`func (o *ContractResponse) GetTimestamp() TimestampRange`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContractResponse) GetTimestampOk() (*TimestampRange, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContractResponse) SetTimestamp(v TimestampRange)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContractResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetBytecode

`func (o *ContractResponse) GetBytecode() *os.File`

GetBytecode returns the Bytecode field if non-nil, zero value otherwise.

### GetBytecodeOk

`func (o *ContractResponse) GetBytecodeOk() (**os.File, bool)`

GetBytecodeOk returns a tuple with the Bytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytecode

`func (o *ContractResponse) SetBytecode(v *os.File)`

SetBytecode sets Bytecode field to given value.

### HasBytecode

`func (o *ContractResponse) HasBytecode() bool`

HasBytecode returns a boolean if a field has been set.

### SetBytecodeNil

`func (o *ContractResponse) SetBytecodeNil(b bool)`

 SetBytecodeNil sets the value for Bytecode to be an explicit nil

### UnsetBytecode
`func (o *ContractResponse) UnsetBytecode()`

UnsetBytecode ensures that no value is present for Bytecode, not even an explicit nil
### GetRuntimeBytecode

`func (o *ContractResponse) GetRuntimeBytecode() *os.File`

GetRuntimeBytecode returns the RuntimeBytecode field if non-nil, zero value otherwise.

### GetRuntimeBytecodeOk

`func (o *ContractResponse) GetRuntimeBytecodeOk() (**os.File, bool)`

GetRuntimeBytecodeOk returns a tuple with the RuntimeBytecode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeBytecode

`func (o *ContractResponse) SetRuntimeBytecode(v *os.File)`

SetRuntimeBytecode sets RuntimeBytecode field to given value.

### HasRuntimeBytecode

`func (o *ContractResponse) HasRuntimeBytecode() bool`

HasRuntimeBytecode returns a boolean if a field has been set.

### SetRuntimeBytecodeNil

`func (o *ContractResponse) SetRuntimeBytecodeNil(b bool)`

 SetRuntimeBytecodeNil sets the value for RuntimeBytecode to be an explicit nil

### UnsetRuntimeBytecode
`func (o *ContractResponse) UnsetRuntimeBytecode()`

UnsetRuntimeBytecode ensures that no value is present for RuntimeBytecode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


