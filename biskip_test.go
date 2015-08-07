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

/*

func TestNormalListInsert(t *testing.T) {
	var (
		key   string
		score int64
		value int
	)
	type node struct {
		key   string
		score int64
		value interface{}
		next  *node
	}
	head := &node{``, 0, nil, nil}
	var nnode, curNode *node
	for i := 0; i < num; i++ {
		key, score, value = strconv.FormatInt(rand.Int63(), 10), rand.Int63(), rand.Int()
		nnode = &node{key, score, value, nil}
		for curNode = head; curNode.next != nil && curNode.next.score <= nnode.score; curNode = curNode.next {
		}
		nnode.next = curNode.next
		curNode.next = nnode
	}
}
*/
/*
func TestSnapshot(t *testing.T) {
	sl.Snapshot()
}*/

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
