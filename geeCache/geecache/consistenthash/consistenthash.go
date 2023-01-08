package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash maps bytes to uint32
type Hash func(data []byte) uint32

// Map contains all hashed keys
type Map struct {
	hash     Hash           //Hash 函数
	replicas int            //虚拟节点倍数(每个节点名称创建的节点个数
	keys     []int          // Sorted 哈希环
	HashMap  map[int]string //虚拟节点与真实节点的映射表
}

// New creates a Map instance
func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		HashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash.
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.HashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}

	hash := int(m.hash([]byte(key)))
	//fmt.Println("hash", hash)
	//Binary search for appropriate replica.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	//fmt.Println("m.keys:", m.keys)
	//fmt.Println("idx:", idx, len(m.keys))
	//fmt.Println("idx%len(m.keys)", idx%len(m.keys))
	//fmt.Println("m.keys[idx%len(m.keys)]", m.keys[idx%len(m.keys)])
	//fmt.Println("m.HashMap[m.keys[idx%len(m.keys)]]", m.HashMap[m.keys[idx%len(m.keys)]])
	return m.HashMap[m.keys[idx%len(m.keys)]]
}
