package transactionlog

import (
	"testing"

	"github.com/mariusfa/golf/httpclient"
)

func TestTransLoggerInHttpClient(t *testing.T) {
	translog := GetLogger()
	_ = httpclient.NewHttpClient(translog)
}
