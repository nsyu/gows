package main

import (
	"fmt"
	"log"

	"github.com/fedesog/webdriver"
	"github.com/tebeka/selenium"
)

func main() {
	chromeDriver := webdriver.NewChromeDriver("./chromedriver.exe")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Windows"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}
	err = session.Url("https://detail.1688.com/offer/627406028649.html?spm=a262eq.12572822.jspytrj4.1.24aa3d3apjDoIn")
	if err != nil {
		log.Println(err)
	}

	//session.Delete()
	//chromeDriver.Stop()

	//time.Sleep(5 * time.Second)
	//element, _ := session.FindElement(selenium.ByCSSSelector, "body").ButtonDown(27)

	// 셀레니움 관련 제어 부분
	productName, _ := session.FindElement(selenium.ByCSSSelector, ".d-title")
	pName, err := productName.Text()
	fmt.Println("product name : " + pName)
	fmt.Println("==================")

	productPrice, _ := session.FindElement(selenium.ByCSSSelector, ".price")
	pPrice, err := productPrice.Text()
	fmt.Println("product price : " + pPrice)
	fmt.Println("==================")

	productOptionName, _ := session.FindElements(selenium.ByCSSSelector, "tr > td.name")
	for _, ele := range productOptionName {
		pOtionName, err := ele.Text()
		if err != nil {
			panic(err)
		}
		fmt.Println("product option : " + pOtionName)
	}
	//log.Println(element)
	//log.Println(1 * time.Second)

	//time.Sleep(1 * time.Second)
	//session.Delete()
	//chromeDriver.Stop()
}
