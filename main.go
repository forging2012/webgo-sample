package main

import (
	"database/sql"
	"time"

	"github.com/bnkamalesh/webgo"
)

//DBStore has all the database handlers and their respective configurations
type DBStore struct {
	MdbCfg *MgoConfig
	Mdb    *MgoStore
	MySQL  *sql.DB
}

//InitMySQL returns a MySQL handler after opening a new connection
func InitMySQL(host, port, username, password, dbname string) (*sql.DB, error) {
	db, err := sql.Open(
		"mysql",
		username+":"+password+"@tcp("+host+":"+port+")/"+dbname,
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	webgo.Log.Println(db.Stats())

	return db, nil
}

func main() {

	// Load configuration from file
	var cfg webgo.Config
	cfg.Load("config.json")

	// To enable HTTPS server
	// cfg.HTTPSOnly = true
	// cfg.CertFile = "/path/to/cert.pem"
	// cfg.KeyFile = "/path/to/.ssl/key.pem"
	// cfg.HTTPSPort = "443"

	// Loading HTML templates
	var t webgo.Templates
	t.Load(map[string]string{
		"Error": cfg.TemplatesBasePath + "/404.html",
	})

	// Initializing context for the app
	var g webgo.Globals

	g.Init(&cfg, t.Tpls)

	//getting mongodb handler
	mdbCfg := MgoConfig{
		Host:     "127.0.0.1",
		Name:     "dbname",
		Username: "user",
		Password: "password",
		Port:     "27017",
	}

	mdb, err := InitMgo(mdbCfg)
	if err != nil {
		webgo.Log.Println("Error getting MongoDB handler: ", err)
	}

	myql, err := InitMySQL("127.0.0.1", "3306", "user", "password", "dbname")
	if err != nil {
		webgo.Log.Println("Error getting MySQL handler: ", err)
	}

	// Adding custom MySQL handler to globals
	var dbs = DBStore{
		MdbCfg: &mdbCfg,
		Mdb:    mdb,
		MySQL:  myql,
	}

	g.App["dbstore"] = &dbs

	// Initializing router with all the required routes
	router := webgo.InitRouter(getRoutes(&g))
	// router.HideAccessLog = true
	router.NotFound = NotFound(g)

	webgo.Start(
		&cfg,
		router,
		time.Second*15,
		time.Second*15,
	)
	// ====
}
