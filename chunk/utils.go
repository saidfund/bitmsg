package chunk

func GetBytes48(src []byte)([48]byte){
	var dest [48]byte

	if len(src) > 48{
		src = src[:48]
	}
	var i = 0
	for ; i< len(src); i++{
		dest[i] = src[i]
	}
	for ; i < 48; i++{
		dest[i] = 0x0
	}
	return dest
}