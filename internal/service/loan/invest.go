package loan

import (
	"bytes"
	"errors"
	"fmt"
	"loan-service/internal/constant"
	"loan-service/internal/entity"
	"loan-service/internal/service/mailer"
	"loan-service/pkg/model/request"

	"github.com/signintech/gopdf"
	log "github.com/sirupsen/logrus"

	"github.com/google/uuid"
)

func (ls *Loan) Invest(req request.InvestLoan, requestedBy string) (success bool, err error) {
	loan := ls.loanRepo.Get(req.LoanID)

	if loan == nil {
		return false, errors.New("loan not exist")
	}

	if loan.Status != constant.Approved {
		return false, errors.New("loan status must be approved")
	}

	newAmount := loan.InvestedAmount + req.Amount

	if newAmount > loan.PrincipalAmount {
		remainingNeeded := loan.PrincipalAmount - loan.InvestedAmount
		return false, fmt.Errorf("remaining amount needed: %d", remainingNeeded)
	}

	existLoanInvestment := ls.loanInvestmentRepo.GetByLoanAndInvestor(loan.ID, requestedBy)

	if existLoanInvestment != nil {
		existLoanInvestment.Amount += req.Amount
		if err = ls.loanInvestmentRepo.Save(existLoanInvestment); err != nil {
			log.Errorf("Error saving loan investment: %s", err.Error())
			return
		}
	} else {
		newInvestment := &entity.LoanInvestment{
			ID:         uuid.New().String(),
			LoanID:     loan.ID,
			InvestorID: requestedBy,
			Amount:     req.Amount,
			Audit: entity.Audit{
				CreatedBy: requestedBy,
				UpdatedBy: requestedBy,
			},
		}
		if err = ls.loanInvestmentRepo.Create(newInvestment); err != nil {
			log.Errorf("Error creating loan investment: %s", err.Error())
			return
		}
	}

	loan.InvestedAmount += req.Amount
	if loan.InvestedAmount == loan.PrincipalAmount {
		loan.Status = constant.Invested
	}

	if err = ls.loanRepo.Save(loan); err != nil {
		log.Errorf("Error saving loan: %s", err.Error())
		return
	}

	go ls.sendInvestorAgreementLetterEmail(loan)

	return true, nil
}

func (ls *Loan) sendInvestorAgreementLetterEmail(loan *entity.Loan) {
	if loan.Status != constant.Invested {
		return
	}

	loanInvestments := ls.loanInvestmentRepo.GetByLoanID(loan.ID)

	for _, inv := range loanInvestments {
		go func(l *entity.Loan, li *entity.LoanInvestment) {
			pdfBuf, err := ls.createInvestorAgreementLetter(l, li)

			if err != nil {
				log.Errorf("Error creating investor agreement letter: %s", err.Error())
				return
			}

			req := &mailer.Request{
				To:      []string{inv.InvestorID},
				Subject: "Investor Agreement Letter",
				Text:    "Investor Agreement Letter",
				Attachment: &mailer.Attachment{
					Name:     "agreement-letter.pdf",
					MimeType: "application/pdf",
					Content:  pdfBuf,
				},
			}
			if err := ls.mailerService.Send(req); err != nil {
				log.Errorf("Error sending email to %s: %s", inv.InvestorID, err.Error())
			}
		}(loan, &inv)
	}
}

func (ls *Loan) createInvestorAgreementLetter(loan *entity.Loan, investment *entity.LoanInvestment) (*bytes.Buffer, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	if err := pdf.AddTTFFont("roboto", "assets/fonts/roboto.ttf"); err != nil {
		return nil, err
	}
	if err := pdf.SetFont("roboto", "", 14); err != nil {
		return nil, err
	}

	pdf.Cell(nil, "LoanID: "+loan.ID)
	pdf.Br(20)
	pdf.Cell(nil, fmt.Sprintf("Principal amount: %d", loan.PrincipalAmount))
	pdf.Br(20)
	pdf.Cell(nil, fmt.Sprintf("Invested amount: %d", loan.InvestedAmount))
	pdf.Br(20)
	pdf.Cell(nil, fmt.Sprintf("Your invested amount: %d", investment.Amount))
	pdf.Br(20)

	var buf bytes.Buffer
	if _, err := pdf.WriteTo(&buf); err != nil {
		return nil, err
	}

	return &buf, nil
}
