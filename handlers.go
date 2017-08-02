package main

import (
	"net/http"

	"github.com/bnkamalesh/webgo"
	"gopkg.in/mgo.v2/bson"

	_ "github.com/go-sql-driver/mysql"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	webgo.R200(w, "Hello world")
	// webgo.SendResponse(w, "Hello world", 200)
}

//NotFound is the 404 handler
func NotFound(g webgo.Globals) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		webgo.Render404(w, g.Templates["Error"])
	})
}

//MongoDB is the handler which fetches data from MOngoDB
func MongoDB(w http.ResponseWriter, r *http.Request) {
	// Access the Globals from context
	wctx := webgo.Context(r)
	dbs := wctx.Route.G.App["dbstore"].(*DBStore)

	// Query string and request body can be accessed directly from `r *http.Request`

	results, err := dbs.Mdb.Get(dbs.MdbCfg.Name, "users", bson.M{"name": wctx.Params["name"]}, nil)
	if err != nil {
		webgo.Log.Println(err)
		webgo.R500(w, "Sorry, an unknown error occurred")
		return
	}

	webgo.R200(w, results)
}

//MySQL is the handler which uses the MySQL DB handler
func MySQL(w http.ResponseWriter, r *http.Request) {
	wctx := webgo.Context(r)
	dbs := wctx.Route.G.App["dbstore"].(*DBStore)

	// Query string and request body can be accessed directly from `r *http.Request`

	rows, err := dbs.MySQL.Query("SELECT * FROM users WHERE name=?", wctx.Params["name"])
	if err != nil {
		webgo.Log.Println(err)
		webgo.R500(w, err.Error())
		return
	}
	defer rows.Close()

	var name, company string
	var _id, age int

	var item = make(map[string]interface{})

	for rows.Next() {
		rows.Scan(&_id, &name, &age, &company)
		item["_id"] = _id
		item["name"] = name
		item["age"] = age
		item["company"] = company
	}

	webgo.R200(w, []interface{}{item})
}
