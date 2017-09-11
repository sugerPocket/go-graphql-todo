package fields
import (
	"go-graphql-todo/models"
	"go-graphql-todo/structs"
	"go-graphql-todo/types"

	"github.com/graphql-go/graphql"
)
/*Update update a todo*/
var Update = &graphql.Field{
	Type:        types.Todo, // the return type for this field
	Description: "Update existing todo, mark it done or not done",
	Args: graphql.FieldConfigArgument{
		"done": &graphql.ArgumentConfig{
			Type: graphql.Boolean,
		},
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// marshall and cast the argument value
		done, _ := params.Args["done"].(bool)
		id, _ := params.Args["id"].(string)
		affectedTodo := structs.Todo{}

		// Search list for todo with id and change the done variable
		for i := 0; i < len(models.List); i++ {
			if models.List[i].ID == id {
				models.List[i].Done = done
				// Assign updated todo so we can return it
				affectedTodo = models.List[i]
				break
			}
		}
		// Return affected todo
		return affectedTodo, nil
	},
}
