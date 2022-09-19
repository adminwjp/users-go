# linux devlop
```cmd
microk8s.exe --help



microk8s.exe status
microk8s is not running. Use microk8s inspect for a deeper inspection.

microk8s.exe start

microk8s.exe status --wait-ready
```

# linux 
# ubuntu
sudo snap install docker
sudo snap remove docker

apt install docker
reboot 
images 不见了
Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon runnin

ubuntu xftp file not upload
owcloud upload file

not used
```conf
vim /etc/vsftpd.conf

anonymous_enable=NO
local_enable=YES
write_enable=YES

 /etc/init.d/vsftpd restart
```
[win k8s kind]
https://www.jianshu.com/p/d42cfa67fa84

https://blog.csdn.net/weixin_45277567/article/details/126305535

https://my.oschina.net/ityzl/blog/5325037

#  (差 k8s 组件  组件在哪 端口都没看见 放的地方不是同一个)
# https://blog.csdn.net/eyeofeagle/article/details/85227414/ 


# https://www.jianshu.com/p/bd293b00fbaf/

https://blog.csdn.net/qq_20466211/article/details/122767891

microk8s status
microk8s enable xx xxx ....

https://snapcraft.io/microk8s
https://microk8s.io/docs
https://microk8s.io/docs/command-reference#heading--microk8s-enable

https://cloud.tencent.com/developer/article/1864390
https://helm.sh/

microk8s enable registry
microk8s images --help
microk8s  -h

chmod 777 /home/software/k8s/

 sudo curl -L  "https://raw.githubusercontent.com/OpsDocker/pullk8s/main/pullk8s.sh"  -o /usr/local/bin/pullk8s
 sudo curl -L  "https://github.com/OpsDocker/pullk8s/blob/main/pullk8s.sh"  -o /usr/local/bin/pullk8s
 
 sudo chmod +x /usr/local/bin/pullk8s
 pullk8s check --microk8s
 
 microk8s enable community
 microk8s enable istio
 Ingress gateways encountered an error
 
 kubectl get pod -n istio-system
 
microk8s ctr i -h
docker images 
microk8s ctr i ls
microk8s ctr image ls

microk8s.helm3

kubectl get svc -n kube-system

microk8s kubectl port-forward -n kubernetes-dashboard --address=0.0.0.0 service/kubernetes-dashboard 10443:443 &
microk8s.kubectl proxy &

microk8s kubectl port-forward -n kube-system --address=0.0.0.0 service/kube-system 8020:8020 &

microk8s kubectl port-forward -n kube-system --address=0.0.0.0 service/go-users 8020:8020 #error
microk8s kubectl port-forward -n kube-system --address=0.0.0.0 service/go-users 30020:30020 #error

microk8s kubectl port-forward -n kube-system --address=0.0.0.0 deployment/go-users 8020:8020
microk8s kubectl port-forward -n kube-system --address=0.0.0.0 deployment/go-users 30020:30020 # timeout

microk8s kubectl get service -n istio-system  service/istio-ingressgateway
microk8s kubectl get service -n istio-system 
microk8s kubectl -n istio-system edit service/istio-ingressgateway
spec:
  externalIps:
  - 192.168.1.6
  
microk8s kubectl label namespace default istio-injection=enable

docker save go_users:v1 >go_user.tar

docker export go_users >go_user.tar

microk8s ctr image import go_user.tar
```sh
microk8s.helm3 list

microk8s.helm3 search hub wordpress
```
```sh
kubectl create -f dashboard-adminrole.yaml
kubectl apply -f dashboard-adminrole.yaml
kubectl get -f dashboard-adminrole.yaml
kubectl delete -f dashboard-adminrole.yaml

kubectl --help
kubectl  delete --help
kubectl  create --help
kubectl  get --help
kubectl  apply --help

# ....

# token is empty
# https://www.cnblogs.com/jishuzhaichen/p/16348711.html

kubectl proxy --port=8001
curl http://localhost:8001/api/

#3600*24*365=31536000

重新安装 token 失效
kubectl delete pods -n kubernetes-dashboard
kubectl delete services -n kubernetes-dashboard
kubectl delete deployments -n kubernetes-dashboard
kubectl delete  -f dashboard.yaml

kubectl apply  -f dashboard.yaml
kubectl apply -f  dashboard-adminuser.yaml
kubectl apply -f dashboard-adminuser1.yaml
curl 'http://127.0.0.1:8001/api/v1/namespaces/kubernetes-dashboard/serviceaccounts/admin-user/token' -H "Content-Type:application/json" -X POST -d '{"kind":"TokenRequest","apiVersion":"authentication.k8s.io/v1","metadata":{"name":"admin-user","namespace":"kubernetes-dashboard"},"spec":{"audiences":["https://kubernetes.default.svc"],"expirationSeconds":31536000}}'

# 3600s
curl 'http://127.0.0.1:8001/api/v1/namespaces/kubernetes-dashboard/serviceaccounts/admin-user/token' -H "Content-Type:application/json" -X POST -d '{}'

# kubernetes-dashboard
# the server could not find the requested resource
# https://www.cnblogs.com/l-hh/p/14833146.html
https://github.com/kubernetes/dashboard/releases
#1.25.0(current) > 1.24.0(new)
kubernetesui/dashboard:v2.6.1
kubernetesui/metrics-scraper:v1.0.8

kubernetesui/dashboard:v2.5.0
kubernetesui/metrics-scraper:v1.0.7

kubernetes-dashboard beat
kubectl get service 
kubectl get services
kubectl get service -o wide
kubectl get service -n kubernetes-dashboard
kubectl get services -n kubernetes-dashboard
kubectl get deployments
kubectl get deployments -n kubernetes-dashboard

kubectl delete svc kubernetes-dashboard
kubectl delete deployments kubernetes-dashboard
kubectl delete deployments kubernetes-dashboard -n kubernetes-dashboard

# https://www.cnblogs.com/smileblogs/p/16421407.html
# 清理顺序：pod>pvc>pv
kubectl delete pod  xx
kubectl delete pvc  xx
kubectl delete pv   xx

# https://blog.csdn.net/qq_43114229/article/details/124208635
kubectl delete -f   pod.yaml

kubectl config view

kubectl delete -f  dashboard-adminuser.yaml
vim dashboard-adminuser1.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard

 
kubectl apply -f dashboard-adminuser1.yaml
 
vim  dashboard-adminrole.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard  

kubectl apply -f dashboard-adminrole.yaml
 
sudo snap install docker
pullk8s check --microk8s
snap alias microk8s.docker pullk8s


docker pull 
docker pull  kubernetesui/metrics-scraper:v1.0.6
docker pull  kubernetesui/dashboard:v2.1.0
snap install kubeadm --classic

# kubeadm config images list --kubernates-version=v1.25.0 --v=5
kubeadm config images list

# 隔离环境没关系
# https://www.jianshu.com/p/70bc515b4e0e
## 使用阿里源

curl -fsSL http://mirrors.aliyun.com/docker-ce/linux/ubuntu/gpg | apt-key add - #安装GPG证书
add-apt-repository "deb [arch=amd64] https://mirrors.aliyun.com/docker-ce/linux/ubuntu $(lsb_release -cs) stable" #写入docker-ce软件源信息

for i in `kubeadm config images list`; do docker pull registry.aliyuncs.com/google_containers/${i:11}; done

for i in `kubeadm config images list`; do docker pull ${i:11}; done


snap install microk8s --classic

vim /var/snap/microk8s/current/args/kubelet
--pod-infra-container-image=s7799653/pause:3.1

vim /var/snap/microk8s/current/args/containerd-template.toml
sandbox_image = "s7799653/pause:3.1"

[plugins.cri.registry.mirrors."local.insecure-registry.io"]

endpoint = ["http://localhost:32000"]

microk8s stop
microk8s start
microk8s enable dns dashboard

kubectl get nodes
kubectl get nodes -w
kubectl describe node wjp-virtual-machine

snap info microk8s 

snap alias microk8s.kubectl kubectl

ls /snap/bin

echo 'export PATH=$PATH:/snap/bin'>>~/.bashrc &&source ~/.bashrc

mkdir -p $HOME/.kube

microk8s.kubectl config view --raw >$HOME/.kube/config

vim $HOME/.kube/config

microk8s.enable dns
vim dashboard.yaml
microk8s.kubectl  create -f dashboard.yaml

microk8s.enable dashboard

vim dashboard-adminuser.yaml
microk8s.kubectl  apply -f dashboard-adminuser.yaml

microk8s.kubectl -n kubernetes-dashboard describe secret $(microk8s.kubectl -n kubernetes-dashboard  get secret | grep admin-user | awk '{print $1}')

microk8s.kubectl -n kube-system  describe secret $(microk8s.kubectl -n kube-system   get secret | grep admin-user | awk '{print $1}')

microk8s.kubectl -n kubernetes-dashboard describe secret $(microk8s kubectl -n kubernetes-dashboard get secret | grep default-token | cut -d " " -f1)

 token=$(microk8s kubectl -n kube-system get secret | grep default-token | cut -d " " -f1)
 microk8s kubectl -n kube-system describe secret $token
 
  token=$(microk8s kubectl -n kubernetes-dashboard get secret | grep default-token | cut -d " " -f1)
 microk8s kubectl -n kubernetes-dashboard describe secret $token
 
 kubectl -n kube-dashboard describe $(kubectl -n kube-dashboard get secret -n kube-dashboard -o name | grep namespace) | grep token
 
  kubectl -n kube-system describe $(kubectl -n kube-system get secret -n kube-system -o name | grep namespace) | grep token
  
netstat -nptl

snap alias microk8s.kubectl kubectl



microk8s.kubectl get po -n kubernetes-dashboard

kubectl get pods -n kube-system


kubectl describe pod kubernetes-dashboard-568987f649-ljm6q

kubectl get all --all-namespaces
kubectl get rc,services --all-namespaces

kubectl get svc -n kubernetes-dashboard

# https://www.cnblogs.com/shanyou/p/16212194.html



```
