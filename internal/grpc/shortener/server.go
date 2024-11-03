package shortener

import (
	"context"
	shortenerv1 "github.com/pmentoring/shortener-protoc/gen/go/shortener"
	"google.golang.org/grpc"
	"shortener-smile/internal/shortener/service"
)

type Shortener interface {
	Shorten(ctx context.Context, req *shortenerv1.UrlShortenRequest) (*shortenerv1.UrlShortenResponse, error)
	Unshorten(ctx context.Context, req *shortenerv1.UrlUnshortenRequest) (*shortenerv1.UrlUnshortenResponse, error)
}

type shortenerServer struct {
	shortenerv1.UnimplementedShortenerServer
	shortener   *service.ShortenLinkService
	unshortener *service.UrlShortenerService
}

func RegisterShortenerServer(server *grpc.Server, shortener *service.ShortenLinkService, unshortener *service.UrlShortenerService) {
	shortenerv1.RegisterShortenerServer(server, &shortenerServer{
		shortener:   shortener,
		unshortener: unshortener,
	})
}

func (s *shortenerServer) Shorten(ctx context.Context, req *shortenerv1.UrlShortenRequest) (*shortenerv1.UrlShortenResponse, error) {
	link, err := s.shortener.CreateShortenLink(req.GetTitle(), req.GetUrl())

	if err != nil {
		return nil, err
	}

	return &shortenerv1.UrlShortenResponse{Url: link.ShortenLink}, nil
}

func (s *shortenerServer) Unshorten(ctx context.Context, req *shortenerv1.UrlUnshortenRequest) (*shortenerv1.UrlUnshortenResponse, error) {
	link, err := s.unshortener.GetLinkByCode(req.GetUrl())

	if err != nil {
		return nil, err
	}

	return &shortenerv1.UrlUnshortenResponse{Url: link.FullLink}, nil
}
