package protoyaml

import (
	"github.com/ghodss/yaml"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Unmarshal reads the given []byte into the given proto.Message.
// The provided message must be mutable (e.g., a non-nil pointer to a message).
func Unmarshal(b []byte, m proto.Message) error {
	return UnmarshalOptions{}.Unmarshal(b, m)
}

// UnmarshalOptions is a configurable YAML format parser.
type UnmarshalOptions struct {
	// If AllowPartial is set, input for messages that will result in missing
	// required fields will not return an error.
	AllowPartial bool

	// If DiscardUnknown is set, unknown fields are ignored.
	DiscardUnknown bool

	// Resolver is used for looking up types when unmarshaling
	// google.protobuf.Any messages or extension fields.
	// If nil, this defaults to using protoregistry.GlobalTypes.
	Resolver interface {
		protoregistry.MessageTypeResolver
		protoregistry.ExtensionTypeResolver
	}
}

// Unmarshal reads the given []byte and populates the given proto.Message
// using options in the UnmarshalOptions object.
// It will clear the message first before setting the fields.
// If it returns an error, the given message may be partially set.
// The provided message must be mutable (e.g., a non-nil pointer to a message).
func (o UnmarshalOptions) Unmarshal(b []byte, m proto.Message) error {
	j, err := yaml.YAMLToJSON(b)
	if err != nil {
		return err
	}
	juo := protojson.UnmarshalOptions{
		AllowPartial:   o.AllowPartial,
		DiscardUnknown: o.DiscardUnknown,
		Resolver:       o.Resolver,
	}
	return juo.Unmarshal(j, m)
}
