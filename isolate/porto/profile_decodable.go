package porto

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Profile) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "registry":
			z.Registry, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Registry")
				return
			}
		case "repository":
			z.Repository, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Repository")
				return
			}
		case "network_mode":
			z.NetworkMode, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "NetworkMode")
				return
			}
		case "network":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Network")
				return
			}
			if z.Network == nil {
				z.Network = make(map[string]string, zb0002)
			} else if len(z.Network) > 0 {
				for key := range z.Network {
					delete(z.Network, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 string
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Network")
					return
				}
				za0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Network", za0001)
					return
				}
				z.Network[za0001] = za0002
			}
		case "cwd":
			z.Cwd, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Cwd")
				return
			}
		case "binds":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Binds")
				return
			}
			if cap(z.Binds) >= int(zb0003) {
				z.Binds = (z.Binds)[:zb0003]
			} else {
				z.Binds = make([]string, zb0003)
			}
			for za0003 := range z.Binds {
				z.Binds[za0003], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Binds", za0003)
					return
				}
			}
		case "container":
			var zb0004 uint32
			zb0004, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Container")
				return
			}
			if z.Container == nil {
				z.Container = make(map[string]string, zb0004)
			} else if len(z.Container) > 0 {
				for key := range z.Container {
					delete(z.Container, key)
				}
			}
			for zb0004 > 0 {
				zb0004--
				var za0004 string
				var za0005 string
				za0004, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Container")
					return
				}
				za0005, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Container", za0004)
					return
				}
				z.Container[za0004] = za0005
			}
		case "volume":
			var zb0005 uint32
			zb0005, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Volume")
				return
			}
			if z.Volume == nil {
				z.Volume = make(map[string]string, zb0005)
			} else if len(z.Volume) > 0 {
				for key := range z.Volume {
					delete(z.Volume, key)
				}
			}
			for zb0005 > 0 {
				zb0005--
				var za0006 string
				var za0007 string
				za0006, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Volume")
					return
				}
				za0007, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Volume", za0006)
					return
				}
				z.Volume[za0006] = za0007
			}
		case "extravolumes":
			var zb0006 uint32
			zb0006, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "ExtraVolumes")
				return
			}
			if cap(z.ExtraVolumes) >= int(zb0006) {
				z.ExtraVolumes = (z.ExtraVolumes)[:zb0006]
			} else {
				z.ExtraVolumes = make([]VolumeProfile, zb0006)
			}
			for za0008 := range z.ExtraVolumes {
				err = z.ExtraVolumes[za0008].DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "ExtraVolumes", za0008)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Profile) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "registry"
	err = en.Append(0x89, 0xa8, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.Registry)
	if err != nil {
		err = msgp.WrapError(err, "Registry")
		return
	}
	// write "repository"
	err = en.Append(0xaa, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79)
	if err != nil {
		return
	}
	err = en.WriteString(z.Repository)
	if err != nil {
		err = msgp.WrapError(err, "Repository")
		return
	}
	// write "network_mode"
	err = en.Append(0xac, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.NetworkMode)
	if err != nil {
		err = msgp.WrapError(err, "NetworkMode")
		return
	}
	// write "network"
	err = en.Append(0xa7, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Network)))
	if err != nil {
		err = msgp.WrapError(err, "Network")
		return
	}
	for za0001, za0002 := range z.Network {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Network")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Network", za0001)
			return
		}
	}
	// write "cwd"
	err = en.Append(0xa3, 0x63, 0x77, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Cwd)
	if err != nil {
		err = msgp.WrapError(err, "Cwd")
		return
	}
	// write "binds"
	err = en.Append(0xa5, 0x62, 0x69, 0x6e, 0x64, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Binds)))
	if err != nil {
		err = msgp.WrapError(err, "Binds")
		return
	}
	for za0003 := range z.Binds {
		err = en.WriteString(z.Binds[za0003])
		if err != nil {
			err = msgp.WrapError(err, "Binds", za0003)
			return
		}
	}
	// write "container"
	err = en.Append(0xa9, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Container)))
	if err != nil {
		err = msgp.WrapError(err, "Container")
		return
	}
	for za0004, za0005 := range z.Container {
		err = en.WriteString(za0004)
		if err != nil {
			err = msgp.WrapError(err, "Container")
			return
		}
		err = en.WriteString(za0005)
		if err != nil {
			err = msgp.WrapError(err, "Container", za0004)
			return
		}
	}
	// write "volume"
	err = en.Append(0xa6, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Volume)))
	if err != nil {
		err = msgp.WrapError(err, "Volume")
		return
	}
	for za0006, za0007 := range z.Volume {
		err = en.WriteString(za0006)
		if err != nil {
			err = msgp.WrapError(err, "Volume")
			return
		}
		err = en.WriteString(za0007)
		if err != nil {
			err = msgp.WrapError(err, "Volume", za0006)
			return
		}
	}
	// write "extravolumes"
	err = en.Append(0xac, 0x65, 0x78, 0x74, 0x72, 0x61, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.ExtraVolumes)))
	if err != nil {
		err = msgp.WrapError(err, "ExtraVolumes")
		return
	}
	for za0008 := range z.ExtraVolumes {
		err = z.ExtraVolumes[za0008].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "ExtraVolumes", za0008)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Profile) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "registry"
	o = append(o, 0x89, 0xa8, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79)
	o = msgp.AppendString(o, z.Registry)
	// string "repository"
	o = append(o, 0xaa, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79)
	o = msgp.AppendString(o, z.Repository)
	// string "network_mode"
	o = append(o, 0xac, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65)
	o = msgp.AppendString(o, z.NetworkMode)
	// string "network"
	o = append(o, 0xa7, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b)
	o = msgp.AppendMapHeader(o, uint32(len(z.Network)))
	for za0001, za0002 := range z.Network {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	// string "cwd"
	o = append(o, 0xa3, 0x63, 0x77, 0x64)
	o = msgp.AppendString(o, z.Cwd)
	// string "binds"
	o = append(o, 0xa5, 0x62, 0x69, 0x6e, 0x64, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Binds)))
	for za0003 := range z.Binds {
		o = msgp.AppendString(o, z.Binds[za0003])
	}
	// string "container"
	o = append(o, 0xa9, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72)
	o = msgp.AppendMapHeader(o, uint32(len(z.Container)))
	for za0004, za0005 := range z.Container {
		o = msgp.AppendString(o, za0004)
		o = msgp.AppendString(o, za0005)
	}
	// string "volume"
	o = append(o, 0xa6, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65)
	o = msgp.AppendMapHeader(o, uint32(len(z.Volume)))
	for za0006, za0007 := range z.Volume {
		o = msgp.AppendString(o, za0006)
		o = msgp.AppendString(o, za0007)
	}
	// string "extravolumes"
	o = append(o, 0xac, 0x65, 0x78, 0x74, 0x72, 0x61, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.ExtraVolumes)))
	for za0008 := range z.ExtraVolumes {
		o, err = z.ExtraVolumes[za0008].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "ExtraVolumes", za0008)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Profile) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "registry":
			z.Registry, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Registry")
				return
			}
		case "repository":
			z.Repository, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Repository")
				return
			}
		case "network_mode":
			z.NetworkMode, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "NetworkMode")
				return
			}
		case "network":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Network")
				return
			}
			if z.Network == nil {
				z.Network = make(map[string]string, zb0002)
			} else if len(z.Network) > 0 {
				for key := range z.Network {
					delete(z.Network, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Network")
					return
				}
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Network", za0001)
					return
				}
				z.Network[za0001] = za0002
			}
		case "cwd":
			z.Cwd, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Cwd")
				return
			}
		case "binds":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Binds")
				return
			}
			if cap(z.Binds) >= int(zb0003) {
				z.Binds = (z.Binds)[:zb0003]
			} else {
				z.Binds = make([]string, zb0003)
			}
			for za0003 := range z.Binds {
				z.Binds[za0003], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Binds", za0003)
					return
				}
			}
		case "container":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Container")
				return
			}
			if z.Container == nil {
				z.Container = make(map[string]string, zb0004)
			} else if len(z.Container) > 0 {
				for key := range z.Container {
					delete(z.Container, key)
				}
			}
			for zb0004 > 0 {
				var za0004 string
				var za0005 string
				zb0004--
				za0004, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Container")
					return
				}
				za0005, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Container", za0004)
					return
				}
				z.Container[za0004] = za0005
			}
		case "volume":
			var zb0005 uint32
			zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Volume")
				return
			}
			if z.Volume == nil {
				z.Volume = make(map[string]string, zb0005)
			} else if len(z.Volume) > 0 {
				for key := range z.Volume {
					delete(z.Volume, key)
				}
			}
			for zb0005 > 0 {
				var za0006 string
				var za0007 string
				zb0005--
				za0006, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Volume")
					return
				}
				za0007, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Volume", za0006)
					return
				}
				z.Volume[za0006] = za0007
			}
		case "extravolumes":
			var zb0006 uint32
			zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ExtraVolumes")
				return
			}
			if cap(z.ExtraVolumes) >= int(zb0006) {
				z.ExtraVolumes = (z.ExtraVolumes)[:zb0006]
			} else {
				z.ExtraVolumes = make([]VolumeProfile, zb0006)
			}
			for za0008 := range z.ExtraVolumes {
				bts, err = z.ExtraVolumes[za0008].UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "ExtraVolumes", za0008)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Profile) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Registry) + 11 + msgp.StringPrefixSize + len(z.Repository) + 13 + msgp.StringPrefixSize + len(z.NetworkMode) + 8 + msgp.MapHeaderSize
	if z.Network != nil {
		for za0001, za0002 := range z.Network {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += 4 + msgp.StringPrefixSize + len(z.Cwd) + 6 + msgp.ArrayHeaderSize
	for za0003 := range z.Binds {
		s += msgp.StringPrefixSize + len(z.Binds[za0003])
	}
	s += 10 + msgp.MapHeaderSize
	if z.Container != nil {
		for za0004, za0005 := range z.Container {
			_ = za0005
			s += msgp.StringPrefixSize + len(za0004) + msgp.StringPrefixSize + len(za0005)
		}
	}
	s += 7 + msgp.MapHeaderSize
	if z.Volume != nil {
		for za0006, za0007 := range z.Volume {
			_ = za0007
			s += msgp.StringPrefixSize + len(za0006) + msgp.StringPrefixSize + len(za0007)
		}
	}
	s += 13 + msgp.ArrayHeaderSize
	for za0008 := range z.ExtraVolumes {
		s += z.ExtraVolumes[za0008].Msgsize()
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *VolumeProfile) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "target":
			z.Target, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Target")
				return
			}
		case "properties":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Properties")
				return
			}
			if z.Properties == nil {
				z.Properties = make(map[string]string, zb0002)
			} else if len(z.Properties) > 0 {
				for key := range z.Properties {
					delete(z.Properties, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 string
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Properties")
					return
				}
				za0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Properties", za0001)
					return
				}
				z.Properties[za0001] = za0002
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *VolumeProfile) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "target"
	err = en.Append(0x82, 0xa6, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74)
	if err != nil {
		return
	}
	err = en.WriteString(z.Target)
	if err != nil {
		err = msgp.WrapError(err, "Target")
		return
	}
	// write "properties"
	err = en.Append(0xaa, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Properties)))
	if err != nil {
		err = msgp.WrapError(err, "Properties")
		return
	}
	for za0001, za0002 := range z.Properties {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Properties")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Properties", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *VolumeProfile) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "target"
	o = append(o, 0x82, 0xa6, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74)
	o = msgp.AppendString(o, z.Target)
	// string "properties"
	o = append(o, 0xaa, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73)
	o = msgp.AppendMapHeader(o, uint32(len(z.Properties)))
	for za0001, za0002 := range z.Properties {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *VolumeProfile) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "target":
			z.Target, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Target")
				return
			}
		case "properties":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Properties")
				return
			}
			if z.Properties == nil {
				z.Properties = make(map[string]string, zb0002)
			} else if len(z.Properties) > 0 {
				for key := range z.Properties {
					delete(z.Properties, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Properties")
					return
				}
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Properties", za0001)
					return
				}
				z.Properties[za0001] = za0002
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *VolumeProfile) Msgsize() (s int) {
	s = 1 + 7 + msgp.StringPrefixSize + len(z.Target) + 11 + msgp.MapHeaderSize
	if z.Properties != nil {
		for za0001, za0002 := range z.Properties {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	return
}
