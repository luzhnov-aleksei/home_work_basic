package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleGet(t *testing.T) {
	expected := "Request method: GET; RequestURI: http://localhost:10001"
	req := httptest.NewRequest(http.MethodGet, "http://localhost:10001", nil)
	w := httptest.NewRecorder()
	Handle(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, expected, string(body))
}

func TestHandlePost(t *testing.T) {
	msg := "978-5-04-154507-9; 978-5-17-109413-3; 978-5-04-116639-7; "
	data := `[{"id":"978-5-04-154507-9","title":"Моби Дик","author":"Герман Мелвилл","year":1851,
	"size":576,"rate":9},{"id":"978-5-17-109413-3","title":"Приключения Тома Сойера","author":"Марк Твен",
	"year":1876,"size":90},{"id":"978-5-04-116639-7","title":"Приключения Гекльберри Финна",
	"author":"Марк Твен","year":1884,"size":95,"rate":9.3}]`
	data = strings.Join(strings.Split(data, "\n\t"), "")
	body := strings.NewReader(data)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:10001", body)
	w := httptest.NewRecorder()
	Handle(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, msg, string(respBody))
}

func TestHandleIncorrectMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "http://localhost:10001", nil)
	w := httptest.NewRecorder()
	Handle(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 405, resp.StatusCode)
	assert.Equal(t, "", string(respBody))
}

func TestHandlePostError(t *testing.T) {
	msg := "invalid character 'o' in literal null (expecting 'u')"
	data := "not json"
	body := strings.NewReader(data)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:10001", body)
	w := httptest.NewRecorder()
	Handle(w, req)
	resp := w.Result()
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, msg, string(respBody))
}
