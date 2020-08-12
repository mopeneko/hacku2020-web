package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHeartRate(t *testing.T) {
	// POST /heart_rate
	reqJSON := map[string]interface{}{
		"count": 120,
	}
	reqJSONText, _ := json.Marshal(reqJSON)
	req := httptest.NewRequest(http.MethodPost, "/heart_rate", bytes.NewReader(reqJSONText))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"status": true, "message": ""}`, rec.Body.String())

	// GET /heart_rate
	req = httptest.NewRequest(http.MethodGet, "/heart_rate", nil)
	rec = httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"status": true, "message": "120"}`, rec.Body.String())
}
