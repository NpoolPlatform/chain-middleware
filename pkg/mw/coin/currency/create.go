package currency

type createHandler struct {
	*Handler
}

func (h *Handler) CreateCurrency(ctx context.Context) (*npool.Currency, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	return h.GetCurrency(ctx)
}
