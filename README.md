
# Getting Started with APIMATIC Calculator

## Introduction

Simple calculator API hosted on APIMATIC

### Requirements

The SDK requires **Go version 1.18 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## Installation

The following section explains how to use the apimaticcalculator library in a new project.

### 1. Add SDK as a Dependency to the Application

- Add the following lines to your application's `go.mod` file:

```go
replace apimaticcalculator => ".\\APIMATIC Calculator" // local path to the SDK

require apimaticcalculator v0.0.0
```

- Resolve the dependencies in the updated `go.mod` file, using the `go get` command.

## Test the SDK

`Go Testing` is used as the testing framework. To run the unit tests of the SDK, navigate to the `tests` directory of the SDK and run the following command in the terminal:

```bash
$ go test
```

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `environment` | Environment | The API environment. <br> **Default: `Environment.PRODUCTION`** |
| `httpConfiguration` | `https.HttpConfiguration` | Configurable http client options. |

The API client can be initialized as follows:

```go
config := apimaticcalculator.ConfigurationFactory(
    apimaticcalculator.WithEnvironment(apimaticcalculator.PRODUCTION),
)
config.LoadFromEnvironment()

client := apimaticcalculator.NewClient(config)
```

## List of APIs

* [Simple Calculator](doc/controllers/simple-calculator.md)

## Classes Documentation

* [HttpConfiguration](doc/http-configuration.md)
* [RetryConfiguration](doc/retry-configuration.md)

