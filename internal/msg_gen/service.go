package msg_gen

type Service interface {
	GenGoodMorning() string
	GenGoodNight() string
}

type service struct {
	cfg       Config
	emojisSet []string
}

type MessagesConfig struct {
	GoodNight   []string
	GoodMorning []string
}

type Config struct {
	MessagesCfg MessagesConfig
}

func New(cfg Config) (Service, error) {
	if cfg.MessagesCfg.GoodNight == nil ||
		len(cfg.MessagesCfg.GoodNight) == 0 {
		return nil, newErr("cfg.MessagesCfg.GoodNight require", nil)
	}

	if cfg.MessagesCfg.GoodMorning == nil ||
		len(cfg.MessagesCfg.GoodMorning) == 0 {
		return nil, newErr("cfg.MessagesCfg.GoodMorning require", nil)
	}

	return &service{
		cfg: cfg,
		emojisSet: []string{
			"🥺", "🥰", "🤗", "💞", "❤️",
			"🙈", "💘", "✨", "💖", "💕",
			"💗", "💓", "💝", "💟", "💖",
			"💕", "💗", "💓", "💘", "💝",
			"💞", "💟", "💖", "💕", "💗",
			"💓", "💘", "💝", "💞", "💟",
			"💖", "💕", "💗", "💓", "💘",
			"💝", "💞", "💟", "💖", "💕",
			"💗", "💓", "💘", "💝", "💞",
			"💟", "💖", "💕", "💗", "💓",
			"💘", "💝", "💞", "💟", "💖",
			"💕", "💗", "💓", "💘", "💝",
			"💞", "💟", "💖", "💕", "💗",
			"💓", "💘", "💝", "💞", "💟",
			"💖", "💕", "💗", "💓", "💘",
			"💝", "💞", "💟", "💖", "💕",
			"💗", "💓", "💘", "💝", "💞",
			"💟", "💖", "💕", "💗", "💓",
			"💘", "💝", "💞", "💟", "💖",
			"💕", "💗", "💓", "💘", "💝",
			"💞", "💟", "💖", "💕", "💗",
			"💓", "💘", "💝", "💞", "💟",
		},
	}, nil
}
