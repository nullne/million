import requests


form_data = {
    'username': 'le.yu',
    'password': '1234',
    'daterange': '2015-10-01 - 2015-12-01'
}

login = requests.post('http://192.168.15.211:8910/', data=form_data)
print login.content
