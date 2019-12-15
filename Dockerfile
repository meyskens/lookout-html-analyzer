FROM alpine

ADD ./build/bin/lookout-html-analyzer /bin/lookout-html-analyzer

CMD ["/bin/lookout-html-analyzer"]
