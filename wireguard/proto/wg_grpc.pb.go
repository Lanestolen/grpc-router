// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WireguardClient is the client API for Wireguard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WireguardClient interface {
	InitializeI(ctx context.Context, in *IReq, opts ...grpc.CallOption) (*IResp, error)
	AddPeer(ctx context.Context, in *AddPReq, opts ...grpc.CallOption) (*AddPResp, error)
	DelPeer(ctx context.Context, in *DelPReq, opts ...grpc.CallOption) (*DelPResp, error)
	ListPeers(ctx context.Context, in *ListPeersReq, opts ...grpc.CallOption) (*ListPeersResp, error)
	ManageNIC(ctx context.Context, in *ManageNICReq, opts ...grpc.CallOption) (*ManageNICResp, error)
	GetPeerStatus(ctx context.Context, in *PeerStatusReq, opts ...grpc.CallOption) (*PeerStatusResp, error)
	GetNICInfo(ctx context.Context, in *NICInfoReq, opts ...grpc.CallOption) (*NICInfoResp, error)
	GenPublicKey(ctx context.Context, in *PubKeyReq, opts ...grpc.CallOption) (*PubKeyResp, error)
	GenPrivateKey(ctx context.Context, in *PrivKeyReq, opts ...grpc.CallOption) (*PrivKeyResp, error)
	GetPrivateKey(ctx context.Context, in *PrivKeyReq, opts ...grpc.CallOption) (*PrivKeyResp, error)
	GetPublicKey(ctx context.Context, in *PubKeyReq, opts ...grpc.CallOption) (*PubKeyResp, error)
}

type wireguardClient struct {
	cc grpc.ClientConnInterface
}

func NewWireguardClient(cc grpc.ClientConnInterface) WireguardClient {
	return &wireguardClient{cc}
}

func (c *wireguardClient) InitializeI(ctx context.Context, in *IReq, opts ...grpc.CallOption) (*IResp, error) {
	out := new(IResp)
	err := c.cc.Invoke(ctx, "/Wireguard/InitializeI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) AddPeer(ctx context.Context, in *AddPReq, opts ...grpc.CallOption) (*AddPResp, error) {
	out := new(AddPResp)
	err := c.cc.Invoke(ctx, "/Wireguard/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) DelPeer(ctx context.Context, in *DelPReq, opts ...grpc.CallOption) (*DelPResp, error) {
	out := new(DelPResp)
	err := c.cc.Invoke(ctx, "/Wireguard/DelPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) ListPeers(ctx context.Context, in *ListPeersReq, opts ...grpc.CallOption) (*ListPeersResp, error) {
	out := new(ListPeersResp)
	err := c.cc.Invoke(ctx, "/Wireguard/ListPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) ManageNIC(ctx context.Context, in *ManageNICReq, opts ...grpc.CallOption) (*ManageNICResp, error) {
	out := new(ManageNICResp)
	err := c.cc.Invoke(ctx, "/Wireguard/ManageNIC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) GetPeerStatus(ctx context.Context, in *PeerStatusReq, opts ...grpc.CallOption) (*PeerStatusResp, error) {
	out := new(PeerStatusResp)
	err := c.cc.Invoke(ctx, "/Wireguard/GetPeerStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) GetNICInfo(ctx context.Context, in *NICInfoReq, opts ...grpc.CallOption) (*NICInfoResp, error) {
	out := new(NICInfoResp)
	err := c.cc.Invoke(ctx, "/Wireguard/GetNICInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) GenPublicKey(ctx context.Context, in *PubKeyReq, opts ...grpc.CallOption) (*PubKeyResp, error) {
	out := new(PubKeyResp)
	err := c.cc.Invoke(ctx, "/Wireguard/GenPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) GenPrivateKey(ctx context.Context, in *PrivKeyReq, opts ...grpc.CallOption) (*PrivKeyResp, error) {
	out := new(PrivKeyResp)
	err := c.cc.Invoke(ctx, "/Wireguard/GenPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) GetPrivateKey(ctx context.Context, in *PrivKeyReq, opts ...grpc.CallOption) (*PrivKeyResp, error) {
	out := new(PrivKeyResp)
	err := c.cc.Invoke(ctx, "/Wireguard/GetPrivateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wireguardClient) GetPublicKey(ctx context.Context, in *PubKeyReq, opts ...grpc.CallOption) (*PubKeyResp, error) {
	out := new(PubKeyResp)
	err := c.cc.Invoke(ctx, "/Wireguard/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WireguardServer is the server API for Wireguard service.
// All implementations must embed UnimplementedWireguardServer
// for forward compatibility
type WireguardServer interface {
	InitializeI(context.Context, *IReq) (*IResp, error)
	AddPeer(context.Context, *AddPReq) (*AddPResp, error)
	DelPeer(context.Context, *DelPReq) (*DelPResp, error)
	ListPeers(context.Context, *ListPeersReq) (*ListPeersResp, error)
	ManageNIC(context.Context, *ManageNICReq) (*ManageNICResp, error)
	GetPeerStatus(context.Context, *PeerStatusReq) (*PeerStatusResp, error)
	GetNICInfo(context.Context, *NICInfoReq) (*NICInfoResp, error)
	GenPublicKey(context.Context, *PubKeyReq) (*PubKeyResp, error)
	GenPrivateKey(context.Context, *PrivKeyReq) (*PrivKeyResp, error)
	GetPrivateKey(context.Context, *PrivKeyReq) (*PrivKeyResp, error)
	GetPublicKey(context.Context, *PubKeyReq) (*PubKeyResp, error)
	mustEmbedUnimplementedWireguardServer()
}

// UnimplementedWireguardServer must be embedded to have forward compatible implementations.
type UnimplementedWireguardServer struct {
}

func (UnimplementedWireguardServer) InitializeI(context.Context, *IReq) (*IResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeI not implemented")
}
func (UnimplementedWireguardServer) AddPeer(context.Context, *AddPReq) (*AddPResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (UnimplementedWireguardServer) DelPeer(context.Context, *DelPReq) (*DelPResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelPeer not implemented")
}
func (UnimplementedWireguardServer) ListPeers(context.Context, *ListPeersReq) (*ListPeersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPeers not implemented")
}
func (UnimplementedWireguardServer) ManageNIC(context.Context, *ManageNICReq) (*ManageNICResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageNIC not implemented")
}
func (UnimplementedWireguardServer) GetPeerStatus(context.Context, *PeerStatusReq) (*PeerStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeerStatus not implemented")
}
func (UnimplementedWireguardServer) GetNICInfo(context.Context, *NICInfoReq) (*NICInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNICInfo not implemented")
}
func (UnimplementedWireguardServer) GenPublicKey(context.Context, *PubKeyReq) (*PubKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenPublicKey not implemented")
}
func (UnimplementedWireguardServer) GenPrivateKey(context.Context, *PrivKeyReq) (*PrivKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenPrivateKey not implemented")
}
func (UnimplementedWireguardServer) GetPrivateKey(context.Context, *PrivKeyReq) (*PrivKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}
func (UnimplementedWireguardServer) GetPublicKey(context.Context, *PubKeyReq) (*PubKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}
func (UnimplementedWireguardServer) mustEmbedUnimplementedWireguardServer() {}

// UnsafeWireguardServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WireguardServer will
// result in compilation errors.
type UnsafeWireguardServer interface {
	mustEmbedUnimplementedWireguardServer()
}

func RegisterWireguardServer(s grpc.ServiceRegistrar, srv WireguardServer) {
	s.RegisterService(&Wireguard_ServiceDesc, srv)
}

func _Wireguard_InitializeI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).InitializeI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/InitializeI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).InitializeI(ctx, req.(*IReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).AddPeer(ctx, req.(*AddPReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_DelPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelPReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).DelPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/DelPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).DelPeer(ctx, req.(*DelPReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_ListPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPeersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).ListPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/ListPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).ListPeers(ctx, req.(*ListPeersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_ManageNIC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageNICReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).ManageNIC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/ManageNIC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).ManageNIC(ctx, req.(*ManageNICReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_GetPeerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).GetPeerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/GetPeerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).GetPeerStatus(ctx, req.(*PeerStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_GetNICInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NICInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).GetNICInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/GetNICInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).GetNICInfo(ctx, req.(*NICInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_GenPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).GenPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/GenPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).GenPublicKey(ctx, req.(*PubKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_GenPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).GenPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/GenPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).GenPrivateKey(ctx, req.(*PrivKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrivKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/GetPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).GetPrivateKey(ctx, req.(*PrivKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wireguard_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WireguardServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Wireguard/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WireguardServer).GetPublicKey(ctx, req.(*PubKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Wireguard_ServiceDesc is the grpc.ServiceDesc for Wireguard service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wireguard_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Wireguard",
	HandlerType: (*WireguardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeI",
			Handler:    _Wireguard_InitializeI_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _Wireguard_AddPeer_Handler,
		},
		{
			MethodName: "DelPeer",
			Handler:    _Wireguard_DelPeer_Handler,
		},
		{
			MethodName: "ListPeers",
			Handler:    _Wireguard_ListPeers_Handler,
		},
		{
			MethodName: "ManageNIC",
			Handler:    _Wireguard_ManageNIC_Handler,
		},
		{
			MethodName: "GetPeerStatus",
			Handler:    _Wireguard_GetPeerStatus_Handler,
		},
		{
			MethodName: "GetNICInfo",
			Handler:    _Wireguard_GetNICInfo_Handler,
		},
		{
			MethodName: "GenPublicKey",
			Handler:    _Wireguard_GenPublicKey_Handler,
		},
		{
			MethodName: "GenPrivateKey",
			Handler:    _Wireguard_GenPrivateKey_Handler,
		},
		{
			MethodName: "GetPrivateKey",
			Handler:    _Wireguard_GetPrivateKey_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _Wireguard_GetPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wg.proto",
}