package main

import (
	"triva/internal"
)

// @title Triva - API Docs
// @version 2.0
// @description An API Documentation for Triva
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
  app := internal.App{}
  app.Init()
}