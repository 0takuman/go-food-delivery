package biz

import (
	"context"
	"errors"
	notemodel "food-delivery/modules/note/model"
)

type CreateNoteStore interface {
	CreateNote(context context.Context, data *notemodel.Notes) error
}

type createNoteBiz struct {
	store CreateNoteStore
}

func NewCreateNoteBiz(store CreateNoteStore) *createNoteBiz {
	return &createNoteBiz{store: store}
}

func (biz *createNoteBiz) CreateNote(context context.Context, data *notemodel.Notes) error {
	if data.Title == "" {
		return errors.New("title is empty")
	}

	if err := biz.store.CreateNote(context, data); err != nil {
		return err
	}
	return nil
}
