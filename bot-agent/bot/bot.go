package bot

import (
	"fmt"
	"time"
)

const Exception string = "No such command!"
const ExceptionLanguage string = "This language does not exist."
const InputError string = "Input Error"

type bots struct {
	RussianBot
	EnglishBot
}

type Bot interface {
	Pronouns()
	PrintName()
	FirstMethod()
	SecondMethod()
	ThirdMethod()
	FourMethod()
	FiveMethod()
}
type RussianBot struct {
	Name string
}
type EnglishBot struct {
	Name string
}

func (rb RussianBot) Pronouns() {
	fmt.Print("Вы: ")
}
func (eb EnglishBot) Pronouns() {
	fmt.Print("You: ")
}

func (rb RussianBot) PrintName() {
	fmt.Print(rb.Name + ": ")
}
func (eb EnglishBot) PrintName() {
	fmt.Print(eb.Name + ": ")
}

func ReadingText(text string, language string) string {
	switch language {
	case "Russian":
		russianMap := map[string]string{
			"Привет": "1",
			"Время":  "2",
			"Дата":   "3",
			"День":   "4",
			"Пока":   "5",
			"1":      "1",
			"2":      "2",
			"3":      "3",
			"4":      "4",
			"5":      "5",
		}
		if val, ok := russianMap[text]; ok {
			return val
		}
	case "English":
		englishMap := map[string]string{
			"Hello": "1",
			"Time":  "2",
			"Date":  "3",
			"Day":   "4",
			"Bye":   "5",
			"1":     "1",
			"2":     "2",
			"3":     "3",
			"4":     "4",
			"5":     "5",
		}
		if val, ok := englishMap[text]; ok {
			return val
		}
	default:
		fmt.Println(Exception)
	}
	return ""
}

func (rb RussianBot) FirstMethod() {
	fmt.Println("Привет, я ", rb.Name)
}

func (eb EnglishBot) FirstMethod() {
	fmt.Println("Hello, I am ", eb.Name)
}

func getTime(name string) time.Time {
	t := time.Now()
	loc, _ := time.LoadLocation(name)
	t = t.In(loc)
	return t
}
func (rb RussianBot) SecondMethod() {
	name := "Europe/Minsk"
	fmt.Println("Сейчас ", getTime(name).Format("15:04:00"))
}

func (eb EnglishBot) SecondMethod() {
	name := "Europe/London"
	fmt.Println("Now is ", getTime(name).Format("15:04:00"))
}

func getDate() string {
	t := time.Time(time.Now())
	return t.Format("January 2,2006")
}

func (rb RussianBot) ThirdMethod() {
	fmt.Println("Сегодня ", getDate())
}

func (eb EnglishBot) ThirdMethod() {
	fmt.Println("Today is ", getDate())
}

func getWeekday() string {
	t := time.Time(time.Now())
	return t.Format("Monday")
}

func (rb RussianBot) FourMethod() {
	fmt.Println("Сегодня ", getWeekday())
}

func (eb EnglishBot) FourMethod() {
	fmt.Println("Today is ", getWeekday())
}

func (rb RussianBot) FiveMethod() {
	fmt.Println("Пока")
}

func (eb EnglishBot) FiveMethod() {
	fmt.Println("Bye")
}
