package main

import "fmt"

// https://leetcode.cn/problems/qJnOS7/

func longestCommonSubsequenceMin(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	if text1[0] == text2[0] {
		return 1 + longestCommonSubsequenceMin(text1[1:], text2[1:])
	}

	max1 := longestCommonSubsequenceMin(text1, text2[1:])
	max2 := longestCommonSubsequenceMin(text1[1:], text2)
	return max(max1, max2)
}

// 这种是递归的想法，写起来很简单，但是，太慢了
func longestCommonSubsequenceDigui(text1 string, text2 string) int {
	text1, text2 = clear(text1, text2)
	fmt.Println(text1, text2)
	return longestCommonSubsequenceMin(text1, text2)
}

var maxMap map[int]int

// 这个其实也是有问题的，虽然过了，但是时间还是不快
func longestCommonSubsequenceV2(text1 string, text2 string) int {
	text1, text2 = clear(text1, text2)

	maxMap = make(map[int]int)

	println(text1, text2)
	return longest(text1, text2)
}

func longest(text1, text2 string) int {
	key := len(text1)*10000 + len(text2)
	if maxMap[key] > 0 {
		return maxMap[key]
	}

	// 每个字母就是两种，一个是要，一个是不要
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	if text1[0] == text2[0] {
		return 1 + longest(text1[1:], text2[1:])
	}
	var i int
	for ; i < len(text2); i++ {
		if text1[0] == text2[i] {
			// println(text1[0], text2[i], i)
			break
		}
	}
	var max1, max2, max3 int
	if i < len(text2) {
		max1 = longest(text1[1:], text2[i+1:]) + 1
	}
	for i = 0; i < len(text1); i++ {
		if text2[0] == text1[i] {
			break
		}
	}
	if i < len(text1) {
		max2 = longest(text1[i+1:], text2[1:]) + 1
	}
	max3 = longest(text1[1:], text2[1:])
	// println(text1, text2, max(max1, max2))

	maxLen := max(max(max1, max2), max3)
	maxMap[key] = maxLen
	return maxLen
}

func clear(text1, text2 string) (string, string) {
	// 处理一下两个字符串
	charMap1, charMap2 := make(map[int32]bool), make(map[int32]bool)
	for _, c := range text1 {
		charMap1[c] = true
	}
	for _, c := range text2 {
		charMap2[c] = true
	}
	charList1, charList2 := make([]int32, 0), make([]int32, 0)
	for _, c := range text1 {
		if charMap2[c] {
			charList1 = append(charList1, c)
		}
	}
	for _, c := range text2 {
		if charMap1[c] {
			charList2 = append(charList2, c)
		}
	}
	return string(charList1), string(charList2)
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}

func longestCommonSubsequence(text1 string, text2 string) int {
	text1, text2 = clear(text1, text2)

	// 经典最长子序列算法
	resMap := make([][]int, len(text1)+1)
	resMap[0] = make([]int, len(text2)+1)
	for i := 0; i < len(text1); i++ {
		resMap[i+1] = make([]int, len(text2)+1)
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				resMap[i+1][j+1] = resMap[i][j] + 1
			} else {
				resMap[i+1][j+1] = max(resMap[i+1][j], resMap[i][j+1])
			}
		}
	}
	return resMap[len(text1)][len(text2)]
}
func main() {
	fmt.Println(longestCommonSubsequence("pmjghexybyrgzczy", "hafcdqbgncrcbihkd"))
	fmt.Println(longestCommonSubsequence("fmtclsfaxchgjavqrifqbkzspazw", "nczivetoxqjclwbwtibqvelwxsdaz"))
	fmt.Println(longestCommonSubsequence("aebc", "abdc"))
	fmt.Println(longestCommonSubsequence("bsbininm", "jmjkbkjkv"))
	str1 := "fcvafurqjylclorwfoladwfqzkbebslwnmpmlkbezkxoncvwhstwzwpqxqtyxozkpgtgtsjobujezgrkvevklmludgtyrmjaxyputqbyxqvupojutsjwlwluzsbmvyxifqtglwvcnkfsfglwjwrmtyxmdgjifyjwrsnenuvsdedsbqdovwzsdghclcdexmtsbexwrszihcpibwpidixmpmxshwzmjgtadmtkxqfkrsdqjcrmxkbkfoncrcvoxuvcdytajgfwrcxivixanuzerebuzklyhezevonqdsrkzetsrgfgxibqpmfuxcrinetyzkvudghgrytsvwzkjulmhanankxqfihenuhmfsfkfepibkjmzybmlkzozmluvybyzsleludsxkpinizoraxonmhwtkfkhudizepyzijafqlepcbihofepmjqtgrsxorunshgpazovuhktatmlcfklafivivefyfubunszyvarcrkpsnglkduzaxqrerkvcnmrurkhkpargvcxefovwtapedaluhclmzynebczodwropwdenqxmrutuhehadyfspcpuxyzodifqdqzgbwhodcjonypyjwbwxepcpujerkrelunstebopkncdazexsbezmhynizsvarafwfmnclerafejgnizcbsrcvcnwrolofyzulcxaxqjqzunedidulspslebifinqrchyvapkzmzwbwjgbyrqhqpolwjijmzyduzerqnadapudmrazmzadstozytonuzarizszubkzkhenaxivytmjqjgvgzwpgxefatetoncjgjsdilmvgtgpgbibexwnexstipkjylalqnupexytkradwxmlmhsnmzuxcdkfkxyfgrmfqtajatgjctenqhkvyrgvapctqtyrufcdobibizihuhsrsterozotytubefutaxcjarknynetipehoduxyjstufwvkvwvwnuletybmrczgtmxctuny"
	str2 := "nohgdazargvalupetizezqpklktojqtqdivcpsfgjopaxwbkvujilqbclehulatshehmjqhyfkpcfwxovajkvankjkvevgdovazmbgtqfwvejczsnmbchkdibstklkxarwjqbqxwvixavkhylqvghqpifijohudenozotejoxavkfkzcdqnoxydynavwdylwhatslyrwlejwdwrmpevmtwpahatwlaxmjmdgrebmfyngdcbmbgjcvqpcbadujkxaxujudmbejcrevuvcdobolcbstifedcvmngnqhudixgzktcdqngxmruhcxqxypwhahobudelivgvynefkjqdyvalmvudcdivmhghqrelurodwdsvuzmjixgdexonwjczghalsjopixsrwjixuzmjgxydqnipelgrivkzkxgjchibgnqbknstspujwdydszohqjsfuzstyjgnwhsrebmlwzkzijgnmnczmrehspihspyfedabotwvwxwpspypctizyhcxypqzctwlspszonsrmnyvmhsvqtkbyhmhwjmvazaviruzqxmbczaxmtqjexmdudypovkjklynktahupanujylylgrajozobsbwpwtohkfsxeverqxylwdwtojoxydepybavwhgdehafurqtcxqhuhkdwxkdojipolctcvcrsvczcxedglgrejerqdgrsvsxgjodajatsnixutihwpivihadqdotsvyrkxehodybapwlsjexixgponcxifijchejoxgxebmbclczqvkfuzgxsbshqvgfcraxytaxeviryhexmvqjybizivyjanwxmpojgxgbyhcruvqpafwjslkbohqlknkdqjixsfsdurgbsvclmrcrcnulinqvcdqhcvwdaxgvafwravunurqvizqtozuxinytafopmhchmxsxgfanetmdcjalmrolejidylkjktunqhkxchyjmpkvsfgnybsjedmzkrkhwryzan"

	fmt.Println(longestCommonSubsequence(str1, str2))
}
