package module

import (
	"go.uber.org/fx"

	activityHTTP "todo/service/activity/delivery/http"
	activityRepo "todo/service/activity/repository/mysql"
	activityUsecase "todo/service/activity/usecase"
)

var Module = fx.Options(
	fx.Provide(
		activityHTTP.New,
		activityUsecase.New,
		activityRepo.New,
	),
)
