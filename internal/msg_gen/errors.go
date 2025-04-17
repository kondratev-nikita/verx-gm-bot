package msg_gen

import (
	"errors"
	"fmt"
)

func newErr(text string, err error) error {
	if err == nil {
		return errors.New("msg_gen_svc: " + text)
	}
	return fmt.Errorf("msg_gen_svc: %s: %v", text, err)
}
