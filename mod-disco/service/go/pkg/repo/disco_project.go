package repo

import (
	"context"
	"github.com/getcouragenow/mod/mod-disco/service/go/pkg/dao"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	discoRpc "github.com/getcouragenow/mod/mod-disco/service/go/rpc/v2"
	sharedAuth "github.com/getcouragenow/sys-share/sys-account/service/go/pkg/shared"
	sysCoreSvc "github.com/getcouragenow/sys/sys-core/service/go/pkg/coredb"
)

func (md *ModDiscoRepo) NewDiscoProject(ctx context.Context, in *discoRpc.NewDiscoProjectRequest) (*discoRpc.DiscoProject, error) {
	if in == nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot insert disco project: %v", sharedAuth.Error{Reason: sharedAuth.ErrInvalidParameters})
	}
	dp, err := md.store.InsertDiscoProject(in)
	if err != nil {
		return nil, err
	}
	return dp, nil
}

func (md *ModDiscoRepo) GetDiscoProject(ctx context.Context, in *discoRpc.IdRequest) (*discoRpc.DiscoProject, error) {
	if in == nil || in.DiscoProjectId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "cannot get disco project: %v", sharedAuth.Error{Reason: sharedAuth.ErrInvalidParameters})
	}
	dp, err := md.store.GetDiscoProject(map[string]interface{}{
		"project_id": in.DiscoProjectId,
	})
	if err != nil {
		return nil, err
	}
	return dp.ToPkgDiscoProject()
}

func (md *ModDiscoRepo) ListDiscoProject(ctx context.Context, in *discoRpc.ListRequest) (*discoRpc.ListResponse, error) {
	if in == nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot list survey user: %v", sharedAuth.Error{Reason: sharedAuth.ErrInvalidParameters})
	}
	var cursor int64
	orderBy := in.OrderBy
	var err error
	filter := map[string]interface{}{}
	if in.Filters != nil && len(in.Filters) > 0 {
		filter, err = sysCoreSvc.UnmarshalToMap(in.Filters)
		if err != nil {
			return nil, err
		}
	}
	if in.IdRequest != nil {
		if in.IdRequest.SysAccountProjectId != "" {
			filter["sys_account_project_ref_id"] = in.IdRequest.SysAccountProjectId
		}
	}
	cursor, err = md.getCursor(in.GetCurrentPageId())
	if err != nil {
		return nil, err
	}
	if in.GetIsDescending() {
		orderBy += " DESC"
	} else {
		orderBy += " ASC"
	}
	limit := in.PerPageEntries
	if limit == 0 {
		limit = dao.DefaultLimit
	}
	daoDiscoProjects, next, err := md.store.ListDiscoProject(filter, orderBy, limit, cursor)
	if err != nil {
		return nil, err
	}
	var pkgDiscoProjects []*discoRpc.DiscoProject
	for _, su := range daoDiscoProjects {
		surveyUser, err := su.ToPkgDiscoProject()
		if err != nil {
			return nil, err
		}
		pkgDiscoProjects = append(pkgDiscoProjects, surveyUser)
	}
	return &discoRpc.ListResponse{
		DiscoProjects: pkgDiscoProjects,
		NextPageId:    *next,
	}, nil
}

func (md *ModDiscoRepo) UpdateDiscoProject(ctx context.Context, in *discoRpc.UpdateDiscoProjectRequest) (*discoRpc.DiscoProject, error) {
	if in == nil || in.ProjectId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "cannot update disco project: %v", sharedAuth.Error{Reason: sharedAuth.ErrInvalidParameters})
	}
	if err := md.store.UpdateDiscoProject(in); err != nil {
		return nil, err
	}
	daoDiscoProject, err := md.store.GetDiscoProject(map[string]interface{}{"project_id": in.ProjectId})
	if err != nil {
		return nil, err
	}
	return daoDiscoProject.ToPkgDiscoProject()
}

func (md *ModDiscoRepo) DeleteDiscoProject(ctx context.Context, in *discoRpc.IdRequest) (*emptypb.Empty, error) {
	if in == nil || (in.DiscoProjectId == "" && in.SysAccountProjectId == "" && in.SysAccountOrgId == "") {
		return nil, status.Errorf(codes.InvalidArgument, "cannot delete disco project: %v", sharedAuth.Error{Reason: sharedAuth.ErrInvalidParameters})
	}
	err := md.store.DeleteDiscoProject(in.DiscoProjectId, in.SysAccountProjectId, in.SysAccountOrgId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
