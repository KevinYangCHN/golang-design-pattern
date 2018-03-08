package command

func ExampleCommand() {
	c1 := NewStartCommand(&MotherBoard{})
	c2 := NewRebootCommand(&MotherBoard{})

	box := NewBox(c1, c2)

	box.PressButton1()
	box.PressButton2()

	// Output:
	// system starting
	// system rebooting
}
