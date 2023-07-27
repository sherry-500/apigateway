namespace go demo

//--------------------request & response--------------
struct ApiReq {
    1: required string data(api.body = 'data'),
}

struct ApiResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

//----------------------service-------------------
service ApigatewayService {
    ApiResp Gateway(1: ApiReq req)(api.post = '/apigateway/:svcName/:methodName')
}