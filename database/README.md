# Database

This is a package for database setup. More specific postgresql.
The package will assume a migrations folder. This is using golang-migrate to perform migrations. However this supports also inserting user and password to do this:
```sql
CREATE USER {{.User}} WITH PASSWORD '{{.Password}}';
```


The package will setup a connection to a postgresql database and perform migrations.


For local development there is a setting to start up a postgresql docker container.


For testing there is a method to be used as testmain which will start up a postgresql docker container and stop it after the tests are done.

### Dependencies
- docker (needs to be installed)
- golang-migrate (go dependency)
- testcontainers (go dependency)

## Usage
Setup in main function or beginning of app. This will setup a docker container and db connection.
```go
// assuming dbConfig is populated
if dbConfig.StartupLocal == "true" {
    containerCleanUp, err := database.SetupContainer(dbConfig)
    if err != nil {
        panic(err)
    }
    defer containerCleanUp()
}

// Sets up db connection and migrations
db, err := dbSetup(dbConfig) 
if err != nil {
    panic(err)
}
defer db.Close()

// Do stuff with db connection
```


For testing use the testmain method
```go
// For instance in a file called main_test.go

// package variable to be used in tests
var dbConfig database.DbConfig

func TestMain(m *testing.M) {
	dbConfig = database.DbConfig{
		User:        "test",
		Password:    "test",
		Name:        "test",
		AppUser:     "app_user",
		AppPassword: "app_password",
	}
	database.SinglePostgresTestMain(m, &dbConfig, "../../../migrations")
}
```

In a test file:
```go
func TestInsertGetTodos(t *testing.T) {
    // using dbConfig variable from testmain
	db, err := database.Setup(&dbConfig)
	if err != nil {
		t.Fatalf("Failed to setup database: %v", err)
	}
	defer db.Close()
	todoRepository := NewTodoRepository(db)

    // Do stuff with todoRepository
}
```
