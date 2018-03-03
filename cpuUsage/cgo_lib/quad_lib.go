package cgo_lib

func Quad(i int) int {
	return Double(Double(i))
}
