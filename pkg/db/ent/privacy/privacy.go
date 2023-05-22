// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/chain-middleware/pkg/db/ent"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns an formatted wrapped Allow decision.
func Allowf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Allow)...)
}

// Denyf returns an formatted wrapped Deny decision.
func Denyf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Deny)...)
}

// Skipf returns an formatted wrapped Skip decision.
func Skipf(format string, a ...interface{}) error {
	return fmt.Errorf(format+": %w", append(a, Skip)...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// MutationRuleFunc type is an adapter which allows the use of
// ordinary functions as mutation rules.
type MutationRuleFunc func(context.Context, ent.Mutation) error

// EvalMutation returns f(ctx, m).
func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	return f(ctx, m)
}

// QueryMutationRule is an interface which groups query and mutation rules.
type QueryMutationRule interface {
	QueryRule
	MutationRule
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return fixedDecision{Allow}
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return fixedDecision{Deny}
}

type fixedDecision struct {
	decision error
}

func (f fixedDecision) EvalQuery(context.Context, ent.Query) error {
	return f.decision
}

func (f fixedDecision) EvalMutation(context.Context, ent.Mutation) error {
	return f.decision
}

type contextDecision struct {
	eval func(context.Context) error
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return contextDecision{eval}
}

func (c contextDecision) EvalQuery(ctx context.Context, _ ent.Query) error {
	return c.eval(ctx)
}

func (c contextDecision) EvalMutation(ctx context.Context, _ ent.Mutation) error {
	return c.eval(ctx)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		if m.Op().Is(op) {
			return rule.EvalMutation(ctx, m)
		}
		return Skip
	})
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AppCoinQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AppCoinQueryRuleFunc func(context.Context, *ent.AppCoinQuery) error

// EvalQuery return f(ctx, q).
func (f AppCoinQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppCoinQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AppCoinQuery", q)
}

// The AppCoinMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AppCoinMutationRuleFunc func(context.Context, *ent.AppCoinMutation) error

// EvalMutation calls f(ctx, m).
func (f AppCoinMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AppCoinMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AppCoinMutation", m)
}

// The CoinBaseQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CoinBaseQueryRuleFunc func(context.Context, *ent.CoinBaseQuery) error

// EvalQuery return f(ctx, q).
func (f CoinBaseQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CoinBaseQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CoinBaseQuery", q)
}

// The CoinBaseMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CoinBaseMutationRuleFunc func(context.Context, *ent.CoinBaseMutation) error

// EvalMutation calls f(ctx, m).
func (f CoinBaseMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CoinBaseMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CoinBaseMutation", m)
}

// The CoinDescriptionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CoinDescriptionQueryRuleFunc func(context.Context, *ent.CoinDescriptionQuery) error

// EvalQuery return f(ctx, q).
func (f CoinDescriptionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CoinDescriptionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CoinDescriptionQuery", q)
}

// The CoinDescriptionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CoinDescriptionMutationRuleFunc func(context.Context, *ent.CoinDescriptionMutation) error

// EvalMutation calls f(ctx, m).
func (f CoinDescriptionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CoinDescriptionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CoinDescriptionMutation", m)
}

// The CoinExtraQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CoinExtraQueryRuleFunc func(context.Context, *ent.CoinExtraQuery) error

// EvalQuery return f(ctx, q).
func (f CoinExtraQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CoinExtraQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CoinExtraQuery", q)
}

// The CoinExtraMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CoinExtraMutationRuleFunc func(context.Context, *ent.CoinExtraMutation) error

// EvalMutation calls f(ctx, m).
func (f CoinExtraMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CoinExtraMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CoinExtraMutation", m)
}

// The CurrencyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type CurrencyQueryRuleFunc func(context.Context, *ent.CurrencyQuery) error

// EvalQuery return f(ctx, q).
func (f CurrencyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CurrencyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.CurrencyQuery", q)
}

// The CurrencyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type CurrencyMutationRuleFunc func(context.Context, *ent.CurrencyMutation) error

// EvalMutation calls f(ctx, m).
func (f CurrencyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.CurrencyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.CurrencyMutation", m)
}

// The ExchangeRateQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ExchangeRateQueryRuleFunc func(context.Context, *ent.ExchangeRateQuery) error

// EvalQuery return f(ctx, q).
func (f ExchangeRateQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ExchangeRateQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ExchangeRateQuery", q)
}

// The ExchangeRateMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ExchangeRateMutationRuleFunc func(context.Context, *ent.ExchangeRateMutation) error

// EvalMutation calls f(ctx, m).
func (f ExchangeRateMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ExchangeRateMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ExchangeRateMutation", m)
}

// The FiatCurrencyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FiatCurrencyQueryRuleFunc func(context.Context, *ent.FiatCurrencyQuery) error

// EvalQuery return f(ctx, q).
func (f FiatCurrencyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FiatCurrencyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FiatCurrencyQuery", q)
}

// The FiatCurrencyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FiatCurrencyMutationRuleFunc func(context.Context, *ent.FiatCurrencyMutation) error

// EvalMutation calls f(ctx, m).
func (f FiatCurrencyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FiatCurrencyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FiatCurrencyMutation", m)
}

// The FiatCurrencyTypeQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type FiatCurrencyTypeQueryRuleFunc func(context.Context, *ent.FiatCurrencyTypeQuery) error

// EvalQuery return f(ctx, q).
func (f FiatCurrencyTypeQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FiatCurrencyTypeQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.FiatCurrencyTypeQuery", q)
}

// The FiatCurrencyTypeMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type FiatCurrencyTypeMutationRuleFunc func(context.Context, *ent.FiatCurrencyTypeMutation) error

// EvalMutation calls f(ctx, m).
func (f FiatCurrencyTypeMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.FiatCurrencyTypeMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.FiatCurrencyTypeMutation", m)
}

// The SettingQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SettingQueryRuleFunc func(context.Context, *ent.SettingQuery) error

// EvalQuery return f(ctx, q).
func (f SettingQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SettingQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.SettingQuery", q)
}

// The SettingMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SettingMutationRuleFunc func(context.Context, *ent.SettingMutation) error

// EvalMutation calls f(ctx, m).
func (f SettingMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.SettingMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.SettingMutation", m)
}

// The TranQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type TranQueryRuleFunc func(context.Context, *ent.TranQuery) error

// EvalQuery return f(ctx, q).
func (f TranQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TranQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.TranQuery", q)
}

// The TranMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type TranMutationRuleFunc func(context.Context, *ent.TranMutation) error

// EvalMutation calls f(ctx, m).
func (f TranMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.TranMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.TranMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q ent.Query) (Filter, error) {
	switch q := q.(type) {
	case *ent.AppCoinQuery:
		return q.Filter(), nil
	case *ent.CoinBaseQuery:
		return q.Filter(), nil
	case *ent.CoinDescriptionQuery:
		return q.Filter(), nil
	case *ent.CoinExtraQuery:
		return q.Filter(), nil
	case *ent.CurrencyQuery:
		return q.Filter(), nil
	case *ent.ExchangeRateQuery:
		return q.Filter(), nil
	case *ent.FiatCurrencyQuery:
		return q.Filter(), nil
	case *ent.FiatCurrencyTypeQuery:
		return q.Filter(), nil
	case *ent.SettingQuery:
		return q.Filter(), nil
	case *ent.TranQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m ent.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *ent.AppCoinMutation:
		return m.Filter(), nil
	case *ent.CoinBaseMutation:
		return m.Filter(), nil
	case *ent.CoinDescriptionMutation:
		return m.Filter(), nil
	case *ent.CoinExtraMutation:
		return m.Filter(), nil
	case *ent.CurrencyMutation:
		return m.Filter(), nil
	case *ent.ExchangeRateMutation:
		return m.Filter(), nil
	case *ent.FiatCurrencyMutation:
		return m.Filter(), nil
	case *ent.FiatCurrencyTypeMutation:
		return m.Filter(), nil
	case *ent.SettingMutation:
		return m.Filter(), nil
	case *ent.TranMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("ent/privacy: unexpected mutation type %T for mutation filter", m)
	}
}
