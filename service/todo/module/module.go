package module

import (
	"go.uber.org/fx"

	todoHTTP "todo/service/todo/delivery/http"
	todoRepo "todo/service/todo/repository/mysql"
	todoUsecase "todo/service/todo/usecase"
)

var Module = fx.Options(
	fx.Provide(
		todoHTTP.New,
		todoUsecase.New,
		todoRepo.New,
	),
)
