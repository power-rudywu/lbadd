package executor

import (
	"fmt"

	"github.com/oklog/ulid"
	"github.com/rs/zerolog"
	"github.com/tomarrell/lbadd/internal/executor/command"
)

var _ Executor = (*simpleExecutor)(nil)

type simpleExecutor struct {
	log zerolog.Logger
}

func newSimpleExecutor(log zerolog.Logger) *simpleExecutor {
	return &simpleExecutor{
		log: log,
	}
}

func (e *simpleExecutor) Execute(cmd command.Command) (Result, error) {
	tempDataSources := make(map[ulid.ULID]Result)

	// execute all dependencies of that command recursively
	for _, dependency := range cmd.Dependencies {
		result, err := e.Execute(dependency)
		if err != nil {
			return nil, fmt.Errorf("execute dependency: %w", err)
		}
		tempDataSources[dependency.ID] = result
	}

	// after all dependencies are executed, handle the actual command
	switch cmd.Type {
	case command.Select:
		return e.executeSelect(cmd, tempDataSources)
	}
	return nil, fmt.Errorf("unimplemented")
}

func (e *simpleExecutor) executeSelect(cmd command.Command, tempDataSources map[ulid.ULID]Result) (Result, error) {

	return nil, fmt.Errorf("unimplemented")
}
