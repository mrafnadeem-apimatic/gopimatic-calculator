package tests

import (
    "apimaticcalculator/models"
    "github.com/apimatic/go-core-runtime/testHelper"
    "testing"
)

func TestSimpleCalculatorControllerMultiply(t *testing.T) {
    operation := nil
    x := 4
    y := 5
    apiResponse, err := SimpleCalculatorController.GetCalculate(operation, x, y)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expected := `20`
    testHelper.NativeBodyMatcher(t, expected, apiResponse.Data)
}
