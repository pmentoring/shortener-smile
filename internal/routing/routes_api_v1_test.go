package routing

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"shortener-smile/database"
	"shortener-smile/internal/common"
	"strings"
	"testing"
)

type ShortenUrlRequest struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func TestShortenUrl(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		return
	}

	router := gin.Default()

	Register(router, db, &common.ApplicationContext{
		InstanceId: "01",
		AppBaseUrl: "http://localhost:8000/",
	})

	w := httptest.NewRecorder()

	request := ShortenUrlRequest{
		URL:   "https://google.com/hui",
		Title: "sukaaa",
	}

	jsonBody, _ := json.Marshal(request)

	req, _ := http.NewRequest("POST", "/shorten", strings.NewReader(string(jsonBody)))
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}
