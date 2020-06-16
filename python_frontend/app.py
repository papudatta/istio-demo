# This is the landing page for istio-demo.
# This landing page will fetch BGP ASN and prefix
# from golang microservices. There are 2
# versions of golang microservices

from flask import Flask, render_template, request
import requests

app = Flask(__name__)
@app.route('/healthz')
def health():
    return('i am cool')

@app.route('/', methods=['GET'])
def landing():
    try:
        # Below choice depends on the type of LB
        ip_address = request.headers.get('X-Forwarded-For'). \
                                         split(',')[0]
        # 
        # ip_address = request.remote_addr
        # ip_address = "122.171.152.151"
    except:
        # There will be an exception in goserv if this happens.
        ip_address = "There was a problem"

    payload = {'ipaddress': ip_address}
    r = requests.get("http://goserv:9091", params=payload)
    resp = r.json()
    asn = resp["asn"]
    prefix = resp["prefix"]
    return render_template('index.html',
                            ip_address=ip_address,
                            asn=asn,
                            prefix=prefix)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)