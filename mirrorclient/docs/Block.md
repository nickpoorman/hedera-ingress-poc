# Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**GasUsed** | Pointer to **NullableInt64** |  | [optional] 
**HapiVersion** | Pointer to **NullableString** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**LogsBloom** | Pointer to **NullableString** | A hex encoded 256-byte array with 0x prefix | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] 
**PreviousHash** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | Pointer to [**BlockTimestamp**](BlockTimestamp.md) |  | [optional] 

## Methods

### NewBlock

`func NewBlock() *Block`

NewBlock instantiates a new Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWithDefaults

`func NewBlockWithDefaults() *Block`

NewBlockWithDefaults instantiates a new Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Block) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Block) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Block) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Block) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGasUsed

`func (o *Block) GetGasUsed() int64`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *Block) GetGasUsedOk() (*int64, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *Block) SetGasUsed(v int64)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *Block) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.

### SetGasUsedNil

`func (o *Block) SetGasUsedNil(b bool)`

 SetGasUsedNil sets the value for GasUsed to be an explicit nil

### UnsetGasUsed
`func (o *Block) UnsetGasUsed()`

UnsetGasUsed ensures that no value is present for GasUsed, not even an explicit nil
### GetHapiVersion

`func (o *Block) GetHapiVersion() string`

GetHapiVersion returns the HapiVersion field if non-nil, zero value otherwise.

### GetHapiVersionOk

`func (o *Block) GetHapiVersionOk() (*string, bool)`

GetHapiVersionOk returns a tuple with the HapiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHapiVersion

`func (o *Block) SetHapiVersion(v string)`

SetHapiVersion sets HapiVersion field to given value.

### HasHapiVersion

`func (o *Block) HasHapiVersion() bool`

HasHapiVersion returns a boolean if a field has been set.

### SetHapiVersionNil

`func (o *Block) SetHapiVersionNil(b bool)`

 SetHapiVersionNil sets the value for HapiVersion to be an explicit nil

### UnsetHapiVersion
`func (o *Block) UnsetHapiVersion()`

UnsetHapiVersion ensures that no value is present for HapiVersion, not even an explicit nil
### GetHash

`func (o *Block) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Block) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Block) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Block) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetLogsBloom

`func (o *Block) GetLogsBloom() string`

GetLogsBloom returns the LogsBloom field if non-nil, zero value otherwise.

### GetLogsBloomOk

`func (o *Block) GetLogsBloomOk() (*string, bool)`

GetLogsBloomOk returns a tuple with the LogsBloom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogsBloom

`func (o *Block) SetLogsBloom(v string)`

SetLogsBloom sets LogsBloom field to given value.

### HasLogsBloom

`func (o *Block) HasLogsBloom() bool`

HasLogsBloom returns a boolean if a field has been set.

### SetLogsBloomNil

`func (o *Block) SetLogsBloomNil(b bool)`

 SetLogsBloomNil sets the value for LogsBloom to be an explicit nil

### UnsetLogsBloom
`func (o *Block) UnsetLogsBloom()`

UnsetLogsBloom ensures that no value is present for LogsBloom, not even an explicit nil
### GetName

`func (o *Block) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Block) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Block) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Block) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumber

`func (o *Block) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Block) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Block) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Block) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPreviousHash

`func (o *Block) GetPreviousHash() string`

GetPreviousHash returns the PreviousHash field if non-nil, zero value otherwise.

### GetPreviousHashOk

`func (o *Block) GetPreviousHashOk() (*string, bool)`

GetPreviousHashOk returns a tuple with the PreviousHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousHash

`func (o *Block) SetPreviousHash(v string)`

SetPreviousHash sets PreviousHash field to given value.

### HasPreviousHash

`func (o *Block) HasPreviousHash() bool`

HasPreviousHash returns a boolean if a field has been set.

### GetSize

`func (o *Block) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Block) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Block) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Block) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *Block) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Block) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetTimestamp

`func (o *Block) GetTimestamp() BlockTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Block) GetTimestampOk() (*BlockTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Block) SetTimestamp(v BlockTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Block) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


