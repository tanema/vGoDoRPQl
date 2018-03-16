// This file was generated by github.com/vektah/gqlgen, DO NOT EDIT

package gql

import (
	"bytes"
	context "context"
	strconv "strconv"

	data "github.com/tanema/vGoDoRPQl/api/data"
	graphql "github.com/vektah/gqlgen/graphql"
	errors "github.com/vektah/gqlgen/neelance/errors"
	introspection "github.com/vektah/gqlgen/neelance/introspection"
	query "github.com/vektah/gqlgen/neelance/query"
	schema "github.com/vektah/gqlgen/neelance/schema"
)

func MakeExecutableSchema(resolvers Resolvers) graphql.ExecutableSchema {
	return &executableSchema{resolvers}
}

type Resolvers interface {
	Mutation_createTodo(ctx context.Context, text string, done *bool) (data.Todo, error)
	Mutation_updateTodos(ctx context.Context, ids []int, changes map[string]interface{}) ([]data.Todo, error)
	Mutation_deleteTodos(ctx context.Context, ids []int) ([]data.Todo, error)
	Query_todo(ctx context.Context, id int) (*data.Todo, error)
	Query_todos(ctx context.Context, status *string) ([]data.Todo, error)
}

type executableSchema struct {
	resolvers Resolvers
}

func (e *executableSchema) Schema() *schema.Schema {
	return parsedSchema
}

func (e *executableSchema) Query(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation, recover graphql.RecoverFunc) *graphql.Response {
	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx, recover: recover}

	data := ec._Query(op.Selections)
	var buf bytes.Buffer
	data.MarshalGQL(&buf)

	return &graphql.Response{
		Data:   buf.Bytes(),
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Mutation(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation, recover graphql.RecoverFunc) *graphql.Response {
	ec := executionContext{resolvers: e.resolvers, variables: variables, doc: doc, ctx: ctx, recover: recover}

	data := ec._Mutation(op.Selections)
	var buf bytes.Buffer
	data.MarshalGQL(&buf)

	return &graphql.Response{
		Data:   buf.Bytes(),
		Errors: ec.Errors,
	}
}

func (e *executableSchema) Subscription(ctx context.Context, doc *query.Document, variables map[string]interface{}, op *query.Operation, recover graphql.RecoverFunc) func() *graphql.Response {
	return graphql.OneShot(&graphql.Response{Errors: []*errors.QueryError{{Message: "subscriptions are not supported"}}})
}

type executionContext struct {
	errors.Builder
	resolvers Resolvers
	variables map[string]interface{}
	doc       *query.Document
	ctx       context.Context
	recover   graphql.RecoverFunc
}

var mutationImplementors = []string{"Mutation"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Mutation(sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, mutationImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createTodo":
			out.Values[i] = ec._Mutation_createTodo(field)
		case "updateTodos":
			out.Values[i] = ec._Mutation_updateTodos(field)
		case "deleteTodos":
			out.Values[i] = ec._Mutation_deleteTodos(field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Mutation_createTodo(field graphql.CollectedField) graphql.Marshaler {
	var arg0 string
	if tmp, ok := field.Args["text"]; ok {
		var err error

		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	var arg1 *bool
	if tmp, ok := field.Args["done"]; ok {
		var err error
		var ptr1 bool

		ptr1, err = graphql.UnmarshalBoolean(tmp)
		arg1 = &ptr1
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res, err := ec.resolvers.Mutation_createTodo(ec.ctx, arg0, arg1)
	if err != nil {
		ec.Error(err)
		return graphql.Null
	}
	return ec._Todo(field.Selections, &res)
}

func (ec *executionContext) _Mutation_updateTodos(field graphql.CollectedField) graphql.Marshaler {
	var arg0 []int
	if tmp, ok := field.Args["ids"]; ok {
		var err error
		rawIf1 := tmp.([]interface{})
		arg0 = make([]int, len(rawIf1))
		for idx1 := range rawIf1 {

			arg0[idx1], err = graphql.UnmarshalInt(rawIf1[idx1])
		}
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	var arg1 map[string]interface{}
	if tmp, ok := field.Args["changes"]; ok {
		var err error

		arg1, err = graphql.UnmarshalMap(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res, err := ec.resolvers.Mutation_updateTodos(ec.ctx, arg0, arg1)
	if err != nil {
		ec.Error(err)
		return graphql.Null
	}
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler { return ec._Todo(field.Selections, &res[idx1]) }())
	}
	return arr1
}

func (ec *executionContext) _Mutation_deleteTodos(field graphql.CollectedField) graphql.Marshaler {
	var arg0 []int
	if tmp, ok := field.Args["ids"]; ok {
		var err error
		rawIf1 := tmp.([]interface{})
		arg0 = make([]int, len(rawIf1))
		for idx1 := range rawIf1 {

			arg0[idx1], err = graphql.UnmarshalInt(rawIf1[idx1])
		}
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res, err := ec.resolvers.Mutation_deleteTodos(ec.ctx, arg0)
	if err != nil {
		ec.Error(err)
		return graphql.Null
	}
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler { return ec._Todo(field.Selections, &res[idx1]) }())
	}
	return arr1
}

var queryImplementors = []string{"Query"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Query(sel []query.Selection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, queryImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "todo":
			out.Values[i] = ec._Query_todo(field)
		case "todos":
			out.Values[i] = ec._Query_todos(field)
		case "__schema":
			out.Values[i] = ec._Query___schema(field)
		case "__type":
			out.Values[i] = ec._Query___type(field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Query_todo(field graphql.CollectedField) graphql.Marshaler {
	var arg0 int
	if tmp, ok := field.Args["id"]; ok {
		var err error

		arg0, err = graphql.UnmarshalInt(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		res, err := ec.resolvers.Query_todo(ec.ctx, arg0)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		if res == nil {
			return graphql.Null
		}
		return ec._Todo(field.Selections, res)
	})
}

func (ec *executionContext) _Query_todos(field graphql.CollectedField) graphql.Marshaler {
	var arg0 *string
	if tmp, ok := field.Args["status"]; ok {
		var err error
		var ptr1 string

		ptr1, err = graphql.UnmarshalString(tmp)
		arg0 = &ptr1
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	return graphql.Defer(func() (ret graphql.Marshaler) {
		defer func() {
			if r := recover(); r != nil {
				userErr := ec.recover(r)
				ec.Error(userErr)
				ret = graphql.Null
			}
		}()
		res, err := ec.resolvers.Query_todos(ec.ctx, arg0)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
		arr1 := graphql.Array{}
		for idx1 := range res {
			arr1 = append(arr1, func() graphql.Marshaler { return ec._Todo(field.Selections, &res[idx1]) }())
		}
		return arr1
	})
}

func (ec *executionContext) _Query___schema(field graphql.CollectedField) graphql.Marshaler {
	res := ec.introspectSchema()
	if res == nil {
		return graphql.Null
	}
	return ec.___Schema(field.Selections, res)
}

func (ec *executionContext) _Query___type(field graphql.CollectedField) graphql.Marshaler {
	var arg0 string
	if tmp, ok := field.Args["name"]; ok {
		var err error

		arg0, err = graphql.UnmarshalString(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res := ec.introspectType(arg0)
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

var todoImplementors = []string{"Todo"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) _Todo(sel []query.Selection, obj *data.Todo) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, todoImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Todo")
		case "id":
			out.Values[i] = ec._Todo_id(field, obj)
		case "text":
			out.Values[i] = ec._Todo_text(field, obj)
		case "done":
			out.Values[i] = ec._Todo_done(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) _Todo_id(field graphql.CollectedField, obj *data.Todo) graphql.Marshaler {
	res := obj.ID
	return graphql.MarshalInt(res)
}

func (ec *executionContext) _Todo_text(field graphql.CollectedField, obj *data.Todo) graphql.Marshaler {
	res := obj.Text
	return graphql.MarshalString(res)
}

func (ec *executionContext) _Todo_done(field graphql.CollectedField, obj *data.Todo) graphql.Marshaler {
	res := obj.Done
	return graphql.MarshalBoolean(res)
}

var __DirectiveImplementors = []string{"__Directive"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Directive(sel []query.Selection, obj *introspection.Directive) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __DirectiveImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(field, obj)
		case "description":
			out.Values[i] = ec.___Directive_description(field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(field, obj)
		case "args":
			out.Values[i] = ec.___Directive_args(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Directive_name(field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Directive_description(field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Directive_locations(field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Locations()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler { return graphql.MarshalString(res[idx1]) }())
	}
	return arr1
}

func (ec *executionContext) ___Directive_args(field graphql.CollectedField, obj *introspection.Directive) graphql.Marshaler {
	res := obj.Args()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(field.Selections, res[idx1])
		}())
	}
	return arr1
}

var __EnumValueImplementors = []string{"__EnumValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___EnumValue(sel []query.Selection, obj *introspection.EnumValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __EnumValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(field, obj)
		case "description":
			out.Values[i] = ec.___EnumValue_description(field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___EnumValue_name(field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_description(field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.IsDeprecated()
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(field graphql.CollectedField, obj *introspection.EnumValue) graphql.Marshaler {
	res := obj.DeprecationReason()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __FieldImplementors = []string{"__Field"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Field(sel []query.Selection, obj *introspection.Field) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __FieldImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(field, obj)
		case "description":
			out.Values[i] = ec.___Field_description(field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(field, obj)
		case "type":
			out.Values[i] = ec.___Field_type(field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(field, obj)
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Field_name(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Field_description(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Field_args(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Args()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Field_type(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.Type()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.IsDeprecated()
	return graphql.MarshalBoolean(res)
}

func (ec *executionContext) ___Field_deprecationReason(field graphql.CollectedField, obj *introspection.Field) graphql.Marshaler {
	res := obj.DeprecationReason()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __InputValueImplementors = []string{"__InputValue"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___InputValue(sel []query.Selection, obj *introspection.InputValue) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __InputValueImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(field, obj)
		case "description":
			out.Values[i] = ec.___InputValue_description(field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(field, obj)
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___InputValue_name(field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Name()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_description(field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___InputValue_type(field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.Type()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(field graphql.CollectedField, obj *introspection.InputValue) graphql.Marshaler {
	res := obj.DefaultValue()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

var __SchemaImplementors = []string{"__Schema"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Schema(sel []query.Selection, obj *introspection.Schema) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __SchemaImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(field, obj)
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(field, obj)
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Schema_types(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.Types()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Schema_queryType(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.QueryType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.MutationType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.SubscriptionType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(field graphql.CollectedField, obj *introspection.Schema) graphql.Marshaler {
	res := obj.Directives()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Directive(field.Selections, res[idx1])
		}())
	}
	return arr1
}

var __TypeImplementors = []string{"__Type"}

// nolint: gocyclo, errcheck, gas, goconst
func (ec *executionContext) ___Type(sel []query.Selection, obj *introspection.Type) graphql.Marshaler {
	fields := graphql.CollectFields(ec.doc, sel, __TypeImplementors, ec.variables)
	out := graphql.NewOrderedMap(len(fields))
	for i, field := range fields {
		out.Keys[i] = field.Alias

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(field, obj)
		case "name":
			out.Values[i] = ec.___Type_name(field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}

	return out
}

func (ec *executionContext) ___Type_kind(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Kind()
	return graphql.MarshalString(res)
}

func (ec *executionContext) ___Type_name(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Name()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Type_description(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Description()
	if res == nil {
		return graphql.Null
	}
	return graphql.MarshalString(*res)
}

func (ec *executionContext) ___Type_fields(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	var arg0 bool
	if tmp, ok := field.Args["includeDeprecated"]; ok {
		var err error

		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res := obj.Fields(arg0)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Field(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_interfaces(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.Interfaces()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_possibleTypes(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.PossibleTypes()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___Type(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_enumValues(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	var arg0 bool
	if tmp, ok := field.Args["includeDeprecated"]; ok {
		var err error

		arg0, err = graphql.UnmarshalBoolean(tmp)
		if err != nil {
			ec.Error(err)
			return graphql.Null
		}
	}
	res := obj.EnumValues(arg0)
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___EnumValue(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_inputFields(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.InputFields()
	arr1 := graphql.Array{}
	for idx1 := range res {
		arr1 = append(arr1, func() graphql.Marshaler {
			if res[idx1] == nil {
				return graphql.Null
			}
			return ec.___InputValue(field.Selections, res[idx1])
		}())
	}
	return arr1
}

func (ec *executionContext) ___Type_ofType(field graphql.CollectedField, obj *introspection.Type) graphql.Marshaler {
	res := obj.OfType()
	if res == nil {
		return graphql.Null
	}
	return ec.___Type(field.Selections, res)
}

var parsedSchema = schema.MustParse("schema {\n  query: Query\n  mutation: Mutation\n}\n\nenum TodoStatus {\n  SHOW_ALL\n  SHOW_COMPLETED\n  SHOW_ACTIVE\n}\n\ntype Query {\n  todo(id: Int!): Todo\n  todos(status: TodoStatus): [Todo!]!\n}\n\ntype Mutation {\n  createTodo(text: String!, done: Boolean): Todo!\n  updateTodos(ids: [Int!], changes: Map!): [Todo]\n  deleteTodos(ids: [Int!]): [Todo]\n}\n\ntype Todo {\n  id: Int!\n  text: String!\n  done: Boolean!\n}\n")

func (ec *executionContext) introspectSchema() *introspection.Schema {
	return introspection.WrapSchema(parsedSchema)
}

func (ec *executionContext) introspectType(name string) *introspection.Type {
	t := parsedSchema.Resolve(name)
	if t == nil {
		return nil
	}
	return introspection.WrapType(t)
}