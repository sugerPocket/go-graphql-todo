package fields
import (
	"go-graphql-todo/models"
	"go-graphql-todo/structs"
	"go-graphql-todo/types"

	"github.com/graphql-go/graphql"
)
/*Todo get a todo by id*/
var Todo = &graphql.Field{
	Type:        types.Todo,
	Description: "Get single todo",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		idQuery, isOK := params.Args["id"].(string)
		if isOK {
			// Search for el with id
			for _, todo := range models.List {
				if todo.ID == idQuery {
					return todo, nil
				}
			}
		}

		return structs.Todo{}, nil
	},
}