//kubeadm init --apiserver-advertise-address=10.0.8.9 --image-repository registry.aliyuncs.com/google_containers --kubernetes-version v1.14.0 --service-cidr=10.1.0.0/16 --pod-network-cidr=10.244.0.0/16
//
//
//
//mkdir -p $HOME/.kube
//sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
//sudo chown $(id -u):$(id -g) $HOME/.kube/config
//
//
//apiVersion: v1
//kind: ConfigMap
//Metadata:
//name: Nginx-config
//data:
//Nginx.conf: |
//	`user nginx;
//worker_processes auto;
//error_log /var/log/nginx/error.log;
//pid /run/nginx.pid;
//
//# Load dynamic modules. See /usr/share/doc/nginx/README.dynamic.
//include /usr/share/nginx/modules/*.conf;
//
//events {
//    worker_connections 1024;
//}
//
//http {
//    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
//                      '$status $body_bytes_sent "$http_referer" '
//                      '"$http_user_agent" "$http_x_forwarded_for"';
//
//    access_log  /var/log/nginx/access.log  main;
//
//    sendfile            on;
//    tcp_nopush          on;
//    tcp_nodelay         on;
//    keepalive_timeout   65;
//    types_hash_max_size 2048;
//
//    include             /etc/nginx/mime.types;
//    default_type        application/octet-stream;
//    include /etc/nginx/default.d/*.conf;
//
//
//    server {
//        listen       80 default_server;
//        listen       [::]:80 default_server;
//        server_name  _;
//        root         /usr/share/nginx/html;
//
//        # Load configuration files for the default server block.
//
//        location / {
//        }
//
//        error_page 404 /404.html;
//        location = /404.html {
//        }
//
//        error_page 500 502 503 504 /50x.html;
//        location = /50x.html {
//        }
//    }
//
//
//
//}
//include /etc/nginx/conf.d/*.conf;`
//
//
//other.conf: |
//	`stream {
//    upstream a1 {
//        server sg.low.im:80;
//        #server 18.181.194.226:80;
//    }
//    server {
//        listen 17113;
//        proxy_pass a1;
//    }
//    upstream a3 {
//        server hkt.rmbks.top:10150;
//    }
//    server {
//        listen 17115;
//        proxy_pass a3;
//    }
//    upstream a4 {
//        server hk.low.im:80;
//    }
//    server {
//        listen 17116;
//        listen 17116 udp;
//        proxy_pass a4;
//    }
//    upstream a5 {
//        server us.low.im:80;
//    }
//    server {
//        listen 17117;
//        listen 17117 udp;
//        proxy_pass a5;
//    }
//
//
//}`

