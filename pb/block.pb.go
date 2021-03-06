// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

package pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_meshplus_bitxhub_kit_types "github.com/meshplus/bitxhub-kit/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Block struct {
	BlockHeader          *BlockHeader                               `protobuf:"bytes,1,opt,name=block_header,json=blockHeader,proto3" json:"block_header,omitempty"`
	Transactions         []*Transaction                             `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	BlockHash            github_com_meshplus_bitxhub_kit_types.Hash `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"block_hash"`
	Signature            []byte                                     `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	Extra                []byte                                     `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return m.Size()
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetBlockHeader() *BlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Block) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

type BlockHeader struct {
	Number               uint64                                     `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	StateRoot            github_com_meshplus_bitxhub_kit_types.Hash `protobuf:"bytes,2,opt,name=state_root,json=stateRoot,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"state_root"`
	TxRoot               github_com_meshplus_bitxhub_kit_types.Hash `protobuf:"bytes,3,opt,name=tx_root,json=txRoot,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"tx_root"`
	ReceiptRoot          github_com_meshplus_bitxhub_kit_types.Hash `protobuf:"bytes,4,opt,name=receipt_root,json=receiptRoot,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"receipt_root"`
	ParentHash           github_com_meshplus_bitxhub_kit_types.Hash `protobuf:"bytes,5,opt,name=parent_hash,json=parentHash,proto3,customtype=github.com/meshplus/bitxhub-kit/types.Hash" json:"parent_hash"`
	Timestamp            int64                                      `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Version              []byte                                     `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{1}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetVersion() []byte {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterType((*BlockHeader)(nil), "pb.BlockHeader")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0xaa, 0xdb, 0x30,
	0x18, 0x85, 0xab, 0x38, 0x76, 0xb8, 0xb2, 0xa1, 0x54, 0x94, 0x22, 0x2e, 0x25, 0xd7, 0xdc, 0x29,
	0x14, 0x62, 0x43, 0xf2, 0x06, 0x99, 0x02, 0x9d, 0xea, 0xb6, 0x73, 0x90, 0x5c, 0xd5, 0x16, 0x89,
	0x2d, 0x21, 0xfd, 0x2e, 0xee, 0x23, 0xf5, 0x4d, 0x32, 0x76, 0xee, 0x10, 0x4a, 0x5e, 0xa2, 0x6b,
	0x89, 0xe4, 0x90, 0xdc, 0xd5, 0x9b, 0xcf, 0xf1, 0xaf, 0x4f, 0xe7, 0xe7, 0x08, 0xc7, 0xfc, 0xa0,
	0xca, 0x7d, 0xa6, 0x8d, 0x02, 0x45, 0x26, 0x9a, 0x3f, 0xbe, 0x01, 0xc3, 0x5a, 0xcb, 0x4a, 0x90,
	0xaa, 0xf5, 0xf6, 0xe3, 0xb2, 0x92, 0x50, 0x77, 0x3c, 0x2b, 0x55, 0x93, 0x57, 0xaa, 0x52, 0xb9,
	0xb3, 0x79, 0xf7, 0xdd, 0x29, 0x27, 0xdc, 0x97, 0x1f, 0x7f, 0xfe, 0x87, 0x70, 0xb8, 0xb9, 0x50,
	0xc9, 0x0a, 0x27, 0x0e, 0xbf, 0xab, 0x05, 0xfb, 0x26, 0x0c, 0x45, 0x29, 0x5a, 0xc4, 0xab, 0xd7,
	0x99, 0xe6, 0x99, 0x1b, 0xd8, 0x3a, 0xbb, 0xf0, 0x19, 0xbc, 0x20, 0x6b, 0x9c, 0xdc, 0x25, 0xb0,
	0x74, 0x92, 0x06, 0xd7, 0x33, 0x5f, 0x6e, 0x7e, 0xf1, 0x62, 0x88, 0x7c, 0xc2, 0x78, 0xb8, 0x88,
	0xd9, 0x9a, 0x06, 0x29, 0x5a, 0x24, 0x9b, 0xd5, 0xf1, 0xf4, 0xf4, 0xea, 0xcf, 0xe9, 0xe9, 0xc3,
	0x5d, 0xfa, 0x46, 0xd8, 0x5a, 0x1f, 0x3a, 0x9b, 0x73, 0x09, 0x7d, 0xdd, 0xf1, 0xe5, 0x5e, 0x42,
	0x0e, 0x3f, 0xb5, 0xb0, 0xd9, 0x96, 0xd9, 0xba, 0x78, 0xf0, 0x49, 0x98, 0xad, 0xc9, 0x7b, 0xfc,
	0x60, 0x65, 0xd5, 0x32, 0xe8, 0x8c, 0xa0, 0xd3, 0x0b, 0xb1, 0xb8, 0x19, 0xe4, 0x2d, 0x0e, 0x45,
	0x0f, 0x86, 0xd1, 0xd0, 0xfd, 0xf1, 0xe2, 0xf9, 0x57, 0x80, 0xe3, 0xbb, 0xc5, 0xc8, 0x3b, 0x1c,
	0xb5, 0x5d, 0xc3, 0x87, 0xcd, 0xa7, 0xc5, 0xa0, 0x2e, 0x71, 0x2d, 0x30, 0x10, 0x3b, 0xa3, 0x14,
	0xd0, 0xc9, 0xf8, 0xb8, 0x8e, 0x52, 0x28, 0x05, 0xe4, 0x23, 0x9e, 0x41, 0xef, 0x79, 0xe3, 0xd7,
	0x8f, 0xa0, 0x77, 0xb0, 0xaf, 0x38, 0x31, 0xa2, 0x14, 0x52, 0x83, 0x27, 0x4e, 0x47, 0x13, 0xe3,
	0x81, 0xe3, 0xb0, 0x9f, 0x71, 0xac, 0x99, 0x11, 0x2d, 0xf8, 0x9a, 0xc2, 0xd1, 0x54, 0xec, 0x31,
	0xd7, 0x9e, 0x40, 0x36, 0xc2, 0x02, 0x6b, 0x34, 0x8d, 0x52, 0xb4, 0x08, 0x8a, 0x9b, 0x41, 0x28,
	0x9e, 0xfd, 0x10, 0xc6, 0x4a, 0xd5, 0xd2, 0x99, 0x6b, 0xea, 0x2a, 0x37, 0xc9, 0xf1, 0x3c, 0x47,
	0xbf, 0xcf, 0x73, 0xf4, 0xf7, 0x3c, 0x47, 0x3c, 0x72, 0x4f, 0x77, 0xfd, 0x3f, 0x00, 0x00, 0xff,
	0xff, 0xa2, 0x6f, 0xb9, 0xa7, 0x0f, 0x03, 0x00, 0x00,
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Extra) > 0 {
		i -= len(m.Extra)
		copy(dAtA[i:], m.Extra)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Extra)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.BlockHash.Size()
		i -= size
		if _, err := m.BlockHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Transactions) > 0 {
		for iNdEx := len(m.Transactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Transactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlock(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeader != nil {
		{
			size, err := m.BlockHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Timestamp != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.ParentHash.Size()
		i -= size
		if _, err := m.ParentHash.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.ReceiptRoot.Size()
		i -= size
		if _, err := m.ReceiptRoot.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.TxRoot.Size()
		i -= size
		if _, err := m.TxRoot.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.StateRoot.Size()
		i -= size
		if _, err := m.StateRoot.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBlock(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Number != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeader != nil {
		l = m.BlockHeader.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if len(m.Transactions) > 0 {
		for _, e := range m.Transactions {
			l = e.Size()
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	l = m.BlockHash.Size()
	n += 1 + l + sovBlock(uint64(l))
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.Extra)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BlockHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovBlock(uint64(m.Number))
	}
	l = m.StateRoot.Size()
	n += 1 + l + sovBlock(uint64(l))
	l = m.TxRoot.Size()
	n += 1 + l + sovBlock(uint64(l))
	l = m.ReceiptRoot.Size()
	n += 1 + l + sovBlock(uint64(l))
	l = m.ParentHash.Size()
	n += 1 + l + sovBlock(uint64(l))
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockHeader == nil {
				m.BlockHeader = &BlockHeader{}
			}
			if err := m.BlockHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transactions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transactions = append(m.Transactions, &Transaction{})
			if err := m.Transactions[len(m.Transactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BlockHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extra", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extra = append(m.Extra[:0], dAtA[iNdEx:postIndex]...)
			if m.Extra == nil {
				m.Extra = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StateRoot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxRoot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiptRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReceiptRoot.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ParentHash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = append(m.Version[:0], dAtA[iNdEx:postIndex]...)
			if m.Version == nil {
				m.Version = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBlock
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
