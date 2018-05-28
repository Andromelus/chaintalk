package chain

import "fmt"
import block "florianc/block"
//import tx_pool "florianc/tx_pool"



func Init() {
	fmt.Println("Initiating chain")

}

func Mine() {
	createNewBlock()
}


func createNewBlock(){
	nBlock := new(block.Block)
	nBlock.Id = 1
	fmt.Println("created new block: Number ", nBlock.Id)
}