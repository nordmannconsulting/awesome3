package awesome3

import (
	"fmt"
	"github.com/gorilla/mux"
)

func Test() {
	fmt.Println("main")
	fmt.Println(mux.NewRouter())
}
