// main.go
package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	elliot := &Person{
		Name: "Elliot",
		Age:  24,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	// Marshal
	data, err := proto.Marshal(elliot)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	// Unmarshal
	newElliot := &Person{}
	err = proto.Unmarshal(data, newElliot)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newElliot.GetAge())
	fmt.Println(newElliot.GetName())
	fmt.Println(newElliot.GetSocialFollowers().GetTwitter())
	fmt.Println(newElliot.GetSocialFollowers().GetYoutube())
}
