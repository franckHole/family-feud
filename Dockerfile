ARG FROM_IMAGE=amd64/golang:1.20-alpine
ARG PROD_IMAGE=scratch

FROM ${FROM_IMAGE} as base

ARG USER_UID=user

ENV USER_UID=${USER_UID}
ENV WORKDIR=/app

WORKDIR ${WORKDIR}

# Add user to not run tests and prod as root and create passwd file with user only for prod stage
RUN addgroup ${USER_UID} && adduser -D -G ${USER_UID} ${USER_UID} && grep ${USER_UID} /etc/passwd > /etc/passwd-prod

# Check go modules
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify

# Copy repository
COPY ./ ./

FROM base as test

# Avoid permissions errors on tests
RUN chown -R ${USER_UID}:${USER_UID} $WORKDIR
RUN apk add make
RUN make install-dev

USER ${USER_UID}

FROM base as builder
ARG ARCH=amd64
ARG APP_VERSION
RUN GOARCH=${ARCH} go build -ldflags="-w -s ${APP_VERSION:+-X github.com/franciscolkdo/family-feud/cmd.Version=${APP_VERSION}}" -o /family-feud

FROM ${PROD_IMAGE} as prod

COPY --from=base /etc/passwd-prod /etc/passwd
COPY --from=builder /family-feud /family-feud
COPY --from=builder /app/config/game.json /config/game.json

ARG USER_UID=user
ARG APP_VERSION
ARG COMMIT_ID

ENV APP_VERSION=${APP_VERSION}
ENV COMMIT_ID=${COMMIT_ID}
# Enable colors in container
ENV COLORTERM=truecolor

USER ${USER_UID}

ENTRYPOINT [ "/family-feud" ]
