// generated version: 0.3.5

package controller

import (
	"context"

	"github.com/labstack/echo/v4"
)

type PostHeartRateController struct {
}

func NewPostHeartRateController() *PostHeartRateController {
	p := &PostHeartRateController{}
	return p
}

// PostHeartRate
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Param Count body integer WIP:${isRequire} WIP:${description}
// @Success 200 {object} PostHeartRateResponse
// @Failure 400 {object} WIP
// @Router /heart_rate [POST]
func (p *PostHeartRateController) PostHeartRate(
	ctx context.Context, c echo.Context, req *PostHeartRateRequest,
) (res *PostHeartRateResponse, err error) {
	panic("require implements.")
}
