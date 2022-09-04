package http

import (
	"testing"
	"time"
)

func TestSendRequest(t *testing.T) {
	go Serve()
	time.Sleep(1 * time.Second)
	res, err := Get("http://127.0.0.1:8080/hello", map[string]string{"Content-Length": "0"}, map[string]string{"name": "alice"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))
	res, err = Post("http://127.0.0.1:8080/user", map[string]string{"Content-Type": "application/json", "Locale": "zh_CN"}, []byte(`{"username":"bob","email":"bob@gmail.com","gender":"male"}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))
	res, err = Post("http://127.0.0.1:8080/user", map[string]string{"Content-Type": "application/json", "Locale": "zh_CN"}, []byte(`{"email":"bob@gmail.com","gender":"male"}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))
}

func TestSendGethRequest(t *testing.T) {
	res, err := Post("https://bsc-dataseed1.binance.org/", map[string]string{"Content-Type": "application/json"}, []byte(`{"jsonrpc":"2.0","method":"net_listening","params":[],"id":67}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))
}
