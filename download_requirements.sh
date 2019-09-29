#!/bin/bash
curl -sL https://deb.nodesource.com/setup_8.x | sudo -E bash -
sudo apt install -y python3 python3-setuptools nodejs npm golang
sudo npm i -g npm
npm install pug
npm install pug-cli -g
npm install less -g
npm install angular@1.6.5
npm install angularjs-datepicker
ng -v
sudo mv node_modules web/js
#cd ./web/js
#git clone https://github.com/angular/quickstart.git quickstart
