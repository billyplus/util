package yi

//Gua 一次算卦
type Gua struct {
	GuaXiang SixXiang `json:"guaxiang"` //卦象
	BenGua   *Trigram `json:"bengua"`   //本卦
	FirstGua *Trigram `json:"firstgua"` //首卦
	Changed  [6]uint  `json:"changed"`  //是否有变卦
	BianGua  *Trigram `json:"biangua"`  //变卦
}

//NewGua 新建一卦
func NewGua(guaXiang SixXiang) *Gua {
	gua := new(Gua)
	gua.GuaXiang = guaXiang
	gua.BenGua = NewBenGua(guaXiang)
	hasBianGua := false
	for i := 0; i < 6; i++ {
		if guaXiang[i] > 1 {
			hasBianGua = true
			gua.Changed[i] = 1
		}
	}
	if hasBianGua {
		gua.BianGua = NewBianGua(guaXiang)
	}
	return gua
}
