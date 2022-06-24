/*
Zinc Search engine API

Zinc Search engine API documents https://docs.zincsearch.com

API version: 0.2.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// MetaQuery struct for MetaQuery
type MetaQuery struct {
	Bool *MetaBoolQuery `json:"bool,omitempty"`
	Exists *MetaExistsQuery `json:"exists,omitempty"`
	// simple, PrefixQuery
	Fuzzy *map[string]MetaFuzzyQuery `json:"fuzzy,omitempty"`
	Ids *MetaIdsQuery `json:"ids,omitempty"`
	// simple, MatchQuery
	Match *map[string]MetaMatchQuery `json:"match,omitempty"`
	MatchAll map[string]interface{} `json:"match_all,omitempty"`
	// simple, MatchBoolPrefixQuery
	MatchBoolPrefix *map[string]MetaMatchBoolPrefixQuery `json:"match_bool_prefix,omitempty"`
	MatchNone map[string]interface{} `json:"match_none,omitempty"`
	// simple, MatchPhraseQuery
	MatchPhrase *map[string]MetaMatchPhraseQuery `json:"match_phrase,omitempty"`
	// simple, MatchPhrasePrefixQuery
	MatchPhrasePrefix *map[string]MetaMatchPhrasePrefixQuery `json:"match_phrase_prefix,omitempty"`
	MultiMatch *MetaMultiMatchQuery `json:"multi_match,omitempty"`
	// .
	Prefix *map[string]MetaPrefixQuery `json:"prefix,omitempty"`
	QueryString *MetaQueryStringQuery `json:"query_string,omitempty"`
	// simple, FuzzyQuery
	Range *map[string]MetaRangeQuery `json:"range,omitempty"`
	// simple, FuzzyQuery
	Regexp *map[string]MetaRegexpQuery `json:"regexp,omitempty"`
	SimpleQueryString *MetaSimpleQueryStringQuery `json:"simple_query_string,omitempty"`
	// simple, TermQuery
	Term *map[string]MetaTermQuery `json:"term,omitempty"`
	// .
	Terms *map[string]map[string]interface{} `json:"terms,omitempty"`
	// simple, WildcardQuery
	Wildcard *map[string]MetaWildcardQuery `json:"wildcard,omitempty"`
}

// NewMetaQuery instantiates a new MetaQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaQuery() *MetaQuery {
	this := MetaQuery{}
	return &this
}

// NewMetaQueryWithDefaults instantiates a new MetaQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaQueryWithDefaults() *MetaQuery {
	this := MetaQuery{}
	return &this
}

// GetBool returns the Bool field value if set, zero value otherwise.
func (o *MetaQuery) GetBool() MetaBoolQuery {
	if o == nil || o.Bool == nil {
		var ret MetaBoolQuery
		return ret
	}
	return *o.Bool
}

// GetBoolOk returns a tuple with the Bool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetBoolOk() (*MetaBoolQuery, bool) {
	if o == nil || o.Bool == nil {
		return nil, false
	}
	return o.Bool, true
}

// HasBool returns a boolean if a field has been set.
func (o *MetaQuery) HasBool() bool {
	if o != nil && o.Bool != nil {
		return true
	}

	return false
}

// SetBool gets a reference to the given MetaBoolQuery and assigns it to the Bool field.
func (o *MetaQuery) SetBool(v MetaBoolQuery) {
	o.Bool = &v
}

// GetExists returns the Exists field value if set, zero value otherwise.
func (o *MetaQuery) GetExists() MetaExistsQuery {
	if o == nil || o.Exists == nil {
		var ret MetaExistsQuery
		return ret
	}
	return *o.Exists
}

// GetExistsOk returns a tuple with the Exists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetExistsOk() (*MetaExistsQuery, bool) {
	if o == nil || o.Exists == nil {
		return nil, false
	}
	return o.Exists, true
}

// HasExists returns a boolean if a field has been set.
func (o *MetaQuery) HasExists() bool {
	if o != nil && o.Exists != nil {
		return true
	}

	return false
}

// SetExists gets a reference to the given MetaExistsQuery and assigns it to the Exists field.
func (o *MetaQuery) SetExists(v MetaExistsQuery) {
	o.Exists = &v
}

// GetFuzzy returns the Fuzzy field value if set, zero value otherwise.
func (o *MetaQuery) GetFuzzy() map[string]MetaFuzzyQuery {
	if o == nil || o.Fuzzy == nil {
		var ret map[string]MetaFuzzyQuery
		return ret
	}
	return *o.Fuzzy
}

// GetFuzzyOk returns a tuple with the Fuzzy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetFuzzyOk() (*map[string]MetaFuzzyQuery, bool) {
	if o == nil || o.Fuzzy == nil {
		return nil, false
	}
	return o.Fuzzy, true
}

// HasFuzzy returns a boolean if a field has been set.
func (o *MetaQuery) HasFuzzy() bool {
	if o != nil && o.Fuzzy != nil {
		return true
	}

	return false
}

// SetFuzzy gets a reference to the given map[string]MetaFuzzyQuery and assigns it to the Fuzzy field.
func (o *MetaQuery) SetFuzzy(v map[string]MetaFuzzyQuery) {
	o.Fuzzy = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *MetaQuery) GetIds() MetaIdsQuery {
	if o == nil || o.Ids == nil {
		var ret MetaIdsQuery
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetIdsOk() (*MetaIdsQuery, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *MetaQuery) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given MetaIdsQuery and assigns it to the Ids field.
func (o *MetaQuery) SetIds(v MetaIdsQuery) {
	o.Ids = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *MetaQuery) GetMatch() map[string]MetaMatchQuery {
	if o == nil || o.Match == nil {
		var ret map[string]MetaMatchQuery
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetMatchOk() (*map[string]MetaMatchQuery, bool) {
	if o == nil || o.Match == nil {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *MetaQuery) HasMatch() bool {
	if o != nil && o.Match != nil {
		return true
	}

	return false
}

// SetMatch gets a reference to the given map[string]MetaMatchQuery and assigns it to the Match field.
func (o *MetaQuery) SetMatch(v map[string]MetaMatchQuery) {
	o.Match = &v
}

// GetMatchAll returns the MatchAll field value if set, zero value otherwise.
func (o *MetaQuery) GetMatchAll() map[string]interface{} {
	if o == nil || o.MatchAll == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MatchAll
}

// GetMatchAllOk returns a tuple with the MatchAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetMatchAllOk() (map[string]interface{}, bool) {
	if o == nil || o.MatchAll == nil {
		return nil, false
	}
	return o.MatchAll, true
}

// HasMatchAll returns a boolean if a field has been set.
func (o *MetaQuery) HasMatchAll() bool {
	if o != nil && o.MatchAll != nil {
		return true
	}

	return false
}

// SetMatchAll gets a reference to the given map[string]interface{} and assigns it to the MatchAll field.
func (o *MetaQuery) SetMatchAll(v map[string]interface{}) {
	o.MatchAll = v
}

// GetMatchBoolPrefix returns the MatchBoolPrefix field value if set, zero value otherwise.
func (o *MetaQuery) GetMatchBoolPrefix() map[string]MetaMatchBoolPrefixQuery {
	if o == nil || o.MatchBoolPrefix == nil {
		var ret map[string]MetaMatchBoolPrefixQuery
		return ret
	}
	return *o.MatchBoolPrefix
}

// GetMatchBoolPrefixOk returns a tuple with the MatchBoolPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetMatchBoolPrefixOk() (*map[string]MetaMatchBoolPrefixQuery, bool) {
	if o == nil || o.MatchBoolPrefix == nil {
		return nil, false
	}
	return o.MatchBoolPrefix, true
}

// HasMatchBoolPrefix returns a boolean if a field has been set.
func (o *MetaQuery) HasMatchBoolPrefix() bool {
	if o != nil && o.MatchBoolPrefix != nil {
		return true
	}

	return false
}

// SetMatchBoolPrefix gets a reference to the given map[string]MetaMatchBoolPrefixQuery and assigns it to the MatchBoolPrefix field.
func (o *MetaQuery) SetMatchBoolPrefix(v map[string]MetaMatchBoolPrefixQuery) {
	o.MatchBoolPrefix = &v
}

// GetMatchNone returns the MatchNone field value if set, zero value otherwise.
func (o *MetaQuery) GetMatchNone() map[string]interface{} {
	if o == nil || o.MatchNone == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MatchNone
}

// GetMatchNoneOk returns a tuple with the MatchNone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetMatchNoneOk() (map[string]interface{}, bool) {
	if o == nil || o.MatchNone == nil {
		return nil, false
	}
	return o.MatchNone, true
}

// HasMatchNone returns a boolean if a field has been set.
func (o *MetaQuery) HasMatchNone() bool {
	if o != nil && o.MatchNone != nil {
		return true
	}

	return false
}

// SetMatchNone gets a reference to the given map[string]interface{} and assigns it to the MatchNone field.
func (o *MetaQuery) SetMatchNone(v map[string]interface{}) {
	o.MatchNone = v
}

// GetMatchPhrase returns the MatchPhrase field value if set, zero value otherwise.
func (o *MetaQuery) GetMatchPhrase() map[string]MetaMatchPhraseQuery {
	if o == nil || o.MatchPhrase == nil {
		var ret map[string]MetaMatchPhraseQuery
		return ret
	}
	return *o.MatchPhrase
}

// GetMatchPhraseOk returns a tuple with the MatchPhrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetMatchPhraseOk() (*map[string]MetaMatchPhraseQuery, bool) {
	if o == nil || o.MatchPhrase == nil {
		return nil, false
	}
	return o.MatchPhrase, true
}

// HasMatchPhrase returns a boolean if a field has been set.
func (o *MetaQuery) HasMatchPhrase() bool {
	if o != nil && o.MatchPhrase != nil {
		return true
	}

	return false
}

// SetMatchPhrase gets a reference to the given map[string]MetaMatchPhraseQuery and assigns it to the MatchPhrase field.
func (o *MetaQuery) SetMatchPhrase(v map[string]MetaMatchPhraseQuery) {
	o.MatchPhrase = &v
}

// GetMatchPhrasePrefix returns the MatchPhrasePrefix field value if set, zero value otherwise.
func (o *MetaQuery) GetMatchPhrasePrefix() map[string]MetaMatchPhrasePrefixQuery {
	if o == nil || o.MatchPhrasePrefix == nil {
		var ret map[string]MetaMatchPhrasePrefixQuery
		return ret
	}
	return *o.MatchPhrasePrefix
}

// GetMatchPhrasePrefixOk returns a tuple with the MatchPhrasePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetMatchPhrasePrefixOk() (*map[string]MetaMatchPhrasePrefixQuery, bool) {
	if o == nil || o.MatchPhrasePrefix == nil {
		return nil, false
	}
	return o.MatchPhrasePrefix, true
}

// HasMatchPhrasePrefix returns a boolean if a field has been set.
func (o *MetaQuery) HasMatchPhrasePrefix() bool {
	if o != nil && o.MatchPhrasePrefix != nil {
		return true
	}

	return false
}

// SetMatchPhrasePrefix gets a reference to the given map[string]MetaMatchPhrasePrefixQuery and assigns it to the MatchPhrasePrefix field.
func (o *MetaQuery) SetMatchPhrasePrefix(v map[string]MetaMatchPhrasePrefixQuery) {
	o.MatchPhrasePrefix = &v
}

// GetMultiMatch returns the MultiMatch field value if set, zero value otherwise.
func (o *MetaQuery) GetMultiMatch() MetaMultiMatchQuery {
	if o == nil || o.MultiMatch == nil {
		var ret MetaMultiMatchQuery
		return ret
	}
	return *o.MultiMatch
}

// GetMultiMatchOk returns a tuple with the MultiMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetMultiMatchOk() (*MetaMultiMatchQuery, bool) {
	if o == nil || o.MultiMatch == nil {
		return nil, false
	}
	return o.MultiMatch, true
}

// HasMultiMatch returns a boolean if a field has been set.
func (o *MetaQuery) HasMultiMatch() bool {
	if o != nil && o.MultiMatch != nil {
		return true
	}

	return false
}

// SetMultiMatch gets a reference to the given MetaMultiMatchQuery and assigns it to the MultiMatch field.
func (o *MetaQuery) SetMultiMatch(v MetaMultiMatchQuery) {
	o.MultiMatch = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *MetaQuery) GetPrefix() map[string]MetaPrefixQuery {
	if o == nil || o.Prefix == nil {
		var ret map[string]MetaPrefixQuery
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetPrefixOk() (*map[string]MetaPrefixQuery, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *MetaQuery) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given map[string]MetaPrefixQuery and assigns it to the Prefix field.
func (o *MetaQuery) SetPrefix(v map[string]MetaPrefixQuery) {
	o.Prefix = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *MetaQuery) GetQueryString() MetaQueryStringQuery {
	if o == nil || o.QueryString == nil {
		var ret MetaQueryStringQuery
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetQueryStringOk() (*MetaQueryStringQuery, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *MetaQuery) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given MetaQueryStringQuery and assigns it to the QueryString field.
func (o *MetaQuery) SetQueryString(v MetaQueryStringQuery) {
	o.QueryString = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *MetaQuery) GetRange() map[string]MetaRangeQuery {
	if o == nil || o.Range == nil {
		var ret map[string]MetaRangeQuery
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetRangeOk() (*map[string]MetaRangeQuery, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *MetaQuery) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given map[string]MetaRangeQuery and assigns it to the Range field.
func (o *MetaQuery) SetRange(v map[string]MetaRangeQuery) {
	o.Range = &v
}

// GetRegexp returns the Regexp field value if set, zero value otherwise.
func (o *MetaQuery) GetRegexp() map[string]MetaRegexpQuery {
	if o == nil || o.Regexp == nil {
		var ret map[string]MetaRegexpQuery
		return ret
	}
	return *o.Regexp
}

// GetRegexpOk returns a tuple with the Regexp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetRegexpOk() (*map[string]MetaRegexpQuery, bool) {
	if o == nil || o.Regexp == nil {
		return nil, false
	}
	return o.Regexp, true
}

// HasRegexp returns a boolean if a field has been set.
func (o *MetaQuery) HasRegexp() bool {
	if o != nil && o.Regexp != nil {
		return true
	}

	return false
}

// SetRegexp gets a reference to the given map[string]MetaRegexpQuery and assigns it to the Regexp field.
func (o *MetaQuery) SetRegexp(v map[string]MetaRegexpQuery) {
	o.Regexp = &v
}

// GetSimpleQueryString returns the SimpleQueryString field value if set, zero value otherwise.
func (o *MetaQuery) GetSimpleQueryString() MetaSimpleQueryStringQuery {
	if o == nil || o.SimpleQueryString == nil {
		var ret MetaSimpleQueryStringQuery
		return ret
	}
	return *o.SimpleQueryString
}

// GetSimpleQueryStringOk returns a tuple with the SimpleQueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetSimpleQueryStringOk() (*MetaSimpleQueryStringQuery, bool) {
	if o == nil || o.SimpleQueryString == nil {
		return nil, false
	}
	return o.SimpleQueryString, true
}

// HasSimpleQueryString returns a boolean if a field has been set.
func (o *MetaQuery) HasSimpleQueryString() bool {
	if o != nil && o.SimpleQueryString != nil {
		return true
	}

	return false
}

// SetSimpleQueryString gets a reference to the given MetaSimpleQueryStringQuery and assigns it to the SimpleQueryString field.
func (o *MetaQuery) SetSimpleQueryString(v MetaSimpleQueryStringQuery) {
	o.SimpleQueryString = &v
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *MetaQuery) GetTerm() map[string]MetaTermQuery {
	if o == nil || o.Term == nil {
		var ret map[string]MetaTermQuery
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetTermOk() (*map[string]MetaTermQuery, bool) {
	if o == nil || o.Term == nil {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *MetaQuery) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given map[string]MetaTermQuery and assigns it to the Term field.
func (o *MetaQuery) SetTerm(v map[string]MetaTermQuery) {
	o.Term = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *MetaQuery) GetTerms() map[string]map[string]interface{} {
	if o == nil || o.Terms == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetTermsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Terms == nil {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *MetaQuery) HasTerms() bool {
	if o != nil && o.Terms != nil {
		return true
	}

	return false
}

// SetTerms gets a reference to the given map[string]map[string]interface{} and assigns it to the Terms field.
func (o *MetaQuery) SetTerms(v map[string]map[string]interface{}) {
	o.Terms = &v
}

// GetWildcard returns the Wildcard field value if set, zero value otherwise.
func (o *MetaQuery) GetWildcard() map[string]MetaWildcardQuery {
	if o == nil || o.Wildcard == nil {
		var ret map[string]MetaWildcardQuery
		return ret
	}
	return *o.Wildcard
}

// GetWildcardOk returns a tuple with the Wildcard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaQuery) GetWildcardOk() (*map[string]MetaWildcardQuery, bool) {
	if o == nil || o.Wildcard == nil {
		return nil, false
	}
	return o.Wildcard, true
}

// HasWildcard returns a boolean if a field has been set.
func (o *MetaQuery) HasWildcard() bool {
	if o != nil && o.Wildcard != nil {
		return true
	}

	return false
}

// SetWildcard gets a reference to the given map[string]MetaWildcardQuery and assigns it to the Wildcard field.
func (o *MetaQuery) SetWildcard(v map[string]MetaWildcardQuery) {
	o.Wildcard = &v
}

func (o MetaQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bool != nil {
		toSerialize["bool"] = o.Bool
	}
	if o.Exists != nil {
		toSerialize["exists"] = o.Exists
	}
	if o.Fuzzy != nil {
		toSerialize["fuzzy"] = o.Fuzzy
	}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	if o.Match != nil {
		toSerialize["match"] = o.Match
	}
	if o.MatchAll != nil {
		toSerialize["match_all"] = o.MatchAll
	}
	if o.MatchBoolPrefix != nil {
		toSerialize["match_bool_prefix"] = o.MatchBoolPrefix
	}
	if o.MatchNone != nil {
		toSerialize["match_none"] = o.MatchNone
	}
	if o.MatchPhrase != nil {
		toSerialize["match_phrase"] = o.MatchPhrase
	}
	if o.MatchPhrasePrefix != nil {
		toSerialize["match_phrase_prefix"] = o.MatchPhrasePrefix
	}
	if o.MultiMatch != nil {
		toSerialize["multi_match"] = o.MultiMatch
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.QueryString != nil {
		toSerialize["query_string"] = o.QueryString
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.Regexp != nil {
		toSerialize["regexp"] = o.Regexp
	}
	if o.SimpleQueryString != nil {
		toSerialize["simple_query_string"] = o.SimpleQueryString
	}
	if o.Term != nil {
		toSerialize["term"] = o.Term
	}
	if o.Terms != nil {
		toSerialize["terms"] = o.Terms
	}
	if o.Wildcard != nil {
		toSerialize["wildcard"] = o.Wildcard
	}
	return json.Marshal(toSerialize)
}

type NullableMetaQuery struct {
	value *MetaQuery
	isSet bool
}

func (v NullableMetaQuery) Get() *MetaQuery {
	return v.value
}

func (v *NullableMetaQuery) Set(val *MetaQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaQuery(val *MetaQuery) *NullableMetaQuery {
	return &NullableMetaQuery{value: val, isSet: true}
}

func (v NullableMetaQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

