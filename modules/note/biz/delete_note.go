package biz

import (
	"context"
	"errors"
	nodemodel "food-delivery/modules/note/model"
)

type DeleteNoteStore interface {
	FindDataWithCondition(context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*nodemodel.Notes, error)
	Delete(context context.Context, id int) error
}

type deleteNoteBiz struct {
	store DeleteNoteStore
}

func NewDeleteNoteBiz(store DeleteNoteStore) *deleteNoteBiz {
	return &deleteNoteBiz{store: store}
}

func (biz *deleteNoteBiz) DeleteNote(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status != 1 {
		return errors.New("note is deleted")
	}

	if err := biz.store.Delete(context, id); err != nil {
		return err
	}
	return nil
}
