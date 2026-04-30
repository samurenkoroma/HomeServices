package farm

import (
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/module"
	"samurenkoroma/services/internal/application/query"
	"samurenkoroma/services/internal/core/domain/repository"
	farmCommands "samurenkoroma/services/internal/modules/farm/application/commands"
	farmQueries "samurenkoroma/services/internal/modules/farm/application/queries"
	farmProjections "samurenkoroma/services/internal/modules/farm/infrastructure/projections"
	"samurenkoroma/services/pkg/utils"
)

type farmModule struct {
	Commands []module.CommandHandler
	Queries  []module.QueryHandler
}

func NewModule(uowFactory repository.Factory) module.Module {
	farmProvider := farmProjections.NewFarmProjectionsProvider(uowFactory.DB())
	h := farmCommands.NewFarmObjectHandler(uowFactory)
	return &farmModule{
		Commands: []module.CommandHandler{{
			Name:    "CreateObject",
			Handler: h.Create,
			Decoder: utils.DecodeJSON[farmCommands.CreateFarmObjectCmd],
		}, {
			Name:    "UpdateObject",
			Handler: h.Update,
			Decoder: utils.DecodeJSON[farmCommands.UpdateFarmObjectCommand],
		}, {
			Name:    "DeleteObject",
			Handler: h.Delete,
			Decoder: utils.DecodeJSON[farmCommands.DeleteFarmObjectCommand],
		}},
		Queries: []module.QueryHandler{{
			Name:    "GetCurrentFarm",
			Handler: farmQueries.NewFarmHandler(farmProvider.Objects()).CurrentFarm,
			Decoder: utils.DecodeJSON[farmQueries.GetCurrentFarmQuery],
		}, {
			Name:    "GetObjects",
			Handler: farmQueries.NewFarmHandler(farmProvider.Objects()).GetPhysicalObjects,
			Decoder: utils.DecodeJSON[farmQueries.GetPhysicalObjectsQuery],
		},
			//{
			//	Name:    "MyOrganizations",
			//	Handler: farmQueries.NewMyOrganizationsHandler(),
			//	Decoder: utils.DecodeJSON[farmQueries.GetMyOrganizationsQuery],
			//},
			{
				Name:    "CurrentOrganization",
				Handler: nil,
				Decoder: nil,
			},
		},
	}

}
func (f *farmModule) RegisterCommands(router command.Router) {
	for _, cmd := range f.Commands {
		router.Register(cmd.Name, cmd.Handler, cmd.Decoder)
	}
}

func (f *farmModule) RegisterQueries(router query.Router) {
	for _, q := range f.Queries {
		router.Register(q.Name, q.Handler, q.Decoder)
	}
}
