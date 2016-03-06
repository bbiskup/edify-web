all: build up



######### Deployment (Ansible-based; obsolete)

# # Provisioning public-facing demo server
# provision-public:
# 	ansible-playbook -i deploy/inventory/ deploy/playbook_demoserver.yml

# # Provision development container
# provision-dev:
# 	ansible-playbook -i deploy/inventory/ deploy/playbook_devserver.yml -v

######################### Dockerised

all: build up

build:
	docker-compose build

up:
	docker-compose up -d

bash:
	./docker-cmd.sh bash

test:
	# Run test in current dir and all subdirectories
	./docker-cmd.sh go test ./...

test-race:
	# Run test in current dir and all subdirectories
	./docker-cmd.sh go test -race -bench . ./...

test-multi-cpu:
	./docker-cmd.sh go test -cpu=1,2,3,4,5,6,7,8 ./...

test-verbose:
	./docker-cmd.sh go test -v ./...

bench:
	# Run test in current dir and all subdirectories
	./docker-cmd.sh go test -bench . ./...

get-deps:
	./docker-cmd.sh go get -t ./...

cover:
	./docker-cmd.sh go test -cover ./...

# Coverage for a particular package
# go test -coverprofile=coverage.out ./...

# Show package coverage in web browser
# go tool cover -html=coverage.out

quality: mccabe nyet 

mccabe:
	./docker-cmd.sh gocyclo -over 9 .

nyet:
	./docker-cmd.sh go-nyet ./...

clean:
	./docker-cmd.sh go clean ./...
	git clean -f -d