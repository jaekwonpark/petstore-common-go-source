/*
 * Generated file common/v1.a1/response/response_model.go.  Product version: 1.0.0-RELEASE
 *
 * Part of the Common API and Schema definitions
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */
package response
import (
  "errors"
  "fmt"
  "encoding/json"
  import1 "common/v1.a1/config"
)

type ApiLink struct {
  ObjectType string `json:"object_type,omitempty"`
  Href *string `json:"href,omitempty"`
  Rel *string `json:"rel,omitempty"`
}
func (p *ApiLink) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.v1.a1.response.ApiLink"
  }
  return p.ObjectType
}
func (p *ApiLink) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

type ApiResponseMetadata struct {
  ObjectType string `json:"object_type,omitempty"`
  Flags []import1.Flag `json:"flags,omitempty"`
  Links []ApiLink `json:"links,omitempty"`
  Messages []import1.Message `json:"messages,omitempty"`
}
func (p *ApiResponseMetadata) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.v1.a1.response.ApiResponseMetadata"
  }
  return p.ObjectType
}
func (p *ApiResponseMetadata) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

