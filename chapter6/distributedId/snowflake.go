package distributedId

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

// https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch6-cloud/ch6-01-dist-id.html
func UseSnowFlake() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println("node: ", id.Node(), "step: ", id.Step(), "time: ", id.Time())
	}
}
