package shared

import (
	kafkaClient "github.com/mehdihadeli/store-golang-microservice-sample/pkg/kafka"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/logger"
	"github.com/mehdihadeli/store-golang-microservice-sample/pkg/mediatr"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/config"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/internal/products/contracts/repositories"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/internal/products/features/creating_product"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/internal/products/features/deleting_product"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/internal/products/features/getting_product_by_id"
	"github.com/mehdihadeli/store-golang-microservice-sample/services/catalogs/internal/products/features/updating_product"
	"github.com/pkg/errors"
)

func NewMediator(log logger.Logger, cfg *config.Config, pgRepo repositories.ProductRepository, kafkaProducer kafkaClient.Producer) (*mediatr.Mediator, error) {

	md := mediatr.New()

	err := md.Register(
		creating_product.NewCreateProductHandler(log, cfg, pgRepo, kafkaProducer),
		updating_product.NewUpdateProductHandler(log, cfg, pgRepo, kafkaProducer),
		deleting_product.NewDeleteProductHandler(log, cfg, pgRepo, kafkaProducer),
		getting_product_by_id.NewGetProductByIdHandler(log, cfg, pgRepo),
	)

	if err != nil {
		return nil, errors.Wrap(err, "error while registering handlers in the mediator")
	}

	return &md, nil
}
