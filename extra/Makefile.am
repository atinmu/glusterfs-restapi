EXTRA_DIST = restconfig.json.in glusterrestd.service.in

restconfdir = $(sysconfdir)/glusterfs
restconf_DATA = restconfig.json

CLEANFILES = restconfig.json glusterrestd.service

SYSTEMD_DIR = @systemddir@

install-exec-local:
	@if [ -d $(SYSTEMD_DIR) ]; then \
		$(mkdir_p) $(DESTDIR)$(SYSTEMD_DIR); \
		$(INSTALL_PROGRAM) glusterrestd.service $(DESTDIR)$(SYSTEMD_DIR)/; \
	fi
