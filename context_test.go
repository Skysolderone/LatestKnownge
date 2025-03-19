package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

// basic
func TestContext(t *testing.T) {
	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()
	t.Log("ctx genarate success")
}

// transport ctx
func TestTransport(t *testing.T) {
	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel()
	go worker(ctxWithCancel)
	time.Sleep(time.Second * 3)
	cancel()
	time.Sleep(time.Second)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("receive cancel")
			return
		default:
			fmt.Println("workdering...")
			time.Sleep(time.Second)
		}
	}
}

// context realize cancel
func TestRealizeCancel(t *testing.T) {
	ctx := context.Background()
	cwc, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-cwc.Done():
			t.Log("receive cancel")
		case <-time.After(time.Second * 2):
			t.Log("workder finish")
		}
	}()
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}

// time out control
func TestContextTimeOutControl(t *testing.T) {
	ctx := context.Background()
	cwt, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	select {
	case <-cwt.Done():
		t.Log("TIME out finish")
	case <-time.After(time.Second * 3):
		t.Log("workering...")
	}
}

// value transfer
func TestContextTransferValue(t *testing.T) {
	cwv := context.WithValue(context.Background(), "key", "i am value")
	value := cwv.Value("key")
	t.Log(value)
}

// context process http time out
func handler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	select {
	case <-time.After(time.Second * 3):
		fmt.Fprintln(w, "hello world")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("处理请求", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
}

func TestContextProcessHttp(t *testing.T) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
