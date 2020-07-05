package orders

import "errors"

func invalidUser() error {
	return errors.New("the user does not exist")
}
