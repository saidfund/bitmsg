package chunk

func GetHash(b byte) [32]byte {
	var a [32]byte
	for i := 0; i < 32; i++ {
		a[i] = b
	}
	return a
}
