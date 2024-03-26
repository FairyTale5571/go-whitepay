package whitepay

import (
	"fmt"
	"net/url"
	"time"
)

type Filters func() url.Values

func Page(page int) Filters {
	return func() url.Values {
		return url.Values{"page": {fmt.Sprintf("%d", page)}}
	}
}

func PerPage(perPage int) Filters {
	return func() url.Values {
		return url.Values{"per_page": {fmt.Sprintf("%d", perPage)}}
	}
}

func DateFrom(dateFrom time.Time) Filters {
	return func() url.Values {
		return url.Values{"date_from": {dateFrom.Format("2006-01-02+15:04:05")}}
	}
}

func DateTo(dateTo time.Time) Filters {
	return func() url.Values {
		return url.Values{"date_to": {dateTo.Format("2006-01-02+15:04:05")}}
	}
}

func CompletedAtFrom(completedAtFrom time.Time) Filters {
	return func() url.Values {
		return url.Values{"completed_at_from": {completedAtFrom.Format("2006-01-02+15:04:05")}}
	}
}

func CompletedAtTo(completedAtTo time.Time) Filters {
	return func() url.Values {
		return url.Values{"completed_at_to": {completedAtTo.Format("2006-01-02+15:04:05")}}
	}
}

func OrderStatusID(orderStatusID string) Filters {
	return func() url.Values {
		return url.Values{"order_status_id": {orderStatusID}}
	}
}

func ExternalOrderID(externalOrderID string) Filters {
	return func() url.Values {
		return url.Values{"external_order_id": {externalOrderID}}
	}
}
