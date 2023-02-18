module minialert/cmd

go 1.20

require (
	github.com/spf13/cobra v1.6.1
	minialert/client v0.0.0-00010101000000-000000000000
	minialert/server v0.0.0-00010101000000-000000000000
)

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace minialert/server => ../server
replace minialert/client => ../client
