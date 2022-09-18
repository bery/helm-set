run:
	go build .
	helm plugin uninstall myplugin
	helm plugin install .