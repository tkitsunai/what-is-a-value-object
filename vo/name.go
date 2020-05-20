package vo

type Name struct {
	value string
}

func NewName(value string) Name {
	return Name{
		value: value,
	}
}
