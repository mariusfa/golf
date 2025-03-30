package applog

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAppLog_Info(t *testing.T) {
	out := captureStdout(func() {
		SetAppName("testapp")
		Info("test payload")
	})

	if !strings.Contains(out, "\"log_type\":\"APP\"") {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", "\"log_type\": \"APP\"", out)
	}
	if !strings.Contains(out, "\"app_name\":\"testapp\"") {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", "\"app_name\": \"testapp\"", out)
	}
	if !strings.Contains(out, "\"payload\":\"test payload\"") {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", "\"message\": \"test payload\"", out)
	}
	if !strings.Contains(out, "\"log_level\":\"INFO\"") {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", "\"log_level\": \"info\"", out)
	}
}

func TestAppLog_Error(t *testing.T) {
	out := captureStdout(func() {
		SetAppName("testapp")
		Error("test payload")
	})

	if !strings.Contains(out, "\"log_level\":\"ERROR\"") {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", "\"log_level\": \"ERROR\"", out)
	}
}

func TestAppLog_Infof(t *testing.T) {
	out := captureStdout(func() {
		SetAppName("testapp")
		Infof("test payload %s", "formatted")
	})

	if !strings.Contains(out, "\"payload\":\"test payload formatted\"") {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", "\"message\": \"test payload formatted\"", out)
	}
}

func TestAppLog_Errorf(t *testing.T) {
	out := captureStdout(func() {
		SetAppName("testapp")
		Errorf("test payload", fmt.Errorf("error occured"))
	})

	if !strings.Contains(out, "\"payload\":\"test payload\"") {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", "\"message\": \"test payload formatted\"", out)
	}
	if !strings.Contains(out, "\"error\":\"error occured\"") {
		t.Errorf("Expected output to contain %q, but was not found. Ouput: %q", "\"error\": \"error occured\"", out)
	}
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
