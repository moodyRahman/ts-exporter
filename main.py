from dotenv import load_dotenv
import os
from flask import *
import requests as r

load_dotenv()  # Loads variables from .env into the environment
app = Flask(__name__)

@app.route("/")
def home():
    out = r.get("https://api.tailscale.com/api/v2/tailnet/TBfy7apG4t11CNTRL/devices", headers={"Authorization":"Bearer " + os.getenv("TS_AUTHKEY") } )
    return out.json()

if __name__ == "__main__":
    app.run(host="0.0.0.0", debug=True)
    