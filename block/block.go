package block

import tx "florianc/tx"


type Block struct {
	Id int
	Hash string
	Txs []tx.Tx
}