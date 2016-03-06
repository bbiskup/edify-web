FROM golang:1.6
MAINTAINER bernhard.biskup@gmx.de


WORKDIR /
RUN wget -v https://nodejs.org/dist/v5.7.1/node-v5.7.1-linux-x64.tar.gz
RUN tar xzvf node-v5.7.1-linux-x64.tar.gz
RUN rm node-v5.7.1-linux-x64.tar.gz
ENV PATH=/node-v5.7.1-linux-x64/bin/:$PATH
RUN npm install -g bower

RUN mkdir -p /go/src/github.com/bbiskup
RUN useradd -m -d /go/src/github.com/bbiskup/edify-web dev
RUN chown -R dev:dev /go/src/
USER dev
WORKDIR /go/src/github.com/bbiskup/edify-web

# Install edify
RUN go get github.com/bbiskup/edify

# COPY . /go/src/github.com/bbiskup/edify-web
ADD bower.json bower.json
RUN bower install

ADD ./main.go main.go
ADD ./handlers handlers
ADD ./defs defs
ADD static static
ADD scripts scripts

# Get EDIFACT specifications
RUN edify download_specs
RUN edify extract_specs
RUN ls -la

RUN ./scripts/get_test_deps.sh
RUN go get -t ./...
RUN go build -v

EXPOSE 8001
ENTRYPOINT ["/bin/bash", "-c"]
