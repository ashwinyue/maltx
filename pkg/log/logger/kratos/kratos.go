// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/onex.
//

package kratos

import (
	krtlog "github.com/go-kratos/kratos/v2/log"

	"github.com/ashwinyue/maltx/pkg/log"
)

func NewLogger(id, name, version string) krtlog.Logger {
	return krtlog.With(log.Default(),
		"ts", krtlog.DefaultTimestamp,
		"caller", krtlog.DefaultCaller,
		"service.id", info.ID,
		"service.name", info.Name,
		"service.version", info.Version,
	)
}
