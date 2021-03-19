package main

import (
	"fmt"

	"github.com/woodsmur/go-mod/greeting/api"
	"github.com/woodsmur/go-mod/greeting/hello"

	v2api "github.com/woodsmur/go-mod/v2/greeting/api"
	v2hello "github.com/woodsmur/go-mod/v2/greeting/hello"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(api.HelloApi())

	fmt.Println(v2api.API())
	fmt.Println(v2hello.Hello())
}
