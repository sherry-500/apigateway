namespace go demo

//--------------------request & response--------------
struct Resp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct IdlMap {
    1: required string svcName(api.body='svcName'),
    2: required string idl(api.body='idl'),
    3: required string type(api.body='type'),
}

struct ListReq {
    1: required i32 page_id(api.query='page_id'),
    2: required i32 page_size(api.query='page_size')
}

//----------------------service-------------------
service IdlManage {
    IdlMap CreateIdl(1: IdlMap req)(api.post = '/idl-manage')
    Resp DeleteIdl()(api.delete = '/idl-manage/:svcname')
    IdlMap UpdateIdl(1: IdlMap req)(api.patch = '/idl-manage/:svcname')
    IdlMap GetIdlByName()(api.get = '/idl-manage/:svcname')
    list<IdlMap> ListIdl(1: ListReq req)(api.get='/idl-manage')
}