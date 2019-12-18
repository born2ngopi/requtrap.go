package action

import (
	"fmt"
	"github.com/libidev/requtrap.go/cli/config"
	"github.com/libidev/requtrap.go/http"
)

// Help Action
// Display help message
func Help() {
	fmt.Println(" ____                 _____")
	fmt.Println("|  _ \\ ___  __ _ _   |_   _| __ __ _ _ __")
	fmt.Println("| |_) / _ \\/ _` | | | || || '__/ _` | '_ \\")
	fmt.Println("|  _ <  __/ (_| | |_| || || | | (_| | |_) |")
	fmt.Println("|_| \\_\\___|\\__, |\\__,_||_||_|  \\__,_| .__/")
	fmt.Println("	      |_|                   |_|")
	fmt.Println("RequTrap - Fast and Configurable API Gateway")

	fmt.Println(`
Usage: 
	requtrap [command]

Commands: 
	start [config.yaml]	starting API Gateway
	stop               	stopping API Gateway
	help               	show this help message


Example:
	requtrap start config.yaml
	requtrap stop
`)
}

// Start Action
// Running API gateway server
func Start(conf *config.Yaml) {
	http.Serve(conf)
}
