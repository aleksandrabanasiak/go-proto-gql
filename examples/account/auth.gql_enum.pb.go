// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/account/auth.proto

package account // import "github.com/danielvladco/go-proto-gql/examples/account"

import fmt "fmt"
import io "io"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/danielvladco/go-proto-gql"
import _ "github.com/mwitkow/go-proto-validators@v0.0.0-20180403085117-0950a7990007"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (c *Roles1) UnmarshalGQL(v interface{}) error {
	code, ok := v.(string)
	if ok {
		*c = Roles1(Roles1_value[code])
		return nil
	}
	return fmt.Errorf("cannot unmarshal Roles1 enum")
}

func (c Roles1) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, "%q", c.String())
}
