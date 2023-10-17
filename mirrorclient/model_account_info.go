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

// checks if the AccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountInfo{}

// AccountInfo struct for AccountInfo
type AccountInfo struct {
	// Network entity ID in the format of `shard.realm.num`
	Account NullableString `json:"account"`
	// RFC4648 no-padding base32 encoded account alias
	Alias NullableString `json:"alias"`
	AutoRenewPeriod NullableInt64 `json:"auto_renew_period"`
	Balance NullableBalance `json:"balance"`
	CreatedTimestamp NullableString `json:"created_timestamp"`
	// Whether the account declines receiving a staking reward
	DeclineReward bool `json:"decline_reward"`
	Deleted NullableBool `json:"deleted"`
	EthereumNonce NullableInt64 `json:"ethereum_nonce"`
	// A network entity encoded as an EVM address in hex.
	EvmAddress Nullable*os.File `json:"evm_address"`
	ExpiryTimestamp NullableString `json:"expiry_timestamp"`
	Key NullableKey `json:"key"`
	MaxAutomaticTokenAssociations NullableInt32 `json:"max_automatic_token_associations"`
	Memo NullableString `json:"memo"`
	// The pending reward in tinybars the account will receive in the next reward payout. Note the value is updated at the end of each staking period and there may be delay to reflect the changes in the past staking period. 
	PendingReward *int64 `json:"pending_reward,omitempty"`
	ReceiverSigRequired NullableBool `json:"receiver_sig_required"`
	StakedAccountId string `json:"staked_account_id"`
	// The id of the node to which this account is staking
	StakedNodeId NullableInt64 `json:"staked_node_id"`
	StakePeriodStart string `json:"stake_period_start"`
}

// NewAccountInfo instantiates a new AccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountInfo(account NullableString, alias NullableString, autoRenewPeriod NullableInt64, balance NullableBalance, createdTimestamp NullableString, declineReward bool, deleted NullableBool, ethereumNonce NullableInt64, evmAddress Nullable*os.File, expiryTimestamp NullableString, key NullableKey, maxAutomaticTokenAssociations NullableInt32, memo NullableString, receiverSigRequired NullableBool, stakedAccountId string, stakedNodeId NullableInt64, stakePeriodStart string) *AccountInfo {
	this := AccountInfo{}
	this.Account = account
	this.Alias = alias
	this.AutoRenewPeriod = autoRenewPeriod
	this.Balance = balance
	this.CreatedTimestamp = createdTimestamp
	this.DeclineReward = declineReward
	this.Deleted = deleted
	this.EthereumNonce = ethereumNonce
	this.EvmAddress = evmAddress
	this.ExpiryTimestamp = expiryTimestamp
	this.Key = key
	this.MaxAutomaticTokenAssociations = maxAutomaticTokenAssociations
	this.Memo = memo
	this.ReceiverSigRequired = receiverSigRequired
	this.StakedAccountId = stakedAccountId
	this.StakedNodeId = stakedNodeId
	this.StakePeriodStart = stakePeriodStart
	return &this
}

// NewAccountInfoWithDefaults instantiates a new AccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountInfoWithDefaults() *AccountInfo {
	this := AccountInfo{}
	return &this
}

// GetAccount returns the Account field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountInfo) GetAccount() string {
	if o == nil || o.Account.Get() == nil {
		var ret string
		return ret
	}

	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// SetAccount sets field value
func (o *AccountInfo) SetAccount(v string) {
	o.Account.Set(&v)
}

// GetAlias returns the Alias field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountInfo) GetAlias() string {
	if o == nil || o.Alias.Get() == nil {
		var ret string
		return ret
	}

	return *o.Alias.Get()
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Alias.Get(), o.Alias.IsSet()
}

// SetAlias sets field value
func (o *AccountInfo) SetAlias(v string) {
	o.Alias.Set(&v)
}

// GetAutoRenewPeriod returns the AutoRenewPeriod field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *AccountInfo) GetAutoRenewPeriod() int64 {
	if o == nil || o.AutoRenewPeriod.Get() == nil {
		var ret int64
		return ret
	}

	return *o.AutoRenewPeriod.Get()
}

// GetAutoRenewPeriodOk returns a tuple with the AutoRenewPeriod field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetAutoRenewPeriodOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoRenewPeriod.Get(), o.AutoRenewPeriod.IsSet()
}

// SetAutoRenewPeriod sets field value
func (o *AccountInfo) SetAutoRenewPeriod(v int64) {
	o.AutoRenewPeriod.Set(&v)
}

// GetBalance returns the Balance field value
// If the value is explicit nil, the zero value for Balance will be returned
func (o *AccountInfo) GetBalance() Balance {
	if o == nil || o.Balance.Get() == nil {
		var ret Balance
		return ret
	}

	return *o.Balance.Get()
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetBalanceOk() (*Balance, bool) {
	if o == nil {
		return nil, false
	}
	return o.Balance.Get(), o.Balance.IsSet()
}

// SetBalance sets field value
func (o *AccountInfo) SetBalance(v Balance) {
	o.Balance.Set(&v)
}

// GetCreatedTimestamp returns the CreatedTimestamp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountInfo) GetCreatedTimestamp() string {
	if o == nil || o.CreatedTimestamp.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedTimestamp.Get()
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetCreatedTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedTimestamp.Get(), o.CreatedTimestamp.IsSet()
}

// SetCreatedTimestamp sets field value
func (o *AccountInfo) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp.Set(&v)
}

// GetDeclineReward returns the DeclineReward field value
func (o *AccountInfo) GetDeclineReward() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DeclineReward
}

// GetDeclineRewardOk returns a tuple with the DeclineReward field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetDeclineRewardOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeclineReward, true
}

// SetDeclineReward sets field value
func (o *AccountInfo) SetDeclineReward(v bool) {
	o.DeclineReward = v
}

// GetDeleted returns the Deleted field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *AccountInfo) GetDeleted() bool {
	if o == nil || o.Deleted.Get() == nil {
		var ret bool
		return ret
	}

	return *o.Deleted.Get()
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Deleted.Get(), o.Deleted.IsSet()
}

// SetDeleted sets field value
func (o *AccountInfo) SetDeleted(v bool) {
	o.Deleted.Set(&v)
}

// GetEthereumNonce returns the EthereumNonce field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *AccountInfo) GetEthereumNonce() int64 {
	if o == nil || o.EthereumNonce.Get() == nil {
		var ret int64
		return ret
	}

	return *o.EthereumNonce.Get()
}

// GetEthereumNonceOk returns a tuple with the EthereumNonce field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetEthereumNonceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EthereumNonce.Get(), o.EthereumNonce.IsSet()
}

// SetEthereumNonce sets field value
func (o *AccountInfo) SetEthereumNonce(v int64) {
	o.EthereumNonce.Set(&v)
}

// GetEvmAddress returns the EvmAddress field value
// If the value is explicit nil, the zero value for *os.File will be returned
func (o *AccountInfo) GetEvmAddress() *os.File {
	if o == nil || o.EvmAddress.Get() == nil {
		var ret *os.File
		return ret
	}

	return *o.EvmAddress.Get()
}

// GetEvmAddressOk returns a tuple with the EvmAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetEvmAddressOk() (**os.File, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvmAddress.Get(), o.EvmAddress.IsSet()
}

// SetEvmAddress sets field value
func (o *AccountInfo) SetEvmAddress(v *os.File) {
	o.EvmAddress.Set(&v)
}

// GetExpiryTimestamp returns the ExpiryTimestamp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountInfo) GetExpiryTimestamp() string {
	if o == nil || o.ExpiryTimestamp.Get() == nil {
		var ret string
		return ret
	}

	return *o.ExpiryTimestamp.Get()
}

// GetExpiryTimestampOk returns a tuple with the ExpiryTimestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetExpiryTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiryTimestamp.Get(), o.ExpiryTimestamp.IsSet()
}

// SetExpiryTimestamp sets field value
func (o *AccountInfo) SetExpiryTimestamp(v string) {
	o.ExpiryTimestamp.Set(&v)
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for Key will be returned
func (o *AccountInfo) GetKey() Key {
	if o == nil || o.Key.Get() == nil {
		var ret Key
		return ret
	}

	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetKeyOk() (*Key, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// SetKey sets field value
func (o *AccountInfo) SetKey(v Key) {
	o.Key.Set(&v)
}

// GetMaxAutomaticTokenAssociations returns the MaxAutomaticTokenAssociations field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *AccountInfo) GetMaxAutomaticTokenAssociations() int32 {
	if o == nil || o.MaxAutomaticTokenAssociations.Get() == nil {
		var ret int32
		return ret
	}

	return *o.MaxAutomaticTokenAssociations.Get()
}

// GetMaxAutomaticTokenAssociationsOk returns a tuple with the MaxAutomaticTokenAssociations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetMaxAutomaticTokenAssociationsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAutomaticTokenAssociations.Get(), o.MaxAutomaticTokenAssociations.IsSet()
}

// SetMaxAutomaticTokenAssociations sets field value
func (o *AccountInfo) SetMaxAutomaticTokenAssociations(v int32) {
	o.MaxAutomaticTokenAssociations.Set(&v)
}

// GetMemo returns the Memo field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AccountInfo) GetMemo() string {
	if o == nil || o.Memo.Get() == nil {
		var ret string
		return ret
	}

	return *o.Memo.Get()
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memo.Get(), o.Memo.IsSet()
}

// SetMemo sets field value
func (o *AccountInfo) SetMemo(v string) {
	o.Memo.Set(&v)
}

// GetPendingReward returns the PendingReward field value if set, zero value otherwise.
func (o *AccountInfo) GetPendingReward() int64 {
	if o == nil || IsNil(o.PendingReward) {
		var ret int64
		return ret
	}
	return *o.PendingReward
}

// GetPendingRewardOk returns a tuple with the PendingReward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetPendingRewardOk() (*int64, bool) {
	if o == nil || IsNil(o.PendingReward) {
		return nil, false
	}
	return o.PendingReward, true
}

// HasPendingReward returns a boolean if a field has been set.
func (o *AccountInfo) HasPendingReward() bool {
	if o != nil && !IsNil(o.PendingReward) {
		return true
	}

	return false
}

// SetPendingReward gets a reference to the given int64 and assigns it to the PendingReward field.
func (o *AccountInfo) SetPendingReward(v int64) {
	o.PendingReward = &v
}

// GetReceiverSigRequired returns the ReceiverSigRequired field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *AccountInfo) GetReceiverSigRequired() bool {
	if o == nil || o.ReceiverSigRequired.Get() == nil {
		var ret bool
		return ret
	}

	return *o.ReceiverSigRequired.Get()
}

// GetReceiverSigRequiredOk returns a tuple with the ReceiverSigRequired field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetReceiverSigRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiverSigRequired.Get(), o.ReceiverSigRequired.IsSet()
}

// SetReceiverSigRequired sets field value
func (o *AccountInfo) SetReceiverSigRequired(v bool) {
	o.ReceiverSigRequired.Set(&v)
}

// GetStakedAccountId returns the StakedAccountId field value
func (o *AccountInfo) GetStakedAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakedAccountId
}

// GetStakedAccountIdOk returns a tuple with the StakedAccountId field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetStakedAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakedAccountId, true
}

// SetStakedAccountId sets field value
func (o *AccountInfo) SetStakedAccountId(v string) {
	o.StakedAccountId = v
}

// GetStakedNodeId returns the StakedNodeId field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *AccountInfo) GetStakedNodeId() int64 {
	if o == nil || o.StakedNodeId.Get() == nil {
		var ret int64
		return ret
	}

	return *o.StakedNodeId.Get()
}

// GetStakedNodeIdOk returns a tuple with the StakedNodeId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountInfo) GetStakedNodeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StakedNodeId.Get(), o.StakedNodeId.IsSet()
}

// SetStakedNodeId sets field value
func (o *AccountInfo) SetStakedNodeId(v int64) {
	o.StakedNodeId.Set(&v)
}

// GetStakePeriodStart returns the StakePeriodStart field value
func (o *AccountInfo) GetStakePeriodStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StakePeriodStart
}

// GetStakePeriodStartOk returns a tuple with the StakePeriodStart field value
// and a boolean to check if the value has been set.
func (o *AccountInfo) GetStakePeriodStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StakePeriodStart, true
}

// SetStakePeriodStart sets field value
func (o *AccountInfo) SetStakePeriodStart(v string) {
	o.StakePeriodStart = v
}

func (o AccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account.Get()
	toSerialize["alias"] = o.Alias.Get()
	toSerialize["auto_renew_period"] = o.AutoRenewPeriod.Get()
	toSerialize["balance"] = o.Balance.Get()
	toSerialize["created_timestamp"] = o.CreatedTimestamp.Get()
	toSerialize["decline_reward"] = o.DeclineReward
	toSerialize["deleted"] = o.Deleted.Get()
	toSerialize["ethereum_nonce"] = o.EthereumNonce.Get()
	toSerialize["evm_address"] = o.EvmAddress.Get()
	toSerialize["expiry_timestamp"] = o.ExpiryTimestamp.Get()
	toSerialize["key"] = o.Key.Get()
	toSerialize["max_automatic_token_associations"] = o.MaxAutomaticTokenAssociations.Get()
	toSerialize["memo"] = o.Memo.Get()
	if !IsNil(o.PendingReward) {
		toSerialize["pending_reward"] = o.PendingReward
	}
	toSerialize["receiver_sig_required"] = o.ReceiverSigRequired.Get()
	toSerialize["staked_account_id"] = o.StakedAccountId
	toSerialize["staked_node_id"] = o.StakedNodeId.Get()
	toSerialize["stake_period_start"] = o.StakePeriodStart
	return toSerialize, nil
}

type NullableAccountInfo struct {
	value *AccountInfo
	isSet bool
}

func (v NullableAccountInfo) Get() *AccountInfo {
	return v.value
}

func (v *NullableAccountInfo) Set(val *AccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountInfo(val *AccountInfo) *NullableAccountInfo {
	return &NullableAccountInfo{value: val, isSet: true}
}

func (v NullableAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

