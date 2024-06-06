# Health pakke

Dette er en pakke for helsesjekk av en app. Helsesjekken er en HTTP GET request til `/health` som returnerer en HTTP status code 200.

Per default så spørr AWS på dette endepunktet for å sjekke om en app er oppe og kjører. Hvis appen ikke svarer så vil ikke AWS sende trafikk til appen.

## Ta i bruk

```go
import (
	"net/http"

	"github.com/sparebank1utvikling/gal/health"
)
func main() {
	router := http.NewServeMux()
	health.RegisterRoute(router)
	http.ListenAndServe(":8080", router)
}
```