// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/ashwinyue/maltx. The professional
// version of this repository is https://github.com/onexstack/onex.

//go:build wireinject
// +build wireinject

package apiserver

import (
	"github.com/ashwinyue/maltx/pkg/authz"
	"github.com/google/wire"

	"github.com/ashwinyue/maltx/internal/apiserver/biz"
	"github.com/ashwinyue/maltx/internal/apiserver/cache"
	"github.com/ashwinyue/maltx/internal/apiserver/pkg/validation"
	"github.com/ashwinyue/maltx/internal/apiserver/store"
	ginmw "github.com/ashwinyue/maltx/internal/pkg/middleware/gin"
	"github.com/ashwinyue/maltx/internal/pkg/server"
)

func InitializeWebServer(*Config) (server.Server, error) {
	wire.Build(
		wire.NewSet(NewWebServer, wire.FieldsOf(new(*Config), "ServerMode")),
		wire.Struct(new(ServerConfig), "*"), // * 表示注入全部字段
		wire.NewSet(store.ProviderSet, cache.ProviderSet, biz.ProviderSet),
		ProvideDB,    // 提供数据库实例
		ProvideRedis, // 提供Redis实例
		validation.ProviderSet,
		wire.NewSet(
			wire.Struct(new(UserRetriever), "*"),
			wire.Bind(new(ginmw.UserRetriever), new(*UserRetriever)),
		),
		authz.ProviderSet,
	)
	return nil, nil
}
