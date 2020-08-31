package main

import "bookReadingNote/project/CircuitAndHystrix/example/circuitExample/circuitManager"

func main() {
	//body, err := circuitManager.CirManager.Get("http://www.google.com/robots.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(body))
	circuitManager.CirManager.FakeMany()
}
