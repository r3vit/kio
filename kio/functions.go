package kio

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

//Choose(id, ...options)
func (kio *Kio) templateChoose(values ...string) string {
	if kio.reference[values[0]] != "" {
		return kio.reference[values[0]]
	}

	var option int

	color.Set(color.FgYellow)
	fmt.Println("Choose an option for", "\"", values[0], "\"")
	color.Unset()

	for i, value := range values {
		if i != 0 {
			fmt.Println("[", i, "]", value)
		}
	}
	fmt.Print("You choose option: ")
	fmt.Scanf("%d", &option)

	for i, value := range values {
		if i == option {
			kio.reference[values[0]] = value
			return value
		}
	}
	return "error"

}

//IsSet(section name) for optional
func (kio *Kio) templateIsSet(name string) bool {
	if kio.reference[name] != "" {
		fmt.Println("----", kio.reference[name])
		return true
	}

	var str string
	color.Set(color.FgYellow)
	fmt.Println("Do you want optional ", name, "?")
	color.Unset()
	fmt.Print("No/Yes: ")

	fmt.Scanf("%s", &str)

	str = strings.ToLower(str)

	if str == "y" || str == "yes" || str == "ye" {
		kio.reference[name] = "yes"
		return true
	}

	kio.reference[name] = "no"
	return false
}

//custom(id)
func (kio *Kio) templateCustom(name string) string {
	if kio.reference[name] != "" {
		return kio.reference[name]
	}

	var str string
	color.Set(color.FgYellow)
	fmt.Print("Custom text for ", name, ": ")
	color.Unset()

	fmt.Scanf("%s", &str)

	//save data in reference
	kio.reference[name] = str

	return str
}
