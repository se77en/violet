package index

import (
	"testing"

	"github.com/cosmtrek/violet/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewInvert(t *testing.T) {
	path, err := utils.TempDir("", true)
	assert.Nil(t, err)
	invert, err := NewInvert(path, "field0", TString, segmenter())
	assert.Nil(t, err)
	assert.NotNil(t, invert)
}

func TestInvert_addDocument_saveTmpInvert_mergeTmpInvert_searchTerm(t *testing.T) {
	path, err := utils.TempDir("", true)
	assert.Nil(t, err)
	invert, err := NewInvert(path, "field1", TString, segmenter())
	assert.Nil(t, err)
	assert.NotNil(t, invert)
	invert.addDocument(uint64(0), "浓烟下的诗歌电台 陈鸿宇 理想三旬")
	invert.addDocument(uint64(1), "雨后有车驶来")
	invert.addDocument(uint64(2), "驶过暮色苍白")
	invert.addDocument(uint64(3), "旧铁皮往南开 恋人已不在")
	invert.addDocument(uint64(4), "收听浓烟下的 诗歌电台")
	invert.addDocument(uint64(5), "不动情的咳嗽 至少看起来")
	invert.addDocument(uint64(6), "归途也还可爱")
	invert.addDocument(uint64(7), "琴弦少了姿态")
	invert.addDocument(uint64(8), "再不见那夜里 听歌的小孩")
	invert.addDocument(uint64(9), "时光匆匆独白")
	invert.addDocument(uint64(10), "将颠沛磨成卡带")
	invert.addDocument(uint64(11), "已枯卷的情怀 踏碎成年代")
	invert.addDocument(uint64(12), "就老去吧 孤独别醒来")
	invert.saveTmpInvert()
	invert.addDocument(uint64(13), "你渴望的离开")
	invert.addDocument(uint64(14), "只是无处停摆")
	invert.addDocument(uint64(15), "就歌唱吧 眼睛眯起来")
	invert.addDocument(uint64(16), "而热泪的崩坏")
	invert.addDocument(uint64(17), "只是没抵达的存在")
	invert.addDocument(uint64(18), "青春又醉倒在")
	invert.addDocument(uint64(19), "籍籍无名的怀")
	invert.addDocument(uint64(20), "靠嬉笑来虚度 聚散得慷慨")
	invert.addDocument(uint64(21), "辗转却去不到 对的站台")
	invert.addDocument(uint64(22), "如果漂泊是成长 必经的路牌")
	invert.addDocument(uint64(23), "你迷醒岁月中")
	invert.addDocument(uint64(24), "那贫瘠的未来")
	invert.addDocument(uint64(25), "像遗憾季节里 未结果的爱")
	invert.addDocument(uint64(26), "弄脏了每一页诗")
	invert.addDocument(uint64(27), "吻最疼痛的告白")
	invert.addDocument(uint64(28), "而风声吹到这 已不需要释怀")
	invert.addDocument(uint64(29), "就老去吧 孤独别醒来")
	invert.saveTmpInvert()
	invert.addDocument(uint64(30), "你渴望的离开")
	invert.addDocument(uint64(31), "只是无处停摆")
	invert.addDocument(uint64(32), "就歌唱吧 眼睛眯起来")
	invert.addDocument(uint64(33), "而热泪的崩坏")
	invert.addDocument(uint64(34), "只是没抵达的存在")
	invert.addDocument(uint64(35), "就甜蜜地忍耐")
	invert.addDocument(uint64(36), "繁星润湿窗台")
	invert.addDocument(uint64(37), "光影跳动着像在 困倦里说爱")
	invert.addDocument(uint64(38), "再无谓的感慨")
	invert.addDocument(uint64(39), "以为明白")
	invert.addDocument(uint64(40), "梦倒塌的地方 今已爬满青苔")
	err = invert.saveTmpInvert()
	assert.Nil(t, err)
	err = invert.mergeTmpInvert()
	assert.Nil(t, err)
	docs1, found1 := invert.searchTerm("孤独")
	assert.True(t, found1)
	expected1 := []Doc{{DocID: 12}, {DocID: 29}}
	assert.EqualValues(t, expected1, docs1)
	docs2, found2 := invert.searchTerm("歌唱")
	assert.True(t, found2)
	expected2 := []Doc{{DocID: 15}, {DocID: 32}}
	assert.EqualValues(t, expected2, docs2)
}
