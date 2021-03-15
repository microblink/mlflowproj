// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mlflow

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

// ModelRegistryServiceClient is the client API for ModelRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelRegistryServiceClient interface {
	// Throws ``RESOURCE_ALREADY_EXISTS`` if a registered model with the given name exists.
	CreateRegisteredModel(ctx context.Context, in *CreateRegisteredModel, opts ...grpc.CallOption) (*CreateRegisteredModel_Response, error)
	RenameRegisteredModel(ctx context.Context, in *RenameRegisteredModel, opts ...grpc.CallOption) (*RenameRegisteredModel_Response, error)
	UpdateRegisteredModel(ctx context.Context, in *UpdateRegisteredModel, opts ...grpc.CallOption) (*UpdateRegisteredModel_Response, error)
	DeleteRegisteredModel(ctx context.Context, in *DeleteRegisteredModel, opts ...grpc.CallOption) (*DeleteRegisteredModel_Response, error)
	GetRegisteredModel(ctx context.Context, in *GetRegisteredModel, opts ...grpc.CallOption) (*GetRegisteredModel_Response, error)
	SearchRegisteredModels(ctx context.Context, in *SearchRegisteredModels, opts ...grpc.CallOption) (*SearchRegisteredModels_Response, error)
	ListRegisteredModels(ctx context.Context, in *ListRegisteredModels, opts ...grpc.CallOption) (*ListRegisteredModels_Response, error)
	GetLatestVersions(ctx context.Context, in *GetLatestVersions, opts ...grpc.CallOption) (*GetLatestVersions_Response, error)
	CreateModelVersion(ctx context.Context, in *CreateModelVersion, opts ...grpc.CallOption) (*CreateModelVersion_Response, error)
	UpdateModelVersion(ctx context.Context, in *UpdateModelVersion, opts ...grpc.CallOption) (*UpdateModelVersion_Response, error)
	TransitionModelVersionStage(ctx context.Context, in *TransitionModelVersionStage, opts ...grpc.CallOption) (*TransitionModelVersionStage_Response, error)
	DeleteModelVersion(ctx context.Context, in *DeleteModelVersion, opts ...grpc.CallOption) (*DeleteModelVersion_Response, error)
	GetModelVersion(ctx context.Context, in *GetModelVersion, opts ...grpc.CallOption) (*GetModelVersion_Response, error)
	SearchModelVersions(ctx context.Context, in *SearchModelVersions, opts ...grpc.CallOption) (*SearchModelVersions_Response, error)
	GetModelVersionDownloadUri(ctx context.Context, in *GetModelVersionDownloadUri, opts ...grpc.CallOption) (*GetModelVersionDownloadUri_Response, error)
	SetRegisteredModelTag(ctx context.Context, in *SetRegisteredModelTag, opts ...grpc.CallOption) (*SetRegisteredModelTag_Response, error)
	SetModelVersionTag(ctx context.Context, in *SetModelVersionTag, opts ...grpc.CallOption) (*SetModelVersionTag_Response, error)
	DeleteRegisteredModelTag(ctx context.Context, in *DeleteRegisteredModelTag, opts ...grpc.CallOption) (*DeleteRegisteredModelTag_Response, error)
	DeleteModelVersionTag(ctx context.Context, in *DeleteModelVersionTag, opts ...grpc.CallOption) (*DeleteModelVersionTag_Response, error)
}

type modelRegistryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelRegistryServiceClient(cc grpc.ClientConnInterface) ModelRegistryServiceClient {
	return &modelRegistryServiceClient{cc}
}

func (c *modelRegistryServiceClient) CreateRegisteredModel(ctx context.Context, in *CreateRegisteredModel, opts ...grpc.CallOption) (*CreateRegisteredModel_Response, error) {
	out := new(CreateRegisteredModel_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/createRegisteredModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) RenameRegisteredModel(ctx context.Context, in *RenameRegisteredModel, opts ...grpc.CallOption) (*RenameRegisteredModel_Response, error) {
	out := new(RenameRegisteredModel_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/renameRegisteredModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) UpdateRegisteredModel(ctx context.Context, in *UpdateRegisteredModel, opts ...grpc.CallOption) (*UpdateRegisteredModel_Response, error) {
	out := new(UpdateRegisteredModel_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/updateRegisteredModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) DeleteRegisteredModel(ctx context.Context, in *DeleteRegisteredModel, opts ...grpc.CallOption) (*DeleteRegisteredModel_Response, error) {
	out := new(DeleteRegisteredModel_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/deleteRegisteredModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) GetRegisteredModel(ctx context.Context, in *GetRegisteredModel, opts ...grpc.CallOption) (*GetRegisteredModel_Response, error) {
	out := new(GetRegisteredModel_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/getRegisteredModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) SearchRegisteredModels(ctx context.Context, in *SearchRegisteredModels, opts ...grpc.CallOption) (*SearchRegisteredModels_Response, error) {
	out := new(SearchRegisteredModels_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/searchRegisteredModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) ListRegisteredModels(ctx context.Context, in *ListRegisteredModels, opts ...grpc.CallOption) (*ListRegisteredModels_Response, error) {
	out := new(ListRegisteredModels_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/listRegisteredModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) GetLatestVersions(ctx context.Context, in *GetLatestVersions, opts ...grpc.CallOption) (*GetLatestVersions_Response, error) {
	out := new(GetLatestVersions_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/getLatestVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) CreateModelVersion(ctx context.Context, in *CreateModelVersion, opts ...grpc.CallOption) (*CreateModelVersion_Response, error) {
	out := new(CreateModelVersion_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/createModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) UpdateModelVersion(ctx context.Context, in *UpdateModelVersion, opts ...grpc.CallOption) (*UpdateModelVersion_Response, error) {
	out := new(UpdateModelVersion_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/updateModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) TransitionModelVersionStage(ctx context.Context, in *TransitionModelVersionStage, opts ...grpc.CallOption) (*TransitionModelVersionStage_Response, error) {
	out := new(TransitionModelVersionStage_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/transitionModelVersionStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) DeleteModelVersion(ctx context.Context, in *DeleteModelVersion, opts ...grpc.CallOption) (*DeleteModelVersion_Response, error) {
	out := new(DeleteModelVersion_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/deleteModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) GetModelVersion(ctx context.Context, in *GetModelVersion, opts ...grpc.CallOption) (*GetModelVersion_Response, error) {
	out := new(GetModelVersion_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/getModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) SearchModelVersions(ctx context.Context, in *SearchModelVersions, opts ...grpc.CallOption) (*SearchModelVersions_Response, error) {
	out := new(SearchModelVersions_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/searchModelVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) GetModelVersionDownloadUri(ctx context.Context, in *GetModelVersionDownloadUri, opts ...grpc.CallOption) (*GetModelVersionDownloadUri_Response, error) {
	out := new(GetModelVersionDownloadUri_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/getModelVersionDownloadUri", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) SetRegisteredModelTag(ctx context.Context, in *SetRegisteredModelTag, opts ...grpc.CallOption) (*SetRegisteredModelTag_Response, error) {
	out := new(SetRegisteredModelTag_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/setRegisteredModelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) SetModelVersionTag(ctx context.Context, in *SetModelVersionTag, opts ...grpc.CallOption) (*SetModelVersionTag_Response, error) {
	out := new(SetModelVersionTag_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/setModelVersionTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) DeleteRegisteredModelTag(ctx context.Context, in *DeleteRegisteredModelTag, opts ...grpc.CallOption) (*DeleteRegisteredModelTag_Response, error) {
	out := new(DeleteRegisteredModelTag_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/deleteRegisteredModelTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelRegistryServiceClient) DeleteModelVersionTag(ctx context.Context, in *DeleteModelVersionTag, opts ...grpc.CallOption) (*DeleteModelVersionTag_Response, error) {
	out := new(DeleteModelVersionTag_Response)
	err := c.cc.Invoke(ctx, "/mlflow.ModelRegistryService/deleteModelVersionTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelRegistryServiceServer is the server API for ModelRegistryService service.
// All implementations must embed UnimplementedModelRegistryServiceServer
// for forward compatibility
type ModelRegistryServiceServer interface {
	// Throws ``RESOURCE_ALREADY_EXISTS`` if a registered model with the given name exists.
	CreateRegisteredModel(context.Context, *CreateRegisteredModel) (*CreateRegisteredModel_Response, error)
	RenameRegisteredModel(context.Context, *RenameRegisteredModel) (*RenameRegisteredModel_Response, error)
	UpdateRegisteredModel(context.Context, *UpdateRegisteredModel) (*UpdateRegisteredModel_Response, error)
	DeleteRegisteredModel(context.Context, *DeleteRegisteredModel) (*DeleteRegisteredModel_Response, error)
	GetRegisteredModel(context.Context, *GetRegisteredModel) (*GetRegisteredModel_Response, error)
	SearchRegisteredModels(context.Context, *SearchRegisteredModels) (*SearchRegisteredModels_Response, error)
	ListRegisteredModels(context.Context, *ListRegisteredModels) (*ListRegisteredModels_Response, error)
	GetLatestVersions(context.Context, *GetLatestVersions) (*GetLatestVersions_Response, error)
	CreateModelVersion(context.Context, *CreateModelVersion) (*CreateModelVersion_Response, error)
	UpdateModelVersion(context.Context, *UpdateModelVersion) (*UpdateModelVersion_Response, error)
	TransitionModelVersionStage(context.Context, *TransitionModelVersionStage) (*TransitionModelVersionStage_Response, error)
	DeleteModelVersion(context.Context, *DeleteModelVersion) (*DeleteModelVersion_Response, error)
	GetModelVersion(context.Context, *GetModelVersion) (*GetModelVersion_Response, error)
	SearchModelVersions(context.Context, *SearchModelVersions) (*SearchModelVersions_Response, error)
	GetModelVersionDownloadUri(context.Context, *GetModelVersionDownloadUri) (*GetModelVersionDownloadUri_Response, error)
	SetRegisteredModelTag(context.Context, *SetRegisteredModelTag) (*SetRegisteredModelTag_Response, error)
	SetModelVersionTag(context.Context, *SetModelVersionTag) (*SetModelVersionTag_Response, error)
	DeleteRegisteredModelTag(context.Context, *DeleteRegisteredModelTag) (*DeleteRegisteredModelTag_Response, error)
	DeleteModelVersionTag(context.Context, *DeleteModelVersionTag) (*DeleteModelVersionTag_Response, error)
	mustEmbedUnimplementedModelRegistryServiceServer()
}

// UnimplementedModelRegistryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedModelRegistryServiceServer struct {
}

func (UnimplementedModelRegistryServiceServer) CreateRegisteredModel(context.Context, *CreateRegisteredModel) (*CreateRegisteredModel_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRegisteredModel not implemented")
}
func (UnimplementedModelRegistryServiceServer) RenameRegisteredModel(context.Context, *RenameRegisteredModel) (*RenameRegisteredModel_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameRegisteredModel not implemented")
}
func (UnimplementedModelRegistryServiceServer) UpdateRegisteredModel(context.Context, *UpdateRegisteredModel) (*UpdateRegisteredModel_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegisteredModel not implemented")
}
func (UnimplementedModelRegistryServiceServer) DeleteRegisteredModel(context.Context, *DeleteRegisteredModel) (*DeleteRegisteredModel_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegisteredModel not implemented")
}
func (UnimplementedModelRegistryServiceServer) GetRegisteredModel(context.Context, *GetRegisteredModel) (*GetRegisteredModel_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRegisteredModel not implemented")
}
func (UnimplementedModelRegistryServiceServer) SearchRegisteredModels(context.Context, *SearchRegisteredModels) (*SearchRegisteredModels_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRegisteredModels not implemented")
}
func (UnimplementedModelRegistryServiceServer) ListRegisteredModels(context.Context, *ListRegisteredModels) (*ListRegisteredModels_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegisteredModels not implemented")
}
func (UnimplementedModelRegistryServiceServer) GetLatestVersions(context.Context, *GetLatestVersions) (*GetLatestVersions_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestVersions not implemented")
}
func (UnimplementedModelRegistryServiceServer) CreateModelVersion(context.Context, *CreateModelVersion) (*CreateModelVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModelVersion not implemented")
}
func (UnimplementedModelRegistryServiceServer) UpdateModelVersion(context.Context, *UpdateModelVersion) (*UpdateModelVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateModelVersion not implemented")
}
func (UnimplementedModelRegistryServiceServer) TransitionModelVersionStage(context.Context, *TransitionModelVersionStage) (*TransitionModelVersionStage_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransitionModelVersionStage not implemented")
}
func (UnimplementedModelRegistryServiceServer) DeleteModelVersion(context.Context, *DeleteModelVersion) (*DeleteModelVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModelVersion not implemented")
}
func (UnimplementedModelRegistryServiceServer) GetModelVersion(context.Context, *GetModelVersion) (*GetModelVersion_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelVersion not implemented")
}
func (UnimplementedModelRegistryServiceServer) SearchModelVersions(context.Context, *SearchModelVersions) (*SearchModelVersions_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchModelVersions not implemented")
}
func (UnimplementedModelRegistryServiceServer) GetModelVersionDownloadUri(context.Context, *GetModelVersionDownloadUri) (*GetModelVersionDownloadUri_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelVersionDownloadUri not implemented")
}
func (UnimplementedModelRegistryServiceServer) SetRegisteredModelTag(context.Context, *SetRegisteredModelTag) (*SetRegisteredModelTag_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRegisteredModelTag not implemented")
}
func (UnimplementedModelRegistryServiceServer) SetModelVersionTag(context.Context, *SetModelVersionTag) (*SetModelVersionTag_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetModelVersionTag not implemented")
}
func (UnimplementedModelRegistryServiceServer) DeleteRegisteredModelTag(context.Context, *DeleteRegisteredModelTag) (*DeleteRegisteredModelTag_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRegisteredModelTag not implemented")
}
func (UnimplementedModelRegistryServiceServer) DeleteModelVersionTag(context.Context, *DeleteModelVersionTag) (*DeleteModelVersionTag_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteModelVersionTag not implemented")
}
func (UnimplementedModelRegistryServiceServer) mustEmbedUnimplementedModelRegistryServiceServer() {}

// UnsafeModelRegistryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelRegistryServiceServer will
// result in compilation errors.
type UnsafeModelRegistryServiceServer interface {
	mustEmbedUnimplementedModelRegistryServiceServer()
}

func RegisterModelRegistryServiceServer(s grpc.ServiceRegistrar, srv ModelRegistryServiceServer) {
	s.RegisterService(&ModelRegistryService_ServiceDesc, srv)
}

func _ModelRegistryService_CreateRegisteredModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegisteredModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).CreateRegisteredModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/createRegisteredModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).CreateRegisteredModel(ctx, req.(*CreateRegisteredModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_RenameRegisteredModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameRegisteredModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).RenameRegisteredModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/renameRegisteredModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).RenameRegisteredModel(ctx, req.(*RenameRegisteredModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_UpdateRegisteredModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegisteredModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).UpdateRegisteredModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/updateRegisteredModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).UpdateRegisteredModel(ctx, req.(*UpdateRegisteredModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_DeleteRegisteredModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegisteredModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).DeleteRegisteredModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/deleteRegisteredModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).DeleteRegisteredModel(ctx, req.(*DeleteRegisteredModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_GetRegisteredModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRegisteredModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).GetRegisteredModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/getRegisteredModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).GetRegisteredModel(ctx, req.(*GetRegisteredModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_SearchRegisteredModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRegisteredModels)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).SearchRegisteredModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/searchRegisteredModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).SearchRegisteredModels(ctx, req.(*SearchRegisteredModels))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_ListRegisteredModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegisteredModels)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).ListRegisteredModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/listRegisteredModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).ListRegisteredModels(ctx, req.(*ListRegisteredModels))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_GetLatestVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestVersions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).GetLatestVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/getLatestVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).GetLatestVersions(ctx, req.(*GetLatestVersions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_CreateModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).CreateModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/createModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).CreateModelVersion(ctx, req.(*CreateModelVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_UpdateModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateModelVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).UpdateModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/updateModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).UpdateModelVersion(ctx, req.(*UpdateModelVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_TransitionModelVersionStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransitionModelVersionStage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).TransitionModelVersionStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/transitionModelVersionStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).TransitionModelVersionStage(ctx, req.(*TransitionModelVersionStage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_DeleteModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModelVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).DeleteModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/deleteModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).DeleteModelVersion(ctx, req.(*DeleteModelVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_GetModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).GetModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/getModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).GetModelVersion(ctx, req.(*GetModelVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_SearchModelVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchModelVersions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).SearchModelVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/searchModelVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).SearchModelVersions(ctx, req.(*SearchModelVersions))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_GetModelVersionDownloadUri_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelVersionDownloadUri)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).GetModelVersionDownloadUri(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/getModelVersionDownloadUri",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).GetModelVersionDownloadUri(ctx, req.(*GetModelVersionDownloadUri))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_SetRegisteredModelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRegisteredModelTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).SetRegisteredModelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/setRegisteredModelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).SetRegisteredModelTag(ctx, req.(*SetRegisteredModelTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_SetModelVersionTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetModelVersionTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).SetModelVersionTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/setModelVersionTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).SetModelVersionTag(ctx, req.(*SetModelVersionTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_DeleteRegisteredModelTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRegisteredModelTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).DeleteRegisteredModelTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/deleteRegisteredModelTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).DeleteRegisteredModelTag(ctx, req.(*DeleteRegisteredModelTag))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelRegistryService_DeleteModelVersionTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteModelVersionTag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelRegistryServiceServer).DeleteModelVersionTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mlflow.ModelRegistryService/deleteModelVersionTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelRegistryServiceServer).DeleteModelVersionTag(ctx, req.(*DeleteModelVersionTag))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelRegistryService_ServiceDesc is the grpc.ServiceDesc for ModelRegistryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelRegistryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mlflow.ModelRegistryService",
	HandlerType: (*ModelRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createRegisteredModel",
			Handler:    _ModelRegistryService_CreateRegisteredModel_Handler,
		},
		{
			MethodName: "renameRegisteredModel",
			Handler:    _ModelRegistryService_RenameRegisteredModel_Handler,
		},
		{
			MethodName: "updateRegisteredModel",
			Handler:    _ModelRegistryService_UpdateRegisteredModel_Handler,
		},
		{
			MethodName: "deleteRegisteredModel",
			Handler:    _ModelRegistryService_DeleteRegisteredModel_Handler,
		},
		{
			MethodName: "getRegisteredModel",
			Handler:    _ModelRegistryService_GetRegisteredModel_Handler,
		},
		{
			MethodName: "searchRegisteredModels",
			Handler:    _ModelRegistryService_SearchRegisteredModels_Handler,
		},
		{
			MethodName: "listRegisteredModels",
			Handler:    _ModelRegistryService_ListRegisteredModels_Handler,
		},
		{
			MethodName: "getLatestVersions",
			Handler:    _ModelRegistryService_GetLatestVersions_Handler,
		},
		{
			MethodName: "createModelVersion",
			Handler:    _ModelRegistryService_CreateModelVersion_Handler,
		},
		{
			MethodName: "updateModelVersion",
			Handler:    _ModelRegistryService_UpdateModelVersion_Handler,
		},
		{
			MethodName: "transitionModelVersionStage",
			Handler:    _ModelRegistryService_TransitionModelVersionStage_Handler,
		},
		{
			MethodName: "deleteModelVersion",
			Handler:    _ModelRegistryService_DeleteModelVersion_Handler,
		},
		{
			MethodName: "getModelVersion",
			Handler:    _ModelRegistryService_GetModelVersion_Handler,
		},
		{
			MethodName: "searchModelVersions",
			Handler:    _ModelRegistryService_SearchModelVersions_Handler,
		},
		{
			MethodName: "getModelVersionDownloadUri",
			Handler:    _ModelRegistryService_GetModelVersionDownloadUri_Handler,
		},
		{
			MethodName: "setRegisteredModelTag",
			Handler:    _ModelRegistryService_SetRegisteredModelTag_Handler,
		},
		{
			MethodName: "setModelVersionTag",
			Handler:    _ModelRegistryService_SetModelVersionTag_Handler,
		},
		{
			MethodName: "deleteRegisteredModelTag",
			Handler:    _ModelRegistryService_DeleteRegisteredModelTag_Handler,
		},
		{
			MethodName: "deleteModelVersionTag",
			Handler:    _ModelRegistryService_DeleteModelVersionTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model_registry.proto",
}
