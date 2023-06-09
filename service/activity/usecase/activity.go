package usecase

import (
	"todo/lib/database_transaction"
	"todo/service/activity"
)

type ActivityUsecase struct {
	activityRepo       activity.Repository
	transactionManager database_transaction.Client
}

func New(
	activityRepo activity.Repository,
	transactionManager database_transaction.Client,
) activity.Usecase {
	return &ActivityUsecase{
		activityRepo:       activityRepo,
		transactionManager: transactionManager,
	}
}
