package utils
import (
	"time"
	"math/rand"
)
func init() {
	rand.Seed(time.Now().UnixNano())
}
/*RandStringRunes 在字符串中取随机串*/
func RandStringRunes(n int, letterRunes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}