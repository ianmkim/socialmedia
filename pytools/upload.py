import json
import os
from pymongo import MongoClient

# all the ethnicity columns to convert to a more logical format
ethnicity_columns = [
    "American Indian or Alaska Native",
    "Asian",
    "Black or African American",
    "Hispanic or Latino",
    "Middle Eastern",
    "Native Hawaiian or Other Pacific Islander",
    "White",
    "Other"
]

'''
takes in a dict and removes all the ethnicity colums
returns a single "race"
'''
def determine_race(item):
    ethnicity = ""
    for col in ethnicity_columns:
        if item[col] != "":
            ethnicity = item[col]
        del item[col]
    return ethnicity

f_user_data = open("DALI_Data.json", "r")
user_data = json.loads(f_user_data.read())
f_user_data.close()

client = MongoClient(os.getenv("CONNSTRING"))
mdb = client['socialmedia']
col = mdb['users']

for item in user_data:
    # add columns needed to go in the db
    item['likes'] = 0
    item['race'] = determine_race(item)

    # insert in db
    col.insert_one(res)





