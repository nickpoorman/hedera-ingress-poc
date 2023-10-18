package hapi

import (
	"context"
	"fmt"

	"github.com/nickpoorman/hedera-ingress/proto"
	pb "google.golang.org/protobuf/proto"
)

type HapiServer struct {
	nftBasedLimits bool
	// accountsCache  *groupcache.Group
}

func NewHapiServer(nftBasedLimits bool) *HapiServer {
	h := &HapiServer{
		nftBasedLimits: nftBasedLimits,
	}

	// Create a group cache for account information.
	// h.accountsCache = groupcache.NewGroup("accounts", 1<<30,
	// 	groupcache.GetterFunc(h.getAccountFromCache),
	// )

	return h
}

// func (h HapiServer) getAccountFromCache(ctx groupcache.Context, key string, dest groupcache.Sink) error {
// 	// When a cache miss occurs, generate the greeting
// 	result := fmt.Sprintf("Hello, %s!", key)
// 	return dest.SetString(result)
// }

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (h HapiServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// TODO: Implement this

	// TODO: Validate signed transaction is valid.
	if _, err := getSignedTransaction(tx); err != nil {
		// Return the same error that Services does when it encounters invalid SignedTransaction bytes.
		return fmt.Errorf("invaild signed transaction bytes: %w", err)
	}

	// Verify the payer account hasn't gone over the limits.
	if h.nftBasedLimits {
		if err := checkNFTBasedLimits(ctx, tx); err != nil {
			return fmt.Errorf("failed nft based limits check: %w", err)
		}
	}

	return nil
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (h HapiServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// TODO: Implement this
	return nil, nil
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (h HapiServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// TODO: Implement this
	return nil
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (h HapiServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// TODO: Implement this
	return nil, nil
}

func checkNFTBasedLimits(ctx context.Context, tx *proto.Transaction) error {
	// TODO: Implement this We'll need to validate a few things:
	//
	// 1. Does this account have one of the testnet NFTs that let them do
	// things?
	//
	// -- It's probably more efficient to periodically poll the mirror node for
	// tokens/listTokenBalancesById and cache the results since there are a
	// small number of NFTs that will allow this.
	//
	// 2. If they have the account create NFT, how many accounts have they
	// already created?
	//
	// -- This is tricky because we need to go fetch a payer account and then get all the transactions where it has created an account.
	// -- transactions/listTransactions to get all the CRYPTOCREATEACCOUNT transactions for an account.
	// -- Optomization: We'll probably want to poll this regurlarly and cache the results for any active accounts to keep things fast.
	//
	// 3. If they have the NFT create NFT, how many NFTs have they already
	// created?
	//
	// -- Same as how many accounts they have created.
	return nil
}

func getSignedTransaction(tx *proto.Transaction) (*proto.SignedTransaction, error) {
	var stx proto.SignedTransaction
	return &stx, pb.Unmarshal(tx.GetSignedTransactionBytes(), &stx)
}
