
.PHONY: all cos crypto pastebin tinyurl user

cos:
	cd cos && go generate; cd -

crypto:
	cd crypto && go generate; cd -

pastebin:
	cd pastebin && go generate; cd -

tinyurl:
	cd tinyurl && go generate; cd -

user:
	cd user && go generate; cd -

all : cos crypto pastebin tinyurl user
	@echo "pb complied"