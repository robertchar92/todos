package usecase

import (
	"todo/service/activity"
)

type ActivityUsecase struct {
	activityRepo activity.Repository
}

func New(
	activityRepo activity.Repository,
) activity.Usecase {
	return &ActivityUsecase{
		activityRepo: activityRepo,
	}
}
