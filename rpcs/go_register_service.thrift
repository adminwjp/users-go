//thrift-0.16.0.exe -gen java go_register_service.thrift
//thrift-0.16.0.exe -gen go go_register_service.thrift
//thrift-0.16.0.exe -gen netstd go_register_service.thrift
namespace netstd com.user_system.ThriftRpcServices
namespace java com.mall.user.rpc.thrift.gouser
namespace cpp com.user_system.thriftRpcs.services
//https://blog.csdn.net/weixin_30443075/article/details/99351633
//server is not null
namespace go registers
struct REmptyThrift {

}
struct NameThrift {
	1:string name
}
struct ServiceThrift {
  1:string ip
  2:i32 port
  3:string name
}
struct ServiceApiThrift {
  1:string ip
  2:i32 port
  3:string name
  4:i32 status
}

struct ServiceReplyThrift{
  1:i32 status
  2:string msg
}

service RegisterThriftService {
  ServiceReplyThrift Register (1:ServiceThrift input),

  ServiceReplyThrift Callback (),

  ServiceApiThrift Check (1:NameThrift input),

  ServiceApiThrift Get (1:NameThrift input)
}
