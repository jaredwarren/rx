############################
# 1. build executable binary
############################
# FROM golang:1.13.0-stretch AS builder

# ENV GO111MODULE=on \
#     CGO_ENABLED=1

# # Create appuser
# ENV USER=appuser
# ENV UID=10001
# # See https://stackoverflow.com/a/55757473/12429735RUN 
# RUN adduser \    
#     --disabled-password \    
#     --gecos "" \    
#     --home "/nonexistent" \    
#     --shell "/sbin/nologin" \    
#     --no-create-home \    
#     --uid "${UID}" \    
#     "${USER}"

############################
# 2. Create the minimal runtime image
############################
FROM scratch

EXPOSE 8080

# Import from builder.
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /etc/passwd /etc/passwd
# COPY --from=builder /etc/group /etc/group


COPY --chown=0:0 web-docker web

COPY templates templates/ 

# Set up the app to run as a non-root user inside the /data folder
# User ID 65534 is usually user 'nobody'. 
# The executor of this image should still specify a user during setup.
# COPY --chown=65534:0 --from=builder /data /data

USER 65534
# USER appuser:appuser
# WORKDIR /data

ENTRYPOINT ["/web"]

# BUILD WITH!!!!
# GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o web