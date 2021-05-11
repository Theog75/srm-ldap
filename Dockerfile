FROM debian

RUN apt update -y && apt install -y ldap-utils
COPY looper.sh /
COPY srm-ldap /
COPY run.sh /
RUN chmod 777 /looper.sh && chmod 777 /srm-ldap && chmod 777 /run.sh

USER 1001
#CMD ["/looper.sh"]
CMD ["/srm-ldap"]

