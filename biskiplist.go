// Copyright 2009 . All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// @author sunlnx
package biskiplist

import (
	"errors"
	"fmt"
	`math/rand`
)

type BiSkiplist struct {
	nIndex   map[string]*biNode
	head     *biNode
	maxLevel int
	level    int
}

type biNode struct {
	key      string
	score    int64
	value    interface{}
	forward  []*biNode
	backward []*biNode
}

func NewBiSkiplist(maxLevel int) *BiSkiplist {
	return &BiSkiplist{
		nIndex:   make(map[string]*biNode),
		head:     newBiNode(`#`, 0, nil, maxLevel),
		maxLevel: maxLevel,
		level:    1,
	}
}

func newBiNode(key string, score int64, value interface{}, level int) *biNode {
	return &biNode{
		key:      key,
		score:    score,
		value:    value,
		forward:  make([]*biNode, level),
		backward: make([]*biNode, level),
	}
}

func (sl *BiSkiplist) Total() int {
	return len(sl.nIndex)
}

func (sl *BiSkiplist) Set(key string, score int64, value interface{}) bool {
	if _, ok := sl.nIndex[key]; !ok {
		var level = 0
		for r := rand.Int(); (r & 1) == 1; r >>= 1 {
			level++
			if level == sl.level {
				if sl.level < sl.maxLevel {
					sl.level++
					break
				} else {
					r = rand.Int()
				}
			}
		}
		var node *biNode = newBiNode(key, score, value, level+1)
		var curBiNode *biNode = sl.head
		for i := sl.level - 1; i >= 0; i-- {
			for ; curBiNode.forward[i] != nil; curBiNode = curBiNode.forward[i] {
				if curBiNode.forward[i].score > score {
					break
				}
			}
			if i <= level {
				node.forward[i] = curBiNode.forward[i]
				if curBiNode.forward[i] != nil {
					curBiNode.forward[i].backward[i] = node
				}
				curBiNode.forward[i] = node
				node.backward[i] = curBiNode
			}
		}
		sl.nIndex[key] = node
		return true
	}
	return false
}

func (sl *BiSkiplist) Promt(key string, score int64) {
	if value, err := sl.Remove(key); err == nil {
		sl.Set(key, score, value)
	}
}

func (sl *BiSkiplist) Get(key string) (interface{}, bool) {
	peer, ok := sl.nIndex[key]
	return peer.value, ok
}

func (sl *BiSkiplist) Score(key string) (int64, error) {
	if node, ok := sl.nIndex[key]; ok {
		return node.score, nil
	}
	return -1, NoExistKeyErr
}

func (sl *BiSkiplist) Contains(key string) bool {
	_, ok := sl.nIndex[key]
	return ok
}

func (sl *BiSkiplist) Remove(key string) (interface{}, error) {
	if node, ok := sl.nIndex[key]; ok && node != nil && node != sl.head {
		delete(sl.nIndex, key)
		for i := len(node.forward) - 1; i >= 0; i-- {
			if node.forward[i] != nil {
				node.forward[i].backward[i] = node.backward[i]
			}
			node.backward[i].forward[i] = node.forward[i]
		}
		return node.value, nil
	}
	return nil, NoExistKeyErr
}

// 获取以key结点为中心, score值离其最小的的limit个节点
func (sl *BiSkiplist) RangeByKey(key string, limit int) []interface{} {
	if node, ok := sl.nIndex[key]; ok {
		var (
			forNode, bacNode *biNode
			i                int
			res              []interface{} = make([]interface{}, limit)
		)
		forNode, bacNode = node.forward[0], node.backward[0]
		for forNode != nil && bacNode != sl.head && i < limit {
			if forNode.score-node.score < node.score-bacNode.score {
				res[i] = forNode.value
				forNode = forNode.forward[0]
				i++
			} else {
				res[i] = bacNode.value
				bacNode = bacNode.backward[0]
				i++
			}
		}
		for forNode != nil && i < limit {
			res[i] = forNode.value
			forNode = forNode.forward[0]
			i++
		}
		for bacNode != sl.head && i < limit {
			res[i] = bacNode.value
			bacNode = bacNode.backward[0]
			i++
		}
		return res
	}
	return nil
}

// 打印biskiplist的节点快照，用于分析性能
func (sl *BiSkiplist) Snapshot() {
	var node *biNode
	for i := sl.level - 1; i >= 0; i-- {
		fmt.Println()
		fmt.Print(i, ` level: `)
		for node = sl.head.forward[i]; node != nil; node = node.forward[i] {
			fmt.Print(node.key, `:`, node.score, `    `)
		}
		fmt.Println()

	}
}

func (sl *BiSkiplist) Iterate() *BiIterate {
	return &BiIterate{
		sl:        sl,
		curBiNode: sl.head,
	}
}

type BiIterate struct {
	sl        *BiSkiplist
	curBiNode *biNode
}

func (it *BiIterate) Next() interface{} {
	if it.curBiNode.forward[0] != nil {
		it.curBiNode = it.curBiNode.forward[0]
		return it.curBiNode.value
	}
	return nil
}

var (
	NoExistKeyErr = errors.New(`key not exist in biskiplist`)
)
