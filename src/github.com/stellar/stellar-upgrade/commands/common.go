package commands

import (
	"github.com/fatih/color"
)

const oldNetworkAlphabet = "gsphnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCr65jkm8oFqi1tuvAxyz"

var (
	Cyan  = color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	Green = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	Red   = color.New(color.FgRed).Add(color.Bold).SprintFunc()
)
