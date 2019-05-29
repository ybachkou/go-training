package main

import (
	"bot/bot"
	"errors"
	"fmt"
	"strings"
)

func main() {
	var language string
	fmt.Print("Language: ")
	_, err := fmt.Scanln(&language)
	if err != nil {
		fmt.Println(errors.New(bot.InputError))
	}
	if strings.EqualFold(language, "Russian") || strings.EqualFold(language, "English") {
		var text string
		exit := true
		for exit {
			createBot(language).Pronouns()
			_, err := fmt.Scanln(&text)
			if err != nil {
				fmt.Println(errors.New(bot.InputError))
			}
			wordsForLanguage := bot.ReadingText(text, language)
			createBot(language).PrintName()
			switch wordsForLanguage {
			case "1":
				createBot(language).FirstMethod()
			case "2":
				createBot(language).SecondMethod()
			case "3":
				createBot(language).ThirdMethod()
			case "4":
				createBot(language).FourMethod()
			case "5":
				createBot(language).FiveMethod()
				exit = false
			default:
				fmt.Println(bot.Exception)
			}
		}
	} else {
		fmt.Println(errors.New(bot.ExceptionLanguage))
	}
}

func createBot(language string) bot.Bot {
	switch language {
	case "Russian":
		r := bot.RussianBot{"Иван"}
		return r
	case "English":
		e := bot.EnglishBot{"John"}
		return e
	default:
		fmt.Println(bot.ExceptionLanguage)
		return nil
	}
}
