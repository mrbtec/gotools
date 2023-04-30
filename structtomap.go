package gotools
func StructToMap(i interface{}) (map[string]interface{}, error) {
    t := reflect.TypeOf(i)
    v := reflect.ValueOf(i)

    if v.Kind() == reflect.Ptr {
        v = v.Elem()
        t = t.Elem()
    }

    if v.Kind() != reflect.Struct {
        return nil, fmt.Errorf("input is not a struct")
    }

    var data = make(map[string]interface{})

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)

        // Ignora campos não exportados
        if field.PkgPath != "" {
            continue
        }

        data[field.Name] = value.Interface()
    }

    return data, nil
}
