package main

import (
    "log"
    //"CC-web/src/api-gateway/handler"
)

func main() {
    router := SetupRouter()

    log.Println("API Gateway is running on port 8080...")
    router.Run(":8080")
}
