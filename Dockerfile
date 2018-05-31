FROM scratch

ENV PORT 8000
ENV INTERNAL_PORT 8001

EXPOSE $PORT
EXPOSE $INTERNAL_PORT

COPY ./bin/linux-amd64/goworkshop /

CMD ["/goworkshop"]
