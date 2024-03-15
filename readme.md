# temporal-dsl

Simple project to experiment with temporal DSL in go. 

This project was bootstrapped using the DSL sample - see [here](https://github.com/temporalio/samples-go/tree/main/dsl)

## Running the Project

Starting the server:

`temporal server start-dev --db-filename t-dsl.db --ui-port 8080`

Start the worker:

`go run worker/main.go`

Start the workflow:

`go run starter/main.go`