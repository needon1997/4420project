package suffixArray_test

//func TestName(t *testing.T) {
//	str := util.GenRandomStrRepeat(int(math.Pow(2, 20)), 26, 32)
//	sa2 := suffixArray.CreateSuffixArray(str)
//	//rlfmi := sa2.ToRLFMI()
//	mtfmi := sa2.ToWTFMI()
//	//sadcsa := gvcsa.MakeSADCSArray(*sa2)
//	//gvcsa := gvcsa.MakeGVCSArray(*sa2)
//	//fmt.Println(size.Of(str))
//	//fmt.Println(size.Of(sa2))
//	//fmt.Println(size.Of(rlfmi))
//	fmt.Println(size.Of(mtfmi))
//	//fmt.Println(size.Of(sadcsa))
//	//fmt.Println(size.Of(gvcsa))
//	for i := 0; i < 100000; i++ {
//		r1 := mtfmi.Search(str[i : i+20])
//		r2 := sa2.Search(str[i : i+20])
//		if r1 != r2 {
//			t.Error("w")
//		}
//	}
//}

//
//func TestCreateSuffixArray(t *testing.T) {
//	//str := "mmiissiissiippii"
//	buf, _ := ioutil.ReadFile("./test")
//	str := string(buf)
//	sa1 := suffixTree.NewSuffixTree(str).ToSuffixArray()
//	sa2 := suffixArray.CreateSuffixArray(str)
//	for i := 0; i < len(str); i++ {
//		if sa1.POS[i] != sa2.POS[i] {
//			t.Error("wrong")
//		}
//	}
//}
//
//func TestSuffixArray_BwtTransform(t *testing.T) {
//	str := "AGTAGTCAGTAC"
//	sa := suffixArray.CreateSuffixArray(str)
//	fmt.Println(sa.POS)
//	fmt.Println(sa.BwtTransform())
//}
//var alphabet = 4
//var str = util.GenRandomStr(int(math.Pow(2,28)),alphabet)
//var sa2 = suffixArray.CreateSuffixArray(str)
//var wtfmi = sa2.ToWTFMI()
//var rlfmi = sa2.ToRLFMI()
//var gv = gvcsa.MakeGVCSArray(*sa2)
//var sad = gvcsa.MakeSADCSArray(*sa2)
//var l = len(str)
//
//func BenchmarkWTFMI_Search(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		if i % 2 == 0{
//			str1 := util.GenRandomStr(20,alphabet)
//			wtfmi.Search(str1)
//		}else {
//			j := i % (i - 20)
//			wtfmi.Search(str[j : j+20])
//		}
//	}
//}
//func BenchmarkRLFMI_Search(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		if i % 2 == 0{
//			str1 := util.GenRandomStr(20,alphabet)
//			rlfmi.Search(str1)
//		}else {
//			j := i % (i - 20)
//			rlfmi.Search(str[j : j+20])
//		}
//	}
//}
//func BenchmarkSuffixArray_Search(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		if i % 2 == 0{
//			str1 := util.GenRandomStr(20,alphabet)
//			sa2.Search(str1)
//		}else {
//			j := i % (i - 20)
//			sa2.Search(str[j : j+20])
//		}
//	}
//}
//func BenchmarkWTFMI_Search2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		if i % 2 == 0{
//			str1 := util.GenRandomStr(20,alphabet)
//			gv.Search(str1)
//		}else {
//			j := i % (i - 20)
//			gv.Search(str[j : j+20])
//		}
//	}
//}
//func BenchmarkRLFMI_Search2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		if i % 2 == 0{
//			str1 := util.GenRandomStr(20,alphabet)
//			sad.Search(str1)
//		}else {
//			j := i % (i - 20)
//			sad.Search(str[j : j+20])
//		}
//	}
//}
