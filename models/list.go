package models
import (
	"time"
	"go-graphql-todo/structs"
)
/*List global todo list*/
var List = append(
	[]structs.Todo{},
	structs.Todo{
		ID: "a",
		Title:"a",
		Content: "A todo not to forget",
		Start: time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC),
		End: time.Date(2017, 9, 14, 0, 0, 0, 0, time.UTC),
	},
	structs.Todo{
		ID: "b",
		Title:"b",
		Content: "This is the most important",
		Start: time.Date(2017, 9, 13, 0, 0, 0, 0, time.UTC),
		End: time.Date(2017, 9, 14, 17, 0, 0, 0, time.UTC),
	},
	structs.Todo{
		ID: "c",
		Title:"c",
		Content: "Please do this or else",
		Start: time.Date(2017, 9, 15, 0, 0, 0, 0, time.UTC),
		End: time.Date(2017, 9, 16, 17, 0, 0, 0, time.UTC),
	},
)