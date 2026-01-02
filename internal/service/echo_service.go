package service

import "web-based/internal/domain"

func Echo(message string) domain.EchoResponse {
	return domain.EchoResponse{
		Echo: message,
	}
}
