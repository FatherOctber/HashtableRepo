package xmlrpc

import (
	"HashtableRepo/hashtable"
	"bufio"
	"bytes"
	"fmt"
	"github.com/divan/gorilla-xmlrpc/xml"
	"github.com/gorilla/rpc"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"time"
)

func Create(isType string) {

	if isType == "server" {
		Server()
	} else {
		Client()
	}
}

type MessageService struct{}

func (ms *MessageService) Print(r *http.Request, args *struct{ Who string }, reply *struct{ Message string }) error {
	log.Println("Call RPC Print")
	//hashtable.Print()

	var p hashtable.Pair
	hashtable.Put(p.New(1, "one"))
	/*hashtable.Put(p.New(2, "two"))
	hashtable.Put(p.New(5, "five"))
	hashtable.Put(p.New(7, "seven"))
	hashtable.Put(p.New(12, "twelve"))
	hashtable.Put(p.New(49, "fourty-nine"))
	hashtable.Put(p.New(33, "thirty-three"))
	hashtable.Put(p.New(8, "eigth"))*/

	str := hashtable.ToString()
	log.Println("Hashtable: " + str)
	reply.Message = "Hello, " + args.Who + "!"
	return nil
}

func (ms *MessageService) clear(r *http.Request, args *struct{ temp string }, reply *struct{ Message string }) error {
	log.Println("Call RPC clear")

	isClear := hashtable.Clear()
	log.Println("Is clear hashtable: " + strconv.FormatBool(isClear))
	if isClear {
		reply.Message = "The hashtable cleaned"
	} else {
		reply.Message = "Unable to clean the table"
	}

	return nil
}

func Server() {
	hashtable.Init(8)

	RPC := rpc.NewServer()
	xmlrpcCodec := xml.NewCodec()
	RPC.RegisterCodec(xmlrpcCodec, "text/xml")
	RPC.RegisterService(new(MessageService), "")
	http.Handle("/hashtable", RPC)

	log.Println("Starting XML-RPC Server on localhost:8081/hashtable")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func XmlRpcCall(method string, args struct{ Who string }) (reply struct{ Message string }, err error) {

	buf, _ := xml.EncodeClientRequest(method, &args)

	resp, err := http.Post("http://localhost:8081/hashtable", "text/xml", bytes.NewBuffer(buf))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = xml.DecodeClientResponse(resp.Body, &reply)
	return
}

func Client() {

	log.Println("Run client...")
	fmt.Println("Enter \"exit\" command for logout")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter the command >: ")
		cmd, _ := reader.ReadString('\n')

		if cmd == "exit\n" {
			break
		}

		executeCommand(cmd)

		/*
			reply, err := XmlRpcCall("MessageService." + foo, nil)
			if err != nil{
				log.Fatal(err)
			}

			log.Printf("Response: %s\n", reply.Message)
		*/
	}
	//time.Sleep(5000 * time.Millisecond)
	//reply, err := XmlRpcCall("MessageService.Reply", struct{ Who string }{"User 1"})
	//reply, err := XmlRpcCall("Service.Print", struct{ Who string }{"User 1"})
	/*
		reply, err := XmlRpcCall("MessageService.Print", struct{ Who string }{"User 1"})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response: %s\n", reply.Message)
	*/
}

func executeCommand(cmd string) {

	var foo string = ""
	ob := strings.Index(cmd, "(")
	cb := strings.Index(cmd, ")")

	if strings.HasPrefix(cmd, "clear") {
		foo = "clear"

		reply, err := XmlRpcCall("MessageService."+foo, struct{ Temp string }{"temp"})
		if err != nil {
			log.Fatal(err)
		}
	}

	if strings.HasPrefix(cmd, "containsKey") {
		foo = "containsKey"
		key := cmd[ob+1 : cb]
		//log.Println("Function: " + foo)
		//log.Println("Key is: " + key)

		reply, err := XmlRpcCall("MessageService."+foo, struct{ Key string }{key})
		if err != nil {
			log.Fatal(err)
		}
	}

	if strings.HasPrefix(cmd, "containsValue") {
		foo = "containsValue"
		value := cmd[ob+1 : cb]

		reply, err := XmlRpcCall("MessageService."+foo, struct{ Value string }{value})
		if err != nil {
			log.Fatal(err)
		}
	}

	if strings.HasPrefix(cmd, "get") {
		foo = "get"
		key := cmd[ob+1 : cb]

		reply, err := XmlRpcCall("MessageService."+foo, struct{ Key string }{key})
		if err != nil {
			log.Fatal(err)
		}
	}

	if strings.HasPrefix(cmd, "isEmpty") {
		foo = "isEmpty"

		reply, err := XmlRpcCall("MessageService."+foo, struct{ Temp string }{"temp"})
		if err != nil {
			log.Fatal(err)
		}
	}

	if strings.HasPrefix(cmd, "put") {
		foo = "put"

		comma := strings.Index(cmd, ",")
		key := cmd[ob+1 : comma]
		value := cmd[comma+2 : cb]

		var p hashtable.Pair

		reply, err := XmlRpcCall("MessageService."+foo, struct{ pair Pair }{p.New(key, value)})
		if err != nil {
			log.Fatal(err)
		}
	}

	if strings.HasPrefix(cmd, "remove") {
		foo = "remove"
		key := cmd[ob+1 : cb]

		reply, err := XmlRpcCall("MessageService."+foo, struct{ Key string }{key})
		if err != nil {
			log.Fatal(err)
		}
	}

	if strings.HasPrefix(cmd, "size") {
		foo = "size"

		reply, err := XmlRpcCall("MessageService."+foo, struct{ Temp string }{"temp"})
		if err != nil {
			log.Fatal(err)
		}
	}

	if strings.HasPrefix(cmd, "toString") {
		foo = "toString"

		reply, err := XmlRpcCall("MessageService."+foo, struct{ Temp string }{"temp"})
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Printf("Response: %s\n", reply.Message)

	foo = ""
}
