* My favorite distro is [linux mint])(www.linuxmint.com)
* After install use "Software Manager" for install new app
* use [xrdp](www.xrdp.org) for remote desktop server

```bash
sudo apt-get install tightvncserver
sudo apt-get insatll xrdp
sudo tail /var/log/xrdp-sesman.log
sudo echo “cinnamon” > ~/.Xclients
sudo chmod +x ~/.Xclients
sudo service restart xrdp.service```
