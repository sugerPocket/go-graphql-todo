package services
import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "root:@127.0.0.1")
}