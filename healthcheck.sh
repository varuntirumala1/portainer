#!/bin/bash
serv=cloudflared
sstat=$(pidof $serv | wc -l )
if [ $sstat -gt 0 ]
then
echo "$serv is running fine!!!"
else
echo "$serv is down/dead"
service $serv start
echo "$serv service is UP now!!!"
fi
