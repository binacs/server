
.PHONY:all static user crypto

static:
	cd static/v1 && go generate; cd -

user:
	cd user/v1 && go generate; cd -

crypto:
	cd crypto/v1 && go generate; cd -

all : static user crypto
	@echo "pb complied"