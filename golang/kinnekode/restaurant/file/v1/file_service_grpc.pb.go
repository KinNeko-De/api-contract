// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: kinnekode/restaurant/file/v1/file_service.proto

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
	FileService_StoreFile_FullMethodName        = "/kinnekode.restaurant.file.v1.FileService/StoreFile"
	FileService_StoreRevision_FullMethodName    = "/kinnekode.restaurant.file.v1.FileService/StoreRevision"
	FileService_DownloadFile_FullMethodName     = "/kinnekode.restaurant.file.v1.FileService/DownloadFile"
	FileService_DownloadRevision_FullMethodName = "/kinnekode.restaurant.file.v1.FileService/DownloadRevision"
)

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	// Stores a new file
	StoreFile(ctx context.Context, opts ...grpc.CallOption) (FileService_StoreFileClient, error)
	// Stores a new revision to an existing file
	StoreRevision(ctx context.Context, opts ...grpc.CallOption) (FileService_StoreRevisionClient, error)
	// Download the latest revision of the file
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (FileService_DownloadFileClient, error)
	// Download a specific revision of the file
	DownloadRevision(ctx context.Context, in *DownloadRevisionRequest, opts ...grpc.CallOption) (FileService_DownloadRevisionClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) StoreFile(ctx context.Context, opts ...grpc.CallOption) (FileService_StoreFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], FileService_StoreFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceStoreFileClient{stream}
	return x, nil
}

type FileService_StoreFileClient interface {
	Send(*StoreFileRequest) error
	CloseAndRecv() (*StoreFileResponse, error)
	grpc.ClientStream
}

type fileServiceStoreFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceStoreFileClient) Send(m *StoreFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceStoreFileClient) CloseAndRecv() (*StoreFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StoreFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) StoreRevision(ctx context.Context, opts ...grpc.CallOption) (FileService_StoreRevisionClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[1], FileService_StoreRevision_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceStoreRevisionClient{stream}
	return x, nil
}

type FileService_StoreRevisionClient interface {
	Send(*StoreRevisionRequest) error
	CloseAndRecv() (*StoreFileResponse, error)
	grpc.ClientStream
}

type fileServiceStoreRevisionClient struct {
	grpc.ClientStream
}

func (x *fileServiceStoreRevisionClient) Send(m *StoreRevisionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceStoreRevisionClient) CloseAndRecv() (*StoreFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StoreFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (FileService_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[2], FileService_DownloadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_DownloadFileClient interface {
	Recv() (*DownloadFileResponse, error)
	grpc.ClientStream
}

type fileServiceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceDownloadFileClient) Recv() (*DownloadFileResponse, error) {
	m := new(DownloadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) DownloadRevision(ctx context.Context, in *DownloadRevisionRequest, opts ...grpc.CallOption) (FileService_DownloadRevisionClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[3], FileService_DownloadRevision_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceDownloadRevisionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileService_DownloadRevisionClient interface {
	Recv() (*DownloadFileResponse, error)
	grpc.ClientStream
}

type fileServiceDownloadRevisionClient struct {
	grpc.ClientStream
}

func (x *fileServiceDownloadRevisionClient) Recv() (*DownloadFileResponse, error) {
	m := new(DownloadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	// Stores a new file
	StoreFile(FileService_StoreFileServer) error
	// Stores a new revision to an existing file
	StoreRevision(FileService_StoreRevisionServer) error
	// Download the latest revision of the file
	DownloadFile(*DownloadFileRequest, FileService_DownloadFileServer) error
	// Download a specific revision of the file
	DownloadRevision(*DownloadRevisionRequest, FileService_DownloadRevisionServer) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) StoreFile(FileService_StoreFileServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreFile not implemented")
}
func (UnimplementedFileServiceServer) StoreRevision(FileService_StoreRevisionServer) error {
	return status.Errorf(codes.Unimplemented, "method StoreRevision not implemented")
}
func (UnimplementedFileServiceServer) DownloadFile(*DownloadFileRequest, FileService_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileServiceServer) DownloadRevision(*DownloadRevisionRequest, FileService_DownloadRevisionServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadRevision not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_StoreFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).StoreFile(&fileServiceStoreFileServer{stream})
}

type FileService_StoreFileServer interface {
	SendAndClose(*StoreFileResponse) error
	Recv() (*StoreFileRequest, error)
	grpc.ServerStream
}

type fileServiceStoreFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceStoreFileServer) SendAndClose(m *StoreFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceStoreFileServer) Recv() (*StoreFileRequest, error) {
	m := new(StoreFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileService_StoreRevision_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).StoreRevision(&fileServiceStoreRevisionServer{stream})
}

type FileService_StoreRevisionServer interface {
	SendAndClose(*StoreFileResponse) error
	Recv() (*StoreRevisionRequest, error)
	grpc.ServerStream
}

type fileServiceStoreRevisionServer struct {
	grpc.ServerStream
}

func (x *fileServiceStoreRevisionServer) SendAndClose(m *StoreFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceStoreRevisionServer) Recv() (*StoreRevisionRequest, error) {
	m := new(StoreRevisionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownloadFile(m, &fileServiceDownloadFileServer{stream})
}

type FileService_DownloadFileServer interface {
	Send(*DownloadFileResponse) error
	grpc.ServerStream
}

type fileServiceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceDownloadFileServer) Send(m *DownloadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FileService_DownloadRevision_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRevisionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileServiceServer).DownloadRevision(m, &fileServiceDownloadRevisionServer{stream})
}

type FileService_DownloadRevisionServer interface {
	Send(*DownloadFileResponse) error
	grpc.ServerStream
}

type fileServiceDownloadRevisionServer struct {
	grpc.ServerStream
}

func (x *fileServiceDownloadRevisionServer) Send(m *DownloadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kinnekode.restaurant.file.v1.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StoreFile",
			Handler:       _FileService_StoreFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "StoreRevision",
			Handler:       _FileService_StoreRevision_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _FileService_DownloadFile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DownloadRevision",
			Handler:       _FileService_DownloadRevision_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kinnekode/restaurant/file/v1/file_service.proto",
}
