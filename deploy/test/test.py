import requests
import time
import multiprocessing
import os
import json
import math
import random

if __name__ == "__main__":
    # 为线程定义一个函数



    def print_time():
        count = 0

        appnames = ["wechat","qq","alipay"]
        msgtypes = ["login","logout","register"]

        while True:
            parms = [{
                "Appname":random.choice(appnames),
                "Msgtype":random.choice(msgtypes),
                "Timestamp": math.floor(time.time()) + random.randint(-15,15),
                "Content":{
                    "test1":123,
                    "test2":"456"
                },
                "count":count,
            }]

            if random.randint(0,9) < -3:
                parms.append({
                "Appname":random.choice(appnames),
                "Msgtype":random.choice(msgtypes),
                "Timestamp": math.floor(time.time()) - random.randint(-12*3600, 12*2600),
                "Content":{
                    "test1":123,
                    "test2":"456"
                },
                "count":count,
            })

            print(json.dumps(parms))
            r = requests.post(url="http://127.0.0.1:8888/", data=json.dumps(parms))
            print(r, r.text, count)
            count += 1
            time.sleep(1)


    # 创建两个线程
    print('Parent process {0} is Running'.format(os.getpid()))
    for i in range(5):
        p = multiprocessing.Process(target=print_time)
        print('process start')
        p.start()
        p.join()
    print('Process close')

