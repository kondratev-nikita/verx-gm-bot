package msg_gen

import (
	"math/rand/v2"
	"strings"
)

func (s *service) GenGoodMorning() string {
	randIndex := rand.N(len(s.cfg.MessagesCfg.GoodMorning))
	randTemplate := s.cfg.MessagesCfg.GoodMorning[randIndex]

	return s.fmtTextWithEmojis(randTemplate)
}

func (s *service) GenGoodNight() string {
	randIndex := rand.N(len(s.cfg.MessagesCfg.GoodNight))
	randTemplate := s.cfg.MessagesCfg.GoodNight[randIndex]

	return s.fmtTextWithEmojis(randTemplate)
}

// setEmojis принимается текст только с {emoji}, меняет на случайные эмодзи
func (s *service) fmtTextWithEmojis(text string) string {
	const target = "{emoji}"
	emojisCount := strings.Count(text, target)

	for i := 0; i < emojisCount; i++ {
		text = strings.Replace(text, target, s.randEmoji(), 1)
	}

	return text
}

func (s *service) randEmoji() string {
	return s.emojisSet[rand.N(len(s.emojisSet))]
}
