package es

func test()  {
	/**
	docker pull blacktop/kibana:7.4
	docker tag blacktop/kibana:7.4  kb:74
	docker run --init -d --name kb -e elasticsearch.hosts="http://{service}:9200"  -p 5601:5601  kb:74
	 */
	// date group  - ad_report_type group
	/*
		GET /datacenter/_search
		{
			"size": 0,
			"aggs": {
				"group_by_date_num": {
					"terms": {
						"field": "date_num",
						"order": {
							"_key": "asc"
						}
					},
					"aggs": {
						"group_by_report_type": {
							"terms": {
								"field": "ad_report_type"
							},
							"aggs": {
								"count": {
									"sum": {
										"field": "ad_report_count"
									}
								}
							}
						}

					}
				}
			}
		}
	*/
	/*query := elastic.NewBoolQuery()
	field_ad_report_type, field_ad_report_count, field_date_num := "ad_report_type", "ad_report_count", "date_num"
	group_ad_report_type, group_date_num := "group_ad_report_type", "group_by_date_num"
	sum_ad_report_count := "sum_ad_report_count"
	if len(cpx) > 0 {
		switch tag {
		case ">=":
			query.Filter(elastic.NewRangeQuery(cpx).Gte(value))
		case "<":
			query.Filter(elastic.NewRangeQuery(cpx).Lt(value))
		}
	}
	if startDate.Unix() > 0 {
		query = query.Filter(elastic.NewRangeQuery(field_date_num).Gte(startDate.Format("20060102")))
	}
	if endDate.Unix() > 0 {
		query = query.Filter(elastic.NewRangeQuery(field_date_num).Lte(endDate.Format("20060102")))
	}
	sumCpx := elastic.NewTermsAggregation().Field(field_date_num).Order("_key", true).
		SubAggregation(group_ad_report_type, elastic.NewTermsAggregation().Field(field_ad_report_type).
			SubAggregation(sum_ad_report_count, elastic.NewSumAggregation().Field(field_ad_report_count)))
	searchResult, _ := EsClient().Search().
		Query(query).
		Aggregation(group_date_num, sumCpx).
		Pretty(true).
		Do(context.Background())
	groupDate, _ := searchResult.Aggregations.Terms(group_date_num)
	var resp []domain.AdListEffect
	for _, bucket := range groupDate.Buckets {
		groupCpx, status := bucket.Aggregations.Terms(group_ad_report_type)
		effect := domain.AdListEffect{Date: uint64(bucket.Key.(float64))}
		for _, cpxBucket := range groupCpx.Buckets {
			if status {
				cpx, _ := cpxBucket.ValueCount(sum_ad_report_count)
				switch cpxBucket.Key {
				case "cpa":
					effect.Cpa = uint64(*cpx.Value)
				case "cpc":
					effect.Cpc = uint64(*cpx.Value)
				case "cpm":
					effect.Cpm = uint64(*cpx.Value)
				}
			}
		}
		resp = append(resp, effect)
	}
	return resp, nil*/
}
