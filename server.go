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
)


type MomentumServer struct {
	Addr string
	listener *net.Listener
	stat     Stat
	lock     *sync.Mutex
	funcs    map[string]interface{}
	working  chan bool
	logger   *log.Logger

}

func Create(addr string)*MomentumServer{
	moment := new(MomentumServer)
	moment.Addr = addr
	moment.listener = initTCPServer(addr)
	moment.stat = Stat{}
	moment.funcs = map[string]interface{}{}
	moment.logger = log.New(os.Stderr, "", log.LstdFlags)
	return moment
}

func(moment *MomentumServer) Start(){
	moment.logger.Print("Try to start server")
	moment.serverRunning()
}

func (moment* MomentumServer) RegisterFunc(title string, f interface{}) {
	err := moment.checkFunc(title, f)
	if err != nil {

	}
}

//IsRunning returns true if server is running and false in otherwise
func (moment* MomentumServer) IsRunning() bool {
	select {
	case <- moment.working:
		return true
	default:
		return false
	}
}

//SendMessage provides sending message to the server
func (moment *MomentumServer) SendMessage(msg string) {

}

func initTCPServer(addr string)* Listener{
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
	size, err := conn.Read(buff)
	if err != nil {
		log.Fatal(err)
	}
	json.UnMarshal(buff)

}

func (moment *MomentumServer) serverRunning() {
	for {
		if conn, err := moment.listener.Accept(); err != nil {
			moment.logger.Fatal(err)
		}

		go getRequest(conn)
	}
}

//helpful method 
func (moment *MomentumServer) checkFunc(title string, f interface{}) error {
	if title == "" {
		return errors.New("Title of function is empty")
	}

	item := reflect.ValueOf(f).Kind()
	if item != "func" {
		return errors.New("This type is not a function")
	}

	return nil	
}