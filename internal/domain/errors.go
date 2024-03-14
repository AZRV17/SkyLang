package domain

type ErrInvalidPassword struct{}

func (e *ErrInvalidPassword) Error() string {
	return "invalid password"
}
