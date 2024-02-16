package metadata

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetContext(ctx context.Context, key string) string {
	var values []string

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		values = md.Get(key)
	}

	if len(values) > 0 {
		return values[0]
	}

	return ""
}

func SetContext(ctx context.Context, key string, value string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(key, value))
}

func GetSessionContextToken(ctx context.Context) string {
	return GetContext(ctx, "token")
}

func SetSessionContextToken(ctx context.Context, token string) context.Context {
	return SetContext(ctx, "token", token)
}

func GetBotContextToken(ctx context.Context) string {
	return GetContext(ctx, "bot_api_token")
}

func SetBotContextToken(ctx context.Context, token string) context.Context {
	return SetContext(ctx, "bot_api_token", token)
}
