FROM alpine:3.13
ADD cloudflared /etc/init.d/
RUN apk add --no-cache \
        openssl \
        curl \
        ca-certificates \
        nano \
        libc6-compat \
        bash \
        wget \
        openrc \
    && curl -s -O https://bin.equinox.io/c/VdrWdbjqyF/cloudflared-stable-linux-amd64.tgz \
    && tar zxf cloudflared-stable-linux-amd64.tgz \
    && mv cloudflared /bin \
    && rm cloudflared-stable-linux-amd64.tgz \
    && curl -s https://api.github.com/repos/portainer/portainer/releases/latest \
       | grep "browser_download_url.*portainer-[^extended].*-linux-amd64\.tar\.gz" \
       | cut -d ":" -f 2,3 \
       | tr -d \" \
       | wget -qi - \
   && tarball="$(find . -name "*linux-amd64.tar.gz")" \
   && tar -xzf $tarball \
   && mv portainer/portainer / \
   && chmod +x /portainer \
   && rm $tarball

RUN chmod +x /etc/init.d/cloudflared \
   && rc-update add cloudflared

ADD healthcheck.sh /
RUN chmod +x /healthcheck.sh
HEALTHCHECK --interval=1m CMD /healthcheck.sh

COPY Argo ./argo/
VOLUME /data
WORKDIR /

EXPOSE 9000
EXPOSE 8000

ENTRYPOINT ["/portainer"]
