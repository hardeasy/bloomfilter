package bloomfilter

import (
	"strconv"
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
	b.Add("0")
	b.Add("haha2")
	b.Add("haha3")
	ok := b.Check("10")
	t.Log("in", ok)
}

//100kb 100w
func TestTemp(t *testing.T) {
	b := NewBloomFilter(1024 * 100)
	for i := 1; i < 1000000; i++ {
		b.Add(strconv.Itoa(i))
	}
	var ecount int = 0
	var ok bool
	for j := 1000001; j < 2000001; j++ {
		ok = b.Check(strconv.Itoa(j))
		if ok {
			ecount++
		}

	}
	t.Log(ecount)

}

func BenchmarkAdd(b *testing.B) {
	bf := NewBloomFilter(10240)
	for i := 1; i < b.N; i++ {
		_ = bf.Add(string(i))
	}
}
