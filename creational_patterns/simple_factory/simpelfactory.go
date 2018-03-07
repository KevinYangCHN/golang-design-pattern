package simplefactory

import "fmt"

type API interface {
	Say(name string) string
}

type hiAPI struct {
}

func (hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type helloAPI struct {
}

func (helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// NewAPI return API instance by type
func NewAPI(t int) API {
	switch t {
	case 1:
		return &hiAPI{}
	case 2:
		return &helloAPI{}
	default:
		return nil
	}
}
