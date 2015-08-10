package biskiplist

import (
	"fmt"
	//	`math/rand`
	"strconv"
	"testing"
)

var num int = 32
var findnum int = 100000
var sl *BiSkiplist = NewBiSkiplist(32)

func TestBiSkiplistInsert(t *testing.T) {
	var (
		key   string
		score int64
		value int
	)
	for i := 0; i < num; i++ {
		key, score, value = strconv.FormatInt(int64(i), 10), int64(i), i
		sl.Set(key, score, value)
	}
}

func TestRemove(t *testing.T) {
	var key string
	for i := 0; i < num; i++ {
		key = strconv.FormatInt(int64(i), 10)
		sl.Remove(key)
	}
}

func TestSkiplistIterate(t *testing.T) {
	it := sl.Iterate()
	for node := it.Next(); node != nil; node = it.Next() {
		fmt.Print(node, `  `)
	}
}

/*
func TestSkiplistContain(t *testing.T) {
	fmt.Println(sl.Contains(6))
}
*/
/*
func TestBiSkiplistGet(t *testing.T) {
	fmt.Println(sl.Get(rand.Intn(num)), sl.level)
}
*/

func TestRange(t *testing.T) {
	fmt.Println(sl.RangeByKey(`564`, 40))
}
