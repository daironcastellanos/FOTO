package scripts

import (
	"os"
	"time"

	"github.com/tebeka/selenium"
)

const (
	baseURL          = "https://www.bing.com/create?toWww=1&redig=30C04FC20A7A4BA08168CC7BB17EDFD6"
	savePath         = "C:/dev/personal/GPT CREATE/image/scripts/TAT"
	chromeDriverPath = "C:/path/to/chromedriver.exe"
)

var (
	scenic = []string{}

	products = [][]string{
		scenic,
	}

	imageXPaths = []string{
		`//*[@id="mmComponent_images_1"]/ul[1]/li[1]/div/div/a/div/img`,
		`//*[@id="mmComponent_images_1"]/ul[1]/li[2]/div/div/a/div/img`,
		`//*[@id="mmComponent_images_1"]/ul[2]/li[1]/div/div/a/div/img`,
		`//*[@id="mmComponent_images_1"]/ul[2]/li[2]/div/div/a/div/img`,
	}
)

func Bing() {
	service, err := selenium.NewChromeDriverService(chromeDriverPath, 9515, nil, nil)
	if err != nil {
		println("Error starting the ChromeDriver service:", err)
		return
	}
	defer service.Stop()

	webDriver, err := selenium.NewRemote(selenium.Capabilities{
		"browserName": "chrome",
	}, "http://localhost:9515/wd/hub")
	if err != nil {
		println("Error starting the web driver:", err)
		return
	}
	defer webDriver.Quit()

	for _, productSet := range products {
		for _, product := range productSet {
			err = webDriver.Get(baseURL)
			if err != nil {
				println("Error navigating to the base URL:", err)
				return
			}

			var loginButton selenium.WebElement
			err = webDriver.Wait(func(webDriver selenium.WebDriver) (bool, error) {
				var err error
				loginButton, err = webDriver.FindElement(selenium.ByXPATH, `/html/body/div/form[1]/div/div/div[2]/div[1]/div/div/div/div/div[1]/div[4]/div/div/div/div[4]/div/div/div/div[2]/a`)
				if err != nil {
					println("No login required, continuing...")
					return true, nil
				}
				return true, nil
			}, time.Second * 5)
			if err != nil {
			println("Error waiting for the login button:", err)
			return
			}		err = loginButton.Click()
			if err != nil {
				println("Error clicking the login button:", err)
				return
			}
			println("Logging in with Face...")
	
			var searchBox selenium.WebElement
			err = webDriver.Wait(func(webDriver selenium.WebDriver) (bool, error) {
				var err error
				searchBox, err = webDriver.FindElement(selenium.ByXPATH, `//*[@id="sb_form_q"]`)
				return err == nil, nil
			}, time.Second*10)
			if err != nil {
				println("Error waiting for the search box:", err)
				return
			}
	
			err = searchBox.Clear()
			if err != nil {
				println("Error clearing the search box:", err)
				return
			}
	
			err = searchBox.SendKeys(product)
			if err != nil {
				println("Error sending keys to the search box:", err)
				return
			}
	
			err = searchBox.Submit()
			if err != nil {
				println("Error submitting the search form:", err)
				return
			}
	
			productFolderName := productSet[0]
			productFolder := savePath + "/" + productFolderName
			err = os.MkdirAll(productFolder, 0755)
			if err != nil {
				println("Error creating the product folder:", err)
				return
			}
	
			for index, xpath := range imageXPaths {
				println(f"Waiting for image {index + 1} to be generated for {product}...")
				var generatedImage selenium.WebElement
				err = webDriver.Wait(func(webDriver selenium.WebDriver) (bool, error) {
					var err error
					generatedImage, err = webDriver.FindElement(selenium.ByXPATH, xpath)
					return err == nil, nil
				}, time.Second*30)
				if err != nil {
					println(f"Error waiting for the image {index + 1}:", err)
					continue
				}
	
				generatedImageURL, err := generatedImage.GetAttribute("src")
				if err != nil {
					println(f"Error getting the URL of the image {index + 1}:", err)
					continue
				}
				println(f""Saving image {:= os.Create(productFolder + "/" + generatedImageURL + ".jpg")
				if err != nil {
				println(f"Error creating the image file for image {index + 1}:", err)
				continue
				}
				defer file.Close()			_, err = file.WriteString(generatedImageURL)
				if err != nil {
					println(f"Error writing the image data for image {index + 1}:", err)
					continue
				}
	
				println(f"Image {index + 1} for {product} saved successfully.")
			}
		}
	}
	
	
