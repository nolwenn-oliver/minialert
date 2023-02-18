module minialert

go 1.20

require minialert/cmd v0.0.0-00010101000000-000000000000

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	minialert/client v0.0.0-00010101000000-000000000000 // indirect
	minialert/server v0.0.0-00010101000000-000000000000 // indirect
)

replace minialert/cmd => ./cmd
replace minialert/server => ./server
replace minialert/client => ./client
