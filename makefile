.PHONY: private_ask,ask,set, build,install,alias-instructions,all

ask:
	git ai diff ${path}	

set:
	git ai set key gpt-turbo-key
 
rm :
	git ai rm key 

all: 
	go build -o git-ai .
	sudo mv git-ai /usr/local/bin
	sudo cp -r git-ai-storage /usr/local/bin
	alias git-ai='/usr/local/bin/git-ai'
