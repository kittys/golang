* My favorite distro is [linux mint])(www.linuxmint.com)
* After install use "Software Manager" for install new app. But sometime it's fail to run after install. For example after install xrdp I cannot remote desktop from my machine

### Remote Desktop Server using [xrdp](www.xrdp.org)
because xrdp have depencencies, and cannot use cinnamon as desktop. Here's [solution i found](http://c-nergy.be/blog/?p=5305)
```
sudo apt-get install tightvncserver
sudo apt-get insatll xrdp
---- restart machine ----
sudo apt-get install xfce4
echo xfce4-session >~/.xsession
sudo service xrdp restart
```
tip: you can see log file using `sudo tail /var/log/xrdp-sesman.lo`g
After finish found 2 more problem 
* everytime I log in always get xcfe desktop not cinnamon, not only RDP. But it's OK due to xcfe is lightweight
* cannot see character i type in terminal --> because font color and background color is same, need change it in terminal profile (right click --> show menu bar --> edit --> profile...)

### Remote Desktop Client using [remmina](http://sourceforge.net/projects/remmina/)
if you install from software manage. It's will not have rdp plugin so you cannot remote to windows PC. Please try this
``` 
sudo apt-add-repository ppa:remmina-ppa-team/remmina-next
sudo apt-get update
sudo apt-get install remmina remmina-plugin-rdp
```

### Remote Desktop from anywhere with [TeamViewer]()
go to teamviewer web site. Download deb file then doubleclick to install or
```
wget http://download.teamviewer.com/download/teamviewer_linux.deb; 
sudo dpkg -i teamviewer_linux.deb; 
```

### [Google Chrome](https://www.google.com/chrome/browser/features.html) Web Browser instead of Firefox
* try this `wget -c wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb; sudo dpkg -i google-chrome-*.deb; sudo apt-get install -f`
* recommend extension [Postman](https://chrome.google.com/webstore/detail/postman-rest-client-packa/fhbjgbiflinjbdggehcddcbncdddomop?utm_source=chrome-ntp-launcher), [Postman Interceptor](https://chrome.google.com/webstore/detail/postman-interceptor/aicmkgpgakddgnaphhhpliifpcfhicfo), [Nicer Invertor](https://chrome.google.com/webstore/detail/nicer-inverter/oichlckdgnbjkmhaebnnhibamjgpndkm)
