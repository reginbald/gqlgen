// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type QueryResolver interface {
	DirectiveSingleNullableArg(ctx context.Context, arg1 *string) (*string, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) dir_populate_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.dir_populate_argsValue(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["value"] = arg0
	return args, nil
}
func (ec *executionContext) dir_populate_argsValue(
	ctx context.Context,
	rawArgs map[string]any,
) (string, error) {
	if _, ok := rawArgs["value"]; !ok {
		var zeroVal string
		return zeroVal, nil
	}

	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("value"))
	if tmp, ok := rawArgs["value"]; ok {
		return ec.unmarshalNString2string(ctx, tmp)
	}

	var zeroVal string
	return zeroVal, nil
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Query___type_argsName(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["name"] = arg0
	return args, nil
}
func (ec *executionContext) field_Query___type_argsName(
	ctx context.Context,
	rawArgs map[string]any,
) (string, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
	if tmp, ok := rawArgs["name"]; ok {
		return ec.unmarshalNString2string(ctx, tmp)
	}

	var zeroVal string
	return zeroVal, nil
}

func (ec *executionContext) field_Query_directiveSingleNullableArg_args(ctx context.Context, rawArgs map[string]any) (map[string]any, error) {
	var err error
	args := map[string]any{}
	arg0, err := ec.field_Query_directiveSingleNullableArg_argsArg1(ctx, rawArgs)
	if err != nil {
		return nil, err
	}
	args["arg1"] = arg0
	return args, nil
}
func (ec *executionContext) field_Query_directiveSingleNullableArg_argsArg1(
	ctx context.Context,
	rawArgs map[string]any,
) (*string, error) {
	ctx = graphql.WithPathContext(ctx, graphql.NewPathWithField("arg1"))
	directive0 := func(ctx context.Context) (any, error) {
		tmp, ok := rawArgs["arg1"]
		if !ok {
			var zeroVal *string
			return zeroVal, nil
		}
		return ec.unmarshalOString2ᚖstring(ctx, tmp)
	}

	directive1 := func(ctx context.Context) (any, error) {
		value, err := ec.unmarshalNString2string(ctx, "test")
		if err != nil {
			var zeroVal *string
			return zeroVal, err
		}
		if ec.directives.Populate == nil {
			var zeroVal *string
			return zeroVal, errors.New("directive populate is not implemented")
		}
		return ec.directives.Populate(ctx, rawArgs, directive0, value)
	}
	directive2 := func(ctx context.Context) (any, error) {
		if ec.directives.Noop == nil {
			var zeroVal *string
			return zeroVal, errors.New("directive noop is not implemented")
		}
		return ec.directives.Noop(ctx, rawArgs, directive1)
	}

	tmp, err := directive2(ctx)
	if err != nil {
		var zeroVal *string
		return zeroVal, graphql.ErrorOnPath(ctx, err)
	}
	if data, ok := tmp.(*string); ok {
		return data, nil
	} else if tmp == nil {
		var zeroVal *string
		return zeroVal, nil
	} else {
		var zeroVal *string
		return zeroVal, graphql.ErrorOnPath(ctx, fmt.Errorf(`unexpected type %T from directive, should be *string`, tmp))
	}
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Query_directiveSingleNullableArg(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query_directiveSingleNullableArg(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Query().DirectiveSingleNullableArg(rctx, fc.Args["arg1"].(*string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalOString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query_directiveSingleNullableArg(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query_directiveSingleNullableArg_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___type(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectType(fc.Args["name"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Type)
	fc.Result = res
	return ec.marshalO__Type2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐType(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___type(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "kind":
				return ec.fieldContext___Type_kind(ctx, field)
			case "name":
				return ec.fieldContext___Type_name(ctx, field)
			case "description":
				return ec.fieldContext___Type_description(ctx, field)
			case "fields":
				return ec.fieldContext___Type_fields(ctx, field)
			case "interfaces":
				return ec.fieldContext___Type_interfaces(ctx, field)
			case "possibleTypes":
				return ec.fieldContext___Type_possibleTypes(ctx, field)
			case "enumValues":
				return ec.fieldContext___Type_enumValues(ctx, field)
			case "inputFields":
				return ec.fieldContext___Type_inputFields(ctx, field)
			case "ofType":
				return ec.fieldContext___Type_ofType(ctx, field)
			case "specifiedByURL":
				return ec.fieldContext___Type_specifiedByURL(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Type", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Query___type_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Query___schema(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (any, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.introspectSchema()
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*introspection.Schema)
	fc.Result = res
	return ec.marshalO__Schema2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋgraphqlᚋintrospectionᚐSchema(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Query___schema(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Query",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "description":
				return ec.fieldContext___Schema_description(ctx, field)
			case "types":
				return ec.fieldContext___Schema_types(ctx, field)
			case "queryType":
				return ec.fieldContext___Schema_queryType(ctx, field)
			case "mutationType":
				return ec.fieldContext___Schema_mutationType(ctx, field)
			case "subscriptionType":
				return ec.fieldContext___Schema_subscriptionType(ctx, field)
			case "directives":
				return ec.fieldContext___Schema_directives(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type __Schema", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var queryImplementors = []string{"Query"}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Query",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Query")
		case "directiveSingleNullableArg":
			field := field

			innerFunc := func(ctx context.Context, _ *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Query_directiveSingleNullableArg(ctx, field)
				return res
			}

			rrm := func(ctx context.Context) graphql.Marshaler {
				return ec.OperationContext.RootResolverMiddleware(ctx,
					func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return rrm(innerCtx) })
		case "__type":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___type(ctx, field)
			})
		case "__schema":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Query___schema(ctx, field)
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

// endregion ***************************** type.gotpl *****************************
