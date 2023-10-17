# ContractCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | Pointer to **NullableString** | Hexadecimal block number or the string \&quot;latest\&quot;, \&quot;pending\&quot;, \&quot;earliest\&quot;. Defaults to \&quot;latest\&quot;. | [optional] 
**Data** | Pointer to **Nullable*os.File** | Hexadecimal method signature and encoded parameters. | [optional] 
**Estimate** | Pointer to **NullableBool** | Whether gas estimation is called. Defaults to false. | [optional] 
**From** | Pointer to **Nullable*os.File** | The 20-byte hexadecimal EVM address the transaction is sent from. | [optional] 
**Gas** | Pointer to **NullableInt64** | Gas provided for the transaction execution. Defaults to 15000000. | [optional] 
**GasPrice** | Pointer to **NullableInt64** | Gas price used for each paid gas. | [optional] 
**To** | ***os.File** | The 20-byte hexadecimal EVM address the transaction is directed to. | 
**Value** | Pointer to **NullableInt64** | Value sent with this transaction. Defaults to 0. | [optional] 

## Methods

### NewContractCallRequest

`func NewContractCallRequest(to *os.File, ) *ContractCallRequest`

NewContractCallRequest instantiates a new ContractCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractCallRequestWithDefaults

`func NewContractCallRequestWithDefaults() *ContractCallRequest`

NewContractCallRequestWithDefaults instantiates a new ContractCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *ContractCallRequest) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ContractCallRequest) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ContractCallRequest) SetBlock(v string)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *ContractCallRequest) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### SetBlockNil

`func (o *ContractCallRequest) SetBlockNil(b bool)`

 SetBlockNil sets the value for Block to be an explicit nil

### UnsetBlock
`func (o *ContractCallRequest) UnsetBlock()`

UnsetBlock ensures that no value is present for Block, not even an explicit nil
### GetData

`func (o *ContractCallRequest) GetData() *os.File`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContractCallRequest) GetDataOk() (**os.File, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContractCallRequest) SetData(v *os.File)`

SetData sets Data field to given value.

### HasData

`func (o *ContractCallRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ContractCallRequest) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ContractCallRequest) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetEstimate

`func (o *ContractCallRequest) GetEstimate() bool`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *ContractCallRequest) GetEstimateOk() (*bool, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *ContractCallRequest) SetEstimate(v bool)`

SetEstimate sets Estimate field to given value.

### HasEstimate

`func (o *ContractCallRequest) HasEstimate() bool`

HasEstimate returns a boolean if a field has been set.

### SetEstimateNil

`func (o *ContractCallRequest) SetEstimateNil(b bool)`

 SetEstimateNil sets the value for Estimate to be an explicit nil

### UnsetEstimate
`func (o *ContractCallRequest) UnsetEstimate()`

UnsetEstimate ensures that no value is present for Estimate, not even an explicit nil
### GetFrom

`func (o *ContractCallRequest) GetFrom() *os.File`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ContractCallRequest) GetFromOk() (**os.File, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ContractCallRequest) SetFrom(v *os.File)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ContractCallRequest) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ContractCallRequest) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ContractCallRequest) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetGas

`func (o *ContractCallRequest) GetGas() int64`

GetGas returns the Gas field if non-nil, zero value otherwise.

### GetGasOk

`func (o *ContractCallRequest) GetGasOk() (*int64, bool)`

GetGasOk returns a tuple with the Gas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGas

`func (o *ContractCallRequest) SetGas(v int64)`

SetGas sets Gas field to given value.

### HasGas

`func (o *ContractCallRequest) HasGas() bool`

HasGas returns a boolean if a field has been set.

### SetGasNil

`func (o *ContractCallRequest) SetGasNil(b bool)`

 SetGasNil sets the value for Gas to be an explicit nil

### UnsetGas
`func (o *ContractCallRequest) UnsetGas()`

UnsetGas ensures that no value is present for Gas, not even an explicit nil
### GetGasPrice

`func (o *ContractCallRequest) GetGasPrice() int64`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *ContractCallRequest) GetGasPriceOk() (*int64, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *ContractCallRequest) SetGasPrice(v int64)`

SetGasPrice sets GasPrice field to given value.

### HasGasPrice

`func (o *ContractCallRequest) HasGasPrice() bool`

HasGasPrice returns a boolean if a field has been set.

### SetGasPriceNil

`func (o *ContractCallRequest) SetGasPriceNil(b bool)`

 SetGasPriceNil sets the value for GasPrice to be an explicit nil

### UnsetGasPrice
`func (o *ContractCallRequest) UnsetGasPrice()`

UnsetGasPrice ensures that no value is present for GasPrice, not even an explicit nil
### GetTo

`func (o *ContractCallRequest) GetTo() *os.File`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContractCallRequest) GetToOk() (**os.File, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContractCallRequest) SetTo(v *os.File)`

SetTo sets To field to given value.


### GetValue

`func (o *ContractCallRequest) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContractCallRequest) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContractCallRequest) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *ContractCallRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ContractCallRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ContractCallRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


