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

/*
func (ms *MessageService) Print(r *http.Request, args *rpcArgs, reply *rpcReply) error {
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
	hashtable.Put(p.New(8, "eigth"))

	str := hashtable.ToString()
	log.Println("Hashtable: " + str)
	reply.Message = "Hello, " + args.Value + "!"
	return nil
}*/

func (ms *MessageService) Clear(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC clear")

	isClear := strconv.FormatBool(hashtable.Clear())
	log.Println("Is clear hashtable: " + isClear)
	reply.Message = isClear

	return nil
}

func (ms *MessageService) Put(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC put")

	var p hashtable.Pair
	val := hashtable.Put(p.New(args.Key, args.Value))

	log.Println("Is put to hashtable, value: " + val)

	reply.Message = val

	return nil
}

func (ms *MessageService) ToString(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC toString")

	str := hashtable.ToString()

	log.Println("Is toString hashtable: " + str)

	reply.Message = str

	return nil
}

func (ms *MessageService) ContainsKey(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC containsKey")

	isContains := strconv.FormatBool(hashtable.ContainsKey(args.Key))

	log.Println("Is containsKey in hashtable: " + isContains)

	reply.Message = isContains

	return nil
}

func (ms *MessageService) ContainsValue(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC containsValue")

	isContains := strconv.FormatBool(hashtable.ContainsValue(args.Value))

	log.Println("Is containsValue in hashtable: " + isContains)

	reply.Message = isContains

	return nil
}

func (ms *MessageService) Get(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC get")

	val := hashtable.Get(args.Key)

	log.Println("Is get from hashtable, value: " + val)

	reply.Message = val

	return nil
}

func (ms *MessageService) IsEmpty(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC isEmpty")

	isEmpty := strconv.FormatBool(hashtable.IsEmpty())
	log.Println("Is empty hashtable: " + isEmpty)
	reply.Message = isEmpty

	return nil
}

func (ms *MessageService) Remove(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC remove")

	val := hashtable.Remove(args.Key)

	log.Println("Is remove from hashtable, value: " + val)

	reply.Message = val

	return nil
}

func (ms *MessageService) Size(r *http.Request, args *struct {
	Key   int
	Value string
}, reply *struct{ Message string }) error {
	log.Println("Call RPC size")

	size := strconv.Itoa(hashtable.Size())
	log.Println("Size of hashtable: " + size)
	reply.Message = size

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

/*
type rpcArgs struct {
	Key   int
	Value string
}

type rpcReply struct {
	Message string
}*/

func XmlRpcCall(method string, args struct {
	Key   int
	Value string
},) (reply struct{ Message string }, err error) {

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

		//cmd = cmd[:len(cmd)-1]

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

	log.Println("Call executeCommand")
	//var args rpcArgs
	//var reply rpcReply
	ob := strings.Index(cmd, "(")
	cb := strings.Index(cmd, ")")

	if strings.HasPrefix(cmd, "clear") {
		foo := "Clear"

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{0, "null"})

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response by \"clear\": %s\n", reply.Message)
	}

	if strings.HasPrefix(cmd, "containsKey") {
		foo := "ContainsKey"

		key, _ := strconv.Atoi(cmd[ob+1 : cb])

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{key, "null"})

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response by \"containsKey\": %s\n", reply.Message)
	}

	if strings.HasPrefix(cmd, "containsValue") {
		foo := "ContainsValue"

		value := cmd[ob+1 : cb]

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{0, value})

		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response by \"containsValue\": %s\n", reply.Message)
	}

	if strings.HasPrefix(cmd, "get") {
		foo := "Get"
		key, _ := strconv.Atoi(cmd[ob+1 : cb])

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{key, "null"})

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response by \"get\": %s\n", reply.Message)
	}

	if strings.HasPrefix(cmd, "isEmpty") {
		foo := "IsEmpty"

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{0, "null"})

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response by \"isEmpty\": %s\n", reply.Message)
	}

	if strings.HasPrefix(cmd, "put") {
		foo := "Put"

		comma := strings.Index(cmd, ",")
		key, _ := strconv.Atoi(cmd[ob+1 : comma])
		value := cmd[comma+2 : cb]

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{key, value})

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response by \"put\": %s\n", reply.Message)
	}

	if strings.HasPrefix(cmd, "remove") {
		foo := "Remove"
		key, _ := strconv.Atoi(cmd[ob+1 : cb])

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{key, "null"})

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response by \"remove\": %s\n", reply.Message)
	}

	if strings.HasPrefix(cmd, "size") {
		foo := "Size"

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{0, "null"})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response by \"size\": %s\n", reply.Message)
	}

	if strings.HasPrefix(cmd, "toString") {
		foo := "ToString"

		reply, err := XmlRpcCall("MessageService."+foo, struct {
			Key   int
			Value string
		}{0, "null"})

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Response by \"toString\": %s\n", reply.Message)
	}
}
