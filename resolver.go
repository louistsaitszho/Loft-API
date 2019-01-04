package loft

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/louistsaitszho/loft/models"
)

// Resolver is the entry point to how the query tree got processed
// dbAreYouSureAboutThis is temporary only. All queries should be prepared on server start-up (fail-fast)
type Resolver struct {
	dbAreYouSureAboutThis *sql.DB
	memberCountStmt       *sql.Stmt
	membersStmt           *sql.Stmt
	taskCountStmt         *sql.Stmt
	tasksStmt             *sql.Stmt
	eventCountStmt        *sql.Stmt
	eventsStmt            *sql.Stmt
	noteCountStml         *sql.Stmt
	notesStmt             *sql.Stmt
	loftStmt              *sql.Stmt
	joinRequestCountStmt  *sql.Stmt
	joinRequestStmt       *sql.Stmt
}

// NewResolver is essentially the constructor for Resolver. It reminds user that they should give Resolver a db to access
// TODO: unless it's COUNT don't use *
func NewResolver(db *sql.DB) *Resolver {
	memberCountStmt, mcErr := db.Prepare("SELECT COUNT(*) FROM loft.member WHERE loft.member.loft_id = $1;")
	if mcErr != nil {
		log.Panicf("Invalid query for member count: '%v'\n", mcErr)
	}
	membersStmt, mErr := db.Prepare("SELECT * FROM loft.member WHERE loft.member.loft_id = $1;")
	if mErr != nil {
		log.Panicf("Invalid query for members: '%v'\n", mErr)
	}
	taskCountStmt, tcErr := db.Prepare("SELECT COUNT(*) FROM loft.task WHERE loft.task.loft_id = $1;")
	if tcErr != nil {
		log.Panicf("Invalid query for task count: '%v'\n", tcErr)
	}
	tasksStmt, tErr := db.Prepare("SELECT * FROM loft.task WHERE loft.task.loft_id = $1;")
	if tErr != nil {
		log.Panicf("Invalid query for tasks: '%v'\n", tErr)
	}
	eventCountStmt, ecErr := db.Prepare("SELECT COUNT(*) FROM loft.event WHERE loft.event.loft_id = $1;")
	if ecErr != nil {
		log.Panicf("Invalid query for event count: '%v'\n", ecErr)
	}
	eventsStmt, eErr := db.Prepare("SELECT * FROM loft.event WHERE loft.event.loft_id = $1;")
	if eErr != nil {
		log.Panicf("Invalid query for events: '%v'\n", eErr)
	}
	noteCountStmt, ncErr := db.Prepare("SELECT COUNT(*) FROM loft.note WHERE loft.note.loft_id = $1;")
	if ncErr != nil {
		log.Panicf("Invalid query for note count: '%v'\n", ncErr)
	}
	notesStmt, nErr := db.Prepare("SELECT * FROM loft.note WHERE loft.note.loft_id = $1;")
	if nErr != nil {
		log.Panicf("Invalid query for note: '%v'\n", nErr)
	}
	loftStmt, lErr := db.Prepare("SELECT * FROM loft.loft WHERE loft.loft.id = $1;")
	if lErr != nil {
		log.Panicf("Invalid query for loft: '%v'\n", lErr)
	}
	joinRequestCountStmt, jrcErr := db.Prepare("SELECT COUNT(*) FROM loft.join_request WHERE loft.join_request.loft_id = $1;")
	if jrcErr != nil {
		log.Panicf("Invalid query for join request count: '%v'\n", jrcErr)
	}
	joinRequestStmt, jrErr := db.Prepare("SELECT * FROM loft.join_request WHERE loft.join_request.loft_id = $1;")
	if jrErr != nil {
		log.Panicf("Invalid query for join requests: '%v'\n", jrErr)
	}

	return &Resolver{
		dbAreYouSureAboutThis: db,
		memberCountStmt:       memberCountStmt,
		membersStmt:           membersStmt,
		taskCountStmt:         taskCountStmt,
		tasksStmt:             tasksStmt,
		eventCountStmt:        eventCountStmt,
		eventsStmt:            eventsStmt,
		noteCountStml:         noteCountStmt,
		notesStmt:             notesStmt,
		loftStmt:              loftStmt,
		joinRequestCountStmt:  joinRequestCountStmt,
		joinRequestStmt:       joinRequestStmt,
	}
}

// Loft returns a resolver that is able to resolve struct Loft
func (r *Resolver) Loft() LoftResolver {
	return &loftResolver{r}
}

// Mutation returns a resolver that is able to resolve mutations
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query returns a resolver that is able to resolve query
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type loftResolver struct{ *Resolver }

func (r *loftResolver) ID(ctx context.Context, obj *models.Loft) (string, error) {
	//TODO: replace it with external marshaler (https://gqlgen.com/reference/scalars/)
	return obj.ID.String(), nil
}
func (r *loftResolver) MembersCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.memberCountStmt.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) Members(ctx context.Context, obj *models.Loft) ([]Member, error) {
	panic("not implemented")
}
func (r *loftResolver) TasksCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.taskCountStmt.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) Tasks(ctx context.Context, obj *models.Loft) ([]Task, error) {
	panic("not implemented")
}
func (r *loftResolver) EventsCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.eventCountStmt.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) Events(ctx context.Context, obj *models.Loft) ([]Event, error) {
	panic("not implemented")
}
func (r *loftResolver) Notes(ctx context.Context, obj *models.Loft) ([]Note, error) {
	panic("not implemented")
}
func (r *loftResolver) NotesCount(ctx context.Context, obj *models.Loft) (int, error) {
	row := r.noteCountStml.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}
func (r *loftResolver) JoinRequestsCount(ctx context.Context, obj *models.Loft) (int, error) {
	panic("not implemented")
}
func (r *loftResolver) JoinRequests(ctx context.Context, obj *models.Loft) ([]JoinRequest, error) {
	row := r.joinRequestCountStmt.QueryRow(obj.ID)
	var count int
	if scanErr := row.Scan(&count); scanErr != nil {
		return -1, scanErr
	}
	return count, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTask(ctx context.Context, input NewTask) (Task, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateEvent(ctx context.Context, input NewEvent) (Event, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateRequest(ctx context.Context, input NewRequest) (JoinRequest, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLoft(ctx context.Context, input NewLoft) (models.Loft, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateLoftAndMember(ctx context.Context, input *NewLoftNewMember) (LoftAndFirstMember, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Lofts(ctx context.Context) ([]models.Loft, error) {
	query := "SELECT loft.loft.id, loft.loft.\"name\", loft.loft.join_code, loft.loft.created_at FROM loft.loft;"
	rows, queryErr := r.dbAreYouSureAboutThis.Query(query)
	if queryErr != nil {
		panic("not implemented: find out what can happen, and what is recoverable and how")
	}
	defer rows.Close()

	var lofts []models.Loft
	for rows.Next() {
		var (
			id              uuid.UUID
			name            string
			joinCode        string
			createdAtString string //TODO: parse it to time instead?
		)
		if scanErr := rows.Scan(&id, &name, &joinCode, &createdAtString); scanErr != nil {
			panic("not implemented: should i skip this one or go straight to nil, error?")
		}
		lofts = append(lofts, models.Loft{
			ID:       id,
			Name:     name,
			JoinCode: joinCode,
			//TODO: created at  missing
		})
	}

	return lofts, nil
}
func (r *queryResolver) Loft(ctx context.Context, id string) (*models.Loft, error) {
	panic("not implemented")
}
func (r *queryResolver) Echo(ctx context.Context) (Echo, error) {
	echo := Echo{
		Time:   time.Now().UTC().Format(time.RFC3339),
		Format: "RFC3339",
	}
	return echo, nil
}
