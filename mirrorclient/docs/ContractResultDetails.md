# ContractResultDetails

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
**Hash** | Pointer to **string** | The hex encoded transaction hash | [optional] 
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
**Logs** | Pointer to [**[]ContractResultLog**](ContractResultLog.md) |  | [optional] 
**StateChanges** | Pointer to [**[]ContractResultStateChange**](ContractResultStateChange.md) |  | [optional] 

## Methods

### NewContractResultDetails

`func NewContractResultDetails() *ContractResultDetails`

NewContractResultDetails instantiates a new ContractResultDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractResultDetailsWithDefaults

`func NewContractResultDetailsWithDefaults() *ContractResultDetails`

NewContractResultDetailsWithDefaults instantiates a new ContractResultDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessList

`func (o *ContractResultDetails) GetAccessList() string`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *ContractResultDetails) GetAccessListOk() (*string, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *ContractResultDetails) SetAccessList(v string)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *ContractResultDetails) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### SetAccessListNil

`func (o *ContractResultDetails) SetAccessListNil(b bool)`

 SetAccessListNil sets the value for AccessList to be an explicit nil

### UnsetAccessList
`func (o *ContractResultDetails) UnsetAccessList()`

UnsetAccessList ensures that no value is present for AccessList, not even an explicit nil
### GetAddress

`func (o *ContractResultDetails) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ContractResultDetails) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ContractResultDetails) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ContractResultDetails) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAmount

`func (o *ContractResultDetails) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ContractResultDetails) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ContractResultDetails) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *ContractResultDetails) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *ContractResultDetails) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *ContractResultDetails) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetBlockGasUsed

`func (o *ContractResultDetails) GetBlockGasUsed() int64`

GetBlockGasUsed returns the BlockGasUsed field if non-nil, zero value otherwise.

### GetBlockGasUsedOk

`func (o *ContractResultDetails) GetBlockGasUsedOk() (*int64, bool)`

GetBlockGasUsedOk returns a tuple with the BlockGasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockGasUsed

`func (o *ContractResultDetails) SetBlockGasUsed(v int64)`

SetBlockGasUsed sets BlockGasUsed field to given value.

### HasBlockGasUsed

`func (o *ContractResultDetails) HasBlockGasUsed() bool`

HasBlockGasUsed returns a boolean if a field has been set.

### SetBlockGasUsedNil

`func (o *ContractResultDetails) SetBlockGasUsedNil(b bool)`

 SetBlockGasUsedNil sets the value for BlockGasUsed to be an explicit nil

### UnsetBlockGasUsed
`func (o *ContractResultDetails) UnsetBlockGasUsed()`

UnsetBlockGasUsed ensures that no value is present for BlockGasUsed, not even an explicit nil
### GetBlockHash

`func (o *ContractResultDetails) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ContractResultDetails) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ContractResultDetails) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *ContractResultDetails) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### SetBlockHashNil

`func (o *ContractResultDetails) SetBlockHashNil(b bool)`

 SetBlockHashNil sets the value for BlockHash to be an explicit nil

### UnsetBlockHash
`func (o *ContractResultDetails) UnsetBlockHash()`

UnsetBlockHash ensures that no value is present for BlockHash, not even an explicit nil
### GetBlockNumber

`func (o *ContractResultDetails) GetBlockNumber() int64`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ContractResultDetails) GetBlockNumberOk() (*int64, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ContractResultDetails) SetBlockNumber(v int64)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *ContractResultDetails) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### SetBlockNumberNil

`func (o *ContractResultDetails) SetBlockNumberNil(b bool)`

 SetBlockNumberNil sets the value for BlockNumber to be an explicit nil

### UnsetBlockNumber
`func (o *ContractResultDetails) UnsetBlockNumber()`

UnsetBlockNumber ensures that no value is present for BlockNumber, not even an explicit nil
### GetBloom

`func (o *ContractResultDetails) GetBloom() *os.File`

GetBloom returns the Bloom field if non-nil, zero value otherwise.

### GetBloomOk

`func (o *ContractResultDetails) GetBloomOk() (**os.File, bool)`

GetBloomOk returns a tuple with the Bloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloom

`func (o *ContractResultDetails) SetBloom(v *os.File)`

SetBloom sets Bloom field to given value.

### HasBloom

`func (o *ContractResultDetails) HasBloom() bool`

HasBloom returns a boolean if a field has been set.

### GetCallResult

`func (o *ContractResultDetails) GetCallResult() string`

GetCallResult returns the CallResult field if non-nil, zero value otherwise.

### GetCallResultOk

`func (o *ContractResultDetails) GetCallResultOk() (*string, bool)`

GetCallResultOk returns a tuple with the CallResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallResult

`func (o *ContractResultDetails) SetCallResult(v string)`

SetCallResult sets CallResult field to given value.

### HasCallResult

`func (o *ContractResultDetails) HasCallResult() bool`

HasCallResult returns a boolean if a field has been set.

### SetCallResultNil

`func (o *ContractResultDetails) SetCallResultNil(b bool)`

 SetCallResultNil sets the value for CallResult to be an explicit nil

### UnsetCallResult
`func (o *ContractResultDetails) UnsetCallResult()`

UnsetCallResult ensures that no value is present for CallResult, not even an explicit nil
### GetChainId

`func (o *ContractResultDetails) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ContractResultDetails) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ContractResultDetails) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ContractResultDetails) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### SetChainIdNil

`func (o *ContractResultDetails) SetChainIdNil(b bool)`

 SetChainIdNil sets the value for ChainId to be an explicit nil

### UnsetChainId
`func (o *ContractResultDetails) UnsetChainId()`

UnsetChainId ensures that no value is present for ChainId, not even an explicit nil
### GetContractId

`func (o *ContractResultDetails) GetContractId() string`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *ContractResultDetails) GetContractIdOk() (*string, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *ContractResultDetails) SetContractId(v string)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *ContractResultDetails) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### SetContractIdNil

`func (o *ContractResultDetails) SetContractIdNil(b bool)`

 SetContractIdNil sets the value for ContractId to be an explicit nil

### UnsetContractId
`func (o *ContractResultDetails) UnsetContractId()`

UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
### GetCreatedContractIds

`func (o *ContractResultDetails) GetCreatedContractIds() []*string`

GetCreatedContractIds returns the CreatedContractIds field if non-nil, zero value otherwise.

### GetCreatedContractIdsOk

`func (o *ContractResultDetails) GetCreatedContractIdsOk() (*[]*string, bool)`

GetCreatedContractIdsOk returns a tuple with the CreatedContractIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedContractIds

`func (o *ContractResultDetails) SetCreatedContractIds(v []*string)`

SetCreatedContractIds sets CreatedContractIds field to given value.

### HasCreatedContractIds

`func (o *ContractResultDetails) HasCreatedContractIds() bool`

HasCreatedContractIds returns a boolean if a field has been set.

### SetCreatedContractIdsNil

`func (o *ContractResultDetails) SetCreatedContractIdsNil(b bool)`

 SetCreatedContractIdsNil sets the value for CreatedContractIds to be an explicit nil

### UnsetCreatedContractIds
`func (o *ContractResultDetails) UnsetCreatedContractIds()`

UnsetCreatedContractIds ensures that no value is present for CreatedContractIds, not even an explicit nil
### GetErrorMessage

`func (o *ContractResultDetails) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ContractResultDetails) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ContractResultDetails) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ContractResultDetails) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ContractResultDetails) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ContractResultDetails) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetFailedInitcode

`func (o *ContractResultDetails) GetFailedInitcode() string`

GetFailedInitcode returns the FailedInitcode field if non-nil, zero value otherwise.

### GetFailedInitcodeOk

`func (o *ContractResultDetails) GetFailedInitcodeOk() (*string, bool)`

GetFailedInitcodeOk returns a tuple with the FailedInitcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedInitcode

`func (o *ContractResultDetails) SetFailedInitcode(v string)`

SetFailedInitcode sets FailedInitcode field to given value.

### HasFailedInitcode

`func (o *ContractResultDetails) HasFailedInitcode() bool`

HasFailedInitcode returns a boolean if a field has been set.

### GetFrom

`func (o *ContractResultDetails) GetFrom() *os.File`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ContractResultDetails) GetFromOk() (**os.File, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ContractResultDetails) SetFrom(v *os.File)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ContractResultDetails) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFunctionParameters

`func (o *ContractResultDetails) GetFunctionParameters() string`

GetFunctionParameters returns the FunctionParameters field if non-nil, zero value otherwise.

### GetFunctionParametersOk

`func (o *ContractResultDetails) GetFunctionParametersOk() (*string, bool)`

GetFunctionParametersOk returns a tuple with the FunctionParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionParameters

`func (o *ContractResultDetails) SetFunctionParameters(v string)`

SetFunctionParameters sets FunctionParameters field to given value.

### HasFunctionParameters

`func (o *ContractResultDetails) HasFunctionParameters() bool`

HasFunctionParameters returns a boolean if a field has been set.

### SetFunctionParametersNil

`func (o *ContractResultDetails) SetFunctionParametersNil(b bool)`

 SetFunctionParametersNil sets the value for FunctionParameters to be an explicit nil

### UnsetFunctionParameters
`func (o *ContractResultDetails) UnsetFunctionParameters()`

UnsetFunctionParameters ensures that no value is present for FunctionParameters, not even an explicit nil
### GetGasLimit

`func (o *ContractResultDetails) GetGasLimit() int64`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *ContractResultDetails) GetGasLimitOk() (*int64, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *ContractResultDetails) SetGasLimit(v int64)`

SetGasLimit sets GasLimit field to given value.

### HasGasLimit

`func (o *ContractResultDetails) HasGasLimit() bool`

HasGasLimit returns a boolean if a field has been set.

### GetGasPrice

`func (o *ContractResultDetails) GetGasPrice() string`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ContractResultDetails) GetGasPriceOk() (*string, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ContractResultDetails) SetGasPrice(v string)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *ContractResultDetails) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### SetGasPriceNil

`func (o *ContractResultDetails) SetGasPriceNil(b bool)`

 SetGasPriceNil sets the value for GasPrice to be an explicit nil

### UnsetGasPrice
`func (o *ContractResultDetails) UnsetGasPrice()`

UnsetGasPrice ensures that no value is present for GasPrice, not even an explicit nil
### GetGasUsed

`func (o *ContractResultDetails) GetGasUsed() int64`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ContractResultDetails) GetGasUsedOk() (*int64, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ContractResultDetails) SetGasUsed(v int64)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *ContractResultDetails) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.

### SetGasUsedNil

`func (o *ContractResultDetails) SetGasUsedNil(b bool)`

 SetGasUsedNil sets the value for GasUsed to be an explicit nil

### UnsetGasUsed
`func (o *ContractResultDetails) UnsetGasUsed()`

UnsetGasUsed ensures that no value is present for GasUsed, not even an explicit nil
### GetHash

`func (o *ContractResultDetails) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ContractResultDetails) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ContractResultDetails) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ContractResultDetails) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMaxFeePerGas

`func (o *ContractResultDetails) GetMaxFeePerGas() string`

GetMaxFeePerGas returns the MaxFeePerGas field if non-nil, zero value otherwise.

### GetMaxFeePerGasOk

`func (o *ContractResultDetails) GetMaxFeePerGasOk() (*string, bool)`

GetMaxFeePerGasOk returns a tuple with the MaxFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFeePerGas

`func (o *ContractResultDetails) SetMaxFeePerGas(v string)`

SetMaxFeePerGas sets MaxFeePerGas field to given value.

### HasMaxFeePerGas

`func (o *ContractResultDetails) HasMaxFeePerGas() bool`

HasMaxFeePerGas returns a boolean if a field has been set.

### SetMaxFeePerGasNil

`func (o *ContractResultDetails) SetMaxFeePerGasNil(b bool)`

 SetMaxFeePerGasNil sets the value for MaxFeePerGas to be an explicit nil

### UnsetMaxFeePerGas
`func (o *ContractResultDetails) UnsetMaxFeePerGas()`

UnsetMaxFeePerGas ensures that no value is present for MaxFeePerGas, not even an explicit nil
### GetMaxPriorityFeePerGas

`func (o *ContractResultDetails) GetMaxPriorityFeePerGas() string`

GetMaxPriorityFeePerGas returns the MaxPriorityFeePerGas field if non-nil, zero value otherwise.

### GetMaxPriorityFeePerGasOk

`func (o *ContractResultDetails) GetMaxPriorityFeePerGasOk() (*string, bool)`

GetMaxPriorityFeePerGasOk returns a tuple with the MaxPriorityFeePerGas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPriorityFeePerGas

`func (o *ContractResultDetails) SetMaxPriorityFeePerGas(v string)`

SetMaxPriorityFeePerGas sets MaxPriorityFeePerGas field to given value.

### HasMaxPriorityFeePerGas

`func (o *ContractResultDetails) HasMaxPriorityFeePerGas() bool`

HasMaxPriorityFeePerGas returns a boolean if a field has been set.

### SetMaxPriorityFeePerGasNil

`func (o *ContractResultDetails) SetMaxPriorityFeePerGasNil(b bool)`

 SetMaxPriorityFeePerGasNil sets the value for MaxPriorityFeePerGas to be an explicit nil

### UnsetMaxPriorityFeePerGas
`func (o *ContractResultDetails) UnsetMaxPriorityFeePerGas()`

UnsetMaxPriorityFeePerGas ensures that no value is present for MaxPriorityFeePerGas, not even an explicit nil
### GetNonce

`func (o *ContractResultDetails) GetNonce() int64`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ContractResultDetails) GetNonceOk() (*int64, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ContractResultDetails) SetNonce(v int64)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ContractResultDetails) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### SetNonceNil

`func (o *ContractResultDetails) SetNonceNil(b bool)`

 SetNonceNil sets the value for Nonce to be an explicit nil

### UnsetNonce
`func (o *ContractResultDetails) UnsetNonce()`

UnsetNonce ensures that no value is present for Nonce, not even an explicit nil
### GetR

`func (o *ContractResultDetails) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *ContractResultDetails) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *ContractResultDetails) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *ContractResultDetails) HasR() bool`

HasR returns a boolean if a field has been set.

### SetRNil

`func (o *ContractResultDetails) SetRNil(b bool)`

 SetRNil sets the value for R to be an explicit nil

### UnsetR
`func (o *ContractResultDetails) UnsetR()`

UnsetR ensures that no value is present for R, not even an explicit nil
### GetResult

`func (o *ContractResultDetails) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ContractResultDetails) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ContractResultDetails) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *ContractResultDetails) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetS

`func (o *ContractResultDetails) GetS() string`

GetS returns the S field if non-nil, zero value otherwise.

### GetSOk

`func (o *ContractResultDetails) GetSOk() (*string, bool)`

GetSOk returns a tuple with the S field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS

`func (o *ContractResultDetails) SetS(v string)`

SetS sets S field to given value.

### HasS

`func (o *ContractResultDetails) HasS() bool`

HasS returns a boolean if a field has been set.

### SetSNil

`func (o *ContractResultDetails) SetSNil(b bool)`

 SetSNil sets the value for S to be an explicit nil

### UnsetS
`func (o *ContractResultDetails) UnsetS()`

UnsetS ensures that no value is present for S, not even an explicit nil
### GetStatus

`func (o *ContractResultDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContractResultDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContractResultDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContractResultDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *ContractResultDetails) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ContractResultDetails) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ContractResultDetails) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ContractResultDetails) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTo

`func (o *ContractResultDetails) GetTo() *os.File`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContractResultDetails) GetToOk() (**os.File, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContractResultDetails) SetTo(v *os.File)`

SetTo sets To field to given value.

### HasTo

`func (o *ContractResultDetails) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ContractResultDetails) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ContractResultDetails) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetTransactionIndex

`func (o *ContractResultDetails) GetTransactionIndex() int64`

GetTransactionIndex returns the TransactionIndex field if non-nil, zero value otherwise.

### GetTransactionIndexOk

`func (o *ContractResultDetails) GetTransactionIndexOk() (*int64, bool)`

GetTransactionIndexOk returns a tuple with the TransactionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIndex

`func (o *ContractResultDetails) SetTransactionIndex(v int64)`

SetTransactionIndex sets TransactionIndex field to given value.

### HasTransactionIndex

`func (o *ContractResultDetails) HasTransactionIndex() bool`

HasTransactionIndex returns a boolean if a field has been set.

### SetTransactionIndexNil

`func (o *ContractResultDetails) SetTransactionIndexNil(b bool)`

 SetTransactionIndexNil sets the value for TransactionIndex to be an explicit nil

### UnsetTransactionIndex
`func (o *ContractResultDetails) UnsetTransactionIndex()`

UnsetTransactionIndex ensures that no value is present for TransactionIndex, not even an explicit nil
### GetType

`func (o *ContractResultDetails) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContractResultDetails) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContractResultDetails) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *ContractResultDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ContractResultDetails) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ContractResultDetails) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetV

`func (o *ContractResultDetails) GetV() int32`

GetV returns the V field if non-nil, zero value otherwise.

### GetVOk

`func (o *ContractResultDetails) GetVOk() (*int32, bool)`

GetVOk returns a tuple with the V field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV

`func (o *ContractResultDetails) SetV(v int32)`

SetV sets V field to given value.

### HasV

`func (o *ContractResultDetails) HasV() bool`

HasV returns a boolean if a field has been set.

### SetVNil

`func (o *ContractResultDetails) SetVNil(b bool)`

 SetVNil sets the value for V to be an explicit nil

### UnsetV
`func (o *ContractResultDetails) UnsetV()`

UnsetV ensures that no value is present for V, not even an explicit nil
### GetLogs

`func (o *ContractResultDetails) GetLogs() []ContractResultLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ContractResultDetails) GetLogsOk() (*[]ContractResultLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ContractResultDetails) SetLogs(v []ContractResultLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ContractResultDetails) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetStateChanges

`func (o *ContractResultDetails) GetStateChanges() []ContractResultStateChange`

GetStateChanges returns the StateChanges field if non-nil, zero value otherwise.

### GetStateChangesOk

`func (o *ContractResultDetails) GetStateChangesOk() (*[]ContractResultStateChange, bool)`

GetStateChangesOk returns a tuple with the StateChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChanges

`func (o *ContractResultDetails) SetStateChanges(v []ContractResultStateChange)`

SetStateChanges sets StateChanges field to given value.

### HasStateChanges

`func (o *ContractResultDetails) HasStateChanges() bool`

HasStateChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


