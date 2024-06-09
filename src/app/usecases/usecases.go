package usecases

import (
	userUC "ex_service/src/app/usecases/user"
)

// register all usecase here in order to call it from hanlder
type AllUseCases struct {
	UserUC        userUC.UserUCInterface
}
