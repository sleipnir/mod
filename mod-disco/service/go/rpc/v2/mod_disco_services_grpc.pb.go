// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v2

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SurveyServiceClient is the client API for SurveyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SurveyServiceClient interface {
	// Projects
	NewSurveyProject(ctx context.Context, in *NewSurveyProjectRequest, opts ...grpc.CallOption) (*SurveyProject, error)
	GetSurveyProject(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SurveyProject, error)
	ListSurveyProject(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	UpdateSurveyProject(ctx context.Context, in *UpdateSurveyProjectRequest, opts ...grpc.CallOption) (*SurveyProject, error)
	DeleteSurveyProject(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Users
	NewSurveyUser(ctx context.Context, in *NewSurveyUserRequest, opts ...grpc.CallOption) (*SurveyUser, error)
	GetSurveyUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SurveyUser, error)
	ListSurveyUser(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	UpdateSurveyUser(ctx context.Context, in *UpdateSurveyUserRequest, opts ...grpc.CallOption) (*SurveyUser, error)
	DeleteSurveyUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// DiscoProjects
	NewDiscoProject(ctx context.Context, in *NewDiscoProjectRequest, opts ...grpc.CallOption) (*DiscoProject, error)
	GetDiscoProject(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*DiscoProject, error)
	ListDiscoProject(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	UpdateDiscoProject(ctx context.Context, in *UpdateSurveyProjectRequest, opts ...grpc.CallOption) (*DiscoProject, error)
	DeleteDiscoProject(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type surveyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSurveyServiceClient(cc grpc.ClientConnInterface) SurveyServiceClient {
	return &surveyServiceClient{cc}
}

var surveyServiceNewSurveyProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "NewSurveyProject",
}

func (c *surveyServiceClient) NewSurveyProject(ctx context.Context, in *NewSurveyProjectRequest, opts ...grpc.CallOption) (*SurveyProject, error) {
	out := new(SurveyProject)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/NewSurveyProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceGetSurveyProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "GetSurveyProject",
}

func (c *surveyServiceClient) GetSurveyProject(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SurveyProject, error) {
	out := new(SurveyProject)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/GetSurveyProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceListSurveyProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "ListSurveyProject",
}

func (c *surveyServiceClient) ListSurveyProject(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/ListSurveyProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceUpdateSurveyProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "UpdateSurveyProject",
}

func (c *surveyServiceClient) UpdateSurveyProject(ctx context.Context, in *UpdateSurveyProjectRequest, opts ...grpc.CallOption) (*SurveyProject, error) {
	out := new(SurveyProject)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/UpdateSurveyProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceDeleteSurveyProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "DeleteSurveyProject",
}

func (c *surveyServiceClient) DeleteSurveyProject(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/DeleteSurveyProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceNewSurveyUserStreamDesc = &grpc.StreamDesc{
	StreamName: "NewSurveyUser",
}

func (c *surveyServiceClient) NewSurveyUser(ctx context.Context, in *NewSurveyUserRequest, opts ...grpc.CallOption) (*SurveyUser, error) {
	out := new(SurveyUser)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/NewSurveyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceGetSurveyUserStreamDesc = &grpc.StreamDesc{
	StreamName: "GetSurveyUser",
}

func (c *surveyServiceClient) GetSurveyUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SurveyUser, error) {
	out := new(SurveyUser)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/GetSurveyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceListSurveyUserStreamDesc = &grpc.StreamDesc{
	StreamName: "ListSurveyUser",
}

func (c *surveyServiceClient) ListSurveyUser(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/ListSurveyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceUpdateSurveyUserStreamDesc = &grpc.StreamDesc{
	StreamName: "UpdateSurveyUser",
}

func (c *surveyServiceClient) UpdateSurveyUser(ctx context.Context, in *UpdateSurveyUserRequest, opts ...grpc.CallOption) (*SurveyUser, error) {
	out := new(SurveyUser)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/UpdateSurveyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceDeleteSurveyUserStreamDesc = &grpc.StreamDesc{
	StreamName: "DeleteSurveyUser",
}

func (c *surveyServiceClient) DeleteSurveyUser(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/DeleteSurveyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceNewDiscoProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "NewDiscoProject",
}

func (c *surveyServiceClient) NewDiscoProject(ctx context.Context, in *NewDiscoProjectRequest, opts ...grpc.CallOption) (*DiscoProject, error) {
	out := new(DiscoProject)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/NewDiscoProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceGetDiscoProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "GetDiscoProject",
}

func (c *surveyServiceClient) GetDiscoProject(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*DiscoProject, error) {
	out := new(DiscoProject)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/GetDiscoProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceListDiscoProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "ListDiscoProject",
}

func (c *surveyServiceClient) ListDiscoProject(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/ListDiscoProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceUpdateDiscoProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "UpdateDiscoProject",
}

func (c *surveyServiceClient) UpdateDiscoProject(ctx context.Context, in *UpdateSurveyProjectRequest, opts ...grpc.CallOption) (*DiscoProject, error) {
	out := new(DiscoProject)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/UpdateDiscoProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var surveyServiceDeleteDiscoProjectStreamDesc = &grpc.StreamDesc{
	StreamName: "DeleteDiscoProject",
}

func (c *surveyServiceClient) DeleteDiscoProject(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v2.mod_disco.services.SurveyService/DeleteDiscoProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SurveyServiceService is the service API for SurveyService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterSurveyServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type SurveyServiceService struct {
	// Projects
	NewSurveyProject    func(context.Context, *NewSurveyProjectRequest) (*SurveyProject, error)
	GetSurveyProject    func(context.Context, *IdRequest) (*SurveyProject, error)
	ListSurveyProject   func(context.Context, *ListRequest) (*ListResponse, error)
	UpdateSurveyProject func(context.Context, *UpdateSurveyProjectRequest) (*SurveyProject, error)
	DeleteSurveyProject func(context.Context, *IdRequest) (*empty.Empty, error)
	// Users
	NewSurveyUser    func(context.Context, *NewSurveyUserRequest) (*SurveyUser, error)
	GetSurveyUser    func(context.Context, *IdRequest) (*SurveyUser, error)
	ListSurveyUser   func(context.Context, *ListRequest) (*ListResponse, error)
	UpdateSurveyUser func(context.Context, *UpdateSurveyUserRequest) (*SurveyUser, error)
	DeleteSurveyUser func(context.Context, *IdRequest) (*empty.Empty, error)
	// DiscoProjects
	NewDiscoProject    func(context.Context, *NewDiscoProjectRequest) (*DiscoProject, error)
	GetDiscoProject    func(context.Context, *IdRequest) (*DiscoProject, error)
	ListDiscoProject   func(context.Context, *ListRequest) (*ListResponse, error)
	UpdateDiscoProject func(context.Context, *UpdateSurveyProjectRequest) (*DiscoProject, error)
	DeleteDiscoProject func(context.Context, *IdRequest) (*empty.Empty, error)
}

func (s *SurveyServiceService) newSurveyProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.NewSurveyProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method NewSurveyProject not implemented")
	}
	in := new(NewSurveyProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.NewSurveyProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/NewSurveyProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.NewSurveyProject(ctx, req.(*NewSurveyProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) getSurveyProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.GetSurveyProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method GetSurveyProject not implemented")
	}
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetSurveyProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/GetSurveyProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSurveyProject(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) listSurveyProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.ListSurveyProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method ListSurveyProject not implemented")
	}
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ListSurveyProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/ListSurveyProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListSurveyProject(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) updateSurveyProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.UpdateSurveyProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method UpdateSurveyProject not implemented")
	}
	in := new(UpdateSurveyProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.UpdateSurveyProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/UpdateSurveyProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateSurveyProject(ctx, req.(*UpdateSurveyProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) deleteSurveyProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.DeleteSurveyProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method DeleteSurveyProject not implemented")
	}
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DeleteSurveyProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/DeleteSurveyProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteSurveyProject(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) newSurveyUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.NewSurveyUser == nil {
		return nil, status.Errorf(codes.Unimplemented, "method NewSurveyUser not implemented")
	}
	in := new(NewSurveyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.NewSurveyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/NewSurveyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.NewSurveyUser(ctx, req.(*NewSurveyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) getSurveyUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.GetSurveyUser == nil {
		return nil, status.Errorf(codes.Unimplemented, "method GetSurveyUser not implemented")
	}
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetSurveyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/GetSurveyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetSurveyUser(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) listSurveyUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.ListSurveyUser == nil {
		return nil, status.Errorf(codes.Unimplemented, "method ListSurveyUser not implemented")
	}
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ListSurveyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/ListSurveyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListSurveyUser(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) updateSurveyUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.UpdateSurveyUser == nil {
		return nil, status.Errorf(codes.Unimplemented, "method UpdateSurveyUser not implemented")
	}
	in := new(UpdateSurveyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.UpdateSurveyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/UpdateSurveyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateSurveyUser(ctx, req.(*UpdateSurveyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) deleteSurveyUser(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.DeleteSurveyUser == nil {
		return nil, status.Errorf(codes.Unimplemented, "method DeleteSurveyUser not implemented")
	}
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DeleteSurveyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/DeleteSurveyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteSurveyUser(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) newDiscoProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.NewDiscoProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method NewDiscoProject not implemented")
	}
	in := new(NewDiscoProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.NewDiscoProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/NewDiscoProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.NewDiscoProject(ctx, req.(*NewDiscoProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) getDiscoProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.GetDiscoProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method GetDiscoProject not implemented")
	}
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetDiscoProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/GetDiscoProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetDiscoProject(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) listDiscoProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.ListDiscoProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method ListDiscoProject not implemented")
	}
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.ListDiscoProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/ListDiscoProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.ListDiscoProject(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) updateDiscoProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.UpdateDiscoProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method UpdateDiscoProject not implemented")
	}
	in := new(UpdateSurveyProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.UpdateDiscoProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/UpdateDiscoProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.UpdateDiscoProject(ctx, req.(*UpdateSurveyProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *SurveyServiceService) deleteDiscoProject(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	if s.DeleteDiscoProject == nil {
		return nil, status.Errorf(codes.Unimplemented, "method DeleteDiscoProject not implemented")
	}
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.DeleteDiscoProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/v2.mod_disco.services.SurveyService/DeleteDiscoProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.DeleteDiscoProject(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegisterSurveyServiceService registers a service implementation with a gRPC server.
func RegisterSurveyServiceService(s grpc.ServiceRegistrar, srv *SurveyServiceService) {
	sd := grpc.ServiceDesc{
		ServiceName: "v2.mod_disco.services.SurveyService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "NewSurveyProject",
				Handler:    srv.newSurveyProject,
			},
			{
				MethodName: "GetSurveyProject",
				Handler:    srv.getSurveyProject,
			},
			{
				MethodName: "ListSurveyProject",
				Handler:    srv.listSurveyProject,
			},
			{
				MethodName: "UpdateSurveyProject",
				Handler:    srv.updateSurveyProject,
			},
			{
				MethodName: "DeleteSurveyProject",
				Handler:    srv.deleteSurveyProject,
			},
			{
				MethodName: "NewSurveyUser",
				Handler:    srv.newSurveyUser,
			},
			{
				MethodName: "GetSurveyUser",
				Handler:    srv.getSurveyUser,
			},
			{
				MethodName: "ListSurveyUser",
				Handler:    srv.listSurveyUser,
			},
			{
				MethodName: "UpdateSurveyUser",
				Handler:    srv.updateSurveyUser,
			},
			{
				MethodName: "DeleteSurveyUser",
				Handler:    srv.deleteSurveyUser,
			},
			{
				MethodName: "NewDiscoProject",
				Handler:    srv.newDiscoProject,
			},
			{
				MethodName: "GetDiscoProject",
				Handler:    srv.getDiscoProject,
			},
			{
				MethodName: "ListDiscoProject",
				Handler:    srv.listDiscoProject,
			},
			{
				MethodName: "UpdateDiscoProject",
				Handler:    srv.updateDiscoProject,
			},
			{
				MethodName: "DeleteDiscoProject",
				Handler:    srv.deleteDiscoProject,
			},
		},
		Streams:  []grpc.StreamDesc{},
		Metadata: "mod_disco_services.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewSurveyServiceService creates a new SurveyServiceService containing the
// implemented methods of the SurveyService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewSurveyServiceService(s interface{}) *SurveyServiceService {
	ns := &SurveyServiceService{}
	if h, ok := s.(interface {
		NewSurveyProject(context.Context, *NewSurveyProjectRequest) (*SurveyProject, error)
	}); ok {
		ns.NewSurveyProject = h.NewSurveyProject
	}
	if h, ok := s.(interface {
		GetSurveyProject(context.Context, *IdRequest) (*SurveyProject, error)
	}); ok {
		ns.GetSurveyProject = h.GetSurveyProject
	}
	if h, ok := s.(interface {
		ListSurveyProject(context.Context, *ListRequest) (*ListResponse, error)
	}); ok {
		ns.ListSurveyProject = h.ListSurveyProject
	}
	if h, ok := s.(interface {
		UpdateSurveyProject(context.Context, *UpdateSurveyProjectRequest) (*SurveyProject, error)
	}); ok {
		ns.UpdateSurveyProject = h.UpdateSurveyProject
	}
	if h, ok := s.(interface {
		DeleteSurveyProject(context.Context, *IdRequest) (*empty.Empty, error)
	}); ok {
		ns.DeleteSurveyProject = h.DeleteSurveyProject
	}
	if h, ok := s.(interface {
		NewSurveyUser(context.Context, *NewSurveyUserRequest) (*SurveyUser, error)
	}); ok {
		ns.NewSurveyUser = h.NewSurveyUser
	}
	if h, ok := s.(interface {
		GetSurveyUser(context.Context, *IdRequest) (*SurveyUser, error)
	}); ok {
		ns.GetSurveyUser = h.GetSurveyUser
	}
	if h, ok := s.(interface {
		ListSurveyUser(context.Context, *ListRequest) (*ListResponse, error)
	}); ok {
		ns.ListSurveyUser = h.ListSurveyUser
	}
	if h, ok := s.(interface {
		UpdateSurveyUser(context.Context, *UpdateSurveyUserRequest) (*SurveyUser, error)
	}); ok {
		ns.UpdateSurveyUser = h.UpdateSurveyUser
	}
	if h, ok := s.(interface {
		DeleteSurveyUser(context.Context, *IdRequest) (*empty.Empty, error)
	}); ok {
		ns.DeleteSurveyUser = h.DeleteSurveyUser
	}
	if h, ok := s.(interface {
		NewDiscoProject(context.Context, *NewDiscoProjectRequest) (*DiscoProject, error)
	}); ok {
		ns.NewDiscoProject = h.NewDiscoProject
	}
	if h, ok := s.(interface {
		GetDiscoProject(context.Context, *IdRequest) (*DiscoProject, error)
	}); ok {
		ns.GetDiscoProject = h.GetDiscoProject
	}
	if h, ok := s.(interface {
		ListDiscoProject(context.Context, *ListRequest) (*ListResponse, error)
	}); ok {
		ns.ListDiscoProject = h.ListDiscoProject
	}
	if h, ok := s.(interface {
		UpdateDiscoProject(context.Context, *UpdateSurveyProjectRequest) (*DiscoProject, error)
	}); ok {
		ns.UpdateDiscoProject = h.UpdateDiscoProject
	}
	if h, ok := s.(interface {
		DeleteDiscoProject(context.Context, *IdRequest) (*empty.Empty, error)
	}); ok {
		ns.DeleteDiscoProject = h.DeleteDiscoProject
	}
	return ns
}

// UnstableSurveyServiceService is the service API for SurveyService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableSurveyServiceService interface {
	// Projects
	NewSurveyProject(context.Context, *NewSurveyProjectRequest) (*SurveyProject, error)
	GetSurveyProject(context.Context, *IdRequest) (*SurveyProject, error)
	ListSurveyProject(context.Context, *ListRequest) (*ListResponse, error)
	UpdateSurveyProject(context.Context, *UpdateSurveyProjectRequest) (*SurveyProject, error)
	DeleteSurveyProject(context.Context, *IdRequest) (*empty.Empty, error)
	// Users
	NewSurveyUser(context.Context, *NewSurveyUserRequest) (*SurveyUser, error)
	GetSurveyUser(context.Context, *IdRequest) (*SurveyUser, error)
	ListSurveyUser(context.Context, *ListRequest) (*ListResponse, error)
	UpdateSurveyUser(context.Context, *UpdateSurveyUserRequest) (*SurveyUser, error)
	DeleteSurveyUser(context.Context, *IdRequest) (*empty.Empty, error)
	// DiscoProjects
	NewDiscoProject(context.Context, *NewDiscoProjectRequest) (*DiscoProject, error)
	GetDiscoProject(context.Context, *IdRequest) (*DiscoProject, error)
	ListDiscoProject(context.Context, *ListRequest) (*ListResponse, error)
	UpdateDiscoProject(context.Context, *UpdateSurveyProjectRequest) (*DiscoProject, error)
	DeleteDiscoProject(context.Context, *IdRequest) (*empty.Empty, error)
}
