package levenshtein

import (
	"testing"
)

func TestLevenScore(t *testing.T) {
	samples := []struct {
		a      string
		b      string
		expect float32
	}{
		{"", "", 1.0},
		{"", "a", 0.0},
		{"a", "", 0.0},
		{"a", "a", 1.0},
		{"a", "b", 0.0},
		{"kitten", "sitting", 0.6153846153846154},
		{"example", "samples", 0.7142857142857143},
		{"creasty", "Creasty", 0.8571428571428571},
		{"creasty", "Creastyi", 0.8},
		{"levenshtein", "frankenstein", 0.6086956521739131},
		{"distance", "difference", 0.5555555555555556},
		{"哇哈哈哈", "娃哈哈", 0.5714285714285714},
		{"哇哈哈哈", "娃哈哈a", 0.5},
		{
			"哺乳动物的红细胞在发育成熟过程中细胞核逐渐退化,为携带氧气的血红蛋白腾出空间.哺乳动物成熟的红细胞一般不含细胞器且取材方便,是研究细胞膜的好材料.根据材料回答问题: (1)在正常机体的红细胞内不会存在的物质是（ ）. A: \\(\\ce{O_{2}}\\) B: \\(\\ce{CO_{2}}\\) C: 抗体 D: 血红蛋白 (2)简述红细胞膜的制备过程___. (3)红细胞细胞膜上具有识别功能的物质成分是___;人的红细胞与草履虫细胞相比,细胞膜的差别主要体现在哪方面?___. (4)用人的口腔上皮细胞能否制备得到较纯净的细胞膜?___.请分析原因:___.",
			"某研究性学习小组欲研究细胞膜的结构和成分,如果你是课题组成员,请设计一个简易实验得到较纯净的细胞膜,并对其成分进行鉴定. (1)应选取人的哪种细胞作实验材料（ ） A: 成熟红细胞 B: 神经细胞 C: 白细胞 D: 口腔上皮细胞 (2)将选取的上述材料放入___中,由于渗透作用,一段时间后细胞将破裂. (3)经过(2)的实验步骤后,再用___法获得较纯净的细胞膜.",
			0.27213822894168466,
		},
	}

	for _, sample := range samples {
		a := []byte(sample.a)
		b := []byte(sample.b)
		actual := LevenScore(a, b)
		if actual != sample.expect {
			t.Errorf("Expected the score of `%s` and `%s` to be %f, but was %f", sample.a, sample.b, sample.expect, actual)
		}
	}
}

func TestLevenScoreInt8(t *testing.T) {
	samples := []struct {
		a      string
		b      string
		expect float32
	}{
		{"", "", 1.0},
		{"", "a", 0.0},
		{"a", "", 0.0},
		{"a", "a", 1.0},
		{"a", "b", 0.0},
		{"kitten", "sitting", 0.6153846153846154},
		{"example", "samples", 0.7142857142857143},
		{"creasty", "Creasty", 0.8571428571428571},
		{"creasty", "Creastyi", 0.8},
		{"levenshtein", "frankenstein", 0.6086956521739131},
		{"distance", "difference", 0.5555555555555556},
		{"哇哈哈哈", "娃哈哈", 0.5714285714285714},
		{"哇哈哈哈", "娃哈哈a", 0.5},
		{
			"哺乳动物的红细胞在发育成熟过程中细胞核逐渐退化,为携带氧气的血红蛋白腾出空间.哺乳动物成熟的红细胞一般不含细胞器且取材方便,是研究细胞膜的好材料.根据材料回答问题: (1)在正常机体的红细胞内不会存在的物质是（ ）. A: \\(\\ce{O_{2}}\\) B: \\(\\ce{CO_{2}}\\) C: 抗体 D: 血红蛋白 (2)简述红细胞膜的制备过程___. (3)红细胞细胞膜上具有识别功能的物质成分是___;人的红细胞与草履虫细胞相比,细胞膜的差别主要体现在哪方面?___. (4)用人的口腔上皮细胞能否制备得到较纯净的细胞膜?___.请分析原因:___.",
			"某研究性学习小组欲研究细胞膜的结构和成分,如果你是课题组成员,请设计一个简易实验得到较纯净的细胞膜,并对其成分进行鉴定. (1)应选取人的哪种细胞作实验材料（ ） A: 成熟红细胞 B: 神经细胞 C: 白细胞 D: 口腔上皮细胞 (2)将选取的上述材料放入___中,由于渗透作用,一段时间后细胞将破裂. (3)经过(2)的实验步骤后,再用___法获得较纯净的细胞膜.",
			0.27213822894168466,
		},
	}

	for _, sample := range samples {
		a := []byte(sample.a)
		a1 := make([]int8, len(a))
		for i, v := range a {
			a1[i] = int8(v)
		}
		b := []byte(sample.b)
		b1 := make([]int8, len(b))
		for i, v := range b {
			b1[i] = int8(v)
		}
		actual := LevenScoreInt8(a1, b1)
		if actual != sample.expect {
			t.Errorf("Expected the score of `%s` and `%s` to be %f, but was %f", sample.a, sample.b, sample.expect, actual)
		}
	}
}
