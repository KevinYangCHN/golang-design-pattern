package builder

type Builder interface {
	Part1()
	Part2()
}

type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) GetResult() int {
	return b.result
}
