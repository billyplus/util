package yi

//Trigram 六十四卦
type Trigram struct {
	ID           uint           `json:"id"`
	Gong         BaGong         `json:"gong"`        //所属八宫
	IndexInGroup uint           `json:"indexingong"` //在宫内的序号
	GongElement  Elements       `json:"ge"`          //所属八宫的五行
	Yaos         *SixYao        `json:"yaos"`        //六爻 0、1对应阳、阴
	Trigrams     [2]BaGua       `json:"trigrams"`    //内外卦
	YaoElement   [6]Elements    `json:"ye"`          //水一，火二，木三，金四，土五
	TG           [6]TianGan     `json:"tg"`          //天干
	DZ           [6]DiZhi       `json:"dz"`          //地支
	Relation     [6]SixRelation `json:"relation"`    //六亲
	ShiIndex     uint           `json:"shi"`         //世爻
	YingIndex    uint           `json:"ying"`        //应爻
	Body         DiZhi          `json:"body"`        //卦身
}

//NewTrigram 根据id算出新卦
func NewTrigram(id uint) *Trigram {
	tri := new(Trigram)

	return tri
}

//NewBenGua 根据卦象出本卦
func NewBenGua(guaxiang SixXiang) *Trigram {
	tri := new(Trigram)
	yaos := new(SixYao)

	//根据卦算排本卦六爻阴阳
	for i := 0; i < 6; i++ {
		switch guaxiang[i] {
		case GXDan, GXChong:
			yaos[i] = YangYao
		case GXChai, GXJiao:
			yaos[i] = YingYao
		default:
			yaos[i] = YangYao
		}
	}
	tri.Yaos = yaos
	tri.translateTrigram()
	return tri
}

//NewBianGua 根据卦象出变卦
func NewBianGua(guaxiang SixXiang) *Trigram {
	tri := new(Trigram)
	yaos := new(SixYao)
	//根据卦算排本卦六爻阴阳
	for i := 0; i < 6; i++ {
		switch guaxiang[i] {
		case GXDan, GXJiao:
			yaos[i] = YangYao
		case GXChai, GXChong:
			yaos[i] = YingYao
		default:
			yaos[i] = YangYao
		}
	}
	tri.Yaos = yaos
	tri.translateTrigram()
	return tri
}

//解析六爻
func (t *Trigram) translateTrigram() {
	t.calcIndex() //算序号
	t.calcTGDZ()  //算天干地支
	t.GongElement = GroupElements[t.Gong]
	for i := 0; i < 6; i++ {
		t.Relation[i] = GetRelation(t.GongElement, t.YaoElement[i])
	}
}

//计算该卦的序号ID、所在宫groupID、在宫内的序号IndexInGroup
func (t *Trigram) calcIndex() {
	var yaos = t.Yaos
	guaindex := uint(0)
	for i := uint(0); i < 6; i++ {
		guaindex += uint(yaos[i] << (5 - i))
	}

	t.ID = guaindex

	if t.ID%9 == 0 {
		t.Gong = BaGongList[t.ID/9]
		t.IndexInGroup = 0
	} else {
		for i, num := range TrigramGroupParams {
			x := uint(t.ID) ^ num
			if x%9 == 0 {
				t.Gong = BaGongList[x/9]
				t.IndexInGroup = uint(i + 1)
				break
			}
		}
	}

	t.ShiIndex = ShiIndexParams[t.IndexInGroup]
	t.YingIndex = YingIndexParams[t.IndexInGroup]
}

//计算天干地支和五行
func (t *Trigram) calcTGDZ() {
	var yaos = t.Yaos
	for i := uint(0); i < 2; i++ {
		tri := int(yaos[i*3]<<2 + yaos[i*3+1]<<1 + yaos[i*3+2])
		t.Trigrams[i] = BaGua(tri)
		for j := uint(0); j < 3; j++ {
			k := i*3 + j
			t.TG[k] = LiuYaoBaGuaTianGan[tri][i]
			t.DZ[k] = LiuYaoBaGuaDiZhi[tri][k]
			t.YaoElement[k] = DiZhiElements[t.DZ[k]]
			//t.Index += yaos[k] << (5 - k)
		}
	}
}

func (t *Trigram) calcBody() {
	yaos := t.Yaos
	s := yaos[t.ShiIndex]
	var start uint
	if s == YingYao {
		start = uint(DZWu)
	} else {
		start = uint(DZZi)
	}
	body := DiZhi(start + t.ShiIndex - 1)
	t.Body = body
}
