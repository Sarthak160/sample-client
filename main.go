package main

import (
	"github.com/Sarthak160/go-client/client"
	"github.com/Sarthak160/go-client/service"
)

func main() {
	cl := &client.HTTPClient{}
	svc := &service.Service{Cl: cl}
	svc.CheckGoogle()
}
