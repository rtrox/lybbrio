# Developing on Lybbrio

## Calibre Library

To access calibre, we are directly integrating with the `metadata.db` database calibre generates, using a repository pattern in [`internal/calibre`](internal/calibre/). These models are then exposed via lybbr.io's internal API server for use in the front-end, and in 3rd party integrations.

## The GraphQL API

The GraphQL API is implemented using [`ent`](https://entgo.io/docs/getting-started) and [`gqlgen`](https://gqlgen.com/), with GraphQL Resolvers Generated via [`entgql`](https://entgo.io/docs/graphql/), an integration between `ent` and `gqlgen` written by the maintainers of `ent`.

### Adding Fields and Edges to Existing Objects

Editing existing objects is done by editing the underlying Ent Schemas, which can be found in the [schema](./internal/ent/schema) directory. Docs on schema definition can be found in the [ent docs](https://entgo.io/docs/schema-def). Some basic style guidance:

- All Edges should generally point outward in the graph. As an example, most things describe books, which means almost all edges on the book object will be `From`. By contrast `User`s are generally owners of the things they related to, and therefore most edges on books will be `To`.
- 1:N and M:N edges to `Book` should almost always be Paginated, which is handled via [Relay Cursors](https://entgo.io/docs/tutorial-todo-gql-paginate/). Otherwise, Pagination should follow common sense -- if an edge might have tens or hundreds of nodes, it should be paginated. If it will always be in the single digits, it does not need to be paginated.
- Ordering is also generally expected to be common sense, if it seems like something you would want to order by, add an [OrderField](https://entgo.io/docs/tutorial-todo-gql-paginate/#add-annotations-to-schema).

Once you've made the edits you want to make, run `go generate .` in the root of the repo. This will generate all `ent` codegen, after which it will run the `entgql` generation to pull the changes into the GraphQL resolvers.

### Adding New Fields

Adding new fields is pretty similar, but you need to first generate a new object. This should be done from *inside the [internal](./internal/) folder:

```bash
go run -mod=mod entgo.io/ent/cmd/ent new <ObjectName>
# Example: go run -mod=mod entgo.io/ent/cmd/ent new Author
```

This will create a new ent object in the [schema](./internal/ent/schema/) directory. This project uses [segmentio/ksuid](https://github.com/segmentio/ksuid)s for IDs via a custom mixin, and [`entgql`](https://entgo.io/docs/getting-started) for codegen, so some `Annotations` and `Mixins` are likely required:

```go
func (Author) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(), // Adds pagination to this object's resolver.
		entgql.QueryField(), // Exposes this object in the root `Query` resolver.
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()), // Exposes Create<Object> and Update<Object> in the Mutation resolver.
		entgql.MultiOrder(), // Allows Ordering by multiple order fields. This should be set on any object with even a single `OrderField` to keep the API consistent if other `OrderField`s are added later.
	}
}

func (Author) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ksuid.MixinWithPrefix("atr"), // This is the prefix for IDs for this object. these must always be exactly 3 characters, and should be unique to this object.
	}
}
```

- **Pagination**: Pagination is handled via [Relay Cursors](https://entgo.io/docs/tutorial-todo-gql-paginate/). For paginated objects, consider carefully what fields should allow ordering. If at least one `OrderField` is added to the object, make sure you add the `entgql.MultiOrder()` annotation, so that the api remains consistent if we add other OrderFields later. This annotation allows for multiple sequential Ordering fields on a single query (i.e, sort by author then title).
- **Mutation**: valid options are `entgql.MutationCreate()`, `entgql.MutationUpdate()`, and `entgql.MutationDelete()`. Consider carefully what mutations should be allowed on this object.
- **KSUIDs**: IDs in this server are Prefixed [`ksuid`](https://github.com/segmentio/ksuid)s. The Prefixes allow us to follow the [Relay `Node` interface](https://entgo.io/docs/tutorial-todo-gql-node), which simplifies introspection of Object IDs. These prefixes are exactly three characters in length, and should be unique to the object. Any Edge Schemas should, if possible, be named `<first letter of left edge><first letter of right edge>x`, for example, `SeriesBook` which describes the edge between Series and Book, is `sbx`. You can see the custom mixin in the [`ksuid` folder](./internal/ent/schema/ksuid/) in the schema directory. To determine what prefixes are already being used, you can check the `prefixMap` in the generated [`ksuid.go`](./internal/ent/ksuid.go) library.

Once you've created the Schema for your new object, again, just run `go generate .` in the root of the repo.

## Gotchas

### Request ID should NOT be considered secure

Request IDs are meant to promote traceability, by creating an ID that can be used to correlate individual requests through the multiple distributed systems and multiple pieces of the codebase. As such, always assume the request id was set by a malicious actor, and do not use it for anything other than the tracing function for which it was intended.
