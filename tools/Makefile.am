EXTRA_DIST = gluster-rest.py.in peer_restcli.py.in

restdir = $(libexecdir)/glusterfs/restapi
rest_SCRIPTS = gluster-rest.py

restpeerdir = $(libexecdir)/glusterfs/
restpeer_SCRIPTS = peer_restcli.py

install-exec-hook:
	$(mkdir_p) $(DESTDIR)$(sbindir)
	rm -f $(DESTDIR)$(sbindir)/gluster-rest
	ln -s $(libexecdir)/glusterfs/rest/gluster-rest.py \
		$(DESTDIR)$(sbindir)/gluster-rest

uninstall-hook:
	rm -f $(DESTDIR)$(sbindir)/gluster-rest
