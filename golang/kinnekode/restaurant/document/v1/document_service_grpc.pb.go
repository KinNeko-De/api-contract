// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: kinnekode/restaurant/document/v1/document_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DocumentService_GeneratePreview_FullMethodName = "/kinnekode.restaurant.document.v1.DocumentService/GeneratePreview"
)

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentServiceClient interface {
	// Generates a preview of a document and send the document directly back
	// Documents are not stored
	GeneratePreview(ctx context.Context, in *GeneratePreviewRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GeneratePreviewResponse], error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) GeneratePreview(ctx context.Context, in *GeneratePreviewRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GeneratePreviewResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DocumentService_ServiceDesc.Streams[0], DocumentService_GeneratePreview_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GeneratePreviewRequest, GeneratePreviewResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DocumentService_GeneratePreviewClient = grpc.ServerStreamingClient[GeneratePreviewResponse]

// DocumentServiceServer is the server API for DocumentService service.
// All implementations must embed UnimplementedDocumentServiceServer
// for forward compatibility.
type DocumentServiceServer interface {
	// Generates a preview of a document and send the document directly back
	// Documents are not stored
	GeneratePreview(*GeneratePreviewRequest, grpc.ServerStreamingServer[GeneratePreviewResponse]) error
	mustEmbedUnimplementedDocumentServiceServer()
}

// UnimplementedDocumentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDocumentServiceServer struct{}

func (UnimplementedDocumentServiceServer) GeneratePreview(*GeneratePreviewRequest, grpc.ServerStreamingServer[GeneratePreviewResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GeneratePreview not implemented")
}
func (UnimplementedDocumentServiceServer) mustEmbedUnimplementedDocumentServiceServer() {}
func (UnimplementedDocumentServiceServer) testEmbeddedByValue()                         {}

// UnsafeDocumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServiceServer will
// result in compilation errors.
type UnsafeDocumentServiceServer interface {
	mustEmbedUnimplementedDocumentServiceServer()
}

func RegisterDocumentServiceServer(s grpc.ServiceRegistrar, srv DocumentServiceServer) {
	// If the following call pancis, it indicates UnimplementedDocumentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DocumentService_ServiceDesc, srv)
}

func _DocumentService_GeneratePreview_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GeneratePreviewRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DocumentServiceServer).GeneratePreview(m, &grpc.GenericServerStream[GeneratePreviewRequest, GeneratePreviewResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DocumentService_GeneratePreviewServer = grpc.ServerStreamingServer[GeneratePreviewResponse]

// DocumentService_ServiceDesc is the grpc.ServiceDesc for DocumentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kinnekode.restaurant.document.v1.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GeneratePreview",
			Handler:       _DocumentService_GeneratePreview_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kinnekode/restaurant/document/v1/document_service.proto",
}
