#!/bin/bash

/usr/bin/ldapsearch  -h  $LDAPHOST -b $LDAPBASEDN -D "$LDAPBINDDN" -w $LDAPBINDPASSWORD s sub "(objectClass=$OBJCLASS)" -E pr=1000/noprompt 
