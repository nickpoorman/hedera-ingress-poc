// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.4
// source: get_account_details.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// Gets all the information about an account, including balance and allowances. This does not get the list of
// account records.
type GetAccountDetailsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Account details sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// The account ID for which information is requested
	AccountId *AccountID `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *GetAccountDetailsQuery) Reset() {
	*x = GetAccountDetailsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_account_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDetailsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDetailsQuery) ProtoMessage() {}

func (x *GetAccountDetailsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_get_account_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDetailsQuery.ProtoReflect.Descriptor instead.
func (*GetAccountDetailsQuery) Descriptor() ([]byte, []int) {
	return file_get_account_details_proto_rawDescGZIP(), []int{0}
}

func (x *GetAccountDetailsQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetAccountDetailsQuery) GetAccountId() *AccountID {
	if x != nil {
		return x.AccountId
	}
	return nil
}

// *
// Response when the client sends the node GetAccountDetailsQuery
type GetAccountDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// *
	// Details of the account (a state proof can be generated for this)
	AccountDetails *GetAccountDetailsResponse_AccountDetails `protobuf:"bytes,2,opt,name=account_details,json=accountDetails,proto3" json:"account_details,omitempty"`
}

func (x *GetAccountDetailsResponse) Reset() {
	*x = GetAccountDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_account_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDetailsResponse) ProtoMessage() {}

func (x *GetAccountDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_account_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetAccountDetailsResponse) Descriptor() ([]byte, []int) {
	return file_get_account_details_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountDetailsResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *GetAccountDetailsResponse) GetAccountDetails() *GetAccountDetailsResponse_AccountDetails {
	if x != nil {
		return x.AccountDetails
	}
	return nil
}

// *
// A granted allowance of hbar transfers for a spender relative to the owner account.
type GrantedCryptoAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The account ID of the spender of the hbar allowance.
	Spender *AccountID `protobuf:"bytes,1,opt,name=spender,proto3" json:"spender,omitempty"`
	// *
	// The amount of the spender's allowance in tinybars.
	Amount int64 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GrantedCryptoAllowance) Reset() {
	*x = GrantedCryptoAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_account_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantedCryptoAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantedCryptoAllowance) ProtoMessage() {}

func (x *GrantedCryptoAllowance) ProtoReflect() protoreflect.Message {
	mi := &file_get_account_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantedCryptoAllowance.ProtoReflect.Descriptor instead.
func (*GrantedCryptoAllowance) Descriptor() ([]byte, []int) {
	return file_get_account_details_proto_rawDescGZIP(), []int{2}
}

func (x *GrantedCryptoAllowance) GetSpender() *AccountID {
	if x != nil {
		return x.Spender
	}
	return nil
}

func (x *GrantedCryptoAllowance) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// *
// A granted allowance for all the NFTs of a token for a spender relative to the owner account.
type GrantedNftAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The token that the allowance pertains to.
	TokenId *TokenID `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// *
	// The account ID of the spender that has been granted access to all NFTs of the owner
	Spender *AccountID `protobuf:"bytes,2,opt,name=spender,proto3" json:"spender,omitempty"`
}

func (x *GrantedNftAllowance) Reset() {
	*x = GrantedNftAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_account_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantedNftAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantedNftAllowance) ProtoMessage() {}

func (x *GrantedNftAllowance) ProtoReflect() protoreflect.Message {
	mi := &file_get_account_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantedNftAllowance.ProtoReflect.Descriptor instead.
func (*GrantedNftAllowance) Descriptor() ([]byte, []int) {
	return file_get_account_details_proto_rawDescGZIP(), []int{3}
}

func (x *GrantedNftAllowance) GetTokenId() *TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *GrantedNftAllowance) GetSpender() *AccountID {
	if x != nil {
		return x.Spender
	}
	return nil
}

// *
// A granted allowance of fungible token transfers for a spender relative to the owner account.
type GrantedTokenAllowance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The token that the allowance pertains to.
	TokenId *TokenID `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	// *
	// The account ID of the token allowance spender.
	Spender *AccountID `protobuf:"bytes,2,opt,name=spender,proto3" json:"spender,omitempty"`
	// *
	// The amount of the spender's token allowance.
	Amount int64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GrantedTokenAllowance) Reset() {
	*x = GrantedTokenAllowance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_account_details_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantedTokenAllowance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantedTokenAllowance) ProtoMessage() {}

func (x *GrantedTokenAllowance) ProtoReflect() protoreflect.Message {
	mi := &file_get_account_details_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantedTokenAllowance.ProtoReflect.Descriptor instead.
func (*GrantedTokenAllowance) Descriptor() ([]byte, []int) {
	return file_get_account_details_proto_rawDescGZIP(), []int{4}
}

func (x *GrantedTokenAllowance) GetTokenId() *TokenID {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *GrantedTokenAllowance) GetSpender() *AccountID {
	if x != nil {
		return x.Spender
	}
	return nil
}

func (x *GrantedTokenAllowance) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GetAccountDetailsResponse_AccountDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The account ID for which this information applies
	AccountId *AccountID `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// *
	// The Contract Account ID comprising of both the contract instance and the cryptocurrency
	// account owned by the contract instance, in the format used by Solidity
	ContractAccountId string `protobuf:"bytes,2,opt,name=contract_account_id,json=contractAccountId,proto3" json:"contract_account_id,omitempty"`
	// *
	// If true, then this account has been deleted, it will disappear when it expires, and all
	// transactions for it will fail except the transaction to extend its expiration date
	Deleted bool `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// *
	// [Deprecated] The Account ID of the account to which this is proxy staked. If proxyAccountID is null,
	// or is an invalid account, or is an account that isn't a node, then this account is
	// automatically proxy staked to a node chosen by the network, but without earning payments.
	// If the proxyAccountID account refuses to accept proxy staking , or if it is not currently
	// running a node, then it will behave as if proxyAccountID was null.
	//
	// Deprecated: Do not use.
	ProxyAccountId *AccountID `protobuf:"bytes,4,opt,name=proxy_account_id,json=proxyAccountId,proto3" json:"proxy_account_id,omitempty"`
	// *
	// The total number of tinybars proxy staked to this account
	ProxyReceived int64 `protobuf:"varint,5,opt,name=proxy_received,json=proxyReceived,proto3" json:"proxy_received,omitempty"`
	// *
	// The key for the account, which must sign in order to transfer out, or to modify the
	// account in any way other than extending its expiration date.
	Key *Key `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	// *
	// The current balance of account in tinybars
	Balance uint64 `protobuf:"varint,7,opt,name=balance,proto3" json:"balance,omitempty"`
	// *
	// If true, no transaction can transfer to this account unless signed by this account's key
	ReceiverSigRequired bool `protobuf:"varint,8,opt,name=receiver_sig_required,json=receiverSigRequired,proto3" json:"receiver_sig_required,omitempty"`
	// *
	// The TimeStamp time at which this account is set to expire
	ExpirationTime *Timestamp `protobuf:"bytes,9,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	// *
	// The duration for expiration time will extend every this many seconds. If there are
	// insufficient funds, then it extends as long as possible. If it is empty when it expires,
	// then it is deleted.
	AutoRenewPeriod *Duration `protobuf:"bytes,10,opt,name=auto_renew_period,json=autoRenewPeriod,proto3" json:"auto_renew_period,omitempty"`
	// *
	// All tokens related to this account
	TokenRelationships []*TokenRelationship `protobuf:"bytes,11,rep,name=token_relationships,json=tokenRelationships,proto3" json:"token_relationships,omitempty"`
	// *
	// The memo associated with the account
	Memo string `protobuf:"bytes,12,opt,name=memo,proto3" json:"memo,omitempty"`
	// *
	// The number of NFTs owned by this account
	OwnedNfts int64 `protobuf:"varint,13,opt,name=owned_nfts,json=ownedNfts,proto3" json:"owned_nfts,omitempty"`
	// *
	// The maximum number of tokens that an Account can be implicitly associated with.
	MaxAutomaticTokenAssociations int32 `protobuf:"varint,14,opt,name=max_automatic_token_associations,json=maxAutomaticTokenAssociations,proto3" json:"max_automatic_token_associations,omitempty"`
	// *
	// The alias of this account
	Alias []byte `protobuf:"bytes,15,opt,name=alias,proto3" json:"alias,omitempty"`
	// *
	// The ledger ID the response was returned from; please see <a href="https://github.com/hashgraph/hedera-improvement-proposal/blob/master/HIP/hip-198.md">HIP-198</a> for the network-specific IDs.
	LedgerId []byte `protobuf:"bytes,16,opt,name=ledger_id,json=ledgerId,proto3" json:"ledger_id,omitempty"`
	// *
	// All of the hbar allowances approved by the account owner.
	GrantedCryptoAllowances []*GrantedCryptoAllowance `protobuf:"bytes,17,rep,name=granted_crypto_allowances,json=grantedCryptoAllowances,proto3" json:"granted_crypto_allowances,omitempty"`
	// *
	// All of the non-fungible token allowances approved by the account owner.
	GrantedNftAllowances []*GrantedNftAllowance `protobuf:"bytes,18,rep,name=granted_nft_allowances,json=grantedNftAllowances,proto3" json:"granted_nft_allowances,omitempty"`
	// *
	// All of the fungible token allowances approved by the account owner.
	GrantedTokenAllowances []*GrantedTokenAllowance `protobuf:"bytes,19,rep,name=granted_token_allowances,json=grantedTokenAllowances,proto3" json:"granted_token_allowances,omitempty"`
}

func (x *GetAccountDetailsResponse_AccountDetails) Reset() {
	*x = GetAccountDetailsResponse_AccountDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_account_details_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountDetailsResponse_AccountDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountDetailsResponse_AccountDetails) ProtoMessage() {}

func (x *GetAccountDetailsResponse_AccountDetails) ProtoReflect() protoreflect.Message {
	mi := &file_get_account_details_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountDetailsResponse_AccountDetails.ProtoReflect.Descriptor instead.
func (*GetAccountDetailsResponse_AccountDetails) Descriptor() ([]byte, []int) {
	return file_get_account_details_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetAccountDetailsResponse_AccountDetails) GetAccountId() *AccountID {
	if x != nil {
		return x.AccountId
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetContractAccountId() string {
	if x != nil {
		return x.ContractAccountId
	}
	return ""
}

func (x *GetAccountDetailsResponse_AccountDetails) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

// Deprecated: Do not use.
func (x *GetAccountDetailsResponse_AccountDetails) GetProxyAccountId() *AccountID {
	if x != nil {
		return x.ProxyAccountId
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetProxyReceived() int64 {
	if x != nil {
		return x.ProxyReceived
	}
	return 0
}

func (x *GetAccountDetailsResponse_AccountDetails) GetKey() *Key {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *GetAccountDetailsResponse_AccountDetails) GetReceiverSigRequired() bool {
	if x != nil {
		return x.ReceiverSigRequired
	}
	return false
}

func (x *GetAccountDetailsResponse_AccountDetails) GetExpirationTime() *Timestamp {
	if x != nil {
		return x.ExpirationTime
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetAutoRenewPeriod() *Duration {
	if x != nil {
		return x.AutoRenewPeriod
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetTokenRelationships() []*TokenRelationship {
	if x != nil {
		return x.TokenRelationships
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *GetAccountDetailsResponse_AccountDetails) GetOwnedNfts() int64 {
	if x != nil {
		return x.OwnedNfts
	}
	return 0
}

func (x *GetAccountDetailsResponse_AccountDetails) GetMaxAutomaticTokenAssociations() int32 {
	if x != nil {
		return x.MaxAutomaticTokenAssociations
	}
	return 0
}

func (x *GetAccountDetailsResponse_AccountDetails) GetAlias() []byte {
	if x != nil {
		return x.Alias
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetLedgerId() []byte {
	if x != nil {
		return x.LedgerId
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetGrantedCryptoAllowances() []*GrantedCryptoAllowance {
	if x != nil {
		return x.GrantedCryptoAllowances
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetGrantedNftAllowances() []*GrantedNftAllowance {
	if x != nil {
		return x.GrantedNftAllowances
	}
	return nil
}

func (x *GetAccountDetailsResponse_AccountDetails) GetGrantedTokenAllowances() []*GrantedTokenAllowance {
	if x != nil {
		return x.GrantedTokenAllowances
	}
	return nil
}

var File_get_account_details_proto protoreflect.FileDescriptor

var file_get_account_details_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x6c, 0x69,
	0x76, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0xfc, 0x08, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x58, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0xd5, 0x07, 0x0a, 0x0e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2f,
	0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0f,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x5f,
	0x72, 0x65, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x49, 0x0a, 0x13, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x12, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d,
	0x65, 0x6d, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x5f, 0x6e, 0x66, 0x74,
	0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x4e, 0x66,
	0x74, 0x73, 0x12, 0x47, 0x0a, 0x20, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x1d, 0x6d, 0x61,
	0x78, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x59,
	0x0a, 0x19, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x17, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x16, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x66, 0x74, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x4e, 0x66, 0x74, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x14, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x4e, 0x66,
	0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x18, 0x67,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x13, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x16, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x22, 0x5c, 0x0a, 0x16, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a,
	0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x6c, 0x0a, 0x13, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x4e, 0x66, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22,
	0x86, 0x01, 0x0a, 0x15, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x07, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x07, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_get_account_details_proto_rawDescOnce sync.Once
	file_get_account_details_proto_rawDescData = file_get_account_details_proto_rawDesc
)

func file_get_account_details_proto_rawDescGZIP() []byte {
	file_get_account_details_proto_rawDescOnce.Do(func() {
		file_get_account_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_account_details_proto_rawDescData)
	})
	return file_get_account_details_proto_rawDescData
}

var file_get_account_details_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_get_account_details_proto_goTypes = []interface{}{
	(*GetAccountDetailsQuery)(nil),                   // 0: proto.GetAccountDetailsQuery
	(*GetAccountDetailsResponse)(nil),                // 1: proto.GetAccountDetailsResponse
	(*GrantedCryptoAllowance)(nil),                   // 2: proto.GrantedCryptoAllowance
	(*GrantedNftAllowance)(nil),                      // 3: proto.GrantedNftAllowance
	(*GrantedTokenAllowance)(nil),                    // 4: proto.GrantedTokenAllowance
	(*GetAccountDetailsResponse_AccountDetails)(nil), // 5: proto.GetAccountDetailsResponse.AccountDetails
	(*QueryHeader)(nil),                              // 6: proto.QueryHeader
	(*AccountID)(nil),                                // 7: proto.AccountID
	(*ResponseHeader)(nil),                           // 8: proto.ResponseHeader
	(*TokenID)(nil),                                  // 9: proto.TokenID
	(*Key)(nil),                                      // 10: proto.Key
	(*Timestamp)(nil),                                // 11: proto.Timestamp
	(*Duration)(nil),                                 // 12: proto.Duration
	(*TokenRelationship)(nil),                        // 13: proto.TokenRelationship
}
var file_get_account_details_proto_depIdxs = []int32{
	6,  // 0: proto.GetAccountDetailsQuery.header:type_name -> proto.QueryHeader
	7,  // 1: proto.GetAccountDetailsQuery.account_id:type_name -> proto.AccountID
	8,  // 2: proto.GetAccountDetailsResponse.header:type_name -> proto.ResponseHeader
	5,  // 3: proto.GetAccountDetailsResponse.account_details:type_name -> proto.GetAccountDetailsResponse.AccountDetails
	7,  // 4: proto.GrantedCryptoAllowance.spender:type_name -> proto.AccountID
	9,  // 5: proto.GrantedNftAllowance.token_id:type_name -> proto.TokenID
	7,  // 6: proto.GrantedNftAllowance.spender:type_name -> proto.AccountID
	9,  // 7: proto.GrantedTokenAllowance.token_id:type_name -> proto.TokenID
	7,  // 8: proto.GrantedTokenAllowance.spender:type_name -> proto.AccountID
	7,  // 9: proto.GetAccountDetailsResponse.AccountDetails.account_id:type_name -> proto.AccountID
	7,  // 10: proto.GetAccountDetailsResponse.AccountDetails.proxy_account_id:type_name -> proto.AccountID
	10, // 11: proto.GetAccountDetailsResponse.AccountDetails.key:type_name -> proto.Key
	11, // 12: proto.GetAccountDetailsResponse.AccountDetails.expiration_time:type_name -> proto.Timestamp
	12, // 13: proto.GetAccountDetailsResponse.AccountDetails.auto_renew_period:type_name -> proto.Duration
	13, // 14: proto.GetAccountDetailsResponse.AccountDetails.token_relationships:type_name -> proto.TokenRelationship
	2,  // 15: proto.GetAccountDetailsResponse.AccountDetails.granted_crypto_allowances:type_name -> proto.GrantedCryptoAllowance
	3,  // 16: proto.GetAccountDetailsResponse.AccountDetails.granted_nft_allowances:type_name -> proto.GrantedNftAllowance
	4,  // 17: proto.GetAccountDetailsResponse.AccountDetails.granted_token_allowances:type_name -> proto.GrantedTokenAllowance
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_get_account_details_proto_init() }
func file_get_account_details_proto_init() {
	if File_get_account_details_proto != nil {
		return
	}
	file_timestamp_proto_init()
	file_duration_proto_init()
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	file_crypto_add_live_hash_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_get_account_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountDetailsQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_account_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountDetailsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_account_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantedCryptoAllowance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_account_details_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantedNftAllowance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_account_details_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantedTokenAllowance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_get_account_details_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountDetailsResponse_AccountDetails); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_get_account_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_account_details_proto_goTypes,
		DependencyIndexes: file_get_account_details_proto_depIdxs,
		MessageInfos:      file_get_account_details_proto_msgTypes,
	}.Build()
	File_get_account_details_proto = out.File
	file_get_account_details_proto_rawDesc = nil
	file_get_account_details_proto_goTypes = nil
	file_get_account_details_proto_depIdxs = nil
}
