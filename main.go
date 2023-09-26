package main

import (
	"fmt"

	"github.com/vkaushik/typist/internal/check"
	"github.com/vkaushik/typist/pkg/error"
)

func main() {
	fmt.Println("typist test")
	// fe, he := evaluate.GetErrors(masterText, testText)
	fe, he := check.GetErrors(Master, Test)
	fmt.Println("=========================")
	error.PrintFullMistakes(fe)
	fmt.Println("=========================")
	error.PrintHalfMistakes(he)
}

var masterText = "so Hello world how are world"
var testText = "so h e.l l o world how are world"

func PrintErrors(fullErrors []error.TypingError, halfErrors []error.TypingError) {
	fmt.Println("Full Errors:")
	for _, err := range fullErrors {
		fmt.Println(err.Error())
	}

	fmt.Println("Half Errors:")
	for _, err := range halfErrors {
		fmt.Println(err.Error())
	}
}

var Master string = `Budget after budget and more particularly this budget has been having a hit at the States.  Before I go into what has been stated about the Budget proposals and how they are affecting the States.  I would like to take from the document itself.  In the year 1997-98, among the taxes which the Government of India collected Rs.192 crores out of Rs.514 crores went to the States.  In the current budget, just Rs.28 crores out of Rs.615 crores will go to the States.  When it comes to the income tax, Rs.10 crores are charged net and the Union Government will have no share.  On the other hand, according to the Budget at a glance, there will be minus Rs.94 crores of the share of the States which means they will have to shell down.  This is one instance as to how there has been continuous shrinking of the share of the States in terms of the Union finances.  Then, Sir, in the total expenditure out of Rs.73,000 crores, transfer to the State will account for about Rs.26,000 crores for 1998-99 and amounts to 33%.  In the previous year, out of Rs.66,100 crores, Rs.24,000 crores was the share of the States and it works out to 37%.  Even from that point of view, there has been a shrinking of the share.  When it comes to the question of the gross revenue and the States share, it has come down to 40%, a fact which has been made clear in the recent National Development Council meeting also.  I have given three instances but I do not want to go into the details which I have done last year.  I could have done that also but these documents themselves show how the resource crunch has affected the Budget proposals of the Finance Minister.  I will read three statements of a particular Chief Minister.  He says that in the Central Budget proposals for the year 1999-2000, an additional revenue of Rs.100 crores from Income Tax could go only to the Centre.  This is one statement mad by one Chief Minister of a State of India.  Then, the Chief Minister also referred to a new levy on the transfer of wealth through inheritance which he said was nothing but Estate Duty in another form.
Budget after budget and more particularly this budget has been having a hit at the States.  Before I go into what has been stated about the Budget proposals and how they are affecting the States.  I would like to take from the document itself.  In the year 1997-98, among the taxes which the Government of India collected Rs.192 crores out of Rs.514 crores went to the States.  In the current budget, just Rs.28 crores out of Rs.615 crores will go to the States.  When it comes to the income tax, Rs.10 crores are charged net and the Union Government will have no share.  On the other hand, according to the Budget at a glance, there will be minus Rs.94 crores of the share of the States which means they will have to shell down.  This is one instance as to how there has been continuous shrinking of the share of the States in terms of the Union finances.  Then, Sir, in the total expenditure out of Rs.73,000 crores, transfer to the State will account for about Rs.26,000 crores for 1998-99 and amounts to 33%.  In the previous year, out of Rs.66,100 crores, Rs.24,000 crores was the share of the States and it works out to 37%.  Even from that point of view, there has been a shrinking of the share.  When it comes to the question of the gross revenue and the States share, it has come down to 40%, a fact which has been made clear in the recent National Development Council meeting also.  I have given three instances but I do not want to go into the details which I have done last year.  I could have done that also but these documents themselves show how the resource crunch has affected the Budget proposals of the Finance Minister.  I will read three statements of a particular Chief Minister.  He says that in the Central Budget proposals for the year 1999-2000, an additional revenue of Rs.100 crores from Income Tax could go only to the Centre.  This is one statement mad by one Chief Minister of a State of India.  Then, the Chief Minister also referred to a new levy on the transfer of wealth through inheritance which he said was nothing but Estate Duty in another form.`

var Test string = `Budget after budget and more particularlly this budget has been having a hit at the States. Before I go into what has been stated about the budget proposals ans how they are affecting the States. I would like to take from the document itself. In the year 1997-98, among the taxes which the Government of India collected Rs. 192 crores out of Rs. 514 crores went to the States. In the current budget, just Rs. 28 crores out of Rs. 615 crores will go the States. When it comes to the income tax, Rs. 10 crores are charged net and the Union Government will have no share. On the other hand, according to the budget at a glance, there will be munus Rs. 94 crores of the share of the States which means they will have to shell down. This is one instance as to how there has been continuous shrinking of the share of the States in terms of the Union finances.Then, Sir, in the total expenditure out of Rs, 73,000 crores, transfer to the State will account for about Rs. 26,000 crores for 1998-99 and amounts to 33%. In the previous year, out of Rs, 66,100 crores, Rs.24000 crores was the share of the States and it works outof 37%. Even from that point of view, therre has been a shrinking of the share. When it comes to the question of the gross revenue and the states share, it has come down to 40%,. a fact which has been made clear in the recent National DevelopmentCouncil meeting also. I have given three instances but I don not want to go into the details which I have done last year. I could have done that also but these documents themselves show how the resourcse crunch has affected The Budget proposals of the Finance Minister. I will ead three statements of a particular Chief Minister. He says that in the Central Budget proposals for the year 1999-2000, an additional revenue fo Rs. 100 crores from Income Tax could go only to the Centre. This is one statement mad by one Chief Minister of a State of india. Then, the chief Minister also referred toa new levy on the transfer of wealth through inhyeritance which he said was nothing but Estate Duty in another  form.
Budget after `
