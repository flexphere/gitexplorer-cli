package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"github.com/flexphere/gitexplorer-cli/options"
	"github.com/manifoldco/promptui"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func main() {
	primaryOptions := options.NewList()
	secondaryOptions := options.NewMap("secondary")
	tertiaryOptions := options.NewMap("tertiary")

	primaryLabel, err := PromptSelection("I want to", primaryOptions.GetLabels())
	errorHandler(err)
	primaryValue := primaryOptions.GetValue(primaryLabel)

	secondaryLabel, err := PromptSelection(primaryLabel, secondaryOptions.GetLabels(primaryValue))
	errorHandler(err)
	secondaryValue := secondaryOptions.GetValue(primaryValue, secondaryLabel)
	if secondaryValue.HasUsage() {
		PrintAndExit(secondaryValue)
		return
	}

	tertiaryLabel, err := PromptSelection(secondaryLabel, tertiaryOptions.GetLabels(secondaryValue.Value))
	errorHandler(err)

	tertiaryValue := tertiaryOptions.GetValue(secondaryValue.Value, tertiaryLabel)
	PrintAndExit(tertiaryValue)
}

func PrintAndExit(result options.Option) {
	fmt.Println("\n" + promptui.Styler(promptui.FGGreen)(result.Usage) + "\n")
	if result.Note != "" {
		fmt.Println(result.Note)
	}

	fmt.Printf("\nCopied to clipboard!\n")
	clipboard.WriteAll(result.Usage)
}

func PromptSelection(prefix string, options []string) (string, error) {
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
