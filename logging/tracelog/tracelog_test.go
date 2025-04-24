package tracelog

import (
	"bytes"
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/mariusfa/golf/request"
)

func TestTraceLog_Info(t *testing.T) {
	ctx := setupContext()
	output := captureStdout(func() {
		SetAppName("testapp")
		Info(ctx, "test payload")
	})

	expected := "\"log_level\":\"INFO\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"log_type\":\"TRACE\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"app_name\":\"testapp\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"payload\":\"test payload\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"request_id\":\"123\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"user_id\":\"test_username\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
}

func TestTraceLog_Error(t *testing.T) {
	ctx := setupContext()
	output := captureStdout(func() {
		SetAppName("testapp")
		Error(ctx, "test error payload")
	})

	expected := "\"log_level\":\"ERROR\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"log_type\":\"TRACE\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"app_name\":\"testapp\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"payload\":\"test error payload\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"request_id\":\"123\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"user_id\":\"test_username\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
}

func TestTraceLog_Errorf(t *testing.T) {
	ctx := setupContext()
	output := captureStdout(func() {
		SetAppName("testapp")
		Errorf(ctx, "test error payload", errors.New("error happened"))
	})

	expected := "\"log_level\":\"ERROR\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"payload\":\"test error payload: error happened\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
}

func TestTraceLog_MissingContext(t *testing.T) {
	output := captureStdout(func() {
		SetAppName("testapp")
		Info(context.Background(), "test payload")
	})

	expected := "\"request_id\":\"\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
	}
	expected = "\"user_id\":\"\""
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but was not found. Output: %q", expected, output)
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
