from flask import Flask,render_template,abort,jsonify,request, redirect,url_for

# python3 -m flask run
app = Flask(__name__)



@app.route("/")
def home():
    return ("Flask worked")
