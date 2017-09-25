package fields
import (
	"time"
	
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
		"title": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"start": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"end": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		// marshall and cast the argument value
		content, _ := params.Args["content"].(string)
		title, _ := params.Args["title"].(string)
		start, _ := params.Args["start"].(string)
		end, _ := params.Args["end"].(string)

		// figure out new id
		newID := utils.RandStringRunes(8, letterRunes)

		// perform mutation operation here
		// for e.g. create a schemas.Todo and save to DB.
		Start, _ := time.Parse("YYYY/MM/DD HH:mm", start)
		End, _ := time.Parse("YYYY/MM/DD HH:mm", end)
		newTodo := structs.Todo{
			ID:   newID,
			Title: title,
			Content: content,
			Start: Start,
			End: End,
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
