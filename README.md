# Webgo Sample

A simple and bare minimum web app built using webgo framework which showcases most of the APIs and abilities of the framework. It shows you how to use a custom Database handler/wrapper which you wrote, adding custom errors for your App etc. The source is simple enough for anyone to understand :)

Both MySQL and MongoDB handlers are being used, which would help in understanding how to add DB handlers and similarly anything required for the to the app context.

## How to run?

`go run *.go`


## Contents of MySQL table `users`

```
+-----+------+------+---------+
| _id | name | age  | company |
+-----+------+------+---------+
|   1 | John |   99 | Github  |
+-----+------+------+---------+ 
```

## Contents of MongoDB collection `users`

```
{
	_id: <mongo ID>,
	name: "John",
	age: 99,
	company: "Github"
}
```