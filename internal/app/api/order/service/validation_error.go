package service

type MyError struct{}

func (m *MyError) Error() string {
	return "boom"
}
