// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

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

type Ret int32

const (
	Ret_OK                Ret = 0
	Ret_InvalidNonce      Ret = 1
	Ret_Dup               Ret = 2
	Ret_InvalidUntilBlock Ret = 3
	Ret_BadSig            Ret = 4
	Ret_NotReady          Ret = 5
	Ret_Busy              Ret = 6
	Ret_BadChainId        Ret = 7
	Ret_QuotaNotEnough    Ret = 8
)

var Ret_name = map[int32]string{
	0: "OK",
	1: "InvalidNonce",
	2: "Dup",
	3: "InvalidUntilBlock",
	4: "BadSig",
	5: "NotReady",
	6: "Busy",
	7: "BadChainId",
	8: "QuotaNotEnough",
}
var Ret_value = map[string]int32{
	"OK":                0,
	"InvalidNonce":      1,
	"Dup":               2,
	"InvalidUntilBlock": 3,
	"BadSig":            4,
	"NotReady":          5,
	"Busy":              6,
	"BadChainId":        7,
	"QuotaNotEnough":    8,
}

func (x Ret) String() string {
	return proto.EnumName(Ret_name, int32(x))
}
func (Ret) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{0}
}

type VerifyTxReq struct {
	ValidUntilBlock      uint64   `protobuf:"varint,1,opt,name=valid_until_block,json=validUntilBlock" json:"valid_until_block,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Signature            []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Crypto               Crypto   `protobuf:"varint,4,opt,name=crypto,enum=Crypto" json:"crypto,omitempty"`
	TxHash               []byte   `protobuf:"bytes,5,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Signer               []byte   `protobuf:"bytes,6,opt,name=signer,proto3" json:"signer,omitempty"`
	Nonce                string   `protobuf:"bytes,7,opt,name=nonce" json:"nonce,omitempty"`
	ChainId              uint32   `protobuf:"varint,8,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`
	Quota                uint64   `protobuf:"varint,9,opt,name=quota" json:"quota,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTxReq) Reset()         { *m = VerifyTxReq{} }
func (m *VerifyTxReq) String() string { return proto.CompactTextString(m) }
func (*VerifyTxReq) ProtoMessage()    {}
func (*VerifyTxReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{0}
}
func (m *VerifyTxReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTxReq.Unmarshal(m, b)
}
func (m *VerifyTxReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTxReq.Marshal(b, m, deterministic)
}
func (dst *VerifyTxReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTxReq.Merge(dst, src)
}
func (m *VerifyTxReq) XXX_Size() int {
	return xxx_messageInfo_VerifyTxReq.Size(m)
}
func (m *VerifyTxReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTxReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTxReq proto.InternalMessageInfo

func (m *VerifyTxReq) GetValidUntilBlock() uint64 {
	if m != nil {
		return m.ValidUntilBlock
	}
	return 0
}

func (m *VerifyTxReq) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *VerifyTxReq) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *VerifyTxReq) GetCrypto() Crypto {
	if m != nil {
		return m.Crypto
	}
	return Crypto_SECP
}

func (m *VerifyTxReq) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *VerifyTxReq) GetSigner() []byte {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *VerifyTxReq) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *VerifyTxReq) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *VerifyTxReq) GetQuota() uint64 {
	if m != nil {
		return m.Quota
	}
	return 0
}

type VerifyTxResp struct {
	TxHash               []byte   `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	Ret                  Ret      `protobuf:"varint,2,opt,name=ret,enum=Ret" json:"ret,omitempty"`
	Signer               []byte   `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty"`
	ReceivedChainId      uint32   `protobuf:"varint,4,opt,name=received_chain_id,json=receivedChainId" json:"received_chain_id,omitempty"`
	ExpectedChainId      uint32   `protobuf:"varint,5,opt,name=expected_chain_id,json=expectedChainId" json:"expected_chain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyTxResp) Reset()         { *m = VerifyTxResp{} }
func (m *VerifyTxResp) String() string { return proto.CompactTextString(m) }
func (*VerifyTxResp) ProtoMessage()    {}
func (*VerifyTxResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{1}
}
func (m *VerifyTxResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyTxResp.Unmarshal(m, b)
}
func (m *VerifyTxResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyTxResp.Marshal(b, m, deterministic)
}
func (dst *VerifyTxResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyTxResp.Merge(dst, src)
}
func (m *VerifyTxResp) XXX_Size() int {
	return xxx_messageInfo_VerifyTxResp.Size(m)
}
func (m *VerifyTxResp) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyTxResp.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyTxResp proto.InternalMessageInfo

func (m *VerifyTxResp) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *VerifyTxResp) GetRet() Ret {
	if m != nil {
		return m.Ret
	}
	return Ret_OK
}

func (m *VerifyTxResp) GetSigner() []byte {
	if m != nil {
		return m.Signer
	}
	return nil
}

func (m *VerifyTxResp) GetReceivedChainId() uint32 {
	if m != nil {
		return m.ReceivedChainId
	}
	return 0
}

func (m *VerifyTxResp) GetExpectedChainId() uint32 {
	if m != nil {
		return m.ExpectedChainId
	}
	return 0
}

type VerifyBlockReq struct {
	Id                   uint64         `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Reqs                 []*VerifyTxReq `protobuf:"bytes,2,rep,name=reqs" json:"reqs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VerifyBlockReq) Reset()         { *m = VerifyBlockReq{} }
func (m *VerifyBlockReq) String() string { return proto.CompactTextString(m) }
func (*VerifyBlockReq) ProtoMessage()    {}
func (*VerifyBlockReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{2}
}
func (m *VerifyBlockReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyBlockReq.Unmarshal(m, b)
}
func (m *VerifyBlockReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyBlockReq.Marshal(b, m, deterministic)
}
func (dst *VerifyBlockReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyBlockReq.Merge(dst, src)
}
func (m *VerifyBlockReq) XXX_Size() int {
	return xxx_messageInfo_VerifyBlockReq.Size(m)
}
func (m *VerifyBlockReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyBlockReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyBlockReq proto.InternalMessageInfo

func (m *VerifyBlockReq) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VerifyBlockReq) GetReqs() []*VerifyTxReq {
	if m != nil {
		return m.Reqs
	}
	return nil
}

type VerifyBlockResp struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Ret                  Ret      `protobuf:"varint,2,opt,name=ret,enum=Ret" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyBlockResp) Reset()         { *m = VerifyBlockResp{} }
func (m *VerifyBlockResp) String() string { return proto.CompactTextString(m) }
func (*VerifyBlockResp) ProtoMessage()    {}
func (*VerifyBlockResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{3}
}
func (m *VerifyBlockResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyBlockResp.Unmarshal(m, b)
}
func (m *VerifyBlockResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyBlockResp.Marshal(b, m, deterministic)
}
func (dst *VerifyBlockResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyBlockResp.Merge(dst, src)
}
func (m *VerifyBlockResp) XXX_Size() int {
	return xxx_messageInfo_VerifyBlockResp.Size(m)
}
func (m *VerifyBlockResp) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyBlockResp.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyBlockResp proto.InternalMessageInfo

func (m *VerifyBlockResp) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VerifyBlockResp) GetRet() Ret {
	if m != nil {
		return m.Ret
	}
	return Ret_OK
}

type BlockTxHashes struct {
	Height               uint64           `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	TxHashes             [][]byte         `protobuf:"bytes,2,rep,name=tx_hashes,json=txHashes,proto3" json:"tx_hashes,omitempty"`
	BlockGasLimit        uint64           `protobuf:"varint,3,opt,name=block_gas_limit,json=blockGasLimit" json:"block_gas_limit,omitempty"`
	AccountGasLimit      *AccountGasLimit `protobuf:"bytes,4,opt,name=account_gas_limit,json=accountGasLimit" json:"account_gas_limit,omitempty"`
	CheckQuota           bool             `protobuf:"varint,5,opt,name=check_quota,json=checkQuota" json:"check_quota,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BlockTxHashes) Reset()         { *m = BlockTxHashes{} }
func (m *BlockTxHashes) String() string { return proto.CompactTextString(m) }
func (*BlockTxHashes) ProtoMessage()    {}
func (*BlockTxHashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{4}
}
func (m *BlockTxHashes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTxHashes.Unmarshal(m, b)
}
func (m *BlockTxHashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTxHashes.Marshal(b, m, deterministic)
}
func (dst *BlockTxHashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTxHashes.Merge(dst, src)
}
func (m *BlockTxHashes) XXX_Size() int {
	return xxx_messageInfo_BlockTxHashes.Size(m)
}
func (m *BlockTxHashes) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTxHashes.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTxHashes proto.InternalMessageInfo

func (m *BlockTxHashes) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockTxHashes) GetTxHashes() [][]byte {
	if m != nil {
		return m.TxHashes
	}
	return nil
}

func (m *BlockTxHashes) GetBlockGasLimit() uint64 {
	if m != nil {
		return m.BlockGasLimit
	}
	return 0
}

func (m *BlockTxHashes) GetAccountGasLimit() *AccountGasLimit {
	if m != nil {
		return m.AccountGasLimit
	}
	return nil
}

func (m *BlockTxHashes) GetCheckQuota() bool {
	if m != nil {
		return m.CheckQuota
	}
	return false
}

type BlockTxHashesReq struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockTxHashesReq) Reset()         { *m = BlockTxHashesReq{} }
func (m *BlockTxHashesReq) String() string { return proto.CompactTextString(m) }
func (*BlockTxHashesReq) ProtoMessage()    {}
func (*BlockTxHashesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{5}
}
func (m *BlockTxHashesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTxHashesReq.Unmarshal(m, b)
}
func (m *BlockTxHashesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTxHashesReq.Marshal(b, m, deterministic)
}
func (dst *BlockTxHashesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTxHashesReq.Merge(dst, src)
}
func (m *BlockTxHashesReq) XXX_Size() int {
	return xxx_messageInfo_BlockTxHashesReq.Size(m)
}
func (m *BlockTxHashesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTxHashesReq.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTxHashesReq proto.InternalMessageInfo

func (m *BlockTxHashesReq) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Miscellaneous struct {
	ChainId              uint32   `protobuf:"varint,1,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Miscellaneous) Reset()         { *m = Miscellaneous{} }
func (m *Miscellaneous) String() string { return proto.CompactTextString(m) }
func (*Miscellaneous) ProtoMessage()    {}
func (*Miscellaneous) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{6}
}
func (m *Miscellaneous) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Miscellaneous.Unmarshal(m, b)
}
func (m *Miscellaneous) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Miscellaneous.Marshal(b, m, deterministic)
}
func (dst *Miscellaneous) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Miscellaneous.Merge(dst, src)
}
func (m *Miscellaneous) XXX_Size() int {
	return xxx_messageInfo_Miscellaneous.Size(m)
}
func (m *Miscellaneous) XXX_DiscardUnknown() {
	xxx_messageInfo_Miscellaneous.DiscardUnknown(m)
}

var xxx_messageInfo_Miscellaneous proto.InternalMessageInfo

func (m *Miscellaneous) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

type MiscellaneousReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MiscellaneousReq) Reset()         { *m = MiscellaneousReq{} }
func (m *MiscellaneousReq) String() string { return proto.CompactTextString(m) }
func (*MiscellaneousReq) ProtoMessage()    {}
func (*MiscellaneousReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_a4735469b19badfe, []int{7}
}
func (m *MiscellaneousReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MiscellaneousReq.Unmarshal(m, b)
}
func (m *MiscellaneousReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MiscellaneousReq.Marshal(b, m, deterministic)
}
func (dst *MiscellaneousReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MiscellaneousReq.Merge(dst, src)
}
func (m *MiscellaneousReq) XXX_Size() int {
	return xxx_messageInfo_MiscellaneousReq.Size(m)
}
func (m *MiscellaneousReq) XXX_DiscardUnknown() {
	xxx_messageInfo_MiscellaneousReq.DiscardUnknown(m)
}

var xxx_messageInfo_MiscellaneousReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VerifyTxReq)(nil), "VerifyTxReq")
	proto.RegisterType((*VerifyTxResp)(nil), "VerifyTxResp")
	proto.RegisterType((*VerifyBlockReq)(nil), "VerifyBlockReq")
	proto.RegisterType((*VerifyBlockResp)(nil), "VerifyBlockResp")
	proto.RegisterType((*BlockTxHashes)(nil), "BlockTxHashes")
	proto.RegisterType((*BlockTxHashesReq)(nil), "BlockTxHashesReq")
	proto.RegisterType((*Miscellaneous)(nil), "Miscellaneous")
	proto.RegisterType((*MiscellaneousReq)(nil), "MiscellaneousReq")
	proto.RegisterEnum("Ret", Ret_name, Ret_value)
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_a4735469b19badfe) }

var fileDescriptor_auth_a4735469b19badfe = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x65, 0x6c, 0xc7, 0x71, 0x6e, 0x5e, 0x93, 0x11, 0x0f, 0xf3, 0x90, 0x6a, 0x79, 0x81, 0xa2,
	0x2c, 0xb2, 0x08, 0x2b, 0x24, 0x36, 0xa4, 0x20, 0xa8, 0x80, 0x22, 0x86, 0xc2, 0xd6, 0x9a, 0xda,
	0x43, 0x3c, 0x6a, 0xb0, 0x13, 0x7b, 0x5c, 0x25, 0x7b, 0x24, 0x3e, 0x87, 0x0f, 0xe1, 0xa7, 0xd0,
	0x5c, 0x3b, 0x34, 0x06, 0x75, 0xe7, 0x7b, 0xee, 0x43, 0xe7, 0x9c, 0x7b, 0xc7, 0x00, 0xa2, 0xd2,
	0xe9, 0x7c, 0x53, 0xe4, 0x3a, 0x7f, 0x44, 0x2f, 0xd7, 0x79, 0x7c, 0x15, 0xa7, 0x42, 0x65, 0x35,
	0x12, 0xfe, 0xb0, 0xa0, 0xff, 0x55, 0x16, 0xea, 0xdb, 0xfe, 0x62, 0xc7, 0xe5, 0x96, 0xcd, 0x60,
	0x72, 0x2d, 0xd6, 0x2a, 0x89, 0xaa, 0x4c, 0xab, 0x75, 0x84, 0xf5, 0x3e, 0x09, 0xc8, 0xd4, 0xe1,
	0x63, 0x4c, 0x7c, 0x31, 0xf8, 0xd2, 0xc0, 0x8c, 0x81, 0x93, 0x8a, 0x32, 0xf5, 0xad, 0x80, 0x4c,
	0x07, 0x1c, 0xbf, 0xd9, 0x13, 0xe8, 0x95, 0x6a, 0x95, 0x09, 0x5d, 0x15, 0xd2, 0xb7, 0x31, 0x71,
	0x03, 0xb0, 0x13, 0x70, 0xe3, 0x62, 0xbf, 0xd1, 0xb9, 0xef, 0x04, 0x64, 0x3a, 0x5a, 0x74, 0xe7,
	0xa7, 0x18, 0xf2, 0x06, 0x66, 0x0f, 0xa0, 0xab, 0x77, 0x11, 0x4e, 0xed, 0x60, 0xb3, 0xab, 0x77,
	0x6f, 0xcd, 0xdc, 0xfb, 0xe0, 0x9a, 0x31, 0xb2, 0xf0, 0xdd, 0x1a, 0xaf, 0x23, 0x76, 0x17, 0x3a,
	0x59, 0x9e, 0xc5, 0xd2, 0xef, 0x06, 0x64, 0xda, 0xe3, 0x75, 0xc0, 0x1e, 0x82, 0x87, 0x22, 0x23,
	0x95, 0xf8, 0x5e, 0x40, 0xa6, 0x43, 0xde, 0xc5, 0xf8, 0x2c, 0x31, 0x0d, 0xdb, 0x2a, 0xd7, 0xc2,
	0xef, 0xa1, 0xa8, 0x3a, 0x08, 0x7f, 0x11, 0x18, 0xdc, 0xd8, 0x50, 0x6e, 0x8e, 0x89, 0x90, 0x7f,
	0x88, 0xd8, 0x85, 0xd4, 0xa8, 0x79, 0xb4, 0x70, 0xe6, 0x5c, 0x6a, 0x6e, 0x80, 0x23, 0x82, 0x76,
	0x8b, 0xe0, 0x0c, 0x26, 0x85, 0x8c, 0xa5, 0xba, 0x96, 0x49, 0xf4, 0x97, 0x93, 0x83, 0x9c, 0xc6,
	0x87, 0xc4, 0x69, 0xc3, 0x6d, 0x06, 0x13, 0xb9, 0xdb, 0xc8, 0x58, 0x1f, 0xd7, 0x76, 0xea, 0xda,
	0x43, 0xa2, 0xa9, 0x0d, 0x97, 0x30, 0xaa, 0x09, 0xe3, 0x2e, 0xcc, 0xea, 0x46, 0x60, 0xa9, 0xa4,
	0xd9, 0x95, 0xa5, 0x12, 0x16, 0x80, 0x53, 0xc8, 0x6d, 0xe9, 0x5b, 0x81, 0x3d, 0xed, 0x2f, 0x06,
	0xf3, 0xa3, 0x35, 0x73, 0xcc, 0x84, 0xcf, 0x61, 0xdc, 0x9a, 0x51, 0x6e, 0xfe, 0x1b, 0x72, 0x8b,
	0xdc, 0xf0, 0x37, 0x81, 0x21, 0x76, 0x5d, 0xa0, 0x2d, 0xb2, 0x34, 0x06, 0xa4, 0x52, 0xad, 0x52,
	0xdd, 0x74, 0x37, 0x11, 0x7b, 0x0c, 0xbd, 0xc6, 0x49, 0x59, 0x73, 0x19, 0x70, 0x4f, 0x1f, 0x9a,
	0x9e, 0xc2, 0x18, 0x4f, 0x2c, 0x5a, 0x89, 0x32, 0x5a, 0xab, 0xef, 0x4a, 0xa3, 0x7d, 0x0e, 0x1f,
	0x22, 0xfc, 0x46, 0x94, 0xef, 0x0d, 0xc8, 0x5e, 0xc0, 0x44, 0xc4, 0x71, 0x5e, 0x65, 0xfa, 0xa8,
	0xd2, 0xb8, 0xd8, 0x5f, 0xd0, 0xf9, 0xcb, 0x3a, 0x73, 0x28, 0xe6, 0x63, 0xd1, 0x06, 0xd8, 0x09,
	0xf4, 0xe3, 0x54, 0xc6, 0x57, 0x51, 0xbd, 0x79, 0xe3, 0xa8, 0xc7, 0x01, 0xa1, 0x4f, 0xb8, 0xfe,
	0x19, 0xd0, 0x96, 0x18, 0x63, 0xe7, 0x2d, 0x7a, 0xc2, 0x19, 0x0c, 0x3f, 0xa8, 0x32, 0x96, 0xeb,
	0xb5, 0xc8, 0x64, 0x5e, 0x95, 0xad, 0x63, 0x23, 0xad, 0x63, 0x0b, 0x19, 0xd0, 0x56, 0x2d, 0x97,
	0xdb, 0xd9, 0x4f, 0x02, 0x36, 0x97, 0x9a, 0xb9, 0x60, 0x7d, 0x7c, 0x47, 0xef, 0x30, 0x0a, 0x83,
	0xb3, 0x0c, 0x9f, 0xd6, 0xb9, 0xb9, 0x5d, 0x4a, 0x58, 0x17, 0xec, 0x57, 0xd5, 0x86, 0x5a, 0xec,
	0x1e, 0x4c, 0x9a, 0xd4, 0xcd, 0xab, 0xa3, 0x36, 0x03, 0x70, 0x97, 0x22, 0xf9, 0xac, 0x56, 0xd4,
	0x61, 0x03, 0xf0, 0xce, 0x73, 0xcd, 0xa5, 0x48, 0xf6, 0xb4, 0xc3, 0x3c, 0x70, 0x96, 0x55, 0xb9,
	0xa7, 0x2e, 0x1b, 0x01, 0x2c, 0xc5, 0xe1, 0x58, 0x68, 0x97, 0x31, 0x18, 0xa1, 0xd4, 0xf3, 0x5c,
	0xbf, 0xce, 0xf2, 0x6a, 0x95, 0x52, 0xef, 0xd2, 0xc5, 0x5f, 0xc0, 0xb3, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x15, 0x8b, 0xa1, 0xf2, 0x22, 0x04, 0x00, 0x00,
}