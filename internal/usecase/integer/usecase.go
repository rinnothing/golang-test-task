package integer

import (
	"context"
	"slices"

	"github.com/rinnothing/golang-test-task/internal/model"
	"github.com/rinnothing/golang-test-task/pkg/transaction"
)

const ReviewersNum = 2

type Usecase interface {
	AddInteger(context.Context, model.Integer) ([]model.Integer, error)
}

type IntegerRepo interface {
	AddInteger(context.Context, model.Integer) error
	ListIntegers(context.Context) ([]model.Integer, error)
}

var _ Usecase = &impl{}

type impl struct {
	repo IntegerRepo
	tr   transaction.Transactor
}

func New(repo IntegerRepo, tr transaction.Transactor) *impl {
	return &impl{repo: repo, tr: tr}
}

func (u *impl) AddInteger(ctx context.Context, num model.Integer) ([]model.Integer, error) {
	var res []model.Integer
	err := u.tr.DoAtomically(ctx, func(ctx context.Context) error {
		err := u.repo.AddInteger(ctx, num)
		if err != nil {
			return err
		}

		res, err = u.repo.ListIntegers(ctx)
		if err != nil {
			return err
		}

		slices.Sort(res)
		return nil
	})

	return res, err
}
