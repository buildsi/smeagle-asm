FROM golang:bullseye as gobase
# docker build -t ghcr.io/buildsi/smeagleasm .
FROM ghcr.io/buildsi/smeagle
COPY --from=gobase /usr/local/go/ /usr/local/go/
WORKDIR /src/
COPY . /src/
ENV PATH /usr/local/go/bin:/code/build/standalone:${PATH}
RUN make
ENTRYPOINT ["/bin/bash"]
