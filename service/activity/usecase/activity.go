package usecase

import "todo/service/activity"

type ActivityUsecase struct{}

func New() activity.Usecase {
	return &ActivityUsecase{}
}
