.PHONY: private_ask,ask,set, build,install,alias-instructions,all

private_ask:
	go run . diff ~/Desktop/teknasyon-staj/language-learning-api main init

ask:
	go run . ask diff ${path}	

set:
	go run . set key berkay-gpt-turbo

set1:
	go run . set key berkay-gpt-turboo

rm :
	go run . rm key berkay-gpt-turboo

all: 
	go build -o git-ai .
	sudo mv git-ai /usr/local/bin
	sudo cp -r git-ai-storage /usr/local/bin
	alias git-ai='/usr/local/bin/git-ai'
