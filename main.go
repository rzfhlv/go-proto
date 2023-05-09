package main

import (
	"fmt"
	pb "go-proto/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	products := &pb.Products{
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Blue Flanel",
				Price: 100000.00,
				Stock: 10,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    2,
				Name:  "Nike Air Jordan",
				Price: 2000000.00,
				Stock: 20,
				Category: &pb.Category{
					Id:   2,
					Name: "Shoes",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("marshall error: ", err)
	}

	fmt.Println("products: ", products)
	fmt.Println("data: ", data)

	testProducts := &pb.Products{}
	if err = proto.Unmarshal(data, testProducts); err != nil {
		log.Fatal("unmarshall error: ", err)
	}

	fmt.Println("TestProducts: ", testProducts)

	for _, product := range testProducts.GetData() {
		fmt.Println("name: ", product.GetName())
		fmt.Println("name string: ", product.Name)
		fmt.Println("price: ", product.Price)
	}
}
