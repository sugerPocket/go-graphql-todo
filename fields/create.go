package fields
import (
	"go-graphql-todo/models"
	"go-graphql-todo/structs"
	"go-graphql-todo/utils"
	"go-graphql-todo/types"

	"github.com/graphql-go/graphql"
)
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
/*Create add a todo*/
var Create = &graphql.Field{
	Type:        types.Todo, // the return type for this field
	Description: "Create new todo",
	Args: graphql.FieldConfigArgument{
		"text": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		text, _ := params.Args["text"].(string)

		// figure out new id
		newID := utils.RandStringRunes(8, letterRunes)

		// perform mutation operation here
		// for e.g. create a schemas.Todo and save to DB.
		newTodo := structs.Todo{
			ID:   newID,
			Text: text,
			Done: false,
		}

		models.List = append(models.List, newTodo)

		// return the new schemas.Todo object that we supposedly save to DB
		// Note here that
		// - we are returning a `schemas.Todo` struct instance here
		// - we previously specified the return Type to be `TodoType`
		// - `schemas.Todo` struct maps to `TodoType`, as defined in `TodoType` ObjectConfig`
		return newTodo, nil
	},
}
