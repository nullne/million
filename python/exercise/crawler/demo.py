#! /usr/bin/env python
# coding=utf-8

__author__ = 'le.yu'


from datetime import datetime
from bs4 import BeautifulSoup
import requests
from bottle import post, run, request
import json

WORK_HOURS = 11
WEEKEND_HOURS = 4


def parse_overtime(content, fare, reason, etc):
    soup = BeautifulSoup(content, "html.parser")
    table = soup.find(id="TBL_TRecord")
    tr = table.find("tr")
    date = []
    count = 0
    while True:
        if count == 0:
            pass
        else:
            tr.find_all("td")
            td = tr.find_all("td")
            tdate, start, end = td[4].string, td[5].string, td[6].string
            week = datetime.strptime(tdate, "%Y-%m-%d").weekday()
            fire = False
            if start is not None and end is not None:
                fire = True
            elif etc:
                if start is None and end is not None:
                    fire = True
                    start = "10:00"
                elif start is not None and end is None:
                    fire = True
                    end = "18:00"
            if fire:
                tstart = datetime.strptime(start, "%H:%M")
                tend = datetime.strptime(end, "%H:%M")
                delta = (tend - tstart).total_seconds() / 3600.0
                if week < 5:
                    if delta >=  WORK_HOURS:
                        date.append([tdate + " " + start, tdate + " " +  end, "%.1f" % delta, fare, reason, "日常加班"])
                else:
                    if delta >=  WEEKEND_HOURS:
                        date.append([tdate + " " + start, tdate + " " +  end, "%.1f" % delta, fare, reason, "周末加班"])

        count += 1
        tr = tr.findNextSibling()
        if tr is None:
            break
    return date


def crawler(username, password, daterange):
    date = daterange.split(" - ")
    login_data = {
        'username': username,
        'password': password}
    login = requests.post('http://att.chinacache.com/Login.action', data=login_data)
    cookies = "JSESSIONID=" + login.cookies['JSESSIONID']
    query_data = {
        'f_empinfofullname': '',
        'f_orgid_op': '',
        'f_orgid': -1,
        'f_fdate': date[0],
        'f_fdate_2': date[1]
    }
    fetch_list = requests.post('http://att.chinacache.com/TRecordList.action', cookies=login.cookies, data=query_data)
    return fetch_list.content


def test():
    content = crawler("le.yu", "1234", "2015-10-01 - 2015-12-01")
    overtime = parse_overtime(content)
    print overtime

@post('/')
def main():
    username = request.forms.get('username')
    password = request.forms.get('password')
    daterange = request.forms.get('daterange')
    fare = request.forms.get("fare")
    reason = request.forms.get("reason")
    etc = request.forms.get("etc")

    content = crawler(username, password, daterange)
    overtime = parse_overtime(content, fare, reason, etc)
    return json.dumps(overtime)




if __name__ == "__main__":
    # test()
    # main()
    run(host='127.0.0.1', port=8080)

