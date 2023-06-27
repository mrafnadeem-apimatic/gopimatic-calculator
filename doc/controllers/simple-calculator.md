# Simple Calculator

```go
simpleCalculatorController := client.SimpleCalculatorController
```

## Class Name

`SimpleCalculatorController`


# Get Calculate

Calculates the expression using the specified operation.

```go
GetCalculate(input GetCalculateInput) (
    https.ApiResponse[float64],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `operation` | [`models.OperationTypeEnum`](../../doc/models/operation-type-enum.md) | Template, Required | The operator to apply on the variables |
| `x` | `float64` | Query, Required | The LHS value |
| `y` | `float64` | Query, Required | The RHS value |

## Response Type

`float64`

## Example Usage

```go
operation := models.OperationTypeEnum("MULTIPLY")

x := 222.14

y := 165.14

apiResponse, err := simpleCalculatorController.GetCalculate(operation, x, y)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

