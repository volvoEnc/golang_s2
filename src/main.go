package main

import (
	"fmt"
	"test2/src/config"
)

func main() {
	fmt.Print(`Loading config...`)
	conf, err := config.NewConfig()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(conf.Database.TablePrefix)

}
