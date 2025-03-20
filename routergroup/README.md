# Router groups 

En `RouterGroup` kan brukes for å gruppere en samling endepunkter under samme prefix. 
Den kan også brukes til å definere et sett med `Middleware` som skal gjelde for endepunktene
som er registrert på denne gruppen.

## Ta i bruk

Eksmpel på bruk:
```go
    sb1skAdapter := sb1sk.NewAdapter(config.Sb1skTokenPath, config.Sb1skHost)
    router := http.NewServeMux()
    sb1Middlwares := routergroup.SB1Middlewares(sb1skAdapter)

    routerGroup := &routergroup.RouterGroup{
        Prefix: "/api/common/banking", 
        Middlewares: sb1Middlwares, 
        Mux: router
    }
```

Middlewares defineres som en funksjon som tar en `http.Handler` som argument og returnerer en `http.Handler`.:

```go
    logTime := func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            next.ServeHTTP(w, r)
            log.Printf("Request took: %v", time.Since(start))
        })
    }

    router := http.NewServeMux()
    routerGroup := &routergroup.RouterGroup{
        Prefix: "/api/common/banking", 
        Middlewares: []routergroup.Middleware{logTime}, 
        Mux: router
    }
```


Registrering av routes:

```go
    routergroup.HandleFunc("GET", "/watchlists", getAll)
    routergroup.HandleFunc("POST", "/watchlists", post)

    func getAll (w http.ResponseWriter, r *http.Request) {
        // ...
        w.WriteHeader(http.StatusOk)
    }

    func post (w http.ResponseWriter, r *http.Request) {
        //...
        w.WriteHeader(http.StatusCreated)
    }
```
