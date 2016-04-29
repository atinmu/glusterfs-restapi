SUBDIRS = src extra tools doc

install-data-hook:
	$(INSTALL) -d -m 755 $(DESTDIR)$(localstatedir)/log/glusterfs/rest

goget:
	@GO15VENDOREXPERIMENT=1 GOPATH=@RESTAPI_GOPATH@ \
		go get github.com/Sirupsen/logrus
	@GO15VENDOREXPERIMENT=1 GOPATH=@RESTAPI_GOPATH@ \
		go get github.com/dgrijalva/jwt-go
	@GO15VENDOREXPERIMENT=1 GOPATH=@RESTAPI_GOPATH@ \
		go get github.com/gorilla/handlers
	@GO15VENDOREXPERIMENT=1 GOPATH=@RESTAPI_GOPATH@ \
		go get github.com/gorilla/mux

