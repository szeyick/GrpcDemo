package clients

import "go.uber.org/fx"

var In = fx.Provide(
	NewLogger,
)