package note

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoColl = "note"
)

type (
	//Repository repository interface
	Repository interface {
		Create(n *Note) error
		FindOne(id bson.ObjectId) (*Note, error)
		FindAll() (*Notes, error)
	}

	mongoRepository struct {
		db *mgo.Database
	}
)

// NewMongoRepository creating new Note repository
func NewMongoRepository(db *mgo.Database) Repository {
	return &mongoRepository{db}
}

// CreateNote to save note into mongodb
func (m *mongoRepository) Create(n *Note) error {
	coll := m.db.C(mongoColl)

	n.ID = bson.NewObjectId()
	if err := coll.Insert(n); err != nil {
		return err
	}
	return nil
}

func (m *mongoRepository) FindOne(id bson.ObjectId) (*Note, error) {
	// WIP: Not implemented yet
	return &Note{}, nil
}

// GetAllNote to fetch all stored notes
func (m *mongoRepository) FindAll() (*Notes, error) {
	var notes Notes

	coll := m.db.C(mongoColl)

	if err := coll.Find(bson.M{}).All(&notes); err != nil {
		return nil, err
	}
	return &notes, nil
}
