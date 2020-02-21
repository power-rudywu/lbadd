package command

import (
	"github.com/oklog/ulid"
	"github.com/tomarrell/lbadd/internal/parser/ast"
)

// Command is the intermediate representation (IR) of an SQL ast.
type Command struct {
	// ID is the ID of this command.
	ID ulid.ULID
	// Type is the type of the command, and represents the type of the SQL
	// statement that this command represents.
	Type Type
	// DataSources is a list of tables or views, that need to be accessed for
	// this command. For SELECT commands, this contains the relations in the
	// FROM clause. If the FROM clause contains a nested query that computes a
	// table, the respective entry in this list refers to a temporary table
	DataSources []DataSource
	// DataTarget is a table or view, whose data is affected by this
	// command (i. e. write). For an INSERT command, this is the table where the
	// values will be inserted into.
	DataTarget   DataTarget
	Dependencies []Command
}

type DataSource struct {
	// Name is the name of the data source, i. e. the name of a table or view.
	Name string
	// Temporary indicates whether the data source is temporary. A data source
	// is temporary, if it was created while executing this command, and does
	// not leave the scope of this command. Temporary data sources are data
	// sources, that were created while executing the dependency commands of
	// this command.
	Temporary bool
}

type DataTarget struct {
	Name string
}

// From converts the given (*ast.SQLStmt) to the IR, which is a
// (command.Command).
func From(stmt *ast.SQLStmt) (Command, error) {
	return Command{}, nil
}
