package httpapi

import (
	"context"
	"net/http"
	"samurenkoroma/services/internal/growing/application"
)

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {

	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}

func TransactionMiddleware(uowFactory application.UnitOfWorkFactory) Middleware {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			uow, err := uowFactory.New(ctx)
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			ctx = context.WithValue(ctx, "uow", uow)

			defer func() {
				if rec := recover(); rec != nil {
					uow.Rollback()
					panic(rec)
				}
			}()

			next.ServeHTTP(w, r.WithContext(ctx))

			uow.Commit()
		})
	}
}
