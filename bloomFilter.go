package bloomfilter

//"fmt"

type bloomFilter struct {
	data     []byte
	bitSize  int64
	byteSize int64
}

//初始值数组
var seed []int = []int{3, 5, 7, 11, 13, 31, 37}

//添加值
func (self *bloomFilter) Add(strs string) error {
	//hash值数组
	hasharr := self.hashStrs(strs)
	//fmt.Println(hasharr)

	//存储
	var byteIndex int64
	var bitIndex int64
	var temp byte
	for i := 0; i < len(hasharr); i++ {
		byteIndex = hasharr[i] / 8
		bitIndex = hasharr[i] % 8
		//fmt.Println("arr", hasharr[i], "byteIndex", byteIndex, "bitIndex", bitIndex)
		temp = 1
		//fmt.Printf("begin:%b\n", self.data[byteIndex])
		self.data[byteIndex] = (temp << uint(bitIndex)) | self.data[byteIndex]
		//fmt.Printf("after:%b\n", self.data[byteIndex])
	}
	//fmt.Printf("selft.data:%b\n", self.data)
	return nil
}

//检测值
func (self *bloomFilter) Check(strs string) bool {
	//hash值数组
	hasharr := self.hashStrs(strs)
	var byteIndex int64
	var bitIndex int64
	var temp byte

	for i := 0; i < len(hasharr); i++ {
		byteIndex = hasharr[i] / 8
		bitIndex = hasharr[i] % 8
		temp = 1
		if (self.data[byteIndex] & (temp << uint(bitIndex))) == 0 {
			return false
		}
	}
	//检测
	return true
}

//得到hash值数组
func (self *bloomFilter) hashStrs(strs string) []int64 {
	hasharr := make([]int64, len(seed))
	var item int64
	for i := 0; i < len(seed); i++ {
		item = self.hash(strs, seed[i])
		hasharr[i] = item
	}
	return hasharr
}

//单个hash值
func (self *bloomFilter) hash(strs string, seed int) int64 {
	hash := 5381
	for i := 0; i < len(strs)-1; i++ {
		hash += (hash << 5) + int(strs[i])
	}
	hash += seed
	return int64(hash) % self.bitSize
}

//新建
func NewBloomFilter(byteSize int64) *bloomFilter {
	b := &bloomFilter{
		data:     make([]byte, byteSize),
		bitSize:  int64(byteSize * 8),
		byteSize: byteSize,
	}
	return b

}
