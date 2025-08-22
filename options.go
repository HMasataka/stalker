package stalker

type Options struct {
	skipFrames int
}

type Option func(*Options)

func SkipFrame(n int) Option {
	return func(g *Options) {
		g.skipFrames = n
	}
}
