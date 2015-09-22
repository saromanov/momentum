package momentum

import
(
	"github.com/golang/protobuf/proto"
	"log"
)

func Serialize(data interface{}) {
	data, err := proto.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
}


func Deserialize() {

}

