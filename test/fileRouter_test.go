package test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-app/router"
)


func TestUploadSingleRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := router.SetupRouter()

    // 開檔
	file, err := os.Open("./asset/img/1.png")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
    // 產生boundary
	writer := multipart.NewWriter(body)

    // 讀檔並寫到body, 填寫form的field key跟一些內容
	part, err := writer.CreateFormFile("file", filepath.Base("./img/test.png"))
	if err != nil {
		t.Error(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error(err)
	}
	_ = writer.Close()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/file/uploadSingle", body)
	req.Header.Add("Content-type", writer.FormDataContentType())

	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}