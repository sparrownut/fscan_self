package utils

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func Printsuc(text string, args ...any) string {
	c := color.New(color.FgHiGreen, color.Bold)
	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[+]")
		if err != nil {
		}
	}
	fmt.Printf(text+"\n", args...)
	return "[+]" + fmt.Sprintf(text+"\n", args...)

}
func Printerr(text string, args ...any) string {
	c := color.New(color.FgHiRed)
	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[-]")
		if err != nil {
		}
	}
	fmt.Printf(text+"\n", args...)
	return "[-]" + fmt.Sprintf(text+"\n", args...)
}
func Printminfo(text string, args ...any) string {
	c := color.New(color.FgYellow)
	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[>]")
		if err != nil {
		}
	}
	fmt.Printf(text+"\n", args...)
	return "[>]" + fmt.Sprintf(text+"\n", args...)
}
func Printhinfo(text string, args ...any) string {
	c := color.New(color.FgHiYellow, color.Bold)

	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[!]")
		if err != nil {
		}
	}
	fmt.Printf(text+"\n", args...)
	return "[!]" + fmt.Sprintf(text+"\n", args...)
}
func Printcritical(text string, args ...any) {
	c := color.New(color.FgHiBlue, color.BgHiRed, color.Bold)
	if !strings.Contains(text, "-") && !strings.Contains(text, ">") && !strings.Contains(text, "!") && !strings.Contains(text, "+") {
		_, err := c.Print("[■]")

		if err != nil {
			return
		}

	}
	_, err := c.Printf(text+"\n", args...)

	if err != nil {
		return
	}
	fmt.Print()
}
