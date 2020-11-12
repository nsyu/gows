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

	searchBox, err := session.FindElement(selenium.ByCSSSelector, "#home-header-searchbox")
	searchBox.SendKeys("女装市场")

	searchBtn, err := session.FindElement(selenium.ByCSSSelector, "button.single")
	searchBtn.Click()

	time.Sleep(5 * time.Second)

	productList, _ := session.FindElements(selenium.ByCSSSelector, ".card-container")
	for _, product := range productList {
		productDetail, err := product.Text()
		fmt.Println(productDetail)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("==================")
		productImgEle, err := product.FindElement(selenium.ByCSSSelector, "a")
		fmt.Println(productImgEle)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("==================")
		productImgHref, err := productImgEle.GetAttribute("href")
		fmt.Println(productImgHref)
		if err != nil {
			fmt.Println(err)
		}
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
