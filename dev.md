This is software development instruction for linux user

### [Docker](https://www.docker.com/) for app container
* from [install guide](https://docs.docker.com/installation/ubuntulinux/) it's easy like typing `wget -qO- https://get.docker.com/ | sh` and done. But I have errors after install
* if you type `sudo docker -d` to run daemon error message will show that no apparmor install so `sudo apt-get install apparmor`
* `docker run hello-world` to test your docker installation
