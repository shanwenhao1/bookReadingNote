package hashFunc

import "testing"

func BenchmarkMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5Hash()
	}
}

func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha1Hash()
	}
}

/*
	MurmurHash比MD5和sha1性能有巨大提升, 因此更适合用于负载均衡
*/
func BenchmarkMurmurHash32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur32()
	}
}

func BenchmarkMurmurHash64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		murmur64()
	}
}
