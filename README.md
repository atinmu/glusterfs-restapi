# GlusterFS REST APIs

## Install

Golang is required to build glusterrest.

	git clone https://github.com/aravindavk/glusterfs-restapi.git
	./autogen.sh
	./configure
	make goget
	make
	make install

## Usage

Use gluster-rest command to enable/disable REST Server

	gluster-rest enable|disable

Create/Manage Application using,

	gluster-rest app-add <APP_ID> <APP_SECRET>
	gluster-rest app-reset <APP_ID> <APP_SECRET>
	gluster-rest app-del <APP_ID>

## Configuration
By default rest server runs in port 8080, can be changed using config command,

	gluster-rest config-set port 9000
	gluster-rest config-set https_enabled true|false
	gluster-rest config-set auth_enabled true|false

Reset all configuration to defaults using,

	gluster-rest config-reset

Reset any specific config using,

	gluster-rest config-reset --name port

View all configurations,

	glusterrest config-get [--name=NAME]

Finally start the REST server,

	glusterrest start|stop|reload|status|restart


## Help
`man gluster-rest` for more details.
