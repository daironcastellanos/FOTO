import os
import time
import uuid
from urllib import request
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

BASE_URL = "https://www.bing.com/create?toWww=1&redig=30C04FC20A7A4BA08168CC7BB17EDFD6"
SAVE_PATH = "C:/Users/ericd/dev/school/Software_engineer\Freel\my-app/Go_Backend/generated"
CHROME_DRIVER_PATH = "C:/path/to/chromedriver.exe"

photo_prompts = [
"Create an abstract portrait using AI photography techniques",
"Generate an otherworldly landscape that defies reality",
"Combine elements of nature and technology to produce a surreal image",
"Design a futuristic cityscape unlike anything seen before",
"Incorporate elements of dream and fantasy into an AI-generated image",
"Create a space-themed scene with intricate details and vivid colors",
"Generate an underwater world filled with strange and wondrous creatures",
"Design a surrealist cityscape with twisted architecture and distorted perspectives",
"Produce a unique interpretation of a classic masterpiece using AI photography",
"Create an otherworldly scene that blurs the lines between science and magic."
]

products= [

photo_prompts
  
]

image_xpaths = [
    '//*[@id="mmComponent_images_1"]/ul[1]/li[1]/div/div/a/div/img',
    '//*[@id="mmComponent_images_1"]/ul[1]/li[2]/div/div/a/div/img',
    '//*[@id="mmComponent_images_1"]/ul[2]/li[1]/div/div/a/div/img',
    '//*[@id="mmComponent_images_1"]/ul[2]/li[2]/div/div/a/div/img'
]


def main():
    with webdriver.Chrome(executable_path=CHROME_DRIVER_PATH) as driver:
        for product_set in products:
            for product in product_set:
                driver.get(BASE_URL)

                try:
                    login_button = WebDriverWait(driver, 5).until(
                        EC.presence_of_element_located((By.XPATH, '/html/body/div/form[1]/div/div/div[2]/div[1]/div/div/div/div/div[1]/div[4]/div/div/div/div[4]/div/div/div/div[2]/a'))
                    )
                    login_button.click()
                    print("Logging in with Face...")
                except Exception as e:
                    print("No login required, continuing...")

                search_box = WebDriverWait(driver, 10).until(
                    EC.presence_of_element_located((By.XPATH, '//*[@id="sb_form_q"]'))
                )
                search_box.clear()
                search_box.send_keys(f"{product}")
                search_box.submit()

                for index, xpath in enumerate(image_xpaths, start=1):
                    print(f"Waiting for image {index} to be generated for {product}...")
                    try:
                        generated_image = WebDriverWait(driver, 30).until(
                            EC.presence_of_element_located((By.XPATH, xpath))
                        )
                        generated_image_url = generated_image.get_attribute("src")

                        response = request.urlopen(generated_image_url)
                        if response.status == 200:
                            print(f"Saving image {index} for {product}...")
                            file_name = f"{product}_{index}_{str(uuid.uuid4())}.jpg"
                            with open(os.path.join(SAVE_PATH, file_name), "wb") as file:
                                file.write(response.read())
                            print(f"Image {index} for {product} saved successfully as {file_name}.")
                        else:
                            print(f"Failed to download image {index} for {product} (status code {response.status}).")
                    except Exception as e:
                      print(f"Error while downloading image {index} for {product}: {e}")

if __name__ == "__main__":
  main()

