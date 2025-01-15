package unit

import (
	"testing"
	"gitgub.com/sut-final-lab/entity"
	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func TestProductPositive(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Product is positive`, func(t *testing.T) {
		product := entity.Products{
			Name: "Test",
			Price: 120,
			SKU: "AEW12348",
		}

		ok, err := govalidator.ValidateStruct(product)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}