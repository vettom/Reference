from flask import Flask,render_template,abort,jsonify,request, redirect,url_for
from dbload import DB,save_db

# python3 -m flask run
app = Flask(__name__)



@app.route("/")
def home():
    return  render_template("welcome.html", Data="Server list", Index=0, DB=DB)

@app.route("/form/", methods=["GET", "POST"])
def form():
    if request.method == "POST":
        #Get data to a dict and add to DB
        New_SRV = { "Server": request.form["ServerName"], "IP": request.form["IPA"]  }
        DB.append(New_SRV)
        save_db()
        return redirect(url_for("server_view", Index=len(DB)-1))
    else:
        return render_template("forms.html")

@app.route("/delete/<int:Index>", methods=["GET", "POST"])
def delete(Index):
    if request.method == "POST":
    #     Delete
        del DB[Index]
        save_db()
        return redirect(url_for("server_view", Index=Index-1))
    else:
        try:
            Srv = DB[Index]["Server"]
            IP = DB[Index]["IP"]
            return render_template("delete.html", Data="Delete Server", Srv=Srv, IP=IP, Index=Index)
        except:
            abort(404)

@app.route("/new/<int:Index>")
def server_view(Index):
    try:
        Srv = DB[Index]["Server"]
        IP = DB[Index]["IP"]
        return render_template("view.html", Data="Servers Info", Srv=Srv, IP=IP, Index=Index)
    except IndexError:
        Index=0
        Srv = DB[Index]["Server"]
        IP = DB[Index]["IP"]
        return render_template("view.html", Data="Index Error, loading first index", Srv=Srv, IP=IP, Index=Index)
    except:
        abort(404)

@app.route("/api/new/")
def api_full():
    return jsonify(DB)

@app.route("/test/")
def test():
    return render_template("testing.html")

@app.route("/api/new/<int:Index>")
def api_new(Index):
    try:
        return DB[Index]
    except IndexError:
        abort(404)


# Below is to display 404 Page not found if not valid URL
@app.route("/<PGNotFound>/")
def pagenotfound(PGNotFound):
    return "Are you lost? "


''' Adding static folder and placing files in it will be loadable as /static/file.gif '''
