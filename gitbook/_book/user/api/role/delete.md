



1.3删除

###### get

​	http://www.user.com/role/remove/1

| name | type | length | comment |
| :--: | ---- | ------ | ------- |
|  id  | long |        | 角色Id  |

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