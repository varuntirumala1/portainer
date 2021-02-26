FROM varuntirumala1/alpine:latest
COPY /etc/services.d/ /etc/services.d/
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

EXPOSE 9000
EXPOSE 8000

ENTRYPOINT ["/init","/portainer"]
