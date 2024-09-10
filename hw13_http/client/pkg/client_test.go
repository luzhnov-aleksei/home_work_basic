package client

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	server "github.com/luzhnov-aleksei/home_work_basic/hw13_http/server/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetData(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(server.Handle))
	defer srv.Close()
	client := NewClient(srv.URL, "")
	resp, err := client.GetData()
	expected := "200 OK Request method: GET; RequestURI: /"
	assert.Nil(t, err)
	assert.Equal(t, expected, resp)
}

func TestPostData(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		server.Handle(w, req)
	}))

	defer svr.Close()
	client := NewClient(svr.URL, "")
	msg := "200 OK 978-5-04-154507-9; 978-5-17-109413-3; 978-5-04-116639-7; "
	newMessage := `[{"id":"978-5-04-154507-9","title":"Моби Дик","author":"Герман Мелвилл","year":1851,
	"size":576,"rate":9},{"id":"978-5-17-109413-3","title":"Приключения Тома Сойера","author":"Марк Твен",
	"year":1876,"size":90},{"id":"978-5-04-116639-7","title":"Приключения Гекльберри Финна",
	"author":"Марк Твен","year":1884,"size":95,"rate":9.3}]`

	newMessage = strings.Join(strings.Split(newMessage, "\n\t"), "")
	resp, err := client.PostData(newMessage)
	assert.Nil(t, err)
	assert.Equal(t, msg, resp)
}

func TestPostDataError(t *testing.T) {
	expected := "200 OK Request method: GET; RequestURI: /"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		server.Handle(w, req)
	}))

	defer svr.Close()
	client := NewClient(svr.URL, "")
	resp, err := client.GetData()
	assert.Nil(t, err)
	assert.Equal(t, expected, resp)
	msg := "400 Bad Request invalid character 'o' in literal null (expecting 'u')"
	newMessage := "not json"
	resp, err = client.PostData(newMessage)
	assert.Nil(t, err)
	assert.Equal(t, msg, resp)
}
