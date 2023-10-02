import requests
from bs4 import BeautifulSoup
import openpyxl
url = 'https://shopee.tw/'
response = requests.get(url)
soup = BeautifulSoup(response.text, 'html.parser')
product_names = soup.select('.shopee-item-card__text-name')
 # 建立 Excel 檔案
workbook = openpyxl.Workbook()
sheet = workbook.active
 # 將商品名稱寫入 Excel 檔案
for i, product_name in enumerate(product_names, start=1):
    sheet.cell(row=i, column=1, value=product_name.text.strip())
 # 儲存 Excel 檔案
workbook.save('product_names.xlsx')