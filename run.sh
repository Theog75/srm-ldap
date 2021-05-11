export LDAPBINDDN=$(cat /ldapauth/binddn)
export LDAPBINDPASSWORD=$(cat /ldapauth/password)
export LDAPBASEDN=$(cat /ldapauth/basedn)
export LDAPHOST=$(cat /ldapauth/ldaphost)

/srm-ldap