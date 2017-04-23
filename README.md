# Webgo Sample

A simple and bare minimum web app built using webgo framework which showcases most of the APIs and abilities of the framework. It shows you how to use a custom Database handler/wrapper which you wrote, adding custom errors for your App etc. The source is simple enough for anyone to understand :)

Both MySQL and MongoDB handlers are being used, which would help in understanding how to add DB handlers and similarly anything required for the to the app context.

### How to run?

1. Start a MySQL & MongoDB server

2. MySQL - create a table `users`, with the following columns `_id,name,age,company` (as shown in the sample data below).

3. Setup a MongoDB server, create a collection `users` and insert a sample document (as shown in the sample data below)

4. Update the configurations accordingly in `main.go`.

5. Start the server by running this command `$ go run *.go` in the terminal.

If all good, you'll see the following message on the terminal.

`Starting HTTP server, listening on ':8000'`

You can try the following links to test

1. `http://localhost:8000` - Hello world response with a CORS middleware, and a post response middleware.

2. `http://localhost:8000/auth` - Sample authentication middleware

3. `http://localhost:8000/mgodb/John` - Result fetched from MongoDB

4. `http://localhost:8000/mysql/John` - Result fetched from MySQL

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