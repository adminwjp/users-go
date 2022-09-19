# Summary

* [Introduction](README.md)

* 用户
    * [表结构](user/user_table.md)
    * Api
        * [1.0 用户](user/api/user.md)
            * [1.1 登录](user/api/user/login.md)
            * [1.2 注册](user/api/user/register.md)
            * [1.3 用户是否存在](user/api/user/exists.md)
            * [1.4 修改手机号](user/api/user/update_phone.md)
            * [1.5 修改邮箱](user/api/user/update_email.md)
            * [1.6修改密码](user/api/user/update_pwd.md)
            * [1.7 查询用户](user/api/user/list.md)
        * [1.0 管理员](user/api/admin.md)
            * [1.1 登录](user/api/admin/login.md)
            * [1.2 注册](user/api/admin/register.md)
            * [1.3 用户是否存在](user/api/admin/exists.md)
            * [1.4 修改手机号](user/api/admin/update_phone.md)
            * [1.5 修改邮箱](user/api/admin/update_email.md)
            * [1.6 修改密码](user/api/admin/update_pwd.md)
            * [1.7 查询用户](user/api/admin/list.md)
        * [1.0 角色](user/api/role.md)
            * [1.1 添加](user/api/role/add.md)
            * [1.2 修改](user/api/role/update.md)
            * [1.3 删除](user/api/role/delete.md)
            * [1.4 查询](user/api/role/list.md)
        * [1.0 邮箱](user/api/email.md)
            * [1.1 添加](user/api/email/add.md)
            * [1.2 修改](user/api/email/update.md)
            * [1.3 删除](user/api/email/delete.md)
            * [1.4 查询](user/api/email/list.md)
        * [1.0 短信](user/api/sms.md)
            * [1.1 添加](user/api/sms/add.md)
            * [1.2 修改](user/api/sms/update.md)
            * [1.2 删除](user/api/rolsmse/delete.md)
            * [1.3 查询](user/api/sms/list.md)
        * [1.0 支付](user/api/pay_secret.md)
            * [1.1 添加](user/api/pay_secret/add.md)
            * [1.2 修改](user/api/role/update.md)
            * [1.3 删除](user/api/pay_secret/delete.md)
            * [1.4 查询](user/api/pay_secret/list.md)       
    * Rpc(Grpc、thrift)
        * [1.0 用户](user/rpc/user.md)
            * [1.1 登录](user/rpc/user/login.md)
            * [1.2 注册](user/rpc/user/register.md)
            * [1.3 用户是否存在](user/rpc/user/exists.md)
            * [1.4 修改手机号](user/rpc/user/update_phone.md)
            * [1.5 修改邮箱](user/rpc/user/update_email.md)
            * [1.6 修改密码](user/rpc/user/update_pwd.md)
        * [1.0 管理员](user/rpc/admin.md)
            * [1.1 登录](user/rpc/admin/login.md)
            * [1.2 注册](user/rpc/admin/register.md)
            * [1.3 用户是否存在](user/rpc/admin/exists.md)
            * [1.4 修改手机号](user/rpc/admin/update_phone.md)
            * [1.5 修改邮箱](user/rpc/admin/update_email.md)
            * [1.6 修改密码](user/rpc/admin/update_pwd.md)
        * [1.0 数据层](user/data/data.md)
            * [1.1 gorm(jinzhu)数据层](user/data/db/gorm_jinzhu.md)
            * [1.1 orm(bee)数据层](user/data/db/bee_orm.md)
            * [1.2 es数据层](user/data/es/es.md)
            * [1.3 mong数据层](user/data/mong/mong.md)
        * [1.0 缓存](user/cache/cache.md)
            * [1.1 本地缓存](user/cache/local.md)
            * [1.2 Redis 缓存](user/cache/redis.md)
        * [1.0 web](user/web/web.md)
            * [1.1 http](user/web/http.md)
            * [1.2 bee](user/web/bee.md)
            * [1.3 gin](user/web/gin.md)   