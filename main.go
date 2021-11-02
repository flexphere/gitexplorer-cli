package main

import (
	"fmt"
	"os"

	"prompt/options"

	"github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func main() {
	primaryOptions := options.NewList("./primary-options.json")
	secondaryOptions := options.NewMap("./secondary-options.json")
	tertiaryOptions := options.NewMap("./tertiary-options.json")

	primaryLabel, err := promptSelection("I want to", primaryOptions.GetLabels())
	errorHandler(err)

	primaryValue := primaryOptions.GetValue(primaryLabel)

	secondaryLabel, err := promptSelection(primaryLabel, secondaryOptions.GetLabels(primaryValue))
	errorHandler(err)

	secondaryValue := secondaryOptions.GetValue(primaryValue, secondaryLabel)

	if secondaryValue.HasUsage() {
		fmt.Println(secondaryValue.Usage)
		if secondaryValue.Note != "" {
			fmt.Println(secondaryValue.Note)
		}
		clipboard.WriteAll(secondaryValue.Usage)
		return
	}

	tertiaryLabel, err := promptSelection(secondaryLabel, tertiaryOptions.GetLabels(secondaryValue.Value))
	errorHandler(err)

	tertiaryValue := tertiaryOptions.GetValue(secondaryValue.Value, tertiaryLabel)
	fmt.Println(tertiaryValue.Usage)
	if tertiaryValue.Note != "" {
		fmt.Println(tertiaryValue.Note)
	}

	clipboard.WriteAll(tertiaryValue.Usage)
}

func promptSelection(prefix string, options []string) (string, error) {
	prompt := promptui.Select{
		Label: prefix,
		Items: options,
	}
	_, selected, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return selected, nil
}
