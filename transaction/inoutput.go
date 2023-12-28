package transaction

type TxOutput struct {
	Value     int    // 转出的资产值
	ToAddress []byte // 资产的接收者的地址
}

type TxInput struct {
	TxID        []byte // 指明支持本次交易的前置交易信息
	OutIdx      int    // 指明是前置交易信息中的第几个Output
	FromAddress []byte // 资产转出者的地址
}
