package main

import (
	"fmt"

	"github.com/Luc1808/go-task-api/internal/db"
)

func main() {
	db.InitDB()

	fmt.Println("Database connected successfully!")
}
