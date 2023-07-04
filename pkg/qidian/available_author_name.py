import time

import requests

if __name__ == '__main__':
    base_url = "https://write.qq.com/ccauthorweb/author/ismayeditauthorname?authorName={}&t={}"
    resp = requests.get(base_url.format("xx", time.time_ns()))
    if resp.status_code / 200 == 1:
        print(resp.json())
