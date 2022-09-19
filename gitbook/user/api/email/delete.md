



1.3删除

###### get

​	http://www.user.com/role/remove/1

| name | type | length | comment |
| :--: | ---- | ------ | ------- |
|  id  | long |        | 角色Id  |

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

1.3.1删除

###### post

http://www.user.com/role/remove/1

| name | type | length | comment |
| :--: | ---- | ------ | ------- |
|  id  | long |        | 角色Id  |

###### post json

```json
{"ids":[1,2,3]}
```

##### post xml

```xml
<user>
    <id>1</id>
    <id>2</id>
    <id>3</id>
</user>
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