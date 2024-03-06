package handler

type Service interface {}

type Handler struct {

}

func NewHandler(service *Service) *Handler {
	return &Handler{
		
		// logger: logger.GetLoggerInstance(),
	}
}