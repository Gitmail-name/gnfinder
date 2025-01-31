package verifier

import "github.com/machinebox/graphql"

type graphqlResponse struct {
	NameResolver struct {
		Responses []response
	}
}

type response struct {
	MatchedDataSources int
	SuppliedInput      string
	QualitySummary     string
	Results            []dataResult
	PreferredResults   []dataResult
}

type dataResult struct {
	Classification classification
	DataSource     dataSource
	TaxonID        string
	Name           name
	CanonicalName  canonical
	AcceptedName   acceptedName
	Synonym        bool
	MatchType      matchType
}

type dataSource struct {
	ID    int
	Title string
}

type name struct {
	ID    string
	Value string
}

type canonical struct {
	ValueRanked string
}

type classification struct {
	Path      string
	PathRanks string
	PathIDs   string
}

type acceptedName struct {
	Name name
}

type matchType struct {
	Kind                 string
	VerbatimEditDistance int
	StemEditDistance     int
}

func graphqlRequest() *graphql.Request {
	req := graphql.NewRequest(`
query($names: [name!]!, $sources: [Int!]) {
  nameResolver(names: $names,
		preferredDataSourceIds: $sources,
		advancedResolution: true
    bestMatchOnly: true) {
    responses {
      total
      suppliedInput
      qualitySummary
      matchedDataSources
      results {
				name { id value }
				canonicalName { valueRanked }
        taxonId
				classification { path pathRanks pathIds }
        dataSource { id title }
        acceptedName { name { value } }
        synonym
        matchType {
					kind
					verbatimEditDistance
					stemEditDistance
				}
      }
      preferredResults {
				name { id value }
				canonicalName { valueRanked }
        taxonId
				classification { path pathRanks pathIds }
        dataSource { id title }
        acceptedName { name { value } }
        synonym
        matchType {
					kind
					verbatimEditDistance
					stemEditDistance
				}
      }
    }
  }
}`)
	return req
}
