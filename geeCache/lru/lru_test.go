package lru

import (
	"fmt"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

type AAA struct {
	key string
	val string
}

func (d AAA) Len() int {
	return len(d.key) + len(d.val)
}

func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	//lru.Add("key1", AAA{
	//	key: "key1111",
	//	val: "val2222",
	//})
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	} else {
		fmt.Println(v.Len(), v)
	}
	if _, ok := lru.Get("key2"); !ok {
		t.Fatalf("cache miss key2 failed")
	}

}
