package module

import (
	"samurenkoroma/services/internal/application/command"
	"samurenkoroma/services/internal/application/query"
)

type CommandHandler struct {
	Name    string
	Handler command.Handler
	Decoder command.DecoderFunc
}

type QueryHandler struct {
	Name    string
	Handler query.Handler
	Decoder query.Decoder
}

type Module interface {
	RegisterCommands(router command.Router)
	RegisterQueries(router query.Router)
}
