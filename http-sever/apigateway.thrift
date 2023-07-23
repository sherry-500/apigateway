namespace go demo

//--------------------request & response--------------
struct ApiReq {
    1: required string data(api.body = 'data'),
}

struct ApiResp {
    1: required i32 resp(api.body='resp'),
}

struct IdlResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct IdlMap {
    1: required string svcName(api.body='svcName'),
    2: required string path(api.body='path'),
}

struct Service {
    1: required string name(api.body='name'),
}

//----------------------service-------------------
service ApigatewayService {
    ApiResp Gateway(1: ApiReq req)(api.post = '/apigateway/:svcName/:methodName')
    IdlResp CreateIdl(1: IdlMap req)(api.post = '/idl-manage/create')
    IdlResp DeleteIdl(1: Service svc)(api.post = '/idl-manage/delete')
    IdlResp UpdateIdl(1: IdlMap req)(api.post = '/idl-manage/update')
    IdlMap SearchIdl(1: Service svc)(api.post = '/idl-manage/search')
}