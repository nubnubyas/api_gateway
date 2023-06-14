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

service calculatorService {
    calculatorResp calculate(1: calculatorReq request) (api.post="/calculator/get") 
}
