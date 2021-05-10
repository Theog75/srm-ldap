#!/bin/bash

/srm-dhcp &

export SAMBASERVER=$(cat /sambaauth/server)
export SAMBASHARE=$(cat /sambaauth/share)
export SMBAUSER=$(cat /sambaauth/username)
export SAMBAPASSWORD=$(cat /sambaauth/password)
DOW=$(date +"%a")
DHCPFILENAME="DhcpSrvLog-"$DOW".log"

while true; do
    echo "Fetching file "${DHCPFILENAME}${DOW};
    sleep 5;
    smbclient -U $SMBAUSER  //$SAMBASERVER/$SAMBASHARE -c "lcd /tmp; get "$LOGFILE"" $SAMBAPASSWORD  
    # cat  $DHCPFILENAME|grep -v ID|grep "," > /tmp/out.csv 
done

