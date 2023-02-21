import json
def load_db():
    try:
        with open ("db.json") as F:
            return json.load(F)
    except Exception as E:
        print(f" Failed to open DB file. \ {E}")

def save_db():
    with open("db.json", 'w') as X:
        return json.dump(DB, X)
DB = load_db()

