from curses import raw
import re
import time
import json
import datetime
import os
import driver
from venv import create
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.chrome.service import Service
from bs4 import BeautifulSoup

SBI_USER = os.getenv("sbi_user")
SBI_PASS = os.getenv("sbi_pass")
CSV_DIR = "/csv/"
LOGIN_URL = "https://site1.sbisec.co.jp/ETGate/"
PORT_URL = "https://site1.sbisec.co.jp/ETGate/?_ControlID=WPLETpfR001Control&_PageID=DefaultPID&_DataStoreID=DSWPLETpfR001Control&_ActionID=DefaultAID&getFlg=on"

# HTMLテーブルデータからCSVを作成
def createCSV(table_data):
    outputCSV = ""
    m = []
    tbody = table_data.find('tbody')
    trs = tbody.find_all('tr')
    for tr in trs:
        r = []
        for td in tr.find_all('td'):
            td_text_without_comma = td.text.replace(',', '')
            r.append(td_text_without_comma)
        m.append(r)
    for r in m:
        outputCSV += ','.join(r)

    return outputCSV

# 作成した文字列データ(CSV)を指定場所に書き込み
def writeCSV(rawoutputCSV):
    filename = str(datetime.date.today()) + '.csv'
    outputCSV = reshapeCSV(rawoutputCSV)
    with open(CSV_DIR + filename, mode='w') as f:
        f.write(outputCSV)

    print(outputCSV)

# 作成した文字列データから空行などを消してCSVフォーマットを整える
def reshapeCSV(rawoutputCSV):
    outputCSV = rawoutputCSV.replace(',\n', ',')
    return outputCSV

if __name__ == '__main__':
    print('Program start')

    # ブラウザ起動
    driver = driver.get_driver()
    print('Get driver')
     
    # ログインURLにアクセス
    driver.get(LOGIN_URL)
    element = driver.find_element(by = By.NAME, value = 'ACT_login')
    input_user_id = driver.find_element(by = By.NAME, value = 'user_id')
    input_user_id.send_keys(SBI_USER)
    input_user_password = driver.find_element(by = By.NAME, value = 'user_password')
    input_user_password.send_keys(SBI_PASS)

    # ログインボタンを押す
    driver.find_element(by = By.NAME, value = 'ACT_login').click()
    print('Login')
    time.sleep(5)

    # ポートフォリオページに移動
    driver.get(PORT_URL)
    print('Move portfolio page')
    time.sleep(5)

    soup = BeautifulSoup(driver.page_source, "html.parser")

    # ポートフォリオの１テーブル目を取得
    table_data = soup.find("table", bgcolor="#9fbf99", cellpadding="4", cellspacing="1", width="100%")
    
    fetch_data = createCSV(table_data)
    print('create CSV')
    
    writeCSV(fetch_data)
    print('write CSV')

    # ブラウザを閉じる
    driver.quit()
