package visitor

func ExampleRequestVisitor() {
	cc := &CustomerCol{}
	cc.Add(NewEnterpriseCustomer("A company"))
	cc.Add(NewEnterpriseCustomer("B company"))
	cc.Add(NewIndividualCustomer("bob"))
	cc.Accept(&ServiceRequestVisitor{})
	// Output:
	// serving enterprise customer A company
	// serving enterprise customer B company
	// serving individual customer bob
}

func ExampleAnalysis() {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Accept(&AnalysisVisitor{})
	// Output:
	// analysis enterprise customer A company
	// analysis enterprise customer B company
}
