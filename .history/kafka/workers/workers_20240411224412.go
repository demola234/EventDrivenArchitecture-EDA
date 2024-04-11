package main

func main() {

	topic := "comments"

	worker, err := connectConstumer([]string{"localhost:29093"})
	if err!= nil {
		panic(err)
	}
		

	consumer

}
