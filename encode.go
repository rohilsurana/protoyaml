package protoyaml

import (
	"github.com/ghodss/yaml"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Marshal writes the given proto.Message in JSON format using default options.
// Do not depend on the output being stable. It may change over time across
// different versions of the program.
func Marshal(m proto.Message) ([]byte, error) {
	return MarshalOptions{}.Marshal(m)
}

// MarshalOptions is a configurable YAML format marshaler.
type MarshalOptions struct {
	// Indent specifies the set of indentation characters to use in a multiline
	// formatted output such that every entry is preceded by Indent and
	// terminated by a newline. If non-empty, then Multiline is treated as true.
	// Indent can only be composed of space or tab characters.
	Indent string

	// UseProtoNames uses proto field name instead of lowerCamelCase name in YAML
	// field names.
	UseProtoNames bool

	// UseEnumNumbers emits enum values as numbers.
	UseEnumNumbers bool

	// EmitUnpopulated specifies whether to emit unpopulated fields. It does not
	// emit unpopulated oneof fields or unpopulated extension fields.
	// The YAML value emitted for unpopulated fields are as follows:
	//  ╔═══════╤════════════════════════════╗
	//  ║ YAML  │ Protobuf field             ║
	//  ╠═══════╪════════════════════════════╣
	//  ║ false │ proto3 boolean fields      ║
	//  ║ 0     │ proto3 numeric fields      ║
	//  ║ ""    │ proto3 string/bytes fields ║
	//  ║ null  │ proto2 scalar fields       ║
	//  ║ null  │ message fields             ║
	//  ║ []    │ list fields                ║
	//  ║ {}    │ map fields                 ║
	//  ╚═══════╧════════════════════════════╝
	EmitUnpopulated bool

	// Resolver is used for looking up types when expanding google.protobuf.Any
	// messages. If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver interface {
		protoregistry.ExtensionTypeResolver
		protoregistry.MessageTypeResolver
	}
}

// Marshal marshals the given proto.Message in the YAML format using options in
// MarshalOptions. Do not depend on the output being stable. It may change over
// time across different versions of the program.
func (o MarshalOptions) Marshal(m proto.Message) ([]byte, error) {
	jmo := protojson.MarshalOptions{
		Multiline:       false,
		Indent:          o.Indent,
		UseProtoNames:   o.UseProtoNames,
		UseEnumNumbers:  o.UseEnumNumbers,
		EmitUnpopulated: o.EmitUnpopulated,
		Resolver:        o.Resolver,
	}
	j, err := jmo.Marshal(m)
	if err != nil {
		return nil, err
	}
	return yaml.JSONToYAML(j)
}
