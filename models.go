package almaclient

// type Meta struct {
// 	ID     string `json:"_id,omitempty"`
// 	Name   string `json:"name,omitempty"`
// 	Alias  string `json:"alias,omitempty"`
// 	Weapon string `json:"weapon,omitempty"`
// }

type ServiceCatalog struct {
	Services []*Service `json:"services"`
}

type Service struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Meta   *Meta  `json:"meta"`
}

type Meta struct {
	Engine       string   `json:"engine"`
	Team         string   `json:"team"`
	Environments []string `json:"environments"`
	SubApps      []string `json:"subapps"`

	PubSub   PubSub   `json:"pubsub"`
	CloudSql CloudSql `json:"cloudsql"`
	GCS      Gcs      `json:"gcs"`
}

type PubSub struct {
	Enabled bool    `json:"enabled"`
	Topics  []Topic `json:"topics"`
}

type Topic struct {
	Name          string         `json:"name"`
	Subscriptions []Subscription `json:"subscriptions"`
}

type Subscription struct {
	Name string `json:"name"`
}

type CloudSql struct {
	Enabled   bool       `json:"enabled"`
	Databases []Database `json:"databases"`
}

type Database struct {
	Name string `json:"name"`
}

type Gcs struct {
	Enabled bool     `json:"enabled"`
	Buckets []Bucket `json:"buckets"`
}

type Bucket struct {
	Name        string   `json:"name"`
	Public      bool     `json:"public"`
	Permissions []string `json:"roles"`
}
