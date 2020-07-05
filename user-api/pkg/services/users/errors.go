package users

import "errors"

func invalidDelete() error {
	return errors.New("it wasn't possible to delete the orders from the user, please contact the support")
}