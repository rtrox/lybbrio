package graph

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"

	"lybbrio/internal/db"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/argon2id"
	"lybbrio/internal/ent/schema/permissions"
	"lybbrio/internal/ent/user"
	"lybbrio/internal/viewer"
)

type graphqlTestContext struct {
	Client       *client.Client
	DBClient     *ent.Client
	AdminCtx     context.Context
	UserCtx      context.Context
	teardownFunc func()
}

func (tc graphqlTestContext) User() *ent.User {
	uid, _ := viewer.FromContext(tc.UserCtx).UserID()
	user := tc.DBClient.User.Query().Where(user.ID(uid)).OnlyX(tc.UserCtx)
	return user
}

func (tc graphqlTestContext) Teardown() {
	tc.teardownFunc()
	tc.DBClient.Close()
}

func testMiddleware(user *ent.User, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := viewer.NewContext(r.Context(), user.ID, permissions.NewPermissions(permissions.Admin))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func setupGQLTest(t *testing.T, testName string, teardown ...func()) graphqlTestContext {
	dbClient := db.OpenTest(t, testName)

	adminCtx := viewer.NewSystemAdminContext(context.Background())
	perms := dbClient.UserPermissions.Create().SetAdmin(true).SaveX(adminCtx)
	user := dbClient.User.Create().
		SetUsername("some_user_name").
		SetEmail("some_email").
		SetUserPermissions(perms).
		SaveX(adminCtx)

	userCtx := viewer.NewContext(
		context.Background(),
		user.ID,
		permissions.NewPermissions(permissions.Admin),
	)
	argon2idConfig := argon2id.Config{
		Memory:      64,
		Iterations:  1,
		Parallelism: 1,
		SaltLen:     16,
		KeyLen:      32,
	}
	handler := handler.NewDefaultServer(NewSchema(dbClient, argon2idConfig))
	handler.Use(
		entgql.Transactioner{TxOpener: dbClient},
	)
	client := client.New(testMiddleware(user, handler))

	ret := graphqlTestContext{
		Client:   client,
		DBClient: dbClient,
		AdminCtx: adminCtx,
		UserCtx:  userCtx,
	}
	if len(teardown) > 0 {
		ret.teardownFunc = teardown[0]
	} else {
		ret.teardownFunc = func() {}
	}
	return ret
}

func Test_CreateBook(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateBook")
	defer tc.Teardown()

	var resp struct {
		CreateBook struct {
			ID            string
			Title         string
			Sort          string
			Path          string
			ISBN          string
			SeriesIndex   float64
			Description   string
			PublishedDate string
		}
	}

	mutation := `
	mutation {
		createBook(input: {
			title:"some_title",
			sort:"some_sort",
			path:"some_path",
			isbn:"some_isbn",
			seriesIndex:1.0,
			description: "i am a book",
			publishedDate: "2020-01-01T00:00:00Z",
		}){
		  id
		  title
		  sort
		  path
		  isbn
		  seriesIndex
		  description
		  publishedDate
		}
	  }`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateBook.ID)
	require.Equal("some_title", resp.CreateBook.Title)
	require.Equal("some_sort", resp.CreateBook.Sort)
	require.Equal("some_path", resp.CreateBook.Path)
	require.Equal("some_isbn", resp.CreateBook.ISBN)
	require.Equal(1.0, resp.CreateBook.SeriesIndex)
	require.Equal("i am a book", resp.CreateBook.Description)
	require.Equal("2020-01-01T00:00:00Z", resp.CreateBook.PublishedDate)
}

func Test_CreateBook_With_Author_And_Series(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateBook_With_Author_And_Series")
	defer tc.Teardown()

	author := tc.DBClient.Author.Create().
		SetName("some_author").
		SetSort("some_author_sort").
		SaveX(tc.AdminCtx)
	series := tc.DBClient.Series.Create().
		SetName("some_series").
		SetSort("some_series_sort").
		SaveX(tc.AdminCtx)

	var resp struct {
		CreateBook struct {
			ID      string
			Title   string
			Sort    string
			Path    string
			Authors []struct {
				ID   string
				Name string
			}
			Series []struct {
				ID   string
				Name string
			}
		}
	}

	mutation := `
	mutation f{
		createBook(input: {
			title:"some_title",
			sort:"some_sort",
			path:"some_path",
			authorIDs: ["%s"],
			seriesIDs: ["%s"],
		}){
			id
			title
			sort
			path
			authors {
				id
				name
			}
			series {
				id
				name
			}
		}
	}`
	err := tc.Client.Post(fmt.Sprintf(mutation, author.ID, series.ID), &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateBook.ID)
	require.Equal("some_title", resp.CreateBook.Title)
	require.Equal("some_sort", resp.CreateBook.Sort)
	require.Equal("some_path", resp.CreateBook.Path)
	require.Len(resp.CreateBook.Authors, 1)
	require.Equal(author.ID.String(), resp.CreateBook.Authors[0].ID)
	require.Equal("some_author", resp.CreateBook.Authors[0].Name)
	require.Len(resp.CreateBook.Series, 1)
	require.Equal(series.ID.String(), resp.CreateBook.Series[0].ID)
	require.Equal("some_series", resp.CreateBook.Series[0].Name)
}

func Test_UpdateBook(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdateBook")
	defer tc.Teardown()

	book := tc.DBClient.Book.Create().
		SetTitle("some_title").
		SetSort("some_sort").
		SetPath("some_path").
		SetIsbn("some_isbn").
		SetSeriesIndex(1.0).
		SetDescription("i am a book").
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdateBook struct {
			ID            string
			Title         string
			Sort          string
			Path          string
			ISBN          string
			SeriesIndex   float64
			Description   string
			PublishedDate string
		}
	}

	mutation := `
	mutation {
		updateBook(id: "%s", input: {
			title:"some_other_title",
			sort:"some_other_sort",
			path:"some_other_path",
			isbn:"some_other_isbn",
			seriesIndex:2.0,
			description: "i am another book",
			publishedDate: "2020-01-02T00:00:00Z",
		}){
		  id
		  title
		  sort
		  path
		  isbn
		  seriesIndex
		  description
		  publishedDate
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, book.ID), &resp)
	require.NoError(err)

	require.Equal(book.ID.String(), resp.UpdateBook.ID)
	require.Equal("some_other_title", resp.UpdateBook.Title)
	require.Equal("some_other_sort", resp.UpdateBook.Sort)
	require.Equal("some_other_path", resp.UpdateBook.Path)
	require.Equal("some_other_isbn", resp.UpdateBook.ISBN)
	require.Equal(2.0, resp.UpdateBook.SeriesIndex)
	require.Equal("i am another book", resp.UpdateBook.Description)
	require.Equal("2020-01-02T00:00:00Z", resp.UpdateBook.PublishedDate)
}

func Test_CreateAuthor(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateAuthor")
	defer tc.Teardown()

	var resp struct {
		CreateAuthor struct {
			ID   string
			Name string
			Sort string
		}
	}

	mutation := `
	mutation {
		createAuthor(input: {
			name:"some_name",
			sort:"some_sort",
		}){
		  id
		  name
		  sort
		}
	  }`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateAuthor.ID)
	require.Equal("some_name", resp.CreateAuthor.Name)
	require.Equal("some_sort", resp.CreateAuthor.Sort)
}

func Test_UpdateAuthor(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdateAuthor")
	defer tc.Teardown()

	author := tc.DBClient.Author.Create().
		SetName("some_name").
		SetSort("some_sort").
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdateAuthor struct {
			ID   string
			Name string
			Sort string
		}
	}

	mutation := `
	mutation {
		updateAuthor(id: "%s", input: {
			name:"some_other_name",
			sort:"some_other_sort",
		}){
		  id
		  name
		  sort
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, author.ID), &resp)
	require.NoError(err)

	require.Equal(author.ID.String(), resp.UpdateAuthor.ID)
	require.Equal("some_other_name", resp.UpdateAuthor.Name)
	require.Equal("some_other_sort", resp.UpdateAuthor.Sort)
}

func Test_CreateShelf(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateShelf")
	defer tc.Teardown()

	var resp struct {
		CreateShelf struct {
			ID          string
			Name        string
			Description string
			Public      bool
			User        struct {
				ID string
			}
		}
	}

	mutation := `
	mutation {
		createShelf(input: {
			name:"some_name",
			description:"some_description",
			public:true,
		}){
		  id
		  name
		  description
		  public
		  user {
			id
		  }
		}
	  }`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateShelf.ID)
	require.Equal("some_name", resp.CreateShelf.Name)
	require.Equal("some_description", resp.CreateShelf.Description)
	require.Equal(tc.User().ID.String(), resp.CreateShelf.User.ID)
	require.True(resp.CreateShelf.Public)
}

func Test_UpdateShelf(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdateShelf")
	defer tc.Teardown()

	shelf := tc.DBClient.Shelf.Create().
		SetName("some_name").
		SetDescription("some_description").
		SetPublic(true).
		SetUser(tc.User()).
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdateShelf struct {
			ID          string
			Name        string
			Description string
			Public      bool
		}
	}

	mutation := `
	mutation {
		updateShelf(id: "%s", input: {
			name:"some_other_name",
			description:"some_other_description",
			public:false,
		}){
		  id
		  name
		  description
		  public
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, shelf.ID), &resp)
	require.NoError(err)

	require.Equal(shelf.ID.String(), resp.UpdateShelf.ID)
	require.Equal("some_other_name", resp.UpdateShelf.Name)
	require.Equal("some_other_description", resp.UpdateShelf.Description)
	require.False(resp.UpdateShelf.Public)
}

func Test_CreateTag(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateTag")
	defer tc.Teardown()

	var resp struct {
		CreateTag struct {
			ID   string
			Name string
		}
	}

	mutation := `
	mutation {
		createTag(input: {
			name:"some_name",
		}){
		  id
		  name
		}
	  }`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateTag.ID)
	require.Equal("some_name", resp.CreateTag.Name)
}

func Test_UpdateTag(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdateTag")
	defer tc.Teardown()

	tag := tc.DBClient.Tag.Create().
		SetName("some_name").
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdateTag struct {
			ID   string
			Name string
		}
	}

	mutation := `
	mutation {
		updateTag(id: "%s", input: {
			name:"some_other_name",
		}){
		  id
		  name
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, tag.ID), &resp)
	require.NoError(err)

	require.Equal(tag.ID.String(), resp.UpdateTag.ID)
	require.Equal("some_other_name", resp.UpdateTag.Name)
}

func Test_CreatePublisher(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreatePublisher")
	defer tc.Teardown()

	var resp struct {
		CreatePublisher struct {
			ID   string
			Name string
		}
	}

	mutation := `
	mutation {
		createPublisher(input: {
			name:"some_name",
		}){
		  id
		  name
		}
	  }`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreatePublisher.ID)
	require.Equal("some_name", resp.CreatePublisher.Name)
}

func Test_UpdatePublisher(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdatePublisher")
	defer tc.Teardown()

	publisher := tc.DBClient.Publisher.Create().
		SetName("some_name").
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdatePublisher struct {
			ID   string
			Name string
		}
	}

	mutation := `
	mutation {
		updatePublisher(id: "%s", input: {
			name:"some_other_name",
		}){
		  id
		  name
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, publisher.ID), &resp)
	require.NoError(err)

	require.Equal(publisher.ID.String(), resp.UpdatePublisher.ID)
	require.Equal("some_other_name", resp.UpdatePublisher.Name)
}

func Test_CreateLanguage(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateLanguage")
	defer tc.Teardown()

	var resp struct {
		CreateLanguage struct {
			ID   string
			Code string
		}
	}

	mutation := `
	mutation {
		createLanguage(input: {
			code:"en",
		}){
		  id
		  code
		}
	  }`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateLanguage.ID)
	require.Equal("en", resp.CreateLanguage.Code)
}

func Test_UpdateLanguage(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdateLanguage")
	defer tc.Teardown()

	language := tc.DBClient.Language.Create().
		SetCode("en").
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdateLanguage struct {
			ID   string
			Code string
		}
	}

	mutation := `
	mutation {
		updateLanguage(id: "%s", input: {
			code:"de",
		}){
		  id
		  code
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, language.ID), &resp)
	require.NoError(err)

	require.Equal(language.ID.String(), resp.UpdateLanguage.ID)
	require.Equal("de", resp.UpdateLanguage.Code)
}

func Test_CreateSeries(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateSeries")
	defer tc.Teardown()

	var resp struct {
		CreateSeries struct {
			ID   string
			Name string
			Sort string
		}
	}

	mutation := `
	mutation {
		createSeries(input: {
			name:"some_name",
			sort:"some_sort",
		}){
		  id
		  name
		  sort
		}
	  }`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateSeries.ID)
	require.Equal("some_name", resp.CreateSeries.Name)
	require.Equal("some_sort", resp.CreateSeries.Sort)
}

func Test_UpdateSeries(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdateSeries")
	defer tc.Teardown()

	series := tc.DBClient.Series.Create().
		SetName("some_name").
		SetSort("some_sort").
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdateSeries struct {
			ID   string
			Name string
			Sort string
		}
	}

	mutation := `
	mutation {
		updateSeries(id: "%s", input: {
			name:"some_other_name",
			sort:"some_other_sort",
		}){
		  id
		  name
		  sort
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, series.ID), &resp)
	require.NoError(err)

	require.Equal(series.ID.String(), resp.UpdateSeries.ID)
	require.Equal("some_other_name", resp.UpdateSeries.Name)
	require.Equal("some_other_sort", resp.UpdateSeries.Sort)
}

func Test_CreateIdentifier(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateIdentifier")
	defer tc.Teardown()

	var resp struct {
		CreateIdentifier struct {
			ID    string
			Type  string
			Value string
		}
	}

	book := tc.DBClient.Book.Create().
		SetTitle("some_title").
		SetSort("some_sort").
		SetPath("some_path").
		SaveX(tc.AdminCtx)

	mutation := `
	mutation {
		createIdentifier(input: {
			type:"some_source",
			value:"some_value",
			bookID:"%s",
		}){
		  id
		  type
		  value
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, book.ID), &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateIdentifier.ID)
	require.Equal("some_source", resp.CreateIdentifier.Type)
	require.Equal("some_value", resp.CreateIdentifier.Value)
}

func Test_UpdateIdentifier(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdateIdentifier")
	defer tc.Teardown()
	book := tc.DBClient.Book.Create().
		SetTitle("some_title").
		SetSort("some_sort").
		SetPath("some_path").
		SaveX(tc.AdminCtx)

	identifier := tc.DBClient.Identifier.Create().
		SetType("some_source").
		SetValue("some_value").
		SetBook(book).
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdateIdentifier struct {
			ID    string
			Type  string
			Value string
		}
	}

	mutation := `
	mutation {
		updateIdentifier(id: "%s", input: {
			type:"some_other_source",
			value:"some_other_value",
		}){
		  id
		  type
		  value
		}
	  }`
	err := tc.Client.Post(fmt.Sprintf(mutation, identifier.ID), &resp)
	require.NoError(err)

	require.Equal(identifier.ID.String(), resp.UpdateIdentifier.ID)
	require.Equal("some_other_source", resp.UpdateIdentifier.Type)
	require.Equal("some_other_value", resp.UpdateIdentifier.Value)
}

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateUser")
	defer tc.Teardown()

	var resp struct {
		CreateUser struct {
			ID              string
			Username        string
			Email           string
			UserPermissions struct {
				ID    string
				Admin bool
			}
		}
	}
	mutation := `
	mutation {
		createUser(input: {
			username:"some_username",
			email:"an_email",
			userPermissions: {
				admin: true,
			}
		}){
		  id
		  username
		  email
		  userPermissions {
			admin
		  }
		}
	  }`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateUser.ID)
	require.Equal("some_username", resp.CreateUser.Username)
	require.Equal("an_email", resp.CreateUser.Email)
	require.Equal(true, resp.CreateUser.UserPermissions.Admin)

}

func Test_UpdateUser(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_UpdateUser")
	defer tc.Teardown()

	perms := tc.DBClient.UserPermissions.Create().SaveX(tc.AdminCtx)
	user := tc.DBClient.User.Create().
		SetUsername("some_username").
		SetEmail("an_email").
		SetUserPermissions(perms).
		SaveX(tc.AdminCtx)

	var resp struct {
		UpdateUser struct {
			ID       string
			Username string
			Email    string
		}
	}

	mutation := `
	mutation {
		updateUser(id: "%s", input: {
			username:"some_other_username",
			email:"some_other_email",
		}){
			id
			username
			email
		}
	}`
	err := tc.Client.Post(fmt.Sprintf(mutation, user.ID), &resp)
	require.NoError(err)

	require.Equal(user.ID.String(), resp.UpdateUser.ID)
	require.Equal("some_other_username", resp.UpdateUser.Username)
	require.Equal("some_other_email", resp.UpdateUser.Email)
}

func Test_CreateTask(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupGQLTest(t, "Test_CreateTask")
	defer tc.Teardown()

	var resp struct {
		CreateTask struct {
			ID         string
			Type       string
			Status     string
			CreateTime string
			UpdateTime string
			UserID     string
		}
	}

	mutation := `
	mutation {
		createTask(input: {
			type: noop,
		}){
			id
			type
			status
			createTime
			updateTime
			userID
		}
	}`
	err := tc.Client.Post(mutation, &resp)
	require.NoError(err)

	require.NotEmpty(resp.CreateTask.ID)
	require.Equal("noop", resp.CreateTask.Type)
	require.Equal("pending", resp.CreateTask.Status)
	require.NotEmpty(resp.CreateTask.CreateTime)
	require.NotEmpty(resp.CreateTask.UpdateTime)
	require.Equal(tc.User().ID.String(), resp.CreateTask.UserID)
}
