package transactionlog

import (
	"testing"

	"github.com/mariusfa/golf/httpclient"
)

func TestTransLoggerInHttpClient(t *testing.T) {
	TransLog := NewTransLogger("test")
	_ = httpclient.NewHttpClient(TransLog)
}
