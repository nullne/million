from datetime import datetime
from bs4 import BeautifulSoup

WORK_HOURS = 11


def overtime(work_hours=WORK_HOURS):
    data = open("./data.html")
    soup = BeautifulSoup(data, "html.parser")
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
            if start is not None and end is not None:
                delta = datetime.strptime(end, "%H:%M") - datetime.strptime(start, "%H:%M")
                if delta.seconds / 3600 >=  work_hours:
                    date.append([tdate + " " + start, tdate + " " +  end])

        count += 1
        tr = tr.findNextSibling()
        if tr is None:
            break
    return date

print overtime()
