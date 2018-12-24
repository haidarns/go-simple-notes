package main

import (
	"github.com/haidarns/iot-platform-go/app/api/note"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mgo "gopkg.in/mgo.v2"
)

const (
	MongoHost = "127.0.0.1"
	MongoPort = "27017"
	MongoUser = "root"
	MongoPass = "example"
	MongoDB   = "note-api"
)

func main() {
	// MongoDB configuration
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{MongoHost},
		Username: MongoUser,
		Password: MongoPass,
	})
	if err != nil {
		panic(err)
	}
	db := session.DB(MongoDB)
	defer session.Close()

	// Echo configuration
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Note api configuration
	noteRepo := note.NewMongoRepository(db)
	noteUcase := note.NewNoteUsecase(noteRepo)
	note.NewHTTPHandler(e, noteUcase)

	e.Logger.Fatal(e.Start(":1234"))
	// e.Start(":1234")
}
