import csv
import sys

with open('names.csv', 'w') as csvfile:
    writer = csv.writer(sys.stdout)
    # writer = csv.DictWriter(csvfile, fieldnames=fieldnames)

    # writer.writeheader()
    writer.writerow({'head': "ufck", 'first_name': 'Lovely', 'last_name': 'Spam'})
    # writer.writerow({'first_name': 'Wonderful', 'last_name': 'Spam'})
    writer.writerow([1, 2, 3], (1,2,3))
    writer.writerow((1, 2, 3))
    writer.writerow((1, 2, 3))
