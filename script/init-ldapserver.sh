#!/bin/bash
# Author: Han Hao
# Email: 136698493@qq.com
# Only Run Once When You First Install OpenLDAP


function usage(){
    cat <<EOF
$0 -D <domain> -u <username> -s <password>
Example: $0 -D home.org -u admin -s 111111
EOF
}


while getopts "D:s:u:" opt;do
    case $opt in
        D) domain=$OPTARG
        ;;
        s) rootpw=$OPTARG 
        ;;
        u) cn=$OPTARG
        ;;
        *) echo "Invalid option: " $OPTARG
        ;;
    esac
done

# check args
[ -z $domain ] && usage && exit 1
[ -z $rootpw ] && usage && exit 1
[ -z $cn ] && usage && exit 1

# domain is home.org ==> olcSuffix is dc=home,dc=org
olcSuffix=`echo $domain |awk -F '.' '{print "dc="$1",dc="$2}'`
# olcRootDN cn=admin,dc=home,dc=org
olcRootDN="cn=$cn,$olcSuffix"
# get password to sha password
shaRootPW=`slappasswd -h {SHA} -s $rootpw`

cat <<EOF | ldapmodify -Y EXTERNAL -H ldapi:///
dn: olcDatabase={0}config,cn=config
changetype: modify
add: olcRootPW
olcRootPW: $shaRootPW
EOF

# add default schema
if [ ! -e /root/.openldap-shell ];then
    echo "Add default schema"
    ls /etc/openldap/schema/*.ldif | xargs -I {} ldapadd -Y EXTERNAL -H ldapi:/// -f {} > /dev/null
fi
# modify olcSuffix olcRootDN
cat <<EOF | ldapmodify -Y EXTERNAL -H ldapi:///
dn: olcDatabase={1}monitor,cn=config
changetype: modify
replace: olcAccess
olcAccess: {0} to * by dn.base="gidNumber=0+uidNumber=0,cn=peercred,cn=external,cn=auth" read by dn.base="$olcRootDN" read by * none 

dn: olcDatabase={2}hdb,cn=config
changetype: modify
replace: olcSuffix
olcSuffix: $olcSuffix

dn: olcDatabase={2}hdb,cn=config
changetype: modify
replace: olcRootDN
olcRootDN: $olcRootDN

dn: olcDatabase={2}hdb,cn=config
changetype: modify
add: olcRootPW
olcRootPW: $shaRootPW

dn: olcDatabase={2}hdb,cn=config
changetype: modify
add: olcAccess
olcAccess: {0}to attrs=userPassword,shadowLastChange by dn="$olcRootDN" write by anonymous auth by self write by * none
olcAccess: {1}to dn.base="" by * read
olcAccess: {2}to * by dn="$olcRootDN" write by * read
EOF

echo "Add $olcRootDN Manager"

# add user and group
echo "Add users and groups"
left_dc=`echo $olcSuffix |awk -F ',' '{print $1}'| awk -F '=' '{print $2}'`
cat << EOF | ldapadd -D "$olcRootDN" -H ldapi:/// -w $rootpw
dn: $olcSuffix
objectClass: top
objectClass: dcObject
objectclass: organization
o: root_ldap
dc: $left_dc

EOF

cat <<EOF
You olcSuffix is $olcSuffix
You olcRootDN is $olcRootDN
You olcRootPW is $rootpw
EOF
