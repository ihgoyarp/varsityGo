package service

import "web-based/internal/domain"

func CheckHealth() string {
	return domain.HealthMessage()
}
