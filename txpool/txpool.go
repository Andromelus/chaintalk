package txpool

import (
	tx "florianc/chaintalk/tx"
	"fmt"
)

var txpool []tx.Tx
var txpoolqueue []tx.Tx

// Init : initiate txpool
func Init() {
	fmt.Println("Initiating txpool")
}

// QueueTx : Queue a transaction to be added into the transaction pool
func QueueTx(tx tx.Tx) {
	var length = len(txpoolqueue)
	fmt.Println("Adding a transaction to txpoolqueue at position %#v\n ", length)
	txpoolqueue[length] = tx
	checkTxValidity(tx)
}

func addTxToPool(tx tx.Tx) {
	var length = len(txpool)
	fmt.Println("Adding a transaction to txpool at position %#v\n ", length)
	txpool[length] = tx
}

func checkTxValidity(tx tx.Tx) {
	var isValid = false
	// check
	if isValid {
		addTxToPool(tx)
	} else {
		fmt.Println("The transaction %#v\n is not valid and has not been added to txpool", tx.Hash)
	}
}
