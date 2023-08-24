Initialize and configure the entire app in this package. This package would import all required configs from `internal/app` directory
and would be use at the `main.go` file. Some example of what we could do in this package is
initialized and injecting dependencies such as database connection, logger, etc.

Note that the name file should be matched with the name of app executable, but since this repo is only a boilerplate
so the naming the file to cmd.go should be ok.