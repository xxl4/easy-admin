# K8S Install easy-admin
> > Build a full-featured administrative interface quickly easy-admin

### 1、Ready
> tips: [how to install k8s or k3s](https://nicelizhi.github.io/easy-admin/guide/install/howtoinstallk8sork3s)
> tips: [how to delopy mysql on k8s]  
> tips: [how to delopy pg on k8s]  


### 2、Configure

[Configure Docs](https://nicelizhi.github.io/easy-admin/guide/configure/)

### 3、Start It
```
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: easy-admin
spec:
  selector:
    matchLabels:
      app: easy-admin
  template:
    metadata:
      labels:
        app: easy-admin
    spec:
      containers:
      - name: easy-admin
        image: nicesteven/easy-admin:latest
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 8080
        volumeMounts:
        - name: easy-admin
          mountPath: /temp
        - name: easy-admin
          mountPath: /static
        - name: easy-admin-config
          mountPath: /config/
          readOnly: true
    volumes:
      - name: easy-admin
        persistentVolumeClaim:
          claimName: easy-admin
      - name: easy-admin-config
        configMap:
          name: settings-admin
---
apiVersion: v1
kind: Service
metadata:
  name: easy-admin-svc
spec:
  selector:
    app: easy-admin-port
  ports:
    - port: 8080
      targetPort: 8080
---
````
```
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: easy-admin
  namespace: easy-admin
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: "1Mi"
  volumeName:
  storageClassName: nfs-csi
```


### 4、Test it

```
curl -v http://127.0.0.1:8000
```

### Issue Submit
https://github.com/nicelizhi/easy-admin/issues