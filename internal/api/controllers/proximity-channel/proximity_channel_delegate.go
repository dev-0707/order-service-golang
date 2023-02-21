package proximity_channel

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ProximityChannelControllerDelegate struct {
}

func (p ProximityChannelControllerDelegate) CheckCoupon(c *gin.Context, params CheckCouponParams) {
	log.Info().Msg("Esecuzione check coupon")
}

func NewProximityChannelControllerDelegate() *ProximityChannelControllerDelegate {
	return &ProximityChannelControllerDelegate{}
}

func (p ProximityChannelControllerDelegate) CheckGaranzia(c *gin.Context, params CheckGaranziaParams) {
}
func (p ProximityChannelControllerDelegate) CheckGaranziaDeferred(c *gin.Context, params CheckGaranziaDeferredParams) {
}
func (p ProximityChannelControllerDelegate) CheckIban(c *gin.Context, params CheckIbanParams)     {}
func (p ProximityChannelControllerDelegate) CheckModulo(c *gin.Context, params CheckModuloParams) {}
func (p ProximityChannelControllerDelegate) CheckTarga(c *gin.Context, params CheckTargaParams)   {}
func (p ProximityChannelControllerDelegate) CreaCliente(c *gin.Context, params CreaClienteParams) {}
func (p ProximityChannelControllerDelegate) AggiornaCliente(c *gin.Context, codiceCliente int64, params AggiornaClienteParams) {
}
func (p ProximityChannelControllerDelegate) GetByCodiceFiscale(c *gin.Context, codiceFiscale string, params GetByCodiceFiscaleParams) {
	log.Info().Msg("Esecuzione GetByCodiceFiscale")
}
func (p ProximityChannelControllerDelegate) AttivazioneContratto(c *gin.Context, params AttivazioneContrattoParams) {
}
func (p ProximityChannelControllerDelegate) GetFirmaContratto(c *gin.Context, params GetFirmaContrattoParams) {
}
func (p ProximityChannelControllerDelegate) GetDettaglioContratto(c *gin.Context, id int64, params GetDettaglioContrattoParams) {
}
func (p ProximityChannelControllerDelegate) AggiungiTelepass(c *gin.Context, id int64, params AggiungiTelepassParams) {
}
func (p ProximityChannelControllerDelegate) DisattivaTelepass(c *gin.Context, id int64, params DisattivaTelepassParams) {
}
func (p ProximityChannelControllerDelegate) SostituisciTelepass(c *gin.Context, id int64, params SostituisciTelepassParams) {
}
func (p ProximityChannelControllerDelegate) AttivaOpzioni(c *gin.Context, id int64, params AttivaOpzioniParams) {
}
func (p ProximityChannelControllerDelegate) DisattivaOpzione(c *gin.Context, id int64, params DisattivaOpzioneParams) {
}
