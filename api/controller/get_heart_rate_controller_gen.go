// generated version: 0.3.5

package controller

import (
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
)

type GetHeartRateController struct {
}

func NewGetHeartRateController() *GetHeartRateController {
	g := &GetHeartRateController{}
	return g
}

// GetHeartRate
// @Summary Returns latest heart rate
// @Description Returns heart rate which is set to *repository.HeartRateRepository
// @Accept json
// @Produce json
// @Success 200 {object} GetHeartRateResponse
// @Failure 400 {object} WIP
// @Router /heart_rate [GET]
func (g *GetHeartRateController) GetHeartRate(
	ctx context.Context, c echo.Context, req *GetHeartRateRequest,
) (res *GetHeartRateResponse, err error) {
	res = new(GetHeartRateResponse)

	count := heartRateRepository.Get()

	res.Status = true
	res.Message = strconv.Itoa(int(count))
	return
}
