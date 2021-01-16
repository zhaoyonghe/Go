import os

def check_empty(d):
    if not os.listdir(d):
        print(d)

for d in os.listdir("."):
    # print(type(d))
    # print(d)
    if os.path.isdir(d):
        check_empty(d)