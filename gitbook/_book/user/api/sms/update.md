1.2 修改

###### post

​	http://www.user.com/role/update

|    name     | type   | length | comment  |
| :---------: | ------ | ------ | -------- |
|     id      | long   |        | 角色Id   |
|    name     | string |        | 名称     |
| description | string |        | 描述     |
|  parent_id  | long   |        | 父角色Id |

###### post json
```json
{"id":1,"name":"admin","description":"123","parent_id":1}
```

##### post xml

```xml
<role>
     <id>1</id>
	<name>admin</name>
	<description>admin</description>
	<parent_id>1</parent_id>
</role>
```