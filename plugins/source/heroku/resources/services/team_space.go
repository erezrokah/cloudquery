// Code generated by codegen; DO NOT EDIT.

package services

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func TeamSpaces() *schema.Table {
	return &schema.Table{
		Name:        "heroku_team_spaces",
		Description: `https://devcenter.heroku.com/articles/platform-api-reference#team-space`,
		Resolver:    fetchTeamSpaces,
		Columns: []schema.Column{
			{
				Name:     "cidr",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CIDR"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "data_cidr",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataCIDR"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "organization",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Organization"),
			},
			{
				Name:     "region",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "shield",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Shield"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "team",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Team"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchTeamSpaces(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	items := make([]heroku.Team, 0, 10)
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.TeamList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		items = append(items, v...)
	}

	for _, it := range items {
		nextRange = &heroku.ListRange{
			Field: "id",
			Max:   1000,
		}
		// Roundtripper middleware in client/pagination.go
		// sets the nextRange value after each request
		for nextRange.Max != 0 {
			ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
			v, err := c.Heroku.TeamSpaceList(ctxWithRange, it.ID, nextRange)
			if err != nil {
				return errors.WithStack(err)
			}
			res <- v
		}
	}
	return nil
}
