// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"lybbrio/internal/ent/schema","Package":"lybbrio/internal/ent","Schemas":[{"name":"Author","config":{"Table":""},"edges":[{"name":"books","type":"Book","annotations":{"EntGQL":{"OrderField":"BOOKS_COUNT","RelayConnection":true}}}],"fields":[{"name":"calibre_id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"sort","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"link","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["calibre_id"]},{"unique":true,"fields":["name"]},{"unique":true,"fields":["sort"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"atr"}}},{"name":"Book","config":{"Table":""},"edges":[{"name":"authors","type":"Author","ref_name":"books","inverse":true},{"name":"publisher","type":"Publisher","ref_name":"books","inverse":true},{"name":"series","type":"Series","ref_name":"books","inverse":true},{"name":"identifiers","type":"Identifier","ref_name":"book","inverse":true},{"name":"tags","type":"Tag","ref_name":"books","inverse":true},{"name":"language","type":"Language","ref_name":"books","inverse":true},{"name":"shelf","type":"Shelf","ref_name":"books","inverse":true}],"fields":[{"name":"calibre_id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"title","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TITLE"}}},{"name":"sort","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"published_date","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"PUB_DATE"}}},{"name":"path","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}},{"name":"isbn","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"ISBN"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0}},{"name":"series_index","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["calibre_id"]},{"fields":["title"]},{"fields":["sort"]},{"fields":["published_date"]},{"fields":["isbn"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"bok"}}},{"name":"BookFile","config":{"Table":""},"edges":[{"name":"book","type":"Book","unique":true,"required":true}],"fields":[{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"path","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"size","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"comment":"Size in bytes"},{"name":"format","type":{"Type":6,"Ident":"bookfile.Format","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"enums":[{"N":"AZW3","V":"AZW3"},{"N":"EPUB","V":"EPUB"},{"N":"KEPUB","V":"KEPUB"},{"N":"PDF","V":"PDF"},{"N":"CBC","V":"CBC"},{"N":"CBR","V":"CBR"},{"N":"CB7","V":"CB7"},{"N":"CBZ","V":"CBZ"},{"N":"CBT","V":"CBT"}],"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"QueryField":{}},"KSUID":{"Prefix":"fil"}}},{"name":"Identifier","config":{"Table":""},"edges":[{"name":"book","type":"Book","unique":true,"required":true}],"fields":[{"name":"calibre_id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"type","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}}},{"name":"value","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"VALUE"}}}],"indexes":[{"unique":true,"fields":["calibre_id"]},{"fields":["type","value"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"idn"}}},{"name":"Language","config":{"Table":""},"edges":[{"name":"books","type":"Book","annotations":{"EntGQL":{"OrderField":"BOOKS_COUNT","RelayConnection":true}}}],"fields":[{"name":"calibre_id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"code","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"unique":true,"fields":["calibre_id"]},{"fields":["code"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"lng"}}},{"name":"Publisher","config":{"Table":""},"edges":[{"name":"books","type":"Book","annotations":{"EntGQL":{"OrderField":"BOOKS_COUNT","RelayConnection":true}}}],"fields":[{"name":"calibre_id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}}],"indexes":[{"unique":true,"fields":["calibre_id"]},{"fields":["name"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"pub"}}},{"name":"Series","config":{"Table":""},"edges":[{"name":"books","type":"Book"}],"fields":[{"name":"calibre_id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"sort","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}}],"indexes":[{"unique":true,"fields":["calibre_id"]},{"fields":["name"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"srs"}}},{"name":"Shelf","config":{"Table":""},"edges":[{"name":"user","type":"User","field":"user_id","unique":true,"required":true,"immutable":true},{"name":"books","type":"Book","annotations":{"EntGQL":{"OrderField":"BOOKS_COUNT","RelayConnection":true}}}],"fields":[{"name":"public","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"user_id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}},{"name":"description","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["name"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":true,"MixinIndex":1},{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"MutationInputs":[{}],"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"shf"}}},{"name":"Tag","config":{"Table":""},"edges":[{"name":"books","type":"Book","annotations":{"EntGQL":{"OrderField":"BOOKS_COUNT","RelayConnection":true}}}],"fields":[{"name":"calibre_id","type":{"Type":13,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"NAME"}}}],"indexes":[{"unique":true,"fields":["calibre_id"]},{"fields":["name"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"MutationInputs":[{"IsCreate":true},{}],"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"tag"}}},{"name":"Task","config":{"Table":""},"edges":[{"name":"user","type":"User","field":"user_id","unique":true,"immutable":true}],"fields":[{"name":"create_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"update_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":1}},{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":2}},{"name":"type","type":{"Type":6,"Ident":"task_enums.TaskType","PkgPath":"lybbrio/internal/ent/schema/task_enums","PkgName":"task_enums","Nillable":false,"RType":{"Name":"TaskType","Ident":"task_enums.TaskType","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/task_enums","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Values":{"In":[],"Out":[{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":null}]}}}},"enums":[{"N":"noop","V":"noop"},{"N":"calibre_import","V":"calibre_import"}],"default":true,"default_value":"noop","default_kind":24,"immutable":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"TYPE"}}},{"name":"status","type":{"Type":6,"Ident":"task_enums.Status","PkgPath":"lybbrio/internal/ent/schema/task_enums","PkgName":"task_enums","Nillable":false,"RType":{"Name":"Status","Ident":"task_enums.Status","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/task_enums","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Values":{"In":[],"Out":[{"Name":"","Ident":"[]string","Kind":23,"PkgPath":"","Methods":null}]}}}},"enums":[{"N":"pending","V":"pending"},{"N":"in_progress","V":"in_progress"},{"N":"success","V":"success"},{"N":"failure","V":"failure"}],"default":true,"default_value":"pending","default_kind":24,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"STATUS"}}},{"name":"progress","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":0,"default_kind":14,"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"comment":"Progress of the task. 0-1"},{"name":"message","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"comment":"Message of the task"},{"name":"error","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"comment":"Error message of the task"},{"name":"user_id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"optional":true,"immutable":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"comment":"The user who created this task. Empty for System Task"},{"name":"is_system_task","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"comment":"Whether this task is created by the system"}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"MultiOrder":true,"QueryField":{},"RelayConnection":true},"KSUID":{"Prefix":"tsk"}}},{"name":"User","config":{"Table":""},"edges":[{"name":"shelves","type":"Shelf","ref_name":"user","inverse":true},{"name":"user_permissions","type":"UserPermissions","unique":true,"required":true,"immutable":true}],"fields":[{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"username","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntGQL":{"OrderField":"USERNAME"}}},{"name":"password_hash","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"sensitive":true},{"name":"email","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"unique":true,"validators":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}}],"indexes":[{"fields":["username"]},{"fields":["password_hash"]}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"EntGQL":{"MutationInputs":[{},{"IsCreate":true}],"QueryField":{}},"KSUID":{"Prefix":"usr"}}},{"name":"UserPermissions","config":{"Table":""},"edges":[{"name":"user","type":"User","field":"user_id","ref_name":"user_permissions","unique":true,"inverse":true}],"fields":[{"name":"id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"default":true,"default_kind":19,"position":{"Index":0,"MixedIn":true,"MixinIndex":1}},{"name":"user_id","type":{"Type":7,"Ident":"ksuid.ID","PkgPath":"lybbrio/internal/ent/schema/ksuid","PkgName":"ksuid","Nillable":false,"RType":{"Name":"ID","Ident":"ksuid.ID","Kind":24,"PkgPath":"lybbrio/internal/ent/schema/ksuid","Methods":{"MarshalGQL":{"In":[{"Name":"Writer","Ident":"io.Writer","Kind":20,"PkgPath":"io","Methods":null}],"Out":[]},"Scan":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"String":{"In":[],"Out":[{"Name":"string","Ident":"string","Kind":24,"PkgPath":"","Methods":null}]},"UnmarshalGQL":{"In":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null}],"Out":[{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]},"Value":{"In":[],"Out":[{"Name":"","Ident":"interface {}","Kind":20,"PkgPath":"","Methods":null},{"Name":"error","Ident":"error","Kind":20,"PkgPath":"","Methods":null}]}}}},"optional":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0}},{"name":"Admin","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":1,"MixedIn":false,"MixinIndex":0}},{"name":"CanCreatePublic","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":2,"MixedIn":false,"MixinIndex":0}},{"name":"CanEdit","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"default":true,"default_value":false,"default_kind":1,"position":{"Index":3,"MixedIn":false,"MixinIndex":0}}],"policy":[{"Index":0,"MixedIn":true,"MixinIndex":0},{"Index":0,"MixedIn":false,"MixinIndex":0}],"annotations":{"KSUID":{"Prefix":"prm"}}}],"Features":["entql","privacy","schema/snapshot","sql/upsert","namedges"]}`
