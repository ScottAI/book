package main

import (
	"./protocol"
	"github.com/gogo/protobuf/proto"
	"log"
)

func main()  {
	u := &protocol.UserInfo{
		Message: *proto.String("testInfo"),
		Length: *proto.Int32(10),
	}

	data,err := proto.Marshal(u)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newInfo := &protocol.UserInfo{}
	err = proto.Unmarshal(data,newInfo)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Println(newInfo.GetMessage())
}