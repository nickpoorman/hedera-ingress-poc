# StakingRewardsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rewards** | Pointer to [**[]StakingReward**](StakingReward.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewStakingRewardsResponse

`func NewStakingRewardsResponse() *StakingRewardsResponse`

NewStakingRewardsResponse instantiates a new StakingRewardsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStakingRewardsResponseWithDefaults

`func NewStakingRewardsResponseWithDefaults() *StakingRewardsResponse`

NewStakingRewardsResponseWithDefaults instantiates a new StakingRewardsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRewards

`func (o *StakingRewardsResponse) GetRewards() []StakingReward`

GetRewards returns the Rewards field if non-nil, zero value otherwise.

### GetRewardsOk

`func (o *StakingRewardsResponse) GetRewardsOk() (*[]StakingReward, bool)`

GetRewardsOk returns a tuple with the Rewards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewards

`func (o *StakingRewardsResponse) SetRewards(v []StakingReward)`

SetRewards sets Rewards field to given value.

### HasRewards

`func (o *StakingRewardsResponse) HasRewards() bool`

HasRewards returns a boolean if a field has been set.

### GetLinks

`func (o *StakingRewardsResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *StakingRewardsResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *StakingRewardsResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *StakingRewardsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


