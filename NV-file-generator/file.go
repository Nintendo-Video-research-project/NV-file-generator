package main

type Header struct {
	Start     [4]byte
	End       [4]byte
	met_start [4]byte
	met_end   [4]byte
	vid_thumb uint32
	___       uint8 // Padding?
	Meta      Metadata
}
type Metadata struct {
	Meta_len   uint32
	Vid_id     uint64
	Release    Timestamp
	Expire     Timestamp
	vid_title  uint16
	Vid_len    uint16
	vid_desc   uint16
	Link_IDs   uint64
	moflex     uint64
	thumb_size uint16
	int_dt_sz  Int_link
}
type Timestamp struct {
	Year    uint16
	Month   uint8
	Day     uint8
	Hours   uint8
	Minutes uint8
	Seconds uint8
	___     uint8 // Padding?
}
type Int_header struct {
	link_num  uint8
	link_addr string
	addr_add  string
}
type Int_link struct {
	meta       uint8
	IlID       uint64
	Unknown    [4]byte
	url_addr   string
	BLC        uint32
	button_li  uint16
	thumb_len  uint16
	thumb_size uint16
}
