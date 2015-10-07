package momentum

import
(
	"sync"
	//"time"
	"net"
	"log"
	"reflect"
	"errors"
	"os"
	"runtime"
	"encoding/json"
	"io/ioutil"
)

type Server interface {
	Start()
	Stop()
}

type MomentumServer struct {
	Addr string
	Listener net.Listener
	Stat     Stat
	Lock     *sync.Mutex
	Funcs    map[string]interface{}
	Working  chan bool
	Logger   *log.Logger

}

func Create(addr string)*MomentumServer{
	moment := new(MomentumServer)
	moment.Addr = addr
	moment.Listener = initTCPServer(addr)
	moment.Stat = Stat{}
	moment.Funcs = map[string]interface{}{}
	moment.Logger = log.New(os.Stderr, "", log.LstdFlags)
	return moment
}

func(moment *MomentumServer) Start(){
	moment.Logger.Print("Try to start server")
	moment.serverRunning()
}

func (moment *MomentumServer) Stop() {
	close(moment.Working)
}

func (moment* MomentumServer) RegisterFunc(title string, f interface{}) error {
	moment.Lock.Lock()
	defer moment.Lock.Unlock()
	err := moment.checkFunc(title, f)
	if err != nil {
		return err
	}

	moment.Funcs[title] = f
	return nil
}

//IsRunning returns true if server is running and false in otherwise
func (moment* MomentumServer) IsRunning() bool {
	select {
	case <- moment.Working:
		return true
	default:
		return false
	}
}

//SendMessage provides sending message to the server
func (moment *MomentumServer) SendMessage(msg string) {

}

//Serialize provides serialization of moment structure
func (moment *MomentumServer) Serialize(outfile string) {
	result, err := json.Marshal(moment)
	if err != nil {
		moment.Logger.Print("Serialize: Error to marshal moment data")
	}

	errwrite := ioutil.WriteFile(outfile, result, 0777)
	if errwrite != nil {
		moment.Logger.Print("Serialize: Error to write data")
	}

}

func initTCPServer(addr string) net.Listener{
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	return l
}

//Getting new request
func getRequest(conn net.Conn){
	buff := make([]byte,1024)
	defer conn.Close()
	_, err := conn.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	//json.Unmarshal(buff)

}

func (moment *MomentumServer) serverRunning() {
	for {
		conn, err := moment.Listener.Accept(); 
		if err != nil {
			moment.Logger.Fatal(err)
		}

		if !<-moment.Working {

		}
		go getRequest(conn)
	}
}

//helpful method 
func (moment *MomentumServer) checkFunc(title string, f interface{}) error {
	if title == "" {
		return errors.New("Title of function is empty")
	}

	//Checking that f is a function
	item := reflect.ValueOf(f).Kind()
	if item != reflect.Func {
		return errors.New("This type is not a function")
	}

	//Checking name of the function
	correctname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	if correctname == "" {
		return errors.New("Function contains invalid name")
	}

	return nil	
}

func (moment *MomentumServer) callMethod(title string){

	//TODO: At the end, send response to the server
}