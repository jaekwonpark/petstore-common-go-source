/*
 * Generated file common/unversioned.a1/utils/utils_model.go.  Product version: 1.0.0-RELEASE
 *
 * Part of the Common API and Schema definitions
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */
package utils

type Namespace struct {
  ObjectType string `json:"object_type,omitempty"`
  Endpoints []string `json:"endpoints,omitempty"`
  ExtId *string `json:"ext_id,omitempty"`
  ModelDependencies []string `json:"model_dependencies,omitempty"`
  Models []string `json:"models,omitempty"`
  Name *string `json:"name,omitempty"`
  Version *string `json:"version,omitempty"`
}
func (p *Namespace) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.unversioned.a1.utils.Namespace"
  }
  return p.ObjectType
}
func (p *Namespace) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

type NamespaceList struct {
  ObjectType string `json:"object_type,omitempty"`
  Data *Namespace `json:"data,omitempty"`
}
func (p *NamespaceList) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.unversioned.a1.utils.NamespaceList"
  }
  return p.ObjectType
}
func (p *NamespaceList) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

