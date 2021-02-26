FROM varuntirumala1/alpine:latest
COPY /etc/ /etc/

RUN apk add --no-cache nginx
RUN cd /tmp \  
    && curl -s https://api.github.com/repos/portainer/portainer/releases/latest \
       | grep "browser_download_url.*portainer-[^extended].*-linux-amd64\.tar\.gz" \
       | cut -d ":" -f 2,3 \
       | tr -d \" \
       | wget -qi - \
   && tarball="$(find . -name "*linux-amd64.tar.gz")" \
   && tar -xzf $tarball \
   && cp -R /tmp/portainer/* / \
   && rm $tarball \
   && rm -rf /tmp/*

RUN mkdir -P /run/nginx \
&& rm /usr/share/nginx/http-default_server.conf \
&& rm /etc/nginx/conf.d/default.conf
COPY /portainer.conf /etc/nginx/http.d/portainer.conf

VOLUME ["/data"]

EXPOSE 443

ENTRYPOINT ["/init","/portainer"]
