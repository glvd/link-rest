package main

func main() {
	rest, err := linkrest.New(18080)
	if err != nil {
		panic(err)
	}
	rest.Start()
}
