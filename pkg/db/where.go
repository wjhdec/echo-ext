package db

import (
	"strings"

	"github.com/wjhdec/echo-ext/pkg/pair"
)

func NewWherePairs() *WherePairs {
	return &WherePairs{pairs: make([]pair.Pair[string, []any], 0)}
}

func NewWherePairsWithParams(param string, value ...any) *WherePairs {
	return &WherePairs{pairs: []pair.Pair[string, []any]{
		*pair.New[string, []any](param, value),
	}}
}

type WherePairs struct {
	pairs []pair.Pair[string, []any]
}

func (w *WherePairs) Append(param string, value ...any) {
	w.pairs = append(w.pairs, pair.Pair[string, []any]{
		Key: param, Value: value,
	})
}

func (w *WherePairs) And(other *WherePairs) *WherePairs {
	w.pairs = append(w.pairs, other.pairs...)
	return w
}

func (w *WherePairs) AndWhere() (whereStr string, values []any) {
	keys := make([]string, 0, len(w.pairs))
	values = make([]any, 0, len(w.pairs))
	for _, item := range w.pairs {
		keys = append(keys, item.Key)
		values = append(values, item.Value...)
	}
	whereStr = strings.Join(keys, " and ")
	return
}
