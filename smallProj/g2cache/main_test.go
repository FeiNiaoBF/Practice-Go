package g2cache

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var dbs = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *Group {
	return NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[Slowdbs] search key", key)
			if v, ok := dbs[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, g2cache *Group) {
	peers := NewHTTPPool(addr)
	peers.Set(addrs...)
	g2cache.RegisterPeers(peers)
	log.Println("g2cachecache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, g2cache *Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := g2cache.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())
		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func setupMockServers(cacheServers map[int]string) (map[int]*httptest.Server, []string) {
	serverHandlers := make(map[int]*httptest.Server)
	var cacheServerURLs []string
	for port, _ := range cacheServers {
		mux := http.NewServeMux()
		mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("mock response"))
		})
		server := httptest.NewServer(mux)
		serverHandlers[port] = server
		cacheServers[port] = server.URL
		cacheServerURLs = append(cacheServerURLs, server.URL)
	}
	return serverHandlers, cacheServerURLs
}

func TestMain(m *testing.M) {
	// Mock database
	dbs = map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}

	// Setup mock cache servers
	cacheServers := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	// Create mock cache server handlers
	serverHandlers, cacheServerURLs := setupMockServers(cacheServers)

	// Setup the cache group
	g2cache := createGroup()

	// Register mock peers
	peers := NewHTTPPool(cacheServers[8001])
	peers.Set(cacheServerURLs...)
	g2cache.RegisterPeers(peers)

	// Start the mock API server
	apiAddr := "http://localhost:9999"
	go startAPIServer(apiAddr, g2cache)

	// Allow time for servers to initialize
	time.Sleep(time.Second)

	// Run tests
	code := m.Run()

	// Tear down mock servers
	for _, server := range serverHandlers {
		server.Close()
	}

	// Exit with test result
	os.Exit(code)
}
