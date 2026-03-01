package uow

import "samurenkoroma/services/internal/common/domain"

type aggregateRegistry struct {
	aggregates []domain.Aggregate
}

func (r *aggregateRegistry) add(a domain.Aggregate) {
	r.aggregates = append(r.aggregates, a)
}

func (r *aggregateRegistry) list() []domain.Aggregate {
	return r.aggregates
}
