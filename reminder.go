package main

import (
	"fmt"
	"github.com/clagraff/argparse"
	"os"
)

func main() {
	/*
	 * Command format:
	 * $reminder [-a <REMIND_TEXT>|-l|-r <REMIND_ID>]
	 *
	 */

	parser := argparse.NewParser("A general purpose CLI TODO list.").Version("0.1a")
	parser.AddHelp().AddVersion()

	add := argparse.NewOption("a", "add", "-a <REMINDER_TEXT>").Default("")
	list := argparse.NewOption("l", "list", "-l will list all current reminders")
	rem := argparse.NewOption("r", "remove", "-r <REMINDER_ID>").Default("-1")

	parser.AddOptions(add, list, rem)

	ns, s, err := parser.Parse(os.Args[1:]...)

	switch err.(type) {
	case argparse.ShowHelpErr:
		return
	case error:
		fmt.Println(err, '\n')
		parser.ShowHelp()
		return
	}
	fmt.Println(ns.String)
	fmt.Println(s)
}
