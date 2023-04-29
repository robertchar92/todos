package usecase

import (
	"todo/lib/database_transaction"
	"todo/service/activity"
	"todo/service/todo"
)

type TodoUsecase struct {
	todoRepo           todo.Repository
	activityRepo       activity.Repository
	transactionManager database_transaction.Client
}

func New(
	todoRepo todo.Repository,
	activityRepo activity.Repository,
	transactionManager database_transaction.Client,
) todo.Usecase {
	return &TodoUsecase{
		todoRepo:           todoRepo,
		activityRepo:       activityRepo,
		transactionManager: transactionManager,
	}
}
