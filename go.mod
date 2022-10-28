module sylr.dev/fix

go 1.18

require (
	filippo.io/age v1.0.0
	github.com/dustin/go-humanize v1.0.0
	github.com/google/uuid v1.3.0
	github.com/iancoleman/strcase v0.2.0
	github.com/lib/pq v1.10.6
	github.com/mattn/go-sqlite3 v1.14.14
	github.com/nats-io/nats-server/v2 v2.8.4
	github.com/nats-io/nats.go v1.16.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/prometheus/client_golang v1.13.0
	github.com/quickfixgo/enum v0.0.0-20210629025633-9afc8539baba
	github.com/quickfixgo/field v0.0.0-20171007195410-74cea5ec78c7
	github.com/quickfixgo/fixt11 v0.0.0-20171007213433-d9788ca97f5d
	github.com/quickfixgo/quickfix v0.6.1-0.20210618140103-31f5ebe90229
	github.com/quickfixgo/tag v0.0.0-20171007194743-cbb465760521
	github.com/rs/zerolog v1.27.0
	github.com/shopspring/decimal v1.3.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/sylr/quickfixgo-fix50sp2/marketdataincrementalrefresh v0.0.0-20220401195242-281940b8a21e
	github.com/sylr/quickfixgo-fix50sp2/marketdatasnapshotfullrefresh v0.0.0-20220401195242-281940b8a21e
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	golang.org/x/term v0.0.0-20220526004731-065cf7ba2467
	sylr.dev/yaml/age/v3 v3.0.0-20220527135827-28ffff5246ba
	sylr.dev/yaml/v3 v3.0.0-20220527135632-500fddf2b049
)

replace (
	github.com/quickfixgo/enum => github.com/sylr/quickfixgo-enum v0.0.0-20220401193143-29a559514373
	github.com/quickfixgo/field => github.com/sylr/quickfixgo-field v0.0.0-20220401193046-ca4cd16301d2
	github.com/quickfixgo/quickfix => github.com/sylr/quickfix-go v0.6.1-0.20221028155147-da1f7761ba49
	github.com/quickfixgo/tag => github.com/sylr/quickfixgo-tag v0.0.0-20220401193001-96cf7367fdfa
)

require (
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/armon/go-proxyproto v0.0.0-20210323213023-7e956b284f0a // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/klauspost/compress v1.15.6 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.2.1-0.20220330180145-442af02fd36a // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/smartystreets/assertions v1.13.0 // indirect
	golang.org/x/net v0.0.0-20220708220712-1185a9018129 // indirect
	golang.org/x/sys v0.0.0-20220615213510-4f61da869c0c // indirect
	golang.org/x/time v0.0.0-20220609170525-579cf78fd858 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
