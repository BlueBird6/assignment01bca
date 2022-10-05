package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	nonce        int
	prev_hash    string
	current_hash string
	transaction  string
}

type Blockchain struct {
	list []*Block
}

func (ls *Blockchain) NewBlock(n int, t string) *Block {
	s := new(Block)
	s.nonce = n
	s.transaction = t

	if VerifyChain(ls) {
		ls.list = append(ls.list, s)
		CalculateHash(ls)
		fmt.Println("Block Added")
		return s
	} else {
		return nil
	}
}

func ListBlocks(stud *Blockchain) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	for i := 0; i < len(stud.list); i++ {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println("nonce:	", stud.list[i].nonce)
		fmt.Println("previous hash:	", stud.list[i].prev_hash)
		fmt.Println("current hash:	", stud.list[i].current_hash)
		fmt.Println("Transaction:	", stud.list[i].transaction)
	}

	fmt.Print("\n\n\n")

}
func (s *Block) getstring() string {

	var r = ""
	r += strconv.Itoa(s.nonce)
	r += s.transaction + s.prev_hash

	return r
}

func CalculateHash(stud *Blockchain) {

	for i := 0; i < len(stud.list); i++ {
		sum := sha256.Sum256([]byte(stud.list[i].getstring()))
		stud.list[i].current_hash = fmt.Sprintf("%x", sum)
		if i < len(stud.list)-1 {
			stud.list[i+1].prev_hash = fmt.Sprintf("%x", sum)
		}
	}

}

func VerifyChain(stud *Blockchain) bool {
	var st = ""
	for i := 0; i < len(stud.list); i++ {
		sum := sha256.Sum256([]byte(stud.list[i].getstring()))
		st = fmt.Sprintf("%x", sum)

		if st != stud.list[i].current_hash {
			fmt.Printf("Block is tempered, Block no %d\n", i)
			return false
		}

	}
	fmt.Println("Blocks are ok! ")
	return true

}
func ChangeBlock(stud *Blockchain, n int, t string) { // n = nonce of the block, t= new transaction or change to the block

	for i := 0; i < len(stud.list); i++ {
		if n == stud.list[i].nonce {

			stud.list[i].transaction = t
			fmt.Println("change done ")
			return
		}
	}

	fmt.Println("block not found!")

}

// func main() {

// 	blockchain := new(Blockchain)
// 	var non = 3000
// 	blockchain.NewBlock(non, "bob to alice: 1")
// 	non++
// 	blockchain.NewBlock(non, "alice to charlie: 2")
// 	non++
// 	blockchain.NewBlock(non, "bob to charlie: 1")
// 	non++
// 	blockchain.NewBlock(non, "charlie to harley: 5")
// 	non++
// 	blockchain.NewBlock(non, "alice to bob: 2")
// 	non++

// 	ListBlocks(blockchain)

// 	fmt.Println("changing a block with nonce 3001")

// 	ChangeBlock(blockchain, 3001, "alice to harley: 10")

// 	ListBlocks(blockchain)

// 	fmt.Println("Adding the new block ")

// 	blockchain.NewBlock(non, "new transaction")
// 	non++

// 	fmt.Println("Recalculating the hashes")

// 	CalculateHash(blockchain)

// 	fmt.Println("verifying the chain ")
// 	VerifyChain(blockchain)

// 	fmt.Println("")
// 	fmt.Println("Adding the new block ")

// 	blockchain.NewBlock(non, "new transaction")
// 	non++

// }
