// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.0
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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DocumentService_GeneratePreview_FullMethodName = "/kinnekode.restaurant.document.v1.DocumentService/GeneratePreview"
)

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentServiceClient interface {
	// Generates a preview of a document and send the document directly back
	// Documents are not stored
	GeneratePreview(ctx context.Context, in *GeneratePreviewRequest, opts ...grpc.CallOption) (DocumentService_GeneratePreviewClient, error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) GeneratePreview(ctx context.Context, in *GeneratePreviewRequest, opts ...grpc.CallOption) (DocumentService_GeneratePreviewClient, error) {
	stream, err := c.cc.NewStream(ctx, &DocumentService_ServiceDesc.Streams[0], DocumentService_GeneratePreview_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &documentServiceGeneratePreviewClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DocumentService_GeneratePreviewClient interface {
	Recv() (*GeneratePreviewResponse, error)
	grpc.ClientStream
}

type documentServiceGeneratePreviewClient struct {
	grpc.ClientStream
}

func (x *documentServiceGeneratePreviewClient) Recv() (*GeneratePreviewResponse, error) {
	m := new(GeneratePreviewResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DocumentServiceServer is the server API for DocumentService service.
// All implementations must embed UnimplementedDocumentServiceServer
// for forward compatibility
type DocumentServiceServer interface {
	// Generates a preview of a document and send the document directly back
	// Documents are not stored
	GeneratePreview(*GeneratePreviewRequest, DocumentService_GeneratePreviewServer) error
	mustEmbedUnimplementedDocumentServiceServer()
}

// UnimplementedDocumentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentServiceServer struct {
}

func (UnimplementedDocumentServiceServer) GeneratePreview(*GeneratePreviewRequest, DocumentService_GeneratePreviewServer) error {
	return status.Errorf(codes.Unimplemented, "method GeneratePreview not implemented")
}
func (UnimplementedDocumentServiceServer) mustEmbedUnimplementedDocumentServiceServer() {}

// UnsafeDocumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServiceServer will
// result in compilation errors.
type UnsafeDocumentServiceServer interface {
	mustEmbedUnimplementedDocumentServiceServer()
}

func RegisterDocumentServiceServer(s grpc.ServiceRegistrar, srv DocumentServiceServer) {
	s.RegisterService(&DocumentService_ServiceDesc, srv)
}

func _DocumentService_GeneratePreview_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GeneratePreviewRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DocumentServiceServer).GeneratePreview(m, &documentServiceGeneratePreviewServer{stream})
}

type DocumentService_GeneratePreviewServer interface {
	Send(*GeneratePreviewResponse) error
	grpc.ServerStream
}

type documentServiceGeneratePreviewServer struct {
	grpc.ServerStream
}

func (x *documentServiceGeneratePreviewServer) Send(m *GeneratePreviewResponse) error {
	return x.ServerStream.SendMsg(m)
}

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
