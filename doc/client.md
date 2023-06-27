
# Client Class Documentation

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

## APIMATIC Calculator Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| simpleCalculator | Gets SimpleCalculatorController |

