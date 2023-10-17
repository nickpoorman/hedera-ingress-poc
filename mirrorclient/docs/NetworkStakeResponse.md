# NetworkStakeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxStakeRewarded** | **int64** | The maximum amount of tinybar that can be staked for reward while still achieving the maximum per-hbar reward rate  | 
**MaxStakingRewardRatePerHbar** | **int64** | The maximum reward rate, in tinybars per whole hbar, that any account can receive in a day | 
**MaxTotalReward** | **int64** | The total tinybars to be paid as staking rewards in the ending period, after applying the settings for the 0.0.800 balance threshold and the maximum stake rewarded  | 
**NodeRewardFeeFraction** | **float32** | The fraction between zero and one of the network and service fees paid to the node reward account 0.0.801 | 
**ReservedStakingRewards** | **int64** | The amount of the staking reward funds of account 0.0.800 reserved to pay pending rewards that have been earned but not collected  | 
**RewardBalanceThreshold** | **int64** | The unreserved tinybar balance of account 0.0.800 required to achieve the maximum per-hbar reward rate  | 
**StakeTotal** | **int64** | The total amount staked to the network in tinybars the start of the current staking period | 
**StakingPeriod** | [**NetworkStakeResponseStakingPeriod**](NetworkStakeResponseStakingPeriod.md) |  | 
**StakingPeriodDuration** | **int64** | The number of minutes in a staking period | 
**StakingPeriodsStored** | **int64** | The number of staking periods for which the reward is stored for each node | 
**StakingRewardFeeFraction** | **float32** | The fraction between zero and one of the network and service fees paid to the staking reward account 0.0.800 | 
**StakingRewardRate** | **int64** | The total number of tinybars to be distributed as staking rewards each period | 
**StakingStartThreshold** | **int64** | The minimum balance of staking reward account 0.0.800 required to active rewards | 
**UnreservedStakingRewardBalance** | **int64** | The unreserved balance of account 0.0.800 at the close of the just-ending period; this value is used to compute the HIP-782 balance ratio  | 

## Methods

### NewNetworkStakeResponse

`func NewNetworkStakeResponse(maxStakeRewarded int64, maxStakingRewardRatePerHbar int64, maxTotalReward int64, nodeRewardFeeFraction float32, reservedStakingRewards int64, rewardBalanceThreshold int64, stakeTotal int64, stakingPeriod NetworkStakeResponseStakingPeriod, stakingPeriodDuration int64, stakingPeriodsStored int64, stakingRewardFeeFraction float32, stakingRewardRate int64, stakingStartThreshold int64, unreservedStakingRewardBalance int64, ) *NetworkStakeResponse`

NewNetworkStakeResponse instantiates a new NetworkStakeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkStakeResponseWithDefaults

`func NewNetworkStakeResponseWithDefaults() *NetworkStakeResponse`

NewNetworkStakeResponseWithDefaults instantiates a new NetworkStakeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxStakeRewarded

`func (o *NetworkStakeResponse) GetMaxStakeRewarded() int64`

GetMaxStakeRewarded returns the MaxStakeRewarded field if non-nil, zero value otherwise.

### GetMaxStakeRewardedOk

`func (o *NetworkStakeResponse) GetMaxStakeRewardedOk() (*int64, bool)`

GetMaxStakeRewardedOk returns a tuple with the MaxStakeRewarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStakeRewarded

`func (o *NetworkStakeResponse) SetMaxStakeRewarded(v int64)`

SetMaxStakeRewarded sets MaxStakeRewarded field to given value.


### GetMaxStakingRewardRatePerHbar

`func (o *NetworkStakeResponse) GetMaxStakingRewardRatePerHbar() int64`

GetMaxStakingRewardRatePerHbar returns the MaxStakingRewardRatePerHbar field if non-nil, zero value otherwise.

### GetMaxStakingRewardRatePerHbarOk

`func (o *NetworkStakeResponse) GetMaxStakingRewardRatePerHbarOk() (*int64, bool)`

GetMaxStakingRewardRatePerHbarOk returns a tuple with the MaxStakingRewardRatePerHbar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStakingRewardRatePerHbar

`func (o *NetworkStakeResponse) SetMaxStakingRewardRatePerHbar(v int64)`

SetMaxStakingRewardRatePerHbar sets MaxStakingRewardRatePerHbar field to given value.


### GetMaxTotalReward

`func (o *NetworkStakeResponse) GetMaxTotalReward() int64`

GetMaxTotalReward returns the MaxTotalReward field if non-nil, zero value otherwise.

### GetMaxTotalRewardOk

`func (o *NetworkStakeResponse) GetMaxTotalRewardOk() (*int64, bool)`

GetMaxTotalRewardOk returns a tuple with the MaxTotalReward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalReward

`func (o *NetworkStakeResponse) SetMaxTotalReward(v int64)`

SetMaxTotalReward sets MaxTotalReward field to given value.


### GetNodeRewardFeeFraction

`func (o *NetworkStakeResponse) GetNodeRewardFeeFraction() float32`

GetNodeRewardFeeFraction returns the NodeRewardFeeFraction field if non-nil, zero value otherwise.

### GetNodeRewardFeeFractionOk

`func (o *NetworkStakeResponse) GetNodeRewardFeeFractionOk() (*float32, bool)`

GetNodeRewardFeeFractionOk returns a tuple with the NodeRewardFeeFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRewardFeeFraction

`func (o *NetworkStakeResponse) SetNodeRewardFeeFraction(v float32)`

SetNodeRewardFeeFraction sets NodeRewardFeeFraction field to given value.


### GetReservedStakingRewards

`func (o *NetworkStakeResponse) GetReservedStakingRewards() int64`

GetReservedStakingRewards returns the ReservedStakingRewards field if non-nil, zero value otherwise.

### GetReservedStakingRewardsOk

`func (o *NetworkStakeResponse) GetReservedStakingRewardsOk() (*int64, bool)`

GetReservedStakingRewardsOk returns a tuple with the ReservedStakingRewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStakingRewards

`func (o *NetworkStakeResponse) SetReservedStakingRewards(v int64)`

SetReservedStakingRewards sets ReservedStakingRewards field to given value.


### GetRewardBalanceThreshold

`func (o *NetworkStakeResponse) GetRewardBalanceThreshold() int64`

GetRewardBalanceThreshold returns the RewardBalanceThreshold field if non-nil, zero value otherwise.

### GetRewardBalanceThresholdOk

`func (o *NetworkStakeResponse) GetRewardBalanceThresholdOk() (*int64, bool)`

GetRewardBalanceThresholdOk returns a tuple with the RewardBalanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardBalanceThreshold

`func (o *NetworkStakeResponse) SetRewardBalanceThreshold(v int64)`

SetRewardBalanceThreshold sets RewardBalanceThreshold field to given value.


### GetStakeTotal

`func (o *NetworkStakeResponse) GetStakeTotal() int64`

GetStakeTotal returns the StakeTotal field if non-nil, zero value otherwise.

### GetStakeTotalOk

`func (o *NetworkStakeResponse) GetStakeTotalOk() (*int64, bool)`

GetStakeTotalOk returns a tuple with the StakeTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeTotal

`func (o *NetworkStakeResponse) SetStakeTotal(v int64)`

SetStakeTotal sets StakeTotal field to given value.


### GetStakingPeriod

`func (o *NetworkStakeResponse) GetStakingPeriod() NetworkStakeResponseStakingPeriod`

GetStakingPeriod returns the StakingPeriod field if non-nil, zero value otherwise.

### GetStakingPeriodOk

`func (o *NetworkStakeResponse) GetStakingPeriodOk() (*NetworkStakeResponseStakingPeriod, bool)`

GetStakingPeriodOk returns a tuple with the StakingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingPeriod

`func (o *NetworkStakeResponse) SetStakingPeriod(v NetworkStakeResponseStakingPeriod)`

SetStakingPeriod sets StakingPeriod field to given value.


### GetStakingPeriodDuration

`func (o *NetworkStakeResponse) GetStakingPeriodDuration() int64`

GetStakingPeriodDuration returns the StakingPeriodDuration field if non-nil, zero value otherwise.

### GetStakingPeriodDurationOk

`func (o *NetworkStakeResponse) GetStakingPeriodDurationOk() (*int64, bool)`

GetStakingPeriodDurationOk returns a tuple with the StakingPeriodDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingPeriodDuration

`func (o *NetworkStakeResponse) SetStakingPeriodDuration(v int64)`

SetStakingPeriodDuration sets StakingPeriodDuration field to given value.


### GetStakingPeriodsStored

`func (o *NetworkStakeResponse) GetStakingPeriodsStored() int64`

GetStakingPeriodsStored returns the StakingPeriodsStored field if non-nil, zero value otherwise.

### GetStakingPeriodsStoredOk

`func (o *NetworkStakeResponse) GetStakingPeriodsStoredOk() (*int64, bool)`

GetStakingPeriodsStoredOk returns a tuple with the StakingPeriodsStored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingPeriodsStored

`func (o *NetworkStakeResponse) SetStakingPeriodsStored(v int64)`

SetStakingPeriodsStored sets StakingPeriodsStored field to given value.


### GetStakingRewardFeeFraction

`func (o *NetworkStakeResponse) GetStakingRewardFeeFraction() float32`

GetStakingRewardFeeFraction returns the StakingRewardFeeFraction field if non-nil, zero value otherwise.

### GetStakingRewardFeeFractionOk

`func (o *NetworkStakeResponse) GetStakingRewardFeeFractionOk() (*float32, bool)`

GetStakingRewardFeeFractionOk returns a tuple with the StakingRewardFeeFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingRewardFeeFraction

`func (o *NetworkStakeResponse) SetStakingRewardFeeFraction(v float32)`

SetStakingRewardFeeFraction sets StakingRewardFeeFraction field to given value.


### GetStakingRewardRate

`func (o *NetworkStakeResponse) GetStakingRewardRate() int64`

GetStakingRewardRate returns the StakingRewardRate field if non-nil, zero value otherwise.

### GetStakingRewardRateOk

`func (o *NetworkStakeResponse) GetStakingRewardRateOk() (*int64, bool)`

GetStakingRewardRateOk returns a tuple with the StakingRewardRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingRewardRate

`func (o *NetworkStakeResponse) SetStakingRewardRate(v int64)`

SetStakingRewardRate sets StakingRewardRate field to given value.


### GetStakingStartThreshold

`func (o *NetworkStakeResponse) GetStakingStartThreshold() int64`

GetStakingStartThreshold returns the StakingStartThreshold field if non-nil, zero value otherwise.

### GetStakingStartThresholdOk

`func (o *NetworkStakeResponse) GetStakingStartThresholdOk() (*int64, bool)`

GetStakingStartThresholdOk returns a tuple with the StakingStartThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakingStartThreshold

`func (o *NetworkStakeResponse) SetStakingStartThreshold(v int64)`

SetStakingStartThreshold sets StakingStartThreshold field to given value.


### GetUnreservedStakingRewardBalance

`func (o *NetworkStakeResponse) GetUnreservedStakingRewardBalance() int64`

GetUnreservedStakingRewardBalance returns the UnreservedStakingRewardBalance field if non-nil, zero value otherwise.

### GetUnreservedStakingRewardBalanceOk

`func (o *NetworkStakeResponse) GetUnreservedStakingRewardBalanceOk() (*int64, bool)`

GetUnreservedStakingRewardBalanceOk returns a tuple with the UnreservedStakingRewardBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreservedStakingRewardBalance

`func (o *NetworkStakeResponse) SetUnreservedStakingRewardBalance(v int64)`

SetUnreservedStakingRewardBalance sets UnreservedStakingRewardBalance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


