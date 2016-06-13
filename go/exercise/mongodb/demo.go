package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	// "log"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("223.202.75.76")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("log-test")
	file, err := db.GridFS("results").OpenId(bson.ObjectIdHex("56e662296e5cfba8e9485b75"))
	if err != nil {
		fmt.Println(err)
		return
		// panic(err)
	}
	fmt.Println(file.Size())

	b := make([]byte, 8192)
	n, err := file.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "\t", string(b))
	err = file.Close()
	if err != nil {
		panic(err)
	}

	/*
	file, err := db.GridFS("results").Create(bson.NewObjectId().Hex())
	file.SetId(bson.ObjectIdHex("56d79ded464f185df1de73e8"))
	fmt.Println(file.Id())
	if err != nil {
		panic(err)
	}
	n, err := file.Write([]byte("fuck you any more"))
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	*/

	/*
		c := session.DB("tests").C("restaurount")
		fmt.Println(c)
		err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
			&Person{"Cla", "+55 53 8402 8510"})
		if err != nil {
			log.Fatal(err)
		}

		result := Person{}
		err = c.Find(bson.M{"name": "Ale"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Phone:", result.Phone)
	*/
}
