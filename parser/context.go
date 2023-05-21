package parser

import (
	bbnparams "github.com/babylonchain/babylon/app/params"

	"github.com/forbole/juno/v4/logging"
	"github.com/forbole/juno/v4/node"

	"github.com/forbole/juno/v4/database"
	"github.com/forbole/juno/v4/modules"
)

// Context represents the context that is shared among different workers
type Context struct {
	EncodingConfig *bbnparams.EncodingConfig
	Node           node.Node
	Database       database.Database
	Logger         logging.Logger
	Modules        []modules.Module
}

// NewContext builds a new Context instance
func NewContext(
	encodingConfig *bbnparams.EncodingConfig, proxy node.Node, db database.Database,
	logger logging.Logger, modules []modules.Module,
) *Context {
	return &Context{
		EncodingConfig: encodingConfig,
		Node:           proxy,
		Database:       db,
		Modules:        modules,
		Logger:         logger,
	}
}
