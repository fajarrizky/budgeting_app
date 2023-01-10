###########################
# Build Container 
###########################
FROM golang:1.19 AS build
WORKDIR /src

# ARG BITBUCKET_SSH_KEY
# RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config
# RUN echo "${BITBUCKET_SSH_KEY}" > /root/.ssh/id_rsa && ssh-keyscan bitbucket.com >> /root/.ssh/known_hosts
# RUN chmod 700 /root/.ssh/id_rsa
# RUN git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o budgetapp

###########################
# Application Container
###########################
FROM alpine:3.17
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/budgetapp /app/budgetapp
COPY --from=build /src/env/app.env /app/.env
COPY --from=build /src/migrations/ /app/migrations/ 
CMD ["./budgetapp", "api"]