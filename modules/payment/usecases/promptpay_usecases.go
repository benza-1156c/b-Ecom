package usecases

import (
	"fmt"
	"strings"
)

type PaymentUsacase interface {
	GeneratePromptPayPayload(phone, amount string) string
}

type paymentUsecase struct{}

func NewPaymentUsacase() PaymentUsacase {
	return &paymentUsecase{}
}

func (p *paymentUsecase) GeneratePromptPayPayload(phone, amount string) string {
	amountFormatted := FormatAmount(amount)
	formattedPhone := FormatPhone(phone)

	aid := "A000000677010111"
	merchant := makeTLV("00", aid) + makeTLV("01", formattedPhone)
	tag29 := makeTLV("29", merchant)

	payload := "" +
		makeTLV("00", "01") +
		makeTLV("01", "11") +
		tag29 +
		makeTLV("52", "0000") +
		makeTLV("53", "764") +
		makeTLV("54", amountFormatted) +
		makeTLV("58", "TH") +
		"6304"

	crc := CalculateCRC16(payload)
	return payload + crc
}

func makeTLV(id, value string) string {
	return fmt.Sprintf("%s%02d%s", id, len(value), value)
}

func FormatPhone(phone string) string {
	if strings.HasPrefix(phone, "0") && len(phone) == 10 {
		return "0066" + phone[1:]
	} else if strings.HasPrefix(phone, "66") && len(phone) == 11 {
		return "00" + phone
	}
	return phone
}

func FormatAmount(amount string) string {
	if !strings.Contains(amount, ".") {
		return amount + ".00"
	}
	return amount
}

func CalculateCRC16(input string) string {
	polynomial := 0x1021
	result := 0xFFFF

	for _, c := range input {
		result ^= int(c) << 8
		for i := 0; i < 8; i++ {
			if (result & 0x8000) != 0 {
				result = (result << 1) ^ polynomial
			} else {
				result <<= 1
			}
		}
	}
	result &= 0xFFFF
	return fmt.Sprintf("%04X", result)
}
