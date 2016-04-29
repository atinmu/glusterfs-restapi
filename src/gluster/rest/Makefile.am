# ldflags required to generate debuginfo packages
# Ref: http://fedoraproject.org/wiki/PackagingDrafts/Go#Debuginfo

sbin_PROGRAMS = glusterrestd
glusterrestd_SOURCES =

glusterrestd$(EXEEXT):
	@GO15VENDOREXPERIMENT=1 GOPATH=@DATADIR@/gocode:@RESTAPI_GOPATH@ \
		go build -ldflags "-B 0x@RESTAPI_BUILD_ID@" -v -a -x -o ./glusterrestd

CLEANFILES = glusterrestd vars.go

EXTRA_DIST = glusterrestd.go vars.go.in handlers_peers.go handlers_volumes.go \
	middleware_extra.go middleware_jwt.go routes.go
