package interceptors

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TokenAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	auth, err := extractHeader(ctx, "authorization")

	if err != nil {
		return ctx, err
	}

	const prefix = "Bearer "

	if !strings.HasPrefix(auth, prefix) {
		return ctx, status.Error(codes.Unauthenticated, `missing "Bearer " prefix in "Authorization" header`)
	}

	if strings.TrimPrefix(auth, prefix) != "abcdef123" {
		return ctx, status.Error(codes.Unauthenticated, "invalid token")
	}

	return handler(ctx, req)
}

func extractHeader(ctx context.Context, header string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return "", status.Error(codes.Unauthenticated, "no header in request")
	}

	authHeader, ok := md[header]

	if !ok {
		return "", status.Error(codes.Unauthenticated, "more than 1 header in request")
	}

	return authHeader[0], nil
}
