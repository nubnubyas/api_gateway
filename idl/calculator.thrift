namespace go calculator

struct calculatorReq {
    1: string name (api.form="name")
    2: i64 num1 (api.form="num1")
    3: i64 num2 (api.form="num2")
    4: string operation (api.form="operation")
}

struct calculatorResp {
    1: string message
}

struct capCalculatorReq {
    // change the name of the field to grades
    //1: list<i64> num1 (api.form="grades")
    1: list<double> num1 (api.form="grades")
}

struct capCalculatorResp {
    1: double message
}

service calculatorService {
    calculatorResp calculate(1: calculatorReq request)  (api.get="/calculatorService/get")
    capCalculatorResp capCalculate(1: capCalculatorReq request) (api.post="/calculatorService/cap") 
}