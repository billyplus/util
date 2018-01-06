package yi

//GetRelation 五行相克关系
//			兄弟	父母	子孙	官鬼	妻财
//			比合	生我	我生	克我	我克
//返回值	0		1		2		3		4
func GetRelation(self Elements, target Elements) SixRelation {
	relation := ElementsRelations[self][target]
	return relation
}
