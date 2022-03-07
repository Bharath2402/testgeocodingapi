package schema

const GeocodeRespPositive = `{
    "$schema": "http://json-schema.org/draft-07/schema",
    "$id": "http://example.com/example.json",
    "type": "array",
    "title": "The root schema",
    "description": "The root schema comprises the entire JSON document.",
    "default": {},
    "examples": [
        {
            "plus_code": {
                "compound_code": "XJCQ+634 Bengaluru, Karnataka, India",
                "global_code": "7J4VXJCQ+634"
            },
            "results": [
                {
                    "address_components": [
                        {
                            "long_name": "1073",
                            "short_name": "1073",
                            "types": [
                                "street_number"
                            ]
                        },
                        {
                            "long_name": "11th Main Road",
                            "short_name": "11th Main Rd",
                            "types": [
                                "route"
                            ]
                        },
                        {
                            "long_name": "HAL 2nd Stage",
                            "short_name": "HAL 2nd Stage",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_3"
                            ]
                        },
                        {
                            "long_name": "Appareddipalya",
                            "short_name": "Appareddipalya",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_2"
                            ]
                        },
                        {
                            "long_name": "Indiranagar",
                            "short_name": "Indiranagar",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_1"
                            ]
                        },
                        {
                            "long_name": "Bengaluru",
                            "short_name": "Bengaluru",
                            "types": [
                                "locality",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        },
                        {
                            "long_name": "560038",
                            "short_name": "560038",
                            "types": [
                                "postal_code"
                            ]
                        }
                    ],
                    "formatted_address": "1073, 11th Main Rd, HAL 2nd Stage, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560038, India",
                    "geometry": {
                        "location": {
                            "lat": 12.9704227,
                            "lng": 77.6377444
                        },
                        "location_type": "ROOFTOP",
                        "viewport": {
                            "northeast": {
                                "lat": 12.9717716802915,
                                "lng": 77.63909338029151
                            },
                            "southwest": {
                                "lat": 12.9690737197085,
                                "lng": 77.6363954197085
                            }
                        }
                    },
                    "place_id": "ChIJQagc8IgXrjsRyBIWMfs2oWk",
                    "plus_code": {
                        "compound_code": "XJCQ+53 Bengaluru, Karnataka, India",
                        "global_code": "7J4VXJCQ+53"
                    },
                    "types": [
                        "street_address"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "1074",
                            "short_name": "1074",
                            "types": [
                                "street_number"
                            ]
                        },
                        {
                            "long_name": "11th Main Road",
                            "short_name": "11th Main Rd",
                            "types": [
                                "route"
                            ]
                        },
                        {
                            "long_name": "Appareddipalya",
                            "short_name": "Appareddipalya",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_2"
                            ]
                        },
                        {
                            "long_name": "Indiranagar",
                            "short_name": "Indiranagar",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_1"
                            ]
                        },
                        {
                            "long_name": "Bengaluru",
                            "short_name": "Bengaluru",
                            "types": [
                                "locality",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        },
                        {
                            "long_name": "560008",
                            "short_name": "560008",
                            "types": [
                                "postal_code"
                            ]
                        }
                    ],
                    "formatted_address": "1074, 11th Main Rd, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560008, India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 12.9705685,
                                "lng": 77.63877040000001
                            },
                            "southwest": {
                                "lat": 12.9704977,
                                "lng": 77.63748629999999
                            }
                        },
                        "location": {
                            "lat": 12.9705331,
                            "lng": 77.63812829999999
                        },
                        "location_type": "GEOMETRIC_CENTER",
                        "viewport": {
                            "northeast": {
                                "lat": 12.9718820802915,
                                "lng": 77.6394773302915
                            },
                            "southwest": {
                                "lat": 12.9691841197085,
                                "lng": 77.63677936970849
                            }
                        }
                    },
                    "place_id": "ChIJfaYLa6cWrjsRZOyvDeJ7QFA",
                    "types": [
                        "route"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "XJCQ+63",
                            "short_name": "XJCQ+63",
                            "types": [
                                "plus_code"
                            ]
                        },
                        {
                            "long_name": "Bengaluru",
                            "short_name": "Bengaluru",
                            "types": [
                                "locality",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        }
                    ],
                    "formatted_address": "XJCQ+63 Bengaluru, Karnataka, India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 12.970625,
                                "lng": 77.63775
                            },
                            "southwest": {
                                "lat": 12.9705,
                                "lng": 77.637625
                            }
                        },
                        "location": {
                            "lat": 12.9705003,
                            "lng": 77.6377112
                        },
                        "location_type": "GEOMETRIC_CENTER",
                        "viewport": {
                            "northeast": {
                                "lat": 12.9719114802915,
                                "lng": 77.63903648029151
                            },
                            "southwest": {
                                "lat": 12.9692135197085,
                                "lng": 77.6363385197085
                            }
                        }
                    },
                    "place_id": "GhIJMYRSauXwKUARvxKjQtBoU0A",
                    "plus_code": {
                        "compound_code": "XJCQ+63 Bengaluru, Karnataka, India",
                        "global_code": "7J4VXJCQ+63"
                    },
                    "types": [
                        "plus_code"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "3266",
                            "short_name": "3266",
                            "types": [
                                "street_number"
                            ]
                        },
                        {
                            "long_name": "11th Main Road",
                            "short_name": "11th Main Rd",
                            "types": [
                                "route"
                            ]
                        },
                        {
                            "long_name": "Appareddipalya",
                            "short_name": "Appareddipalya",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_2"
                            ]
                        },
                        {
                            "long_name": "Indiranagar",
                            "short_name": "Indiranagar",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_1"
                            ]
                        },
                        {
                            "long_name": "Bengaluru",
                            "short_name": "Bengaluru",
                            "types": [
                                "locality",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        },
                        {
                            "long_name": "560008",
                            "short_name": "560008",
                            "types": [
                                "postal_code"
                            ]
                        }
                    ],
                    "formatted_address": "3266, 11th Main Rd, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560008, India",
                    "geometry": {
                        "location": {
                            "lat": 12.9704796,
                            "lng": 77.6370254
                        },
                        "location_type": "RANGE_INTERPOLATED",
                        "viewport": {
                            "northeast": {
                                "lat": 12.9718285802915,
                                "lng": 77.63837438029151
                            },
                            "southwest": {
                                "lat": 12.9691306197085,
                                "lng": 77.6356764197085
                            }
                        }
                    },
                    "place_id": "ElMzMjY2LCAxMXRoIE1haW4gUmQsIEFwcGFyZWRkaXBhbHlhLCBJbmRpcmFuYWdhciwgQmVuZ2FsdXJ1LCBLYXJuYXRha2EgNTYwMDA4LCBJbmRpYSIbEhkKFAoSCVU0MkunFq47Ea3KiZZAb2ZbEMIZ",
                    "types": [
                        "street_address"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "Appareddipalya",
                            "short_name": "Appareddipalya",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_2"
                            ]
                        },
                        {
                            "long_name": "Indiranagar",
                            "short_name": "Indiranagar",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_1"
                            ]
                        },
                        {
                            "long_name": "Bengaluru",
                            "short_name": "Bengaluru",
                            "types": [
                                "locality",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        }
                    ],
                    "formatted_address": "Appareddipalya, Indiranagar, Bengaluru, Karnataka, India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 12.972636,
                                "lng": 77.641284
                            },
                            "southwest": {
                                "lat": 12.9698529,
                                "lng": 77.6359049
                            }
                        },
                        "location": {
                            "lat": 12.9712206,
                            "lng": 77.6398059
                        },
                        "location_type": "APPROXIMATE",
                        "viewport": {
                            "northeast": {
                                "lat": 12.972636,
                                "lng": 77.641284
                            },
                            "southwest": {
                                "lat": 12.9698529,
                                "lng": 77.6359049
                            }
                        }
                    },
                    "place_id": "ChIJjzgrDqcWrjsR-yflGa4hcEA",
                    "types": [
                        "political",
                        "sublocality",
                        "sublocality_level_2"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "Indiranagar",
                            "short_name": "Indiranagar",
                            "types": [
                                "political",
                                "sublocality",
                                "sublocality_level_1"
                            ]
                        },
                        {
                            "long_name": "Bengaluru",
                            "short_name": "Bengaluru",
                            "types": [
                                "locality",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        }
                    ],
                    "formatted_address": "Indiranagar, Bengaluru, Karnataka, India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 12.985708,
                                "lng": 77.6481491
                            },
                            "southwest": {
                                "lat": 12.9601705,
                                "lng": 77.62831919999999
                            }
                        },
                        "location": {
                            "lat": 12.9783692,
                            "lng": 77.6408356
                        },
                        "location_type": "APPROXIMATE",
                        "viewport": {
                            "northeast": {
                                "lat": 12.985708,
                                "lng": 77.6481491
                            },
                            "southwest": {
                                "lat": 12.9601705,
                                "lng": 77.62831919999999
                            }
                        }
                    },
                    "place_id": "ChIJkQN3GKQWrjsRNhBQJrhGD7U",
                    "types": [
                        "political",
                        "sublocality",
                        "sublocality_level_1"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "560008",
                            "short_name": "560008",
                            "types": [
                                "postal_code"
                            ]
                        },
                        {
                            "long_name": "Bengaluru",
                            "short_name": "Bengaluru",
                            "types": [
                                "locality",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        }
                    ],
                    "formatted_address": "Bengaluru, Karnataka 560008, India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 12.9945063,
                                "lng": 77.65228549999999
                            },
                            "southwest": {
                                "lat": 12.9460051,
                                "lng": 77.6187999
                            }
                        },
                        "location": {
                            "lat": 12.9819939,
                            "lng": 77.6256131
                        },
                        "location_type": "APPROXIMATE",
                        "viewport": {
                            "northeast": {
                                "lat": 12.9945063,
                                "lng": 77.65228549999999
                            },
                            "southwest": {
                                "lat": 12.9460051,
                                "lng": 77.6187999
                            }
                        }
                    },
                    "place_id": "ChIJXRp0jKAWrjsRAMK3zN0Ibzk",
                    "types": [
                        "postal_code"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "Bengaluru",
                            "short_name": "Bengaluru",
                            "types": [
                                "locality",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        }
                    ],
                    "formatted_address": "Bengaluru, Karnataka, India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 13.173706,
                                "lng": 77.8826809
                            },
                            "southwest": {
                                "lat": 12.7342888,
                                "lng": 77.3791981
                            }
                        },
                        "location": {
                            "lat": 12.9715987,
                            "lng": 77.5945627
                        },
                        "location_type": "APPROXIMATE",
                        "viewport": {
                            "northeast": {
                                "lat": 13.173706,
                                "lng": 77.8826809
                            },
                            "southwest": {
                                "lat": 12.7342888,
                                "lng": 77.3791981
                            }
                        }
                    },
                    "place_id": "ChIJbU60yXAWrjsR4E9-UejD3_g",
                    "types": [
                        "locality",
                        "political"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "Bangalore Urban",
                            "short_name": "Bangalore Urban",
                            "types": [
                                "administrative_area_level_2",
                                "political"
                            ]
                        },
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        }
                    ],
                    "formatted_address": "Bangalore Urban, Karnataka, India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 13.22497,
                                "lng": 77.83501
                            },
                            "southwest": {
                                "lat": 12.65785,
                                "lng": 77.32867
                            }
                        },
                        "location": {
                            "lat": 12.9700247,
                            "lng": 77.6536125
                        },
                        "location_type": "APPROXIMATE",
                        "viewport": {
                            "northeast": {
                                "lat": 13.22497,
                                "lng": 77.83501
                            },
                            "southwest": {
                                "lat": 12.65785,
                                "lng": 77.32867
                            }
                        }
                    },
                    "place_id": "ChIJ_Q7NCkQWrjsRn2yP7O-T8Fg",
                    "types": [
                        "administrative_area_level_2",
                        "political"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "Karnataka",
                            "short_name": "KA",
                            "types": [
                                "administrative_area_level_1",
                                "political"
                            ]
                        },
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        }
                    ],
                    "formatted_address": "Karnataka, India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 18.4411689,
                                "lng": 78.5860101
                            },
                            "southwest": {
                                "lat": 11.593352,
                                "lng": 74.04960609999999
                            }
                        },
                        "location": {
                            "lat": 15.3172775,
                            "lng": 75.7138884
                        },
                        "location_type": "APPROXIMATE",
                        "viewport": {
                            "northeast": {
                                "lat": 18.4411689,
                                "lng": 78.5860101
                            },
                            "southwest": {
                                "lat": 11.593352,
                                "lng": 74.04960609999999
                            }
                        }
                    },
                    "place_id": "ChIJj0i_N0xaozsR1Xx10YzS8UE",
                    "types": [
                        "administrative_area_level_1",
                        "political"
                    ]
                },
                {
                    "address_components": [
                        {
                            "long_name": "India",
                            "short_name": "IN",
                            "types": [
                                "country",
                                "political"
                            ]
                        }
                    ],
                    "formatted_address": "India",
                    "geometry": {
                        "bounds": {
                            "northeast": {
                                "lat": 35.513327,
                                "lng": 97.39535869999999
                            },
                            "southwest": {
                                "lat": 6.4626999,
                                "lng": 68.1097
                            }
                        },
                        "location": {
                            "lat": 20.593684,
                            "lng": 78.96288
                        },
                        "location_type": "APPROXIMATE",
                        "viewport": {
                            "northeast": {
                                "lat": 35.513327,
                                "lng": 97.39535869999999
                            },
                            "southwest": {
                                "lat": 6.4626999,
                                "lng": 68.1097
                            }
                        }
                    },
                    "place_id": "ChIJkbeSa_BfYzARphNChaFPjNc",
                    "types": [
                        "country",
                        "political"
                    ]
                }
            ],
            "status": "OK"
        }
    ],
    "required": [
        "plus_code",
        "results",
        "status"
    ],
    "properties": {
        "plus_code": {
            "$id": "#/properties/plus_code",
            "type": "object",
            "title": "The plus_code schema",
            "description": "An explanation about the purpose of this instance.",
            "default": {},
            "examples": [
                {
                    "compound_code": "XJCQ+634 Bengaluru, Karnataka, India",
                    "global_code": "7J4VXJCQ+634"
                }
            ],
            "required": [
                "compound_code",
                "global_code"
            ],
            "properties": {
                "compound_code": {
                    "$id": "#/properties/plus_code/properties/compound_code",
                    "type": "string",
                    "title": "The compound_code schema",
                    "description": "An explanation about the purpose of this instance.",
                    "default": "",
                    "examples": [
                        "XJCQ+634 Bengaluru, Karnataka, India"
                    ]
                },
                "global_code": {
                    "$id": "#/properties/plus_code/properties/global_code",
                    "type": "string",
                    "title": "The global_code schema",
                    "description": "An explanation about the purpose of this instance.",
                    "default": "",
                    "examples": [
                        "7J4VXJCQ+634"
                    ]
                }
            },
            "additionalProperties": true
        },
        "results": {
            "$id": "#/properties/results",
            "type": "array",
            "title": "The results schema",
            "description": "An explanation about the purpose of this instance.",
            "default": [],
            "examples": [
                [
                    {
                        "address_components": [
                            {
                                "long_name": "1073",
                                "short_name": "1073",
                                "types": [
                                    "street_number"
                                ]
                            },
                            {
                                "long_name": "11th Main Road",
                                "short_name": "11th Main Rd",
                                "types": [
                                    "route"
                                ]
                            },
                            {
                                "long_name": "HAL 2nd Stage",
                                "short_name": "HAL 2nd Stage",
                                "types": [
                                    "political",
                                    "sublocality",
                                    "sublocality_level_3"
                                ]
                            },
                            {
                                "long_name": "Appareddipalya",
                                "short_name": "Appareddipalya",
                                "types": [
                                    "political",
                                    "sublocality",
                                    "sublocality_level_2"
                                ]
                            },
                            {
                                "long_name": "Indiranagar",
                                "short_name": "Indiranagar",
                                "types": [
                                    "political",
                                    "sublocality",
                                    "sublocality_level_1"
                                ]
                            },
                            {
                                "long_name": "Bengaluru",
                                "short_name": "Bengaluru",
                                "types": [
                                    "locality",
                                    "political"
                                ]
                            },
                            {
                                "long_name": "Bangalore Urban",
                                "short_name": "Bangalore Urban",
                                "types": [
                                    "administrative_area_level_2",
                                    "political"
                                ]
                            },
                            {
                                "long_name": "Karnataka",
                                "short_name": "KA",
                                "types": [
                                    "administrative_area_level_1",
                                    "political"
                                ]
                            },
                            {
                                "long_name": "India",
                                "short_name": "IN",
                                "types": [
                                    "country",
                                    "political"
                                ]
                            },
                            {
                                "long_name": "560038",
                                "short_name": "560038",
                                "types": [
                                    "postal_code"
                                ]
                            }
                        ],
                        "formatted_address": "1073, 11th Main Rd, HAL 2nd Stage, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560038, India",
                        "geometry": {
                            "location": {
                                "lat": 12.9704227,
                                "lng": 77.6377444
                            },
                            "location_type": "ROOFTOP",
                            "viewport": {
                                "northeast": {
                                    "lat": 12.9717716802915,
                                    "lng": 77.63909338029151
                                },
                                "southwest": {
                                    "lat": 12.9690737197085,
                                    "lng": 77.6363954197085
                                }
                            }
                        },
                        "place_id": "ChIJQagc8IgXrjsRyBIWMfs2oWk",
                        "plus_code": {
                            "compound_code": "XJCQ+53 Bengaluru, Karnataka, India",
                            "global_code": "7J4VXJCQ+53"
                        },
                        "types": [
                            "street_address"
                        ]
                    },
                    {
                        "address_components": [
                            {
                                "long_name": "1074",
                                "short_name": "1074",
                                "types": [
                                    "street_number"
                                ]
                            },
                            {
                                "long_name": "11th Main Road",
                                "short_name": "11th Main Rd",
                                "types": [
                                    "route"
                                ]
                            },
                            {
                                "long_name": "Appareddipalya",
                                "short_name": "Appareddipalya",
                                "types": [
                                    "political",
                                    "sublocality",
                                    "sublocality_level_2"
                                ]
                            },
                            {
                                "long_name": "Indiranagar",
                                "short_name": "Indiranagar",
                                "types": [
                                    "political",
                                    "sublocality",
                                    "sublocality_level_1"
                                ]
                            },
                            {
                                "long_name": "Bengaluru",
                                "short_name": "Bengaluru",
                                "types": [
                                    "locality",
                                    "political"
                                ]
                            },
                            {
                                "long_name": "Bangalore Urban",
                                "short_name": "Bangalore Urban",
                                "types": [
                                    "administrative_area_level_2",
                                    "political"
                                ]
                            },
                            {
                                "long_name": "Karnataka",
                                "short_name": "KA",
                                "types": [
                                    "administrative_area_level_1",
                                    "political"
                                ]
                            },
                            {
                                "long_name": "India",
                                "short_name": "IN",
                                "types": [
                                    "country",
                                    "political"
                                ]
                            },
                            {
                                "long_name": "560008",
                                "short_name": "560008",
                                "types": [
                                    "postal_code"
                                ]
                            }
                        ],
                        "formatted_address": "1074, 11th Main Rd, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560008, India",
                        "geometry": {
                            "bounds": {
                                "northeast": {
                                    "lat": 12.9705685,
                                    "lng": 77.63877040000001
                                },
                                "southwest": {
                                    "lat": 12.9704977,
                                    "lng": 77.63748629999999
                                }
                            },
                            "location": {
                                "lat": 12.9705331,
                                "lng": 77.63812829999999
                            },
                            "location_type": "GEOMETRIC_CENTER",
                            "viewport": {
                                "northeast": {
                                    "lat": 12.9718820802915,
                                    "lng": 77.6394773302915
                                },
                                "southwest": {
                                    "lat": 12.9691841197085,
                                    "lng": 77.63677936970849
                                }
                            }
                        },
                        "place_id": "ChIJfaYLa6cWrjsRZOyvDeJ7QFA",
                        "types": [
                            "route"
                        ]
                    }
                ]
            ],
            "additionalItems": true,
            "items": {
                "$id": "#/properties/results/items",
                "anyOf": [
                    {
                        "$id": "#/properties/results/items/anyOf/0",
                        "type": "object",
                        "title": "The first anyOf schema",
                        "description": "An explanation about the purpose of this instance.",
                        "default": {},
                        "examples": [
                            {
                                "address_components": [
                                    {
                                        "long_name": "1073",
                                        "short_name": "1073",
                                        "types": [
                                            "street_number"
                                        ]
                                    },
                                    {
                                        "long_name": "11th Main Road",
                                        "short_name": "11th Main Rd",
                                        "types": [
                                            "route"
                                        ]
                                    },
                                    {
                                        "long_name": "HAL 2nd Stage",
                                        "short_name": "HAL 2nd Stage",
                                        "types": [
                                            "political",
                                            "sublocality",
                                            "sublocality_level_3"
                                        ]
                                    },
                                    {
                                        "long_name": "Appareddipalya",
                                        "short_name": "Appareddipalya",
                                        "types": [
                                            "political",
                                            "sublocality",
                                            "sublocality_level_2"
                                        ]
                                    },
                                    {
                                        "long_name": "Indiranagar",
                                        "short_name": "Indiranagar",
                                        "types": [
                                            "political",
                                            "sublocality",
                                            "sublocality_level_1"
                                        ]
                                    },
                                    {
                                        "long_name": "Bengaluru",
                                        "short_name": "Bengaluru",
                                        "types": [
                                            "locality",
                                            "political"
                                        ]
                                    },
                                    {
                                        "long_name": "Bangalore Urban",
                                        "short_name": "Bangalore Urban",
                                        "types": [
                                            "administrative_area_level_2",
                                            "political"
                                        ]
                                    },
                                    {
                                        "long_name": "Karnataka",
                                        "short_name": "KA",
                                        "types": [
                                            "administrative_area_level_1",
                                            "political"
                                        ]
                                    },
                                    {
                                        "long_name": "India",
                                        "short_name": "IN",
                                        "types": [
                                            "country",
                                            "political"
                                        ]
                                    },
                                    {
                                        "long_name": "560038",
                                        "short_name": "560038",
                                        "types": [
                                            "postal_code"
                                        ]
                                    }
                                ],
                                "formatted_address": "1073, 11th Main Rd, HAL 2nd Stage, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560038, India",
                                "geometry": {
                                    "location": {
                                        "lat": 12.9704227,
                                        "lng": 77.6377444
                                    },
                                    "location_type": "ROOFTOP",
                                    "viewport": {
                                        "northeast": {
                                            "lat": 12.9717716802915,
                                            "lng": 77.63909338029151
                                        },
                                        "southwest": {
                                            "lat": 12.9690737197085,
                                            "lng": 77.6363954197085
                                        }
                                    }
                                },
                                "place_id": "ChIJQagc8IgXrjsRyBIWMfs2oWk",
                                "plus_code": {
                                    "compound_code": "XJCQ+53 Bengaluru, Karnataka, India",
                                    "global_code": "7J4VXJCQ+53"
                                },
                                "types": [
                                    "street_address"
                                ]
                            }
                        ],
                        "required": [
                            "address_components",
                            "formatted_address",
                            "geometry",
                            "place_id",
                            "plus_code",
                            "types"
                        ],
                        "properties": {
                            "address_components": {
                                "$id": "#/properties/results/items/anyOf/0/properties/address_components",
                                "type": "array",
                                "title": "The address_components schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": [],
                                "examples": [
                                    [
                                        {
                                            "long_name": "1073",
                                            "short_name": "1073",
                                            "types": [
                                                "street_number"
                                            ]
                                        },
                                        {
                                            "long_name": "11th Main Road",
                                            "short_name": "11th Main Rd",
                                            "types": [
                                                "route"
                                            ]
                                        }
                                    ]
                                ],
                                "additionalItems": true,
                                "items": {
                                    "$id": "#/properties/results/items/anyOf/0/properties/address_components/items",
                                    "anyOf": [
                                        {
                                            "$id": "#/properties/results/items/anyOf/0/properties/address_components/items/anyOf/0",
                                            "type": "object",
                                            "title": "The first anyOf schema",
                                            "description": "An explanation about the purpose of this instance.",
                                            "default": {},
                                            "examples": [
                                                {
                                                    "long_name": "1073",
                                                    "short_name": "1073",
                                                    "types": [
                                                        "street_number"
                                                    ]
                                                }
                                            ],
                                            "required": [
                                                "long_name",
                                                "short_name",
                                                "types"
                                            ],
                                            "properties": {
                                                "long_name": {
                                                    "$id": "#/properties/results/items/anyOf/0/properties/address_components/items/anyOf/0/properties/long_name",
                                                    "type": "string",
                                                    "title": "The long_name schema",
                                                    "description": "An explanation about the purpose of this instance.",
                                                    "default": "",
                                                    "examples": [
                                                        "1073"
                                                    ]
                                                },
                                                "short_name": {
                                                    "$id": "#/properties/results/items/anyOf/0/properties/address_components/items/anyOf/0/properties/short_name",
                                                    "type": "string",
                                                    "title": "The short_name schema",
                                                    "description": "An explanation about the purpose of this instance.",
                                                    "default": "",
                                                    "examples": [
                                                        "1073"
                                                    ]
                                                },
                                                "types": {
                                                    "$id": "#/properties/results/items/anyOf/0/properties/address_components/items/anyOf/0/properties/types",
                                                    "type": "array",
                                                    "title": "The types schema",
                                                    "description": "An explanation about the purpose of this instance.",
                                                    "default": [],
                                                    "examples": [
                                                        [
                                                            "street_number"
                                                        ]
                                                    ],
                                                    "additionalItems": true,
                                                    "items": {
                                                        "$id": "#/properties/results/items/anyOf/0/properties/address_components/items/anyOf/0/properties/types/items",
                                                        "anyOf": [
                                                            {
                                                                "$id": "#/properties/results/items/anyOf/0/properties/address_components/items/anyOf/0/properties/types/items/anyOf/0",
                                                                "type": "string",
                                                                "title": "The first anyOf schema",
                                                                "description": "An explanation about the purpose of this instance.",
                                                                "default": "",
                                                                "examples": [
                                                                    "street_number"
                                                                ]
                                                            }
                                                        ]
                                                    }
                                                }
                                            },
                                            "additionalProperties": true
                                        }
                                    ]
                                }
                            },
                            "formatted_address": {
                                "$id": "#/properties/results/items/anyOf/0/properties/formatted_address",
                                "type": "string",
                                "title": "The formatted_address schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": "",
                                "examples": [
                                    "1073, 11th Main Rd, HAL 2nd Stage, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560038, India"
                                ]
                            },
                            "geometry": {
                                "$id": "#/properties/results/items/anyOf/0/properties/geometry",
                                "type": "object",
                                "title": "The geometry schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": {},
                                "examples": [
                                    {
                                        "location": {
                                            "lat": 12.9704227,
                                            "lng": 77.6377444
                                        },
                                        "location_type": "ROOFTOP",
                                        "viewport": {
                                            "northeast": {
                                                "lat": 12.9717716802915,
                                                "lng": 77.63909338029151
                                            },
                                            "southwest": {
                                                "lat": 12.9690737197085,
                                                "lng": 77.6363954197085
                                            }
                                        }
                                    }
                                ],
                                "required": [
                                    "location",
                                    "location_type",
                                    "viewport"
                                ],
                                "properties": {
                                    "location": {
                                        "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/location",
                                        "type": "object",
                                        "title": "The location schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": {},
                                        "examples": [
                                            {
                                                "lat": 12.9704227,
                                                "lng": 77.6377444
                                            }
                                        ],
                                        "required": [
                                            "lat",
                                            "lng"
                                        ],
                                        "properties": {
                                            "lat": {
                                                "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/location/properties/lat",
                                                "type": "number",
                                                "title": "The lat schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": 0.0,
                                                "examples": [
                                                    12.9704227
                                                ]
                                            },
                                            "lng": {
                                                "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/location/properties/lng",
                                                "type": "number",
                                                "title": "The lng schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": 0.0,
                                                "examples": [
                                                    77.6377444
                                                ]
                                            }
                                        },
                                        "additionalProperties": true
                                    },
                                    "location_type": {
                                        "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/location_type",
                                        "type": "string",
                                        "title": "The location_type schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": "",
                                        "examples": [
                                            "ROOFTOP"
                                        ]
                                    },
                                    "viewport": {
                                        "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/viewport",
                                        "type": "object",
                                        "title": "The viewport schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": {},
                                        "examples": [
                                            {
                                                "northeast": {
                                                    "lat": 12.9717716802915,
                                                    "lng": 77.63909338029151
                                                },
                                                "southwest": {
                                                    "lat": 12.9690737197085,
                                                    "lng": 77.6363954197085
                                                }
                                            }
                                        ],
                                        "required": [
                                            "northeast",
                                            "southwest"
                                        ],
                                        "properties": {
                                            "northeast": {
                                                "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/viewport/properties/northeast",
                                                "type": "object",
                                                "title": "The northeast schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": {},
                                                "examples": [
                                                    {
                                                        "lat": 12.9717716802915,
                                                        "lng": 77.63909338029151
                                                    }
                                                ],
                                                "required": [
                                                    "lat",
                                                    "lng"
                                                ],
                                                "properties": {
                                                    "lat": {
                                                        "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/viewport/properties/northeast/properties/lat",
                                                        "type": "number",
                                                        "title": "The lat schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            12.9717716802915
                                                        ]
                                                    },
                                                    "lng": {
                                                        "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/viewport/properties/northeast/properties/lng",
                                                        "type": "number",
                                                        "title": "The lng schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            77.63909338029151
                                                        ]
                                                    }
                                                },
                                                "additionalProperties": true
                                            },
                                            "southwest": {
                                                "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/viewport/properties/southwest",
                                                "type": "object",
                                                "title": "The southwest schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": {},
                                                "examples": [
                                                    {
                                                        "lat": 12.9690737197085,
                                                        "lng": 77.6363954197085
                                                    }
                                                ],
                                                "required": [
                                                    "lat",
                                                    "lng"
                                                ],
                                                "properties": {
                                                    "lat": {
                                                        "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/viewport/properties/southwest/properties/lat",
                                                        "type": "number",
                                                        "title": "The lat schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            12.9690737197085
                                                        ]
                                                    },
                                                    "lng": {
                                                        "$id": "#/properties/results/items/anyOf/0/properties/geometry/properties/viewport/properties/southwest/properties/lng",
                                                        "type": "number",
                                                        "title": "The lng schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            77.6363954197085
                                                        ]
                                                    }
                                                },
                                                "additionalProperties": true
                                            }
                                        },
                                        "additionalProperties": true
                                    }
                                },
                                "additionalProperties": true
                            },
                            "place_id": {
                                "$id": "#/properties/results/items/anyOf/0/properties/place_id",
                                "type": "string",
                                "title": "The place_id schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": "",
                                "examples": [
                                    "ChIJQagc8IgXrjsRyBIWMfs2oWk"
                                ]
                            },
                            "plus_code": {
                                "$id": "#/properties/results/items/anyOf/0/properties/plus_code",
                                "type": "object",
                                "title": "The plus_code schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": {},
                                "examples": [
                                    {
                                        "compound_code": "XJCQ+53 Bengaluru, Karnataka, India",
                                        "global_code": "7J4VXJCQ+53"
                                    }
                                ],
                                "required": [
                                    "compound_code",
                                    "global_code"
                                ],
                                "properties": {
                                    "compound_code": {
                                        "$id": "#/properties/results/items/anyOf/0/properties/plus_code/properties/compound_code",
                                        "type": "string",
                                        "title": "The compound_code schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": "",
                                        "examples": [
                                            "XJCQ+53 Bengaluru, Karnataka, India"
                                        ]
                                    },
                                    "global_code": {
                                        "$id": "#/properties/results/items/anyOf/0/properties/plus_code/properties/global_code",
                                        "type": "string",
                                        "title": "The global_code schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": "",
                                        "examples": [
                                            "7J4VXJCQ+53"
                                        ]
                                    }
                                },
                                "additionalProperties": true
                            },
                            "types": {
                                "$id": "#/properties/results/items/anyOf/0/properties/types",
                                "type": "array",
                                "title": "The types schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": [],
                                "examples": [
                                    [
                                        "street_address"
                                    ]
                                ],
                                "additionalItems": true,
                                "items": {
                                    "$id": "#/properties/results/items/anyOf/0/properties/types/items",
                                    "anyOf": [
                                        {
                                            "$id": "#/properties/results/items/anyOf/0/properties/types/items/anyOf/0",
                                            "type": "string",
                                            "title": "The first anyOf schema",
                                            "description": "An explanation about the purpose of this instance.",
                                            "default": "",
                                            "examples": [
                                                "street_address"
                                            ]
                                        }
                                    ]
                                }
                            }
                        },
                        "additionalProperties": true
                    },
                    {
                        "$id": "#/properties/results/items/anyOf/1",
                        "type": "object",
                        "title": "The second anyOf schema",
                        "description": "An explanation about the purpose of this instance.",
                        "default": {},
                        "examples": [
                            {
                                "address_components": [
                                    {
                                        "long_name": "1074",
                                        "short_name": "1074",
                                        "types": [
                                            "street_number"
                                        ]
                                    },
                                    {
                                        "long_name": "11th Main Road",
                                        "short_name": "11th Main Rd",
                                        "types": [
                                            "route"
                                        ]
                                    },
                                    {
                                        "long_name": "Appareddipalya",
                                        "short_name": "Appareddipalya",
                                        "types": [
                                            "political",
                                            "sublocality",
                                            "sublocality_level_2"
                                        ]
                                    },
                                    {
                                        "long_name": "Indiranagar",
                                        "short_name": "Indiranagar",
                                        "types": [
                                            "political",
                                            "sublocality",
                                            "sublocality_level_1"
                                        ]
                                    },
                                    {
                                        "long_name": "Bengaluru",
                                        "short_name": "Bengaluru",
                                        "types": [
                                            "locality",
                                            "political"
                                        ]
                                    },
                                    {
                                        "long_name": "Bangalore Urban",
                                        "short_name": "Bangalore Urban",
                                        "types": [
                                            "administrative_area_level_2",
                                            "political"
                                        ]
                                    },
                                    {
                                        "long_name": "Karnataka",
                                        "short_name": "KA",
                                        "types": [
                                            "administrative_area_level_1",
                                            "political"
                                        ]
                                    },
                                    {
                                        "long_name": "India",
                                        "short_name": "IN",
                                        "types": [
                                            "country",
                                            "political"
                                        ]
                                    },
                                    {
                                        "long_name": "560008",
                                        "short_name": "560008",
                                        "types": [
                                            "postal_code"
                                        ]
                                    }
                                ],
                                "formatted_address": "1074, 11th Main Rd, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560008, India",
                                "geometry": {
                                    "bounds": {
                                        "northeast": {
                                            "lat": 12.9705685,
                                            "lng": 77.63877040000001
                                        },
                                        "southwest": {
                                            "lat": 12.9704977,
                                            "lng": 77.63748629999999
                                        }
                                    },
                                    "location": {
                                        "lat": 12.9705331,
                                        "lng": 77.63812829999999
                                    },
                                    "location_type": "GEOMETRIC_CENTER",
                                    "viewport": {
                                        "northeast": {
                                            "lat": 12.9718820802915,
                                            "lng": 77.6394773302915
                                        },
                                        "southwest": {
                                            "lat": 12.9691841197085,
                                            "lng": 77.63677936970849
                                        }
                                    }
                                },
                                "place_id": "ChIJfaYLa6cWrjsRZOyvDeJ7QFA",
                                "types": [
                                    "route"
                                ]
                            }
                        ],
                        "required": [
                            "address_components",
                            "formatted_address",
                            "geometry",
                            "place_id",
                            "types"
                        ],
                        "properties": {
                            "address_components": {
                                "$id": "#/properties/results/items/anyOf/1/properties/address_components",
                                "type": "array",
                                "title": "The address_components schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": [],
                                "examples": [
                                    [
                                        {
                                            "long_name": "1074",
                                            "short_name": "1074",
                                            "types": [
                                                "street_number"
                                            ]
                                        },
                                        {
                                            "long_name": "11th Main Road",
                                            "short_name": "11th Main Rd",
                                            "types": [
                                                "route"
                                            ]
                                        }
                                    ]
                                ],
                                "additionalItems": true,
                                "items": {
                                    "$id": "#/properties/results/items/anyOf/1/properties/address_components/items",
                                    "anyOf": [
                                        {
                                            "$id": "#/properties/results/items/anyOf/1/properties/address_components/items/anyOf/0",
                                            "type": "object",
                                            "title": "The first anyOf schema",
                                            "description": "An explanation about the purpose of this instance.",
                                            "default": {},
                                            "examples": [
                                                {
                                                    "long_name": "1074",
                                                    "short_name": "1074",
                                                    "types": [
                                                        "street_number"
                                                    ]
                                                }
                                            ],
                                            "required": [
                                                "long_name",
                                                "short_name",
                                                "types"
                                            ],
                                            "properties": {
                                                "long_name": {
                                                    "$id": "#/properties/results/items/anyOf/1/properties/address_components/items/anyOf/0/properties/long_name",
                                                    "type": "string",
                                                    "title": "The long_name schema",
                                                    "description": "An explanation about the purpose of this instance.",
                                                    "default": "",
                                                    "examples": [
                                                        "1074"
                                                    ]
                                                },
                                                "short_name": {
                                                    "$id": "#/properties/results/items/anyOf/1/properties/address_components/items/anyOf/0/properties/short_name",
                                                    "type": "string",
                                                    "title": "The short_name schema",
                                                    "description": "An explanation about the purpose of this instance.",
                                                    "default": "",
                                                    "examples": [
                                                        "1074"
                                                    ]
                                                },
                                                "types": {
                                                    "$id": "#/properties/results/items/anyOf/1/properties/address_components/items/anyOf/0/properties/types",
                                                    "type": "array",
                                                    "title": "The types schema",
                                                    "description": "An explanation about the purpose of this instance.",
                                                    "default": [],
                                                    "examples": [
                                                        [
                                                            "street_number"
                                                        ]
                                                    ],
                                                    "additionalItems": true,
                                                    "items": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/address_components/items/anyOf/0/properties/types/items",
                                                        "anyOf": [
                                                            {
                                                                "$id": "#/properties/results/items/anyOf/1/properties/address_components/items/anyOf/0/properties/types/items/anyOf/0",
                                                                "type": "string",
                                                                "title": "The first anyOf schema",
                                                                "description": "An explanation about the purpose of this instance.",
                                                                "default": "",
                                                                "examples": [
                                                                    "street_number"
                                                                ]
                                                            }
                                                        ]
                                                    }
                                                }
                                            },
                                            "additionalProperties": true
                                        }
                                    ]
                                }
                            },
                            "formatted_address": {
                                "$id": "#/properties/results/items/anyOf/1/properties/formatted_address",
                                "type": "string",
                                "title": "The formatted_address schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": "",
                                "examples": [
                                    "1074, 11th Main Rd, Appareddipalya, Indiranagar, Bengaluru, Karnataka 560008, India"
                                ]
                            },
                            "geometry": {
                                "$id": "#/properties/results/items/anyOf/1/properties/geometry",
                                "type": "object",
                                "title": "The geometry schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": {},
                                "examples": [
                                    {
                                        "bounds": {
                                            "northeast": {
                                                "lat": 12.9705685,
                                                "lng": 77.63877040000001
                                            },
                                            "southwest": {
                                                "lat": 12.9704977,
                                                "lng": 77.63748629999999
                                            }
                                        },
                                        "location": {
                                            "lat": 12.9705331,
                                            "lng": 77.63812829999999
                                        },
                                        "location_type": "GEOMETRIC_CENTER",
                                        "viewport": {
                                            "northeast": {
                                                "lat": 12.9718820802915,
                                                "lng": 77.6394773302915
                                            },
                                            "southwest": {
                                                "lat": 12.9691841197085,
                                                "lng": 77.63677936970849
                                            }
                                        }
                                    }
                                ],
                                "required": [
                                    "bounds",
                                    "location",
                                    "location_type",
                                    "viewport"
                                ],
                                "properties": {
                                    "bounds": {
                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/bounds",
                                        "type": "object",
                                        "title": "The bounds schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": {},
                                        "examples": [
                                            {
                                                "northeast": {
                                                    "lat": 12.9705685,
                                                    "lng": 77.63877040000001
                                                },
                                                "southwest": {
                                                    "lat": 12.9704977,
                                                    "lng": 77.63748629999999
                                                }
                                            }
                                        ],
                                        "required": [
                                            "northeast",
                                            "southwest"
                                        ],
                                        "properties": {
                                            "northeast": {
                                                "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/bounds/properties/northeast",
                                                "type": "object",
                                                "title": "The northeast schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": {},
                                                "examples": [
                                                    {
                                                        "lat": 12.9705685,
                                                        "lng": 77.63877040000001
                                                    }
                                                ],
                                                "required": [
                                                    "lat",
                                                    "lng"
                                                ],
                                                "properties": {
                                                    "lat": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/bounds/properties/northeast/properties/lat",
                                                        "type": "number",
                                                        "title": "The lat schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            12.9705685
                                                        ]
                                                    },
                                                    "lng": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/bounds/properties/northeast/properties/lng",
                                                        "type": "number",
                                                        "title": "The lng schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            77.63877040000001
                                                        ]
                                                    }
                                                },
                                                "additionalProperties": true
                                            },
                                            "southwest": {
                                                "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/bounds/properties/southwest",
                                                "type": "object",
                                                "title": "The southwest schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": {},
                                                "examples": [
                                                    {
                                                        "lat": 12.9704977,
                                                        "lng": 77.63748629999999
                                                    }
                                                ],
                                                "required": [
                                                    "lat",
                                                    "lng"
                                                ],
                                                "properties": {
                                                    "lat": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/bounds/properties/southwest/properties/lat",
                                                        "type": "number",
                                                        "title": "The lat schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            12.9704977
                                                        ]
                                                    },
                                                    "lng": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/bounds/properties/southwest/properties/lng",
                                                        "type": "number",
                                                        "title": "The lng schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            77.63748629999999
                                                        ]
                                                    }
                                                },
                                                "additionalProperties": true
                                            }
                                        },
                                        "additionalProperties": true
                                    },
                                    "location": {
                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/location",
                                        "type": "object",
                                        "title": "The location schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": {},
                                        "examples": [
                                            {
                                                "lat": 12.9705331,
                                                "lng": 77.63812829999999
                                            }
                                        ],
                                        "required": [
                                            "lat",
                                            "lng"
                                        ],
                                        "properties": {
                                            "lat": {
                                                "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/location/properties/lat",
                                                "type": "number",
                                                "title": "The lat schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": 0.0,
                                                "examples": [
                                                    12.9705331
                                                ]
                                            },
                                            "lng": {
                                                "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/location/properties/lng",
                                                "type": "number",
                                                "title": "The lng schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": 0.0,
                                                "examples": [
                                                    77.63812829999999
                                                ]
                                            }
                                        },
                                        "additionalProperties": true
                                    },
                                    "location_type": {
                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/location_type",
                                        "type": "string",
                                        "title": "The location_type schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": "",
                                        "examples": [
                                            "GEOMETRIC_CENTER"
                                        ]
                                    },
                                    "viewport": {
                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/viewport",
                                        "type": "object",
                                        "title": "The viewport schema",
                                        "description": "An explanation about the purpose of this instance.",
                                        "default": {},
                                        "examples": [
                                            {
                                                "northeast": {
                                                    "lat": 12.9718820802915,
                                                    "lng": 77.6394773302915
                                                },
                                                "southwest": {
                                                    "lat": 12.9691841197085,
                                                    "lng": 77.63677936970849
                                                }
                                            }
                                        ],
                                        "required": [
                                            "northeast",
                                            "southwest"
                                        ],
                                        "properties": {
                                            "northeast": {
                                                "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/viewport/properties/northeast",
                                                "type": "object",
                                                "title": "The northeast schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": {},
                                                "examples": [
                                                    {
                                                        "lat": 12.9718820802915,
                                                        "lng": 77.6394773302915
                                                    }
                                                ],
                                                "required": [
                                                    "lat",
                                                    "lng"
                                                ],
                                                "properties": {
                                                    "lat": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/viewport/properties/northeast/properties/lat",
                                                        "type": "number",
                                                        "title": "The lat schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            12.9718820802915
                                                        ]
                                                    },
                                                    "lng": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/viewport/properties/northeast/properties/lng",
                                                        "type": "number",
                                                        "title": "The lng schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            77.6394773302915
                                                        ]
                                                    }
                                                },
                                                "additionalProperties": true
                                            },
                                            "southwest": {
                                                "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/viewport/properties/southwest",
                                                "type": "object",
                                                "title": "The southwest schema",
                                                "description": "An explanation about the purpose of this instance.",
                                                "default": {},
                                                "examples": [
                                                    {
                                                        "lat": 12.9691841197085,
                                                        "lng": 77.63677936970849
                                                    }
                                                ],
                                                "required": [
                                                    "lat",
                                                    "lng"
                                                ],
                                                "properties": {
                                                    "lat": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/viewport/properties/southwest/properties/lat",
                                                        "type": "number",
                                                        "title": "The lat schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            12.9691841197085
                                                        ]
                                                    },
                                                    "lng": {
                                                        "$id": "#/properties/results/items/anyOf/1/properties/geometry/properties/viewport/properties/southwest/properties/lng",
                                                        "type": "number",
                                                        "title": "The lng schema",
                                                        "description": "An explanation about the purpose of this instance.",
                                                        "default": 0.0,
                                                        "examples": [
                                                            77.63677936970849
                                                        ]
                                                    }
                                                },
                                                "additionalProperties": true
                                            }
                                        },
                                        "additionalProperties": true
                                    }
                                },
                                "additionalProperties": true
                            },
                            "place_id": {
                                "$id": "#/properties/results/items/anyOf/1/properties/place_id",
                                "type": "string",
                                "title": "The place_id schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": "",
                                "examples": [
                                    "ChIJfaYLa6cWrjsRZOyvDeJ7QFA"
                                ]
                            },
                            "types": {
                                "$id": "#/properties/results/items/anyOf/1/properties/types",
                                "type": "array",
                                "title": "The types schema",
                                "description": "An explanation about the purpose of this instance.",
                                "default": [],
                                "examples": [
                                    [
                                        "route"
                                    ]
                                ],
                                "additionalItems": true,
                                "items": {
                                    "$id": "#/properties/results/items/anyOf/1/properties/types/items",
                                    "anyOf": [
                                        {
                                            "$id": "#/properties/results/items/anyOf/1/properties/types/items/anyOf/0",
                                            "type": "string",
                                            "title": "The first anyOf schema",
                                            "description": "An explanation about the purpose of this instance.",
                                            "default": "",
                                            "examples": [
                                                "route"
                                            ]
                                        }
                                    ]
                                }
                            }
                        },
                        "additionalProperties": true
                    }
                ]
            }
        },
        "status": {
            "$id": "#/properties/status",
            "type": "string",
            "title": "The status schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "OK"
            ]
        }
    },
    "additionalProperties": true
}`

const GeocodeRespNegative = `{
    "$schema": "http://json-schema.org/draft-07/schema",
    "$id": "http://example.com/example.json",
    "type": "object",
    "title": "The root schema",
    "description": "The root schema comprises the entire JSON document.",
    "default": {},
    "examples": [
        {
            "error_message": "Invalid request. Missing the 'address', 'components', 'latlng' or 'place_id' parameter.",
            "results": [],
            "status": "INVALID_REQUEST"
        }
    ],
    "required": [
        "error_message",
        "results",
        "status"
    ],
    "properties": {
        "error_message": {
            "$id": "#/properties/error_message",
            "type": "string",
            "title": "The error_message schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "Invalid request. Missing the 'address', 'components', 'latlng' or 'place_id' parameter."
            ]
        },
        "results": {
            "$id": "#/properties/results",
            "type": "array",
            "title": "The results schema",
            "description": "An explanation about the purpose of this instance.",
            "default": [],
            "examples": [
                []
            ],
            "additionalItems": true,
            "items": {
                "$id": "#/properties/results/items"
            }
        },
        "status": {
            "$id": "#/properties/status",
            "type": "string",
            "title": "The status schema",
            "description": "An explanation about the purpose of this instance.",
            "default": "",
            "examples": [
                "INVALID_REQUEST"
            ]
        }
    },
    "additionalProperties": true
}`
