package bins

import (
	"errors"
	"time"
)

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, isPrivate bool, name string) (*Bin, error) {
	if id == "" {
		return nil, errors.New("INVALID_ID")
	}

	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}

	newBin := &Bin{
		ID:        id,
		Private:   isPrivate,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return newBin, nil
}
