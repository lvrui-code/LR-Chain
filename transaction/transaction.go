package transaction

type Transaction struct {
	ID     []byte     // 哈希值
	Input  []TxInput  // 标记支持本次转账的前置的交易信息的TxOutput
	Output []TxOutput // 记录本次转账的amount和Reciever
}
