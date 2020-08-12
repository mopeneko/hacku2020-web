package controller

import "github.com/mopeneko/hacku2020-web/api/repository"

var heartRateRepository *repository.HeartRateRepository

func init() {
	heartRateRepository = repository.NewHeartRateRepository()
}

type PostHeartRateRequest struct {
	Count uint `json:"count"`
}

type PostHeartRateResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type GetHeartRateRequest struct {
}

type GetHeartRateResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
