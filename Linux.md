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

### [Xcfe](https://wiki.xfce.org/faq) for desktop environment
* add / remove keyboard shortcut from Application Menu (top left) --> Settings --> Keyboard --> Application Shortcuts
* Useful default short cut 
  * `Ctrl+Alt+Del` = lock screen (xflock4)
  * `Alt+F3` = App Finder
* Lists of all shortcut `xfconf-query -c xfce4-keyboard-shortcuts -l -v | cut -d'/' -f4 | awk '{printf "%30s", $2; print "\t" $1}' | sort | uniq`
* Shortcut for manage window `Applications Menu --> Setting --> Settings Manager --> window manager settings > Keyboard`
* Another place for keyboard shortcuts `Menu → Settings → Settings Editor, select "xfce4-keyboard-shortcuts`

### [tmux](http://code.tutsplus.com/tutorials/intro-to-tmux--net-33889) for terminal multiplexing
like web browse with will have many tabs. [tmux](https://danielmiessler.com/study/tmux/) let's you have many tabs for terminal
* can intall via "software manager"
* still not get it, will come back later
* reload config file `tmux source-file ~/.tmux.conf` from [here](https://gist.github.com/MohamedAlaa/2961058)
* sample tmux config https://raw.github.com/danielmiessler/tmux/master/.tmux.config
* example for setup keybinding http://unix.stackexchange.com/questions/5998/how-do-i-bind-to-shiftleft-right-in-tmux
* tmux man page http://www.openbsd.org/cgi-bin/man.cgi/OpenBSD-current/man1/tmux.1?query=tmux&sec=1#options

### [sublime text](sublimetext.md) for text editor
My favorite text editor, please go to sublime text page

### Samba file sharing
Linux side
 samba -V
 smbpasswd
 testparm #after edit conf
 smbclient -v \\\\198.10.10.71\\share -U sst
 net usershare info --long

Windows side
 net use * /del
 net use t: \\198.10.10.71\share /user:sst passwo
