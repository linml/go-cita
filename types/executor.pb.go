// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ReceiptError int32

const (
	// ExecutionError
	ReceiptError_NotEnoughBaseGas        ReceiptError = 0
	ReceiptError_BlockGasLimitReached    ReceiptError = 1
	ReceiptError_AccountGasLimitReached  ReceiptError = 2
	ReceiptError_InvalidTransactionNonce ReceiptError = 3
	ReceiptError_NotEnoughCash           ReceiptError = 4
	ReceiptError_NoTransactionPermission ReceiptError = 5
	ReceiptError_NoContractPermission    ReceiptError = 6
	ReceiptError_NoCallPermission        ReceiptError = 7
	ReceiptError_ExecutionInternal       ReceiptError = 8
	ReceiptError_TransactionMalformed    ReceiptError = 9
	// EvmError
	ReceiptError_OutOfGas                   ReceiptError = 10
	ReceiptError_BadJumpDestination         ReceiptError = 11
	ReceiptError_BadInstruction             ReceiptError = 12
	ReceiptError_StackUnderflow             ReceiptError = 13
	ReceiptError_OutOfStack                 ReceiptError = 14
	ReceiptError_Internal                   ReceiptError = 15
	ReceiptError_MutableCallInStaticContext ReceiptError = 16
	ReceiptError_OutOfBounds                ReceiptError = 17
	ReceiptError_Reverted                   ReceiptError = 18
)

var ReceiptError_name = map[int32]string{
	0:  "NotEnoughBaseGas",
	1:  "BlockGasLimitReached",
	2:  "AccountGasLimitReached",
	3:  "InvalidTransactionNonce",
	4:  "NotEnoughCash",
	5:  "NoTransactionPermission",
	6:  "NoContractPermission",
	7:  "NoCallPermission",
	8:  "ExecutionInternal",
	9:  "TransactionMalformed",
	10: "OutOfGas",
	11: "BadJumpDestination",
	12: "BadInstruction",
	13: "StackUnderflow",
	14: "OutOfStack",
	15: "Internal",
	16: "MutableCallInStaticContext",
	17: "OutOfBounds",
	18: "Reverted",
}
var ReceiptError_value = map[string]int32{
	"NotEnoughBaseGas":           0,
	"BlockGasLimitReached":       1,
	"AccountGasLimitReached":     2,
	"InvalidTransactionNonce":    3,
	"NotEnoughCash":              4,
	"NoTransactionPermission":    5,
	"NoContractPermission":       6,
	"NoCallPermission":           7,
	"ExecutionInternal":          8,
	"TransactionMalformed":       9,
	"OutOfGas":                   10,
	"BadJumpDestination":         11,
	"BadInstruction":             12,
	"StackUnderflow":             13,
	"OutOfStack":                 14,
	"Internal":                   15,
	"MutableCallInStaticContext": 16,
	"OutOfBounds":                17,
	"Reverted":                   18,
}

func (x ReceiptError) String() string {
	return proto.EnumName(ReceiptError_name, int32(x))
}
func (ReceiptError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{0}
}

type ExecutedHeader struct {
	Prevhash             []byte   `protobuf:"bytes,1,opt,name=prevhash,proto3" json:"prevhash,omitempty"`
	Timestamp            uint64   `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Height               uint64   `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	StateRoot            []byte   `protobuf:"bytes,4,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	TransactionsRoot     []byte   `protobuf:"bytes,5,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	ReceiptsRoot         []byte   `protobuf:"bytes,6,opt,name=receipts_root,json=receiptsRoot,proto3" json:"receipts_root,omitempty"`
	LogBloom             []byte   `protobuf:"bytes,7,opt,name=log_bloom,json=logBloom,proto3" json:"log_bloom,omitempty"`
	GasUsed              uint64   `protobuf:"varint,8,opt,name=gas_used,json=gasUsed" json:"gas_used,omitempty"`
	GasLimit             uint64   `protobuf:"varint,9,opt,name=gas_limit,json=gasLimit" json:"gas_limit,omitempty"`
	Proposer             []byte   `protobuf:"bytes,10,opt,name=proposer,proto3" json:"proposer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutedHeader) Reset()         { *m = ExecutedHeader{} }
func (m *ExecutedHeader) String() string { return proto.CompactTextString(m) }
func (*ExecutedHeader) ProtoMessage()    {}
func (*ExecutedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{0}
}
func (m *ExecutedHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutedHeader.Unmarshal(m, b)
}
func (m *ExecutedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutedHeader.Marshal(b, m, deterministic)
}
func (dst *ExecutedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutedHeader.Merge(dst, src)
}
func (m *ExecutedHeader) XXX_Size() int {
	return xxx_messageInfo_ExecutedHeader.Size(m)
}
func (m *ExecutedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutedHeader proto.InternalMessageInfo

func (m *ExecutedHeader) GetPrevhash() []byte {
	if m != nil {
		return m.Prevhash
	}
	return nil
}

func (m *ExecutedHeader) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ExecutedHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ExecutedHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *ExecutedHeader) GetTransactionsRoot() []byte {
	if m != nil {
		return m.TransactionsRoot
	}
	return nil
}

func (m *ExecutedHeader) GetReceiptsRoot() []byte {
	if m != nil {
		return m.ReceiptsRoot
	}
	return nil
}

func (m *ExecutedHeader) GetLogBloom() []byte {
	if m != nil {
		return m.LogBloom
	}
	return nil
}

func (m *ExecutedHeader) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *ExecutedHeader) GetGasLimit() uint64 {
	if m != nil {
		return m.GasLimit
	}
	return 0
}

func (m *ExecutedHeader) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type LogEntry struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Topics               [][]byte `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{1}
}
func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (dst *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(dst, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *LogEntry) GetTopics() [][]byte {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *LogEntry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ReceiptErrorWithOption struct {
	Error                ReceiptError `protobuf:"varint,1,opt,name=error,enum=ReceiptError" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReceiptErrorWithOption) Reset()         { *m = ReceiptErrorWithOption{} }
func (m *ReceiptErrorWithOption) String() string { return proto.CompactTextString(m) }
func (*ReceiptErrorWithOption) ProtoMessage()    {}
func (*ReceiptErrorWithOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{2}
}
func (m *ReceiptErrorWithOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptErrorWithOption.Unmarshal(m, b)
}
func (m *ReceiptErrorWithOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptErrorWithOption.Marshal(b, m, deterministic)
}
func (dst *ReceiptErrorWithOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptErrorWithOption.Merge(dst, src)
}
func (m *ReceiptErrorWithOption) XXX_Size() int {
	return xxx_messageInfo_ReceiptErrorWithOption.Size(m)
}
func (m *ReceiptErrorWithOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptErrorWithOption.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptErrorWithOption proto.InternalMessageInfo

func (m *ReceiptErrorWithOption) GetError() ReceiptError {
	if m != nil {
		return m.Error
	}
	return ReceiptError_NotEnoughBaseGas
}

type StateRoot struct {
	StateRoot            []byte   `protobuf:"bytes,1,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StateRoot) Reset()         { *m = StateRoot{} }
func (m *StateRoot) String() string { return proto.CompactTextString(m) }
func (*StateRoot) ProtoMessage()    {}
func (*StateRoot) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{3}
}
func (m *StateRoot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateRoot.Unmarshal(m, b)
}
func (m *StateRoot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateRoot.Marshal(b, m, deterministic)
}
func (dst *StateRoot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateRoot.Merge(dst, src)
}
func (m *StateRoot) XXX_Size() int {
	return xxx_messageInfo_StateRoot.Size(m)
}
func (m *StateRoot) XXX_DiscardUnknown() {
	xxx_messageInfo_StateRoot.DiscardUnknown(m)
}

var xxx_messageInfo_StateRoot proto.InternalMessageInfo

func (m *StateRoot) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

type Receipt struct {
	StateRoot            *StateRoot              `protobuf:"bytes,1,opt,name=state_root,json=stateRoot" json:"state_root,omitempty"`
	GasUsed              string                  `protobuf:"bytes,2,opt,name=gas_used,json=gasUsed" json:"gas_used,omitempty"`
	LogBloom             []byte                  `protobuf:"bytes,3,opt,name=log_bloom,json=logBloom,proto3" json:"log_bloom,omitempty"`
	Logs                 []*LogEntry             `protobuf:"bytes,4,rep,name=logs" json:"logs,omitempty"`
	Error                *ReceiptErrorWithOption `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	AccountNonce         uint64                  `protobuf:"varint,6,opt,name=account_nonce,json=accountNonce" json:"account_nonce,omitempty"`
	TransactionHash      []byte                  `protobuf:"bytes,7,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Receipt) Reset()         { *m = Receipt{} }
func (m *Receipt) String() string { return proto.CompactTextString(m) }
func (*Receipt) ProtoMessage()    {}
func (*Receipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{4}
}
func (m *Receipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Receipt.Unmarshal(m, b)
}
func (m *Receipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Receipt.Marshal(b, m, deterministic)
}
func (dst *Receipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Receipt.Merge(dst, src)
}
func (m *Receipt) XXX_Size() int {
	return xxx_messageInfo_Receipt.Size(m)
}
func (m *Receipt) XXX_DiscardUnknown() {
	xxx_messageInfo_Receipt.DiscardUnknown(m)
}

var xxx_messageInfo_Receipt proto.InternalMessageInfo

func (m *Receipt) GetStateRoot() *StateRoot {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *Receipt) GetGasUsed() string {
	if m != nil {
		return m.GasUsed
	}
	return ""
}

func (m *Receipt) GetLogBloom() []byte {
	if m != nil {
		return m.LogBloom
	}
	return nil
}

func (m *Receipt) GetLogs() []*LogEntry {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Receipt) GetError() *ReceiptErrorWithOption {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Receipt) GetAccountNonce() uint64 {
	if m != nil {
		return m.AccountNonce
	}
	return 0
}

func (m *Receipt) GetTransactionHash() []byte {
	if m != nil {
		return m.TransactionHash
	}
	return nil
}

type ReceiptWithOption struct {
	Receipt              *Receipt `protobuf:"bytes,1,opt,name=receipt" json:"receipt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiptWithOption) Reset()         { *m = ReceiptWithOption{} }
func (m *ReceiptWithOption) String() string { return proto.CompactTextString(m) }
func (*ReceiptWithOption) ProtoMessage()    {}
func (*ReceiptWithOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{5}
}
func (m *ReceiptWithOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptWithOption.Unmarshal(m, b)
}
func (m *ReceiptWithOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptWithOption.Marshal(b, m, deterministic)
}
func (dst *ReceiptWithOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptWithOption.Merge(dst, src)
}
func (m *ReceiptWithOption) XXX_Size() int {
	return xxx_messageInfo_ReceiptWithOption.Size(m)
}
func (m *ReceiptWithOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptWithOption.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptWithOption proto.InternalMessageInfo

func (m *ReceiptWithOption) GetReceipt() *Receipt {
	if m != nil {
		return m.Receipt
	}
	return nil
}

type ExecutedInfo struct {
	Header               *ExecutedHeader      `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Receipts             []*ReceiptWithOption `protobuf:"bytes,3,rep,name=receipts" json:"receipts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ExecutedInfo) Reset()         { *m = ExecutedInfo{} }
func (m *ExecutedInfo) String() string { return proto.CompactTextString(m) }
func (*ExecutedInfo) ProtoMessage()    {}
func (*ExecutedInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{6}
}
func (m *ExecutedInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutedInfo.Unmarshal(m, b)
}
func (m *ExecutedInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutedInfo.Marshal(b, m, deterministic)
}
func (dst *ExecutedInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutedInfo.Merge(dst, src)
}
func (m *ExecutedInfo) XXX_Size() int {
	return xxx_messageInfo_ExecutedInfo.Size(m)
}
func (m *ExecutedInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutedInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutedInfo proto.InternalMessageInfo

func (m *ExecutedInfo) GetHeader() *ExecutedHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ExecutedInfo) GetReceipts() []*ReceiptWithOption {
	if m != nil {
		return m.Receipts
	}
	return nil
}

type ConsensusConfig struct {
	BlockGasLimit        uint64           `protobuf:"varint,1,opt,name=block_gas_limit,json=blockGasLimit" json:"block_gas_limit,omitempty"`
	AccountGasLimit      *AccountGasLimit `protobuf:"bytes,2,opt,name=account_gas_limit,json=accountGasLimit" json:"account_gas_limit,omitempty"`
	Nodes                [][]byte         `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	BlockInterval        uint64           `protobuf:"varint,4,opt,name=block_interval,json=blockInterval" json:"block_interval,omitempty"`
	CheckQuota           bool             `protobuf:"varint,5,opt,name=check_quota,json=checkQuota" json:"check_quota,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ConsensusConfig) Reset()         { *m = ConsensusConfig{} }
func (m *ConsensusConfig) String() string { return proto.CompactTextString(m) }
func (*ConsensusConfig) ProtoMessage()    {}
func (*ConsensusConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{7}
}
func (m *ConsensusConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusConfig.Unmarshal(m, b)
}
func (m *ConsensusConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusConfig.Marshal(b, m, deterministic)
}
func (dst *ConsensusConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusConfig.Merge(dst, src)
}
func (m *ConsensusConfig) XXX_Size() int {
	return xxx_messageInfo_ConsensusConfig.Size(m)
}
func (m *ConsensusConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusConfig proto.InternalMessageInfo

func (m *ConsensusConfig) GetBlockGasLimit() uint64 {
	if m != nil {
		return m.BlockGasLimit
	}
	return 0
}

func (m *ConsensusConfig) GetAccountGasLimit() *AccountGasLimit {
	if m != nil {
		return m.AccountGasLimit
	}
	return nil
}

func (m *ConsensusConfig) GetNodes() [][]byte {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *ConsensusConfig) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

func (m *ConsensusConfig) GetCheckQuota() bool {
	if m != nil {
		return m.CheckQuota
	}
	return false
}

type ExecutedResult struct {
	ExecutedInfo         *ExecutedInfo    `protobuf:"bytes,1,opt,name=executed_info,json=executedInfo" json:"executed_info,omitempty"`
	Config               *ConsensusConfig `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ExecutedResult) Reset()         { *m = ExecutedResult{} }
func (m *ExecutedResult) String() string { return proto.CompactTextString(m) }
func (*ExecutedResult) ProtoMessage()    {}
func (*ExecutedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{8}
}
func (m *ExecutedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutedResult.Unmarshal(m, b)
}
func (m *ExecutedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutedResult.Marshal(b, m, deterministic)
}
func (dst *ExecutedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutedResult.Merge(dst, src)
}
func (m *ExecutedResult) XXX_Size() int {
	return xxx_messageInfo_ExecutedResult.Size(m)
}
func (m *ExecutedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutedResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutedResult proto.InternalMessageInfo

func (m *ExecutedResult) GetExecutedInfo() *ExecutedInfo {
	if m != nil {
		return m.ExecutedInfo
	}
	return nil
}

func (m *ExecutedResult) GetConfig() *ConsensusConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type RegisterRequest struct {
	ContractAddress      string   `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress" json:"contract_address,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Port                 string   `protobuf:"bytes,3,opt,name=port" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{9}
}
func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(dst, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *RegisterRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RegisterRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type RegisterResponse struct {
	State                string   `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{10}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type LoadRequest struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	ContractAddress      string   `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress" json:"contract_address,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadRequest) Reset()         { *m = LoadRequest{} }
func (m *LoadRequest) String() string { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()    {}
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{11}
}
func (m *LoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadRequest.Unmarshal(m, b)
}
func (m *LoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadRequest.Marshal(b, m, deterministic)
}
func (dst *LoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadRequest.Merge(dst, src)
}
func (m *LoadRequest) XXX_Size() int {
	return xxx_messageInfo_LoadRequest.Size(m)
}
func (m *LoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadRequest proto.InternalMessageInfo

func (m *LoadRequest) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *LoadRequest) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *LoadRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type LoadResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadResponse) Reset()         { *m = LoadResponse{} }
func (m *LoadResponse) String() string { return proto.CompactTextString(m) }
func (*LoadResponse) ProtoMessage()    {}
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_executor_a2e24703375185b1, []int{12}
}
func (m *LoadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadResponse.Unmarshal(m, b)
}
func (m *LoadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadResponse.Marshal(b, m, deterministic)
}
func (dst *LoadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadResponse.Merge(dst, src)
}
func (m *LoadResponse) XXX_Size() int {
	return xxx_messageInfo_LoadResponse.Size(m)
}
func (m *LoadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadResponse proto.InternalMessageInfo

func (m *LoadResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecutedHeader)(nil), "ExecutedHeader")
	proto.RegisterType((*LogEntry)(nil), "LogEntry")
	proto.RegisterType((*ReceiptErrorWithOption)(nil), "ReceiptErrorWithOption")
	proto.RegisterType((*StateRoot)(nil), "StateRoot")
	proto.RegisterType((*Receipt)(nil), "Receipt")
	proto.RegisterType((*ReceiptWithOption)(nil), "ReceiptWithOption")
	proto.RegisterType((*ExecutedInfo)(nil), "ExecutedInfo")
	proto.RegisterType((*ConsensusConfig)(nil), "ConsensusConfig")
	proto.RegisterType((*ExecutedResult)(nil), "ExecutedResult")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "RegisterResponse")
	proto.RegisterType((*LoadRequest)(nil), "LoadRequest")
	proto.RegisterType((*LoadResponse)(nil), "LoadResponse")
	proto.RegisterEnum("ReceiptError", ReceiptError_name, ReceiptError_value)
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_executor_a2e24703375185b1) }

var fileDescriptor_executor_a2e24703375185b1 = []byte{
	// 1052 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x56, 0xdd, 0x72, 0xdb, 0x36,
	0x13, 0x8d, 0x7e, 0x6c, 0x49, 0x2b, 0x4a, 0x82, 0x30, 0xf9, 0x1c, 0x7e, 0x4a, 0xd3, 0x7a, 0x98,
	0xa6, 0x55, 0xd2, 0x29, 0x67, 0xea, 0x5e, 0xf4, 0xa6, 0xbd, 0x88, 0x5c, 0x4f, 0xe2, 0x8e, 0x63,
	0xa7, 0x70, 0x33, 0xbd, 0x54, 0x61, 0x12, 0x92, 0x30, 0xa6, 0x00, 0x06, 0x00, 0xd5, 0xe4, 0x15,
	0xfa, 0x3a, 0x7d, 0x85, 0xbe, 0x42, 0xdf, 0xa7, 0x03, 0x10, 0x94, 0x68, 0xc5, 0x77, 0xdc, 0xb3,
	0x07, 0x8b, 0xdd, 0xa3, 0xb3, 0xb0, 0x61, 0xc8, 0x3e, 0xb0, 0xa4, 0x30, 0x52, 0xc5, 0xb9, 0x92,
	0x46, 0x4e, 0xd0, 0x4d, 0x26, 0x93, 0xdb, 0x64, 0x45, 0xb9, 0x28, 0x91, 0xe8, 0x9f, 0x26, 0x0c,
	0xcf, 0x1c, 0x89, 0xa5, 0xaf, 0x19, 0x4d, 0x99, 0xc2, 0x13, 0xe8, 0xe6, 0x8a, 0x6d, 0x56, 0x54,
	0xaf, 0xc2, 0xc6, 0x71, 0x63, 0x1a, 0x90, 0x6d, 0x8c, 0x3f, 0x83, 0x9e, 0xe1, 0x6b, 0xa6, 0x0d,
	0x5d, 0xe7, 0x61, 0xf3, 0xb8, 0x31, 0x6d, 0x93, 0x1d, 0x80, 0x8f, 0xe0, 0x70, 0xc5, 0xf8, 0x72,
	0x65, 0xc2, 0x96, 0x4b, 0xf9, 0x08, 0x3f, 0x01, 0xd0, 0x86, 0x1a, 0x36, 0x57, 0x52, 0x9a, 0xb0,
	0xed, 0x6a, 0xf6, 0x1c, 0x42, 0xa4, 0x34, 0xf8, 0x1b, 0x18, 0x1b, 0x45, 0x85, 0xa6, 0x89, 0xe1,
	0x52, 0xe8, 0x92, 0x75, 0xe0, 0x58, 0xa8, 0x9e, 0x70, 0xe4, 0xa7, 0x30, 0x50, 0x2c, 0x61, 0x3c,
	0x37, 0x9e, 0x78, 0xe8, 0x88, 0x41, 0x05, 0x3a, 0xd2, 0x63, 0xe8, 0x65, 0x72, 0x39, 0xbf, 0xc9,
	0xa4, 0x5c, 0x87, 0x9d, 0x72, 0x86, 0x4c, 0x2e, 0x67, 0x36, 0xc6, 0xff, 0x87, 0xee, 0x92, 0xea,
	0x79, 0xa1, 0x59, 0x1a, 0x76, 0x5d, 0x9f, 0x9d, 0x25, 0xd5, 0xef, 0x34, 0x4b, 0xed, 0x39, 0x9b,
	0xca, 0xf8, 0x9a, 0x9b, 0xb0, 0xe7, 0x72, 0x96, 0x7b, 0x61, 0xe3, 0x52, 0x17, 0x99, 0x4b, 0xcd,
	0x54, 0x08, 0x95, 0x2e, 0x65, 0x1c, 0xbd, 0x85, 0xee, 0x85, 0x5c, 0x9e, 0x09, 0xa3, 0x3e, 0xe2,
	0x10, 0x3a, 0x34, 0x4d, 0x15, 0xd3, 0xda, 0xcb, 0x57, 0x85, 0x56, 0x1f, 0x23, 0x73, 0x9e, 0xe8,
	0xb0, 0x79, 0xdc, 0x9a, 0x06, 0xc4, 0x47, 0x18, 0x43, 0x3b, 0xa5, 0x86, 0x3a, 0xd5, 0x02, 0xe2,
	0xbe, 0xa3, 0x9f, 0xe0, 0x88, 0x94, 0x23, 0x9d, 0x29, 0x25, 0xd5, 0xef, 0xdc, 0xac, 0xae, 0x72,
	0x2b, 0x03, 0x7e, 0x0a, 0x07, 0xcc, 0x42, 0xae, 0xfa, 0xf0, 0x64, 0x10, 0xd7, 0x79, 0xa4, 0xcc,
	0x45, 0x2f, 0xa0, 0x77, 0xbd, 0x15, 0xf8, 0xae, 0xfe, 0x8d, 0x3d, 0xfd, 0xa3, 0xbf, 0x9a, 0xd0,
	0xf1, 0x35, 0xf0, 0xf3, 0x4f, 0xa8, 0xfd, 0x13, 0x88, 0xb7, 0xa5, 0xea, 0x3f, 0x5b, 0x5d, 0x47,
	0x6b, 0x85, 0xde, 0x1d, 0x1d, 0x77, 0xfa, 0xb7, 0xf6, 0xf4, 0x7f, 0x02, 0xed, 0x4c, 0x2e, 0x75,
	0xd8, 0x3e, 0x6e, 0x4d, 0xfb, 0x27, 0xbd, 0xb8, 0x12, 0x8e, 0x38, 0x18, 0x7f, 0x5b, 0x8d, 0x77,
	0xe0, 0x2e, 0x7f, 0x14, 0xdf, 0x2f, 0x83, 0x1f, 0xd4, 0xfa, 0x81, 0x26, 0x89, 0x2c, 0x84, 0x99,
	0x0b, 0x29, 0x12, 0xe6, 0xfc, 0xd0, 0x26, 0x81, 0x07, 0x2f, 0x2d, 0x86, 0x9f, 0x43, 0xdd, 0x48,
	0x73, 0x67, 0xed, 0xd2, 0x16, 0xa3, 0x1a, 0xfe, 0x9a, 0xea, 0x55, 0xf4, 0x03, 0x8c, 0xfd, 0x85,
	0x35, 0xc9, 0x23, 0xe8, 0x78, 0x7f, 0x79, 0x49, 0xba, 0x55, 0x57, 0xa4, 0x4a, 0x44, 0x4b, 0x08,
	0xaa, 0x45, 0x3a, 0x17, 0x0b, 0x89, 0xbf, 0xb6, 0xcb, 0x60, 0x17, 0xca, 0x1f, 0x19, 0xc5, 0x77,
	0xf7, 0x8c, 0xf8, 0x34, 0x8e, 0xa1, 0x5b, 0x99, 0x37, 0x6c, 0x39, 0x4d, 0x70, 0xfc, 0x49, 0x0b,
	0x64, 0xcb, 0x89, 0xfe, 0x6d, 0xc0, 0xe8, 0x54, 0x0a, 0xcd, 0x84, 0x2e, 0xf4, 0xa9, 0x14, 0x0b,
	0xbe, 0xc4, 0x5f, 0xc1, 0xc8, 0xad, 0xf6, 0x7c, 0x67, 0xdf, 0x86, 0xd3, 0x61, 0xe0, 0xe0, 0x57,
	0x95, 0x87, 0x7f, 0x84, 0x71, 0xa5, 0xd6, 0x8e, 0xd9, 0x74, 0xfd, 0xa1, 0xf8, 0x65, 0x99, 0xa9,
	0xc8, 0x64, 0x44, 0xef, 0x02, 0xf8, 0x21, 0x1c, 0x08, 0x99, 0xb2, 0xb2, 0xcd, 0x80, 0x94, 0x01,
	0x7e, 0x06, 0xc3, 0xf2, 0x6e, 0x2e, 0x0c, 0x53, 0x1b, 0x9a, 0xb9, 0x0d, 0xaf, 0xae, 0x3e, 0xf7,
	0x20, 0xfe, 0x02, 0xfa, 0xc9, 0x8a, 0x25, 0xb7, 0xf3, 0xf7, 0x85, 0x34, 0xd4, 0xfd, 0xba, 0x5d,
	0x02, 0x0e, 0xfa, 0xd5, 0x22, 0x91, 0xd8, 0xbd, 0x44, 0x84, 0xe9, 0x22, 0x33, 0xf8, 0x04, 0x06,
	0xcc, 0x23, 0x73, 0x2e, 0x16, 0xd2, 0x2b, 0x39, 0x88, 0xeb, 0x42, 0x93, 0x80, 0xd5, 0x65, 0x9f,
	0xc2, 0x61, 0xe2, 0x34, 0xd9, 0x8e, 0xb5, 0xa7, 0x15, 0xf1, 0xf9, 0xe8, 0x0f, 0x18, 0x11, 0xb6,
	0xe4, 0xda, 0x30, 0x45, 0xd8, 0xfb, 0x82, 0x69, 0xeb, 0x7e, 0x94, 0x48, 0x61, 0x14, 0x4d, 0xcc,
	0xbc, 0xbe, 0xc3, 0x3d, 0x32, 0xaa, 0xf0, 0x97, 0x7e, 0x97, 0x87, 0xd0, 0xe4, 0xb9, 0xf7, 0x7d,
	0x93, 0xe7, 0x76, 0x87, 0x73, 0xa9, 0xca, 0x97, 0xaf, 0x47, 0xdc, 0x77, 0x34, 0x05, 0xb4, 0xbb,
	0x41, 0xe7, 0xb6, 0x0f, 0xab, 0xa1, 0x5b, 0x21, 0x5f, 0xb7, 0x0c, 0xa2, 0x1b, 0xe8, 0x5f, 0x48,
	0x9a, 0x56, 0x7d, 0xec, 0x1e, 0xd2, 0xc6, 0x9d, 0x87, 0xf4, 0xbe, 0xfe, 0x9a, 0xf7, 0xf7, 0x87,
	0xa0, 0x75, 0xcb, 0x3e, 0xfa, 0x76, 0xec, 0x67, 0xf4, 0x25, 0x04, 0xe5, 0x1d, 0xbb, 0x4e, 0x36,
	0x34, 0x2b, 0xb6, 0x9d, 0xb8, 0xe0, 0xc5, 0xdf, 0x2d, 0x08, 0xea, 0x1b, 0x87, 0x1f, 0x02, 0xba,
	0x94, 0xe6, 0x4c, 0xc8, 0x62, 0xb9, 0x9a, 0x51, 0xcd, 0x5e, 0x51, 0x8d, 0x1e, 0xe0, 0x10, 0x1e,
	0xce, 0xea, 0xce, 0x22, 0x8c, 0x26, 0x2b, 0x96, 0xa2, 0x06, 0x9e, 0xc0, 0xd1, 0xbe, 0x91, 0x7c,
	0xae, 0x89, 0x1f, 0xc3, 0xa3, 0x73, 0xb1, 0xa1, 0x19, 0x4f, 0x7f, 0xdb, 0xad, 0x9d, 0x5b, 0x51,
	0xd4, 0xc2, 0x63, 0x18, 0x6c, 0x2f, 0x3a, 0xa5, 0x7a, 0x85, 0xda, 0x96, 0x7f, 0x29, 0x6b, 0xd4,
	0xb7, 0x4c, 0xad, 0xb9, 0xd6, 0x5c, 0x0a, 0x74, 0x60, 0x5b, 0xb8, 0x94, 0xa7, 0x7e, 0xec, 0x5a,
	0xe6, 0xb0, 0x6c, 0xf9, 0x94, 0x66, 0x59, 0x0d, 0xed, 0xe0, 0xff, 0xc1, 0xb8, 0xf4, 0x0d, 0x97,
	0xc2, 0xb9, 0x52, 0xd0, 0x0c, 0x75, 0x6d, 0x99, 0xda, 0x0d, 0x6f, 0x68, 0xb6, 0x90, 0x6a, 0xcd,
	0x52, 0xd4, 0xc3, 0x01, 0x74, 0xaf, 0x0a, 0x73, 0xb5, 0xb0, 0x13, 0x03, 0x3e, 0x02, 0x3c, 0xa3,
	0xe9, 0x2f, 0xc5, 0x3a, 0xff, 0x99, 0x69, 0xc3, 0x05, 0xb5, 0x74, 0xd4, 0xc7, 0x18, 0x86, 0x33,
	0x9a, 0x9e, 0x0b, 0x6d, 0x54, 0xe1, 0x4a, 0xa0, 0xc0, 0x62, 0xd7, 0x86, 0x26, 0xb7, 0xef, 0x44,
	0xca, 0xd4, 0x22, 0x93, 0x7f, 0xa2, 0x01, 0x1e, 0x02, 0xb8, 0x6a, 0x2e, 0x81, 0x86, 0xb6, 0xfa,
	0xb6, 0x8b, 0x11, 0xfe, 0x1c, 0x26, 0x6f, 0x0a, 0x43, 0x6f, 0x32, 0x66, 0xfb, 0x3e, 0x17, 0xf6,
	0xc5, 0xe5, 0x89, 0x9d, 0x8e, 0x7d, 0x30, 0x08, 0xe1, 0x11, 0xf4, 0xdd, 0xe9, 0x99, 0x2c, 0x44,
	0xaa, 0xd1, 0xd8, 0x1e, 0x27, 0x6c, 0xc3, 0x94, 0x61, 0x29, 0xc2, 0x27, 0xb7, 0x30, 0x3a, 0xf3,
	0x7f, 0xea, 0xaf, 0x99, 0xda, 0xf0, 0x84, 0xe1, 0xef, 0x2c, 0xa1, 0x34, 0x1f, 0x46, 0xf1, 0x9e,
	0xd3, 0x27, 0xe3, 0x78, 0xdf, 0x99, 0xd1, 0x03, 0xfc, 0x0c, 0xda, 0xd6, 0x21, 0x38, 0x88, 0x6b,
	0x66, 0x9c, 0x0c, 0xe2, 0xba, 0x6d, 0xa2, 0x07, 0x37, 0x87, 0xee, 0x5f, 0x87, 0xef, 0xff, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x2f, 0xf4, 0xc9, 0xaf, 0x5e, 0x08, 0x00, 0x00,
}
