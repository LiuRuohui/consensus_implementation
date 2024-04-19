package main

import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
    "crypto/sha256"
    "encoding/hex"
)

type Block struct {
    Index int
    Timestamp string
    Prehash string
    Hash string
    Data []byte

    delegate *Node// 代理 区块由哪个节点挖出
}


func GenesisBlock()  Block {
    gene := Block{0, time.Now().String(),"", "", []byte("genesis block"), nil}

    gene.Hash = string(blockHash(gene))

    var delegate *Node = new(Node)
    delegate.Name = "Genis Block"
    delegate.Votes = 0;
    gene.delegate = delegate

    return gene
}

func blockHash(block Block) []byte  {

    record := strconv.Itoa(block.Index) + block.Timestamp + block.Prehash + hex.EncodeToString(block.Data)

    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hashed
}


//节点类型
type Node struct {
    Name  string //节点名称
    Votes int    // 被选举的票数
}

func (node *Node)GenerateNewBlock(lastBlock Block, data []byte) Block  {

    var newBlock = Block{lastBlock.Index+1, time.Now().String(), lastBlock.Hash, "", data, nil}

    newBlock.Hash = hex.EncodeToString(blockHash(newBlock))
    newBlock.delegate = node
    return newBlock

}

//创建节点
var NodeArr = make([]Node,10)
func CreateNode() {

    for i := 0; i < 10; i++ {
        name := fmt.Sprintf("NODE %d num", i+1)
        NodeArr[i] = Node{name, 0}
    }

}

//简单模拟投票
func Vote()  {
    for i := 0; i < 10; i++ {
        rand.Seed(time.Now().UnixNano())
        vote := rand.Intn(10) + 1
        NodeArr[i].Votes = vote
    }
}


//选出票数最多的前3位
func SortNodes() []Node  {
    n:= NodeArr
    for i := 0; i<len(n) ;i++  {
        for j := 0; j < len(n)-1 ;j++  {
            if n[j].Votes < n[j+1].Votes {
                n[j],n[j+1] = n[j+1],n[j]
            }
        }
    }

    return n[:3]
}


func main() {

    CreateNode()
    fmt.Println(NodeArr)
    Vote()
    nodes := SortNodes()

    fmt.Println(nodes)


    //创建创世区块
    gene := GenesisBlock()

    lastBlock := gene
    for i:= 0; i< len(nodes) ;i++  {
        //打印新区块信息
        // var node *Node
         node := lastBlock.delegate
         fmt.Println("区块号:",lastBlock.Index,"节点名称:",node.Name,"节点投票数:",node.Votes)

        lastBlock =  nodes[i].GenerateNewBlock(lastBlock,[]byte(fmt.Sprintf("new block %d",i)))
    }

}