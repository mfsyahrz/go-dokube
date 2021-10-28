// Kube

. Beberapa command untuk pods:

```
kubectl get pods
kubectl describe pod <nama-pod>
```

2. Konfigurasi yaml service (tipe nodeport)

```
apiVersion: v1
kind: Service
metadata:
  name: my-service1
spec:
  type: NodePort
  selector:
    name: my-app1
  ports:
    - port: 8080
      targetPort: 8080
```

3. Command apply service:

```
kubectl apply -f 3-create-service.yaml
```

4. Beberapa command untuk service:

```
kubectl get services
kubectl describe service <nama-service>
```

5. Command untuk get service url

```
minikube service <nama-service>
```