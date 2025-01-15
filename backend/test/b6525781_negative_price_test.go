package unit

import (
	"testing"
	"gitgub.com/sut-final-lab/entity"
	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func TestPriceNegative(t *testing.T) {
	g := NewGomegaWithT(t)

	var price float64
	t.Run(`Price is required`, func(t *testing.T) {
		product := entity.Products{
			Name: "Test",
			Price: price, //ผิดตรงนี้ไม่กรอก
			SKU: "AEW12348",
		}

		ok, err := govalidator.ValidateStruct(product)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Price is required"))
	})

	t.Run(`Price is not in range`, func(t *testing.T) {
		product := entity.Products{
			Name: "Test",
			Price: -45, //ผิดตรงนี้
			SKU: "AEW12348",
		}

		ok, err := govalidator.ValidateStruct(product)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Price must be between 1 and 1000"))
	})
}