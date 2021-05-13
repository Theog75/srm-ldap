#!/bin/bash

/srm-ldap &

while true; do
    echo "Fetching LDAP";
    smbclient -U $SMBAUSER  //$SAMBASERVER/$SAMBASHARE -c "lcd /tmp; get "$LOGFILE"" $SAMBAPASSWORD  
    sleep 500;
    # cat  $DHCPFILENAME|grep -v ID|grep "," > /tmp/out.csv 
done

