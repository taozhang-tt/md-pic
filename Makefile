Bin=mdpic

build:
	go build -o ${Bin}
	mv ${Bin} /usr/local/bin/
	cp config.json /etc/mdpic.json

