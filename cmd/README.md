Initialize and configure the entire app in this package. This package would import all required configs from `internal/app` directory
and would be use at the `main.go` file. Some example of what we could do in this package is
initialize and injecting dependencies such as database connection, logger, etc.