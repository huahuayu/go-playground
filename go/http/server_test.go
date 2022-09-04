package http

import (
	"os/exec"
	"testing"
)

func TestServer(t *testing.T) {
	Serve()
}

func TestHttpServer(t *testing.T) {
	go Serve()
	result, _ := exec.Command("curl", "http://127.0.0.1:8080/hello?name=alice").Output()
	t.Log(string(result))
	result, _ = exec.Command("curl", "http://127.0.0.1:8080/hello").Output()
	t.Log(string(result))
	result, _ = exec.Command("curl", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"username":"bob","email":"bob@gmail.com","gender":"male"}`, "http://localhost:8080/user").Output()
	t.Log(string(result))
	result, _ = exec.Command("curl", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"email":"bob@gmail.com"}`, "http://localhost:8080/user").Output()
	t.Log(string(result))
}
