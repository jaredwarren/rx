FROM scratch

EXPOSE 8080

COPY --chown=0:0 user-docker user

# COPY templates templates/ 


USER 65534

ENTRYPOINT ["/user"]

# BUILD WITH!!!!
# GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o user