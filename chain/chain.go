package chain

import (
	block "florianc/chaintalk/block"
	"florianc/chaintalk/txpool"
	"fmt"
)

// Init : Initiate chain
func Init() {
	fmt.Println("Initiating chain")
	txpool.Init()

}

// Mine : Mine a block
func Mine() {
	createNewBlock()
}

func createNewBlock() {
	nBlock := new(block.Block)
	nBlock.ID = 1
	fmt.Println("created new block: Number ", nBlock.ID)
}
