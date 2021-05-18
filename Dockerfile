FROM debian

USER root
RUN apt update -y && apt install -y ldap-utils
COPY looper.sh /
COPY srm-ldap /
COPY run.sh /
COPY ldap-query.sh /
RUN chmod 777 /looper.sh && chmod 777 /srm-ldap && chmod 777 /run.sh 

USER 1001
#CMD ["/looper.sh"]
CMD ["/run.sh"]

