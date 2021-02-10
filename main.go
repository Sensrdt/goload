package main
import (
	"fmt"
	"log"
	"net/http"
)

var (
	serverList = [] *server{
		newServer("http://127.0.0.1:9000"),
		newServer("http://127.0.0.1:9001"),
		newServer("http://127.0.0.1:9002"),
		newServer("http://127.0.0.1:9003"),
	}

	lastIndex = 0
)

func requestForwarding(res http.ResponseWriter, req *http.Request) {
	server, err := checkServerStatus()
	if server == nil {
		http.Error(res, "Couldn't proces request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	server.reverseProxy.ServeHTTP(res, req)
}

func checkServerStatus() (*server, error) {
	for i:=0; i<len(serverList); i++ {
		server := routingToServer()
		if server.health {
			return server, nil
		}
	}

	return nil, fmt.Errorf("All Dead")

}

func routingToServer() *server{
	nextIndex := (lastIndex  + 1) %  len(serverList)
	server := serverList[nextIndex]
	log.Printf("Routing to '%s'", nextIndex)
	lastIndex = nextIndex
	return server
}

func main() {
	http.HandleFunc("/", requestForwarding)
	go startHealthCheck()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
