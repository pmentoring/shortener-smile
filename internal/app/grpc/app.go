package grpc

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"shortener-smile/internal/grpc/shortener"
	"shortener-smile/internal/shortener/service"
)

type AppServer struct {
	gRPCServer *grpc.Server
	port       int
	log        *slog.Logger
}

func NewApp(
	log *slog.Logger,
	port int,
	shortenerService *service.ShortenLinkService,
	unshortenerService *service.UrlShortenerService,
) *AppServer {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived, logging.PayloadSent,
		),
	}

	srv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(InterceptorLogger(log), loggingOpts...),
	))

	shortener.RegisterShortenerServer(srv, shortenerService, unshortenerService)

	return &AppServer{
		gRPCServer: srv,
		port:       port,
		log:        log,
	}
}

func (a *AppServer) MustStart() {
	err := a.Start()

	if err != nil {
		a.log.Error(err.Error())
	}
}

func (a *AppServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		slog.Error(err.Error())
		return err
	}

	if err := a.gRPCServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (a *AppServer) GracefulShutdown() {
	a.gRPCServer.GracefulStop()
}

func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
