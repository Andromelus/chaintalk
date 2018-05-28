package block

import tx "florianc/chaintalk/tx"

type Block struct {
	ID   int
	Hash string
	Txs  []tx.Tx
}
