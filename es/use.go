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
/*func (s *statisticsRepo) commonQuery(condition *domain.EffectTotalCondition) *elastic.BoolQuery {
	query := elastic.NewBoolQuery()
	if condition.PayType != "" {
		query = query.Must(elastic.NewMatchQuery("ad_report_type.keyword", condition.PayType))
	}
	if !condition.StartDate.IsZero() && !condition.EndDate.IsZero() {
		query = query.Must(elastic.NewRangeQuery("ad_report_time").Gte(condition.StartDate).Lt(condition.EndDate))
	}
	for _, scene := range condition.Scenes {
		query = query.Should(elastic.NewTermQuery("ad_report_scene_code.keyword", scene))
		// 解决ES搜索，should和must共用，should失效的问题
		query.MinimumNumberShouldMatch(1)
	}
	switch condition.Tag {
	case ">=":
		query = query.Must(elastic.NewRangeQuery(condition.PayType).Gte(condition.Value))
	}
	for _, id := range condition.Customers {
		query = query.Should(elastic.NewTermQuery("customer_id", id))
		// 解决ES搜索，should和must共用，should失效的问题
		query.MinimumNumberShouldMatch(1)
	}

	for _, id := range condition.Merchants {
		query = query.Should(elastic.NewTermQuery("merchant_id", id))
		// 解决ES搜索，should和must共用，should失效的问题
		query.MinimumNumberShouldMatch(1)
	}

	for _, id := range condition.ShopIndustries {
		query = query.Should(elastic.NewTermQuery("shop_type_code", id))
		// 解决ES搜索，should和must共用，should失效的问题
		query.MinimumNumberShouldMatch(1)
	}

	for _, id := range condition.Cities {
		query = query.Should(elastic.NewTermQuery("code_city", id))
		// 解决ES搜索，should和must共用，should失效的问题
		query.MinimumNumberShouldMatch(1)
	}

	return query
}*/

/*
GET /es_datacenter/_search
{
    "_source": {
        "includes": [
            "serialnum",
            "ad_plan_starttime",
            "ad_plan_endtime",
            "ad_report_scene",
            "ad_plan_advertiser_name",
            "ad_plan_name",
            "merchant_name",
            "shop_type_top_name",
            "shop_type_name",
            "name_prov",
            "name_city",
            "name_coun",
            "customer_name",
            "ad_report_time"
        ]
    },
    "from": 0,
    "query": {
        "bool": {
            "must": [
                {
                    "match": {
                        "ad_plan_id": {
                            "query": 205
                        }
                    }
                },
                {
                    "match": {
                        "ad_report_type": {
                            "query": "cpc"
                        }
                    }
                }
            ]
        }
    },
    "size": 10
}
*/

/*

func (s *statisticsRepo) EffectGroupByDate(condition *domain.EffectCondition) ([]domain.DateValueObject, error) {
query := s.commonQuery(condition)
payTypeAgg := elastic.NewTermsAggregation().Field("date_num").Order("_key", true).Size(2000).
SubAggregation("cpc", elastic.NewSumAggregation().Field("cpc"))
payTypeAgg = payTypeAgg.SubAggregation("cpa", elastic.NewSumAggregation().Field("cpa"))
payTypeAgg = payTypeAgg.SubAggregation("cpm", elastic.NewSumAggregation().Field("cpm"))
//payTypeAgg.SubAggregation("test", elastic.NewBucketSelectorAggregation().AddBucketsPath("cpc", "cpc").Script(elastic.NewScript("params.cpc >=1")))
if condition.EffectRangeCondition.Value > 0 {
payTypeAgg.SubAggregation("test", elastic.NewBucketSelectorAggregation().AddBucketsPath(condition.EffectRangeCondition.PayType, condition.EffectRangeCondition.PayType).Script(elastic.NewScript(fmt.Sprintf("params.%s >= %d",condition.EffectRangeCondition.PayType, condition.EffectRangeCondition.Value))))
}
searchResult, err := EsClient().Search().Index(s.index).Size(2000).
Query(query).
Aggregation("group_by_date_num", payTypeAgg).
Pretty(true).
Do(context.Background())
if err != nil {
return nil, err
}
agg, found := searchResult.Aggregations.Terms("group_by_date_num")
if !found {
return nil, errors.New("should have a terms aggregation called " + "group_by_date_num")
}
var rows []domain.DateValueObject
// 遍历桶数据
for _, userBucket := range agg.Buckets {
// 每一个桶都有一个key值，其实就是分组的值，可以理解为SQL的group by值
effect := map[string]int{}
for i, sum := range userBucket.Aggregations {
switch i {
case "cpa", "cpm", "cpc":
temp := struct {Value float64}{}
_= json.Unmarshal(sum, &temp)
effect[i] = int(temp.Value)
}
}
// 添加上转化率
//effect.Rate = bussiness.CalcRate(effect.Cpc, effect.Cpa)
dateStr := pkg.IntToDateStr(int(userBucket.Key.(float64)))
rows = append(rows, domain.DateValueObject{Date: dateStr, Value: effect})
}

return rows, nil
}
*/
