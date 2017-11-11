package bloomfilter

import (
	"testing"
)

func TestMain(t *testing.T) {
	b := NewBloomFilter(5)
	t.Log(b)
}

func TestBloomFilter_Hash(t *testing.T) {
	b := NewBloomFilter(5)
	result := b.hash("haha", 111)
	t.Log(result)
}

func TestBloomFilter_hashStrs(t *testing.T) {
	b := NewBloomFilter(5)
	hasharr := b.hashStrs("haha")
	t.Log(hasharr)
}

func TestBloomFilter_Add(t *testing.T) {
	b := NewBloomFilter(5)
	b.Add("haha")
	b.Add("hahas")
}

func TestBloomFilter_Check(t *testing.T) {
	b := NewBloomFilter(5)
	b.Add("haha")
	b.Add("haha2")
	b.Add("haha3")
	ok := b.Check("haha3")
	t.Log("in", ok)
}
