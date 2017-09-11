package models
import (
	"go-graphql-todo/structs"
)
/*List global todo list*/
var List = append(
	[]structs.Todo{},
	structs.Todo{ID: "a", Text: "A todo not to forget", Done: false},
	structs.Todo{ID: "b", Text: "This is the most important", Done: false},
	structs.Todo{ID: "c", Text: "Please do this or else", Done: false},
)