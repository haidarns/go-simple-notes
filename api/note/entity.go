package note

import "gopkg.in/mgo.v2/bson"

type (
	// Note type
	Note struct {
		ID    bson.ObjectId `bson:"_id" json:"id"`
		Title string        `bson:"title" json:"title"`
		Body  string        `bson:"body" json:"body"`
	}

	// Notes type [Note]
	Notes []Note
)

// Validate Note model
func (n *Note) Validate() error {
	return nil
}
