package tracelog

import (
	"bytes"
	"context"
	"github.com/mariusfa/golf/request"
	"os"
	"strings"
	"testing"
)

func TestTraceLog_Context(t *testing.T) {
	ctx := setupContext()
	output := captureStdout(func() {
		SetAppName("testapp")
		Info(ctx, "test payload")
	})

	expected := "\"log_type\":\"TRACE\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", expected, output)
	}
	expected = "\"app_name\":\"testapp\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", expected, output)
	}
	expected = "\"payload\":\"test payload\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", expected, output)
	}
	expected = "\"request_id\":\"123\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", expected, output)
	}
	expected = "\"username\":\"test_username\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", expected, output)
	}
}

func TestTraceLog_MissingContext(t *testing.T) {
	output := captureStdout(func() {
		SetAppName("testapp")
		Info(context.Background(), "test payload")
	})

	expected := "\"request_id\":\"\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", expected, output)
	}
	expected = "\"username\":\"\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", expected, output)
	}
}

func setupContext() context.Context {
	ctx := context.Background()
	sessionCtx := request.NewSessionCtx("test_user_id", "test_name", "test_email", "test_username")
	reqIdCtx := request.NewRequestIdCtx("123")
	ctx = context.WithValue(ctx, request.SessionCtxKey, sessionCtx)
	ctx = context.WithValue(ctx, request.RequestIdCtxKey, reqIdCtx)

	return ctx
}

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)
	r.Close()

	return buf.String()
}
