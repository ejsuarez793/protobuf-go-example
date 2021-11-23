module profobuf-example-go

replace github.com/ejsuarez793/protobuf-example-go/src/simplepb => /Users/enriquesuarez/dev/go/src/github.com/ejsuarez793/protobuf-example-go/src/simplepb

replace github.com/ejsuarez793/protobuf-example-go/src/enumpb => /Users/enriquesuarez/dev/go/src/github.com/ejsuarez793/protobuf-example-go/src/enumpb

replace github.com/ejsuarez793/protobuf-example-go/src/complexpb => /Users/enriquesuarez/dev/go/src/github.com/ejsuarez793/protobuf-example-go/src/complexpb

replace github.com/ejsuarez793/protobuf-example-go/src/addressbookpb => /Users/enriquesuarez/dev/go/src/github.com/ejsuarez793/protobuf-example-go/src/addressbookpb

go 1.17

require (
	github.com/ejsuarez793/protobuf-example-go/src/addressbookpb v0.0.0-00010101000000-000000000000
	github.com/ejsuarez793/protobuf-example-go/src/enumpb v0.0.0-00010101000000-000000000000
	github.com/ejsuarez793/protobuf-example-go/src/simplepb v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.0
)

require (
	github.com/ejsuarez793/protobuf-example-go/src/complexpb v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.27.1 // indirect
)
