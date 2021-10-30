package biz

// UserDevice 用户设备
type UserDevice struct {
	BaseModel
	Uid           int64   `gorm:"index;not null;comment:用户id"`
	DeviceType    *int32  `gorm:"not null;comment:用户设备类型(1-Android,2-IOS,3-Windows,4-MacOS,5-Linux,6-Web"`
	DeviceBrand   *string `gorm:"size:50;comment:设备品牌"`
	DeviceModel   *string `gorm:"size:50;comment:设备型号"`
	DeviceVersion *string `gorm:"size:50;comment:设备系统版本"`
	SdkVersion    *string `gorm:"size:30;comment:软件版本"`
	ConnStatus    *bool   `gorm:"index;not null;comment:设备是否在线"`
	ServerAddr    *string `gorm:"index;not null;comment:连接服务器地址"`
}
