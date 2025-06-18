package service

import (
	"fmt"
)

type HTTPClientInterface interface {
	GetGoogle() (string, error)
}

type Service struct {
	Cl HTTPClientInterface // Embeds the client
}

func (s *Service) CheckGoogle() {
	result, err := s.Cl.GetGoogle() // Directly call embedded client method
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Google Response:", result)
}
