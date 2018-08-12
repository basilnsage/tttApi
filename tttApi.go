package tttApi

import (
	"fmt"
	graphql "github.com/graph-gophers/graphql-go"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	# define endpoints into our API
	type Query {
		viewGame(user: Player!, id: ID!): Game
		listGames(): [Game]!
	}
	# specify the mutations (updates) our API supports
	type Mutation {
		makePlayer(name: String!): Player!
		joinGame(user: Player!, id: ID!): Game!
		createGame(user: Player!): Game!
		makeMove(change: Move!): Move!
	}
	# specify the fields that the board needs to implement
	type Game {
		# unique id of the game
		id: ID!
		# list of players participating in game
		players: [Player]!
		# number of turns played
		turns: Int!
		# player whose turn is next
		whoseTurn: Player
	}
	# we gotta have players
	type Player {
		name: String!
		id: ID!
	}
	# we also need a way to encode a move
	type Move {
		# will probably want additional information to enforce proper turn order
		mover: Player!
		x: Int!
		y: Int!
	}
`

type Game struct {
	ID        graphql.ID
	Players   []Player
	Turns     int
	WhoseTurn Player
}

type Player struct {
	Name string
	ID   graphql.ID
}

type Move struct {
	Mover Player
	X     int
	Y     int
}

type Resolver struct{}

func (r *Resolver) viewGame(args struct {
	User Player
	ID   graphql.ID
}) *gameResolver {
}

func (r *Resolver) listGames() []*gameResolver {}

func (r *Resolver) makePlayer(args *struct{ Name string }) *playerResolver {}

func (r *Resolver) joinGame(args *struct {
	User Player
	ID   graphql.ID
}) *gameResolver {
}

func (r *Resolver) createGame(args *struct{ User Player }) *gameResolver {}

func (r *Resolver) makeMove(args *struct{ Change Move }) *moveResolver {}

type gameResolver struct {
	g *Game
}

func (r *gameResolver) ID() graphql.ID {
	return r.g.ID
}

type playerResolver struct {
	p *Player
}

type moveResolver struct {
	m *Move
}
