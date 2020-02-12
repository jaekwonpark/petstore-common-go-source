/*
 * Generated file common/v1.a1/config/config_model.go.  Product version: 1.0.0-RELEASE
 *
 * Part of the Common API and Schema definitions
 *
 * (c) 2020 Nutanix Inc.  All rights reserved
 *
 */
package config

type AbstractModel struct {
  Reserved_ *string `json:"reserved_,omitempty"`
  ObjectType string `json:"object_type,omitempty"`
}
func (p *AbstractModel) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.v1.a1.config.AbstractModel"
  }
  return p.ObjectType
}
func (p *AbstractModel) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

type BaseEnum int

const(
    BASEENUM_UNKNOWN BaseEnum = 0
)

// returns the name of the enum given an ordinal number
func (e BaseEnum) name(index int) string {
    names := [...]string {
        "BASEENUM_UNKNOWN",
    }
    if index < 0 || index > len(names) {
       return "_UNKNOWN"
    }
    return names[index]
}
// returns the enum type given a string value
func (e BaseEnum) index(name string) BaseEnum {
   names := [...]string {
    "BASEENUM_UNKNOWN",
   }
   for idx := range names {
     if names[idx] == name {
        return BaseEnum(idx)
     }
   }

   return BASEENUM_UNKNOWN
}


type Flag struct {
  ObjectType string `json:"object_type,omitempty"`
  Name *string `json:"name,omitempty"`
  Value *bool `json:"value,omitempty"`
}
func (p *Flag) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.v1.a1.config.Flag"
  }
  return p.ObjectType
}
func (p *Flag) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

type KVStringPair struct {
  ObjectType string `json:"object_type,omitempty"`
  Name *string `json:"name,omitempty"`
  Value *string `json:"value,omitempty"`
}
func (p *KVStringPair) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.v1.a1.config.KVStringPair"
  }
  return p.ObjectType
}
func (p *KVStringPair) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

type Message struct {
  ObjectType string `json:"object_type,omitempty"`
  Code *string `json:"code,omitempty"`
  Locale *string `json:"locale,omitempty"`
  Message *string `json:"message,omitempty"`
  Severity *MessageSeverity `json:"severity,omitempty"`
}
func (p *Message) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.v1.a1.config.Message"
  }
  return p.ObjectType
}
func (p *Message) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

type MessageSeverity int

const(
    INFO MessageSeverity = 0
    WARNING MessageSeverity = 1
    ERROR MessageSeverity = 2
    MESSAGESEVERITY_UNKNOWN MessageSeverity = 3
)

// returns the name of the enum given an ordinal number
func (e MessageSeverity) name(index int) string {
    names := [...]string {
        "INFO",
        "WARNING",
        "ERROR",
        "MESSAGESEVERITY_UNKNOWN",
    }
    if index < 0 || index > len(names) {
       return "_UNKNOWN"
    }
    return names[index]
}
// returns the enum type given a string value
func (e MessageSeverity) index(name string) MessageSeverity {
   names := [...]string {
    "INFO",
    "WARNING",
    "ERROR",
    "MESSAGESEVERITY_UNKNOWN",
   }
   for idx := range names {
     if names[idx] == name {
        return MessageSeverity(idx)
     }
   }

   return MESSAGESEVERITY_UNKNOWN
}


type Messages []Message

type Page struct {
  ObjectType string `json:"object_type,omitempty"`
  Limit_ *int32 `json:"limit_,omitempty"`
  Page_ *int32 `json:"page_,omitempty"`
}
func (p *Page) GetObjectType() string {
  if "" == p.ObjectType {
    p.ObjectType = "common.v1.a1.config.Page"
  }
  return p.ObjectType
}
func (p *Page) SetObjectType(ObjectType string) {
  p.ObjectType = ObjectType
}

