package mirror

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const MAINNET_ADDRESS = "https://mainnet-public.mirrornode.hedera.com"
const TESTNET_ADDRESS = "https://testnet.mirrornode.hedera.com"
const NETWORK_NODES_ROUTE = "api/v1/network/nodes"

type StakingPeriod struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Timestamp struct {
	From string  `json:"from"`
	To   *string `json:"to"`
}

type ServiceEndpoint struct {
	IPAddressV4 string `json:"ip_address_v4"`
	Port        int    `json:"port"`
}

type Node struct {
	Description      string            `json:"description"`
	FileID           string            `json:"file_id"`
	MaxStake         int               `json:"max_stake"`
	Memo             string            `json:"memo"`
	MinStake         int               `json:"min_stake"`
	NodeAccountID    string            `json:"node_account_id"`
	NodeCertHash     string            `json:"node_cert_hash"`
	NodeID           int               `json:"node_id"`
	PublicKey        string            `json:"public_key"`
	RewardRateStart  int               `json:"reward_rate_start"`
	ServiceEndpoints []ServiceEndpoint `json:"service_endpoints"`
	Stake            int               `json:"stake"`
	StakeNotRewarded int               `json:"stake_not_rewarded"`
	StakeRewarded    int               `json:"stake_rewarded"`
	StakingPeriod    StakingPeriod     `json:"staking_period"`
	Timestamp        Timestamp         `json:"timestamp"`
}

type Links struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type AddressBook struct {
	Nodes []Node `json:"nodes"`
	Links Links  `json:"links"`
}

func FetchAddressBook(ctx context.Context) (*AddressBook, error) {
	requestURL := fmt.Sprintf("%s/%s", MAINNET_ADDRESS, NETWORK_NODES_ROUTE)
	resp, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("fetching address book: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading address book response body: %w", err)
	}

	var addressBook AddressBook
	if err := json.Unmarshal(body, &addressBook); err != nil {
		return nil, fmt.Errorf("parsing address book response body: %w", err)
	}

	return &addressBook, nil
}
