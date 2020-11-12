package main

import (
	"fmt"
	"time"

	"github.com/fedesog/webdriver"
	"github.com/tebeka/selenium"
)

func main() {
	chromeDriver := webdriver.NewChromeDriver("./chromedriver.exe")
	err := chromeDriver.Start()
	if err != nil {
		fmt.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Windows"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		fmt.Println(err)
	}
	err = session.Url("https://1688.com/")
	if err != nil {
		fmt.Println(err)
	}

	category, _ := session.FindElements(selenium.ByCSSSelector, ".c-menu > p > a")
	for _, categoryElement := range category {
		categoryElementHref, err := categoryElement.GetAttribute("href")
		if err != nil {
			panic(err)
		}

		// NEW TAB
		tab, err := chromeDriver.NewSession(desired, required)
		if err != nil {
			fmt.Println(err)
		}
		err = tab.Url(categoryElementHref)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(60 * time.Second)
		fmt.Println(categoryElementHref)
		fmt.Println("==================")
	}

	// elements, _ := session.FindElements(selenium.ByCSSSelector, ".list > div > a > div.offer-img > img")
	// for _, we := range elements {
	// 	text, err := we.GetAttribute("src")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(text)
	// 	fmt.Println("==================")
	// }

	// time.Sleep(1 * time.Second)
	// session.Delete()
	// chromeDriver.Stop()
}
