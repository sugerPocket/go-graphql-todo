package fields
import (
	"go-graphql-todo/models"
	"go-graphql-todo/types"

	"github.com/graphql-go/graphql"
)
/*Last get last todo*/
var Last = &graphql.Field{
	Type:        types.Todo,
	Description: "Last todo added",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return models.List[len(models.List) - 1], nil
	},
}