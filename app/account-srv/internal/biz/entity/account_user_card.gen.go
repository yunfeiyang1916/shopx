// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameAccountUserCard = "account_user_card"

// AccountUserCard 用户会员卡
type AccountUserCard struct {
	CardID              int32   `gorm:"column:card_id;primaryKey;autoIncrement:true;comment:会员卡编号" json:"card_id"`                          // 会员卡编号
	CardName            string  `gorm:"column:card_name;not null;comment:会员卡名称" json:"card_name"`                                           // 会员卡名称
	UserLevelID         int32   `gorm:"column:user_level_id;not null;comment:会员等级" json:"user_level_id"`                                    // 会员等级
	CardEnable          []uint8 `gorm:"column:card_enable;not null;default:b'0';comment:状态(BOOL):0-禁用;1-开启" json:"card_enable"`             // 状态(BOOL):0-禁用;1-开启
	CardPirce           float64 `gorm:"column:card_pirce;not null;default:0.00;comment:会费（元/期）" json:"card_pirce"`                          // 会费（元/期）
	CardMarketPrice     float64 `gorm:"column:card_market_price;not null;default:0.00;comment:原会费（元/期）" json:"card_market_price"`           // 原会费（元/期）
	CardSaveAmount      float64 `gorm:"column:card_save_amount;not null;default:0.00;comment:预计年省金额" json:"card_save_amount"`               // 预计年省金额
	CardYear            int32   `gorm:"column:card_year;not null;comment:有效期年" json:"card_year"`                                            // 有效期年
	CardMonth           int32   `gorm:"column:card_month;not null;comment:有效期月" json:"card_month"`                                          // 有效期月
	CardDay             int32   `gorm:"column:card_day;not null;comment:有效期日" json:"card_day"`                                              // 有效期日
	CardTime            int64   `gorm:"column:card_time;not null;comment:上线时间" json:"card_time"`                                            // 上线时间
	CardSort            int32   `gorm:"column:card_sort;not null;default:50;comment:排序" json:"card_sort"`                                   // 排序
	UserTypeID          int32   `gorm:"column:user_type_id;not null;default:1;comment:会员类型(ENUM):1-普通用户;2-扩展用户" json:"user_type_id"`        // 会员类型(ENUM):1-普通用户;2-扩展用户
	TransportTypeID     int32   `gorm:"column:transport_type_id;not null;comment:物流模板编号" json:"transport_type_id"`                          // 物流模板编号
	CardBindAddr        []uint8 `gorm:"column:card_bind_addr;not null;default:b'0';comment:唯一地址(BOOL):0-否;1-是" json:"card_bind_addr"`       // 唯一地址(BOOL):0-否;1-是
	CardIsExp           []uint8 `gorm:"column:card_is_exp;not null;default:b'0';comment:体验卡(BOOL):0-否;1-是" json:"card_is_exp"`              // 体验卡(BOOL):0-否;1-是
	CardExpMonth        int32   `gorm:"column:card_exp_month;not null;comment:体验月" json:"card_exp_month"`                                   // 体验月
	CardExpDay          int32   `gorm:"column:card_exp_day;not null;comment:体验日" json:"card_exp_day"`                                       // 体验日
	CardExpFee          float64 `gorm:"column:card_exp_fee;not null;default:0.00;comment:体验费用" json:"card_exp_fee"`                         // 体验费用
	CardExpStatus       []uint8 `gorm:"column:card_exp_status;not null;default:b'0';comment:体验状态(BOOL):0-禁用;1-开启" json:"card_exp_status"`   // 体验状态(BOOL):0-禁用;1-开启
	StoreID             int32   `gorm:"column:store_id;not null;comment:店铺编号" json:"store_id"`                                              // 店铺编号
	ActivityIDDiscount  int32   `gorm:"column:activity_id_discount;not null;comment:活动会员专享折扣编号" json:"activity_id_discount"`                // 活动会员专享折扣编号
	ActivityIDGitfbag   int32   `gorm:"column:activity_id_gitfbag;not null;comment:活动会员开卡赠品编号" json:"activity_id_gitfbag"`                  // 活动会员开卡赠品编号
	ActivityIDMemberDay int32   `gorm:"column:activity_id_member_day;not null;comment:活动会员日（实物）编号" json:"activity_id_member_day"`           // 活动会员日（实物）编号
	ActivityIDBirthday  int32   `gorm:"column:activity_id_birthday;not null;comment:活动生日专享编号" json:"activity_id_birthday"`                  // 活动生日专享编号
	UserFxEnable        []uint8 `gorm:"column:user_fx_enable;not null;default:b'0';comment:允许分销(BOOL):1-启用分销;0-禁用分销" json:"user_fx_enable"` // 允许分销(BOOL):1-启用分销;0-禁用分销
	UserFxRate          float64 `gorm:"column:user_fx_rate;not null;default:0.00;comment:分销佣金比例" json:"user_fx_rate"`                       // 分销佣金比例
	CardTitle           string  `gorm:"column:card_title;not null;comment:分享标题" json:"card_title"`                                          // 分享标题
	CardDesc            string  `gorm:"column:card_desc;not null;comment:分享描述" json:"card_desc"`                                            // 分享描述
	CardKeywords        string  `gorm:"column:card_keywords;not null;comment:搜索关键字" json:"card_keywords"`                                   // 搜索关键字
	CardImg             string  `gorm:"column:card_img;not null;comment:分享图标" json:"card_img"`                                              // 分享图标
	CardFxActive        []uint8 `gorm:"column:card_fx_active;not null;default:b'0';comment:开通推广员(BOOL):1-开通;0-不开通" json:"card_fx_active"`   // 开通推广员(BOOL):1-开通;0-不开通
	CardImgBg           string  `gorm:"column:card_img_bg;not null;comment:背景图片" json:"card_img_bg"`                                        // 背景图片
}

// TableName AccountUserCard's table name
func (*AccountUserCard) TableName() string {
	return TableNameAccountUserCard
}
