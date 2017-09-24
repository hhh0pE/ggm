package ggm

type PgNotifyEvent int8

const (
	PG_NOTIFY_INSERT PgNotifyEvent = iota + 1
	PG_NOTIFY_UPDATE
	PG_NOTIFY_DELETE
)

type pgNotifyParams struct {
}

func (pnp pgNotifyParams) Payload(fields ...interface{}) {

}

func (pnp pgNotifyParams) Name(name string) *pgNotifyParams {
	return &pgNotifyParams{}
}

func PgNotify(events ...PgNotifyEvent) *pgNotifyParams {
	return &pgNotifyParams{}
}
//type modelWithPgNotifyInsert interface {
//	onPgNotifyInsert()
//}
