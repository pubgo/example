package version_entry

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/pubgo/lava/clients/orm"
	"github.com/pubgo/lava/config"
	"github.com/pubgo/lava/internal/example/services/protopb/proto/hello"
	"github.com/pubgo/lava/logger"
)

var _ hello.TestApiServer = (*Service)(nil)

type Service struct {
	Db  *orm.Client   `dix:""`
	Cfg config.Config `dix:""`
}

func (t *Service) Init() {

}

func (t *Service) Router(r gin.IRouter) {
}

func (t *Service) VersionTestCustom(ctx context.Context, req *hello.TestReq) (*hello.TestApiOutput, error) {
	panic("implement me")
}

func (t *Service) Version1(ctx context.Context, req *structpb.Value) (*hello.TestApiOutput1, error) {
	panic("implement me")
}

func (t *Service) Version(ctx context.Context, in *hello.TestReq) (out *hello.TestApiOutput, err error) {
	var log = logger.GetLog(ctx)
	log.Sugar().Infof("Received Helloworld.Call request, name: %s", in.Input)

	if t.Db != nil {
		log.Info("dix db ok", zap.Any("err", t.Db.Ping()))
		log.Info("dix config ok", zap.String("cfg", t.Cfg.ConfigPath()))
	}

	out = &hello.TestApiOutput{
		Msg: in.Input,
	}
	out.Reset()
	time.Sleep(time.Millisecond * 10)
	return
}

func (t *Service) VersionTest(ctx context.Context, in *hello.TestReq) (out *hello.TestApiOutput, err error) {
	out = &hello.TestApiOutput{
		Msg: in.Input + "_test",
	}
	return
}
