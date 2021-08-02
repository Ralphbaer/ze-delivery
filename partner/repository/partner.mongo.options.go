package repository

import (
	"github.com/Ralphbaer/ze-delivery/common/zmongo"
)

// WithPartnerQuery is a custom mongodb query for Find events operations
var WithPartnerQuery = func(p *PartnerQuery) zmongo.MongoQueryBuilderOption {
	return func(b *zmongo.MongoQueryBuilder) {
		/*if strings.TrimSpace(q.UserID) != "" {
			b.Filter["attendees.id"] = q.UserID
		}
		if q.Types != nil {
			b.Filter["type"] = bson.M{"$in": q.Types}
		}
		if q.Statuses != nil {
			b.Filter["current.status.value"] = bson.M{"$in": q.Statuses}
		}
		if q.OrgNames != nil {
			orgNamesStr := strings.Join(q.OrgNames, ",")
			q.OrgNames = strings.Split(orgNamesStr, ",")
			b.Filter["attendees.metadata.orgName"] = bson.M{"$in": q.OrgNames}
		}
		if !q.Date.IsZero() {
			filter := bson.M{}
			filter["$gte"] = q.Date.Truncate(24 * time.Hour)
			filter["$lte"] = q.Date.Truncate(24 * time.Hour).Add(time.Hour*time.Duration(23) + time.Minute*time.Duration(59) + time.Second*time.Duration(59))
			b.Filter["current.startDate"] = filter
		}
		if q.Active != nil {
			b.Filter["active"] = q.Active
		}*/
	}
}

