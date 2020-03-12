binary_name=go_login
binary_path=${PWD}/bin

.PHONY: default install test runserver clean migrate

default:
	@go build -o $(binary_name)

clean:
	@rm ${binary_path}/${binary_name}

install:
	@go build -o ${binary_path}/${binary_name} main.go

runserver: install
	${binary_path}/${binary_name} runserver

migrate: install
	${binary_path}/${binary_name} migrate up

roolback: install
	${binary_path}/${binary_name} migrate down
