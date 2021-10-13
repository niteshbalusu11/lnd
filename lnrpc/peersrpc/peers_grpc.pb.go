// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package peersrpc

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

// PeersClient is the client API for Peers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeersClient interface {
	// lncli: `updatenodeannouncement`
	//UpdateNodeAnnouncement allows the caller to update the node parameters
	//and broadcasts a new version of the node announcement to its peers.
	UpdateNodeAnnouncement(ctx context.Context, in *NodeAnnouncementUpdateRequest, opts ...grpc.CallOption) (*NodeAnnouncementUpdateResponse, error)
}

type peersClient struct {
	cc grpc.ClientConnInterface
}

func NewPeersClient(cc grpc.ClientConnInterface) PeersClient {
	return &peersClient{cc}
}

func (c *peersClient) UpdateNodeAnnouncement(ctx context.Context, in *NodeAnnouncementUpdateRequest, opts ...grpc.CallOption) (*NodeAnnouncementUpdateResponse, error) {
	out := new(NodeAnnouncementUpdateResponse)
	err := c.cc.Invoke(ctx, "/peersrpc.Peers/UpdateNodeAnnouncement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeersServer is the server API for Peers service.
// All implementations must embed UnimplementedPeersServer
// for forward compatibility
type PeersServer interface {
	// lncli: `updatenodeannouncement`
	//UpdateNodeAnnouncement allows the caller to update the node parameters
	//and broadcasts a new version of the node announcement to its peers.
	UpdateNodeAnnouncement(context.Context, *NodeAnnouncementUpdateRequest) (*NodeAnnouncementUpdateResponse, error)
	mustEmbedUnimplementedPeersServer()
}

// UnimplementedPeersServer must be embedded to have forward compatible implementations.
type UnimplementedPeersServer struct {
}

func (UnimplementedPeersServer) UpdateNodeAnnouncement(context.Context, *NodeAnnouncementUpdateRequest) (*NodeAnnouncementUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeAnnouncement not implemented")
}
func (UnimplementedPeersServer) mustEmbedUnimplementedPeersServer() {}

// UnsafePeersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeersServer will
// result in compilation errors.
type UnsafePeersServer interface {
	mustEmbedUnimplementedPeersServer()
}

func RegisterPeersServer(s grpc.ServiceRegistrar, srv PeersServer) {
	s.RegisterService(&Peers_ServiceDesc, srv)
}

func _Peers_UpdateNodeAnnouncement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeAnnouncementUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeersServer).UpdateNodeAnnouncement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peersrpc.Peers/UpdateNodeAnnouncement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeersServer).UpdateNodeAnnouncement(ctx, req.(*NodeAnnouncementUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Peers_ServiceDesc is the grpc.ServiceDesc for Peers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Peers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "peersrpc.Peers",
	HandlerType: (*PeersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateNodeAnnouncement",
			Handler:    _Peers_UpdateNodeAnnouncement_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peersrpc/peers.proto",
}
