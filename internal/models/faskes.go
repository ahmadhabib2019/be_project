package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Faskes struct {
	db.ModelBase
	Id        uuid.UUID  `json:"id,omitempty" column:"name:id;type:uuid;primaryKey;nullable:false;default:gen_random_uuid()"`
	Name      *string    `json:"name,omitempty" column:"name:name;type:varchar;nullable"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestamp;nullable"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"faskes" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	FaskesThroughPoliFaske []*Faskes `json:"faskes_through_poli_faske,omitempty" join:"joinType:manyToMany;through:poli;sourcePrimaryKey:id;sourceForeignKey:faskes_id;targetPrimaryKey:id;targetForeign:faskes_id"`
	PoliFaskes             []*Poli   `json:"poli_faskes,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:faskes_id"`
}
