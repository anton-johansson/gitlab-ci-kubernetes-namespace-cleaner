FROM golang:1.12.17-alpine3.11 AS build
RUN apk add --no-cache git make
COPY . /src/gitlab-ci-kubernetes-namespace-cleaner
WORKDIR /src/gitlab-ci-kubernetes-namespace-cleaner
RUN make linux

FROM alpine:3.11.3
RUN addgroup --system --gid 1000 gitlab-ci-kubernetes-namespace-cleaner \
 && adduser --system --uid 1000 -G gitlab-ci-kubernetes-namespace-cleaner gitlab-ci-kubernetes-namespace-cleaner
USER 1000
COPY --from=build /src/gitlab-ci-kubernetes-namespace-cleaner/bin/gitlab-ci-kubernetes-namespace-cleaner-linux-amd64 /usr/bin/gitlab-ci-kubernetes-namespace-cleaner
ENTRYPOINT ["gitlab-ci-kubernetes-namespace-cleaner", "clean"]
