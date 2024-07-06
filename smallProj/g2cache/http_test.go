package g2cache

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHttp(t *testing.T) {
	var db = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}

	// 创建一个mock缓存
	_ = NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	var getTests = []struct {
		name       string
		path       string
		key        any
		expectedOk bool
	}{
		{"http get is true if Tom", "/_g2cache/scores/Tom", "Tom", true},
		{"http get is false if Bob not expected", "/_g2cache/scores/Bob", "Bob", false},
		{"http get is false if not `false` group", "/_g2cache/false/Tom", "Tom", false},
		// {"http get is false if Error URL", "/_false/scores/Tom", "Tom", false},
	}
	// 1. 启动http服务
	pool := NewHTTPPool("localhost:9999")

	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodGet, "http://localhost:9999"+tt.path, nil)
			response := httptest.NewRecorder()

			pool.ServeHTTP(response, request)

			if tt.expectedOk {
				// 检查响应状态码
				if status := response.Code; status != http.StatusOK {
					t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
				}

				// 检查响应内容
				expected := db[tt.key.(string)]
				actual, err := io.ReadAll(response.Body)
				if err != nil {
					t.Fatalf("failed to read response body: %v", err)
				}
				if string(actual) != expected {
					t.Errorf("handler returned unexpected body: got %v want %v", string(actual), expected)
				}

				// 检查 Content-Type 头
				if contentType := response.Header().Get("Content-Type"); contentType != "application/octet-stream" {
					t.Errorf("handler returned wrong content type: got %v want %v", contentType, "application/octet-stream")
				}
			} else {
				// 检查响应状态码
				// Error is StatusInternalServerError
				if status := response.Code; status != http.StatusInternalServerError && status != http.StatusNotFound && status != http.StatusBadRequest {
					t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
				}
			}

		})
	}

}

// Test HttpGetter

func TestHttpGetter_Get(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/_g2cache/testgroup/foo" {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		w.Write([]byte("bar"))
	}))

	defer server.Close()
	t.Logf("server.URL: %s", server.URL+server.Config.Addr)

	// Create a new httpGetter instance
	getter := &httpGetter{baseURL: server.URL + "/_g2cache/"}

	// Get the value
	value, err := getter.Get("testgroup", "foo")
	if err != nil {
		t.Fatal(err)
	}

	// Check the value
	expected := "bar"
	if string(value) != expected {
		t.Errorf("getter returned unexpected value: got %v want %v", string(value), expected)
	}
}
