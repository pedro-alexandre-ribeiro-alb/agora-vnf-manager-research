package main

import (
	docs "agora-vnf-manager/docs"
	server "agora-vnf-manager/server"
)

// @title AGORA - VNF Manager
// @version 1.0.0
// @description Documentation of the AGORA - VNF Manager API

// @contact.name Andre Brizido; Pedro Ribeiro
// @contact.email andre-d-brizido@alticelabs.com; pedro-alexandre-ribeiro@alticelabs.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @port 3000
func main() {
	docs.SwaggerInfo.Host = "localhost:3000"
	server.Init()
}
