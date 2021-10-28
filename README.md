# Mock API Server

## Introduction
This app allows you to add urls and serve dummy json responses. It contains two handlers, the `DummyHandler` allows you to map a route to a json schema which will be served as the response and `DummyErrorHandler` allows you to add an endpoint that will return a specific error which is useful for mocking error states.

## Running The Server Locally
Complete the following steps to run this app:

* Ensure golang is installed on your machine
* Build the app using `go build -o bin/mock-api-server`
* Run the app using `./bin/mock-api-server`

The dummy api will now be available on `localhost:8080`

## Adding Routes
Routes can be added to the `router.go` file. Example routes have been added which you can copy. You need only change the url and path to the schema and optionally specify a response http status code. If no status is passed the endpoint will return a status of `200`:

```golang
mux.Handle(
    "/api/test-url",
    &h.DummyHandler{
        FilePath: "./schemas/you_json_response.json",
        Status: 201,
    },
)
```

## Rapid Heroku Deployment
It's possible to have a live dummy server up and running within minutes with heroku by doing the following (it's free!)
* Clone this repo and then remove the `.git` folder `rm -rf .git`
* Create an [Heroku](https://signup.heroku.com/login) account if you don't already have one
* Install the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli)
* Sign into you heroku account and create a new app
* On the deploy tab select "Heroku Git" and follow the commands given.
