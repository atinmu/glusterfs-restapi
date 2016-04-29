AC_INIT([glusterrest], m4_esyscmd([tr -d '\n' < VERSION]), [avishwan@redhat.com])
AC_PREFIX_DEFAULT(/usr/local)
AM_INIT_AUTOMAKE

#LT_INIT

case $host in
     *linux*)
        GLUSTERD_WORKDIR_TMP="${localstatedir}/lib/glusterd"
        ;;
     solaris*)
        GLUSTERD_WORKDIR_TMP="${localstatedir}/lib/glusterd"
        ;;
     *netbsd*)
        GLUSTERD_WORKDIR_TMP="${localstatedir}/db/glusterd"
        ;;
     *freebsd*)
        GLUSTERD_WORKDIR_TMP="${localstatedir}/db/glusterd"
        ;;
     darwin*)
        GLUSTERD_WORKDIR_TMP="${localstatedir}/db/glusterd"
        ;;
esac

VERSION=$(tr -d '\n' < VERSION)
AC_SUBST(VERSION)

dnl Fix this early so we can expand with eval later
test "${prefix}" = "NONE" && prefix="${ac_default_prefix}"
test "${exec_prefix}" = "NONE" && exec_prefix='${prefix}'

dnl Default value for sbindir
sbintemp="${sbindir}"
eval sbintemp=\"${sbintemp}\"
eval sbintemp=\"${sbintemp}\"
SBINDIR=${sbintemp}
AC_SUBST(SBINDIR)

sysconfdirtemp="${sysconfdir}"
eval sysconfdirtemp=\"${sysconfdirtemp}\"
SYSCONFDIR=${sysconfdirtemp}
AC_SUBST(SYSCONFDIR)

datadirtemp="${datadir}"
eval datadirtemp=\"${datadirtemp}\"
eval datadirtemp=\"${datadirtemp}\"
DATADIR=${datadirtemp}
AC_SUBST(DATADIR)

localstatedirtemp="${localstatedir}"
eval localstatedirtemp=\"${localstatedirtemp}\"
LOCALSTATEDIR=${localstatedirtemp}
AC_SUBST(LOCALSTATEDIR)

tmp_prefix="${prefix}"
prefix=""
eval GLUSTERD_WORKDIR_TMP=\"${GLUSTERD_WORKDIR_TMP}\"
GLUSTERD_WORKDIR=${GLUSTERD_WORKDIR_TMP}
prefix="${tmp_prefix}"
AC_SUBST(GLUSTERD_WORKDIR)

RESTAPI_GOPATH=$(dirname $(readlink -f $0))/restapi
RESTAPI_BUILD_ID=$(head -c20 /dev/urandom | od -An -tx1 | tr -d ' \n')

AC_SUBST(RESTAPI_GOPATH)
AC_SUBST(RESTAPI_BUILD_ID)

AC_ARG_WITH(systemddir,
            [  --with-systemddir=DIR systemd service files in DIR @<:@/usr/lib/systemd/system@:>@],
            [systemddir=$withval],
            [systemddir='/usr/lib/systemd/system'])
AC_SUBST(systemddir)

AC_ARG_WITH(initdir,
            [  --with-initdir=DIR init.d scripts in DIR @<:@/etc/init.d@:>@],
            [initdir=$withval],
            [initdir='/etc/init.d'])
AC_SUBST(initdir)


AC_CONFIG_FILES([doc/Makefile
		 extra/Makefile
		 extra/restconfig.json
		 extra/glusterrestd.service
		 Makefile
		 src/Makefile
		 src/gluster/Makefile
		 src/gluster/rest/Makefile
		 src/gluster/rest/vars.go
		 src/gluster/cli/Makefile
		 src/gluster/utils/Makefile
		 tools/Makefile
		 tools/gluster-rest.py
		 tools/peer_restcli.py
		 ])

AC_OUTPUT
