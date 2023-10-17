/*
Hedera Mirror Node REST API

The Mirror Node REST API offers the ability to query cryptocurrency transactions and account information from a Hedera managed mirror node.  Base url: [/api/v1](/api/v1)  OpenAPI Spec: [/api/v1/docs/openapi.yml](/api/v1/docs/openapi.yml)

API version: 0.89.0
Contact: mirrornode@hedera.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mirrorclient

import (
	"encoding/json"
	"os"
)

// checks if the ContractLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractLog{}

// ContractLog struct for ContractLog
type ContractLog struct {
	// The hex encoded EVM address of the contract
	Address *string `json:"address,omitempty"`
	Bloom **os.File `json:"bloom,omitempty"`
	// Network entity ID in the format of `shard.realm.num`
	ContractId NullableString `json:"contract_id,omitempty"`
	// The hex encoded data of the contract log
	Data NullableString `json:"data,omitempty"`
	// The index of the contract log in the chain of logs for an execution
	Index *int32 `json:"index,omitempty"`
	// A list of hex encoded topics associated with this log event
	Topics []string `json:"topics,omitempty"`
	// The hex encoded block (record file chain) hash
	BlockHash *string `json:"block_hash,omitempty"`
	// The block height calculated as the number of record files starting from zero since network start.
	BlockNumber *int64 `json:"block_number,omitempty"`
	RootContractId *string `json:"root_contract_id,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	// A hex encoded transaction hash
	TransactionHash *string `json:"transaction_hash,omitempty"`
	// The position of the transaction in the block
	TransactionIndex NullableInt32 `json:"transaction_index,omitempty"`
}

// NewContractLog instantiates a new ContractLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractLog() *ContractLog {
	this := ContractLog{}
	return &this
}

// NewContractLogWithDefaults instantiates a new ContractLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractLogWithDefaults() *ContractLog {
	this := ContractLog{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ContractLog) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ContractLog) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ContractLog) SetAddress(v string) {
	o.Address = &v
}

// GetBloom returns the Bloom field value if set, zero value otherwise.
func (o *ContractLog) GetBloom() *os.File {
	if o == nil || IsNil(o.Bloom) {
		var ret *os.File
		return ret
	}
	return *o.Bloom
}

// GetBloomOk returns a tuple with the Bloom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetBloomOk() (**os.File, bool) {
	if o == nil || IsNil(o.Bloom) {
		return nil, false
	}
	return o.Bloom, true
}

// HasBloom returns a boolean if a field has been set.
func (o *ContractLog) HasBloom() bool {
	if o != nil && !IsNil(o.Bloom) {
		return true
	}

	return false
}

// SetBloom gets a reference to the given *os.File and assigns it to the Bloom field.
func (o *ContractLog) SetBloom(v *os.File) {
	o.Bloom = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContractLog) GetContractId() string {
	if o == nil || IsNil(o.ContractId.Get()) {
		var ret string
		return ret
	}
	return *o.ContractId.Get()
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractLog) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractId.Get(), o.ContractId.IsSet()
}

// HasContractId returns a boolean if a field has been set.
func (o *ContractLog) HasContractId() bool {
	if o != nil && o.ContractId.IsSet() {
		return true
	}

	return false
}

// SetContractId gets a reference to the given NullableString and assigns it to the ContractId field.
func (o *ContractLog) SetContractId(v string) {
	o.ContractId.Set(&v)
}
// SetContractIdNil sets the value for ContractId to be an explicit nil
func (o *ContractLog) SetContractIdNil() {
	o.ContractId.Set(nil)
}

// UnsetContractId ensures that no value is present for ContractId, not even an explicit nil
func (o *ContractLog) UnsetContractId() {
	o.ContractId.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContractLog) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractLog) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *ContractLog) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *ContractLog) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *ContractLog) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ContractLog) UnsetData() {
	o.Data.Unset()
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ContractLog) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ContractLog) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *ContractLog) SetIndex(v int32) {
	o.Index = &v
}

// GetTopics returns the Topics field value if set, zero value otherwise.
func (o *ContractLog) GetTopics() []string {
	if o == nil || IsNil(o.Topics) {
		var ret []string
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetTopicsOk() ([]string, bool) {
	if o == nil || IsNil(o.Topics) {
		return nil, false
	}
	return o.Topics, true
}

// HasTopics returns a boolean if a field has been set.
func (o *ContractLog) HasTopics() bool {
	if o != nil && !IsNil(o.Topics) {
		return true
	}

	return false
}

// SetTopics gets a reference to the given []string and assigns it to the Topics field.
func (o *ContractLog) SetTopics(v []string) {
	o.Topics = v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *ContractLog) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetBlockHashOk() (*string, bool) {
	if o == nil || IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *ContractLog) HasBlockHash() bool {
	if o != nil && !IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *ContractLog) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise.
func (o *ContractLog) GetBlockNumber() int64 {
	if o == nil || IsNil(o.BlockNumber) {
		var ret int64
		return ret
	}
	return *o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetBlockNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.BlockNumber) {
		return nil, false
	}
	return o.BlockNumber, true
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *ContractLog) HasBlockNumber() bool {
	if o != nil && !IsNil(o.BlockNumber) {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given int64 and assigns it to the BlockNumber field.
func (o *ContractLog) SetBlockNumber(v int64) {
	o.BlockNumber = &v
}

// GetRootContractId returns the RootContractId field value if set, zero value otherwise.
func (o *ContractLog) GetRootContractId() string {
	if o == nil || IsNil(o.RootContractId) {
		var ret string
		return ret
	}
	return *o.RootContractId
}

// GetRootContractIdOk returns a tuple with the RootContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetRootContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.RootContractId) {
		return nil, false
	}
	return o.RootContractId, true
}

// HasRootContractId returns a boolean if a field has been set.
func (o *ContractLog) HasRootContractId() bool {
	if o != nil && !IsNil(o.RootContractId) {
		return true
	}

	return false
}

// SetRootContractId gets a reference to the given string and assigns it to the RootContractId field.
func (o *ContractLog) SetRootContractId(v string) {
	o.RootContractId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ContractLog) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ContractLog) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *ContractLog) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *ContractLog) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractLog) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *ContractLog) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *ContractLog) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

// GetTransactionIndex returns the TransactionIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContractLog) GetTransactionIndex() int32 {
	if o == nil || IsNil(o.TransactionIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.TransactionIndex.Get()
}

// GetTransactionIndexOk returns a tuple with the TransactionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContractLog) GetTransactionIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionIndex.Get(), o.TransactionIndex.IsSet()
}

// HasTransactionIndex returns a boolean if a field has been set.
func (o *ContractLog) HasTransactionIndex() bool {
	if o != nil && o.TransactionIndex.IsSet() {
		return true
	}

	return false
}

// SetTransactionIndex gets a reference to the given NullableInt32 and assigns it to the TransactionIndex field.
func (o *ContractLog) SetTransactionIndex(v int32) {
	o.TransactionIndex.Set(&v)
}
// SetTransactionIndexNil sets the value for TransactionIndex to be an explicit nil
func (o *ContractLog) SetTransactionIndexNil() {
	o.TransactionIndex.Set(nil)
}

// UnsetTransactionIndex ensures that no value is present for TransactionIndex, not even an explicit nil
func (o *ContractLog) UnsetTransactionIndex() {
	o.TransactionIndex.Unset()
}

func (o ContractLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Bloom) {
		toSerialize["bloom"] = o.Bloom
	}
	if o.ContractId.IsSet() {
		toSerialize["contract_id"] = o.ContractId.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.Topics) {
		toSerialize["topics"] = o.Topics
	}
	if !IsNil(o.BlockHash) {
		toSerialize["block_hash"] = o.BlockHash
	}
	if !IsNil(o.BlockNumber) {
		toSerialize["block_number"] = o.BlockNumber
	}
	if !IsNil(o.RootContractId) {
		toSerialize["root_contract_id"] = o.RootContractId
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transaction_hash"] = o.TransactionHash
	}
	if o.TransactionIndex.IsSet() {
		toSerialize["transaction_index"] = o.TransactionIndex.Get()
	}
	return toSerialize, nil
}

type NullableContractLog struct {
	value *ContractLog
	isSet bool
}

func (v NullableContractLog) Get() *ContractLog {
	return v.value
}

func (v *NullableContractLog) Set(val *ContractLog) {
	v.value = val
	v.isSet = true
}

func (v NullableContractLog) IsSet() bool {
	return v.isSet
}

func (v *NullableContractLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractLog(val *ContractLog) *NullableContractLog {
	return &NullableContractLog{value: val, isSet: true}
}

func (v NullableContractLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


