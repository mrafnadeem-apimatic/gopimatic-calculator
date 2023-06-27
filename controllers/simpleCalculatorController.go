package controllers

import (
    "apimaticcalculator/models"
    "fmt"
    "github.com/apimatic/go-core-runtime/https"
)

type SimpleCalculatorController struct {
    baseController
}

func NewSimpleCalculatorController(baseController baseController) *SimpleCalculatorController {
    simpleCalculatorController := SimpleCalculatorController{baseController: baseController}
    return &simpleCalculatorController
}

// Calculates the expression using the specified operation.
func (s *SimpleCalculatorController) GetCalculate(
    operation models.OperationTypeEnum,
    x float64,
    y float64) (
    https.ApiResponse[float64],
    error) {
    req := s.prepareRequest("GET", fmt.Sprintf("/%s", operation))
    req.Authenticate(true)
    req.QueryParam("x", x)
    req.QueryParam("y", y)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[float64]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[float64]{ Response: resp}, err
    }
    
    var result float64
    result, err = utilities.DecodeResults[float64](decoder)
    if err != nil {
        return https.ApiResponse[float64]{ Response: resp}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
