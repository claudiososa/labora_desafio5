package models

type Loan struct {
	id          int
	user_id     int
	book_id     int
	loan_date   string
	return_date string
}

func (loan *Loan) GetId() int {
	return loan.id
}

func (loan *Loan) GetUser_id() int {
	return loan.user_id
}

func (loan *Loan) GetBook_id() int {
	return loan.book_id
}

func (loan *Loan) GetLoan_date() string {
	return loan.loan_date
}

func (loan *Loan) GetReturn_date() string {
	return loan.return_date
}

func (loan *Loan) SetId(id int) *Loan {
	loan.id = id
	return loan
}

func (loan *Loan) SetUser_id(user_id int) *Loan {
	loan.user_id = user_id
	return loan
}

func (loan *Loan) SetBook_id(book_id int) *Loan {
	loan.book_id = book_id
	return loan
}

func (loan *Loan) SetLoan_date(loan_date string) *Loan {
	loan.loan_date = loan_date
	return loan
}

func (loan *Loan) SetReturn_date(return_date string) *Loan {
	loan.return_date = return_date
	return loan
}
