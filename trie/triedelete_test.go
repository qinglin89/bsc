package trie

import (
	"fmt"
	"testing"
)

func TestTrieDelete(t *testing.T) {
	fn := &fullNode{}
	fn.Children[0] = valueNode([]byte{'a', 'b'})
	fn.Children[16] = valueNode([]byte{'d', 'e'})
	fn2 := &fullNode{}
	fn2.Children[0] = valueNode([]byte{'m', 'j'})
	fn2.Children[1] = fn

	mpt := &Trie{
		root: fn2,
		//		root: &shortNode{
		//			Key: []byte{5},
		//			Val: fn,
		//		},
	}
	//	fmt.Println(mpt.root.fstring(""))
	sn := mpt.root.(*fullNode)
	//sfn := sn.Children[1].(*fullNode)
	//	fmt.Println(sfn.Children[16].fstring(""))
	v, _, _, _ := mpt.tryGet(mpt.root, []byte{0}, 0)
	fmt.Println("value of key-0", string(v))
	v, _, _, _ = mpt.tryGet(mpt.root, []byte{1, 0, 1, 2, 3, 5, 6}, 0)
	fmt.Println("value of key-1-0 search with prefix match: ", string(v))
	v, _, _, _ = mpt.tryGet(mpt.root, []byte{1, 16, 1, 2, 3, 5, 6}, 0)
	fmt.Println("value of key-1-16 search with prefix match: ", string(v))

	_, n, _ := mpt.delete(mpt.root, nil, []byte{1, 0})
	//fmt.Println("delete result", b, n, e)
	mpt.root = n

	v, _, _, _ = mpt.tryGet(mpt.root, []byte{0}, 0)
	fmt.Println("value of key-0", string(v))

	v, _, _, _ = mpt.tryGet(mpt.root, []byte{1, 16}, 0)
	fmt.Println("value of key-1-16", string(v))
	sn, ok1 := mpt.root.(*fullNode)
	_, ok2 := sn.Children[1].(*shortNode)
	fmt.Println(ok1, ok2)
	mpt.root.(*fullNode).Children[1] = valueNode([]byte{'d', 'e'})
	v, _, _, _ = mpt.tryGet(mpt.root, []byte{1, 16}, 0)
	fmt.Println("value of key-1 searched as 1-16", string(v))

}
