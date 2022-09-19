import sys
import base64
import json
import requests
import subprocess
from requests.auth import HTTPBasicAuth

def main():
    print("test start")

    # 1-1 health check
    url = "http://localhost:8080/"
    res = requests.get(url)
    if res.status_code != 200:
        print("1-1. health check failed")
        sys.exit(1)
    print("1-1. health check ok")

    # 2-1 regist new record
    url = "http://localhost:8080/register/20060102"
    with open("test/20060102.csv") as f:
        cmd = 'curl -sXPOST --data-binary @test/20060102.csv http://localhost:8080/regist/20060102'
        process = (subprocess.Popen(cmd, stdout=subprocess.PIPE, shell=True).communicate()[0]).decode('utf-8')
        if process != '{"created_number":5}':
            print("2-1. regist new record failed")
            print(process)
            sys.exit(1)
    print("2-1. regist new record ok")

    # 2-2 regist already exists
    url = "http://localhost:8080/register/20060102"
    with open("test/20060102.csv") as f:
        cmd = 'curl -sXPOST --data-binary @test/20060102.csv http://localhost:8080/regist/20060102'
        process = (subprocess.Popen(cmd, stdout=subprocess.PIPE, shell=True).communicate()[0]).decode('utf-8')
        if process != '{"skipped_number":5}':
            print("2-2. regist already exists failed")
            print(process)
            sys.exit(1)
    print("2-2. regist already exists ok")
    
if __name__ == "__main__":
    main()
