package yi

var (
	//HeavenlyStems 天干
	HeavenlyStems = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

	//EarthlyBranches 地支
	EarthlyBranches = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

	//Zodiac 十二生肖
	Zodiac = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}

	//TrigramNames 八卦
	TrigramNames = []string{"乾", "兑", "离", "震", "巽", "坎", "艮", "坤"}
)

//BaGua 八卦
type BaGua uint

//八卦定义
const (
	QianGua BaGua = 0 + iota //乾
	DuiGua                   //兑
	LiGua                    //离
	ZhenGua                  //震
	XunGua                   //巽
	KanGua                   //坎
	GenGua                   //艮
	KunGua                   //坤
)

//BaGuaList 八卦列表
var BaGuaList = []BaGua{QianGua, DuiGua, LiGua, ZhenGua, XunGua, KanGua, GenGua, KunGua}

//BaGong 八宫
type BaGong uint

//八宫定义
const (
	QianGong BaGong = 0 + iota //乾
	DuiGong                    //兑
	LiGong                     //离
	ZhenGong                   //震
	XunGong                    //巽
	KanGong                    //坎
	GenGong                    //艮
	KunGong                    //坤
)

//BaGongList 八宫列表
var BaGongList = []BaGong{QianGong, DuiGong, LiGong, ZhenGong, XunGong, KanGong, GenGong, KunGong}

//DiZhi 十二地支
type DiZhi uint

//十二地支定义
const (
	DZZi   DiZhi = 0 + iota //子
	DZChou                  //丑
	DZYin                   //寅
	DZMao                   //卯
	DZChen                  //辰
	DZSi                    //巳
	DZWu                    //午
	DZWei                   //未
	DZShen                  //申
	DZYou                   //酉
	DZXu                    //戌
	DZHai                   //亥
)

//DiZhiList 地支列表
var DiZhiList = []DiZhi{DZZi, DZChou, DZYin, DZMao, DZChen, DZSi, DZWu, DZWei, DZShen, DZYou, DZXu, DZHai}

//TianGan 十天干
type TianGan uint

//十天干定义
const (
	TGJia  TianGan = 0 + iota //甲
	TGYi                      //乙
	TGBing                    //丙
	TGDing                    //丁
	TGWu                      //戊
	TGJi                      //己
	TGGeng                    //庚
	TGXin                     //辛
	TGRen                     //壬
	TGGui                     //癸
)

//TianGanList 十天干列表
var TianGanList = []TianGan{TGJia, TGYi, TGBing, TGDing, TGWu, TGJi, TGGeng, TGXin, TGRen, TGGui}

//Elements 五行元素
type Elements uint

//五行元素
const (
	Water Elements = 0 + iota //水
	Fire                      //火
	Wood                      //木
	Metal                     //金
	Earth                     //土
)

//ElementList 五行列表
var ElementList = []Elements{Water, Fire, Wood, Metal, Earth}

//SixRelation 六亲
type SixRelation uint

//六亲
const (
	Brother SixRelation = 0 + iota //兄弟
	Father                         //父母
	Child                          //子孙
	GuanGui                        //官鬼
	Wife                           //妻财
)

//SixRelationList 六亲列表
var SixRelationList = []SixRelation{Brother, Father, Child, GuanGui, Wife}

//Yao 六爻
type Yao uint

//阴爻、阳爻
const (
	YangYao Yao = 0 + iota //阳爻
	YingYao                //阴爻
)

//SixYao 对应初二三四五上---六爻
type SixYao [6]Yao

//Xiang 卦象
type Xiang uint

//单拆重交
const (
	GXDan   Xiang = 0 + iota //单
	GXChai                   //拆
	GXChong                  //重
	GXJiao                   //交
)

//XiaogList 卦象列表
var XiangList = []Xiang{GXDan, GXChai, GXChong, GXJiao}

//SixXiang 一组卦象，六次
type SixXiang [6]Xiang

// 0	1	2	3	4	5	6	7	8	9	10	11
// 子	丑	寅	卯	辰	巳	午	未	申	酉	戌	亥

//LiuYaoBaGuaDiZhi 六爻八卦每卦对应地支顺序
var LiuYaoBaGuaDiZhi = [8][6]DiZhi{
	{DZZi, DZYin, DZChen, DZWu, DZShen, DZXu},  //乾
	{DZSi, DZMao, DZChou, DZHai, DZYou, DZWei}, //兑
	{DZMao, DZChou, DZHai, DZYou, DZWei, DZSi}, //离
	{DZZi, DZYin, DZChen, DZWu, DZShen, DZXu},  //震
	{DZChou, DZHai, DZYou, DZWei, DZSi, DZMao}, //巽
	{DZYin, DZChen, DZWu, DZShen, DZXu, DZZi},  //坎
	{DZChen, DZWu, DZShen, DZXu, DZZi, DZYin},  //艮
	{DZWei, DZSi, DZMao, DZChou, DZHai, DZYou}, //坤
}

//LiuYaoBaGuaTianGan 六爻八卦每卦对应天干顺序
//	0	1	2	3	4	5	6	7	8	9
//	甲	乙	丙	丁	戊	己	庚	辛	壬	癸
var LiuYaoBaGuaTianGan = [8][2]TianGan{
	{TGJia, TGRen},   //乾
	{TGDing, TGDing}, //兑
	{TGJi, TGJi},     //离
	{TGGeng, TGGeng}, //震
	{TGXin, TGXin},   //巽
	{TGWu, TGWu},     //坎
	{TGBing, TGBing}, //艮
	{TGYi, TGGui},    //坤
}

//DiZhiElements 地支对应五行
//	0	1	2	3	4
//	水	火	木	金	土
var DiZhiElements = [12]Elements{Water, Earth, Wood, Wood, Earth, Fire, Fire, Earth, Metal, Metal, Earth, Water}

//TrigramGroupParams 用来算该卦属于哪一宫
//var TrigramGroupParams = []int{1, 3, 7, 15, 31, 23, 16}
var TrigramGroupParams = []uint{32, 48, 56, 60, 62, 58, 2}

//GroupElements 八宫所属五行
//	乾	兑	离	震	巽	坎	艮	坤
//	0	1	2	3	4
//	水	火	木	金	土
var GroupElements = []Elements{Metal, Metal, Fire, Wood, Wood, Water, Earth, Earth}

//ShiIndexParams 根据所在八宫的位置计算世爻所在
var ShiIndexParams = []uint{5, 0, 1, 2, 3, 4, 3, 2}

//YingIndexParams 根据所在八宫的位置计算应爻所在
var YingIndexParams = []uint{2, 3, 4, 5, 0, 1, 0, 5}

//ElementsRelations 五行相生相克
//	兄弟	父母	子孙	官鬼	妻财
//	比合	生我	我生	克我	我克
//	0		1		2		3		4
var ElementsRelations = [][]SixRelation{
	//水，    火，    木，    金，   土
	{Brother, GuanGui, Child, Wife, Father}, //水
	{Father, Child, Wife, Brother, GuanGui}, //火
	{Child, Brother, Father, GuanGui, Wife}, //木
	{GuanGui, Wife, Brother, Father, Child}, //金
	{Wife, Father, GuanGui, Child, Brother}, //土
}
