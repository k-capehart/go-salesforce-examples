module main

go 1.25.0

require github.com/k-capehart/go-salesforce/v2 v2.5.2

require (
	github.com/forcedotcom/go-soql v0.0.0-20220705175410-00f698360bee // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/jszwec/csvutil v1.10.0 // indirect
	github.com/spf13/afero v1.14.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	k8s.io/apimachinery v0.33.3 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738 // indirect
)

replace github.com/k-capehart/go-salesforce/v2 => ../go-salesforce/
