# ContractResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessList** | Pointer to **NullableString** | The hex encoded access_list of the wrapped ethereum transaction | [optional] 
**Address** | Pointer to **string** | The hex encoded evm address of contract | [optional] 
**Amount** | Pointer to **NullableInt64** | The number of tinybars sent to the function | [optional] 
**BlockGasUsed** | Pointer to **NullableInt64** | The total amount of gas used in the block | [optional] 
**BlockHash** | Pointer to **NullableString** | The hex encoded block (record file chain) hash | [optional] 
**BlockNumber** | Pointer to **NullableInt64** | The block height calculated as the number of record files starting from zero since network start. | [optional] 
**Bloom** | Pointer to [***os.File**](*os.File.md) |  | [optional] 
**CallResult** | Pointer to **NullableString** | The hex encoded result returned by the function | [optional] 
**ChainId** | Pointer to **NullableString** | The hex encoded chain_id of the wrapped ethereum transaction | [optional] 
**ContractId** | Pointer to **NullableString** | Network entity ID in the format of &#x60;shard.realm.num&#x60; | [optional] 
**CreatedContractIds** | Pointer to **[]string** | The list of smart contracts that were created by the function call. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | The message when an error occurs during smart contract execution | [optional] 
**FailedInitcode** | Pointer to **string** | The hex encoded initcode of a failed contract create transaction | [optional] 
**From** | Pointer to ***os.File** | A network entity encoded as an EVM address in hex. | [optional] 
**FunctionParameters** | Pointer to **NullableString** | The hex encoded parameters passed to the function | [optional] 
**GasLimit** | Pointer to **int64** | The maximum units of gas allowed for contract execution | [optional] 
**GasPrice** | Pointer to **NullableString** | The hex encoded gas_price of the wrapped ethereum transaction | [optional] 
**GasUsed** | Pointer to **NullableInt64** | The units of gas used to execute contract | [optional] 
**Hash** | Pointer to **string** | A hex encoded 32 byte hash and it is only populated for Ethereum transaction case | [optional] 
**MaxFeePerGas** | Pointer to **NullableString** | The hex encoded max_fee_per_gas of the wrapped ethereum transaction | [optional] 
**MaxPriorityFeePerGas** | Pointer to **NullableString** | The hex encoded max_priority_fee_per_gas of the wrapped ethereum transaction | [optional] 
**Nonce** | Pointer to **NullableInt64** | The nonce of the wrapped ethereum transaction | [optional] 
**R** | Pointer to **NullableString** | The hex encoded signature_r of the wrapped ethereum transaction | [optional] 
**Result** | Pointer to **string** | The result of the transaction | [optional] 
**S** | Pointer to **NullableString** | The hex encoded signature_s of the wrapped ethereum transaction | [optional] 
**Status** | Pointer to **string** | The status of the transaction, 0x1 for a SUCCESS transaction and 0x0 for all else | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**To** | Pointer to **Nullable*os.File** | A network entity encoded as an EVM address in hex. | [optional] 
**TransactionIndex** | Pointer to **NullableInt64** | The position of the transaction in the block | [optional] 
**Type** | Pointer to **NullableInt32** | The type of the wrapped ethereum transaction, 0 (Pre-Eip1559) or 2 (Post-Eip1559) | [optional] 
**V** | Pointer to **NullableInt32** | The recovery_id of the wrapped ethereum transaction | [optional] 

## Methods

### NewContractResult

`func NewContractResult() *ContractResult`

NewContractResult instantiates a new ContractResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractResultWithDefaults

`func NewContractResultWithDefaults() *ContractResult`

NewContractResultWithDefaults instantiates a new ContractResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessList

`func (o *ContractResult) GetAccessList() string`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *ContractResult) GetAccessListOk() (*string, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *ContractResult) SetAccessList(v string)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *ContractResult) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### SetAccessListNil

`func (o *ContractResult) SetAccessListNil(b bool)`

 SetAccessListNil sets the value for AccessList to be an explicit nil

### UnsetAccessList
`func (o *ContractResult) UnsetAccessList()`

UnsetAccessList ensures that no value is present for AccessList, not even an explicit nil
### GetAddress

`func (o *ContractResult) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractResult) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractResult) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractResult) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAmount

`func (o *ContractResult) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ContractResult) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ContractResult) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ContractResult) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ContractResult) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ContractResult) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetBlockGasUsed

`func (o *ContractResult) GetBlockGasUsed() int64`

GetBlockGasUsed returns the BlockGasUsed field if non-nil, zero value otherwise.

### GetBlockGasUsedOk

`func (o *ContractResult) GetBlockGasUsedOk() (*int64, bool)`

GetBlockGasUsedOk returns a tuple with the BlockGasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockGasUsed

`func (o *ContractResult) SetBlockGasUsed(v int64)`

SetBlockGasUsed sets BlockGasUsed field to given value.

### HasBlockGasUsed

`func (o *ContractResult) HasBlockGasUsed() bool`

HasBlockGasUsed returns a boolean if a field has been set.

### SetBlockGasUsedNil

`func (o *ContractResult) SetBlockGasUsedNil(b bool)`

 SetBlockGasUsedNil sets the value for BlockGasUsed to be an explicit nil

### UnsetBlockGasUsed
`func (o *ContractResult) UnsetBlockGasUsed()`

UnsetBlockGasUsed ensures that no value is present for BlockGasUsed, not even an explicit nil
### GetBlockHash

`func (o *ContractResult) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ContractResult) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ContractResult) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *ContractResult) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### SetBlockHashNil

`func (o *ContractResult) SetBlockHashNil(b bool)`

 SetBlockHashNil sets the value for BlockHash to be an explicit nil

### UnsetBlockHash
`func (o *ContractResult) UnsetBlockHash()`

UnsetBlockHash ensures that no value is present for BlockHash, not even an explicit nil
### GetBlockNumber

`func (o *ContractResult) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ContractResult) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ContractResult) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *ContractResult) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### SetBlockNumberNil

`func (o *ContractResult) SetBlockNumberNil(b bool)`

 SetBlockNumberNil sets the value for BlockNumber to be an explicit nil

### UnsetBlockNumber
`func (o *ContractResult) UnsetBlockNumber()`

UnsetBlockNumber ensures that no value is present for BlockNumber, not even an explicit nil
### GetBloom

`func (o *ContractResult) GetBloom() *os.File`

GetBloom returns the Bloom field if non-nil, zero value otherwise.

### GetBloomOk

`func (o *ContractResult) GetBloomOk() (**os.File, bool)`

GetBloomOk returns a tuple with the Bloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloom

`func (o *ContractResult) SetBloom(v *os.File)`

SetBloom sets Bloom field to given value.

### HasBloom

`func (o *ContractResult) HasBloom() bool`

HasBloom returns a boolean if a field has been set.

### GetCallResult

`func (o *ContractResult) GetCallResult() string`

GetCallResult returns the CallResult field if non-nil, zero value otherwise.

### GetCallResultOk

`func (o *ContractResult) GetCallResultOk() (*string, bool)`

GetCallResultOk returns a tuple with the CallResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallResult

`func (o *ContractResult) SetCallResult(v string)`

SetCallResult sets CallResult field to given value.

### HasCallResult

`func (o *ContractResult) HasCallResult() bool`

HasCallResult returns a boolean if a field has been set.

### SetCallResultNil

`func (o *ContractResult) SetCallResultNil(b bool)`

 SetCallResultNil sets the value for CallResult to be an explicit nil

### UnsetCallResult
`func (o *ContractResult) UnsetCallResult()`

UnsetCallResult ensures that no value is present for CallResult, not even an explicit nil
### GetChainId

`func (o *ContractResult) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ContractResult) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ContractResult) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ContractResult) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### SetChainIdNil

`func (o *ContractResult) SetChainIdNil(b bool)`

 SetChainIdNil sets the value for ChainId to be an explicit nil

### UnsetChainId
`func (o *ContractResult) UnsetChainId()`

UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
### GetContractId

`func (o *ContractResult) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractResult) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractResult) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ContractResult) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### SetContractIdNil

`func (o *ContractResult) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *ContractResult) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetCreatedContractIds

`func (o *ContractResult) GetCreatedContractIds() []*string`

GetCreatedContractIds returns the CreatedContractIds field if non-nil, zero value otherwise.

### GetCreatedContractIdsOk

`func (o *ContractResult) GetCreatedContractIdsOk() (*[]*string, bool)`

GetCreatedContractIdsOk returns a tuple with the CreatedContractIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedContractIds

`func (o *ContractResult) SetCreatedContractIds(v []*string)`

SetCreatedContractIds sets CreatedContractIds field to given value.

### HasCreatedContractIds

`func (o *ContractResult) HasCreatedContractIds() bool`

HasCreatedContractIds returns a boolean if a field has been set.

### SetCreatedContractIdsNil

`func (o *ContractResult) SetCreatedContractIdsNil(b bool)`

 SetCreatedContractIdsNil sets the value for CreatedContractIds to be an explicit nil

### UnsetCreatedContractIds
`func (o *ContractResult) UnsetCreatedContractIds()`

UnsetCreatedContractIds ensures that no value is present for CreatedContractIds, not even an explicit nil
### GetErrorMessage

`func (o *ContractResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ContractResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ContractResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ContractResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ContractResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ContractResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFailedInitcode

`func (o *ContractResult) GetFailedInitcode() string`

GetFailedInitcode returns the FailedInitcode field if non-nil, zero value otherwise.

### GetFailedInitcodeOk

`func (o *ContractResult) GetFailedInitcodeOk() (*string, bool)`

GetFailedInitcodeOk returns a tuple with the FailedInitcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInitcode

`func (o *ContractResult) SetFailedInitcode(v string)`

SetFailedInitcode sets FailedInitcode field to given value.

### HasFailedInitcode

`func (o *ContractResult) HasFailedInitcode() bool`

HasFailedInitcode returns a boolean if a field has been set.

### GetFrom

`func (o *ContractResult) GetFrom() *os.File`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ContractResult) GetFromOk() (**os.File, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ContractResult) SetFrom(v *os.File)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ContractResult) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFunctionParameters

`func (o *ContractResult) GetFunctionParameters() string`

GetFunctionParameters returns the FunctionParameters field if non-nil, zero value otherwise.

### GetFunctionParametersOk

`func (o *ContractResult) GetFunctionParametersOk() (*string, bool)`

GetFunctionParametersOk returns a tuple with the FunctionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionParameters

`func (o *ContractResult) SetFunctionParameters(v string)`

SetFunctionParameters sets FunctionParameters field to given value.

### HasFunctionParameters

`func (o *ContractResult) HasFunctionParameters() bool`

HasFunctionParameters returns a boolean if a field has been set.

### SetFunctionParametersNil

`func (o *ContractResult) SetFunctionParametersNil(b bool)`

 SetFunctionParametersNil sets the value for FunctionParameters to be an explicit nil

### UnsetFunctionParameters
`func (o *ContractResult) UnsetFunctionParameters()`

UnsetFunctionParameters ensures that no value is present for FunctionParameters, not even an explicit nil
### GetGasLimit

`func (o *ContractResult) GetGasLimit() int64`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ContractResult) GetGasLimitOk() (*int64, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ContractResult) SetGasLimit(v int64)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *ContractResult) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetGasPrice

`func (o *ContractResult) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ContractResult) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ContractResult) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *ContractResult) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### SetGasPriceNil

`func (o *ContractResult) SetGasPriceNil(b bool)`

 SetGasPriceNil sets the value for GasPrice to be an explicit nil

### UnsetGasPrice
`func (o *ContractResult) UnsetGasPrice()`

UnsetGasPrice ensures that no value is present for GasPrice, not even an explicit nil
### GetGasUsed

`func (o *ContractResult) GetGasUsed() int64`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ContractResult) GetGasUsedOk() (*int64, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ContractResult) SetGasUsed(v int64)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *ContractResult) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.

### SetGasUsedNil

`func (o *ContractResult) SetGasUsedNil(b bool)`

 SetGasUsedNil sets the value for GasUsed to be an explicit nil

### UnsetGasUsed
`func (o *ContractResult) UnsetGasUsed()`

UnsetGasUsed ensures that no value is present for GasUsed, not even an explicit nil
### GetHash

`func (o *ContractResult) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ContractResult) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ContractResult) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ContractResult) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *ContractResult) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *ContractResult) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *ContractResult) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *ContractResult) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### SetMaxFeePerGasNil

`func (o *ContractResult) SetMaxFeePerGasNil(b bool)`

 SetMaxFeePerGasNil sets the value for MaxFeePerGas to be an explicit nil

### UnsetMaxFeePerGas
`func (o *ContractResult) UnsetMaxFeePerGas()`

UnsetMaxFeePerGas ensures that no value is present for MaxFeePerGas, not even an explicit nil
### GetMaxPriorityFeePerGas

`func (o *ContractResult) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *ContractResult) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *ContractResult) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *ContractResult) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### SetMaxPriorityFeePerGasNil

`func (o *ContractResult) SetMaxPriorityFeePerGasNil(b bool)`

 SetMaxPriorityFeePerGasNil sets the value for MaxPriorityFeePerGas to be an explicit nil

### UnsetMaxPriorityFeePerGas
`func (o *ContractResult) UnsetMaxPriorityFeePerGas()`

UnsetMaxPriorityFeePerGas ensures that no value is present for MaxPriorityFeePerGas, not even an explicit nil
### GetNonce

`func (o *ContractResult) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ContractResult) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ContractResult) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ContractResult) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *ContractResult) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *ContractResult) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetR

`func (o *ContractResult) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *ContractResult) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *ContractResult) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *ContractResult) HasR() bool`

HasR returns a boolean if a field has been set.

### SetRNil

`func (o *ContractResult) SetRNil(b bool)`

 SetRNil sets the value for R to be an explicit nil

### UnsetR
`func (o *ContractResult) UnsetR()`

UnsetR ensures that no value is present for R, not even an explicit nil
### GetResult

`func (o *ContractResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ContractResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ContractResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ContractResult) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetS

`func (o *ContractResult) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *ContractResult) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *ContractResult) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *ContractResult) HasS() bool`

HasS returns a boolean if a field has been set.

### SetSNil

`func (o *ContractResult) SetSNil(b bool)`

 SetSNil sets the value for S to be an explicit nil

### UnsetS
`func (o *ContractResult) UnsetS()`

UnsetS ensures that no value is present for S, not even an explicit nil
### GetStatus

`func (o *ContractResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContractResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *ContractResult) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContractResult) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContractResult) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContractResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTo

`func (o *ContractResult) GetTo() *os.File`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContractResult) GetToOk() (**os.File, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContractResult) SetTo(v *os.File)`

SetTo sets To field to given value.

### HasTo

`func (o *ContractResult) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ContractResult) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ContractResult) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTransactionIndex

`func (o *ContractResult) GetTransactionIndex() int64`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *ContractResult) GetTransactionIndexOk() (*int64, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *ContractResult) SetTransactionIndex(v int64)`

SetTransactionIndex sets TransactionIndex field to given value.

### HasTransactionIndex

`func (o *ContractResult) HasTransactionIndex() bool`

HasTransactionIndex returns a boolean if a field has been set.

### SetTransactionIndexNil

`func (o *ContractResult) SetTransactionIndexNil(b bool)`

 SetTransactionIndexNil sets the value for TransactionIndex to be an explicit nil

### UnsetTransactionIndex
`func (o *ContractResult) UnsetTransactionIndex()`

UnsetTransactionIndex ensures that no value is present for TransactionIndex, not even an explicit nil
### GetType

`func (o *ContractResult) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractResult) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractResult) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ContractResult) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ContractResult) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ContractResult) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetV

`func (o *ContractResult) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ContractResult) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ContractResult) SetV(v int32)`

SetV sets V field to given value.

### HasV

`func (o *ContractResult) HasV() bool`

HasV returns a boolean if a field has been set.

### SetVNil

`func (o *ContractResult) SetVNil(b bool)`

 SetVNil sets the value for V to be an explicit nil

### UnsetV
`func (o *ContractResult) UnsetV()`

UnsetV ensures that no value is present for V, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


