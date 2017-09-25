package types
import (
	"github.com/graphql-go/graphql"
)
/*Todo graphql todo type*/
var Todo = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"start": &graphql.Field{
			Type: graphql.DateTime,
		},
		"end": &graphql.Field{
			Type: graphql.DateTime,
		},
	},
})