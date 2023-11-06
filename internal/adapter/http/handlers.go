package customHTTP

import (
	"github.com/iki-rumondor/init-golang-service/internal/application"
)

type Handlers struct {
	Service *application.Service
}

func NewTaskHandler(service *application.Service) *Handlers {
	return &Handlers{
		Service: service,
	}
}
