package fields
import (
	"go-graphql-todo/models"
	"go-graphql-todo/types"

	"github.com/graphql-go/graphql"
)
/*List get todo list*/
var List = &graphql.Field{
	Type:        graphql.NewList(types.Todo),
	Description: "List of todos",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return models.List, nil
	},
}