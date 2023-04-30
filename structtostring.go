package gotools
func StructToString(data interface{}, delimiter string) (string, error) {
    value := reflect.ValueOf(data)
    if value.Kind() != reflect.Struct {
        return "", fmt.Errorf("expected struct, got %v", value.Kind())
    }

    var buffer bytes.Buffer
    t := value.Type()
    numFields := value.NumField()

    for i := 0; i < numFields; i++ {
        if i > 0 {
            buffer.WriteString(delimiter)
        }

        fieldName := t.Field(i).Name
        fieldValue := fmt.Sprintf("%v", value.Field(i).Interface())

        buffer.WriteString(fieldName)
        buffer.WriteString(":")
        buffer.WriteString(fieldValue)
    }

    return buffer.String(), nil
}
