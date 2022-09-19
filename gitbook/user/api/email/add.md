1.1 添加

###### post

​	http://www.user.com/role/add

|    name     | type   | length | comment  |
| :---------: | ------ | ------ | -------- |
|    name     | string |        | 名称     |
| description | string |        | 描述     |
|  parent_id  | long   |        | 父角色Id |

###### post json
```json
{"name":"admin","description":"123","parent_id":1}
```

##### post xml

```xml
<role>
	<name>admin</name>
	<description>admin</description>
	<parent_id>1</parent_id>
</role>
```

##### response json
```json
{
    
}
```
##### response xml
```xml
<response>
</response>
```