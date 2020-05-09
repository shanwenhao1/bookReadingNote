package hashFunc

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/spaolacci/murmur3"
)

var str = "hello world"

func md5Hash() [16]byte {
	return md5.Sum([]byte(str))
}

func sha1Hash() [20]byte {
	return sha1.Sum([]byte(str))
}

func murmur32() uint32 {
	return murmur3.Sum32([]byte(str))
}

func murmur64() uint64 {
	return murmur3.Sum64([]byte(str))
}

var bucketSize = 10

func murmur64Func(p string) uint64 {
	return murmur3.Sum64([]byte(p))
}

/*
	以murmur3为例，我们先以15810000000开头，造一千万个和手机号类似的数字，然后将计算后的哈希值分十个桶，并观察计数是否均匀
*/
func MurmurHash() {
	var bucketMap = map[uint64]int{}
	for i := 15810000000; i < 15810000000+10000000; i++ {
		hashInt := murmur64Func(fmt.Sprint(i)) % uint64(bucketSize)
		// 各个桶如果分得数则加一, 用于验证hash是否均匀
		bucketMap[hashInt]++
	}
	fmt.Println(bucketMap)
}
