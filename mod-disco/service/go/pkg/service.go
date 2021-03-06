package pkg

import (
	"fmt"
	service "github.com/getcouragenow/mod/mod-disco/service/go"
	"github.com/getcouragenow/mod/mod-disco/service/go/pkg/repo"
	discoRpc "github.com/getcouragenow/mod/mod-disco/service/go/rpc/v2"
	sharedAccountPkg "github.com/getcouragenow/sys-share/sys-account/service/go/pkg"
	"github.com/getcouragenow/sys-share/sys-account/service/go/pkg/interceptor"
	sharedBus "github.com/getcouragenow/sys-share/sys-core/service/go/pkg/bus"
	"github.com/getcouragenow/sys/sys-core/service/go/pkg/coredb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type ModDiscoService struct {
	proxyService      *discoRpc.SurveyServiceService
	ClientInterceptor *interceptor.ClientSide
}

type ModDiscoServiceConfig struct {
	authProxyClient *sharedAccountPkg.SysAccountProxyServiceClient
	store           *coredb.CoreDB
	Cfg             *service.ModDiscoConfig
	bus             *sharedBus.CoreBus
	logger          *logrus.Entry
}

func NewModDiscoServiceConfig(l *logrus.Entry, db *coredb.CoreDB, filepath string, bus *sharedBus.CoreBus, grpcClientOpts grpc.ClientConnInterface) (*ModDiscoServiceConfig, error) {
	var err error
	if db == nil {
		return nil, fmt.Errorf("error creating mod disco service: database is null")
	}
	modDiscoLogger := l.WithFields(logrus.Fields{
		"mod": "mod-disco",
	})
	discoCfg, err := service.NewConfig(filepath)
	if err != nil {
		return nil, err
	}
	newAuthProxyClient := sharedAccountPkg.NewSysAccountProxyServiceClient(grpcClientOpts)
	mdsc := &ModDiscoServiceConfig{
		store:           db,
		Cfg:             discoCfg,
		bus:             bus,
		logger:          modDiscoLogger,
		authProxyClient: newAuthProxyClient,
	}
	return mdsc, nil
}

func NewModDiscoService(cfg *ModDiscoServiceConfig) (*ModDiscoService, error) {
	cfg.logger.Infoln("Initializing Mod-Disco Service")
	discoRepo, err := repo.NewDiscoRepo(cfg.logger, cfg.store, cfg.Cfg, cfg.bus, cfg.authProxyClient)
	if err != nil {
		return nil, err
	}
	discoService := discoRpc.NewSurveyServiceService(discoRepo)
	return &ModDiscoService{
		proxyService:      discoService,
		ClientInterceptor: discoRepo.ClientInterceptor,
	}, nil
}

func (mds *ModDiscoService) RegisterServices(srv *grpc.Server) {
	discoRpc.RegisterSurveyServiceService(srv, mds.proxyService)
}
