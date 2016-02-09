FROM gliderlabs/alpine:3.3
MAINTAINER  Kent Cederström <kent@kentcederstrom.se>

# Update alpine and get some dependcies for the system
RUN apk update && \
	apk add go && \
	rm -rf /var/cache/apk/* && \
	apk add --update -t build-deps git

# Add some enviroment vars
# Sys and go vars
ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

# Set work dir and add source code
WORKDIR /gopath/src/kentos/gooskfotboll
ADD . /gopath/src/kentos/gooskfotboll/

# Download depencies TODO: find a better way to handle this
RUN go get github.com/gorilla/mux
RUN go get github.com/PuerkitoBio/goquery

# Install the application
RUN go install

# Clean deps
RUN apk del --purge build-deps && \
	rm -rf /gopath/src/* && \
	rm -rf /var/cache/apk/*

# Run it
CMD []
EXPOSE 8080
ENTRYPOINT ["gooskfotboll"]