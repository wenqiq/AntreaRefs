 ```yaml
                   properties:
                        matchExpressions:
                          type: array
                          items:
                            type: object
                            properties:
                              key:
                                type: string
                              operator:
                                enum:
                                  - In
                                  - NotIn
                                  - Exists
                                  - DoesNotExist
                                type: string
                              values:
                                type: array
                                items:
                                  type: string
                        matchLabels:
                          type: object
                  type: object


                   properties:
                     matchExpressions:
                       type: array
                       items:
                         type: object
                         properties:
                           key:
                             type: string
                           operator:
                             enum:
                               - In
                               - NotIn
                               - Exists
                               - DoesNotExist
                             type: string
                           values:
                             type: array
                             items:
                               type: string
                     matchLabels:
                       type: object
                   type: object
```