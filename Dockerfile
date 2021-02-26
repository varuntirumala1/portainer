FROM varuntirumala1/alpine-nginx:latest
COPY /etc/services.d/ /etc/services.d/
COPY /portainer-proxy.conf /config/nginx/site-confs/default
RUN rm /etc/services.d/php-fpm/run
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

RUN chmod +x /etc/services.d/cloudflared/run

VOLUME ["/data"]

EXPOSE 443

ENTRYPOINT ["/init","/portainer"]
