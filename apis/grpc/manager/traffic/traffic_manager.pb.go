//
// Copyright (C) 2019 Vdaas.org Vald team ( kpango, kou-m, rinx )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package traffic

import (
	fmt "fmt"
	math "math"

	_ "github.com/danielvladco/go-proto-gql/pb"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("traffic/traffic_manager.proto", fileDescriptor_fe442397473f4bc7) }

var fileDescriptor_fe442397473f4bc7 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x29, 0x4a, 0x4c,
	0x4b, 0xcb, 0x4c, 0xd6, 0x87, 0xd2, 0xf1, 0xb9, 0x89, 0x79, 0x89, 0xe9, 0xa9, 0x45, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0xfc, 0x68, 0xc2, 0x52, 0x32, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9,
	0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5,
	0x10, 0xe5, 0x52, 0x3c, 0x05, 0x49, 0xfa, 0xe9, 0x85, 0x39, 0x10, 0x9e, 0x93, 0xed, 0x89, 0x47,
	0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x18, 0xa5, 0x9f, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x96, 0x92, 0x98, 0x58, 0xac, 0x5f, 0x96, 0x98,
	0x93, 0x02, 0x32, 0xa8, 0x58, 0x3f, 0xbd, 0xa8, 0x20, 0x59, 0x1f, 0x6a, 0x05, 0xcc, 0x25, 0x49,
	0x6c, 0x60, 0x53, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0x8d, 0xc1, 0x22, 0xa3, 0x00,
	0x00, 0x00,
}
