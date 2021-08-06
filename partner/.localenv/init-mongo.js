db.createCollection('partners')

db.partners.createIndex({ document: 1 }, { unique: true })
db.partners.createIndex({ address: "2dsphere" })
