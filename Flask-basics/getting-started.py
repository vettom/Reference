from datetime import datetime
from flask import Flask

# export FLASK_APP=index.py
# export FLASK_ENV=development
# flask run
app = Flask(__name__)
Count = 0


@app.route("/")
def home():
    return "This is Home"


@app.route('/dv')
def dv():
    global Count
    Count += 1
    return f"This id for dv {str(datetime.now())}  Count = {Count} <hr>"

# Below is to display 404 Page not found if not valid URL
@app.route("/<PGNotFound>/")
def pagenotfound(PGNotFound):
    return "Are you lost?"


''' Adding static folder and placing files in it will be loadable as /static/file.gif '''
