FROM debian

RUN apt update -y && apt install -y smbclient
COPY looper.sh /
COPY srm-dhcp /
COPY dhcp.sh /
RUN chmod 777 /looper.sh && chmod 777 /srm-dhcp && chmod 777 /dhcp.sh

USER 1001
#CMD ["/looper.sh"]
CMD ["/dhcp.sh"]

