# yiyaoshuju_html
医药数据网
nginx作反向代理的时候
在/etc/ngnix/sites-enable/   新建yiyaoshuju文件
 yiyaoshuju内容如下:
  server {
      listen 80;
	  listen {::}:80;
	  root /var/www/html/yiyaoshuju; #yiyaoshuju文件夹中放此网站的静态文件
	  index  index.html;
	  server_name yiyaoshuju.cn www.yiyaoshuju.cn;
	  location  /chaxun {
	    proxy_pass htpp://127.0.0.1:8090;
	  }
  }

 