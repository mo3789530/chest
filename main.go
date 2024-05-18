/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/mo3789530/chest/cmd"
	logger "github.com/mo3789530/chest/internal/pkg/logger"
)

func main() {
	logger.NewLogger()
	cmd.Execute()
}
