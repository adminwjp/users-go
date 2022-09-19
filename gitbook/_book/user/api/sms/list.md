1.4 查询
###### get

​	http://www.user.com/role/list/1/10

###### post

​	http://www.user.com/role/list/1/10

|    name     | type   | length | comment  |
| :---------: | ------ | ------ | -------- |
|    name     | string |        | 名称     |
|  parent_id  | long   |        | 父角色Id |

###### post json
```json
{"name":"admin","parent_id":1}
```

##### post xml

```xml
<role>
	<name>admin</name>
	<parent_id>1</parent_id>
</role>
```

1.4.1 查询分类
###### get

​	http://www.user.com/role/tree