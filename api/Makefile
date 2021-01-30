
.PHONY:all crypto pastebin tinyurl user

crypto:
	cd crypto && go generate; cd -

pastebin:
	cd pastebin && go generate; cd -

tinyurl:
	cd tinyurl && go generate; cd -

user:
	cd user && go generate; cd -

all : crypto pastebin tinyurl user
	@echo "pb complied"