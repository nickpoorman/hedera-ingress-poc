package hapi

import (
	"context"

	"github.com/nickpoorman/hedera-ingress/proto"
	"google.golang.org/grpc"
)

func (h *HapiServer) RegisterServices(gsrv *grpc.Server, backendGrpcConn *grpc.ClientConn) error {
	consensusServiceServer := &ConsensusServiceServer{
		backendClient: proto.NewConsensusServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterConsensusServiceServer(gsrv, consensusServiceServer)
	
	cryptoServiceServer := &CryptoServiceServer{
		backendClient: proto.NewCryptoServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterCryptoServiceServer(gsrv, cryptoServiceServer)
	
	fileServiceServer := &FileServiceServer{
		backendClient: proto.NewFileServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterFileServiceServer(gsrv, fileServiceServer)
	
	freezeServiceServer := &FreezeServiceServer{
		backendClient: proto.NewFreezeServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterFreezeServiceServer(gsrv, freezeServiceServer)
	
	networkServiceServer := &NetworkServiceServer{
		backendClient: proto.NewNetworkServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterNetworkServiceServer(gsrv, networkServiceServer)
	
	scheduleServiceServer := &ScheduleServiceServer{
		backendClient: proto.NewScheduleServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterScheduleServiceServer(gsrv, scheduleServiceServer)
	
	smartContractServiceServer := &SmartContractServiceServer{
		backendClient: proto.NewSmartContractServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterSmartContractServiceServer(gsrv, smartContractServiceServer)
	
	tokenServiceServer := &TokenServiceServer{
		backendClient: proto.NewTokenServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterTokenServiceServer(gsrv, tokenServiceServer)
	
	utilServiceServer := &UtilServiceServer{
		backendClient: proto.NewUtilServiceClient(backendGrpcConn),
		hapiServer:    h,
	}
	proto.RegisterUtilServiceServer(gsrv, utilServiceServer)
	
	return nil
}

type ConsensusServiceServer struct {
	proto.UnimplementedConsensusServiceServer
	backendClient proto.ConsensusServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *ConsensusServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *ConsensusServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *ConsensusServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *ConsensusServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *ConsensusServiceServer) CreateTopic(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CreateTopic(ctx, tx)
}

func (s *ConsensusServiceServer) UpdateTopic(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UpdateTopic(ctx, tx)
}

func (s *ConsensusServiceServer) DeleteTopic(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.DeleteTopic(ctx, tx)
}

func (s *ConsensusServiceServer) GetTopicInfo(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetTopicInfo(ctx, query)
}

func (s *ConsensusServiceServer) SubmitMessage(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.SubmitMessage(ctx, tx)
}

type CryptoServiceServer struct {
	proto.UnimplementedCryptoServiceServer
	backendClient proto.CryptoServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *CryptoServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *CryptoServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *CryptoServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *CryptoServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *CryptoServiceServer) CreateAccount(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CreateAccount(ctx, tx)
}

func (s *CryptoServiceServer) UpdateAccount(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UpdateAccount(ctx, tx)
}

func (s *CryptoServiceServer) CryptoTransfer(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CryptoTransfer(ctx, tx)
}

func (s *CryptoServiceServer) CryptoDelete(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CryptoDelete(ctx, tx)
}

func (s *CryptoServiceServer) ApproveAllowances(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.ApproveAllowances(ctx, tx)
}

func (s *CryptoServiceServer) DeleteAllowances(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.DeleteAllowances(ctx, tx)
}

func (s *CryptoServiceServer) AddLiveHash(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.AddLiveHash(ctx, tx)
}

func (s *CryptoServiceServer) DeleteLiveHash(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.DeleteLiveHash(ctx, tx)
}

func (s *CryptoServiceServer) GetLiveHash(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetLiveHash(ctx, query)
}

func (s *CryptoServiceServer) GetAccountRecords(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetAccountRecords(ctx, query)
}

func (s *CryptoServiceServer) CryptoGetBalance(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.CryptoGetBalance(ctx, query)
}

func (s *CryptoServiceServer) GetAccountInfo(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetAccountInfo(ctx, query)
}

func (s *CryptoServiceServer) GetTransactionReceipts(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetTransactionReceipts(ctx, query)
}

func (s *CryptoServiceServer) GetFastTransactionRecord(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetFastTransactionRecord(ctx, query)
}

func (s *CryptoServiceServer) GetTxRecordByTxID(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetTxRecordByTxID(ctx, query)
}

func (s *CryptoServiceServer) GetStakersByAccountID(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetStakersByAccountID(ctx, query)
}

type FileServiceServer struct {
	proto.UnimplementedFileServiceServer
	backendClient proto.FileServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *FileServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *FileServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *FileServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *FileServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *FileServiceServer) CreateFile(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CreateFile(ctx, tx)
}

func (s *FileServiceServer) UpdateFile(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UpdateFile(ctx, tx)
}

func (s *FileServiceServer) DeleteFile(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.DeleteFile(ctx, tx)
}

func (s *FileServiceServer) AppendContent(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.AppendContent(ctx, tx)
}

func (s *FileServiceServer) GetFileContent(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetFileContent(ctx, query)
}

func (s *FileServiceServer) GetFileInfo(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetFileInfo(ctx, query)
}

func (s *FileServiceServer) SystemDelete(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.SystemDelete(ctx, tx)
}

func (s *FileServiceServer) SystemUndelete(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.SystemUndelete(ctx, tx)
}

type FreezeServiceServer struct {
	proto.UnimplementedFreezeServiceServer
	backendClient proto.FreezeServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *FreezeServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *FreezeServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *FreezeServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *FreezeServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *FreezeServiceServer) Freeze(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.Freeze(ctx, tx)
}

type NetworkServiceServer struct {
	proto.UnimplementedNetworkServiceServer
	backendClient proto.NetworkServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *NetworkServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *NetworkServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *NetworkServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *NetworkServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *NetworkServiceServer) GetVersionInfo(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetVersionInfo(ctx, query)
}

func (s *NetworkServiceServer) GetExecutionTime(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetExecutionTime(ctx, query)
}

func (s *NetworkServiceServer) UncheckedSubmit(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UncheckedSubmit(ctx, tx)
}

func (s *NetworkServiceServer) GetAccountDetails(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetAccountDetails(ctx, query)
}

type ScheduleServiceServer struct {
	proto.UnimplementedScheduleServiceServer
	backendClient proto.ScheduleServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *ScheduleServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *ScheduleServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *ScheduleServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *ScheduleServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *ScheduleServiceServer) CreateSchedule(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CreateSchedule(ctx, tx)
}

func (s *ScheduleServiceServer) SignSchedule(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.SignSchedule(ctx, tx)
}

func (s *ScheduleServiceServer) DeleteSchedule(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.DeleteSchedule(ctx, tx)
}

func (s *ScheduleServiceServer) GetScheduleInfo(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetScheduleInfo(ctx, query)
}

type SmartContractServiceServer struct {
	proto.UnimplementedSmartContractServiceServer
	backendClient proto.SmartContractServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *SmartContractServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *SmartContractServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *SmartContractServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *SmartContractServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *SmartContractServiceServer) CreateContract(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CreateContract(ctx, tx)
}

func (s *SmartContractServiceServer) UpdateContract(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UpdateContract(ctx, tx)
}

func (s *SmartContractServiceServer) ContractCallMethod(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.ContractCallMethod(ctx, tx)
}

func (s *SmartContractServiceServer) GetContractInfo(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetContractInfo(ctx, query)
}

func (s *SmartContractServiceServer) ContractCallLocalMethod(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.ContractCallLocalMethod(ctx, query)
}

func (s *SmartContractServiceServer) ContractGetBytecode(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.ContractGetBytecode(ctx, query)
}

func (s *SmartContractServiceServer) GetBySolidityID(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetBySolidityID(ctx, query)
}

func (s *SmartContractServiceServer) GetTxRecordByContractID(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetTxRecordByContractID(ctx, query)
}

func (s *SmartContractServiceServer) DeleteContract(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.DeleteContract(ctx, tx)
}

func (s *SmartContractServiceServer) SystemDelete(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.SystemDelete(ctx, tx)
}

func (s *SmartContractServiceServer) SystemUndelete(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.SystemUndelete(ctx, tx)
}

func (s *SmartContractServiceServer) CallEthereum(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CallEthereum(ctx, tx)
}

type TokenServiceServer struct {
	proto.UnimplementedTokenServiceServer
	backendClient proto.TokenServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *TokenServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *TokenServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *TokenServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *TokenServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *TokenServiceServer) CreateToken(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.CreateToken(ctx, tx)
}

func (s *TokenServiceServer) UpdateToken(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UpdateToken(ctx, tx)
}

func (s *TokenServiceServer) MintToken(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.MintToken(ctx, tx)
}

func (s *TokenServiceServer) BurnToken(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.BurnToken(ctx, tx)
}

func (s *TokenServiceServer) DeleteToken(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.DeleteToken(ctx, tx)
}

func (s *TokenServiceServer) WipeTokenAccount(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.WipeTokenAccount(ctx, tx)
}

func (s *TokenServiceServer) FreezeTokenAccount(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.FreezeTokenAccount(ctx, tx)
}

func (s *TokenServiceServer) UnfreezeTokenAccount(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UnfreezeTokenAccount(ctx, tx)
}

func (s *TokenServiceServer) GrantKycToTokenAccount(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.GrantKycToTokenAccount(ctx, tx)
}

func (s *TokenServiceServer) RevokeKycFromTokenAccount(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.RevokeKycFromTokenAccount(ctx, tx)
}

func (s *TokenServiceServer) AssociateTokens(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.AssociateTokens(ctx, tx)
}

func (s *TokenServiceServer) DissociateTokens(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.DissociateTokens(ctx, tx)
}

func (s *TokenServiceServer) UpdateTokenFeeSchedule(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UpdateTokenFeeSchedule(ctx, tx)
}

func (s *TokenServiceServer) GetTokenInfo(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetTokenInfo(ctx, query)
}

func (s *TokenServiceServer) GetAccountNftInfos(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetAccountNftInfos(ctx, query)
}

func (s *TokenServiceServer) GetTokenNftInfo(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetTokenNftInfo(ctx, query)
}

func (s *TokenServiceServer) GetTokenNftInfos(ctx context.Context, query *proto.Query) (*proto.Response, error) {
	if err := s.validateQuery(ctx, query); err != nil {
		return s.handleQueryValidationError(ctx, query, err)
	}
    return s.backendClient.GetTokenNftInfos(ctx, query)
}

func (s *TokenServiceServer) PauseToken(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.PauseToken(ctx, tx)
}

func (s *TokenServiceServer) UnpauseToken(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.UnpauseToken(ctx, tx)
}

type UtilServiceServer struct {
	proto.UnimplementedUtilServiceServer
	backendClient proto.UtilServiceClient
	hapiServer    *HapiServer
}

// validateTransaction takes a transaction and validates it. If the transaction
// is invalid, it returns an error.
func (s *UtilServiceServer) validateTransaction(ctx context.Context, tx *proto.Transaction) error {
	// For now delegates to the global validateTransaction function.
	return s.hapiServer.validateTransaction(ctx, tx)
}

// handleValidationError takes a transaction and an error and returns a response
// with same error message that would have been returned by the server.
func (s *UtilServiceServer) handleTransactionValidationError(ctx context.Context, tx *proto.Transaction, err error) (*proto.TransactionResponse, error) {
	// For now delegates to the global handleTransactionValidationError function.
	return s.hapiServer.handleTransactionValidationError(ctx, tx, err)
}

// validateQuery takes a query and validates it. If the query
// is invalid, it returns an error.
func (s *UtilServiceServer) validateQuery(ctx context.Context, query *proto.Query) error {
	// For now delegates to the global validateQuery function.
	return s.hapiServer.validateQuery(ctx, query)
}

// handleValidationError takes a query and an error and returns a response
// with same error message that would have been returned by the server.
func (s *UtilServiceServer) handleQueryValidationError(ctx context.Context, query *proto.Query, err error) (*proto.Response, error) {
	// For now delegates to the global handleQueryValidationError function.
	return s.hapiServer.handleQueryValidationError(ctx, query, err)
}

func (s *UtilServiceServer) Prng(ctx context.Context, tx *proto.Transaction) (*proto.TransactionResponse, error) {
	if err := s.validateTransaction(ctx, tx); err != nil {
		return s.handleTransactionValidationError(ctx, tx, err)
	}
    return s.backendClient.Prng(ctx, tx)
}
