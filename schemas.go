package main
import (
	"go-graphql-todo/fields"
	"github.com/graphql-go/graphql"
)

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	// root query
	// we just define a trivial example here, since root query is required.
	// Test with curl
	// curl -g 'http://localhost:8080/graphql?query={last{id,text,done}}'
	Query:    graphql.NewObject(graphql.ObjectConfig{
		Name: "query",
		Fields: graphql.Fields{
			/*curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'*/
			"todo": fields.Todo,
			"last": fields.Last,
			/*curl -g 'http://localhost:8080/graphql?query={list{id,text,done}}'*/
			"list": fields.List,
		},
	}),
	// root query
	// we just define a trivial example here, since root query is required.
	// Test with curl
	// curl -g 'http://localhost:8080/graphql?query={lastTodo{id,text,done}}'
	Mutation: graphql.NewObject(graphql.ObjectConfig{
		Name: "mutation",
		Fields: graphql.Fields{
			/*
				curl -g 'http://localhost:8080/graphql?query=mutation+_{create(text:"My+new+todo"){id,text,done}}'
			*/
			"create": fields.Create,
			/*curl -g 'http://localhost:8080/graphql?query=mutation+_{update(id:"a",done:true){id,text,done}}'*/
			"update": fields.Update,
		},
	}),
})
