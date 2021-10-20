package main

func parseModelToNonDuplicatedContents(value interface{}) (map[string]*TEnum, map[string]*TStruct, map[string]*TException) {
	enums := make(map[string]*TEnum)
	structs := make(map[string]*TStruct)
	exceptions := make(map[string]*TException)
	switch content := value.(type) {
	case *TService:
		for _, m := range content.Methods {
			es, ss, e := parseModelToNonDuplicatedContents(m)
			updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
		}
	case *TStruct:
		structs[content.Name] = content
		for _, f := range content.Fields {
			es, ss, e := parseModelToNonDuplicatedContents(f)
			updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
		}
	case *TEnum:
		enums[content.Name] = content
	case *TMap:
		{
			es, ss, e := parseModelToNonDuplicatedContents(content.KeyType)
			updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
		}
		{
			es, ss, e := parseModelToNonDuplicatedContents(content.ValueType)
			updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
		}
	case *TList:
		es, ss, e := parseModelToNonDuplicatedContents(content.Type)
		updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
	case *TSet:
		es, ss, e := parseModelToNonDuplicatedContents(content.Type)
		updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
	case *TField:
		es, ss, e := parseModelToNonDuplicatedContents(content.Type)
		updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
	case *TMethod:
		{
			es, ss, e := parseModelToNonDuplicatedContents(content.Response)
			updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
		}
		{
			es, ss, e := parseModelToNonDuplicatedContents(content.Exception)
			updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
		}
		{
			for _, f := range content.Parameters {
				es, ss, e := parseModelToNonDuplicatedContents(f)
				updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
			}
		}

	case *TException:
		exceptions[content.Name] = content
		for _, f := range content.Fields {
			es, ss, e := parseModelToNonDuplicatedContents(f)
			updateEnumAndStructs(enums, structs, exceptions, es, ss, e)
		}
	}
	return enums, structs, exceptions
}

func updateEnumAndStructs(basee map[string]*TEnum, bases map[string]*TStruct, baseex map[string]*TException, newe map[string]*TEnum, news map[string]*TStruct, newex map[string]*TException) {
	for k, v := range newe {
		basee[k] = v
	}
	for k, v := range news {
		bases[k] = v
	}
	for k, v := range newex {
		baseex[k] = v
	}
}

func FormatThriftService(service *TService) string {
	enums, structs, exceptions := parseModelToNonDuplicatedContents(service)
	var txt string
	for _, enum := range enums {
		txt += enum.TString() + "\n\n"
	}
	for _, s := range structs {
		txt += s.TString() + "\n\n"
	}
	for _, s := range exceptions {
		txt += s.TString() + "\n\n"
	}
	return txt + service.TString()
}
