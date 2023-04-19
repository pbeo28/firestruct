package main

var firestoreCloudEventJSON = `{
	"oldValue":{},
	"value":{
	   "name":"projects/myproject/databases/(default)/documents/samplecollection/e3a290b9-f307-423f-a0a5-67f5de651e1f",
	   "fields":{
		  "boolData":{
			 "booleanValue":true
		  },
		  "bytesData":{
			 "bytesValue":"SGVsbG8gV29ybGQ="
		  },
		  "doubleData":{
			 "doubleValue":987.123456
		  },
		  "geoPointData":{
			 "geoPointValue":{
				"latitude":51.205005708080876,
				"longitude":3.225345050850536
			 }
		  },
		  "intData":{
			 "integerValue":987654321
		  },
		  "nestedMapData":{
			 "mapValue":{
				"fields":{
				   "nestedArrayData":{
					  "arrayValue":{
						 "values":[
							{
							   "mapValue":{
								  "fields":{
									 "boolData":{
										"booleanValue":true
									 },
									 "bytesData":{
										"bytesValue":"SGVsbG8gV29ybGQ="
									 },
									 "doubleData":{
										"doubleValue":987.123456
									 },
									 "geoPointData":{
										"geoPointValue":{
										   "latitude":51.205005708080876,
										   "longitude":3.225345050850536
										}
									 },
									 "intData":{
										"integerValue":987654321
									 },
									 "nilData":{
										"nullValue":null
									 },
									 "referenceData":{
										"referenceValue":"/reference/path"
									 },
									 "stringData":{
										"stringValue":"Hello World"
									 },
									 "timeData":{
										"timestampValue":"2025-04-14T01:02:03Z"
									 },
									 "uuidData":{
										"stringValue":"e3a290b9-f307-423f-a0a5-67f5de651e1f"
									 }
								  }
							   }
							},
							{
							   "subNestedArrayData":{
								  "arrayValue":{
									 "values":[
										{
										   "timeData":{
											  "timestampValue":"2025-04-14T01:02:03Z"
										   }
										},
										{
										   "stringData":{
											  "stringValue":"Hello World"
										   }
										},
										{
										   "uuidData":{
											  "stringValue":"e3a290b9-f307-423f-a0a5-67f5de651e1f"
										   }
										},
										{
										   "boolData":{
											  "booleanValue":true
										   }
										},
										{
										   "intData":{
											  "integerValue":987654321
										   }
										},
										{
										   "doubleData":{
											  "doubleValue":987.123456
										   }
										},
										{
										   "bytesData":{
											  "bytesValue":"SGVsbG8gV29ybGQ="
										   }
										},
										{
										   "nilData":{
											  "nullValue":null
										   }
										},
										{
										   "referenceData":{
											  "referenceValue":"/reference/path"
										   }
										},
										{
										   "geoPointData":{
											  "geoPointValue":{
												 "latitude":51.205005708080876,
												 "longitude":3.225345050850536
											  }
										   }
										}
									 ]
								  }
							   }
							},
							{
							   "timeData":{
								  "timestampValue":"2025-04-14T01:02:03Z"
							   }
							},
							{
							   "stringData":{
								  "stringValue":"Hello World"
							   }
							},
							{
							   "uuidData":{
								  "stringValue":"e3a290b9-f307-423f-a0a5-67f5de651e1f"
							   }
							},
							{
							   "boolData":{
								  "booleanValue":true
							   }
							},
							{
							   "intData":{
								  "integerValue":987654321
							   }
							},
							{
							   "doubleData":{
								  "doubleValue":987.123456
							   }
							},
							{
							   "bytesData":{
								  "bytesValue":"SGVsbG8gV29ybGQ="
							   }
							},
							{
							   "nilData":{
								  "nullValue":null
							   }
							},
							{
							   "referenceData":{
								  "referenceValue":"/reference/path"
							   }
							},
							{
							   "geoPointData":{
								  "geoPointValue":{
									 "latitude":51.205005708080876,
									 "longitude":3.225345050850536
								  }
							   }
							}
						 ]
					  }
				   }
				}
			 }
		  },
		  "nilData":{
			 "nullValue":null
		  },
		  "referenceData":{
			 "referenceValue":"/reference/path"
		  },
		  "stringData":{
			 "stringValue":"Hello World"
		  },
		  "timeData":{
			 "timestampValue":"2025-04-14T01:02:03Z"
		  },
		  "uuidData":{
			 "stringValue":"e3a290b9-f307-423f-a0a5-67f5de651e1f"
		  }
	   },
	   "createTime":"2023-03-30T14:48:14.184895Z",
	   "updateTime":"2023-03-30T14:48:14.184895Z"
	},
	"updateMask":{
	   
	}
  }`
