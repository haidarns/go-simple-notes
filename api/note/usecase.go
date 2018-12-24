package note

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// Usecase interface Note
	Usecase interface {
		Store(n *Note) error
		GetOne(id bson.ObjectId) (*Note, error)
		GetAll() (*Notes, error)
	}

	uCase struct {
		noteRepo Repository
	}
)

// NewNoteUsecase creating Note usecase
func NewNoteUsecase(noteRepo Repository) Usecase {
	return &uCase{
		noteRepo: noteRepo,
	}
}

func (nu *uCase) Store(n *Note) error {
	return nu.noteRepo.Create(n)
}

func (nu *uCase) GetOne(id bson.ObjectId) (*Note, error) {
	return nu.noteRepo.FindOne(id)
}

func (nu *uCase) GetAll() (*Notes, error) {
	return nu.noteRepo.FindAll()
}
