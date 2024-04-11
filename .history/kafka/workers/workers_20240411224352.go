package main

func main() {

	topic := "comments"

	worker, err := workerconnectConstumer([]string{"localhost:29093"})
	if err != nil {
		panic(err)
	}

}
