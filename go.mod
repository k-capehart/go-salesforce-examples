module main

go 1.25.0

require github.com/k-capehart/go-salesforce/v3 v3.0.0-20251003161401-e3baa43d18c5

require (
	github.com/forcedotcom/go-soql v0.0.0-20220705175410-00f698360bee // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.0 // indirect
	github.com/jszwec/csvutil v1.10.0 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	golang.org/x/text v0.28.0 // indirect
	k8s.io/apimachinery v0.34.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/utils v0.0.0-20250604170112-4c0f3b243397 // indirect
)

replace github.com/k-capehart/go-salesforce/v3 => ../go-salesforce/
