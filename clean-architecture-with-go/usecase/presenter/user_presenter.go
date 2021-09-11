package presenter

import "clean-architecture-with-go/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
