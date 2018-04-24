// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// grpccmd imports
import (
	"io"

	"github.com/spf13/cobra"
	"github.com/sshaman1101/grpccmd"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OrderType int32

const (
	OrderType_ANY OrderType = 0
	OrderType_BID OrderType = 1
	OrderType_ASK OrderType = 2
)

var OrderType_name = map[int32]string{
	0: "ANY",
	1: "BID",
	2: "ASK",
}
var OrderType_value = map[string]int32{
	"ANY": 0,
	"BID": 1,
	"ASK": 2,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

type OrderStatus int32

const (
	OrderStatus_ORDER_INACTIVE OrderStatus = 0
	OrderStatus_ORDER_ACTIVE   OrderStatus = 1
)

var OrderStatus_name = map[int32]string{
	0: "ORDER_INACTIVE",
	1: "ORDER_ACTIVE",
}
var OrderStatus_value = map[string]int32{
	"ORDER_INACTIVE": 0,
	"ORDER_ACTIVE":   1,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}
func (OrderStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

type IdentityLevel int32

const (
	IdentityLevel_ANONYMOUS    IdentityLevel = 0
	IdentityLevel_PSEUDONYMOUS IdentityLevel = 1
	IdentityLevel_IDENTIFIED   IdentityLevel = 2
)

var IdentityLevel_name = map[int32]string{
	0: "ANONYMOUS",
	1: "PSEUDONYMOUS",
	2: "IDENTIFIED",
}
var IdentityLevel_value = map[string]int32{
	"ANONYMOUS":    0,
	"PSEUDONYMOUS": 1,
	"IDENTIFIED":   2,
}

func (x IdentityLevel) String() string {
	return proto.EnumName(IdentityLevel_name, int32(x))
}
func (IdentityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

type DealStatus int32

const (
	DealStatus_DEAL_UNKNOWN  DealStatus = 0
	DealStatus_DEAL_ACCEPTED DealStatus = 1
	DealStatus_DEAL_CLOSED   DealStatus = 2
)

var DealStatus_name = map[int32]string{
	0: "DEAL_UNKNOWN",
	1: "DEAL_ACCEPTED",
	2: "DEAL_CLOSED",
}
var DealStatus_value = map[string]int32{
	"DEAL_UNKNOWN":  0,
	"DEAL_ACCEPTED": 1,
	"DEAL_CLOSED":   2,
}

func (x DealStatus) String() string {
	return proto.EnumName(DealStatus_name, int32(x))
}
func (DealStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

type ChangeRequestStatus int32

const (
	ChangeRequestStatus_REQUEST_UNKNOWN  ChangeRequestStatus = 0
	ChangeRequestStatus_REQUEST_CREATED  ChangeRequestStatus = 1
	ChangeRequestStatus_REQUEST_CANCELED ChangeRequestStatus = 2
	ChangeRequestStatus_REQUEST_REJECTED ChangeRequestStatus = 3
	ChangeRequestStatus_REQUEST_ACCEPTED ChangeRequestStatus = 4
)

var ChangeRequestStatus_name = map[int32]string{
	0: "REQUEST_UNKNOWN",
	1: "REQUEST_CREATED",
	2: "REQUEST_CANCELED",
	3: "REQUEST_REJECTED",
	4: "REQUEST_ACCEPTED",
}
var ChangeRequestStatus_value = map[string]int32{
	"REQUEST_UNKNOWN":  0,
	"REQUEST_CREATED":  1,
	"REQUEST_CANCELED": 2,
	"REQUEST_REJECTED": 3,
	"REQUEST_ACCEPTED": 4,
}

func (x ChangeRequestStatus) String() string {
	return proto.EnumName(ChangeRequestStatus_name, int32(x))
}
func (ChangeRequestStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

type GetOrdersReply struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *GetOrdersReply) Reset()                    { *m = GetOrdersReply{} }
func (m *GetOrdersReply) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersReply) ProtoMessage()               {}
func (*GetOrdersReply) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *GetOrdersReply) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type Benchmarks struct {
	Values []uint64 `protobuf:"varint,1,rep,packed,name=values" json:"values,omitempty"`
}

func (m *Benchmarks) Reset()                    { *m = Benchmarks{} }
func (m *Benchmarks) String() string            { return proto.CompactTextString(m) }
func (*Benchmarks) ProtoMessage()               {}
func (*Benchmarks) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *Benchmarks) GetValues() []uint64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type Deal struct {
	Id             string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Benchmarks     *Benchmarks `protobuf:"bytes,2,opt,name=benchmarks" json:"benchmarks,omitempty"`
	SupplierID     string      `protobuf:"bytes,3,opt,name=supplierID" json:"supplierID,omitempty"`
	ConsumerID     string      `protobuf:"bytes,4,opt,name=consumerID" json:"consumerID,omitempty"`
	MasterID       string      `protobuf:"bytes,5,opt,name=masterID" json:"masterID,omitempty"`
	AskID          string      `protobuf:"bytes,6,opt,name=askID" json:"askID,omitempty"`
	BidID          string      `protobuf:"bytes,7,opt,name=bidID" json:"bidID,omitempty"`
	Duration       uint64      `protobuf:"varint,8,opt,name=duration" json:"duration,omitempty"`
	Price          *BigInt     `protobuf:"bytes,9,opt,name=price" json:"price,omitempty"`
	StartTime      *Timestamp  `protobuf:"bytes,10,opt,name=startTime" json:"startTime,omitempty"`
	EndTime        *Timestamp  `protobuf:"bytes,11,opt,name=endTime" json:"endTime,omitempty"`
	Status         DealStatus  `protobuf:"varint,12,opt,name=status,enum=sonm.DealStatus" json:"status,omitempty"`
	BlockedBalance *BigInt     `protobuf:"bytes,13,opt,name=blockedBalance" json:"blockedBalance,omitempty"`
	TotalPayout    *BigInt     `protobuf:"bytes,14,opt,name=totalPayout" json:"totalPayout,omitempty"`
	LastBillTS     *Timestamp  `protobuf:"bytes,15,opt,name=lastBillTS" json:"lastBillTS,omitempty"`
}

func (m *Deal) Reset()                    { *m = Deal{} }
func (m *Deal) String() string            { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()               {}
func (*Deal) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *Deal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Deal) GetBenchmarks() *Benchmarks {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

func (m *Deal) GetSupplierID() string {
	if m != nil {
		return m.SupplierID
	}
	return ""
}

func (m *Deal) GetConsumerID() string {
	if m != nil {
		return m.ConsumerID
	}
	return ""
}

func (m *Deal) GetMasterID() string {
	if m != nil {
		return m.MasterID
	}
	return ""
}

func (m *Deal) GetAskID() string {
	if m != nil {
		return m.AskID
	}
	return ""
}

func (m *Deal) GetBidID() string {
	if m != nil {
		return m.BidID
	}
	return ""
}

func (m *Deal) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Deal) GetPrice() *BigInt {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Deal) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Deal) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Deal) GetStatus() DealStatus {
	if m != nil {
		return m.Status
	}
	return DealStatus_DEAL_UNKNOWN
}

func (m *Deal) GetBlockedBalance() *BigInt {
	if m != nil {
		return m.BlockedBalance
	}
	return nil
}

func (m *Deal) GetTotalPayout() *BigInt {
	if m != nil {
		return m.TotalPayout
	}
	return nil
}

func (m *Deal) GetLastBillTS() *Timestamp {
	if m != nil {
		return m.LastBillTS
	}
	return nil
}

type Order struct {
	Id             string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DealID         string        `protobuf:"bytes,2,opt,name=dealID" json:"dealID,omitempty"`
	OrderType      OrderType     `protobuf:"varint,3,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
	OrderStatus    OrderStatus   `protobuf:"varint,4,opt,name=orderStatus,enum=sonm.OrderStatus" json:"orderStatus,omitempty"`
	AuthorID       string        `protobuf:"bytes,5,opt,name=authorID" json:"authorID,omitempty"`
	CounterpartyID string        `protobuf:"bytes,6,opt,name=counterpartyID" json:"counterpartyID,omitempty"`
	Duration       uint64        `protobuf:"varint,7,opt,name=duration" json:"duration,omitempty"`
	Price          *BigInt       `protobuf:"bytes,8,opt,name=price" json:"price,omitempty"`
	Netflags       uint64        `protobuf:"varint,9,opt,name=netflags" json:"netflags,omitempty"`
	IdentityLevel  IdentityLevel `protobuf:"varint,10,opt,name=identityLevel,enum=sonm.IdentityLevel" json:"identityLevel,omitempty"`
	Blacklist      string        `protobuf:"bytes,11,opt,name=blacklist" json:"blacklist,omitempty"`
	Tag            []byte        `protobuf:"bytes,12,opt,name=tag,proto3" json:"tag,omitempty"`
	Benchmarks     *Benchmarks   `protobuf:"bytes,13,opt,name=benchmarks" json:"benchmarks,omitempty"`
	FrozenSum      *BigInt       `protobuf:"bytes,14,opt,name=frozenSum" json:"frozenSum,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetDealID() string {
	if m != nil {
		return m.DealID
	}
	return ""
}

func (m *Order) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_ANY
}

func (m *Order) GetOrderStatus() OrderStatus {
	if m != nil {
		return m.OrderStatus
	}
	return OrderStatus_ORDER_INACTIVE
}

func (m *Order) GetAuthorID() string {
	if m != nil {
		return m.AuthorID
	}
	return ""
}

func (m *Order) GetCounterpartyID() string {
	if m != nil {
		return m.CounterpartyID
	}
	return ""
}

func (m *Order) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Order) GetPrice() *BigInt {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Order) GetNetflags() uint64 {
	if m != nil {
		return m.Netflags
	}
	return 0
}

func (m *Order) GetIdentityLevel() IdentityLevel {
	if m != nil {
		return m.IdentityLevel
	}
	return IdentityLevel_ANONYMOUS
}

func (m *Order) GetBlacklist() string {
	if m != nil {
		return m.Blacklist
	}
	return ""
}

func (m *Order) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *Order) GetBenchmarks() *Benchmarks {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

func (m *Order) GetFrozenSum() *BigInt {
	if m != nil {
		return m.FrozenSum
	}
	return nil
}

type BidNetwork struct {
	Overlay  bool `protobuf:"varint,1,opt,name=overlay" json:"overlay,omitempty"`
	Outbound bool `protobuf:"varint,2,opt,name=outbound" json:"outbound,omitempty"`
	Incoming bool `protobuf:"varint,3,opt,name=incoming" json:"incoming,omitempty"`
}

func (m *BidNetwork) Reset()                    { *m = BidNetwork{} }
func (m *BidNetwork) String() string            { return proto.CompactTextString(m) }
func (*BidNetwork) ProtoMessage()               {}
func (*BidNetwork) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *BidNetwork) GetOverlay() bool {
	if m != nil {
		return m.Overlay
	}
	return false
}

func (m *BidNetwork) GetOutbound() bool {
	if m != nil {
		return m.Outbound
	}
	return false
}

func (m *BidNetwork) GetIncoming() bool {
	if m != nil {
		return m.Incoming
	}
	return false
}

type BidResources struct {
	Network    *BidNetwork       `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	Benchmarks map[string]uint64 `protobuf:"bytes,2,rep,name=benchmarks" json:"benchmarks,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *BidResources) Reset()                    { *m = BidResources{} }
func (m *BidResources) String() string            { return proto.CompactTextString(m) }
func (*BidResources) ProtoMessage()               {}
func (*BidResources) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *BidResources) GetNetwork() *BidNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *BidResources) GetBenchmarks() map[string]uint64 {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

type BidOrder struct {
	ID             string        `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Duration       *Duration     `protobuf:"bytes,2,opt,name=duration" json:"duration,omitempty"`
	Price          *Price        `protobuf:"bytes,3,opt,name=price" json:"price,omitempty"`
	Blacklist      *EthAddress   `protobuf:"bytes,4,opt,name=blacklist" json:"blacklist,omitempty"`
	Identity       IdentityLevel `protobuf:"varint,5,opt,name=identity,enum=sonm.IdentityLevel" json:"identity,omitempty"`
	Tag            string        `protobuf:"bytes,6,opt,name=tag" json:"tag,omitempty"`
	CounterpartyID *EthAddress   `protobuf:"bytes,7,opt,name=CounterpartyID" json:"CounterpartyID,omitempty"`
	Resources      *BidResources `protobuf:"bytes,8,opt,name=resources" json:"resources,omitempty"`
}

func (m *BidOrder) Reset()                    { *m = BidOrder{} }
func (m *BidOrder) String() string            { return proto.CompactTextString(m) }
func (*BidOrder) ProtoMessage()               {}
func (*BidOrder) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *BidOrder) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *BidOrder) GetDuration() *Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *BidOrder) GetPrice() *Price {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *BidOrder) GetBlacklist() *EthAddress {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *BidOrder) GetIdentity() IdentityLevel {
	if m != nil {
		return m.Identity
	}
	return IdentityLevel_ANONYMOUS
}

func (m *BidOrder) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *BidOrder) GetCounterpartyID() *EthAddress {
	if m != nil {
		return m.CounterpartyID
	}
	return nil
}

func (m *BidOrder) GetResources() *BidResources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOrdersReply)(nil), "sonm.GetOrdersReply")
	proto.RegisterType((*Benchmarks)(nil), "sonm.Benchmarks")
	proto.RegisterType((*Deal)(nil), "sonm.Deal")
	proto.RegisterType((*Order)(nil), "sonm.Order")
	proto.RegisterType((*BidNetwork)(nil), "sonm.BidNetwork")
	proto.RegisterType((*BidResources)(nil), "sonm.BidResources")
	proto.RegisterType((*BidOrder)(nil), "sonm.BidOrder")
	proto.RegisterEnum("sonm.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("sonm.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterEnum("sonm.IdentityLevel", IdentityLevel_name, IdentityLevel_value)
	proto.RegisterEnum("sonm.DealStatus", DealStatus_name, DealStatus_value)
	proto.RegisterEnum("sonm.ChangeRequestStatus", ChangeRequestStatus_name, ChangeRequestStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Market service

type MarketClient interface {
	// GetOrders returns orders by given filter parameters.
	// Note that set of filters may be changed in the closest future.
	GetOrders(ctx context.Context, in *Count, opts ...grpc.CallOption) (*GetOrdersReply, error)
	// CreateOrder places new order on the Marketplace.
	// Note that current impl of Node API prevents you from
	// creating ASKs orders.
	CreateOrder(ctx context.Context, in *BidOrder, opts ...grpc.CallOption) (*Order, error)
	// GetOrderByID returns order by given ID.
	// If order save an `inactive` status returns error instead.
	GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error)
	// CancelOrder removes active order from the Marketplace.
	CancelOrder(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type marketClient struct {
	cc *grpc.ClientConn
}

func NewMarketClient(cc *grpc.ClientConn) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) GetOrders(ctx context.Context, in *Count, opts ...grpc.CallOption) (*GetOrdersReply, error) {
	out := new(GetOrdersReply)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CreateOrder(ctx context.Context, in *BidOrder, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/CreateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrderByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CancelOrder(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Market/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Market service

type MarketServer interface {
	// GetOrders returns orders by given filter parameters.
	// Note that set of filters may be changed in the closest future.
	GetOrders(context.Context, *Count) (*GetOrdersReply, error)
	// CreateOrder places new order on the Marketplace.
	// Note that current impl of Node API prevents you from
	// creating ASKs orders.
	CreateOrder(context.Context, *BidOrder) (*Order, error)
	// GetOrderByID returns order by given ID.
	// If order save an `inactive` status returns error instead.
	GetOrderByID(context.Context, *ID) (*Order, error)
	// CancelOrder removes active order from the Marketplace.
	CancelOrder(context.Context, *ID) (*Empty, error)
}

func RegisterMarketServer(s *grpc.Server, srv MarketServer) {
	s.RegisterService(&_Market_serviceDesc, srv)
}

func _Market_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Count)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrders(ctx, req.(*Count))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CreateOrder(ctx, req.(*BidOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrderByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CancelOrder(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Market_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _Market_GetOrders_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Market_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _Market_GetOrderByID_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Market_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace.proto",
}

// Begin grpccmd
var _ = grpccmd.RunE

// Market
var _MarketCmd = &cobra.Command{
	Use:   "market [method]",
	Short: "Subcommand for the Market service.",
}

var _Market_GetOrdersCmd = &cobra.Command{
	Use:   "getOrders",
	Short: "Make the GetOrders method call, input-type: sonm.Count output-type: sonm.GetOrdersReply",
	RunE: grpccmd.RunE(
		"GetOrders",
		"sonm.Count",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMarketClient(cc)
		},
	),
}

var _Market_GetOrdersCmd_gen = &cobra.Command{
	Use:   "getOrders-gen",
	Short: "Generate JSON for method call of GetOrders (input-type: sonm.Count)",
	RunE:  grpccmd.TypeToJson("sonm.Count"),
}

var _Market_CreateOrderCmd = &cobra.Command{
	Use:   "createOrder",
	Short: "Make the CreateOrder method call, input-type: sonm.BidOrder output-type: sonm.Order",
	RunE: grpccmd.RunE(
		"CreateOrder",
		"sonm.BidOrder",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMarketClient(cc)
		},
	),
}

var _Market_CreateOrderCmd_gen = &cobra.Command{
	Use:   "createOrder-gen",
	Short: "Generate JSON for method call of CreateOrder (input-type: sonm.BidOrder)",
	RunE:  grpccmd.TypeToJson("sonm.BidOrder"),
}

var _Market_GetOrderByIDCmd = &cobra.Command{
	Use:   "getOrderByID",
	Short: "Make the GetOrderByID method call, input-type: sonm.ID output-type: sonm.Order",
	RunE: grpccmd.RunE(
		"GetOrderByID",
		"sonm.ID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMarketClient(cc)
		},
	),
}

var _Market_GetOrderByIDCmd_gen = &cobra.Command{
	Use:   "getOrderByID-gen",
	Short: "Generate JSON for method call of GetOrderByID (input-type: sonm.ID)",
	RunE:  grpccmd.TypeToJson("sonm.ID"),
}

var _Market_CancelOrderCmd = &cobra.Command{
	Use:   "cancelOrder",
	Short: "Make the CancelOrder method call, input-type: sonm.ID output-type: sonm.Empty",
	RunE: grpccmd.RunE(
		"CancelOrder",
		"sonm.ID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMarketClient(cc)
		},
	),
}

var _Market_CancelOrderCmd_gen = &cobra.Command{
	Use:   "cancelOrder-gen",
	Short: "Generate JSON for method call of CancelOrder (input-type: sonm.ID)",
	RunE:  grpccmd.TypeToJson("sonm.ID"),
}

// Register commands with the root command and service command
func init() {
	grpccmd.RegisterServiceCmd(_MarketCmd)
	_MarketCmd.AddCommand(
		_Market_GetOrdersCmd,
		_Market_GetOrdersCmd_gen,
		_Market_CreateOrderCmd,
		_Market_CreateOrderCmd_gen,
		_Market_GetOrderByIDCmd,
		_Market_GetOrderByIDCmd_gen,
		_Market_CancelOrderCmd,
		_Market_CancelOrderCmd_gen,
	)
}

// End grpccmd

func init() { proto.RegisterFile("marketplace.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 1095 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0x5d, 0x6f, 0xe3, 0x44,
	0x17, 0xc7, 0xe3, 0x24, 0xcd, 0xcb, 0x49, 0xe2, 0x64, 0xa7, 0xd5, 0xca, 0x8a, 0x1e, 0x3d, 0x2a,
	0x06, 0x95, 0x10, 0x41, 0x77, 0xd5, 0x82, 0xb4, 0x20, 0x21, 0x91, 0xd8, 0x06, 0x99, 0xed, 0x26,
	0x65, 0x92, 0x82, 0xf6, 0x86, 0xd5, 0x24, 0x9e, 0x4d, 0x47, 0x71, 0xec, 0x60, 0x8f, 0x8b, 0xc2,
	0x1d, 0x9f, 0x87, 0x4f, 0xc0, 0x05, 0xf7, 0x7c, 0x28, 0x2e, 0xd0, 0x8c, 0x5f, 0x32, 0x4d, 0x5b,
	0x89, 0xbb, 0x9c, 0xff, 0xf9, 0xfb, 0xcc, 0xcc, 0x99, 0xdf, 0x4c, 0x06, 0x9e, 0x6d, 0x48, 0xb4,
	0xa6, 0x7c, 0xeb, 0x93, 0x25, 0x3d, 0xdf, 0x46, 0x21, 0x0f, 0x51, 0x35, 0x0e, 0x83, 0x4d, 0xbf,
	0xbd, 0x60, 0x2b, 0x16, 0xf0, 0x54, 0xeb, 0x77, 0x59, 0x20, 0xd4, 0x80, 0x91, 0x5c, 0xe0, 0x6c,
	0x43, 0x63, 0x4e, 0x36, 0xdb, 0x54, 0x30, 0xbf, 0x00, 0xfd, 0x3b, 0xca, 0xa7, 0x91, 0x47, 0xa3,
	0x18, 0xd3, 0xad, 0xbf, 0x43, 0x1f, 0x42, 0x2d, 0x94, 0xa1, 0xa1, 0x9d, 0x56, 0x06, 0xad, 0x8b,
	0xd6, 0xb9, 0x28, 0x71, 0x2e, 0x2d, 0x38, 0x4b, 0x99, 0x1f, 0x01, 0x8c, 0x69, 0xb0, 0xbc, 0x15,
	0xd3, 0x88, 0xd1, 0x73, 0xa8, 0xdd, 0x11, 0x3f, 0xa1, 0xe9, 0x27, 0x55, 0x9c, 0x45, 0xe6, 0x1f,
	0x55, 0xa8, 0xda, 0x94, 0xf8, 0x48, 0x87, 0x32, 0xf3, 0x0c, 0xed, 0x54, 0x1b, 0x34, 0x71, 0x99,
	0x79, 0xe8, 0x25, 0xc0, 0xa2, 0xf8, 0xdc, 0x28, 0x9f, 0x6a, 0x83, 0xd6, 0x45, 0x2f, 0x1d, 0x67,
	0x5f, 0x16, 0x2b, 0x1e, 0xf4, 0x7f, 0x80, 0x38, 0xd9, 0x6e, 0x7d, 0x46, 0x23, 0xd7, 0x36, 0x2a,
	0xb2, 0x92, 0xa2, 0x88, 0xfc, 0x32, 0x0c, 0xe2, 0x64, 0x23, 0xf3, 0xd5, 0x34, 0xbf, 0x57, 0x50,
	0x1f, 0x1a, 0x1b, 0x12, 0x73, 0x99, 0x3d, 0x92, 0xd9, 0x22, 0x46, 0x27, 0x70, 0x44, 0xe2, 0xb5,
	0x6b, 0x1b, 0x35, 0x99, 0x48, 0x03, 0xa1, 0x2e, 0x98, 0xe7, 0xda, 0x46, 0x3d, 0x55, 0x65, 0x20,
	0xea, 0x78, 0x49, 0x44, 0x38, 0x0b, 0x03, 0xa3, 0x71, 0xaa, 0x0d, 0xaa, 0xb8, 0x88, 0x91, 0x09,
	0x47, 0xdb, 0x88, 0x2d, 0xa9, 0xd1, 0x94, 0x0b, 0x6a, 0x67, 0x0b, 0x62, 0x2b, 0x37, 0xe0, 0x38,
	0x4d, 0xa1, 0xcf, 0xa0, 0x19, 0x73, 0x12, 0xf1, 0x39, 0xdb, 0x50, 0x03, 0xa4, 0xaf, 0x9b, 0xfa,
	0xe6, 0xf9, 0xce, 0xe0, 0xbd, 0x03, 0x7d, 0x02, 0x75, 0x1a, 0x78, 0xd2, 0xdc, 0x7a, 0xdc, 0x9c,
	0xe7, 0xd1, 0x00, 0x6a, 0x31, 0x27, 0x3c, 0x89, 0x8d, 0xf6, 0xa9, 0x36, 0xd0, 0xf3, 0x7e, 0x8a,
	0xfe, 0xcf, 0xa4, 0x8e, 0xb3, 0x3c, 0xfa, 0x1c, 0xf4, 0x85, 0x1f, 0x2e, 0xd7, 0xd4, 0x1b, 0x13,
	0x9f, 0x04, 0x4b, 0x6a, 0x74, 0x1e, 0x99, 0xf0, 0x81, 0x07, 0x9d, 0x43, 0x8b, 0x87, 0x9c, 0xf8,
	0xd7, 0x64, 0x17, 0x26, 0xdc, 0xd0, 0x1f, 0xf9, 0x44, 0x35, 0xa0, 0x17, 0x00, 0x3e, 0x89, 0xf9,
	0x98, 0xf9, 0xfe, 0x7c, 0x66, 0x74, 0x1f, 0x9f, 0xbd, 0x62, 0x31, 0xff, 0xa9, 0xc0, 0x91, 0xa4,
	0xec, 0x01, 0x2e, 0xcf, 0xa1, 0xe6, 0x51, 0xe2, 0xbb, 0xb6, 0x44, 0xa5, 0x89, 0xb3, 0x48, 0x34,
	0x53, 0xf2, 0x38, 0xdf, 0x6d, 0xa9, 0x64, 0x42, 0xcf, 0x47, 0x98, 0xe6, 0x32, 0xde, 0x3b, 0xd0,
	0x25, 0xb4, 0x64, 0x90, 0xb6, 0x43, 0x42, 0xa2, 0x5f, 0x3c, 0x53, 0x3e, 0xc8, 0xfa, 0xa4, 0xba,
	0xc4, 0x86, 0x93, 0x84, 0xdf, 0x86, 0x0a, 0x38, 0x79, 0x8c, 0xce, 0x40, 0x5f, 0x86, 0x49, 0xc0,
	0x69, 0xb4, 0x25, 0x11, 0xdf, 0x15, 0x04, 0x1d, 0xa8, 0xf7, 0xa0, 0xa9, 0x3f, 0x05, 0x4d, 0xe3,
	0x69, 0x68, 0xfa, 0xd0, 0x08, 0x28, 0x7f, 0xef, 0x93, 0x55, 0x2c, 0xd9, 0xaa, 0xe2, 0x22, 0x46,
	0x5f, 0x42, 0x87, 0x79, 0x34, 0xe0, 0x8c, 0xef, 0xae, 0xe8, 0x1d, 0xf5, 0x25, 0x54, 0xfa, 0xc5,
	0x71, 0x5a, 0xc7, 0x55, 0x53, 0xf8, 0xbe, 0x13, 0xfd, 0x0f, 0x9a, 0x0b, 0x9f, 0x2c, 0xd7, 0x3e,
	0x8b, 0xb9, 0xc4, 0xab, 0x89, 0xf7, 0x02, 0xea, 0x41, 0x85, 0x93, 0x95, 0x84, 0xa9, 0x8d, 0xc5,
	0xcf, 0x83, 0x53, 0xdb, 0xf9, 0x0f, 0xa7, 0x76, 0x08, 0xcd, 0xf7, 0x51, 0xf8, 0x1b, 0x0d, 0x66,
	0xc9, 0xe6, 0x51, 0x62, 0xf6, 0x69, 0xf3, 0x67, 0x80, 0x31, 0xf3, 0x26, 0x94, 0xff, 0x1a, 0x46,
	0x6b, 0x64, 0x40, 0x3d, 0xbc, 0xa3, 0x91, 0x4f, 0x76, 0x92, 0x83, 0x06, 0xce, 0x43, 0xd1, 0x8c,
	0x30, 0xe1, 0x8b, 0x30, 0x09, 0x3c, 0x89, 0x43, 0x03, 0x17, 0xb1, 0xc8, 0xb1, 0x60, 0x19, 0x6e,
	0x58, 0xb0, 0x92, 0x3c, 0x34, 0x70, 0x11, 0x9b, 0x7f, 0x69, 0xd0, 0x1e, 0x33, 0x0f, 0xd3, 0x38,
	0x4c, 0xa2, 0x25, 0x15, 0x93, 0xab, 0x07, 0xe9, 0x68, 0x72, 0x88, 0xfd, 0x5a, 0x8a, 0x59, 0xe0,
	0xdc, 0x80, 0xc6, 0x07, 0x17, 0x96, 0xb8, 0x18, 0xcd, 0xc2, 0x5e, 0xd4, 0x54, 0xfa, 0xe0, 0x04,
	0x3c, 0xda, 0xa9, 0xcd, 0xe8, 0x7f, 0x0d, 0xdd, 0x83, 0xb4, 0xe8, 0xf1, 0x9a, 0xee, 0x32, 0xd2,
	0xc5, 0x4f, 0x71, 0xeb, 0xc8, 0xcb, 0x53, 0x2e, 0xad, 0x8a, 0xd3, 0xe0, 0xab, 0xf2, 0x2b, 0xcd,
	0xfc, 0xbb, 0x0c, 0x8d, 0x31, 0xf3, 0x8a, 0x13, 0xe2, 0xda, 0xf9, 0x09, 0x71, 0x6d, 0x34, 0x54,
	0x08, 0x4b, 0xaf, 0x53, 0x3d, 0x3b, 0xfe, 0x99, 0xaa, 0x10, 0xf7, 0x41, 0x4e, 0x5c, 0x45, 0x1a,
	0xb3, 0xfb, 0xfd, 0x5a, 0x48, 0x39, 0x70, 0xe7, 0x2a, 0x19, 0x55, 0xb5, 0x39, 0x0e, 0xbf, 0x1d,
	0x79, 0x5e, 0x44, 0xe3, 0x58, 0x65, 0xe5, 0x05, 0x34, 0x72, 0xb4, 0xe4, 0x21, 0x79, 0x82, 0xbf,
	0xc2, 0x94, 0xc3, 0x95, 0x1e, 0x17, 0x09, 0xd7, 0x2b, 0xd0, 0xad, 0xfb, 0x67, 0xa9, 0xfe, 0xc4,
	0xb8, 0x07, 0x3e, 0xf4, 0x12, 0x9a, 0x51, 0xbe, 0x01, 0xd9, 0x29, 0x42, 0x0f, 0xb7, 0x06, 0xef,
	0x4d, 0xc3, 0x33, 0x68, 0x16, 0x17, 0x04, 0xaa, 0x43, 0x65, 0x34, 0x79, 0xdb, 0x2b, 0x89, 0x1f,
	0x63, 0xd7, 0xee, 0x69, 0x52, 0x99, 0xbd, 0xee, 0x95, 0x87, 0x97, 0xd0, 0x52, 0xee, 0x05, 0x84,
	0x40, 0x9f, 0x62, 0xdb, 0xc1, 0xef, 0xdc, 0xc9, 0xc8, 0x9a, 0xbb, 0x3f, 0x3a, 0xbd, 0x12, 0xea,
	0x41, 0x3b, 0xd5, 0x32, 0x45, 0x1b, 0x7e, 0x03, 0x9d, 0x7b, 0xab, 0x46, 0x1d, 0x68, 0x8e, 0x26,
	0xd3, 0xc9, 0xdb, 0x37, 0xd3, 0x9b, 0x59, 0xfa, 0xc5, 0xf5, 0xcc, 0xb9, 0xb1, 0x73, 0x45, 0x43,
	0x3a, 0x80, 0x6b, 0x3b, 0x93, 0xb9, 0xfb, 0xad, 0xeb, 0xd8, 0xbd, 0xf2, 0x70, 0x0c, 0xb0, 0xbf,
	0xb5, 0x85, 0xdf, 0x76, 0x46, 0x57, 0xef, 0x6e, 0x26, 0xaf, 0x27, 0xd3, 0x9f, 0x26, 0xbd, 0x12,
	0x7a, 0x06, 0x1d, 0xa9, 0x8c, 0x2c, 0xcb, 0xb9, 0x9e, 0x3b, 0x62, 0xca, 0x5d, 0x68, 0x49, 0xc9,
	0xba, 0x9a, 0xce, 0x64, 0x8d, 0xdf, 0x35, 0x38, 0xb6, 0x6e, 0x49, 0xb0, 0xa2, 0x98, 0xfe, 0x92,
	0xd0, 0x98, 0x67, 0xd5, 0x8e, 0xa1, 0x8b, 0x9d, 0x1f, 0x6e, 0x9c, 0xd9, 0x5c, 0x29, 0xa8, 0x88,
	0x16, 0x76, 0x46, 0x69, 0xc9, 0x13, 0xe8, 0x15, 0xe2, 0x68, 0x62, 0x39, 0x57, 0xa2, 0xae, 0xaa,
	0x62, 0xe7, 0x7b, 0xc7, 0x12, 0xde, 0x8a, 0xaa, 0x16, 0x93, 0xaa, 0x5e, 0xfc, 0xa9, 0x41, 0xed,
	0x8d, 0x7c, 0xa7, 0x88, 0x3d, 0x2a, 0x9e, 0x19, 0x28, 0x23, 0x4e, 0x6e, 0x63, 0xff, 0x24, 0x0d,
	0xee, 0x3f, 0x42, 0xcc, 0x12, 0xfa, 0x14, 0x5a, 0x56, 0x44, 0x09, 0xa7, 0x19, 0xf0, 0xc5, 0x8e,
	0xca, 0xb8, 0xaf, 0xbe, 0x4a, 0xcc, 0x12, 0xfa, 0x18, 0xda, 0x79, 0x85, 0xb1, 0x60, 0xa2, 0x91,
	0xe1, 0x67, 0x1f, 0x1a, 0xcf, 0xa0, 0x65, 0x89, 0xbf, 0x33, 0x3f, 0x2d, 0xfb, 0xc0, 0xe7, 0x6c,
	0xb6, 0x7c, 0x67, 0x96, 0x16, 0x35, 0xf9, 0x3c, 0xba, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0xf5, 0x94, 0xc0, 0x69, 0x09, 0x00, 0x00,
}
