package structs
import (
	"time"
)
/*Todo schema define*/
type Todo struct {
	ID      string 	  `json:"id"`
	Title   string	  `json:"title"`
	Content string 	  `json:"content"`
	Start   time.Time `json:"start"`
	End     time.Time `json:"end"`
}