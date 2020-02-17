// Code generated by protoc-gen-go. DO NOT EDIT.
// source: valmessage.proto

package valmessage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/offchainlabs/arbitrum/packages/arb-util/common"
	protocol "github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	valprotocol "github.com/offchainlabs/arbitrum/packages/arb-validator/valprotocol"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TokenTypeBuf struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenTypeBuf) Reset()         { *m = TokenTypeBuf{} }
func (m *TokenTypeBuf) String() string { return proto.CompactTextString(m) }
func (*TokenTypeBuf) ProtoMessage()    {}
func (*TokenTypeBuf) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{0}
}

func (m *TokenTypeBuf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenTypeBuf.Unmarshal(m, b)
}
func (m *TokenTypeBuf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenTypeBuf.Marshal(b, m, deterministic)
}
func (m *TokenTypeBuf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenTypeBuf.Merge(m, src)
}
func (m *TokenTypeBuf) XXX_Size() int {
	return xxx_messageInfo_TokenTypeBuf.Size(m)
}
func (m *TokenTypeBuf) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenTypeBuf.DiscardUnknown(m)
}

var xxx_messageInfo_TokenTypeBuf proto.InternalMessageInfo

func (m *TokenTypeBuf) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type VMConfiguration struct {
	GracePeriod           uint64                `protobuf:"varint,1,opt,name=grace_period,json=gracePeriod,proto3" json:"grace_period,omitempty"`
	EscrowRequired        *common.BigIntegerBuf `protobuf:"bytes,2,opt,name=escrow_required,json=escrowRequired,proto3" json:"escrow_required,omitempty"`
	EscrowCurrency        *common.AddressBuf    `protobuf:"bytes,3,opt,name=escrow_currency,json=escrowCurrency,proto3" json:"escrow_currency,omitempty"`
	AssertKeys            []*common.AddressBuf  `protobuf:"bytes,4,rep,name=assert_keys,json=assertKeys,proto3" json:"assert_keys,omitempty"`
	MaxExecutionStepCount uint32                `protobuf:"varint,5,opt,name=max_execution_step_count,json=maxExecutionStepCount,proto3" json:"max_execution_step_count,omitempty"`
	Owner                 *common.AddressBuf    `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}              `json:"-"`
	XXX_unrecognized      []byte                `json:"-"`
	XXX_sizecache         int32                 `json:"-"`
}

func (m *VMConfiguration) Reset()         { *m = VMConfiguration{} }
func (m *VMConfiguration) String() string { return proto.CompactTextString(m) }
func (*VMConfiguration) ProtoMessage()    {}
func (*VMConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{1}
}

func (m *VMConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMConfiguration.Unmarshal(m, b)
}
func (m *VMConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMConfiguration.Marshal(b, m, deterministic)
}
func (m *VMConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMConfiguration.Merge(m, src)
}
func (m *VMConfiguration) XXX_Size() int {
	return xxx_messageInfo_VMConfiguration.Size(m)
}
func (m *VMConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_VMConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_VMConfiguration proto.InternalMessageInfo

func (m *VMConfiguration) GetGracePeriod() uint64 {
	if m != nil {
		return m.GracePeriod
	}
	return 0
}

func (m *VMConfiguration) GetEscrowRequired() *common.BigIntegerBuf {
	if m != nil {
		return m.EscrowRequired
	}
	return nil
}

func (m *VMConfiguration) GetEscrowCurrency() *common.AddressBuf {
	if m != nil {
		return m.EscrowCurrency
	}
	return nil
}

func (m *VMConfiguration) GetAssertKeys() []*common.AddressBuf {
	if m != nil {
		return m.AssertKeys
	}
	return nil
}

func (m *VMConfiguration) GetMaxExecutionStepCount() uint32 {
	if m != nil {
		return m.MaxExecutionStepCount
	}
	return 0
}

func (m *VMConfiguration) GetOwner() *common.AddressBuf {
	if m != nil {
		return m.Owner
	}
	return nil
}

type UnanimousAssertionValidatorNotification struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Signatures           [][]byte `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnanimousAssertionValidatorNotification) Reset() {
	*m = UnanimousAssertionValidatorNotification{}
}
func (m *UnanimousAssertionValidatorNotification) String() string { return proto.CompactTextString(m) }
func (*UnanimousAssertionValidatorNotification) ProtoMessage()    {}
func (*UnanimousAssertionValidatorNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{2}
}

func (m *UnanimousAssertionValidatorNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnanimousAssertionValidatorNotification.Unmarshal(m, b)
}
func (m *UnanimousAssertionValidatorNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnanimousAssertionValidatorNotification.Marshal(b, m, deterministic)
}
func (m *UnanimousAssertionValidatorNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnanimousAssertionValidatorNotification.Merge(m, src)
}
func (m *UnanimousAssertionValidatorNotification) XXX_Size() int {
	return xxx_messageInfo_UnanimousAssertionValidatorNotification.Size(m)
}
func (m *UnanimousAssertionValidatorNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_UnanimousAssertionValidatorNotification.DiscardUnknown(m)
}

var xxx_messageInfo_UnanimousAssertionValidatorNotification proto.InternalMessageInfo

func (m *UnanimousAssertionValidatorNotification) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *UnanimousAssertionValidatorNotification) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type SignedMessage struct {
	Message              *valprotocol.MessageBuf `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature            []byte                  `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SignedMessage) Reset()         { *m = SignedMessage{} }
func (m *SignedMessage) String() string { return proto.CompactTextString(m) }
func (*SignedMessage) ProtoMessage()    {}
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{3}
}

func (m *SignedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedMessage.Unmarshal(m, b)
}
func (m *SignedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedMessage.Marshal(b, m, deterministic)
}
func (m *SignedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedMessage.Merge(m, src)
}
func (m *SignedMessage) XXX_Size() int {
	return xxx_messageInfo_SignedMessage.Size(m)
}
func (m *SignedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SignedMessage proto.InternalMessageInfo

func (m *SignedMessage) GetMessage() *valprotocol.MessageBuf {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *SignedMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type UnanimousAssertionValidatorRequest struct {
	BeforeHash           *common.HashBuf               `protobuf:"bytes,1,opt,name=beforeHash,proto3" json:"beforeHash,omitempty"`
	BeforeInbox          *common.HashBuf               `protobuf:"bytes,2,opt,name=beforeInbox,proto3" json:"beforeInbox,omitempty"`
	SequenceNum          uint64                        `protobuf:"varint,3,opt,name=sequenceNum,proto3" json:"sequenceNum,omitempty"`
	TimeBounds           *protocol.TimeBoundsBlocksBuf `protobuf:"bytes,4,opt,name=timeBounds,proto3" json:"timeBounds,omitempty"`
	SignedMessages       []*SignedMessage              `protobuf:"bytes,5,rep,name=signedMessages,proto3" json:"signedMessages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UnanimousAssertionValidatorRequest) Reset()         { *m = UnanimousAssertionValidatorRequest{} }
func (m *UnanimousAssertionValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*UnanimousAssertionValidatorRequest) ProtoMessage()    {}
func (*UnanimousAssertionValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{4}
}

func (m *UnanimousAssertionValidatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnanimousAssertionValidatorRequest.Unmarshal(m, b)
}
func (m *UnanimousAssertionValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnanimousAssertionValidatorRequest.Marshal(b, m, deterministic)
}
func (m *UnanimousAssertionValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnanimousAssertionValidatorRequest.Merge(m, src)
}
func (m *UnanimousAssertionValidatorRequest) XXX_Size() int {
	return xxx_messageInfo_UnanimousAssertionValidatorRequest.Size(m)
}
func (m *UnanimousAssertionValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnanimousAssertionValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnanimousAssertionValidatorRequest proto.InternalMessageInfo

func (m *UnanimousAssertionValidatorRequest) GetBeforeHash() *common.HashBuf {
	if m != nil {
		return m.BeforeHash
	}
	return nil
}

func (m *UnanimousAssertionValidatorRequest) GetBeforeInbox() *common.HashBuf {
	if m != nil {
		return m.BeforeInbox
	}
	return nil
}

func (m *UnanimousAssertionValidatorRequest) GetSequenceNum() uint64 {
	if m != nil {
		return m.SequenceNum
	}
	return 0
}

func (m *UnanimousAssertionValidatorRequest) GetTimeBounds() *protocol.TimeBoundsBlocksBuf {
	if m != nil {
		return m.TimeBounds
	}
	return nil
}

func (m *UnanimousAssertionValidatorRequest) GetSignedMessages() []*SignedMessage {
	if m != nil {
		return m.SignedMessages
	}
	return nil
}

type ValidatorRequest struct {
	RequestId *common.HashBuf `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*ValidatorRequest_Unanimous
	//	*ValidatorRequest_UnanimousNotification
	Request              isValidatorRequest_Request `protobuf_oneof:"request"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ValidatorRequest) Reset()         { *m = ValidatorRequest{} }
func (m *ValidatorRequest) String() string { return proto.CompactTextString(m) }
func (*ValidatorRequest) ProtoMessage()    {}
func (*ValidatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{5}
}

func (m *ValidatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorRequest.Unmarshal(m, b)
}
func (m *ValidatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorRequest.Marshal(b, m, deterministic)
}
func (m *ValidatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorRequest.Merge(m, src)
}
func (m *ValidatorRequest) XXX_Size() int {
	return xxx_messageInfo_ValidatorRequest.Size(m)
}
func (m *ValidatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorRequest proto.InternalMessageInfo

func (m *ValidatorRequest) GetRequestId() *common.HashBuf {
	if m != nil {
		return m.RequestId
	}
	return nil
}

type isValidatorRequest_Request interface {
	isValidatorRequest_Request()
}

type ValidatorRequest_Unanimous struct {
	Unanimous *UnanimousAssertionValidatorRequest `protobuf:"bytes,2,opt,name=unanimous,proto3,oneof"`
}

type ValidatorRequest_UnanimousNotification struct {
	UnanimousNotification *UnanimousAssertionValidatorNotification `protobuf:"bytes,3,opt,name=unanimousNotification,proto3,oneof"`
}

func (*ValidatorRequest_Unanimous) isValidatorRequest_Request() {}

func (*ValidatorRequest_UnanimousNotification) isValidatorRequest_Request() {}

func (m *ValidatorRequest) GetRequest() isValidatorRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ValidatorRequest) GetUnanimous() *UnanimousAssertionValidatorRequest {
	if x, ok := m.GetRequest().(*ValidatorRequest_Unanimous); ok {
		return x.Unanimous
	}
	return nil
}

func (m *ValidatorRequest) GetUnanimousNotification() *UnanimousAssertionValidatorNotification {
	if x, ok := m.GetRequest().(*ValidatorRequest_UnanimousNotification); ok {
		return x.UnanimousNotification
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ValidatorRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ValidatorRequest_Unanimous)(nil),
		(*ValidatorRequest_UnanimousNotification)(nil),
	}
}

type UnanimousAssertionFollowerResponse struct {
	Accepted             bool            `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Signature            []byte          `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	AssertionHash        *common.HashBuf `protobuf:"bytes,3,opt,name=assertion_hash,json=assertionHash,proto3" json:"assertion_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UnanimousAssertionFollowerResponse) Reset()         { *m = UnanimousAssertionFollowerResponse{} }
func (m *UnanimousAssertionFollowerResponse) String() string { return proto.CompactTextString(m) }
func (*UnanimousAssertionFollowerResponse) ProtoMessage()    {}
func (*UnanimousAssertionFollowerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{6}
}

func (m *UnanimousAssertionFollowerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnanimousAssertionFollowerResponse.Unmarshal(m, b)
}
func (m *UnanimousAssertionFollowerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnanimousAssertionFollowerResponse.Marshal(b, m, deterministic)
}
func (m *UnanimousAssertionFollowerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnanimousAssertionFollowerResponse.Merge(m, src)
}
func (m *UnanimousAssertionFollowerResponse) XXX_Size() int {
	return xxx_messageInfo_UnanimousAssertionFollowerResponse.Size(m)
}
func (m *UnanimousAssertionFollowerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnanimousAssertionFollowerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnanimousAssertionFollowerResponse proto.InternalMessageInfo

func (m *UnanimousAssertionFollowerResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *UnanimousAssertionFollowerResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *UnanimousAssertionFollowerResponse) GetAssertionHash() *common.HashBuf {
	if m != nil {
		return m.AssertionHash
	}
	return nil
}

type FollowerResponse struct {
	RequestId            *common.HashBuf                     `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Unanimous            *UnanimousAssertionFollowerResponse `protobuf:"bytes,3,opt,name=unanimous,proto3" json:"unanimous,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *FollowerResponse) Reset()         { *m = FollowerResponse{} }
func (m *FollowerResponse) String() string { return proto.CompactTextString(m) }
func (*FollowerResponse) ProtoMessage()    {}
func (*FollowerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b34ccd35396e2606, []int{7}
}

func (m *FollowerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FollowerResponse.Unmarshal(m, b)
}
func (m *FollowerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FollowerResponse.Marshal(b, m, deterministic)
}
func (m *FollowerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FollowerResponse.Merge(m, src)
}
func (m *FollowerResponse) XXX_Size() int {
	return xxx_messageInfo_FollowerResponse.Size(m)
}
func (m *FollowerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FollowerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FollowerResponse proto.InternalMessageInfo

func (m *FollowerResponse) GetRequestId() *common.HashBuf {
	if m != nil {
		return m.RequestId
	}
	return nil
}

func (m *FollowerResponse) GetUnanimous() *UnanimousAssertionFollowerResponse {
	if m != nil {
		return m.Unanimous
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenTypeBuf)(nil), "valmessage.TokenTypeBuf")
	proto.RegisterType((*VMConfiguration)(nil), "valmessage.VMConfiguration")
	proto.RegisterType((*UnanimousAssertionValidatorNotification)(nil), "valmessage.UnanimousAssertionValidatorNotification")
	proto.RegisterType((*SignedMessage)(nil), "valmessage.SignedMessage")
	proto.RegisterType((*UnanimousAssertionValidatorRequest)(nil), "valmessage.UnanimousAssertionValidatorRequest")
	proto.RegisterType((*ValidatorRequest)(nil), "valmessage.ValidatorRequest")
	proto.RegisterType((*UnanimousAssertionFollowerResponse)(nil), "valmessage.UnanimousAssertionFollowerResponse")
	proto.RegisterType((*FollowerResponse)(nil), "valmessage.FollowerResponse")
}

func init() { proto.RegisterFile("valmessage.proto", fileDescriptor_b34ccd35396e2606) }

var fileDescriptor_b34ccd35396e2606 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x6d, 0x92, 0xa6, 0x6d, 0x6e, 0xd2, 0x1f, 0xcd, 0xf7, 0x55, 0x35, 0x15, 0x45, 0xc1, 0x42,
	0x22, 0x1b, 0x12, 0xb5, 0x95, 0x60, 0x81, 0x40, 0xaa, 0x0b, 0xa8, 0x15, 0xb4, 0x42, 0x6e, 0xe9,
	0x82, 0x8d, 0x99, 0xd8, 0x37, 0xce, 0x28, 0xf6, 0x8c, 0x99, 0xf1, 0xa4, 0xc9, 0x2b, 0xb0, 0x62,
	0xc5, 0x53, 0xf0, 0x22, 0xbc, 0x15, 0xf2, 0x4f, 0x1c, 0xd3, 0x3f, 0xca, 0xca, 0x9e, 0x33, 0xe7,
	0xdc, 0x3b, 0x73, 0xee, 0x19, 0xd8, 0x18, 0xd3, 0x20, 0x44, 0xa5, 0xa8, 0x8f, 0xdd, 0x48, 0x8a,
	0x58, 0x10, 0x98, 0x23, 0xdb, 0x3b, 0x63, 0x1a, 0xa4, 0xa8, 0x2b, 0x82, 0x5e, 0xe9, 0x3f, 0xa3,
	0x6e, 0x6f, 0x15, 0x7b, 0x57, 0x36, 0xfe, 0x73, 0x45, 0x18, 0x0a, 0xde, 0xcb, 0x3e, 0x19, 0x68,
	0x3e, 0x81, 0xd6, 0xb9, 0x18, 0x21, 0x3f, 0x9f, 0x46, 0x68, 0xe9, 0x01, 0xf9, 0x1f, 0xea, 0x63,
	0x1a, 0x68, 0x34, 0x2a, 0xed, 0x4a, 0xa7, 0x65, 0x67, 0x0b, 0xf3, 0x57, 0x15, 0xd6, 0x2f, 0x4e,
	0x0e, 0x05, 0x1f, 0x30, 0x5f, 0x4b, 0x1a, 0x33, 0xc1, 0xc9, 0x63, 0x68, 0xf9, 0x92, 0xba, 0xe8,
	0x44, 0x28, 0x99, 0xf0, 0x52, 0xc1, 0xa2, 0xdd, 0x4c, 0xb1, 0x8f, 0x29, 0x44, 0x5e, 0xc3, 0x3a,
	0x2a, 0x57, 0x8a, 0x4b, 0x47, 0xe2, 0x57, 0xcd, 0x24, 0x7a, 0x46, 0xb5, 0x5d, 0xe9, 0x34, 0xf7,
	0x36, 0xbb, 0xf9, 0x21, 0x2c, 0xe6, 0x1f, 0xf3, 0x18, 0x7d, 0x94, 0x96, 0x1e, 0xd8, 0x6b, 0x19,
	0xdb, 0xce, 0xc9, 0xe4, 0x65, 0xa1, 0x77, 0xb5, 0x94, 0xc8, 0xdd, 0xa9, 0x51, 0x4b, 0xf5, 0x64,
	0xa6, 0x3f, 0xf0, 0x3c, 0x89, 0x4a, 0x95, 0xc4, 0x87, 0x39, 0x93, 0xec, 0x43, 0x93, 0x2a, 0x85,
	0x32, 0x76, 0x46, 0x38, 0x55, 0xc6, 0x62, 0xbb, 0x76, 0x8b, 0x10, 0x32, 0xda, 0x7b, 0x9c, 0x2a,
	0xf2, 0x02, 0x8c, 0x90, 0x4e, 0x1c, 0x9c, 0xa0, 0xab, 0x93, 0x5b, 0x3a, 0x2a, 0xc6, 0xc8, 0x71,
	0x85, 0xe6, 0xb1, 0x51, 0x6f, 0x57, 0x3a, 0xab, 0xf6, 0x66, 0x48, 0x27, 0x6f, 0x67, 0xdb, 0x67,
	0x31, 0x46, 0x87, 0xc9, 0x26, 0xe9, 0x40, 0x5d, 0x5c, 0x72, 0x94, 0xc6, 0xd2, 0xad, 0x07, 0xcc,
	0x08, 0x26, 0xc2, 0xd3, 0x4f, 0x9c, 0x72, 0x16, 0x0a, 0xad, 0x0e, 0xd2, 0xce, 0x4c, 0xf0, 0x0b,
	0x1a, 0x30, 0x8f, 0xc6, 0x42, 0x9e, 0x8a, 0x98, 0x0d, 0x98, 0x9b, 0x59, 0xbc, 0x0d, 0x2b, 0xd4,
	0x75, 0x31, 0x8a, 0x31, 0xb3, 0x77, 0xc5, 0x2e, 0xd6, 0xe4, 0x11, 0x80, 0x62, 0x3e, 0xa7, 0xb1,
	0x96, 0xa8, 0x8c, 0x6a, 0xbb, 0xd6, 0x69, 0xd9, 0x25, 0xc4, 0xfc, 0x02, 0xab, 0x67, 0xcc, 0xe7,
	0xe8, 0x9d, 0x64, 0xb1, 0x21, 0xbb, 0xb0, 0x9c, 0x27, 0x28, 0xad, 0xd5, 0xdc, 0xdb, 0xea, 0x96,
	0xc3, 0x93, 0xd3, 0x92, 0x83, 0xce, 0x78, 0xe4, 0x21, 0x34, 0x8a, 0x8a, 0xe9, 0xe4, 0x5a, 0xf6,
	0x1c, 0x30, 0x7f, 0x56, 0xc1, 0xbc, 0xe3, 0x26, 0xc9, 0x14, 0x51, 0xc5, 0xa4, 0x07, 0xd0, 0xc7,
	0x81, 0x90, 0x78, 0x44, 0xd5, 0x30, 0x6f, 0xbd, 0x3e, 0xb3, 0x27, 0xc1, 0xd2, 0x19, 0xcc, 0x29,
	0x64, 0x17, 0x9a, 0xd9, 0xea, 0x98, 0xf7, 0xc5, 0x24, 0x4f, 0xcc, 0x35, 0x45, 0x99, 0x43, 0xda,
	0xd0, 0x54, 0x49, 0x3b, 0xee, 0xe2, 0xa9, 0x0e, 0xd3, 0x90, 0x2c, 0xda, 0x65, 0x88, 0xbc, 0x02,
	0x88, 0x59, 0x88, 0x96, 0xd0, 0xdc, 0x4b, 0xc2, 0x90, 0xd4, 0xdc, 0xe9, 0x16, 0xb7, 0x3f, 0x2f,
	0xf6, 0xac, 0x40, 0xb8, 0xa3, 0x2c, 0x17, 0x73, 0x01, 0x39, 0x80, 0x35, 0x55, 0x76, 0x53, 0x19,
	0xf5, 0x34, 0x4f, 0x0f, 0xba, 0xa5, 0xa7, 0xfa, 0x87, 0xdf, 0xf6, 0x15, 0x81, 0xf9, 0xad, 0x0a,
	0x1b, 0xd7, 0xcc, 0xe9, 0x02, 0xc8, 0xec, 0xd7, 0x61, 0xde, 0x6d, 0xe6, 0x34, 0x72, 0xca, 0xb1,
	0x47, 0x4e, 0xa1, 0xa1, 0x67, 0x96, 0xe7, 0xce, 0x74, 0xcb, 0x47, 0xf8, 0xfb, 0x3c, 0x8e, 0x16,
	0xec, 0x79, 0x09, 0x32, 0x82, 0xcd, 0x62, 0x51, 0x8e, 0x5e, 0xfe, 0xce, 0xf6, 0xef, 0x59, 0xbb,
	0x2c, 0x3d, 0x5a, 0xb0, 0x6f, 0xae, 0x69, 0x35, 0x60, 0x39, 0xbf, 0x89, 0xf9, 0xa3, 0x72, 0x53,
	0x76, 0xde, 0x89, 0x20, 0x10, 0x97, 0x28, 0x6d, 0x54, 0x91, 0xe0, 0x0a, 0xef, 0x7c, 0x00, 0x77,
	0x86, 0x93, 0x3c, 0x87, 0x35, 0x3a, 0x2b, 0xeb, 0x0c, 0x93, 0xe4, 0xd5, 0x6e, 0x36, 0x77, 0xb5,
	0xa0, 0x25, 0x88, 0xf9, 0xbd, 0x02, 0x1b, 0xd7, 0x8e, 0xf1, 0xaf, 0x53, 0xfa, 0x50, 0x9e, 0x52,
	0xed, 0x3e, 0x53, 0xba, 0xda, 0xb2, 0x34, 0x23, 0xeb, 0xcd, 0x67, 0xcb, 0x67, 0xf1, 0x50, 0xf7,
	0x93, 0x8e, 0x3d, 0x31, 0x18, 0xb8, 0x43, 0xca, 0x78, 0x40, 0xfb, 0xaa, 0x47, 0x65, 0x9f, 0xc5,
	0x52, 0x87, 0xbd, 0x88, 0xba, 0xa3, 0x24, 0x65, 0x09, 0xf2, 0x6c, 0x3c, 0x1b, 0x4b, 0x6f, 0xde,
	0xb3, 0xbf, 0x94, 0x66, 0x7d, 0xff, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0xd3, 0x48, 0x4d,
	0x5c, 0x06, 0x00, 0x00,
}