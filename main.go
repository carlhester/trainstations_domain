package main

//import "fmt"
//import "trainstations_domain/stations"
//import "trainstations_domain/storage/file"
import web "trainstations_domain/server/web"

func main() {

	web.StartServer(":8080")
}
