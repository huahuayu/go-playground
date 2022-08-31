package http

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"os/exec"
	"testing"
	"time"
)

func TestHttpServer(t *testing.T) {
	go serve()
	result, _ := exec.Command("curl", "http://127.0.0.1:6000/hello?name=alice").Output()
	t.Log(string(result))
	result, _ = exec.Command("curl", "http://127.0.0.1:6000/hello").Output()
	t.Log(string(result))
	result, _ = exec.Command("curl", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"username":"bob","email":"bob@gmail.com","gender":"male"}`, "http://localhost:6000/user").Output()
	t.Log(string(result))
	result, _ = exec.Command("curl", "-X", "POST", "-H", "Content-Type: application/json", "-d", `{"email":"bob@gmail.com"}`, "http://localhost:6000/user").Output()
	t.Log(string(result))
}

func TestSendRequest(t *testing.T) {
	go serve()
	sendGetRequest(t)
	SendPostRequest(t)
}

func sendGetRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:6000/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("name", "alice")
	req.URL.RawQuery = q.Encode()
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			t.Log("timeout")
			return
		}
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))
}

func SendPostRequest(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var jsonStr = []byte(`{"username":"bob","email":"bob@gmail.com","gender":"male"}`)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://localhost:6000/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			t.Log("timeout")
			return
		}
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(body))
}
