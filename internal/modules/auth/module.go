package auth

import (
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/module"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	"samurenkoroma/services/internal/modules/auth/application/commands/organization"
	"samurenkoroma/services/internal/modules/auth/application/queries"
	"samurenkoroma/services/internal/modules/auth/infrastructure/jwt"
	"samurenkoroma/services/pkg/utils"
)

type authModule struct {
	Commands []module.CommandHandler
	Queries  []module.QueryHandler
}

func NewModule(uowFactory repository.Factory, jwtService *jwt.Service) module.Module {
	h := organization.NewOrganizationHandler(uowFactory, jwtService)
	return &authModule{

		Commands: []module.CommandHandler{
			{
				Name:    "SwitchOrganization",
				Handler: h.Switch,
				Decoder: utils.DecodeJSON[organization.SwitchOrganizationCmd],
			},
			{
				Name:    "CreateOrganization",
				Handler: h.Create,
				Decoder: utils.DecodeJSON[organization.CreateOrganizationCmd],
			},
		},
		Queries: []module.QueryHandler{
			{
				Name:    "Me",
				Handler: queries.NewUserHandler(uowFactory, jwtService).Me,
				Decoder: utils.DecodeJSON[queries.MeQuery],
			},
		},
	}

}
func (f *authModule) RegisterCommands(router command.Router) {
	for _, cmd := range f.Commands {
		router.Register(cmd.Name, cmd.Handler, cmd.Decoder)
	}
}

func (f *authModule) RegisterQueries(router query.Router) {
	for _, q := range f.Queries {
		router.Register(q.Name, q.Handler, q.Decoder)
	}
}
