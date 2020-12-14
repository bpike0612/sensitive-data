package main

import (
	"sensitive-data/controller"
	"sensitive-data/protocol"
)

func main() {
	_ = controller.Process(&protocol.Company{
		Id: 11,
		Owner: &protocol.User{
			Id:    1,
			Name:  "Batman",
			Email: "batman@cave.com",
			Title: &protocol.Title{
				Id:   100001,
				Name: "CLSO - Chief Life Savior Officer",
			},
		},
		CoOwner: &protocol.User{
			Id:    2,
			Name:  "Catwoman",
			Email: "catwoman@box.com",
			Title: &protocol.Title{
				Id:   100002,
				Name: "CCO - Chef Cuddling Officer",
			},
		},
		Size: 3,
	})
}
