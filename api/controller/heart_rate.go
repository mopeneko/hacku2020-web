package controller

type PostHeartRateRequest struct {
	Count uint
}

type PostHeartRateResponse struct {
	Status  bool
	Message string
}

type GetHeartRateRequest struct {
}

type GetHeartRateResponse struct {
	Count uint
}
