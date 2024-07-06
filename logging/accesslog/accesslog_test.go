package accesslog

import (
	"testing"

	"github.com/mariusfa/golf/middleware"
)

func TestAccesslogInMiddleware(t *testing.T) {
	acceslog := GetLogger()
	middleware.AccessLogMiddleware(nil, acceslog)
}
