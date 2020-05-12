package rpc

import (
	"context"

	grpcServer "github.com/iixlabs/virtual-lsobus/rpc/grpc/server"

	"github.com/iixlabs/virtual-lsobus/common/event"
	"github.com/iixlabs/virtual-lsobus/config"
	"github.com/iixlabs/virtual-lsobus/log"
	chainctx "github.com/iixlabs/virtual-lsobus/services/context"
	"go.uber.org/zap"
)

type RPC struct {
	config  *config.Config
	ctx     context.Context
	cancel  context.CancelFunc
	eb      event.EventBus
	cfgFile string
	logger  *zap.SugaredLogger
	cc      *chainctx.ServiceContext
	grpc    *grpcServer.GRPCServer
}

func NewRPC(cfgFile string) (*RPC, error) {
	cc := chainctx.NewServiceContext(cfgFile)
	cfg, _ := cc.Config()
	ctx, cancel := context.WithCancel(context.Background())
	r := RPC{
		eb:      cc.EventBus(),
		config:  cfg,
		cfgFile: cfgFile,
		ctx:     ctx,
		cancel:  cancel,
		logger:  log.NewLogger("rpc"),
		cc:      cc,
	}
	if cfg.RPC.Enable {
		r.grpc = grpcServer.NewGRPCServer()
	}
	return &r, nil
}

func (r *RPC) StopRPC() {
	r.cancel()
	if r.config.RPC.Enable {
		r.grpc.Stop()
	}
}

func (r *RPC) StartRPC() error {
	if r.config.RPC.Enable {
		err := r.grpc.Start(r.config)
		if err != nil {
			return err
		}
	}
	return nil
}
