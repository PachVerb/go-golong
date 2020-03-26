package conv

func Gconv(k Kilogram) Gram {
	return Gram(k * 1000) //声明的类型对应一个Gram方法，用于类型转换操作。 底层数据类型的运算操作类型底层类型结构一样
}

func Kconv(g Gram) Kilogram {
	return Kilogram(g / 1000)
}
