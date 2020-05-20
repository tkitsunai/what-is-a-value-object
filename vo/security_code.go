package vo

import "errors"

type SecurityCode struct {
	value int
}

func NewSecurityCode(value int) (SecurityCode, error) {
	if value < 1300 || value > 9999 {
		return SecurityCode{}, errors.New("証券コードの値範囲外です")
	}

	return SecurityCode{
		value: value,
	}, nil
}
