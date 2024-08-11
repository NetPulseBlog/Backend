package ers

import "fmt"

func ThrowMessage(op string, err error) error {
	return fmt.Errorf("%s: %w", op, err)
}
