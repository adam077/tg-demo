package utils

var ConstData = new(StructConstData)

type StructConstData struct {
}

func init() {
	FillWithDefault(ConstData)
}
