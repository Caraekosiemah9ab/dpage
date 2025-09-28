<img width="1600" height="800" alt="dpage" src="https://github.com/user-attachments/assets/eb70a3c8-871e-4021-889c-5b200f5c0703" />

# Dynamic Page (dpage) — простой, но мощный веб-сервер с поддержкой плагинов <a href="https://ongg.net/threads/55/" target="_blank">[Official Topic]</a>

> Генерируй статические и динамические страницы с помощью шаблонов и расширяемых плагинов на Go.

`dpage` — это лёгкий веб-сервер, который позволяет:
- Отображать страницы через **HTML-шаблоны** (Go templates),
- Динамически подгружать функциональность через **плагины** (`.so`),
- Управлять навигацией через **конфигурационные файлы**.

Nginx Configuration:

```nginx
server {
    listen 80;
    listen [::]:80;
    server_name YOURDOMAINNAME.ZONE;
    return 301 https://YOURDOMAINNAME.ZONE$request_uri;
}

server {
    listen 443 ssl http2;
    listen [::]:443 ssl http2;
    server_name YOURDOMAINNAME.ZONE;

    ssl_certificate /etc/letsencrypt/live/YOURDOMAINNAME.ZONE/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/YOURDOMAINNAME.ZONE/privkey.pem;

    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-SHA256:ECDHE-RSA-AES256-SHA384;
    ssl_prefer_server_ciphers off;
    ssl_session_cache shared:SSL:10m;
    ssl_session_timeout 10m;

    add_header X-Frame-Options DENY;
    add_header X-Content-Type-Options nosniff;
    add_header X-XSS-Protection "1; mode=block";

    client_max_body_size 1024M;

    location / {
        proxy_pass http://Addr:Port; # watch in server.json
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host $host;
        proxy_set_header X-Forwarded-Port $server_port;
        proxy_redirect off;
        proxy_http_version 1.1;
        proxy_cache_bypass $http_upgrade;
    }
}

