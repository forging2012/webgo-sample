// This file contains a wrapper to use over Mgo driver.
// Only basic operations are implemented as of now.
package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MgoConfig Database configuration
type MgoConfig struct {
	Name          string `json:"name"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	AuthSource    string `json:"authSource"`
	MgoDialString string `json:"mgoDialString"`
}

//MgoStore DB Sessions are maintained inside a struct for better caching of the data stores
//Developed based on the answer:
//http://stackoverflow.com/questions/26574594/best-practice-to-maintain-a-mgo-session
type MgoStore struct {
	DbName  string
	Session *mgo.Session
}

// Clone the master session and return
func (ds *MgoStore) getSession() *mgo.Session {
	return ds.Session.Copy()
}

//GetSessionCollection gets the appropriate MongoDB collection
func (ds *MgoStore) GetSessionCollection(dbName, collection string) (*mgo.Session, *mgo.Collection) {
	s := ds.getSession()
	c := s.DB(dbName).C(collection)

	return s, c
}

//Get does a MongoDB Get
func (ds *MgoStore) Get(dbName, collection string, conditions interface{}, resultStruct interface{}) ([]bson.M, error) {

	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	if resultStruct != nil {
		err := c.Find(conditions).All(resultStruct)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil, nil
			}
			return nil, err
		}
		return nil, nil
	}

	var data []bson.M
	err := c.Find(conditions).All(&data)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

//GetAll does a MongoDB GetAll
func (ds *MgoStore) GetAll(dbName, collection string, resultStruct interface{}) ([]bson.M, error) {

	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	if resultStruct != nil {
		err := c.Find(nil).All(resultStruct)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil, nil
			}
			return nil, err
		}
		return nil, nil
	}

	var data []bson.M
	err := c.Find(nil).All(&data)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

//GetOne does a MongoDB GetOne
func (ds *MgoStore) GetOne(dbName, collection string, conditions interface{}, resultStruct interface{}) (bson.M, error) {

	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	if resultStruct != nil {
		err := c.Find(conditions).One(resultStruct)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil, nil
			}
			return nil, err
		}
		return nil, nil
	}

	var data bson.M
	err := c.Find(conditions).One(&data)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

//Save does a MongoDB Save
func (ds *MgoStore) Save(dbName, collection string, data interface{}) error {
	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	err := c.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

//Update does a MongoDB Update - multiple records
func (ds *MgoStore) Update(dbName, collection string, condition, updateData interface{}) error {
	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	err := c.Update(condition, updateData)
	if err != nil {
		return err
	}

	return nil
}

//UpdateId does MongoDB update - single record, by MongoID
func (ds *MgoStore) UpdateId(dbName, collection string, _id, data interface{}) error {
	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	err := c.UpdateId(_id, data)
	if err != nil {
		return err
	}

	return nil
}

//RemoveId does a MongoDB Remove, single record, by MongoID
func (ds *MgoStore) RemoveId(dbName, collection string, id interface{}) error {
	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	err := c.RemoveId(id)
	if err != nil {
		return err
	}

	return nil
}

//Remove does a MongoDB Remove
func (ds *MgoStore) Remove(dbName, collection string, condition interface{}) error {
	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	err := c.Remove(condition)
	if err != nil {
		return err
	}

	return nil
}

// RemoveAll does MongoDB RemoveAll
func (ds *MgoStore) RemoveAll(dbName, collection string) (*mgo.ChangeInfo, error) {
	s, c := ds.GetSessionCollection(dbName, collection)
	defer s.Close()

	info, err := c.RemoveAll(nil)
	if err != nil {
		return nil, err
	}

	return info, nil
}

//newDataStore creates a new data store
func newDataStore(user, pass, host, port, name, authSource, mgoDialString string) (*MgoStore, error) {
	dialString := "mongodb://"

	if len(mgoDialString) > 0 {
		dialString = mgoDialString
	} else {
		if len(user) > 0 && len(pass) > 0 {
			dialString += (user + ":" + pass)
		}

		if len(host) > 0 {
			dialString += ("@" + host)
			if len(port) > 0 {
				dialString += (":" + port)
			}
		}

		if len(name) > 0 {
			dialString += ("/" + name)
		}

		if len(authSource) > 0 {
			dialString += "?authSource=" + authSource
		}
	}

	session, err := mgo.Dial(dialString)
	if err != nil {
		return nil, err
	}
	session.SetSafe(&mgo.Safe{})

	return &MgoStore{DbName: name, Session: session}, nil
}

//InitMgo initializes MongoDB
func InitMgo(dbc MgoConfig) (*MgoStore, error) {
	dStore, err := newDataStore(
		dbc.Username,
		dbc.Password,
		dbc.Host,
		dbc.Port,
		dbc.Name,
		dbc.AuthSource,
		dbc.MgoDialString,
	)

	if err != nil {
		return nil, err
	}

	return dStore, nil
}
