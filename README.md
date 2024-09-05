# location ip address
列出ip所在区域

# 概述
在使用Cloudflare的Worker时，如果配置的后端服务器路由到了美国，由于美国的ip就要选中另一个
# Worker设置中的IP设置

![[IMG-20240524192351133.png]]
# 调用程序选前端ip
1. 在此处复制 https://www.cloudflare.com/ips-v4
2. 覆盖ipaddresslocation_go中的ip.txt
3. go run .
4. 得到iplist1

# 运行CloudflareST
1. 复制上面的iplist1到ip.txt中
2.  ./CloudflareST.exe -tll 50 -tl 200
3. 得到iplist2

# 查看后端具体位置
1. 设置Workerip
2. 在这个页面(https://lastpass.com/disablemultifactor.php?multifactortype=country)具体查看ip位置