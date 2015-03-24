* My favorite distro is [linux mint])(www.linuxmint.com)
* After install use "Software Manager" for install new app. But sometime it's fail to run after install. For example after install xrdp I cannot remote desktop from my machine

##### Remote Desktop Server using [xrdp](www.xrdp.org)
because xrdp have depencencies, and cannot use cinnamon as desktop. Here's [solution i found](http://c-nergy.be/blog/?p=5305)
```
sudo apt-get install tightvncserver
sudo apt-get insatll xrdp
---- restart machine ----
sudo apt-get install xfce4
echo xfce4-session >~/.xsession
sudo service xrdp restart
```
tip: you can see log file using `sudo tail /var/log/xrdp-sesman.log`

