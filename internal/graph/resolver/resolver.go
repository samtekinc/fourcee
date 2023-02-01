package resolver

import (
	"context"
	"unicode"

	"github.com/99designs/gqlgen/graphql"
	"github.com/sheacloud/tfom/internal/api"
	"github.com/sheacloud/tfom/internal/config"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	apiClient api.APIClientInterface
	config    *config.Config
}

func NewResolver(apiClient api.APIClientInterface, config *config.Config) *Resolver {
	return &Resolver{apiClient: apiClient, config: config}
}

func GetPreloads(ctx context.Context) []string {
	return GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
}

func GetNestedPreloads(ctx *graphql.OperationContext, fields []graphql.CollectedField, prefix string) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "." + capitalizeFirstLetter(name)
	}
	return capitalizeFirstLetter(name)
}

func capitalizeFirstLetter(s string) string {
	for i, v := range s {
		return string(unicode.ToUpper(v)) + s[i+1:]
	}
	return s
}
