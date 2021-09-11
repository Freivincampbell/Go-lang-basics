package registry

import (
	"clean-architecture-with-go/interface/controllers"
	ip "clean-architecture-with-go/interface/presenters"
	ir "clean-architecture-with-go/interface/repository"
	"clean-architecture-with-go/usecase/interactor"
	up "clean-architecture-with-go/usecase/presenter"
	ur "clean-architecture-with-go/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
