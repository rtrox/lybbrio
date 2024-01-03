package graph

import (
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/ent/schema/task_enums"
	"testing"

	"entgo.io/contrib/entgql"
	"github.com/stretchr/testify/require"
)

// TODO: Tests for all resolvers

func Test_Node(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Node")
	defer tc.Teardown()

	node, err := tc.Resolver.Query().Node(tc.UserCtx, tc.User().ID)
	require.NoError(err)

	user := node.(*ent.User)
	require.Equal(tc.User().ID, user.ID)
	require.Equal(tc.User().Username, user.Username)
	require.Equal(tc.User().Email, user.Email)

	author := tc.Client.Author.Create().
		SetName("some_author").
		SetSort("some_sort").
		SaveX(tc.AdminCtx)

	node, err = tc.Resolver.Query().Node(tc.UserCtx, author.ID)
	require.NoError(err)

	author = node.(*ent.Author)
	require.Equal(author.ID, author.ID)
	require.Equal(author.Name, author.Name)
	require.Equal(author.Sort, author.Sort)
}

func Test_Nodes(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Nodes")
	defer tc.Teardown()

	nodes, err := tc.Resolver.Query().Nodes(tc.UserCtx, []ksuid.ID{tc.User().ID})
	require.NoError(err)

	require.Len(nodes, 1)

	user := nodes[0].(*ent.User)
	require.Equal(tc.User().ID, user.ID)
	require.Equal(tc.User().Username, user.Username)
	require.Equal(tc.User().Email, user.Email)
}

func Test_Authors(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Authors")
	defer tc.Teardown()

	author := tc.Client.Author.Create().
		SetName("some_author").
		SetSort("some_sort").
		SaveX(tc.AdminCtx)

	author2 := tc.Client.Author.Create().
		SetName("some_author2").
		SetSort("some_sort2").
		SaveX(tc.AdminCtx)

	order := []*ent.AuthorOrder{
		{
			Direction: entgql.OrderDirectionDesc,
			Field:     ent.AuthorOrderFieldSort,
		},
	}
	authors, err := tc.Resolver.Query().Authors(tc.UserCtx, nil, nil, nil, nil, order, nil)
	require.NoError(err)
	require.Len(authors.Edges, 2)

	returnedAuthor2 := authors.Edges[0].Node
	require.Equal(author2.ID, returnedAuthor2.ID)
	require.Equal(author2.Name, returnedAuthor2.Name)
	require.Equal(author2.Sort, returnedAuthor2.Sort)

	returnedAuthor1 := authors.Edges[1].Node
	require.Equal(author.ID, returnedAuthor1.ID)
	require.Equal(author.Name, returnedAuthor1.Name)
	require.Equal(author.Sort, returnedAuthor1.Sort)

}

func Test_Books(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Books")
	defer tc.Teardown()

	book := tc.Client.Book.Create().
		SetTitle("some_title").
		SetSort("some_sort").
		SetPath("some_path").
		SaveX(tc.AdminCtx)

	book2 := tc.Client.Book.Create().
		SetTitle("some_title2").
		SetSort("some_sort2").
		SetPath("some_path2").
		SaveX(tc.AdminCtx)

	order := []*ent.BookOrder{
		{
			Direction: entgql.OrderDirectionDesc,
			Field:     ent.BookOrderFieldSort,
		},
	}
	books, err := tc.Resolver.Query().Books(tc.UserCtx, nil, nil, nil, nil, order, nil)
	require.NoError(err)
	require.Len(books.Edges, 2)

	returnedBook2 := books.Edges[0].Node
	require.Equal(book2.ID, returnedBook2.ID)
	require.Equal(book2.Title, returnedBook2.Title)
	require.Equal(book2.Sort, returnedBook2.Sort)

	returnedBook1 := books.Edges[1].Node
	require.Equal(book.ID, returnedBook1.ID)
	require.Equal(book.Title, returnedBook1.Title)
	require.Equal(book.Sort, returnedBook1.Sort)
}

func Test_Identifiers(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Identifiers")
	defer tc.Teardown()

	book := tc.Client.Book.Create().
		SetTitle("some_title").
		SetSort("some_sort").
		SetPath("some_path").
		SaveX(tc.AdminCtx)

	identifier := tc.Client.Identifier.Create().
		SetType("some_type").
		SetValue("some_value").
		SetBook(book).
		SaveX(tc.AdminCtx)

	identifier2 := tc.Client.Identifier.Create().
		SetType("some_type2").
		SetValue("some_value2").
		SetBook(book).
		SaveX(tc.AdminCtx)

	order := []*ent.IdentifierOrder{
		{
			Direction: entgql.OrderDirectionDesc,
			Field:     ent.IdentifierOrderFieldValue,
		},
	}
	identifiers, err := tc.Resolver.Query().Identifiers(tc.UserCtx, nil, nil, nil, nil, order, nil)
	require.NoError(err)
	require.Len(identifiers.Edges, 2)

	returnedIdentifier2 := identifiers.Edges[0].Node
	require.Equal(identifier2.ID, returnedIdentifier2.ID)
	require.Equal(identifier2.Value, returnedIdentifier2.Value)
	require.Equal(identifier2.Type, returnedIdentifier2.Type)

	returnedIdentifier1 := identifiers.Edges[1].Node
	require.Equal(identifier.ID, returnedIdentifier1.ID)
	require.Equal(identifier.Value, returnedIdentifier1.Value)
	require.Equal(identifier.Type, returnedIdentifier1.Type)
}

func Test_Languages(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Languages")
	defer tc.Teardown()

	language := tc.Client.Language.Create().
		SetCode("some_code").
		SaveX(tc.AdminCtx)

	language2 := tc.Client.Language.Create().
		SetCode("some_code2").
		SaveX(tc.AdminCtx)

	languages, err := tc.Resolver.Query().Languages(tc.UserCtx, nil, nil, nil, nil, nil, nil)
	require.NoError(err)
	require.Len(languages.Edges, 2)

	ids := []ksuid.ID{language.ID, language2.ID}
	for l := range languages.Edges {
		require.Contains(ids, languages.Edges[l].Node.ID)
	}
}

func Test_Publishers(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Publishers")
	defer tc.Teardown()

	publisher := tc.Client.Publisher.Create().
		SetName("some_name").
		SaveX(tc.AdminCtx)

	publisher2 := tc.Client.Publisher.Create().
		SetName("some_name2").
		SaveX(tc.AdminCtx)

	order := []*ent.PublisherOrder{
		{
			Direction: entgql.OrderDirectionDesc,
			Field:     ent.PublisherOrderFieldName,
		},
	}
	publishers, err := tc.Resolver.Query().Publishers(tc.UserCtx, nil, nil, nil, nil, order, nil)
	require.NoError(err)
	require.Len(publishers.Edges, 2)

	returnedPublisher2 := publishers.Edges[0].Node
	require.Equal(publisher2.ID, returnedPublisher2.ID)
	require.Equal(publisher2.Name, returnedPublisher2.Name)

	returnedPublisher1 := publishers.Edges[1].Node
	require.Equal(publisher.ID, returnedPublisher1.ID)
	require.Equal(publisher.Name, returnedPublisher1.Name)
}

// TODO Allow sort on series.Sort
func Test_SeriesSlice(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_SeriesSlice")
	defer tc.Teardown()

	series1 := tc.Client.Series.Create().
		SetName("some_name").
		SetSort("some_sort").
		SaveX(tc.AdminCtx)

	series2 := tc.Client.Series.Create().
		SetName("some_name2").
		SetSort("some_sort2").
		SaveX(tc.AdminCtx)

	order := []*ent.SeriesOrder{
		{
			Direction: entgql.OrderDirectionDesc,
			Field:     ent.SeriesOrderFieldSort,
		},
	}
	seriesSlices, err := tc.Resolver.Query().SeriesSlice(tc.UserCtx, nil, nil, nil, nil, order, nil)
	require.NoError(err)
	require.Len(seriesSlices.Edges, 2)

	returnedSeries2 := seriesSlices.Edges[0].Node
	require.Equal(series2.ID, returnedSeries2.ID)
	require.Equal(series2.Name, returnedSeries2.Name)

	returnedSeries1 := seriesSlices.Edges[1].Node
	require.Equal(series1.ID, returnedSeries1.ID)
	require.Equal(series1.Name, returnedSeries1.Name)
}

func Test_Shelves(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Shelves")
	defer tc.Teardown()

	shelf1 := tc.Client.Shelf.Create().
		SetName("some_name").
		SetUser(tc.User()).
		SaveX(tc.AdminCtx)

	shelf2 := tc.Client.Shelf.Create().
		SetName("some_name2").
		SetUser(tc.User()).
		SaveX(tc.AdminCtx)

	order := []*ent.ShelfOrder{
		{
			Direction: entgql.OrderDirectionDesc,
			Field:     ent.ShelfOrderFieldName,
		},
	}

	shelves, err := tc.Resolver.Query().Shelves(tc.UserCtx, nil, nil, nil, nil, order, nil)
	require.NoError(err)
	require.Len(shelves.Edges, 2)

	returnedShelf2 := shelves.Edges[0].Node
	require.Equal(shelf2.ID, returnedShelf2.ID)
	require.Equal(shelf2.Name, returnedShelf2.Name)

	returnedShelf1 := shelves.Edges[1].Node
	require.Equal(shelf1.ID, returnedShelf1.ID)
	require.Equal(shelf1.Name, returnedShelf1.Name)
}

func Test_Tags(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Tags")
	defer tc.Teardown()

	tag1 := tc.Client.Tag.Create().
		SetName("some_name").
		SaveX(tc.AdminCtx)

	tag2 := tc.Client.Tag.Create().
		SetName("some_name2").
		SaveX(tc.AdminCtx)

	order := []*ent.TagOrder{
		{
			Direction: entgql.OrderDirectionDesc,
			Field:     ent.TagOrderFieldName,
		},
	}

	tags, err := tc.Resolver.Query().Tags(tc.UserCtx, nil, nil, nil, nil, order, nil)
	require.NoError(err)
	require.Len(tags.Edges, 2)

	returnedTag2 := tags.Edges[0].Node
	require.Equal(tag2.ID, returnedTag2.ID)
	require.Equal(tag2.Name, returnedTag2.Name)

	returnedTag1 := tags.Edges[1].Node
	require.Equal(tag1.ID, returnedTag1.ID)
	require.Equal(tag1.Name, returnedTag1.Name)
}

func Test_Tasks(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Tasks")
	defer tc.Teardown()

	task1 := tc.Client.Task.Create().
		SetType(task_enums.TypeCalibreImport).
		SetUser(tc.User()).
		SaveX(tc.AdminCtx)

	task2 := tc.Client.Task.Create().
		SetType(task_enums.TypeNoOp).
		SetUser(tc.User()).
		SaveX(tc.AdminCtx)

	order := []*ent.TaskOrder{
		{
			Direction: entgql.OrderDirectionDesc,
			Field:     ent.TaskOrderFieldType,
		},
	}

	tasks, err := tc.Resolver.Query().Tasks(tc.UserCtx, nil, nil, nil, nil, order, nil)
	require.NoError(err)
	require.Len(tasks.Edges, 2)

	returnedTask2 := tasks.Edges[0].Node
	require.Equal(task2.ID, returnedTask2.ID)
	require.Equal(task2.Type, returnedTask2.Type)

	returnedTask1 := tasks.Edges[1].Node
	require.Equal(task1.ID, returnedTask1.ID)
	require.Equal(task1.Type, returnedTask1.Type)
}

func Test_Users(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	tc := setupTest(t, "Test_Users")
	defer tc.Teardown()

	perms := tc.Client.UserPermissions.Create().SaveX(tc.AdminCtx)
	user1 := tc.Client.User.Create().
		SetUsername("some_user_name_2").
		SetEmail("some_email_2").
		SetUserPermissions(perms).
		SaveX(tc.AdminCtx)

	perms2 := tc.Client.UserPermissions.Create().SaveX(tc.AdminCtx)
	user2 := tc.Client.User.Create().
		SetUsername("some_user_name_3").
		SetEmail("some_email_3").
		SetUserPermissions(perms2).
		SaveX(tc.AdminCtx)

	users, err := tc.Resolver.Query().Users(tc.UserCtx)
	require.NoError(err)
	require.Len(users, 3)

	ids := []ksuid.ID{tc.User().ID, user1.ID, user2.ID}

	for _, u := range users {
		require.Contains(ids, u.ID)
	}
}
