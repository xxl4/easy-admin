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
apiVersion: v1
kind: Service
metadata:
  name: easy-admin
  labels:
    app: easy-admin
    service: easy-admin
spec:
  ports:
  - port: 8000
    name: http
    protocol: TCP
  selector:
    app: easy-admin
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: easy-admin-v1
  labels:
    app: easy-admin
    version: v1
spec:
  replicas: 1
  selector:
    matchLabels:
      app: easy-admin
      version: v1
  template:
    metadata:
      labels:
        app: easy-admin
        version: v1
    spec:
      containers:
      - name: easy-admin
        image: registry.ap-southeast-1.aliyuncs.com/kuops/easy-admin:1.10
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 8000
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

### Issue Submit
https://github.com/nicelizhi/easy-admin/issues