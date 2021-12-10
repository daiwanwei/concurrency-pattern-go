package main

type PayBack func() (money int, err error)

type SuccessFn func(money int)
type FailFn func(err error)

type LoanCompany struct {
	success SuccessFn
	fail    FailFn
}

func (company *LoanCompany) Success(fn SuccessFn) *LoanCompany {
	company.success = fn
	return company
}

func (company *LoanCompany) Fail(fn FailFn) *LoanCompany {
	company.fail = fn
	return company
}

func (company *LoanCompany) Execute(fn PayBack) {
	execFn := func() {
		money, err := fn()
		if err != nil {
			company.fail(err)
		}
		company.success(money)
	}
	go execFn()
}
