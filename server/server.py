#!/bin/python3

from database import Db

db = Db("test")
# print(db.insertPayment("2025-08-12 14:06:33", "qAsti", "Paninaro", "Postepay"))
db.insertDetail("Saccottino al cioccolato", 2, 1, 1000)
print(db.getCity())
print(db.getItem())
print(db.getShop())
print(db.getMethod())
print(db.getPayment())
print(db.getDetail())
print()
for line in db.getFullDetail():
    print(line)
