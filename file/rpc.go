package main

import (
	"fmt"
	"net"
	"net/rpc"
	"strconv"
	"time"
)

type Cluster struct {
	NodeList []*NodeInfo
}

type NodeInfo struct {
	Ip   string
	Port int
}

func (this *NodeInfo) String() string {
	return this.Ip + ":" + strconv.Itoa(this.Port)
}

type Node struct {
	name      string
	nodeInfo  *NodeInfo
	cluster   *Cluster
	rpcServer *rpc.Server
}

func NewNode(ip string, port int) *Node {
	nodeInfo := &NodeInfo{ip, port}
	name := strconv.Itoa(port)
	nodeList := make([]*NodeInfo, 1)
	nodeList[0] = nodeInfo
	cluster := &Cluster{nodeList}
	rpcServer := rpc.NewServer()
	return &Node{name, nodeInfo, cluster, rpcServer}
}

func (this *Node) start() {
	this.rpcServer.Register(this)
	port := strconv.Itoa(this.nodeInfo.Port)
	listener, e := net.Listen("tcp", ":"+port)
	if e != nil {
		fmt.Println(e)
		return
	} else {
		fmt.Println(this.name + ":listen on :" + port)
	}
	go this.loop(listener)
}
func (this *Node) loop(listener net.Listener) {
	for {
		if conn, err := listener.Accept(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(this.name + ":get new client")
			go this.rpcServer.ServeConn(conn)
		}
	}
}

func (this *Node) letMeInRequest(ip string, port int) {
	client, err := rpc.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return
	}
	response := new(NodeInfo)
	callErr := client.Call("Node.LetMeInResponse", this.nodeInfo, response)
	if callErr != nil {
		fmt.Println(callErr)
	} else {
		fmt.Println(this.name + ":response:" + response.String())
		this.cluster.NodeList = append(this.cluster.NodeList, response)
	}
}

func (this *Node) LetMeInResponse(request *NodeInfo, response *NodeInfo) error {
	fmt.Println(this.name + ":request:" + request.String())
	this.cluster.NodeList = append(this.cluster.NodeList, request)
	response.Ip = this.nodeInfo.Ip
	response.Port = this.nodeInfo.Port
	return nil
}
func main() {
	firstNode := NewNode("127.0.0.1", 1234)
	secondNode := NewNode("127.0.0.1", 1235)
	firstNode.start()
	secondNode.start()
	time.Sleep(1 * time.Second)
	firstNode.letMeInRequest("127.0.0.1", 1235)
	time.Sleep(1 * time.Second)
	fmt.Println(firstNode.cluster.NodeList)
	fmt.Println(secondNode.cluster.NodeList)
}
