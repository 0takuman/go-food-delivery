package biz

import (
	"context"
	"food-delivery/common"
	nodemodel "food-delivery/modules/note/model"
	notemodel "food-delivery/modules/note/model"
)

type ListNoteStore interface {
	ListDataWithCondition(context context.Context,
		filter *notemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]notemodel.Notes, error)
}

type listNoteBiz struct {
	store ListNoteStore
}

func NewListNoteBiz(store ListNoteStore) *listNoteBiz {
	return &listNoteBiz{store: store}
}

func (biz *listNoteBiz) ListNote(context context.Context,
	filter *notemodel.Filter,
	paging *common.Paging,
) ([]nodemodel.Notes, error) {

	data, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, err
	}
	return data, nil
}
