This is operation / sysadmin instruction for linux user

### [Docker](https://www.docker.com/) for app container
* from [install guide](https://docs.docker.com/installation/ubuntulinux/) it's easy like typing `wget -qO- https://get.docker.com/ | sh` and done. But I have errors after install
* if you type `sudo docker -d` to run daemon error message will show that no apparmor install so `sudo apt-get install apparmor`
* `docker run hello-world` to test your docker installation
* docker daemon will run after reboot. I also need to add my user to docker group `sudo usermod -aG docker {{my user}}`

#### [docker basic]() search "docker 101" on dotcloudtv youtube
* `docker version` and `docker info` to see overview
* in breif `images ps run stop start rm logs attach inspect commit push pull build`
Get start
* type `docker images` to see how many image you have in system.
* run imageName:Tag -it=interactive in terminal, -d=background. you type `exit` to exit from interactive mode
* `docker run -it ubuntu bash` will download ubuntu image (200MB) then run bash in interactive mode. You'll see
prompt change because you're inside container. try `ps -aux` it's will show only bash is running
* `docker ps` show process inside docker, `docker ps -a` show all include cached image that stop run
* `exit` to go out from container

diff / commit /
 `docker diff {{id}}` to see what's change in that run id
 `docker commit -m "your message# {{id}} {{user/imagename}}` to make new image. you can see by type `docker images`
 `docker login` to connect to hub.docker.io

* TODO: read more [basic docker](https://docs.docker.com/articles/basics)
 
### [shipyard](http://shipyard-project.com/) for docker management
Setup this due to want to know how hard/easy to run docker app
* `docker pull shipyard/deploy` to get shipyard image
* run it `docker run --rm -v /var/run/docker.sock:/var/run/docker.sock shipyard/deploy start` --> read [quick start](http://shipyard-project.com/docs/quickstart/). it's will download rethinkdb and shipyard app then start
* go to localhost:8080 then input user/password --> you need to add engines in shipyard to start uses (I stuck here for few hours)
* to add engine, you need first make docker daemon accept tcp connect on port 2375 using [instruciton here](http://www.virtuallyghetto.com/2014/07/quick-tip-how-to-enable-docker-remote-api.html)
  * `sudo gedit /etc/init/docker.conf`
  * add `DOCKER_OPTS='-H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock'`
  * `sudo service docker restart`
  * THIS IS UNSECURE because docker run as root. you need at least [enable https](https://docs.docker.com/articles/https/)
* add engine to shipyard --> DONE
