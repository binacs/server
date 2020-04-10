
.PHONY:all static user

static:
	cd static/v1 && go generate; cd -

user:
	cd user/v1 && go generate; cd -

all : static user
	@echo "pb complied"