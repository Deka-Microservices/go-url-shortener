package base62

import "math/big"

func Encode(val uint32) string {
	var i big.Int
	i.SetUint64(uint64(val))
	return i.Text(62)
}
