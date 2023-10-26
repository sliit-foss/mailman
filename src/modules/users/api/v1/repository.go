package v1

import (
	"mailman/src/modules/users/api/v1/models"
	"mailman/src/utils"
)

var repository = utils.NewRepository[models.User]("users")

func Repository() utils.Repository[models.User] {
	return repository
}
