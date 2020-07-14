### istio-demo

##### We can go directly to yamls and try those. On the other hand, if you want to push the container images to your own docker hub account, maybe after a change, please use the usual docker build/push commands.
```
docker login
docker build -t docker_username/<image_name>:1 .
docker push docker_username/<image_name>:1
```
##### Update the yaml with the above image. Can do it manually or use sed

**Overview of this demo app** \
FrontEnd runs a Python flask app \
Backend runs a golang service \
v1 of golang service returns ASN of the IP you are connecting from \
v2 returns Country code along with ASN.

**Steps:**
1. I'd run the yamls in the numbered fashion.
2. Can skip the first `1_verify.yaml` as it is only meant to verify
3. Do a 'kubectl delete -f 1_verify.yaml` before proceeding to the remaining
4. Start executing from 2 onwards
5. Notice a loadbalancer being configured after `3-istio-demo-gateway1.yaml` 
6. Run `kubectl patch svc istio-ingressgateway -n istio-system -p '{"spec":{"externalTrafficPolicy":"Local"}}'`
   Else, we will not get expected results when making an http connection.

**Useful commands to inspect above implementations.**

```
kubectl get gateway istio-demo-gateway -o yaml

kubectl get virtualservice istio-demo-vservice -o yaml

kubectl get virtualservice goserv -o yaml
```

```
istioctl proxy-status

istioctl x describe pod $(k get pod -l app=goserv -o jsonpath="{.items[0].metadata.name}")

istioctl pc cluster -n istio-system $(k get pod -n istio-system -l \
      app=istio-ingressgateway -o jsonpath="{.items[0].metadata.name}")

istioctl pc listeners -n istio-system $(k get pod -n istio-system -l \
      app=istio-ingressgateway -o jsonpath="{.items[0].metadata.name}")

istioctl pc endpoints -n istio-system $(k get pod -n istio-system -l \
      app=istio-ingressgateway -o jsonpath="{.items[0].metadata.name}")

istioctl pc routes -n istio-system $(k get pod -n istio-system -l \
      app=istio-ingressgateway -o jsonpath="{.items[0].metadata.name}") -o json

istioctl pc cluster $(k get pod -l app=goserv -o jsonpath="{.items[0].metadata.name}")

istioctl pc cluster $(k get pod -l app=flask -o jsonpath="{.items[0].metadata.name}")

istioctl pc listeners $(k get pod -l app=goserv -o jsonpath="{.items[0].metadata.name}")

istioctl pc listeners $(k get pod -l app=flask -o jsonpath="{.items[0].metadata.name}")

istioctl pc endpoints $(k get pod -l app=goserv -o jsonpath="{.items[0].metadata.name}")

istioctl pc endpoints $(k get pod -l app=flask -o jsonpath="{.items[0].metadata.name}")

istioctl pc routes $(k get pod -l app=flask -o jsonpath="{.items[0].metadata.name}") -o json

istioctl pc routes $(k get pod -l app=goserv -o jsonpath="{.items[0].metadata.name}") -o json
```
