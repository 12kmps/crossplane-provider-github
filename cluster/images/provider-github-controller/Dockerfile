FROM BASEIMAGE
RUN apk --no-cache --repository https://artifactory.internal.cba:443/artifactory/alpine/v3.15/main add ca-certificates bash

ARG ARCH
ARG TINI_VERSION

ADD provider /usr/local/bin/crossplane-github-provider

EXPOSE 8080
ENTRYPOINT ["crossplane-github-provider"]
