package models

import mgo "gopkg.in/mgo.v2"

// GetSession return a mongo session
func GetSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.DialWithInfo(&mgo.DialInfo{
        Addrs:    []string{"localhost"},
        Username: "",
        Password: "",
        Database: "admin",
    })

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return s
}
