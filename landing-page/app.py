# This is the landing page for istio-demo.
# This landing page will fetch BGP info
# from goland microservices. There are 2
# versions of golang microservices

from flask import Flask, render_template, request
from datetime import datetime
import requests

app = Flask(__name__)
@app.route('/', methods=['GET'])
def landing():
    try:
        ip_address = request.headers.get('X-Forwarded-For'). \
                                        split(',')[0]
    except:
        ip_address = "There was a problem"

    payload = {'ipaddress': ip_address}
    r = requests.get("http://goserv1", params=payload)
    #asn, prefix = "asn", "prefix"
    resp = r.json()
    asn = resp["asn"]
    prefix = resp["prefix"]
    return render_template('index.html',
                            ip_address=ip_address,
                            asn=asn,
                            prefix=prefix)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)