package fields
import (
	"time"
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
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"title": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"start": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"end": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		// marshall and cast the argument value
		title, _ := params.Args["title"].(string)
		content, _ := params.Args["content"].(string)
		id, _ := params.Args["id"].(string)
		start, _ := params.Args["start"].(string)
		end, _ := params.Args["end"].(string)
		affectedTodo := structs.Todo{}

		// Search list for todo with id and change the done variable
		for i := 0; i < len(models.List); i++ {
			if models.List[i].ID == id {
				models.List[i].Content = content
				models.List[i].Title = title
				models.List[i].End, _ = time.Parse("YYYY/MM/DD HH:mm", end)
				models.List[i].Start, _ = time.Parse("YYYY/MM/DD HH:mm", start)
				// Assign updated todo so we can return it
				affectedTodo = models.List[i]
				break
			}
		}
		// Return affected todo
		return affectedTodo, nil
	},
}
