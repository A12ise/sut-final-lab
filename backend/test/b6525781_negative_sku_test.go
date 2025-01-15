package unit

import (
	"testing"
	"gitgub.com/sut-final-lab/entity"
	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func TestSKUNegative(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Price is required`, func(t *testing.T) {
		product := entity.Products{
			Name: "Test",
			Price: 800, 
			SKU: "", //ผิดตรงนี้ไม่กรอก
		}

		ok, err := govalidator.ValidateStruct(product)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("SKU is required"))
	})

	t.Run(`Price is not in range`, func(t *testing.T) {
		product := entity.Products{
			Name: "Test",
			Price: 800, 
			SKU: "awd12348",//ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(product)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("SKU must start with 3 uppercase English letters (A-Z) followed by 5 digits (0-9)"))
	})
}