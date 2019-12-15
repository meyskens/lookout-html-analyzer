FROM alpine

RUN apk add --no-cache ca-certificates

ADD ./build/bin/lookout-html-analyzer /bin/lookout-html-analyzer

CMD ["/bin/lookout-html-analyzer"]
