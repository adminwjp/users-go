//thrift-0.16.0.exe -gen java go_config.thrift
//thrift-0.16.0.exe -gen go go_config.thrift
//thrift-0.16.0.exe -gen netstd go_config.thrift
// netstd csharp
namespace netstd com.user_system.ThriftRpcServices
namespace java com.mall.user.rpc.thrift.gouser
namespace cpp com.user_system.thriftRpcs.services
//https://blog.csdn.net/weixin_30443075/article/details/99351633
//server is not null
namespace go configs
struct CEmptyRequest {

}


struct DbThrift   {
  1:i32 id
  2:string  default_db
  3:string  db
  4:i64  ip
  5:string  ip_string
  6:i32  port
  7:i64  count
  8:i64  max_count
  9:i64  space
  10:i64  start_id
  11:i64  end_id
   12:string table
}
struct TableThrift {
  1:i32  id
  2:string  default_db
  3:string  db
  4:i64  ip
  5:string  ip_string
  6:i32  port
  7:i64  count
  8:i64  max_count
  9:i64  space
  10:i64  start_id
  11:i64  end_id
  12:string table
  13:string  default_table
  14:string  ids
  15:string  default_ids
}


struct ReplyThrift {
  1:i32 status
  2:string data
}

service ConfigThriftService {
  ReplyThrift SaveDb (1:DbThrift input) ,

  ReplyThrift SaveTab (1:TableThrift input),

  ReplyThrift ListDb (),

  ReplyThrift ListTab ()
}
